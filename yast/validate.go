package yast

// ValidateDuplicateKeys loops over all Obj Attr.Key values and returns a PosError if they are not unique.
func ValidateDuplicateKeys(node Node) error {
	if obj, ok := node.(*Obj); ok {
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
			allDuplicates := make([]*Attr, 0)
			for _, attr := range obj.Attrs {
				if attr.Key.Value == lastKey {
					allDuplicates = append(allDuplicates, attr)
				}
			}

			details := make([]ErrDetail, 0)
			for i := 1; i < len(allDuplicates); i++ {
				details = append(details, ErrDetail{
					Node:    allDuplicates[i].Key,
					Message: "key also defined here",
				})
			}

			posErr := NewPosError(allDuplicates[0], "duplicate key '"+lastKey+"'", details...)
			posErr.Hint = "duplicate keys are not allowed, so remove them."

			return posErr
		}
	}

	if pnode, ok := node.(Parent); ok {
		for _, n := range pnode.Children() {
			if err := ValidateDuplicateKeys(n); err != nil {
				return err
			}
		}
	}

	return nil
}
