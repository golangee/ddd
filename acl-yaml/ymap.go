package aclyaml

import (
	"fmt"
	"github.com/golangee/architecture/acl"
	"gopkg.in/yaml.v3"
	"sort"
)

type YamlMap struct {
	pos acl.Pos
	m   *yaml.Node
}

func NewYamlMap(filename string, m *yaml.Node) *YamlMap {
	return &YamlMap{m: m, pos: acl.Pos{
		File: filename,
		Line: m.Line,
		Col:  m.Column,
	}}
}

func (n *YamlMap) Pos() acl.Pos {
	return n.pos
}

func (n *YamlMap) Count() int {
	return len(n.m.Content) / 2
}

// Validate checks if each key is unique.
func (n *YamlMap) Validate() error {
	tmp := make([]string, 0, n.Count())
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		tmp = append(tmp, node.Value)
	}

	sort.Strings(tmp)

	lastKey := ""
	for _, s := range tmp {
		if s == lastKey {
			if s != "" {
				return fmt.Errorf(n.pos.String() + ": contains duplicate key '%s'",s)
			}
		}

		lastKey = s
	}

	for i := 1; i < len(n.m.Content); i += 2 {
		wrapped := NewNode(n.pos.File, n.m.Content[i])
		if v, ok := wrapped.(validateable); ok {
			if err := v.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *YamlMap) Names() []acl.Lit {
	tmp := make([]acl.Lit, 0, n.Count())
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		tmp = append(tmp, n.newLit(node))
	}

	return tmp
}

func (n *YamlMap) Name(key string) acl.Lit {
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		if node.Value == key {
			return n.newLit(node)
		}
	}

	return nil
}

func (n *YamlMap) Get(key string) acl.Node {
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		if node.Value == key {
			i++
			node = n.m.Content[i]
			return NewNode(n.pos.File, node)
		}
	}

	return nil
}

func (n *YamlMap) newLit(node *yaml.Node) *YamlLit {
	return NewYamlLit(acl.Pos{
		File: n.pos.File,
		Line: node.Line,
		Col:  node.Column,
	}, node.Value)
}
