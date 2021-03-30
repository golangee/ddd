package golang

import (
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/stereotype"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"strings"
)

const (
	filenameMigrations = "migrations.go"
)

// RenderMigrations expects the result from files.go (RenderFiles).
func RenderMigrations(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := corego.MkFile(dst, modName, pkgName, filenameMigrations)

	tableName := strings.ReplaceAll(corego.ShortModName(file)+"_"+corego.PkgRelativeName(file), "/", "_")
	tableName += "_migration_schema_history"

	if err := renderMigrationFunc(file, src, tableName); err != nil {
		return fmt.Errorf("cannot render migration func: %w", err)
	}

	historyEntity, err := renderMigrationEntity(file, src, tableName)
	if err != nil {
		return fmt.Errorf("cannot render migration entity: %w", err)
	}

	if err := renderMigrationStruct(file, src, tableName); err != nil {
		return fmt.Errorf("cannot render migration struct: %w", err)
	}

	findAllHistory, err := renderStaticFindAll(historyEntity, "DBTX")
	if err != nil {
		return fmt.Errorf("cannot render find all history entities: %w", err)
	}

	findAllHistory.FunName = "readMigrationHistoryTable"
	file.AddFuncs(findAllHistory.SetVisibility(ast.PackagePrivate))

	return nil
}

