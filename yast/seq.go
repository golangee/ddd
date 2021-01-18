package yast

// A Seq represents a sequence of arbitrary nodes.
type Seq struct {
	SeqPos        Pos
	SeqEnd        Pos
	SeqParent     Node
	SeqStereotype string
	Values        []Node
}

func (s *Seq) Pos() Pos {
	return s.SeqPos
}

func (s *Seq) End() Pos {
	return s.SeqEnd
}

func (s *Seq) Parent() Node {
	return s.SeqParent
}

func (s *Seq) Stereotype() string {
	return s.SeqStereotype
}

func (s *Seq) Children() []Node {
	tmp := make([]Node, 0, len(s.Values))
	for _, node := range tmp {
		tmp = append(tmp, node)
	}

	return tmp
}
