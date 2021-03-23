package core

// Visibility indicates the visibility of a type, field or method.
type Visibility int

const (
	Public Visibility = iota
	PackagePrivate
)

// Identifier represents a string used as identifier which must be unique in its scope. Identifiers cannot be
// translated. An Identifier is not qualified.
type Identifier struct {
	// Visibility of this identifier. This looks weired, but this makes life easier with Go.
	// Actually connecting the visibility with the identifier is not so wrong as it looks at first. Imagine
	// a package private implementation which is exported using a public interface. So the fact of exporting
	// a type is really an identifying attribute and not a type attribute in itself.
	Visibility Visibility

	// Value is the original name literal.
	Value string

	// ValueTransformer renames the Identifier into something else, e.g. to fit Go notation or hungarian notations like in AOSP.
	ValueTransformer func(ident Identifier) string

	node
}

func NewIdent(str string) Identifier {
	return Identifier{Value: str}
}

// String returns either the original or a transformed Value.
func (v Identifier) String() string {
	if v.ValueTransformer == nil {
		return v.Value
	}

	return v.ValueTransformer(v)
}
