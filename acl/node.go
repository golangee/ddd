package acl

import "strconv"

// Pos describes a position within a file.
type Pos struct {
	// File contains the absolute file path.
	File string
	// Line denotes the one-based line number in the denoted File.
	Line int
	// Col denotes the one-based column number in the denoted Line.
	Col int
}

func (p Pos) String() string {
	return p.File + ":" + strconv.Itoa(p.Line) + ":" + strconv.Itoa(p.Col)
}

// Node represents the most abstract element of a configuration file.
type Node interface {
	// Pos returns the actual starting position of this Node.
	Pos() Pos
}

// Map provides access to an Object or map-like data structure.
type Map interface {
	Node

	// Count of the declared attributes.
	Count() int

	// Names of the unique attributes returns all key literals.
	Names() []Lit

	// Name returns the according literal for the given attribute key.
	Name(key string) Lit

	// Get returns the value Node which may be a Lit, a Seq or another Map.
	Get(key string) Node
}

// Seq defines a list or array-like object with an ordered access by a zero-based index.
type Seq interface {
	Node

	// Count of the contained entries.
	Count() int

	// Get returns the Lit, Seq or Map at the given index position. If idx is < 0 or >= Count a panic is raised.
	Get(idx int) Node
}

// A Lit defines an atomic value like a string, float or integer type.
type Lit interface {
	Node

	// String returns the uninterpreted natural literal (which is always a string in text-based configurations).
	// If the underlying format is something else, like a binary serialization, a string-marshalled representation
	// is returned.
	String() string
}

// A Doc contains the actual Seq, Map or Lit configuration nodes.
type Doc interface {
	Node

	// String returns a textual (marshalled) representation of the documents contents which matches each contained Pos.
	String() string

	// Root returns the documents root node.
	Root() Node
}

// Pkg represents a named package in which documents and other packages may have been declared.
type Pkg interface {
	Node

	// Names of the uniquely named nodes within this package.
	Names() []Lit

	// Name returns the according literal for the given Pkg or Doc.
	Name(key string) Lit

	// Get returns the Pkg or a Doc at the given index position. If idx is < 0 or >= Count a panic is raised.
	Get(key string) Node
}
