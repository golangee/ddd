package ddd

import "github.com/golangee/ddd/sql"

func MySQL(migrations []*sql.MigrationSpec, structs []*StructSpec, genSpecs []*GenSpec) *LayerSpec {
	return nil
}

type GenSpec struct {
}

func Repository(name, comment string, funcs ...*GenFuncSpec) *GenSpec {
	return &GenSpec{}
}

func Generate(genSpecs ...*GenSpec) []*GenSpec {
	return genSpecs
}

type GenFuncSpec struct {
}

func StatementFunc(name, comment, statement string, in []*ParamSpec, out []*ParamSpec) *GenFuncSpec {
	return &GenFuncSpec{}
}
