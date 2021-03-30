package golang

import (
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
)

const (
	filenameMigrations = "migrations.go"
)

// RenderMigrations expects the result from files.go (RenderFiles).
func RenderMigrations(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := corego.MkFile(dst, modName, pkgName, filenameMigrations)

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
		),
	)
	return nil
}
