package validate

import (
	"github.com/golangee/architecture/yast"
	"reflect"
	"strings"
)

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

// ExpectObj splits the given path by slashes (e.g. a/b/c) and checks that c is a yast.Node.
func ExpectNode(dst *yast.Node, srcPath, expectedType string) Validator {
	return func(src yast.Node) error {
		names := strings.Split(srcPath, "/")
		root := src
		for _, name := range names {
			if obj, ok := root.(*yast.Obj); ok {
				val := obj.Get(name)
				if val == nil {
					posErr := yast.NewPosError(root, "expected element '"+name+"' but is missing")
					posErr.Hint = "add an object identified by key '" + name + "'"
					return posErr
				}

				root = val
			} else {
				posErr := yast.NewPosError(root, "expected element '"+name+"' to be a '"+expectedType+"' but found: "+reflect.TypeOf(root).String())
				posErr.Hint = "change attribute '" + name + "' to refer to '" + expectedType + "'"
				return posErr
			}
		}

		*dst = root
		return nil
	}
}


// ExpectObj splits the given path by slashes (e.g. a/b/c) and checks that c is a yast.Obj.
func ExpectObj(dst **yast.Obj, srcPath string) Validator {
	return func(node yast.Node) error {
		var tmp yast.Node
		err := ExpectNode(&tmp, srcPath, "Object")(node)
		if err == nil {
			*dst = tmp.(*yast.Obj)
		}

		return err
	}
}

// ExpectStr splits the given path by slashes (e.g. a/b/c) and checks that c is a yast.Str.
func ExpectStr(dst **yast.Str, srcPath string) Validator {
	return func(node yast.Node) error {
		var tmp yast.Node
		err := ExpectNode(&tmp, srcPath, "Str")(node)
		if err == nil {
			*dst = tmp.(*yast.Str)
		}

		return err
	}
}

