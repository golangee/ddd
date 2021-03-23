package metamodel

// StrLit represents a string used as a literal. Literals have context specific meaning and cannot be translated,
// like a constant or query expression.
type StrLit struct {
	Value string
	Pos
}

