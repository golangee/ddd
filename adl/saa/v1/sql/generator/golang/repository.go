package golang

import (
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/astutil"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
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

		file.Name = simpleLowerCaseName(repo.TypeName)

		implTypeName := corego.MakePublic(string(src.Dialect) + repo.TypeName + "Impl")
		file.AddTypes(ast.NewStruct(implTypeName).
			SetComment(repo.CommentText()))
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