func renderMigrationFunc(dst *ast.File, src *sql.Ctx, tableName string) error {
	dst.AddFuncs(
		ast.NewFunc("Migrate").
			SetComment("...ensures that the migration history table exists, checks the checksums of all already applied migrations\nand applies all missing migrations in the defined version order.").
			AddParams(ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX"))).
			AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
			SetBody(
				ast.NewBlock(

					ast.NewReturnStmt(ast.NewIdentLit("nil")),
				),
			),

	)

	return nil
}

func renderMigrationStruct(dst *ast.File, src *sql.Ctx, tableName string) error {
	dst.AddTypes(
		ast.NewStruct("migration").
			SetVisibility(ast.PackagePrivate).
			SetComment("...represents a single isolated database migration, which may consist of multiple statements but which must be executed atomically.").
			AddFields(
				ast.NewField("Version", ast.NewSimpleTypeDecl(stdlib.Int64)).
					SetComment("...is the unix timestamp in seconds, at which this migration was defined."),
				ast.NewField("Description", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...describes at least why the migration is needed."),
				ast.NewField("Statements", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl(stdlib.String))).
					SetComment("...contains e.g. CREATE, ALTER or DROP statements to apply."),
				ast.NewField("File", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...is the file path indicating the origin of the statements."),
				ast.NewField("Line", ast.NewSimpleTypeDecl(stdlib.Int32)).
					SetComment("...is the line number indicating the origin of the statements."),
				ast.NewField("Checksum", ast.NewSimpleTypeDecl(stdlib.Int32)).
					SetComment("...is the hex encoded 28 byte sha3-224 checksum of all trimmed statements."),
			).
			AddMethods(
				ast.NewFunc("apply").
					SetVisibility(ast.PackagePrivate).
					SetComment("... executes the statements.").
					SetRecName("m").
					AddParams(ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX"))).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetBody(ast.NewBlock(
						ast.NewRangeStmt(
							nil,
							ast.NewIdent("s"),
							lang.Sel("m", "Statements"),
							ast.NewBlock(
								lang.TryDefine(
									ast.NewIdentLit("_"),
									lang.CallIdent("db", "ExecContext", lang.CallStatic("context.Background")),
									"cannot execute statement",
								),
							),
						),
						lang.Term(),
					)),
			),
	)

	return nil
}

func renderMigrationEntity(dst *ast.File, src *sql.Ctx, tableName string) (*ast.Struct, error) {
	const recName = "m"
	entity :=
		ast.NewStruct("migrationEntry").
			SetVisibility(ast.PackagePrivate).
			SetComment("...represents a row entry from the migration schema history table.").
			SetDefaultRecName(recName).
			AddFields(
				stereotype.FieldFrom(
					ast.NewField("Version", ast.NewSimpleTypeDecl(stdlib.Int64)).
						SetComment("...represents a row entry from the migration schema history table."),
				).SetSQLColumnName("version").Unwrap(),
				ast.NewField("File", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...is the file path indicating the origin of the statements."),
				ast.NewField("Line", ast.NewSimpleTypeDecl(stdlib.Int32)).
					SetComment("...is the line number indicating the origin of the statements."),
				ast.NewField("Checksum", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...is the hex encoded 28 byte sha3-224 checksum of all trimmed statements."),
				ast.NewField("AppliedAt", ast.NewSimpleTypeDecl(stdlib.Int64)).
					SetComment("...is the unix timestamp in seconds when this migration has been applied."),
				ast.NewField("ExecutionDuration", ast.NewSimpleTypeDecl(stdlib.Int64)).
					SetComment("...is the amount of nanoseconds which were needed to apply this migration."),
				ast.NewField("Description", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...describes at least why the migration is needed."),
				ast.NewField("Status", ast.NewSimpleTypeDecl(stdlib.String)).
					SetComment("...is the status of the migration"),

			).
			AddMethods(
				ast.NewFunc("insert").
					SetVisibility(ast.PackagePrivate).
					SetComment("...writes a migrationEntry into the history table.").
					SetRecName(recName).
					AddParams(ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX"))).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetBody(
						ast.NewBlock(
							ast.NewConstDecl(ast.NewSimpleAssign(ast.NewIdent("q"), ast.AssignSimple, ast.NewStrLit(sqlInsertIntoMigrationHistory(tableName, src.Dialect)))),
							lang.TryDefine(
								ast.NewIdent("_"),
								lang.CallIdent("db", "ExecContext",
									sqlArgs("m.Version", "m.File", "m.Line", "m.Checksum", "m.AppliedAt", "m.ExecutionDuration", "m.Description", "m.Status")...),
								"cannot insert migration entry",
							),
							ast.NewReturnStmt(ast.NewIdent("nil")),
						),
					),

				ast.NewFunc("update").
					SetVisibility(ast.PackagePrivate).
					SetComment("...writes a migrationEntry into the history table and replaces the exiting entry identified by version.").
					SetRecName(recName).
					AddParams(ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX"))).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetBody(
						ast.NewBlock(
							ast.NewConstDecl(ast.NewSimpleAssign(ast.NewIdent("q"), ast.AssignSimple, ast.NewStrLit(sqlUpdateIntoMigrationHistory(tableName, src.Dialect)))),
							lang.TryDefine(
								ast.NewIdent("_"),
								lang.CallIdent("db", "ExecContext",
									sqlArgs("m.File", "m.Line", "m.Checksum", "m.AppliedAt", "m.ExecutionDuration", "m.Description", "m.Status", "m.Version")...),
								"cannot update migration entry",
							),
							ast.NewReturnStmt(ast.NewIdent("nil")),
						),
					),

			)

	dst.AddTypes(entity)
	stereotype.StructFrom(entity).SetSQLTableName(tableName)

	return entity, nil
}

func sqlInsertIntoMigrationHistory(tableName string, dialect sql.Dialect) string {
	switch dialect {
	case sql.MySQL:
		return `INSERT INTO ` + tableName + `(version, file, line, checksum, applied_at, execution_duration, description, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	default:
		panic("dialect not implemented: " + string(dialect))
	}
}

func sqlUpdateIntoMigrationHistory(tableName string, dialect sql.Dialect) string {
	switch dialect {
	case sql.MySQL:
		return `UPDATE ` + tableName + ` SET file = ?, line = ?, checksum = ?, applied_at = ?, execution_duration = ?, description = ?, status = ? WHERE version = ?`
	default:
		panic("dialect not implemented: " + string(dialect))
	}
}

func sqlCreateTableMigrationHistory(tableName string, dialect sql.Dialect) string {
	switch dialect {
	case sql.MySQL:
		createMigrationTable := "CREATE TABLE IF NOT EXISTS \"" + tableName
		createMigrationTable += `"
(
    "version"            BIGINT       NOT NULL,
    "file"               VARCHAR(255) NOT NULL,
    "line"               INT          NOT NULL,
    "checksum"           CHAR(32)     NOT NULL,
    "applied_at"         BIGINT       NOT NULL,
    "execution_duration" BIGINT       NOT NULL,
	"description"		 TEXT         NOT NULL,
	"status"			 VARCHAR(255) NOT NULL,
    PRIMARY KEY ("version")
)`
		return createMigrationTable
	default:
		panic("dialect not implemented: " + string(dialect))
	}
}

func sqlArgs(idents ...string) []ast.Expr {
	var r []ast.Expr
	r = append(r, lang.CallStatic("context.Background"))

	for _, ident := range idents {
		r = append(r, ast.NewIdent(ident))
	}

	return r
}
