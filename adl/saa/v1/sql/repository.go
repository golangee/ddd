package sql

import "github.com/golangee/architecture/adl/saa/v1/core"

// A RepoImplSpec declares a repository and methods with data bindings.
type Repository struct {
	Implements core.TypeLit // full qualified name of the implementing interface.
	Methods    []Method     // subset of methods which should be implemented automatically.
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context. The actual types are read from the according ast.Func.
type Method struct {
	Name    core.StrLit // Name must match an interface method of Repository.Implements
	Query   core.StrLit // Query is a literal (usually with placeholders) to execute or query a result set for.
	Mapping Mapping     // Mapping describes what how the sql prepare and scan generator should work. However, multiple returns are not allowed (special to Go only)
}

// Mapping of Prepare and Result
type Mapping interface {
	mappingType()
}

// ExecOne maps 0, 1 or many input parameters defined as selectors into a prepared sql statement.
// There is no result.
type ExecOne struct {
	In []core.StrLit // may contain dots to select a parameter or a parameter field, so either "myParam" or "myParam.Field"
}

func (_ ExecOne) mappingType() {

}

// ExecMany maps a single input slice (or list) parameter into a repeated prepared sql statement.
// There may be other non-array input parameters. This is optimized for high load batch inserts.
// There is no result.
type ExecMany struct {
	// Slice determines the range/list/count target and determines the index 'i' for the declared
	// in parameters.
	Slice core.StrLit

	// This can be a mixture of
	//  myparam
	//  myparam.field
	//  myparam[i]
	//  myparam.field[i]
	//  ...
	// i represents the loop index.
	In []core.StrLit
}

func (_ ExecMany) mappingType() {

}

// ExecOne maps 0, 1 or many input parameters defined as selectors into a prepared sql statement.
//
// The result is either a single struct or a single primitive. The type is determined from the ast.Func.
type QueryOne struct {
	In  []core.StrLit // may contain dots to select a parameter or a parameter field, so either "myParam" or "myParam.Field"
	Out []core.StrLit // may contain dots to select a field or . for the primitive itself
}

func (_ QueryOne) mappingType() {

}

// ExecMany maps 0, 1 or many input parameters defined as selectors into a prepared sql statement.
//
// The result is a slice of either structs or primitives.
type QueryMany struct {
	In  []core.StrLit // may contain dots to select a parameter or a parameter field, so either "myParam" or "myParam.Field"
	Out []core.StrLit // may contain dots to select a field or . for the primitive itself
}

func (_ QueryMany) mappingType() {

}
