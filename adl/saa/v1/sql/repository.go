package sql

import "github.com/golangee/architecture/adl/saa/v1/core"

// A RepoImplSpec declares a repository and methods with data bindings.
type Repository struct {
	Implements core.TypeLit // full qualified name of the implementing interface.
	Methods    []Method     // subset of methods which should be implemented automatically.
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context.
type Method struct {
	Name    core.StrLit // Name must match an interface method of Repository.Implements
	Query   core.StrLit // Query is a literal (usually with placeholders) to execute or query a result set for.
	Prepare []Mapping   // either a parameter name or a parameter name with dot selectors.
	Result  []Mapping   // either a result name, a result type name or a parameter name with dot selectors.
}

// Mapping of Prepare and Result
type Mapping interface {
	mappingType()
}

// MapSelOne maps to a parameter name (either in or out). The natural index of this mapping
// determines its order - either as a placeholder or as a result column.
type MapSelOne struct {
	Sel core.StrLit // may contain dots to select a field, so either "myParam" or "myParam.Field"
}

func (m MapSelOne) mappingType() {

}

// MapSelMany maps to a slice of structs or primitives. The natural index of this mapping
// determines its order - either as a placeholder or as a result column.
type MapSelMany struct {
	Sel []core.StrLit // must be of the form "." or ".Field" and must address each entry of the single parameter
}

func (m MapSelMany) mappingType() {

}

// MapPrimResult maps to specific primitive type and is only allowed as a Result.
type MapPrimResult struct {
	Type core.TypeLit // The full qualified primitive type name. This can also be an alias.
}

func (m MapPrimResult) mappingType() {

}
