package ast

// A Pos is an arbitrary position handle without any meaning.
type Pos int

// Position describes a resolved position within a file.
type Position struct {
	// File contains the absolute file path.
	File string
	// Line denotes the one-based line number in the denoted File.
	Line int
	// Col denotes the one-based column number in the denoted Line.
	Col int
}

// A PositionSet resolves a Pos to a Position.
type PositionSet struct {
	pos map[Pos]Position
}

// Position returns the Position value for the given file position pos.
func (p *PositionSet) Position(pos Pos) Position {
	if p == nil || p.pos == nil {
		return Position{}
	}

	return p.pos[pos]
}

// Add appends the given position and returns a new Pos handle.
func (p *PositionSet) Add(pos Position) Pos {
	if p.pos == nil {
		p.pos = map[Pos]Position{}
	}

	next := Pos(len(p.pos) + 1)
	p.pos[next] = pos

	return next
}

// A Stereotype is used to declare a meta class for a node.
type Stereotype string

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
