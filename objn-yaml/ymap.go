package objnyaml

import (
	"fmt"
	"github.com/golangee/architecture/objn"
	"gopkg.in/yaml.v3"
	"sort"
)

type YamlMap struct {
	pos objn.Pos
	m   *yaml.Node
}

func NewYamlMap(filename string, m *yaml.Node) *YamlMap {
	return &YamlMap{m: m, pos: objn.Pos{
		File: filename,
		Line: m.Line,
		Col:  m.Column,
	}}
}

func (n *YamlMap) Pos() objn.Pos {
	return n.pos
}

func (n *YamlMap) Count() int {
	return len(n.m.Content) / 2
}

func (n *YamlMap) Comment() string {
	return n.m.HeadComment
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
				// pick all duplicates
				var details []objn.PosErrorDetail
				names :=n.Names()
				for _, lit := range names {
					if lit.String() == s{
						details = append(details, objn.NewPosErrorDetailFromDoc())
					}
				}
				return fmt.Errorf(n.pos.String()+": contains duplicate key '%s'", s)
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

func (n *YamlMap) Names() []objn.Lit {
	tmp := make([]objn.Lit, 0, n.Count())
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		tmp = append(tmp, n.newLit(node))
	}

	return tmp
}

func (n *YamlMap) Name(key string) objn.Lit {
	for i := 0; i < len(n.m.Content); i += 2 {
		node := n.m.Content[i]
		if node.Value == key {
			return n.newLit(node)
		}
	}

	return nil
}

func (n *YamlMap) Get(key string) objn.Node {
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
	return NewYamlLit(objn.Pos{
		File: n.pos.File,
		Line: node.Line,
		Col:  node.Column,
	}, node.Value, node.HeadComment)
}
