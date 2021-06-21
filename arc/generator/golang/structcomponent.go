package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/golang"
	"github.com/golangee/src/stdlib"
	"strings"
)

// AddComponent transpiles the given component and may optionally create multiple additional helper types, like
// default implementations and configuration types.
func AddComponent(parent *ast.File, compo *adl.Struct) (component *ast.Struct, _ error) {
	requiresInitStub := false
	for _, method := range compo.Methods {
		if method.StubDefault {
			requiresInitStub = true
			break
		}
	}

	if len(compo.Inject) > 0 {
		requiresInitStub = true
	}

	component = ast.NewStruct(compo.Name.String()).SetComment(compo.Comment.String() + "\n\nThe stereotype of this type is '" + compo.Stereotype.String() + "'.")
	parent.AddTypes(component)

	if requiresInitStub {
		shortName := strings.ToLower(compo.Name.String()[0:1])
		c := ast.NewFunc("New"+golang.MakePublic(compo.Name.String())).SetComment("...allocates and initializes a new "+compo.Name.String()+" instance.").
			AddResults(
				ast.NewParam("", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl(ast.Name(parent.Pkg().Path+"."+compo.Name.String())))),
				ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)),
			)

		injectFieldAssigns := ""
		for _, injection := range compo.Inject {
			p := ast.NewParam(injection.Name.String(), astutil.MakeTypeDecl(injection.Type))
			if injection.Comment.String() != "" {
				p.SetComment(injection.Comment.String())
			}

			c.AddParams(p)
			injectFieldAssigns += shortName + "." + MakePrivate(injection.Name.String()) + "=" + injection.Name.String() + "\n"
		}

		c.SetBody(ast.NewBlock(ast.NewTpl(shortName + " := &" + compo.Name.String() + "{}\n" + injectFieldAssigns + "\nif err := " + shortName + ".init(); err != nil {\nreturn nil, {{.Use \"fmt.Errorf\"}}(\"cannot initialize '" + compo.Name.String() + "': %w\",err)}\n\n return " + shortName + ",nil\n")))

		component.AddFactoryRefs(c)
		parent.AddNodes(c)
	}

	if requiresInitStub {
		defaultComponent := ast.NewStruct(golang.MakePrivate("Default" + compo.Name.String())).
			SetComment("...is an implementation stub for " + compo.Name.String() + ".\nThe sole purpose of this type is to mock the method contract and each method should be shadowed\nby a concrete implementation.").
			SetVisibility(ast.Private)

		defaultComponent.AddMethods(ast.NewFunc("init").
			AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
			SetVisibility(ast.Private).
			SetComment("...is invoked from the constructor/factory function to setup any pre-variants.\nShadow this method as required.").
			SetBody(ast.NewBlock(ast.NewReturnStmt(ast.NewIdentLit("nil")))))

		component.AddEmbedded(ast.NewSimpleTypeDecl(ast.Name(defaultComponent.TypeName)))

		for _, method := range compo.Methods {
			aMethod := ast.NewFunc(method.Name.String()).SetComment(method.Comment.String() + "\nShadow this method as required.")
			for _, param := range method.In {
				aMethod.AddParams(ast.NewParam(param.Name.String(), astutil.MakeTypeDecl(param.Type)).SetComment(param.Comment.String()))
			}

			for _, param := range method.Out {
				aMethod.AddResults(ast.NewParam(param.Name.String(), astutil.MakeTypeDecl(param.Type)).SetComment(param.Comment.String()))
			}

			aMethod.SetBody(ast.NewBlock(ast.NewTpl(`panic("not yet implemented")`)))
			defaultComponent.AddMethods(aMethod)

		}

		parent.AddTypes(defaultComponent)
	}

	for _, decl := range compo.Inject {
		f := ast.NewField(MakePrivate(decl.Name.String()),
			astutil.MakeTypeDecl(decl.Type)).
			SetComment(decl.Comment.String()).
			SetVisibility(ast.Private)

		if decl.Comment.String() == "" {
			switch decl.Stereotype.String() {
			case adl.Cfg:
				f.SetComment("...is the components configuration and injected at construction time.")
			default:
				if decl.Stereotype.String() != "" {
					f.SetComment("...is the components '" + decl.Stereotype.String() + "' and injected at construction time.")
				}
			}
		}

		component.AddFields(f)
	}

	for _, field := range compo.Fields {
		f := ast.NewField(field.Name.String(), astutil.MakeTypeDecl(field.Type)).SetComment(field.Comment.String())
		if field.Private {
			f.SetVisibility(ast.Private)
		}

		component.AddFields(f)
	}

	return component, nil
}
