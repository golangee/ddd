package golang

import (
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
)

// renderStaticFindAll performs a simple selection using the declared sql stereotypes from
// entity. The table name is read from the stereotype entity itself and all fields which
// have the according field names are used for projection. A default sort order is appended
// optionally.
func renderStaticFindAll(entity *ast.Struct, dbType ast.Name) (*ast.Func, error) {
	tableName, _ := stereotype.StructFrom(entity).SQLTableName()

	var scanArgs []ast.Expr

	selection := "SELECT "
	for _, f := range entity.Fields() {
		if col, ok := stereotype.FieldFrom(f).SQLColumnName(); ok {
			selection += col + ", "
			scanArgs = append(scanArgs, ast.NewUnaryExpr(lang.Sel("i", f.FieldName), ast.OpAnd))
		}
	}

	selection = selection[:len(selection)-2] + " FROM " + tableName

	if defaultOrder, ok := stereotype.StructFrom(entity).SQLDefaultOrder(); ok {
		selection += " " + defaultOrder
	}

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
				ast.NewConstDecl(ast.NewSimpleAssign(ast.NewIdent("q"), ast.AssignSimple, ast.NewStrLit(selection))),
				ast.NewVarDecl(ast.NewParam("res", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl(ast.Name(entity.TypeName))))),
				lang.TryDefine(
					ast.NewIdent("rows"),
					lang.CallIdent("db", "QueryContext", lang.CallStatic("context.Background"), ast.NewIdent("q")),
					"cannot query "+entity.TypeName,
				),
				ast.NewDeferStmt(lang.CallIdent("rows", "Close")),
				lang.Term(),
				ast.NewForStmt(nil, lang.CallIdent("rows", "Next"), nil,
					ast.NewBlock(
						ast.NewVarDecl(ast.NewParam("i", ast.NewSimpleTypeDecl(ast.Name(entity.TypeName)))),
						lang.TryDefine(nil, lang.CallIdent("rows", "Scan", scanArgs...), "cannot scan row result"),
						ast.NewAssign(ast.Exprs(ast.NewIdent("res")), ast.AssignSimple, ast.Exprs(lang.Call("append", ast.NewIdent("res"), ast.NewIdent("i")))),
					),
				),
				lang.TryDefine(nil, lang.CallIdent("rows", "Close"), "cannot close rows"),
				lang.TryDefine(nil, lang.CallIdent("rows", "Err"), "query failed"),

				ast.NewReturnStmt(ast.NewIdent("res"), ast.NewIdentLit("nil")),
			),

		)

	return fun, nil
}
