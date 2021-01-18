package yast

// XNode can be used to express markup languages like xml. Note, that every xml node is an XNode and node types
// are represented using stereotypes. See also https://www.w3schools.com/xml/dom_nodetype.asp.
type XNode struct {
	ElemPos        Pos
	ElemEnd        Pos
	ElemStereotype string
	ElemParent     Node
	Space          string // optional name space identifier. Contains the canonical URL, not a short prefix.
	Attrs          []Attr
	Name           string
	Value          string
	Nodes          []XNode
}

func (a *XNode) Pos() Pos {
	return a.ElemPos
}

func (a *XNode) End() Pos {
	return a.ElemEnd
}

func (a *XNode) Parent() Node {
	return a.ElemParent
}

func (a *XNode) Stereotype() string {
	return a.ElemStereotype
}

func (a *XNode) Children() []Node {
	tmp := make([]Node, 0, len(a.Nodes))
	for _, node := range tmp {
		tmp = append(tmp, node)
	}

	return tmp
}
