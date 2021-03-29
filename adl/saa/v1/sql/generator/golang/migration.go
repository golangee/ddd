package golang

import (
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	_ "github.com/pingcap/parser/test_driver"
)

const (
	filenameMigrations = "migrations.gen.go"
)

func RenderMigrations(dst *ast.Prj, src *sql.Ctx) error {
	if len(src.Migrations) == 0 {
		return nil
	}

	if err := RenderOptions(dst, src.Mod.String(), src.Pkg.String(), src.Dialect); err != nil {
		return err
	}

	if err := RenderFiles(dst, src); err != nil {
		return err
	}

	return nil
}

func RenderRepository(dst *ast.Prj, src *sql.RepoImplSpec) error {
	//mod := astutil.MkMod(dst, src.Parent.Parent.Name.String())
	//pkg := astutil.MkPkg(mod, src.Parent.Path.String())

	//	_ = pkg

	return nil
}
