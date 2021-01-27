package validate

import "github.com/golangee/architecture/yast"

// A Validator is just a predicate-like func which returns an error instead of a bool.
type Validator func(node yast.Node) error

// All delegates its check call to all given validators and returns the first failure.
func All(checkers ...Validator) Validator {
	return func(node yast.Node) error {
		for _, checker := range checkers {
			if err := checker(node); err != nil {
				return err
			}
		}

		return nil
	}
}

// Recursive loops over all yast.Obj and yast.Seq values and applies the given validator on each.
func Recursive(v Validator) Validator {
	return func(node yast.Node) error {
		if err := v(node); err != nil {
			return err
		}

		if pnode, ok := node.(yast.Parent); ok {
			for _, n := range pnode.Children() {
				if err := Recursive(v)(n); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

// NoDuplicateKeys loops over all Obj Attr.Key values and returns a PosError if they are not unique. If node
// is not a *yast.Obj nothing happens.
func NoDuplicateKeys() Validator {
	return func(node yast.Node) error {
		if obj, ok := node.(*yast.Obj); ok {
			names := obj.Names()
			lastKey := ""
			foundDuplicate := false
			for _, s := range names {
				if s == lastKey {
					if s != "" {
						foundDuplicate = true
						break
					}
				}

				lastKey = s
			}

			// that is our duplicate key
			if foundDuplicate {
				allDuplicates := make([]*yast.Attr, 0)
				for _, attr := range obj.Attrs {
					if attr.Key.Value == lastKey {
						allDuplicates = append(allDuplicates, attr)
					}
				}

				details := make([]yast.ErrDetail, 0)
				for i := 1; i < len(allDuplicates); i++ {
					details = append(details, yast.ErrDetail{
						Node:    allDuplicates[i].Key,
						Message: "key also defined here",
					})
				}

				posErr := yast.NewPosError(allDuplicates[0], "duplicate key '"+lastKey+"'", details...)
				posErr.Hint = "duplicate keys are not allowed, so remove them."

				return posErr
			}
		}

		return nil
	}
}
