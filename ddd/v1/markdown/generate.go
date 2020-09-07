package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/architecture/ddd/v1/validation"
	"github.com/golangee/plantuml"
	"reflect"
	"strconv"
)

const (
	mainMarkdown = "README.md"
)

func generateDocument(ctx *genctx) error {
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

	return generateLayers(ctx)
}

func generateLayers(ctx *genctx) error {
	for _, bc := range ctx.spec.BoundedContexts() {
		md := ctx.markdown(mainMarkdown).
			H3("The context *" + bc.Name() + "*").
			P(bc.Description())

		for _, layer := range bc.Layers() {
			switch l := layer.(type) {
			case *ddd.CoreLayerSpec:
				dataTypes := 0
				ifaceTypes := 0
				factoryFuncs := 0

				var uml []plantuml.Renderable
				for _, structOrInterface := range l.API() {
					switch t := structOrInterface.(type) {
					case *ddd.StructSpec:
						uml = append(uml, generateUMLForStruct(t))
						dataTypes++
					case *ddd.InterfaceSpec:
						uml = append(uml, generateUMLForInterface(t))
						ifaceTypes++
					default:
						panic("not yet implemented: " + reflect.TypeOf(t).String())
					}
				}

				for _, funcOrStruct := range l.Factories() {
					switch t := funcOrStruct.(type) {
					case *ddd.FuncSpec:
						factoryFuncs++
					case *ddd.StructSpec:
						// ignored
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
						md.P("The data class *" + structOrInterface.Name() + "* " + text.TrimComment(structOrInterface.Comment()))
					case *ddd.InterfaceSpec:
						_, ok := apiIfaceFactory[structOrInterface.Name()]
						if ok {
							md.P("The API interface *" + structOrInterface.Name() + "* " + text.TrimComment(structOrInterface.Comment()))
						} else {
							md.P("The SPI interface *" + structOrInterface.Name() + "* " + text.TrimComment(structOrInterface.Comment()))
						}
					}

				}

				for _, funcOrStruct := range l.Factories() {
					if fun, ok := funcOrStruct.(*ddd.FuncSpec); ok {
						md.H5("Factory *" + fun.Name() + "*")
						md.P("The API factory method *" + fun.Name() + "* " + text.TrimComment(fun.Comment()))
					}
				}

				diagram := md.UML(bc.Name() + " core API")
				for _, renderable := range uml {
					diagram.Add(renderable)
				}

			case *ddd.UseCaseLayerSpec:
				md.H4("The use case or application layer")
				if len(l.UseCases()) == 1 {
					md.P("The following use case is defined.")
				} else {
					md.P("The following " + strconv.Itoa(len(l.UseCases())) + " use cases have been identified.")
				}

				for _, useCase := range l.UseCases() {
					var ucMethods []*ddd.FuncSpec

					md.H5(useCase.Name())
					md.P("The use case *" + useCase.Name() + "* " + text.TrimComment(useCase.Comment()) + "\n" +
						"It contains " + strconv.Itoa(len(useCase.Stories())) + " user stories.")

					md.TableHeader("As a/an", "I want to...", "So that...")
					for _, story := range useCase.Stories() {
						storyModel, err := validation.CheckUserStory(story.Story())
						if err != nil {
							panic("illegal state: must validate before")
						}

						md.TableRow(storyModel.Role, storyModel.Goal, storyModel.Reason)

						ucMethods = append(ucMethods, story.Func())
					}
					md.P("")

					// create the use case diagram
					ucDiag := md.UML("use case-" + useCase.Name())
					addUseCaseDiagram(ucDiag, useCase)
					md.P("")

					tmpIface := ddd.Interface(useCase.Name(), "", ucMethods...)
					diagram := md.UML("iface-" + tmpIface.Name())
					diagram.Add(generateUMLForInterface(tmpIface))
				}

			case *ddd.RestLayerSpec:
				addRestAPI(md, l)
			default:
				panic("not yet implemented: " + reflect.TypeOf(l).String())
			}

		}

	}

	return nil
}
