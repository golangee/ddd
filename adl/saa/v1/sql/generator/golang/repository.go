package golang

import (
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/astutil"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib/lang"
	"strings"
)

func renderRepositories(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()

	for _, repoTypeName := range src.Repositories {
		file := corego.MkFile(dst, modName, pkgName, "tmp.go")
		rNode := astutil.Resolve(file, repoTypeName.String())
		repo, ok := rNode.(*ast.Interface)
		if !ok {
			return fmt.Errorf("unable to resolve required interface type " + repoTypeName.String())
		}

		file.Name = simpleLowerCaseName(repo.TypeName) + ".go"

		implTypeName := corego.MakePublic(string(src.Dialect) + repo.TypeName + "Impl")

		stub := ast.NewStruct(corego.MakePrivate("Abstract" + repo.TypeName)).
			SetVisibility(ast.PackagePrivate).
			SetComment("...provides function stubs for the interface\n" + repoTypeName.String() + ".")

		stub.AddMethods(
			ast.NewFunc("init").
				SetVisibility(ast.PackagePrivate).
				SetComment("...is called after construction from the factory method.").
				SetBody(ast.NewBlock()),
		)
		corego.ImplementFunctions(repo, stub)

		for _, f := range stub.Methods() {
			f.SetComment(f.CommentText() + "\nOverride this method in the embedding type in another file.")
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
