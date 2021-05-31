package golang

import (
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/src/ast"
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

