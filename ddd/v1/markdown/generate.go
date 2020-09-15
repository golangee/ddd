package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/architecture/ddd/v1/validation"
	"github.com/golangee/plantuml"
	"reflect"
	"strconv"
	"strings"
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

	if err := generateLayers(ctx); err != nil {
		return err
	}

	return generateCommandlineArgs(ctx)
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

				for range l.Implementations() {
					factoryFuncs++
				}

				md.H4("The domains core layer").
					Printf("The core layer or API layer of the domain consists of %d data types,\n", dataTypes).
					Printf("%d service or SPI interfaces and %d actual service implementations.\n\n", ifaceTypes, factoryFuncs)

				// returned types from factories are API types, everything else is SPI
				apiIfaceFactory := make(map[string]string)
				for _, impl := range l.Implementations() {
					apiIfaceFactory[impl.Of()] = ""
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

				for _, impl := range l.Implementations() {
					md.H5("Factory *" + impl.Of() + "*")
					md.Print("The API factory method *" + impl.Of() + "Factory* creates an instance.\n")
					if len(impl.Requires()) > 0 {
						md.Print("It requires the interfaces *" + strings.Join(impl.Requires(), "*, *") + "* as dependencies.\n")
					}

					if len(impl.Options().Fields()) > 0 {
						md.Print("The instance must be configured using the following options:\n")
						for _, field := range impl.Options().Fields() {
							md.Print(" * " + field.Name() + " (" + field.Comment() + ")\n")
						}
					}
					md.Print("\n")
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

func generateCommandlineArgs(ctx *genctx) error {
	binName := text.Safename(ctx.spec.Name())
	md := ctx.markdown(mainMarkdown)
	md.H2("usage")
	md.P("The application can be launched from the command line. One can display any available options using the *-help* flag:")
	md.Code("bash", binName+" -help")

	flagCount := 0
	for _, bc := range ctx.spec.BoundedContexts() {
		for _, layer := range bc.Layers() {
			_ = layer.Walk(func(obj interface{}) error {
				if spec, ok := obj.(*ddd.ServiceImplSpec); ok {
					for range spec.Options().Fields() {
						flagCount++
					}
				}
				return nil
			})
		}
	}

	if flagCount == 0 {
		md.P("There are not further options available.")
		return nil
	}

	intro := "The application can be configured using the following command line or environment options.\n"
	if flagCount == 1 {
		intro += "Currently, there is only one option.\n"
	} else {
		intro += "Currently, there are " + strconv.Itoa(flagCount) + " options.\n"
	}
	intro += "At first, the default value is loaded into the variable.\n"
	intro += "Afterwards the environment variable is considered and finally the command line argument takes precedence."
	md.P(intro)

	for _, bc := range ctx.spec.BoundedContexts() {
		for _, layer := range bc.Layers() {
			envPrefix := strings.ToUpper(bc.Name() + "." + layer.Name() + ".")
			err := layer.Walk(func(obj interface{}) error {
				if spec, ok := obj.(*ddd.ServiceImplSpec); ok {
					for _, field := range spec.Options().Fields() {
						envName := strings.ReplaceAll(strings.ToUpper(envPrefix+field.Name()), ".", "_")
						cmdName := strings.ReplaceAll(strings.ToLower(envPrefix+field.Name()), ".", "-")

						md.H3(field.Name())
						myDefaultText := field.Default()
						myDefaultValue := field.Default()
						str := "The bounded context *" + bc.Name() + "* declares in the layer *" + layer.Name() + "* the *" + string(field.TypeName()) + "* option **" + field.Name() + "** which " + text.TrimComment(field.Comment()) + "\n"
						if field.Default() == "" {
							str += "The default value is " + field.Default()
							switch field.TypeName() {
							case ddd.String:
								myDefaultText = "the empty string"
								myDefaultValue = "\"lorem ipsum\""
							case ddd.Int64:
								myDefaultText = "0 (zero)"
								myDefaultValue = "0"
							case ddd.Float64:
								myDefaultText = "0.0"
								myDefaultValue = "0.0"
							case ddd.Bool:
								myDefaultText = "false"
								myDefaultValue = "false"
							case ddd.Duration:
								myDefaultText = "0s"
								myDefaultValue = "0s"
							default:
								myDefaultText = "the native zero value"
							}
						}

						str += "The default value is " + myDefaultText + ".\n"
						str += "The environment variable *" + envName + "* is evaluated, if present and is only overridden by the command line argument *" + cmdName + "*."
						md.P(str)

						if myDefaultValue != "" {
							md.P("Example")
							tmp := "export " + envName + "=" + myDefaultValue + "\n"
							tmp += binName + " -" + cmdName + "=" + myDefaultValue
							md.Code("bash", tmp)
						}

					}
				}

				return nil
			})

			if err != nil {
				return err
			}
		}

	}

	return nil
}
