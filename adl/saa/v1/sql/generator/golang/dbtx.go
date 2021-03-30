package golang

import (
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
)

const (
	filenameDBTX = "dbtx.go"
)

// RenderDBTX emits a dbtx.go file which contains some sql utilities.
func RenderDBTX(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := corego.MkFile(dst, modName, pkgName, filenameDBTX)

	file.AddTypes(
		ast.NewInterface("DBTX").
			SetComment("...abstracts from a concrete sql.DB or sql.Tx dependency.").
			AddMethods(
				ast.NewFunc("ExecContext").
					SetComment("...represents an according call to sql.DB or sql.Tx").
					AddParams(
						ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context")),
						ast.NewParam("query", ast.NewSimpleTypeDecl(stdlib.String)),
						ast.NewParam("args", ast.NewSimpleTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						ast.NewParam("", ast.NewSimpleTypeDecl("database/sql.Result")),
						ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
					),

				ast.NewFunc("QueryContext").
					SetComment("...represents an according call to sql.DB or sql.Tx").
					AddParams(

						ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context")),
						ast.NewParam("query", ast.NewSimpleTypeDecl(stdlib.String)),
						ast.NewParam("args", ast.NewSimpleTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						ast.NewParam("", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl("database/sql.Rows"))),
						ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
					),
			),

	)

	return nil
}
