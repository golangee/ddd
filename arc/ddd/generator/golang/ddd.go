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

	// domain package
	domainTerm := prj.Glossary.Terms[src.Domain.Name.String()]
	domainPkg := golang.MakePkgPath(src.Generator.Go.Module.String(), src.Domain.Name.String())
	domain := astutil.MkPkg(mod, domainPkg)
	domain.SetComment("...contains the core and usecase packages which represent the " + src.Domain.Name.String() + " domain model.\n" + golang2.DeEllipsis(domainTerm.Ident.String(), domainTerm.Description.String()))
	domain.SetPreamble(makePreamble(src.Preamble))

	// core packages
	coreRootPkg := golang.MakePkgPath(domainPkg, "core")
	coreRoot := astutil.MkPkg(mod, coreRootPkg)
	coreRoot.SetComment("...contains the domains primitives like Entities, Values, Repositories, Services (anemic), Events, aggregate roots (non-anemic) or DTOs.")
	coreRoot.SetPreamble(makePreamble(src.Preamble))
	for _, p := range src.Domain.Core {
		if err := renderUserTypes(coreRoot, src, p); err != nil {
			return token.NewPosError(p.Name, "unable to render domain core package").SetCause(err)
		}
	}

	// usecase packages
	usecaseRootPkg := golang.MakePkgPath(domainPkg, "usecase")
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

	for _, repository := range src.Repositories {
		file := ast.NewFile(strings.ToLower(golang2.MakeIdentifier(repository.Name.String())) + ".go")
		file.SetPreamble(makePreamble(srcMod.Preamble))
		pkg.AddFiles(file)
	}

	return nil
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
