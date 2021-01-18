package yast

import "strconv"

// A Stereotype is used to declare a meta class for a node.
type Stereotype = string

const (
	Document  Stereotype = "document"
	Package              = "package"
	File                 = "file"
	Directory            = "directory"
	Source               = "source"
)

// A Pos describes a resolved position within a file.
type Pos struct {
	// File contains the absolute file path.
	File string
	// Line denotes the one-based line number in the denoted File.
	Line int
	// Col denotes the one-based column number in the denoted Line.
	Col int
}

// String returns the content in the "file:line:col" format.
func (p Pos) String() string {
	return p.File + ":" + strconv.Itoa(p.Line) + ":" + strconv.Itoa(p.Col)
}

// A Node represents the common contract
type Node interface {
	// Pos returns the actual starting position of this Node.
	Pos() Pos

	// End is the position of the first char after the node.
	End() Pos

	// Parent returns the parent Node or nil if undefined. This recursive implementation may be considered as
	// unnecessary and even as an anti pattern within an AST but the core feature is to perform semantic validations
	// which requires a lot of down/up iterations through the (entire) AST. Keeping the relational relation
	// at the node level keeps things simple and we don't need to pass (path) contexts everywhere.
	Parent() Node

	// Stereotype returns a declared meta class.
	Stereotype() string
}

// A Parent is a Node and may contain other nodes as children. This is used to simplify algorithms based on Walk.
type Parent interface {
	Node
	Children() []Node // Children returns a defensive copy of the underlying slice. However the Node references are shared.
}
