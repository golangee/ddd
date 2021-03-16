package metamodel

import "strconv"

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
