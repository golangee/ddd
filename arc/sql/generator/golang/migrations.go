package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"strconv"
	"strings"
	"time"
)

const (
	filenameMigrations = "migrations.go"
)

// RenderMigrations expects the result from files.go (RenderFiles).
// TODO how to solve multiple mysql/mariadb migrations due to concurrent k8s replicas? => some DB based locking?
func RenderMigrations(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()
	file := golang.MkFile(dst, modName, pkgName, filenameMigrations)

	tableName := strings.ReplaceAll(golang.ShortModName(file)+"_"+golang.PkgRelativeName(file), "/", "_")
	tableName += "_migration_schema_history"

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

	if err := renderMigrationFunc(file, src, tableName); err != nil {
		return fmt.Errorf("cannot render migration func: %w", err)
	}

	if err := renderMigrationsFunc(file, src); err != nil {
		return fmt.Errorf("cannot render migrations: %w", err)
	}

	return nil
}

// renderMigrationsFunc creates the func which returns all embedded and hardcoded migrations.
func renderMigrationsFunc(dst *ast.File, src *sql.Ctx) error {
	sb := &strings.Builder{}
	sb.WriteString("return []migration {\n")
	lastId := time.Time{}
	for _, migration := range src.Migrations {
		if migration.ID.Unix() <= lastId.Unix() {
			return fmt.Errorf("sql migrations are unordered or not monotonic")
		}

		lastId = migration.ID

		sb.WriteString("{\n")
		sb.WriteString(fmt.Sprintf("Version: %d, // %s\n", migration.ID.Unix(), migration.ID.String()))
		sb.WriteString(fmt.Sprintf("Description: %s,\n", strconv.Quote(migration.Name.String())))
		sb.WriteString(fmt.Sprintf("File: %s,\n", strconv.Quote(migration.Name.Begin().File)))
		sb.WriteString(fmt.Sprintf("Line: %d,\n", migration.Name.Begin().Line))
		sb.WriteString(fmt.Sprintf("Checksum: %s,\n", varMigrationHashName(migration.Name.String())))
		sb.WriteString("Statements: []string{\n")
		for i := range migration.Statements {
			constName := varMigrationStatementName(migration.Name.String(), i)
			sb.WriteString(constName)
			sb.WriteString(",\n")
		}
		sb.WriteString("},\n")
		sb.WriteString("},\n")
	}

	sb.WriteString("}\n")

	dst.AddFuncs(
		ast.NewFunc("migrations").
			SetVisibility(ast.PackagePrivate).
			SetComment("...returns all available migrations in sorted order from oldest to latest.").
			AddResults(
				ast.NewParam("", ast.NewSliceTypeDecl(ast.NewSimpleTypeDecl("migration"))),
			).
			SetBody(ast.NewBlock(ast.NewTpl(sb.String()))),

	)

	return nil
}

