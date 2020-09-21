package golang

import (
	"encoding/hex"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"golang.org/x/crypto/sha3"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func createSQLLayer(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())

	for _, repo := range sql.Repositories() {
		file := ctx.newFile(layerPath, text.Safename(repo.InterfaceName()), "").
			SetPackageDoc(sql.Description())

		iface := bc.SPIServiceByName(repo.InterfaceName())
		if iface == nil {
			panic("illegal state: validate the model first")
		}

		repoSpec := ctx.repoSpecByName(repo.InterfaceName())
		if repoSpec == nil {
			panic("illegal state: validate the model first and define core layer before sql layer")
		}

		impl := src.Implement(repoSpec.iface, true)
		impl.SetName(sql.Name() + repo.InterfaceName())
		impl.AddFields(
			src.NewField("db", src.NewTypeDecl(rslv.assembleQualifier(rMysql, "DBTX"))),
		)
		impl.SetDoc("...is an implementation of the " + pkgNameCore + "." + repo.InterfaceName() + " defined as SPI/driven port in the domain/core layer.\nThe queries are specific for the " + strings.ToLower(sql.Name()) + " dialect.")
		file.AddTypes(impl)

		for _, method := range impl.Methods() {
			sqlSpec := repo.ImplementationByName(method.Name())
			if sqlSpec == nil {
				panic("illegal state: undefined method: validate the model first")
			}
			body := src.NewBlock()
			method.AddBody(body)
			if len(sqlSpec.Row()) == 0 {
				createSQLExec(sqlSpec, method, body)
			} else {
				if method.Results()[0].Decl().Qualifier() == "[]" {
					createSQLQueryMany(sqlSpec, method, body)
				} else {
					createSQLQueryOne(sqlSpec, method, body)
				}
			}

		}

		repoTypeDecl, err := rslv.resolveTypeName(rCore, ddd.TypeName(iface.Name()))
		if err != nil {
			return err
		}

		repoFac := src.NewFunc("New"+impl.Name()).
			SetDoc("...creates a new instance of "+impl.Name()+".").
			AddParams(src.NewParameter("db", src.NewTypeDecl(rslv.assembleQualifier(rMysql, "DBTX")))).
			AddResults(
				src.NewParameter("", repoTypeDecl),
				src.NewParameter("", src.NewTypeDecl("error")),
			).
			AddBody(src.NewBlock("return &" + impl.Name() + "{db:db},nil"))

		file.AddFuncs(repoFac)

		ctx.factorySpecs = append([]*factorySpec{{
			file:        file,
			factoryFunc: repoFac,
		}}, ctx.factorySpecs...)
	}

	migrationFile, err := createSQLMigration(ctx, rslv, bc, sql)
	if err != nil {
		return err
	}

	ctx.addFirstFactorySpec(migrationFile, src.NewFunc("Migrate").
		AddParams(src.NewParameter("db", src.NewTypeDecl(rslv.assembleQualifier(rMysql, "DBTX")))).
		AddResults(src.NewParameter("", src.NewTypeDecl("error"))),
		nil)

	if err := createSQLUtil(ctx, rslv, bc, sql); err != nil {
		return err
	}

	return nil
}

func createSQLQueryMany(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")

	body.Add("r, err := ", method.ReceiverName(), ".db.QueryContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")
	body.If("err!=nil", src.NewBlock(
		"return nil, ", src.NewTypeDecl("fmt.Errorf"), "(\"QueryContext failed: %w\",err)",
	))
	body.AddLine("defer r.Close()")

	body.AddLine("var l ", method.Results()[0].Decl().Clone())
	body.AddLine("for r.Next() {")
	body.AddLine("var i ", method.Results()[0].Decl().Params()[0].Clone())
	body.Add("if err := r.Scan(")
	for _, row := range sqlSpec.Row() {
		body.Add(makeSqlVarAccess("i", string(row)), ",")
	}
	body.AddLine(");err!=nil{")
	body.AddLine("return nil, ", src.NewTypeDecl("fmt.Errorf"), "(\"scan failed: %w\",err)")
	body.AddLine("}")
	body.AddLine("l = append(l, i)")
	body.AddLine("}")

	body.AddLine("err = r.Close()")
	body.Check("err", "cannot close rows", "l")
	body.NewLine()

	body.AddLine("err = r.Err()")
	body.Check("err", "query failed", "l")
	body.NewLine()

	body.AddLine("return l, nil")
}

