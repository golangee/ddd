package golang

import (
	"embed"
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/golang"
	"testing"
)

//go:embed *.sql
var fs embed.FS

func TestRenderRepository(t *testing.T) {
	prj := ast.NewPrj("test")
	ctx := createCtx(t)
	if err := RenderMigrations(prj, ctx); err != nil {
		t.Fatal(core.Explain(err))
	}

	renderer := golang.NewRenderer(golang.Options{})
	a, err := renderer.Render(prj)
	if err != nil {
		fmt.Println(a)
		t.Fatal(err)
	}

	fmt.Println(a)
}

func createCtx(t *testing.T) *sql.Ctx {
	t.Helper()

	mod := core.NewModLit("github.com/worldiety/supportiety")
	pkg := core.NewPkgLit("github.com/worldiety/supportiety/tickets/core")

	return &sql.Ctx{
		Dialect:    sql.MySQL,
		Mod:        mod,
		Pkg:        pkg,
		Migrations: createMigrations(t),
	}
}

func createMigrations(t *testing.T) []*sql.Migration {
	t.Helper()

	entries, err := fs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}

	var migrations []*sql.Migration
	for _, entry := range entries {
		ts, name, err := sql.ParseMigrationName(entry.Name())
		if err != nil {
			t.Fatal(err)
		}

		file, err := fs.Open(entry.Name())
		if err != nil {
			t.Fatal(err)
		}

		stmts, err := sql.ParseStatements(file)
		if err != nil {
			t.Fatal(err)
		}

		migrations = append(migrations, &sql.Migration{
			ID:         ts,
			Name:       core.NewStrLit(name),
			Statements: stmts,
		})
	}

	return migrations
}
