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
	if err := generateFlagsConfigure(envPrefix, "ConfigureFlags", strct, impl.Options()); err != nil {
		return nil, nil, err
	}

	if err := generateSetDefault("Reset", strct, impl.Options()); err != nil {
		return nil, nil, err
	}

	if err := generateParseEnv(envPrefix, "ParseEnv", strct, impl.Options()); err != nil {
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

// generateSetDefault appends a method to reset each struct value to either their natural default
// or to its configured default.
func generateSetDefault(name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	comment := "...restores this instance to the default state.\n"
	rec.AddMethods(fun)
	body := src.NewBlock()
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		comment += " * The default value of " + oType.Name() + " is '"

		if oType.Default() != "" {
			comment += oType.Default()
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = ", oType.Default())
		} else {
			switch oType.TypeName() {
			case "string":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = \"\"")
			case "byte":
				fallthrough
			case "int8":
				fallthrough
			case "int16":
				fallthrough
			case "int32":
				fallthrough
			case "int64":
				fallthrough
			case "float32":
				fallthrough
			case "float64":
				fallthrough
			case "time.Duration":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = 0")
				comment += "0"
			case "bool":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = false")
				comment += "false"
			default:
				return buildErr("TypeName", string(oType.TypeName()), "has no supported default value. You need to define a default literal by using SetDefault()", oType.Pos())
			}
		}
		comment += "'. \n"
	}

	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}

// generateFlagsConfigure appends a method to setup the go flags package to be parsed by the according struct instance.
// The naming is <a>-<b>-<c> for the flags. On unix, camel case is discouraged, so we have only the alternatives
// of . _ or - and we decided for now, to use -.
func generateFlagsConfigure(envPrefix, name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	rec.AddMethods(fun)
	body := src.NewBlock()
	comment := "... configures the flags to be ready to get evaluated. The default values are taken from the struct at calling time.\nAfter calling, use flag.Parse() to load the values. You can only use it once, otherwise the flag package will panic.\nThe following flags will be tied to this instance:\n"
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		envName := strings.ReplaceAll(strings.ToLower(envPrefix+field.Name()), ".", "-")
		fieldComment := field.Name() + " " + text.TrimComment(field.Doc())
		comment += " * " + field.Name() + " is parsed from flag '" + envName + "'\n"
		switch oType.TypeName() {
		case "string":
			body.AddLine(src.NewTypeDecl("flag.StringVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "bool":
			body.AddLine(src.NewTypeDecl("flag.BoolVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "int64":
			body.AddLine(src.NewTypeDecl("flag.Int64Var"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "float64":
			body.AddLine(src.NewTypeDecl("flag.Float64Var"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "time.Duration":
			body.AddLine(src.NewTypeDecl("flag.DurationVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		default:
			return buildErr("TypeName", string(oType.TypeName()), "is not supported to be parsed as a flag", oType.Pos())
		}
	}
	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}

func generateParseEnv(envPrefix, name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	fun.AddResults(src.NewParameter("", src.NewTypeDecl("error")))
	rec.AddMethods(fun)
	body := src.NewBlock()

	comment := "... tries to parse the environment variables into this instance. It will only set those values, which have been actually defined. If values cannot be parsed, an error is returned.\n"
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		envName := strings.ReplaceAll(strings.ToUpper(envPrefix+field.Name()), ".", "_")
		comment += " * " + field.Name() + " is parsed from flag '" + envName + "'\n"

		body.AddLine("if value,ok := ", src.NewTypeDecl("os.LookupEnv"), "(\"", envName, "\");ok{")
		var myTypeDecl *src.TypeDecl
		switch oType.TypeName() {
		case "string":
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = value")
		case "bool":
			myTypeDecl = src.NewTypeDecl("bool")
		case "int64":
			myTypeDecl = src.NewTypeDecl("int64")
		case "float64":
			myTypeDecl = src.NewTypeDecl("float64")
		case "time.Duration":
			myTypeDecl = src.NewTypeDecl("time.Duration")
		default:
			return buildErr("TypeName", string(oType.TypeName()), "is not supported to be parsed as a flag", oType.Pos())
		}

		// this is currently only if no parser is needed, just like plain string
		if myTypeDecl != nil {
			code, err := genParseStr("value", myTypeDecl)
			if err != nil {
				return err
			}
			body.AddLine("v,err := ", code, "")
			body.AddLine("if err != nil {")
			body.AddLine("return ", src.NewTypeDecl("fmt.Errorf"), "(\"cannot parse environment variable '", envName, "': %w\"", ",err)")
			body.AddLine("}")
			body.NewLine()
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = v")
		}

		body.AddLine("}")
		body.NewLine()
	}
	body.Add("return nil")

	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}
