package metamodel

// Path represents a string which is separated by /.
type Path struct {
	// Value is the original path literal.
	Value string

	// ValueTransformer renames the Path into something else, e.g. to make lower-case and dots out of it.
	ValueTransformer func(ident Path) string

	// Pos denotes the original position of the Identifier declaration.
	Pos Pos
}

// String returns either the original or a transformed Value.
func (v Path) String() string {
	if v.ValueTransformer == nil {
		return v.Value
	}

	return v.ValueTransformer(v)
}
