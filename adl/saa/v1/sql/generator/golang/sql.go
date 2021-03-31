package golang

import (
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	_ "github.com/pingcap/parser/test_driver"
)

// RenderSQL takes the sql context and emits the according
// options.go (contains connection options), files.go (contains migration files) and migrations.go (contains
// migration logic).
func RenderSQL(dst *ast.Prj, src *sql.Ctx) error {
	if len(src.Migrations) == 0 {
		return nil
	}

	if err := RenderOptions(dst, src.Mod.String(), src.Pkg.String(), src.Dialect); err != nil {
		return err
	}

	if err := RenderFiles(dst, src); err != nil {
		return err
	}

	if err := RenderDBTX(dst, src); err != nil {
		return err
	}

	if err := RenderMigrations(dst, src); err != nil {
		return err
	}

	if err := renderRepositories(dst, src); err != nil {
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
