package aclyaml

import (
	"github.com/golangee/architecture/acl"
	"gopkg.in/yaml.v3"
)

type YamlSeq struct {
	pos  acl.Pos
	node *yaml.Node
}

func NewYamlSeq(filename string, node *yaml.Node) *YamlSeq {
	return &YamlSeq{node: node, pos: acl.Pos{
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

func (n *YamlSeq) Pos() acl.Pos {
	return n.pos
}

func (n *YamlSeq) Count() int {
	return len(n.node.Content)
}

func (n *YamlSeq) Get(idx int) acl.Node {
	return NewNode(n.pos.File, n.node.Content[idx])
}