func makeSqlVarAccess(name, accessor string) string {
	if accessor == "&." {
		return "&" + name
	}

	if strings.HasPrefix(accessor, "&.") {
		return "&" + name + "." + accessor[2:]
	}

	if accessor == "." {
		return name
	}

	if strings.HasPrefix(accessor, ".") {
		return name + accessor
	}

	return accessor
}

func createSQLQueryOne(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")

	body.AddLine("var i ", method.Results()[0].Decl().Clone())
	body.Add("r, err := ", method.ReceiverName(), ".db.QueryContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")
	body.If("err!=nil", src.NewBlock(
		"return i, ", src.NewTypeDecl("fmt.Errorf"), "(\"QueryContext failed: %w\",err)",
	))
	body.AddLine("defer r.Close()")

	body.AddLine("for r.Next() {")
	body.Add("if err := r.Scan(")
	for _, row := range sqlSpec.Row() {
		body.Add(makeSqlVarAccess("i", string(row)), ",")
	}
	body.AddLine(");err!=nil{")
	body.AddLine("return i, ", src.NewTypeDecl("fmt.Errorf"), "(\"scan failed: %w\",err)")
	body.AddLine("}")

	body.AddLine("err = r.Close()")
	body.Check("err", "cannot close rows", "i")
	body.NewLine()

	body.AddLine("err = r.Err()")
	body.Check("err", "query failed", "i")
	body.NewLine()

	body.AddLine("return i, err")
	body.AddLine("}")

	body.AddLine("return i, ", src.NewTypeDecl("fmt.Errorf(\"empty result set\")"))
}

func createSQLExec(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")
	body.Add("_, err := ", method.ReceiverName(), ".db.ExecContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")

	body.Check("err", "ExecContext failed")
	body.NewLine()

	body.AddLine("return nil")
}

