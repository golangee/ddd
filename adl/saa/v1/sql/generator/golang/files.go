package golang

import (
	"bytes"
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

	for _, migration := range src.Migrations {
		migrationConstBlock := ast.NewConstDecl()
		migrationHashConstBlock := ast.NewConstDecl()

		hashBuf := &bytes.Buffer{}
		for i, statement := range migration.Statements {
			p := parser.New()
			_, _, err := p.Parse(statement.String(), "UTF-8", "UTF-8")
			if err != nil {
				return core.NewPosError(statement, "cannot parse sql statement").SetCause(err)
			}

			hashBuf.WriteString(statement.String())

			constName := varMigrationStatementName(migration.Name.String(), i)
			migrationConstBlock.Add(
				ast.NewSimpleAssign(ast.NewIdent(constName), ast.AssignSimple, ast.NewStrLit(statement.String())).
					SetComment("...is defined in file " + statement.NodePos.File + " line " + strconv.Itoa(statement.NodePos.Line) + "."),
			)

		}

		migrationHash := sha3.Sum224(hashBuf.Bytes())
		migrationHashStr := hex.EncodeToString(migrationHash[:])
		migrationHashConstBlock.Add(
			ast.NewSimpleAssign(ast.NewIdent(varMigrationHashName(migration.Name.String())), ast.AssignSimple, ast.NewStrLit(migrationHashStr)).
				SetComment("...contains the sha3-224 hash of all related and normalized sql statements."),
		)

		file.AddNodes(migrationConstBlock, migrationHashConstBlock)
	}

	return nil
}

func varMigrationStatementName(migrationName string, statementNo int) string {
	return "migrate" + golang.MakeIdentifier(migrationName) + strconv.Itoa(statementNo+1)
}

func varMigrationHashName(migrationName string) string {
	return "migrate" + golang.MakeIdentifier(migrationName) + "Hash"
}
