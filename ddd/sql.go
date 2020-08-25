package ddd

import "github.com/golangee/architecture/sql"

func MySQL(migrations []*sql.MigrationSpec,  genSpecs []*GenSpec) *UseCaseLayerSpec {
	return nil
}

type GenSpec struct {
}

func From(typeName TypeName, funcs ...*GenFuncSpec) *GenSpec {
	return &GenSpec{}
}

func Generate(genSpecs ...*GenSpec) []*GenSpec {
	return genSpecs
}

type GenFuncSpec struct {
}

func StatementFunc(name, statement string) *GenFuncSpec {
	return &GenFuncSpec{}
}
