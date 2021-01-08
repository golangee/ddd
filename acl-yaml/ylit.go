package aclyaml

import (
	"github.com/golangee/architecture/acl"
)

type YamlLit struct {
	pos acl.Pos
	val string
}

func NewYamlLit(pos acl.Pos, val string) *YamlLit {
	return &YamlLit{pos: pos, val: val}
}

func (y YamlLit) Pos() acl.Pos {
	return y.pos
}

func (y YamlLit) String() string {
	return y.val
}
