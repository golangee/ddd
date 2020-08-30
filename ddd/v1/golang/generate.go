package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/plantuml"
	"github.com/golangee/src"
	"path/filepath"
	"reflect"
)

const (
	pkgNameCore  = "core"
	mainMarkdown = "README.md"
)

func generateCmdSrv(ctx *genctx) error {
	md := ctx.markdown(mainMarkdown).
		H1(ctx.spec.Name()).
		P(ctx.spec.Description()).
		H2("Index").
		TOC().
		H2("Architecture")

	md.Println("The server is organized after the domain driven design principles.")
	if len(ctx.spec.BoundedContexts()) == 1 {
		md.Println("However, it currently consists only of exact one bounded context.")
	} else {
		md.Printf("It is separated into the following %d bounded contexts.\n\n", len(ctx.spec.BoundedContexts()))
	}

	ctx.newFile("cmd/"+safename(ctx.spec.Name()+"srv"), "main", "main").
		SetPackageDoc("Package main contains the executable to launch the actual " + ctx.spec.Name() + " server process.").
		AddFuncs(
			src.NewFunc("main"),
		)

	return nil
}

func generateLayers(ctx *genctx) error {
	for _, bc := range ctx.spec.BoundedContexts() {
		md := ctx.markdown(mainMarkdown).
			H3("The context *" + bc.Name() + "*").
			P(bc.Description())

		rslv := newResolver(ctx.mod.Main().Path, bc)
		bcPath := filepath.Join("internal", safename(bc.Name()))
		ctx.newFile(bcPath, "doc", "").SetPackageDoc(
			"Package " + safename(bc.Name()) + " contains all bounded domain API models, the according use cases and \n" +
				"all other port and adapter implementations.\n\n" + bc.Description(),
		)

		for _, layer := range bc.Layers() {
			switch l := layer.(type) {
			case *ddd.CoreLayerSpec:
				dataTypes := 0
				ifaceTypes := 0
				factoryFuncs := 0

				corePath := filepath.Join(bcPath, pkgNameCore)
				ctx.newFile(corePath, "doc", "").SetPackageDoc(l.Description())

				var uml []plantuml.Renderable
				api := ctx.newFile(corePath, "api", "")
				for _, structOrInterface := range l.API() {
					switch t := structOrInterface.(type) {
					case *ddd.StructSpec:
						strct, err := generateStruct(rslv, rUniverse|rCore, t)
						if err != nil {
							return fmt.Errorf("core: %w", err)
						}
						api.AddTypes(strct)
						uml = append(uml, generateUML(strct))
						dataTypes++
					case *ddd.InterfaceSpec:
						iface, err := generateInterface(rslv, rUniverse|rCore, t)
						if err != nil {
							return fmt.Errorf("core: %w", err)
						}
						api.AddTypes(iface)
						uml = append(uml, generateUML(iface))
						ifaceTypes++
					default:
						panic("not yet implemented: " + reflect.TypeOf(t).String())
					}
				}

				facs := ctx.newFile(corePath, "factories", "")
				for _, funcOrStruct := range l.Factories() {
					switch t := funcOrStruct.(type) {
					case *ddd.StructSpec:
						strct, err := generateStruct(rslv, rUniverse|rCore, t)
						if err != nil {
							return fmt.Errorf("%s: %w", layer.Name(), err)
						}
						facs.AddTypes(strct)
					case *ddd.FuncSpec:
						fun, err := generateFactoryFunc(rslv, rUniverse|rCore, t)
						if err != nil {
							return fmt.Errorf("%s: %w", layer.Name(), err)
						}
						facs.AddFuncs(fun)
						factoryFuncs++
					default:
						panic("not yet implemented: " + reflect.TypeOf(t).String())
					}
				}

				md.H4("The domains core layer").
					Printf("The core layer or API layer of the domain consists of %d data types,\n", dataTypes).
					Printf("%d service or SPI interfaces and %d actual service implementations.\n\n", ifaceTypes, factoryFuncs)

				// returned types from factories are API types, everything else is SPI
				apiIfaceFactory := make(map[string]string)
				for _, funcOrStruct := range l.Factories() {
					if fun, ok := funcOrStruct.(*ddd.FuncSpec); ok {
						for _, spec := range fun.Out() {
							apiIfaceFactory[string(spec.TypeName())] = ""
						}
					}
				}

				for _, structOrInterface := range l.API() {
					md.H5("Type *" + structOrInterface.Name() + "*")
					switch structOrInterface.(type) {
					case *ddd.StructSpec:
						md.P("The data class *" + structOrInterface.Name() + "* " + trimComment(structOrInterface.Comment()))
					case *ddd.InterfaceSpec:
						_, ok := apiIfaceFactory[structOrInterface.Name()]
						if ok {
							md.P("The API interface *" + structOrInterface.Name() + "* " + trimComment(structOrInterface.Comment()))
						} else {
							md.P("The SPI interface *" + structOrInterface.Name() + "* " + trimComment(structOrInterface.Comment()))
						}
					}

				}

				for _, funcOrStruct := range l.Factories() {
					if fun, ok := funcOrStruct.(*ddd.FuncSpec); ok {
						md.H5("Factory *" + fun.Name() + "*")
						md.P("The API factory method *" + fun.Name() + "* " + trimComment(fun.Comment()))
					}
				}

				md.H4("UML")
				diagram := md.UML(bc.Name() + " core API")
				for _, renderable := range uml {
					diagram.Add(renderable)
				}

			default:
				panic("not yet implemented: " + reflect.TypeOf(l).String())
			}

		}

	}

	return nil
}

