package golang

import (
	"encoding/hex"
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/golang"
	"github.com/pingcap/parser"
	"golang.org/x/crypto/sha3"
	"strconv"
)

const (
	filenameFiles = "files.go"
)

// RenderFiles normalizes the given migrations and embeds them.
func RenderFiles(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := corego.MkFile(dst, modName, pkgName, filenameFiles)

	migrationConstBlock := ast.NewConstDecl()
	migrationHashConstBlock := ast.NewConstDecl()
	for _, migration := range src.Migrations {
		for i, statement := range migration.Statements {
			p := parser.New()
			_, _, err := p.Parse(statement.String(), "UTF-8", "UTF-8")
			if err != nil {
				return core.NewPosError(statement, "cannot parse sql statement").SetCause(err)
			}

			migrationHash := sha3.Sum224([]byte(statement.String()))
			migrationHashStr := hex.EncodeToString(migrationHash[:])

			constName := "migrate" + golang.MakeIdentifier(migration.Name.String()) + strconv.Itoa(i+1)
			migrationConstBlock.Add(
				ast.NewSimpleAssign(ast.NewIdent(constName), ast.AssignSimple, ast.NewStrLit(statement.String())).
					SetComment("...is defined in file " + statement.NodePos.File + "."),
			)

			migrationHashConstBlock.Add(
				ast.NewSimpleAssign(ast.NewIdent(constName+"Hash"), ast.AssignSimple, ast.NewStrLit(migrationHashStr)),
			)
		}
	}

	file.AddNodes(migrationConstBlock, migrationHashConstBlock)

	return nil
}
