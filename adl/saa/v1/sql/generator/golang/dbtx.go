package golang

import (
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
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

	if err := renderOpen(file, src); err != nil {
		return err
	}

	return nil
}


func renderOpen(file *ast.File, src *sql.Ctx) error {
	file.AddNodes(ast.NewImport("_", "github.com/go-sql-driver/mysql").SetComment("side-effect-only import to load mysql driver"))

	file.AddFuncs(
		ast.NewFunc("Open").
			SetComment("...tries to connect to a mysql compatible database.").
			AddParams(
				ast.NewParam("opts", ast.NewSimpleTypeDecl("Options")),
			).
			AddResults(
				ast.NewParam("", ast.NewSimpleTypeDecl("DBTX")),
				ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
			).
			SetBody(
				ast.NewBlock(
					lang.TryDefine(
						ast.NewIdent("db"),
						lang.CallStatic("database/sql.Open", lang.CallIdent("opts", "DSN")),
						"cannot open mysql database",
					),

					lang.TryDefine(
						nil,
						lang.CallIdent("db", "Ping"),
						"cannot ping mysql database",
					),

					lang.CallIdent("db", "SetConnMaxLifetime", lang.Sel("opts", "ConnMaxLifetime")),
					lang.Term(),
					lang.CallIdent("db", "SetMaxOpenConns", lang.Sel("opts", "MaxOpenConns")),
					lang.Term(),
					lang.CallIdent("db", "SetMaxIdleConns", lang.Sel("opts", "MaxIdleConns")),
					lang.Term(),
					lang.Term(),
					ast.NewReturnStmt(ast.NewIdent("db"), ast.NewIdent("nil")),
				),
			),
	)
	return nil
}