func generateUML(t *src.TypeBuilder) *plantuml.Class {
	class := plantuml.NewClass(t.Name())
	for _, field := range t.Fields() {
		class.AddAttrs(plantuml.Attr{
			Visibility: plantuml.Public,
			Abstract:   false,
			Static:     false,
			Name:       field.Name(),
			Type:       umlifyDeclName(field.Type()),
		})
	}

	for _, fun := range t.Methods() {
		pTmp := ""
		for i, p := range fun.Params() {
			pTmp += p.Name() + " " + umlifyDeclName(p.Decl())
			if i < len(fun.Params())-1 {
				pTmp += ","
			}
		}

		rTmp := ""
		for i, p := range fun.Results() {
			rTmp += p.Name() + " " + umlifyDeclName(p.Decl())
			if i < len(fun.Params())-1 {
				rTmp += ","
			}
		}

		class.AddAttrs(plantuml.Attr{
			Visibility: plantuml.Public,
			Abstract:   true,
			Static:     false,
			Name:       fun.Name() + "("+pTmp+")",
			Type:       "("+rTmp+")",
		})
	}

	return class
}

func generateFactoryFunc(rslv *resolver, scopes resolverScope, fun *ddd.FuncSpec) (*src.FuncBuilder, error) {
	f, err := generateFunc(rslv, scopes, fun)
	if err != nil {
		return nil, err
	}

	// TODO it would be great to inspect the entire AST of the package and create the private function for the dev
	block := src.NewBlock().
		AddLine("// this package private function is implemented by the developer").
		Add("return ", makePackagePrivate(fun.Name()), "(")

	for _, p := range f.Params() {
		block.Add(p.Name(), ",")
	}

	block.AddLine(")")

	f.AddBody(block)
	return f, nil
}

func generateFunc(rslv *resolver, scopes resolverScope, fun *ddd.FuncSpec) (*src.FuncBuilder, error) {
	f := src.NewFunc(fun.Name())
	myComment := fun.Comment()

	for _, param := range fun.In() {
		decl, err := rslv.resolveTypeName(scopes, param.TypeName())
		if err != nil {
			return nil, buildErr("typeName", string(param.TypeName()), err.Error(), param.Pos())
		}
		p := src.NewParameter(param.Name(), decl)
		f.AddParams(p)

		myComment += "\n\n"
		myComment += "The parameter '" + param.Name() + "' "
		myComment += trimComment(param.Comment())
	}

	for _, param := range fun.Out() {
		decl, err := rslv.resolveTypeName(scopes, param.TypeName())
		if err != nil {
			return nil, buildErr("typeName", string(param.TypeName()), err.Error(), param.Pos())
		}
		p := src.NewParameter(param.Name(), decl)
		f.AddResults(p)

		myComment += "\n\n"
		myComment += "The result '"
		if param.Name() != "" {
			myComment += param.Name()
		} else {
			myComment += commentifyDeclName(decl)
		}
		myComment += "' "
		myComment += trimComment(param.Comment())
	}

	f.SetDoc(myComment)

	return f, nil
}

func generateStruct(rslv *resolver, scopes resolverScope, t *ddd.StructSpec) (*src.TypeBuilder, error) {
	s := src.NewStruct(t.Name())
	s.SetDoc(t.Comment())
	for _, field := range t.Fields() {
		decl, err := rslv.resolveTypeName(scopes, field.TypeName())
		if err != nil {
			return nil, buildErr("typeName", string(field.TypeName()), err.Error(), field.Pos())
		}

		f := src.NewField(field.Name(), decl)
		f.SetDoc(field.Comment())

		s.AddFields(f)
	}

	return s, nil
}

func generateInterface(rslv *resolver, scopes resolverScope, t *ddd.InterfaceSpec) (*src.TypeBuilder, error) {
	s := src.NewInterface(t.Name())
	s.SetDoc(t.Comment())
	for _, fun := range t.Funcs() {
		f, err := generateFunc(rslv, scopes, fun)
		if err != nil {
			return nil, err
		}

		s.AddMethods(f)
	}

	return s, nil
}
