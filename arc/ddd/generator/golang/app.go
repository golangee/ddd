package golang

import (
	"fmt"
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
	pkgCore        = "core"
	pkgUsecase     = "usecase"
	pkgInternalApp = "internal/application"
)

func renderApps(dst *ast.Mod, src *adl.Module) error {
	if len(src.Executables) > 0 {
		cmd := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, pkgInternalApp))
		cmd.SetComment("...contains individual applications and dependency injection layers for each executable.")
		cmd.SetPreamble(makePreamble(src.Preamble))

		for _, executable := range src.Executables {
			cmdPkg := astutil.MkPkg(dst, getApplicationPath(dst, executable))
			cmdPkg.SetComment("...defines the application and dependency injection layer for the '" + executable.Name.String() + "' executable.\n\n" + executable.Comment.String())
			cmdPkg.SetPreamble(makePreamble(src.Preamble))

			appStub := ast.NewStruct("defaultApplication").SetComment("...embeds an IoC instance to provide a default behavior").SetVisibility(ast.Private)
			app := ast.NewStruct("Application").SetComment("...embeds the defaultApplication to provide the default application behavior.\nIt also provides the inversion of control injection mechanism for all bounded contexts.")

			appStub.AddMethods(
				ast.NewFunc("init").
					SetVisibility(ast.Private).
					AddParams(ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context"))).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetBody(ast.NewBlock(ast.NewReturnStmt(ast.NewIdentLit("nil")))),

				ast.NewFunc("Run").
					AddParams(ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context"))).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetBody(ast.NewBlock(ast.NewReturnStmt(ast.NewIdentLit("nil")))),
			)

			appConst := ast.NewFunc("New"+app.TypeName).
				AddParams(ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context"))).
				AddResults(
					ast.NewParam("", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl("Application"))),
					ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
				).
				SetBody(ast.NewBlock(ast.NewTpl("a := &Application{}\na." + appStub.TypeName + ".self = a\nif err:=a.init(ctx);err!=nil{\nreturn nil, fmt.Errorf(\"cannot init application: %w\",err)\n}\n\nreturn a,nil\n")))

			cmdPkg.AddFiles(
				ast.NewFile("application.go").
					SetPreamble(makePreamble(src.Preamble)).
					AddNodes(app, appConst, appStub),
			)

			app.AddEmbedded(ast.NewSimpleTypeDecl(ast.Name(astutil.FullQualifiedName(appStub))))
			appStub.AddFields(ast.NewField("self", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl(ast.Name(astutil.FullQualifiedName(app))))).SetVisibility(ast.Private).SetComment("...provides a pointer to the actual Application instance to provide\none level of a quasi-vtable calling indirection for simple method 'overriding'."))
			appStub.SetComment("...aggregates all contained bounded contexts and starts their driver adapters.")
			appStub.SetDefaultRecName(strings.ToLower(appStub.TypeName)[:1])

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
					if err := makeServiceGetter(appStub, service); err != nil {
						return fmt.Errorf("cannot create service: %w", err)
					}
				}

				// the domain use cases
				usecaseServices := findTypes(findPrefixPkgs(dst, golang.MakePkgPath(path.String(), pkgUsecase)), func(s stereotype.Struct) bool {
					return s.IsService()
				})

				for _, service := range usecaseServices {
					if err := makeServiceGetter(appStub, service); err != nil {
						return fmt.Errorf("cannot create service: %w", err)
					}
				}
			}

		}
	}

	return nil
}

func getApplicationPath(mod *ast.Mod, exec *adl.Executable) string {
	return golang.MakePkgPath(mod.Name, pkgInternalApp, golang2.MakeIdentifier(exec.Name.String()))
}

func makeGetter(app *ast.Struct, typ ast.TypeDecl) (*ast.Func, error) {
	if _, ok := typ.(*ast.SimpleTypeDecl); !ok {
		return nil, fmt.Errorf("type must be a SimpleTypeDecl: " + typ.String())
	}

	funName := "get" + golang.GlobalFlatName2(typ)

	var fun *ast.Func
	for _, f := range app.Methods() {
		if f.FunName == funName {
			fun = f
			break
		}
	}

	if fun != nil {
		return fun, nil
	}

	fun = ast.NewFunc(funName).
		SetVisibility(ast.Private).
		SetRecName(app.DefaultRecName).
		SetPtrReceiver(true).
		AddResults(
			ast.NewParam("", typ.Clone()),
			ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
		)

	body := ast.NewBlock()
	fun.SetBody(body)
	app.AddMethods(fun)

	resolvedType := astutil.Resolve(app, typ.String())
	if resolvedType == nil {
		body.Add(ast.NewTpl("panic(\"not implemented\")"))
		return fun, nil
	}

	switch t := resolvedType.(type) {
	case *ast.Struct:
		body.Add(ast.NewTpl(`panic("assemble super config, parse that once and then poke from that")`))
	case *ast.Interface:
		body.Add(ast.NewTpl(`panic("find different implementations and make them configurable, e.g. mysql vs postgres")`))
	default:
		return nil, fmt.Errorf("unsupported resolved getter injection type: %v", t)
	}

	return fun, nil
}

func makeServiceGetter(app, service *ast.Struct) error {
	getter := ast.NewFunc("get"+golang.GlobalFlatName(service)).
		SetRecName(app.DefaultRecName).
		SetPtrReceiver(true).
		AddResults(
			ast.NewParam("", ast.NewTypeDeclPtr(astutil.TypeDecl(service))),
			ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
		).SetVisibility(ast.Private)

	serviceField := ast.NewField(golang.MakePrivate(golang.GlobalFlatName(service)), ast.NewTypeDeclPtr(astutil.TypeDecl(service))).
		SetVisibility(ast.Private)

	body := ast.NewBlock()
	body.Add(
		ast.NewIfStmt(
			ast.NewBinaryExpr(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName)), ast.OpNotEqual, ast.NewIdentLit("nil")),
			ast.NewBlock(
				ast.NewReturnStmt(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName)), ast.NewIdentLit("nil"))),
		),
		lang.Term(),
	)

	factory := service.FactoryRefs[0] // always expecting at least one factory
	factoryFQN := ast.Name(ast.Name(astutil.FullQualifiedName(service)).Qualifier() + "." + factory.FunName)
	var callIdents []ast.Expr
	for _, param := range factory.Params() {
		paramGetter, err := makeGetter(app, param.TypeDecl())
		if err != nil {
			return fmt.Errorf("invalid service parameter: %w", err)
		}

		callParamGetter := ast.NewCallExpr(ast.NewSelExpr(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent("self")), ast.NewIdent(paramGetter.FunName)))
		body.Add(lang.TryDefine(ast.NewIdent(param.ParamName), callParamGetter, "cannot get parameter '"+param.ParamName+"'"))

		callIdents = append(callIdents, ast.NewIdent(param.ParamName))
	}

	body.Add(lang.Term())
	body.Add(lang.TryDefine(ast.NewIdentLit("s"), lang.CallStatic(factoryFQN, callIdents...), "cannot create service '"+service.TypeName+"'"))
	body.Add(lang.Term())
	body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(getter.RecName()), ast.NewIdent(serviceField.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewIdent("s"))))
	body.Add(lang.Term(), lang.Term())
	body.Add(ast.NewReturnStmt(ast.NewIdent("s"), ast.NewIdentLit("nil")))
	getter.SetBody(body)

	app.AddFields(serviceField)
	app.AddMethods(getter)

	return nil
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