func createSQLUtil(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())

	file := ctx.newFile(layerPath, "db", "").
		SetPackageDoc(sql.Description())

	file.AddTypes(
		src.NewInterface("DBTX").
			SetDoc("...abstracts from a concrete sql.DB or sql.Tx dependency.").
			AddMethods(
				src.NewFunc("ExecContext").SetDoc("...represents an according call to sql.DB or sql.Tx").
					AddParams(
						src.NewParameter("ctx", src.NewTypeDecl("context.Context")),
						src.NewParameter("query", src.NewTypeDecl("string")),
						src.NewParameter("args", src.NewTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						src.NewParameter("", src.NewTypeDecl("database/sql.Result")),
						src.NewParameter("", src.NewTypeDecl("error")),
					),

				src.NewFunc("QueryContext").SetDoc("...represents an according call to sql.DB or sql.Tx").
					AddParams(
						src.NewParameter("ctx", src.NewTypeDecl("context.Context")),
						src.NewParameter("query", src.NewTypeDecl("string")),
						src.NewParameter("args", src.NewTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						src.NewParameter("", src.NewPointerDecl(src.NewTypeDecl("database/sql.Rows"))),
						src.NewParameter("", src.NewTypeDecl("error")),
					),
			),
	)

	opts := createMySQLOptions(rslv, text.Safename(ctx.spec.Name()), bc)
	file.AddTypes(opts)

	dbFactory := createMySQLOpen(rslv.assembleQualifier(rMysql, opts.Name()))
	file.AddFuncs(dbFactory)

	file.AddSideEffectImports("github.com/go-sql-driver/mysql")

	ctx.factorySpecs = append([]*factorySpec{{
		file:        file,
		factoryFunc: dbFactory,
		options:     opts,
	}}, ctx.factorySpecs...)

	return nil
}

func createMySQLOptions(rslv *resolver, defaultDBName string, bc *ddd.BoundedContextSpec) *src.TypeBuilder {
	opt := ddd.Struct("Options",

		"...contains the connection options for a MySQL database.",
		ddd.Field("Port", ddd.Int64, "...is the database port to connect.").SetDefault("3306"),
		ddd.Field("User", ddd.String, "...is the database user.").SetDefault("\"root\""),
		ddd.Field("Password", ddd.String, "...is the database user password.").SetDefault(""),
		ddd.Field("Protocol", ddd.String, "...is the protocol to use.").SetDefault("\"tcp\""),
		ddd.Field("Database", ddd.String, "...is the database name.").SetDefault(`"`+defaultDBName+`"`),
		ddd.Field("Address", ddd.String, "...is the host or path to socket.").SetDefault("\"localhost\""),

		// see https://stackoverflow.com/questions/766809/whats-the-difference-between-utf8-general-ci-and-utf8-unicode-ci/766996#766996
		// https://www.percona.com/live/e17/sites/default/files/slides/Collations%20in%20MySQL%208.0.pdf
		//
		// we enforce correct unicode support for mysql and index/sorting collations. For mysql 8.0 using
		// accent insensitive/case insensitive Unicode 9 support utf8mb4_0900_ai_ci would be better but not compatible
		// with mariadb, so we use a fixed older version for reproducibility across different database servers.
		ddd.Field("Collation", ddd.String, "...declares the connections default collation for sorting and indexing.").SetDefault("\"utf8mb4_unicode_520_ci\""),
		ddd.Field("Charset", ddd.String, "...declares the connections default charset encoding for text.").SetDefault("\"utf8mb4\""),
		ddd.Field("MaxAllowedPacket", ddd.Int64, "...is the max packet size in bytes.").SetDefault("4194304"),
		ddd.Field("Timeout", ddd.Duration, "...is the duration until the dial receives a timeout.").SetDefault("30s"),
		ddd.Field("WriteTimeout", ddd.Duration, "...is the duration for the write timeout.").SetDefault("30s"),
		ddd.Field("Tls", ddd.String, "...configures connection security. Valid values are true, false, skip-verify or preferred.").SetDefault("\"false\""),
		ddd.Field("SqlMode", ddd.String, "...is a flag which influences the sql parser.").SetDefault("\"ANSI\""),
		ddd.Field("ConnMaxLifetime", ddd.Duration, "...is the duration of how long pooled connections are kept alive.").SetDefault("3m"),
		ddd.Field("MaxOpenConns", ddd.Int64, "...is the amount of how many open connections can be kept in the pool.").SetDefault("25"),
		ddd.Field("MaxIdleConns", ddd.Int64, "...is the amount of how many open connections can be idle.").SetDefault("25"),

	)

	genOpt, err := generateStruct(rslv, rUniverse, opt)
	if err != nil {
		panic("illegal state: " + err.Error())
	}

	if err := generateSetDefault("Reset", genOpt, opt); err != nil {
		panic("illegal state: " + err.Error())
	}

	envPrefix := "MySQL." + bc.Name() + "."

	if err := generateParseEnv(envPrefix, "ParseEnv", genOpt, opt); err != nil {
		panic("illegal state: " + err.Error())
	}

	if err := generateFlagsConfigure(envPrefix, "ConfigureFlags", genOpt, opt); err != nil {
		panic("illegal state: " + err.Error())
	}

	genOpt.AddMethodToJson("String", true, true, true)
	genOpt.AddMethodFromJson("Parse")

	dsn := src.NewFunc("DSN").SetPointerReceiver(true)
	genOpt.AddMethods(dsn)
	r := dsn.ReceiverName()
	dsn.
		SetDoc("...returns the options as a fully serialized datasource name.\n" +
			"The returned string is of the form:\n" +
			"  username:password@protocol(address)/dbname?param=value").
		AddResults(src.NewParameter("", src.NewTypeDecl("string")))

	body := src.NewBlock().
		AddLine("sb := &", src.NewTypeDecl("strings.Builder{}")).
		AddLine("sb.WriteString(", src.NewTypeDecl("net/url.PathEscape"), "(", r, ".User))").
		AddLine("sb.WriteString(\":\")").
		AddLine("sb.WriteString(", src.NewTypeDecl("net/url.PathEscape"), "(", r, ".Password))").
		AddLine("sb.WriteString(\"@\")").
		AddLine("sb.WriteString(", r, ".Protocol)").
		AddLine("sb.WriteString(\"(\")").
		AddLine("sb.WriteString(", r, ".Address)").
		AddLine("sb.WriteString(\")\")").
		AddLine("sb.WriteString(\"/\")").
		AddLine("sb.WriteString(", r, ".Database)").
		AddLine("sb.WriteString(\"?\")")

	options := map[string]string{
		"Charset":          "charset",
		"Collation":        "collation",
		"MaxAllowedPacket": "maxAllowedPacket",
		"Tls":              "tls",
		"Timeout":          "timeout",
		"WriteTimeout":     "writeTimeout",
		"SqlMode":          "sql_mode",
	}

	var sortedKeys []string
	for k, _ := range options {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		v := options[k]
		body.Add("sb.WriteString(")
		if genOpt.FieldByName(k).Type().Qualifier() == "string" {
			body.Add(src.NewTypeDecl("fmt.Sprintf"), "(\"%s=%s&\",", `"`+v+`",`, src.NewTypeDecl("net/url.QueryEscape"), "(", r, "."+k, "))")
		} else {
			body.Add(src.NewTypeDecl("fmt.Sprintf"), "(\"%s=%v&\",", `"`+v+`",`, r, "."+k, ")")
		}
		body.AddLine(")")
	}

	body.AddLine("return sb.String()")
	dsn.AddBody(body)

	return genOpt
}

func createMySQLOpen(opts src.Qualifier) *src.FuncBuilder {
	return src.NewFunc("Open").
		SetDoc("...tries to connect to a mysql compatible database.").
		AddParams(src.NewParameter("opts", src.NewTypeDecl(opts))).
		AddResults(
			src.NewParameter("", src.NewTypeDecl(src.Qualifier(opts.Path()+".DBTX"))), // actually this is just "database/sql.DB" but our injector in genapp.go cannot match types and interfaces yet
			src.NewParameter("", src.NewTypeDecl("error")),
		).
		AddBody(src.NewBlock().
			AddLine("db,err := ", src.NewTypeDecl("database/sql.Open"), "(\"mysql\", opts.DSN())").
			Check("err", "cannot open database", "nil").NewLine().
			AddLine("err = db.Ping()").
			Check("err", "cannot ping database", "nil").
			NewLine().
			AddLine("db.SetConnMaxLifetime(opts.ConnMaxLifetime)").
			AddLine("db.SetMaxOpenConns(int(opts.MaxOpenConns))").
			AddLine("db.SetMaxIdleConns(int(opts.MaxIdleConns))").
			AddLine("return db,nil"),

		)
}

// createSQLMigrationTooling
func createSQLMigration(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) (*src.FileBuilder, error) {
	const format = "2006-01-02T15:04:05"
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())
	file := ctx.newFile(layerPath, "migration", "")

	if err := createSQLMigrationTooling(ctx, rslv, file, bc, sql); err != nil {
		return nil, err
	}

	migrationBody := src.NewBlock().AddLine("return []", src.NewTypeDecl("migration"), "{")
	file.AddFuncs(src.NewFunc("migrations").
		SetDoc("...returns all available migrations in sorted order from oldest to latest.").
		AddResults(src.NewParameter("", src.NewSliceDecl(src.NewTypeDecl("migration")))).
		AddBody(migrationBody))

	for _, m := range sql.Migrations() {
		t, err := time.Parse(format, m.DateTime())
		if err != nil {
			panic("illegal state: validate the model first: " + err.Error())
		}
		relPath, err := filepath.Rel(ctx.archMod.Main().Dir, m.Pos().File)
		if err != nil {
			panic("illegal state: " + err.Error())
		}

		var statements []string
		for _, s := range m.RawStatements() {
			statements = append(statements, string(s))
		}

		migrationBody.AddLine("{")
		migrationBody.AddLine("Version: ", t.Unix(), ", // ", m.DateTime())
		migrationBody.AddLine("File: \"", relPath, "\",")
		migrationBody.AddLine("Line: ", m.Pos().Line, ",")
		migrationBody.AddLine("Checksum: \"", sqlChecksum(statements), "\",")
		migrationBody.AddLine("Description: ", strconv.Quote(m.Description()), ",")
		migrationBody.AddLine("Statements: []string{")
		for _, statement := range statements {
			migrationBody.AddLine(strconv.Quote(normalizeSQLStatement(statement)), ",")
		}
		migrationBody.AddLine("},")
		migrationBody.AddLine("},")
		migrationBody.AddLine()
	}
	migrationBody.AddLine("}")

	return file, nil
}

