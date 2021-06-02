package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	golang2 "github.com/golangee/src/golang"
	"strings"
)

func RenderModule(dst *ast.Prj, prj *adl.Project, src *adl.Module) error {
	if src.Generator == nil {
		return fmt.Errorf("cannot render a non-target project: %s", src.Name)
	}

	if src.Generator.Go == nil {
		return fmt.Errorf("cannot render a non-go module: %s -> %s", src.Name, src.Generator.OutDir)
	}

	mod := astutil.MkMod(dst, src.Generator.Go.Module.String()).
		SetLang(ast.LangGo).
		SetLangVersion(ast.LangVersionGo16).
		SetOutputDirectory(src.Generator.OutDir.String())

	for _, require := range src.Generator.Go.Requires {
		mod.Require(require.String())
	}

	// domain package
	domainTerm := prj.Glossary.Terms[src.Domain.Name.String()]
	domainRootPkg := golang.MakePkgPath(src.Generator.Go.Module.String(), src.Domain.Name.String(), "domain")
	domain := astutil.MkPkg(mod, domainRootPkg)
	domain.SetComment("...contains the core and usecase packages which represent the " + src.Domain.Name.String() + " domain model.\n" + golang2.DeEllipsis(domainTerm.Ident.String(), domainTerm.Description.String()))
	domain.SetPreamble(makePreamble(src.Preamble))

	// core packages
	coreRootPkg := golang.MakePkgPath(domainRootPkg, "core")
	coreRoot := astutil.MkPkg(mod, coreRootPkg)
	coreRoot.SetComment("...contains the domains primitives like Entities, Values, Repositories, Services (anemic), Events, aggregate roots (non-anemic) or DTOs.")
	coreRoot.SetPreamble(makePreamble(src.Preamble))
	for _, p := range src.Domain.Core {
		if err := renderUserTypes(coreRoot, src, p); err != nil {
			return token.NewPosError(p.Name, "unable to render domain core package").SetCause(err)
		}
	}

	// usecase packages
	usecaseRootPkg := golang.MakePkgPath(domainRootPkg, "usecase")
	usecaseRoot := astutil.MkPkg(mod, usecaseRootPkg)
	usecaseRoot.SetComment("...contains the domains use cases, usually in a service form, which uses an arbitrary composition of the domains primitives.")
	usecaseRoot.SetPreamble(makePreamble(src.Preamble))
	for _, p := range src.Domain.Usecase {
		if err := renderUserTypes(usecaseRoot, src, p); err != nil {
			return token.NewPosError(p.Name, "unable to render domain usecase package").SetCause(err)
		}
	}

	return nil
}

func renderUserTypes(parent *ast.Pkg, srcMod *adl.Module, src *adl.Package) error {
	mod := astutil.Mod(parent)
	pkg := parent
	if src.Name.String() != "" {
		pkg = astutil.MkPkg(mod, golang.MakePkgPath(pkg.Path, src.Name.String()))
		pkg.SetPreamble(makePreamble(srcMod.Preamble))
		pkg.SetComment(src.Comment.String())
	}

	// repos
	if len(src.Repositories) > 0 {
		file := ast.NewFile(strings.ToLower("repositories.go"))
		file.SetPreamble(makePreamble(srcMod.Preamble))
		for _, repository := range src.Repositories {
			iface, err := buildInterface(file, srcMod, src, repository)
			if err != nil {
				return err
			}

			_ = iface
		}
		pkg.AddFiles(file)
	}

	// dtos
	if len(src.DTOs) > 0 {
		file := ast.NewFile(strings.ToLower("dtos.go"))
		file.SetPreamble(makePreamble(srcMod.Preamble))
		for _, dto := range src.DTOs {
			typ, err := buildStruct(file, srcMod, src, dto)
			if err != nil {
				return err
			}

			_ = typ
		}
		pkg.AddFiles(file)
	}

	// services
	if len(src.Services) > 0 {
		file := ast.NewFile(strings.ToLower("services.go"))
		file.SetPreamble(makePreamble(srcMod.Preamble))
		for _, dto := range src.Services {
			_, _, err := buildService(file, srcMod, src, dto)
			if err != nil {
				return err
			}
		}
		pkg.AddFiles(file)
	}

	return nil
}

func buildInterface(parent *ast.File, srcMod *adl.Module, src *adl.Package, iface *adl.Interface) (*ast.Interface, error) {
	aType := ast.NewInterface(iface.Name.String()).SetComment(iface.Comment.String())
	parent.AddTypes(aType)

	for _, method := range iface.Methods {
		aMethod := ast.NewFunc(method.Name.String()).SetComment(method.Comment.String())
		for _, param := range method.In {
			// TODO how to model complex types? TADL AST tree nodes?
			aMethod.AddParams(ast.NewParam(param.Name.String(), ast.NewSimpleTypeDecl(ast.Name(param.TypeName.String()))).SetComment(param.Comment.String()))
		}

		for _, param := range method.Out {
			// TODO how to model complex types? TADL AST tree nodes?
			aMethod.AddResults(ast.NewParam(param.Name.String(), ast.NewSimpleTypeDecl(ast.Name(param.TypeName.String()))).SetComment(param.Comment.String()))
		}

		aType.AddMethods(aMethod)
	}

	return aType, nil
}

func buildStruct(parent *ast.File, srcMod *adl.Module, src *adl.Package, typ *adl.DTO) (*ast.Struct, error) {
	aType := ast.NewStruct(typ.Name.String()).SetComment(typ.Comment.String())
	parent.AddTypes(aType)

	for _, field := range typ.Fields {
		// TODO how to model complex types? TADL AST tree nodes?
		aType.AddFields(ast.NewField(field.Name.String(), ast.NewSimpleTypeDecl(ast.Name(field.TypeName.String()))).SetComment(field.Comment.String()))
	}
	return aType, nil
}

func buildService(parent *ast.File, srcMod *adl.Module, src *adl.Package, typ *adl.Service) (defaultService, service *ast.Struct, _ error) {
	defaultService = ast.NewStruct("Default" + typ.Name.String()).SetComment("...is an implementation stub for " + typ.Name.String() + ".")

	for _, field := range typ.Fields {
		// TODO how to model complex types? TADL AST tree nodes?
		defaultService.AddFields(ast.NewField(field.Name.String(), ast.NewSimpleTypeDecl(ast.Name(field.TypeName.String()))).SetComment(field.Comment.String()))
	}

	service = ast.NewStruct(typ.Name.String()).SetComment(typ.Comment.String()).AddEmbedded(ast.NewSimpleTypeDecl(ast.Name(defaultService.TypeName)))

	parent.AddTypes(defaultService)
	parent.AddTypes(service)

	return defaultService, service, nil
}

func makePreamble(p adl.Preamble) string {
	tmp := ""
	if p.Generator != "" {
		tmp = p.Generator
	}

	if tmp != "" && p.License != "" {
		tmp += "\n\n"
	}

	if p.License != "" {
		tmp += p.License
	}

	return tmp
}
