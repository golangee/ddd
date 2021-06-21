package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	golang2 "github.com/golangee/src/golang"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"strings"
)

const (
	pkgCore    = "core"
	pkgUsecase = "usecase"
)

func renderApps(dst *ast.Mod, src *adl.Module) error {
	if len(src.Executables) > 0 {
		cmd := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "internal/application"))
		cmd.SetComment("...contains individual applications and dependency injection layers for each executable.")
		cmd.SetPreamble(makePreamble(src.Preamble))

		for _, executable := range src.Executables {
			cmdPkg := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "internal", "application", golang2.MakeIdentifier(executable.Name.String())))
			cmdPkg.SetComment("...defines the application and dependency injection layer for the '" + executable.Name.String() + "' executable.\n\n" + executable.Comment.String())
			cmdPkg.SetPreamble(makePreamble(src.Preamble))

			app := ast.NewStruct("Application")
			app.SetComment("...aggregates all contained bounded contexts and starts their driver adapters.")
			for _, path := range executable.BoundedContextPaths {
				bc := astutil.FindPkg(dst, path.String())
				if bc == nil {
					return token.NewPosError(path, "invalid bounded context import path: "+path.String())
				}

				// the domain core
				coreServices := findTypes(findPrefixPkgs(dst, golang.MakePkgPath(path.String(), pkgCore)), func(s stereotype.Struct) bool {
					return s.IsService()
				})

				for _, service := range coreServices {
					makeServiceGetter(app, service)
				}

				// the domain use cases
				usecaseServices := findTypes(findPrefixPkgs(dst, golang.MakePkgPath(path.String(), pkgUsecase)), func(s stereotype.Struct) bool {
					return s.IsService()
				})

				for _, service := range usecaseServices {
					makeServiceGetter(app, service)
				}
			}

			cmdPkg.AddFiles(
				ast.NewFile("application.go").
					SetPreamble(makePreamble(src.Preamble)).
					AddTypes(app),
			)
		}
	}

	return nil
}

func makeServiceGetter(app, service *ast.Struct) {
	getter := ast.NewFunc("get"+golang.GlobalFlatName(service)).
		SetRecName(strings.ToLower(app.TypeName)[:1]).
		SetPtrReceiver(true).
		AddResults(
			ast.NewParam("", ast.NewTypeDeclPtr(astutil.TypeDecl(service))),
			ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
		).SetVisibility(ast.Private)

	serviceField := ast.NewField(golang.MakePrivate(golang.GlobalFlatName(service)), ast.NewTypeDeclPtr(astutil.TypeDecl(service))).
		SetVisibility(ast.Private)

	factory := service.FactoryRefs[0]
	factoryFQN := ast.Name(ast.Name(astutil.FullQualifiedName(service)).Qualifier() + "." + factory.FunName)
	var callIdents []ast.Expr
	for _, param := range factory.Params() {
		getter.AddParams(ast.NewParam(param.ParamName, param.TypeDecl().Clone()))
		callIdents = append(callIdents, ast.NewIdentLit(param.ParamName))
	}

	body := ast.NewBlock()
	body.Add(ast.NewIfStmt(ast.NewBinaryExpr(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName)), ast.OpNotEqual, ast.NewIdentLit("nil")),
		ast.NewBlock(
			ast.NewReturnStmt(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName)), ast.NewIdentLit("nil")))),
	)
	body.Add(lang.Term())
	body.Add(lang.TryDefine(ast.NewIdentLit("s"), lang.CallStatic(factoryFQN, callIdents...), "cannot create service '"+service.TypeName+"'"))
	body.Add(lang.Term())
	body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewIdent("s"))))
	body.Add(lang.Term(), lang.Term())
	body.Add(ast.NewReturnStmt(ast.NewIdent("s"), ast.NewIdentLit("nil")))
	getter.SetBody(body)

	app.AddFields(serviceField)
	app.AddMethods(getter)
}

// findPrefixPkgs returns all packages using the according prefix.
func findPrefixPkgs(mod *ast.Mod, prefix string) []*ast.Pkg {
	var r []*ast.Pkg
	for _, pkg := range mod.Pkgs {
		if strings.HasPrefix(pkg.Path, prefix) {
			r = append(r, pkg)
		}
	}

	return r
}

// findServices returns all annotated services from the package.
func findTypes(pkg []*ast.Pkg, predicate func(s stereotype.Struct) bool) []*ast.Struct {
	var r []*ast.Struct
	for _, a := range pkg {
		r = append(r, stereotype.PkgFrom(a).FindStructs(predicate)...)
	}

	return r
}
