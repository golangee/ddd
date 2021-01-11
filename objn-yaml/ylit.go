package objnyaml

import (
	"github.com/golangee/architecture/objn"
)

type YamlLit struct {
	pos     objn.Pos
	val     string
	comment string
}

func NewYamlLit(pos objn.Pos, val, comment string) *YamlLit {
	return &YamlLit{pos: pos, val: val, comment: comment}
}

func (y YamlLit) Comment() string {
	return y.comment
}

func (y YamlLit) Pos() objn.Pos {
	return y.pos
}

func (y YamlLit) String() string {
	return y.val
}
