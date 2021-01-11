package objnyaml

import (
	"github.com/golangee/architecture/objn"
	"gopkg.in/yaml.v3"
)

type YamlSeq struct {
	pos  objn.Pos
	node *yaml.Node
}

func NewYamlSeq(filename string, node *yaml.Node) *YamlSeq {
	return &YamlSeq{node: node, pos: objn.Pos{
		File: filename,
		Line: node.Line,
		Col:  node.Column,
	}}
}

func (n *YamlSeq) Validate() error {
	for _, node := range n.node.Content {
		wrapped := NewNode(n.pos.File, node)
		if v, ok := wrapped.(validateable); ok {
			if err := v.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *YamlSeq) Comment() string {
	return n.node.HeadComment
}

func (n *YamlSeq) Pos() objn.Pos {
	return n.pos
}

func (n *YamlSeq) Count() int {
	return len(n.node.Content)
}

func (n *YamlSeq) Get(idx int) objn.Node {
	return NewNode(n.pos.File, n.node.Content[idx])
}
