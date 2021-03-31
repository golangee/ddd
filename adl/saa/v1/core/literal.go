package core

// StrLit represents a string used as a literal. Literals have context specific meaning and cannot be translated,
// like a constant or query expression.
type StrLit struct {
	Value string
	node
}

func NewStrLit(str string) StrLit {
	return StrLit{Value: str}
}

func (s StrLit) String() string {
	return s.Value
}

// ModLit is just like a StrLit but with a different semantic.
type ModLit StrLit

func NewModLit(str string) ModLit {
	return ModLit{Value: str}
}

func (s ModLit) String() string {
	return s.Value
}

// PkgLit is just like a StrLit but with a different semantic.
type PkgLit StrLit

func NewPkgLit(str string) PkgLit {
	return PkgLit{Value: str}
}

func (s PkgLit) String() string {
	return s.Value
}

// TypeLit is just like a StrLit but with a different semantic.
type TypeLit StrLit

func NewTypeLit(str string) TypeLit {
	return TypeLit{Value: str}
}

func (s TypeLit) String() string {
	return s.Value
}
