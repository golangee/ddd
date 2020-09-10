package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const (
	pkgNameCore    = "core"
	pkgNameUseCase = "usecase"
	pkgNameRest    = "rest"
)

func generateCmdSrv(ctx *genctx) error {
	ctx.newFile("cmd/"+safename(ctx.spec.Name()+"srv"), "main", "main").
		SetPackageDoc("Package main contains the executable to launch the actual " + ctx.spec.Name() + " server process.").
		AddFuncs(
			src.NewFunc("main"),
		)

	return nil
}

func generateLayers(ctx *genctx) error {
	for _, bc := range ctx.spec.BoundedContexts() {

		rslv := newResolver(ctx.mod.Main().Path, bc)
		bcPath := filepath.Join("internal", safename(bc.Name()))
		ctx.newFile(bcPath, "doc", "").SetPackageDoc(
			"Package " + safename(bc.Name()) + " contains all bounded domain API models, the according use cases and \n" +
				"all other port and adapter implementations.\n\n" + bc.Description(),
		)

		for _, layer := range bc.Layers() {
			switch l := layer.(type) {
			case *ddd.CoreLayerSpec:
				corePath := filepath.Join(bcPath, pkgNameCore)
				ctx.newFile(corePath, "doc", "").SetPackageDoc(l.Description())

				api := ctx.newFile(corePath, "api", "")
				mock := ctx.newFile(corePath, "mock", "")
				for _, structOrInterface := range l.API() {
					switch t := structOrInterface.(type) {
					case *ddd.StructSpec:
						strct, err := generateStruct(rslv, rUniverse|rCore, t)
						if err != nil {
							return fmt.Errorf("core: %w", err)
						}
						api.AddTypes(strct)
					case *ddd.InterfaceSpec:
						iface, err := generateInterface(rslv, rUniverse|rCore, t)

						if err != nil {
							return fmt.Errorf("core: %w", err)
						}
						api.AddTypes(iface)

						if l.IsService(t.Name()) {
							mock.AddTypes(src.ImplementMock(iface))
						}
					default:
						panic("not yet implemented: " + reflect.TypeOf(t).String())
					}
				}

				facs := ctx.newFile(corePath, "factories", "")
				for _, impl := range l.Implementations() {

					fun, opt, err := generateFactoryFunc(strings.ToUpper(bc.Name()+"."+l.Name()+"."), rslv, rUniverse|rCore, impl)
					if err != nil {
						return fmt.Errorf("%s: %w", layer.Name(), err)
					}
					facs.AddTypes(opt)

					facs.AddVars(
						src.NewVar(impl.Of() + "Factory").SetRHS(src.NewBlock(fun)).SetDoc(fun.Doc()),
					)
					fun.SetDoc("")
					fun.SetName("")

				}

			case *ddd.UseCaseLayerSpec:
				usecasePath := filepath.Join(bcPath, pkgNameUseCase)
				ctx.newFile(usecasePath, "doc", "").SetPackageDoc(l.Description())

				for _, useCase := range l.UseCases() {

					api := ctx.newFile(usecasePath, strings.ToLower(useCase.Name()), "")
					uFace := src.NewInterface(useCase.Name())
					myDoc := useCase.Comment()
					myDoc += "\n\nThe following user stories are covered:\n\n"
					api.AddTypes(uFace)

					for _, story := range useCase.Stories() {
						fun, err := generateFunc(rslv, rUniverse|rCore|rUsecase, story.Func())
						if err != nil {
							return fmt.Errorf("%s: %w", layer.Name(), err)
						}
						uFace.AddMethods(fun)
						myDoc += "  * " + story.Story() + "\n"

						for _, strct := range story.Structs() {
							s, err := generateStruct(rslv, rUniverse|rCore|rUsecase, strct)
							if err != nil {
								return fmt.Errorf("%s: %w", layer.Name(), err)
							}
							api.AddTypes(s)
						}
					}

					uFace.SetDoc(myDoc)
				}
			case *ddd.RestLayerSpec:
				if err := createRestLayer(ctx, rslv, bc, l); err != nil {
					return fmt.Errorf("%s: %w", layer.Name(), err)
				}
			default:
				panic("not yet implemented: " + reflect.TypeOf(l).String())
			}

		}

	}

	return nil
}

func generateFactoryFunc(envPrefix string, rslv *resolver, scopes resolverScope, impl *ddd.ServiceImplSpec) (*src.FuncBuilder, *src.TypeBuilder, error) {
	strct, err := generateStruct(rslv, rUniverse|rCore, impl.Options())
	if err != nil {
		return nil, nil, err
	}

	strct.AddMethodToJson("String", true, true, true)
	strct.AddMethodFromJson("Parse")
	if err := generateEnvFlagsConfigure(envPrefix, "ParseEnv", strct, impl.Options()); err != nil {
		return nil, nil, err
	}

	var in ddd.InParams
	in = append(in, ddd.Var("opts", ddd.TypeName(strct.Name()), "... contains the options to create the instance."))
	for _, inj := range impl.Requires() {
		in = append(in, ddd.Var(makePackagePrivate(inj), ddd.TypeName(inj), "... is a non-nil instance."))
	}

	fun := ddd.Func(impl.Of()+"Factory", "... is the factory function to create a new instance of "+impl.Of()+".", in,
		ddd.Out(
			ddd.Return(ddd.TypeName(impl.Of()), "...is the new instance or nil in case of an error."),
			ddd.Err(),
		),
	)

	f, err := generateFunc(rslv, scopes, fun)
	if err != nil {
		return nil, nil, err
	}

	f.AddBody(src.NewBlock("return &", impl.Of(), "Mock", "{}, nil"))
	return f, strct, nil
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
		myComment += text.TrimComment(param.Comment())
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
		myComment += text.TrimComment(param.Comment())
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
		var jsonTags []string
		myName := makePackagePrivate(f.Name())
		if field.JsonName() != "" {
			myName = field.JsonName()
		}
		jsonTags = append(jsonTags, myName)
		if field.Optional() {
			jsonTags = append(jsonTags, "omitempty")
		}
		f.AddTag("json", jsonTags...)

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

func generateEnvFlagsConfigure(envPrefix, name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	rec.AddMethods(fun)
	body := src.NewBlock()
	comment := "... parses the environment variables for the following flags:\n"
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		envName := envPrefix + strings.ToUpper(field.Name())
		fieldComment := field.Name() + " " + text.TrimComment(field.Doc())
		comment += " * " + envName + " as " + string(oType.TypeName()) + ". The default value is '" + oType.Default() + "'. " + fieldComment + "\n"
		switch oType.TypeName() {
		case "string":
			literal := oType.Default()
			if literal == "" {
				literal = `""`
			}
			body.AddLine(src.NewTypeDecl("flag.StringVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", literal, ",", strconv.Quote(fieldComment), ")")
		case "bool":
			literal := oType.Default()
			if literal == "" {
				literal = `false`
			}
			body.AddLine(src.NewTypeDecl("flag.BoolVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", literal, ",", strconv.Quote(fieldComment), ")")
		default:
			return buildErr("TypeName", string(oType.TypeName()), "is not supported to be parsed from an environment variable", oType.Pos())
		}
	}
	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}
