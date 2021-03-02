package types

import "github.com/golangee/architecture/adl/saa/v1/spec"

type Interface struct {
	Name    spec.Identifier
	Comment spec.String
	Methods []Method
}

type Method struct {
	Name    spec.Identifier
	Comment spec.String
	In      []Param
	Out     []Param
}

type Param struct {
	Comment spec.String
	Name    spec.Identifier
	Type    spec.StrLit
}
