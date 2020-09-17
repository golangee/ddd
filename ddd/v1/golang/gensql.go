package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"strings"
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

	opts := createMySQLOptions(rslv, bc)
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

func createMySQLOptions(rslv *resolver, bc *ddd.BoundedContextSpec) *src.TypeBuilder {
	opt := ddd.Struct("Options",

		"...contains the connection options for a MySQL database.",
		ddd.Field("Port", ddd.Int64, "...is the database port to connect.").SetDefault("3306"),
		ddd.Field("User", ddd.String, "...is the database user.").SetDefault("\"root\""),
		ddd.Field("Password", ddd.String, "...is the database user password.").SetDefault(""),
		ddd.Field("Protocol", ddd.String, "...is the protocol to use.").SetDefault("\"tcp\""),
		ddd.Field("Database", ddd.String, "...is the database name."),
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

	for k, v := range options {
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
