package golang

import (
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib/lang"
)

// ImplementFunctions appends all interface methods to the given struct.
// Funcs with errors will return a not-implemented-error, otherwise a panic
// is raised.
func ImplementFunctions(from *ast.Interface, to *ast.Struct) {
	structPkg := astutil.Pkg(to)
	for _, f := range from.Methods() {

		fun := ast.NewFunc(f.FunName).
			SetComment(f.CommentText())

		for _, param := range f.FunParams {
			fun.AddParams(ast.NewParam(param.ParamName, astutil.UseTypeDeclIn(param.ParamTypeDecl, structPkg)))
		}

		for _, param := range f.FunResults {
			fun.AddResults(ast.NewParam(param.ParamName, astutil.UseTypeDeclIn(param.ParamTypeDecl, structPkg)))
		}

		fun.SetBody(ast.NewBlock(
			lang.Panic("not yet implemented: " + fun.FunName),
		))
		to.AddMethods(fun)
	}
}