// sqlChecksum normalizes whitespaces of linebreaks and calculates a partially white space invariant checksum.
// The returned value is the hex encoded value of the first 16 byte of a sha3-256 sum. So the result is always 32 byte
// long.
// Checksums of
//    CREATE TABLE book (id BINARY(16))
// should be equal to
//    CREATE TABLE
//         book
//            (id BINARY(16))
func sqlChecksum(statements []string) string {
	sb := &strings.Builder{}
	for _, statement := range statements {
		sb.WriteString(normalizeSQLStatement(statement))
		sb.WriteString(" ")
	}
	digest := sha3.Sum256([]byte(strings.TrimSpace(sb.String())))
	return hex.EncodeToString(digest[:16])
}

// normalizeSQLStatement normalizes the sql string and removes line breaks and leading/trailing whitespaces.
func normalizeSQLStatement(statement string) string {
	sb := &strings.Builder{}
	lines := strings.Split(statement, "\n")
	for _, line := range lines {
		t := strings.TrimSpace(line)
		if t == "" {
			continue
		}
		sb.WriteString(t)
		sb.WriteString(" ")
	}

	return strings.TrimSpace(sb.String())
}

// createSQLMigrationTooling emits all generic bits to handle database migrations
func createSQLMigrationTooling(ctx *genctx, rslv *resolver, file *src.FileBuilder, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {

	tableName := text.Safename(bc.Name()) + "_migration_schema_history"

	file.PutNamedImport("time", "time")
	file.PutNamedImport("database/sql", "sql")

	file.AddTypes(
		src.NewStruct("migration").
			SetDoc("...represents a single isolated database migration, which may consist of multiple statements but which must be executed atomically.").
			AddFields(
				src.NewField("Version", src.NewTypeDecl("int64")).SetDoc("...is the unix timestamp in seconds, at which this migration was defined."),
				src.NewField("Description", src.NewTypeDecl("string")).SetDoc("...describes at least why the migration is needed."),
				src.NewField("Statements", src.NewSliceDecl(src.NewTypeDecl("string"))).SetDoc("...contains e.g. CREATE, ALTER or DROP statements to apply."),
				src.NewField("File", src.NewTypeDecl("string")).SetDoc("...is the file path indicating the origin of the statements."),
				src.NewField("Line", src.NewTypeDecl("int32")).SetDoc("...is the line number indicating the origin of the statements."),
				src.NewField("Checksum", src.NewTypeDecl("string")).SetDoc("...is the hex encoded first 16 byte sha3-256 checksum of all trimmed statements."),
			).AddMethodToJson("String", true, false, true).
			AddMethods(
				src.NewFunc("Apply").SetDoc("... executes the statements.").
					AddParams(src.NewParameter("db", src.NewTypeDecl("DBTX"))).
					AddResults(src.NewParameter("", src.NewTypeDecl("error"))).
					AddBody(src.NewBlock().
						AddLine("for _, s := range m.Statements {").
						AddLine("_, err := db.ExecContext(", src.NewTypeDecl("context.Background"), "(), s)").
						Check("err", "cannot ExecContext").
						AddLine("}").
						AddLine("return nil"),
					),
			),
		src.NewStruct("migrationEntry").
			SetDoc("...represents a row entry from the migration schema history table.").
			AddFields(
				src.NewField("Version", src.NewTypeDecl("int64")).SetDoc("...is the unix timestamp in seconds, at which this migration was defined."),
				src.NewField("File", src.NewTypeDecl("string")).SetDoc("...is the file path indicating the origin of the statements."),
				src.NewField("Line", src.NewTypeDecl("int32")).SetDoc("...is the line number indicating the origin of the statements."),
				src.NewField("Checksum", src.NewTypeDecl("string")).SetDoc("...is the hex encoded first 16 byte sha3-256 checksum of all trimmed statements."),
				src.NewField("AppliedAt", src.NewTypeDecl("int64")).SetDoc("...is the unix timestamp in seconds when this migration has been applied."),
				src.NewField("ExecutionDuration", src.NewTypeDecl("int64")).SetDoc("...is the amount of nanoseconds which were needed to apply this migration."),
				src.NewField("Description", src.NewTypeDecl("string")).SetDoc("...describes at least why the migration is needed."),
				src.NewField("Status", src.NewTypeDecl("string")).SetDoc("...is the status of the migration"),
			).AddMethodToJson("String", true, false, true).
			AddMethods(
				src.NewFunc("insert").SetDoc("...writes a migrationEntry into the history table.").
					AddParams(
						src.NewParameter("db", src.NewTypeDecl("DBTX")),
					).
					AddResults(src.NewParameter("", src.NewTypeDecl("error"))).
					AddBody(src.NewBlock().
						AddLine(`const q = "INSERT INTO `+tableName+`(version, file, line, checksum, applied_at, execution_duration, description, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"`).
						AddLine("_, err := db.ExecContext(", src.NewTypeDecl("context.Background"), "(), q, m.Version, m.File, m.Line, m.Checksum, m.AppliedAt, m.ExecutionDuration, m.Description, m.Status)").
						Check("err", "cannot ExecContext").
						NewLine().
						AddLine("return nil"),
					),
				src.NewFunc("update").SetDoc("...writes a migrationEntry into the history table and replaces the exiting entry identified by version.").
					AddParams(
						src.NewParameter("db", src.NewTypeDecl("DBTX")),
					).
					AddResults(src.NewParameter("", src.NewTypeDecl("error"))).
					AddBody(src.NewBlock().
						AddLine(`const q = "UPDATE `+tableName+` SET file = ?, line = ?, checksum = ?, applied_at = ?, execution_duration = ?, description = ?, status = ? WHERE version = ?"`).
						AddLine("_, err := db.ExecContext(", src.NewTypeDecl("context.Background"), "(), q,  m.File, m.Line, m.Checksum, m.AppliedAt, m.ExecutionDuration, m.Description, m.Status, m.Version)").
						Check("err", "cannot ExecContext").
						NewLine().
						AddLine("return nil"),
					),
			),
	)

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

	file.AddFuncs(
		src.NewFunc("readMigrationHistoryTable").
			SetDoc("...reads the entire history into memory, which are only a few bytes.").
			AddParams(src.NewParameter("db", src.NewTypeDecl(rslv.assembleQualifier(rMysql, "DBTX")))).
			AddResults(
				src.NewParameter("", src.NewSliceDecl(src.NewTypeDecl(rslv.assembleQualifier(rMysql, "migrationEntry")))),
				src.NewParameter("", src.NewTypeDecl("error")),
			).AddBody(src.NewBlock().
			AddLine("var res []migrationEntry").
			AddLine("rows, err := db.QueryContext(", src.NewTypeDecl("context.Background"), "(),\"", "SELECT version, file, line, checksum, applied_at, execution_duration, description, status FROM "+tableName+" ORDER BY version ASC\")").
			Check("err", "cannot query history", "nil").
			AddLine("defer rows.Close()").
			AddLine("for rows.Next() {").
			AddLine("var i migrationEntry").
			AddLine("if err := rows.Scan(&i.Version, &i.File, &i.Line, &i.Checksum, &i.AppliedAt, &i.ExecutionDuration, &i.Description, &i.Status);err!=nil{").
			AddLine("return nil, ", src.NewTypeDecl("fmt.Errorf"), "(\"scan failed: %w\",err)").
			AddLine("}").
			AddLine("res = append(res, i)").
			AddLine("}").

			AddLine("err = rows.Close()").
			Check("err", "cannot close rows", "res").
			NewLine().

			AddLine("err = rows.Err()").
			Check("err", "query failed", "res").
			NewLine().

			AddLine("return res, nil"),

		),


		src.NewFunc("Migrate").
			SetDoc("...ensures that the migration history table exists, checks the checksums of all already applied migrations\nand applies all missing migrations in the defined version order.").
			AddParams(src.NewParameter("db", src.NewTypeDecl(rslv.assembleQualifier(rMysql, "DBTX")))).
			AddResults(src.NewParameter("", src.NewTypeDecl("error"))).
			AddBody(src.NewBlock().
				AddLine(`
					// if we can start a transaction on our own, do so and invoke recursively
						if db,ok := db.(*sql.DB);ok{
							tx, err := db.BeginTx(context.Background(),nil)
							if err != nil{
								return fmt.Errorf("cannot begin transaction: %w",err)
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

`).
				AddLine("const q = `", createMigrationTable, "`").
				AddLine("if _,err := db.ExecContext(", src.NewTypeDecl("context.Background"), "(), q); err!=nil {").
				AddLine("return ", src.NewTypeDecl("fmt.Errorf"), "(\"cannot create "+tableName+": %w\", err)").
				AddLine("}").
				NewLine().
				AddLine("history, err := readMigrationHistoryTable(db)").
				Check("err", "cannot read history").
				NewLine().
				AddLine(
					`
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
`).


				AddLine("return nil"),
			),
	)

	return nil
}
