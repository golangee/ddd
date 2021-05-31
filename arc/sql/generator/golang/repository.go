package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib/lang"
	"reflect"
	"strings"
)

func renderRepositories(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()

	for _, repository := range src.Repositories {
		repoTypeName := repository.Implements
		file := golang.MkFile(dst, modName, pkgName, "tmp.go")
		rNode := astutil.Resolve(file, repoTypeName.String())
		repo, ok := rNode.(*ast.Interface)
		if !ok {
			return fmt.Errorf("unable to resolve required interface type " + repoTypeName.String())
		}

		file.Name = simpleLowerCaseName(repo.TypeName) + ".go"

		implTypeName := golang.MakePublic(string(src.Dialect) + repo.TypeName + "Impl")

		stub := ast.NewStruct(golang.MakePrivate("Abstract" + repo.TypeName)).
			SetVisibility(ast.PackagePrivate).
			SetComment("...provides function stubs for the interface\n" + repoTypeName.String() + ".")

		stub.AddMethods(
			ast.NewFunc("init").
				SetVisibility(ast.PackagePrivate).
				SetComment("...is called after construction from the factory method.").
				SetBody(ast.NewBlock()),

			ast.NewFunc("context").
				SetVisibility(ast.PackagePrivate).
				SetComment("...is called to get the significant context.\n"+
					"This is not idiomatic, but it is not fine to clutter domain APIs with I/O details either.\n"+
					"This is a compromise to actually customize timeouts a bit, but yes, this could be better.").
				SetBody(ast.NewBlock(ast.NewTpl(`return {{.Use "context.Background"}}()`))),
		)
		golang.ImplementFunctions(repo, stub)

		for _, f := range stub.Methods() {
			f.SetComment(f.CommentText() + "\nOverride this method in the embedding type in another file.")
		}

		for _, m := range repository.Methods {
			var method *ast.Func
			for _, f := range stub.Methods() {
				if f.FunName == m.Name.String() {
					method = f
					break
				}
			}

			if method == nil {
				return token.NewPosError(m.Name, "refers to a non existing interface method")
			}

			method.SetRecName("r")
			if err := implementBody(method, m); err != nil {
				return fmt.Errorf("cannot implement method %s: %w", m.Name, err)
			}
		}

		file.AddNodes(
			ast.NewTpl("// document and assert interface compatibility.\n// Entities must be imported anyway, so we won't loose modularity.\n"),
			ast.NewTpl(`var _ {{.Use (.Get "repoTypeName")}} = ({{.Get "iface"}})(nil)`).
				Put("repoTypeName", repoTypeName.String()).
				Put("iface", implTypeName),
			lang.Term(),
		)

		file.AddTypes(
			ast.NewStruct(implTypeName).
				SetComment(repo.CommentText()).
				AddFields(
					ast.NewField("db", ast.NewSimpleTypeDecl("DBTX")).SetVisibility(ast.PackagePrivate),
				).
				AddEmbedded(ast.NewSimpleTypeDecl(ast.Name(stub.TypeName))),
		)

		file.AddFuncs(
			ast.NewFunc("New" + implTypeName).
				SetComment("...creates a new repository instance.").
				AddParams(
					ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX")),
				).
				AddResults(
					ast.NewParam("", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl(ast.Name(implTypeName)))),
				).
				SetBody(ast.NewBlock(
					ast.NewTpl(`r := &{{.Get "typename"}}{
										db:db,
									}
									r.{{.Get "stubname"}}.db = db
									r.init()

									return r`).
						Put("typename", implTypeName).
						Put("stubname", stub.TypeName),
				)),

		)

		file.AddNodes(stub)
	}

	return nil
}

func implementBody(fun *ast.Func, method sql.Method) error {
	switch m := method.Mapping.(type) {
	case sql.ExecMany:
		return implementExecMany(fun, method.Query, m)
	case sql.ExecOne:
		return implementExecOne(fun, method.Query, m)
	case sql.QueryOne:
		return implementFindOne(fun, method.Query, m)
	case sql.QueryMany:
		return implementFindMany(fun, method.Query, m)
	default:
		panic("not implemented: " + reflect.TypeOf(m).String())
	}
}

// simpleLowerCaseName returns a lowercase name which just contains a..z, nothing else.
func simpleLowerCaseName(str string) string {
	str = strings.ToLower(str)
	sb := &strings.Builder{}
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