// renderMigrationFunc creates the func to perform the actual migration.
func renderMigrationFunc(dst *ast.File, src *sql.Ctx, tableName string) error {
	dst.AddFuncs(
		ast.NewFunc("Migrate").
			SetComment("...ensures that the migration history table exists, checks the checksums of all already applied migrations\nand applies all missing migrations in the defined version order.").
			AddParams(ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX"))).
			AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
			SetBody(
				ast.NewBlock(
					ast.NewTpl(`
						// if we can start a transaction on our own, do so and invoke recursively
						if db,ok := db.(*sql.DB);ok{
							tx, err := db.BeginTx(context.Background(),nil)
							if err != nil{
								return {{.Use "fmt.Errorf"}}("cannot begin transaction: %w",err)
							}
							
							if err := Migrate(tx);err!=nil{
								if suppressedErr := tx.Rollback(); suppressedErr != nil {
									fmt.Println(suppressedErr.Error())
								}
								return err
							}
					
							if err := tx.Commit(); err != nil {
								return err
							}
							
							return nil
						}

						const q = {{.Get "createStatement"}}
						if _, err := db.ExecContext({{.Use "context.Background"}}(), q); err!=nil{
							return {{.Use "fmt.Errorf"}}("cannot create {{.Get "tableName"}}: %w", err)
						}

						history, err := readMigrationHistoryTable(db)
						if err != nil {
							return {{.Use "fmt.Errorf"}}("cannot read history: %w",err)
						}

						availMigrations := migrations()
		
						// check history validity:
						// 1. is everything which has been applied still defined?
						// 2. has any checksum changed?
						// 3. do we have any unclean migration?
						for _, entry := range history {
							found := false
							for _, m := range availMigrations {
								if entry.Status != "success" {
									return fmt.Errorf("found an incomplete migration. Your database is inconsistent and you have to solve this manually. Affected migration: %s", entry.String())
								}
		
								if entry.Version == m.Version {
									if entry.Checksum != m.Checksum {
										return fmt.Errorf("already applied migration %s has been modified. Expected %s but found %s", entry.String(), entry.Checksum, m.Checksum)
									}
		
									found = true
									break
								}
			
							}
		
							if !found {
								return fmt.Errorf("already applied migration %s is undefined", entry.String())
							}
						}
		
						// pick migrations to apply
						for _, m := range availMigrations {
							alreadyApplied := false
							for _, entry := range history {
								if m.Version == entry.Version {
									alreadyApplied = true
									break
								}
							}
		
							if !alreadyApplied {
								start := time.Now()
		
								entry := migrationEntry{
									Version:           m.Version,
									File:              m.File,
									Line:              m.Line,
									Checksum:          m.Checksum,
									AppliedAt:         start.Unix(),
									Description:       m.Description,
									Status:			   "pending",
								}
								err = entry.insert(db)
								if err != nil {
									return fmt.Errorf("unable to insert migration state %s: %w", m.String(), err)
								}
		
								err := m.Apply(db)
								if err != nil {
									return fmt.Errorf("unable to apply migration %s: %w", m.String(), err)
								}
		
								entry.ExecutionDuration = time.Now().Sub(start).Nanoseconds()
								entry.Status = "success"
								err = entry.update(db)
								if err != nil {
									return fmt.Errorf("unable to update migration state %s: %w", m.String(), err)
								}
							}
						}
						
						return nil

					`).Put("createStatement", strconv.Quote(strings.Join(strings.Split(sqlCreateTableMigrationHistory(tableName, src.Dialect), "\n"), " "))).
						Put("tableName", tableName),
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

				stereotype.FieldFrom(
					ast.NewField("File", ast.NewSimpleTypeDecl(stdlib.String)).
						SetComment("...is the file path indicating the origin of the statements."),
				).SetSQLColumnName("file").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("Line", ast.NewSimpleTypeDecl(stdlib.Int32)).
						SetComment("...is the line number indicating the origin of the statements."),
				).SetSQLColumnName("line").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("Checksum", ast.NewSimpleTypeDecl(stdlib.String)).
						SetComment("...is the hex encoded 28 byte sha3-224 checksum of all trimmed statements."),
				).SetSQLColumnName("checksum").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("AppliedAt", ast.NewSimpleTypeDecl(stdlib.Int64)).
						SetComment("...is the unix timestamp in seconds when this migration has been applied."),
				).SetSQLColumnName("applied_at").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("ExecutionDuration", ast.NewSimpleTypeDecl(stdlib.Int64)).
						SetComment("...is the amount of nanoseconds which were needed to apply this migration."),
				).SetSQLColumnName("execution_duration").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("Description", ast.NewSimpleTypeDecl(stdlib.String)).
						SetComment("...describes at least why the migration is needed."),
				).SetSQLColumnName("description").Unwrap(),

				stereotype.FieldFrom(
					ast.NewField("Status", ast.NewSimpleTypeDecl(stdlib.String)).
						SetComment("...is the status of the migration"),
				).SetSQLColumnName("status").Unwrap(),


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
	stereotype.StructFrom(entity).
		SetSQLTableName(tableName).
		SetSQLDefaultOrder("ORDER BY version ASC")

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
