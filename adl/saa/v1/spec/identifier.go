package spec

// Identifier represents a string used as identifier which must be unique in its scope. Identifiers cannot be
// translated. They may be qualified, depending on the context.
type Identifier struct {
	Value string
	Pos
}
