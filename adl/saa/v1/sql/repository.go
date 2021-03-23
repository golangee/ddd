package sql

import "github.com/golangee/architecture/adl/saa/v1/core"

// A RepoImplSpec declares a repository and methods with data bindings.
type RepoImplSpec struct {
	Name       core.Identifier
	Methods    []*Method
	Migrations []*Migration
}

func NewRepoImplSpec(name core.Identifier) *RepoImplSpec {
	return &RepoImplSpec{Name: name}
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context.
type Method struct {
	Parent  *RepoImplSpec
	Name    core.Identifier
	Query   core.StrLit
	Prepare []core.Identifier
	Map     []core.Identifier
}
