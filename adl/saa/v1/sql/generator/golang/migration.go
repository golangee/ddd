package golang

import (
	"encoding/hex"
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/stereotype"
	"github.com/golangee/src/golang"
	"github.com/golangee/src/stdlib"
	"github.com/pingcap/parser"
	_ "github.com/pingcap/parser/test_driver"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"

	"github.com/golangee/architecture/adl/saa/v1/astutil"
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
)

const (
	filenameMigrations = "migrations.gen.go"
)

func RenderMigrations(dst *ast.Prj, dialect sql.Dialect, src []*sql.Migration) error {
	if len(src) == 0 {
		return nil
	}

	mod := astutil.MkMod(dst, src[0].Mod.String())
	mod.SetLang(ast.LangGo)
	pkg := astutil.MkPkg(mod, src[0].Pkg.String())
	file := astutil.MkFile(pkg, filenameMigrations)

	migrationConstBlock := ast.NewConstDecl()
	migrationHashConstBlock := ast.NewConstDecl()
	for _, migration := range src {
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
	_, err := createMySQLOptions(file, dialect, "defaultName")
	if err != nil {
		return fmt.Errorf("unable to create mysql options: %w", err)
	}

	return nil
}

func createMySQLOptions(file *ast.File, dialect sql.Dialect, defaultDBName string) (*ast.Struct, error) {
	opt := ast.NewStruct("Options").
		SetComment("...contains the connection options for a MySQL database.").
		AddFields(
			ast.NewField("Port", ast.NewSimpleTypeDecl(stdlib.Int)).
				SetComment("...is the database port to connect.").
				SetDefault(ast.NewIntLit(3306)),
			ast.NewField("User", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is the database user.").
				SetDefault(ast.NewStrLit("")),
			ast.NewField("Password", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is the database user password.").
				SetDefault(ast.NewStrLit("")),
			ast.NewField("Protocol", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is the protocol to use.").
				SetDefault(ast.NewStrLit("tcp")),
			ast.NewField("Database", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is the database name.").
				SetDefault(ast.NewStrLit(defaultDBName)),
			ast.NewField("Address", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is the host or path to socket or host.").
				SetDefault(ast.NewStrLit("localhost")),

			// see https://stackoverflow.com/questions/766809/whats-the-difference-between-utf8-general-ci-and-utf8-unicode-ci/766996#766996
			// https://www.percona.com/live/e17/sites/default/files/slides/Collations%20in%20MySQL%208.0.pdf
			//
			// we enforce correct unicode support for mysql and index/sorting collations. For mysql 8.0 using
			// accent insensitive/case insensitive Unicode 9 support utf8mb4_0900_ai_ci would be better but not compatible
			// with mariadb, so we use a fixed older version for reproducibility across different database servers.
			ast.NewField("Collation", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...declares the connections default collation for sorting and indexing.").
				SetDefault(ast.NewStrLit("utf8mb4_unicode_520_ci")),
			ast.NewField("Charset", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...declares the connections default charset encoding for text.").
				SetDefault(ast.NewStrLit("utf8mb4")),
			ast.NewField("MaxAllowedPacket", ast.NewSimpleTypeDecl(stdlib.Int64)).
				SetComment("...is the max packet size in bytes.").
				SetDefault(ast.NewIntLit(4194304)),
			ast.NewField("Timeout", ast.NewSimpleTypeDecl(stdlib.Duration)).
				SetComment("...is the duration until the dial receives a timeout.").
				SetDefault(ast.NewIdentLit("30s")),
			ast.NewField("WriteTimeout", ast.NewSimpleTypeDecl(stdlib.Duration)).
				SetComment("...is the duration for the write timeout.").
				SetDefault(ast.NewIdentLit("30s")),
			ast.NewField("Tls", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...configures connection security. Valid values are true, false, skip-verify or preferred.").
				SetDefault(ast.NewBoolLit(false)),
			ast.NewField("SqlMode", ast.NewSimpleTypeDecl(stdlib.String)).
				SetComment("...is a flag which influences the sql parser.").
				SetDefault(ast.NewStrLit("ANSI")),
			ast.NewField("ConnMaxLifetime", ast.NewSimpleTypeDecl(stdlib.Duration)).
				SetComment("...is the duration of how long pooled connections are kept alive.").
				SetDefault(ast.NewIdentLit("3m")),
			ast.NewField("MaxOpenConns", ast.NewSimpleTypeDecl(stdlib.Int64)).
				SetComment("...is the amount of how many open connections can be kept in the pool.").
				SetDefault(ast.NewIntLit(25)),
			ast.NewField("MaxIdleConns", ast.NewSimpleTypeDecl(stdlib.Int64)).
				SetComment("...is the amount of how many open connections can be idle.").
				SetDefault(ast.NewIntLit(25)),
			ast.NewField("Test", ast.NewSimpleTypeDecl(stdlib.Bool)).SetDefault(ast.NewBoolLit(true)),
			ast.NewField("Test2", ast.NewSimpleTypeDecl(stdlib.Float64)).SetDefault(ast.NewBasicLit(ast.TokenFloat, "3.41")),
		)

	file.AddNodes(opt) // add it early, functions may need contextual information like package path

	opt.DefaultRecName = strings.ToLower(opt.TypeName[:1])

	if _, err := corego.AddResetFunc(opt); err != nil {
		return nil, fmt.Errorf("unable to add reset func: %w", err)
	}

	stereotype.Put(opt, stereotype.ConfigureStruct, stereotype.Database, stereotype.MySQL)

	addDSNFunc(opt)

	if _, err := corego.AddParseEnvFunc(string(dialect), opt); err != nil {
		return nil, fmt.Errorf("unable to add env parser func: %w", err)
	}

	if _, err := corego.AddParseFlagFunc(string(dialect), opt); err != nil {
		return nil, fmt.Errorf("unable to add flag parser func: %w", err)
	}
	/*


		envPrefix := "MySQL." + bc.Name() + "."

		if err := generateParseEnv(envPrefix, "ParseEnv", genOpt, opt); err != nil {
			panic("illegal state: " + err.Error())
		}

		if err := generateFlagsConfigure(envPrefix, "ConfigureFlags", genOpt, opt); err != nil {
			panic("illegal state: " + err.Error())
		}

		genOpt.AddMethodToJson("String", true, true, true)
		genOpt.AddMethodFromJson("Parse")



		return genOpt

	*/

	return opt, nil
}

func RenderRepository(dst *ast.Prj, src *sql.RepoImplSpec) error {
	//mod := astutil.MkMod(dst, src.Parent.Parent.Name.String())
	//pkg := astutil.MkPkg(mod, src.Parent.Path.String())

	//	_ = pkg

	return nil
}
