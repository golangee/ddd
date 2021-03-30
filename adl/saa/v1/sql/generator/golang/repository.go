package golang

import (
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
)

// renderStaticFindAll performs a simple selection using the declared sql stereotypes from
// entity. The table name is read from the entity itself and all fields which 
// have the according field names are used for projection.
func renderStaticFindAll(entity *ast.Struct, dbType ast.Name, tableName string) (*ast.Func, error) {
	fun := ast.NewFunc("FindAll"+entity.TypeName).
		SetComment(
			"... reads the entire table '"+tableName+"' into memory.",
		).
		AddParams(
			ast.NewParam("db", ast.NewSimpleTypeDecl(dbType)),
		).
		AddResults(
			ast.NewParam("", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl(ast.Name(entity.TypeName)))),
			ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
		).
		SetBody(
			ast.NewBlock(
				//ast.
			),
		)

	return fun, nil
}
