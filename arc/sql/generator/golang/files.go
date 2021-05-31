package golang

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/sql/dialect/mysql"
	"github.com/golangee/src/ast"
	golang2 "github.com/golangee/src/golang"
	"strconv"
)

const (
	filenameFiles = "files.go"
)

// RenderFiles normalizes the given migrations and embeds them.
func RenderFiles(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := golang.MkFile(dst, modName, pkgName, filenameFiles)

	for _, migration := range src.Migrations {
		migrationConstBlock := ast.NewConstDecl()
		migrationHashConstBlock := ast.NewConstDecl()

		hashBuf := &bytes.Buffer{}
		for i, statement := range migration.Statements {
			_, err := mysql.Parse(statement.String())
			if err != nil {
				return token.NewPosError(statement, "cannot parse sql statement").SetCause(err)
			}

			hashBuf.WriteString(statement.String())

			constName := varMigrationStatementName(migration.Name.String(), i)
			migrationConstBlock.Add(
				ast.NewSimpleAssign(ast.NewIdent(constName), ast.AssignSimple, ast.NewStrLit(statement.String())).
					SetComment("...is defined in file " + statement.BeginPos.File + " line " + strconv.Itoa(statement.BeginPos.Line) + "."),
			)

		}

		migrationHash := sha512.Sum512_224(hashBuf.Bytes())
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
	return "migrate" + golang2.MakeIdentifier(migrationName) + strconv.Itoa(statementNo+1)
}

func varMigrationHashName(migrationName string) string {
	return "migrate" + golang2.MakeIdentifier(migrationName) + "Hash"
}
