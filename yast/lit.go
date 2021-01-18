package yast

// A Str represents an already parsed string literal.
type Str struct {
	ValuePos        Pos
	ValueEnd        Pos
	ValueParent     Node
	ValueStereotype string
	Value           string
}

func (a *Str) Pos() Pos {
	return a.ValuePos
}

func (a *Str) End() Pos {
	return a.ValueEnd
}

func (a *Str) Parent() Node {
	return a.ValueParent
}

func (a *Str) Stereotype() string {
	return a.ValueStereotype
}

func (a *Str) String() string {
	return a.Value
}

// An Int represents an already parsed int64 literal.
type Int struct {
	ValuePos        Pos
	ValueEnd        Pos
	ValueParent     Node
	ValueStereotype string
	Value           int64
}

func (a *Int) Pos() Pos {
	return a.ValuePos
}

func (a *Int) End() Pos {
	return a.ValueEnd
}

func (a *Int) Parent() Node {
	return a.ValueParent
}

func (a *Int) Stereotype() string {
	return a.ValueStereotype
}

// A Float represents an already parsed float64 literal.
type Float struct {
	ValuePos        Pos
	ValueEnd        Pos
	ValueParent     Node
	ValueStereotype string
	Value           float64
}

func (a *Float) Pos() Pos {
	return a.ValuePos
}

func (a *Float) End() Pos {
	return a.ValueEnd
}

func (a *Float) Parent() Node {
	return a.ValueParent
}

func (a *Float) Stereotype() string {
	return a.ValueStereotype
}

// A Bool represents an already parsed bool literal.
type Bool struct {
	ValuePos        Pos
	ValueEnd        Pos
	ValueParent     Node
	ValueStereotype string
	Value           bool
}

func (a *Bool) Pos() Pos {
	return a.ValuePos
}

func (a *Bool) End() Pos {
	return a.ValueEnd
}

func (a *Bool) Parent() Node {
	return a.ValueParent
}

func (a *Bool) Stereotype() string {
	return a.ValueStereotype
}

// A Null represents a void resp. undefined value.
type Null struct {
	ValuePos        Pos
	ValueEnd        Pos
	ValueParent     Node
	ValueStereotype string
}

func (a *Null) Pos() Pos {
	return a.ValuePos
}

func (a *Null) End() Pos {
	return a.ValueEnd
}

func (a *Null) Parent() Node {
	return a.ValueParent
}

func (a *Null) Stereotype() string {
	return a.ValueStereotype
}
