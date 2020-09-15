package golang

import (
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"strconv"
	"strings"
)

const (
	appPkgName  = "application"
	appTypeName = "App"
)

// generateCmdSrv emits the main package
func generateCmdSrv(ctx *genctx) error {
	ctx.newFile("cmd/"+text.Safename(ctx.spec.Name()), "main", "main").
		SetPackageDoc("Package main contains the executable to launch the actual " + ctx.spec.Name() + " server process.").
		AddFuncs(
			src.NewFunc("main").AddBody(src.NewBlock().
				AddLine("a,err :=", src.NewTypeDecl(src.Qualifier(ctx.mod.Main().Path+"/internal/"+appPkgName+".New"+appTypeName)), "()").
				AddLine("if err != nil{").
				AddLine(src.NewTypeDecl("log.Fatal"), "(err)").
				AddLine("}").
				NewLine().
				AddLine("if err := a.Start();err!=nil{").
				AddLine(src.NewTypeDecl("log.Fatal"), "(err)").
				AddLine("}"),
			),


		)

	return nil
}

// generateApp emits the actual application code
func generateApp(ctx *genctx) error {
	ctx.newFile("internal/"+appPkgName, "doc", "").SetPackageDoc("Package " + appPkgName + " contains the dependency injection logic and assembles all layers together.")
	file := ctx.newFile("internal/"+appPkgName, "app", "")

	generateUberOptions(ctx)

	app := src.NewStruct(appTypeName).SetDoc("...is the actual application, which glues all layers together and is launched from the command line.")
	for _, factory := range ctx.factorySpecs {
		serviceType := factory.factoryFunc.Results()[0].Decl().Clone()
		field := src.NewField(makePackagePrivate(serviceType.Qualifier().Name()), serviceType)
		app.AddFields(field)
	}

	constructor := generateAppConstructor(ctx)
	file.AddFuncs(constructor)

	file.AddTypes(app)

	startFun := generateAppStart(ctx)
	app.AddMethods(startFun)

	return nil
}

func generateAppConstructor(ctx *genctx) *src.FuncBuilder {
	fun := src.NewFunc("New" + appTypeName).SetDoc("...creates a new instance of the application and performs all parameter parsing and wiring.")
	fun.AddResults(
		src.NewParameter("", src.NewPointerDecl(src.NewTypeDecl(appTypeName))),
		src.NewParameter("", src.NewTypeDecl("error")),
	)

	body := src.NewBlock()
	if len(ctx.factorySpecs) > 0 {
		body.AddLine("var err error")
	}
	body.AddLine("options := Options{}")
	body.AddLine("options.Reset()")
	body.AddLine("if err := options.ParseEnv(); err!=nil {")
	body.AddLine("return nil,err")
	body.AddLine("}")
	body.AddLine("options.ConfigureFlags()")
	body.AddLine("help := ", src.NewTypeDecl("flag.Bool"), "(\"help\",false,\"shows this help\")")
	body.AddLine(src.NewTypeDecl("flag.Parse()"))
	body.AddLine("if *help {")
	body.AddLine(src.NewTypeDecl("fmt.Println"), "(\"", ctx.spec.Name(), "\")")
	body.AddLine(src.NewTypeDecl("flag.PrintDefaults"), "()")
	body.AddLine(src.NewTypeDecl("os.Exit"), "(0)")
	body.AddLine("}")
	body.AddLine("a := &", appTypeName, "{}")

	knownDependencies := map[string]*src.TypeDecl{}

	for _, factory := range ctx.factorySpecs {
		serviceType := factory.factoryFunc.Results()[0].Decl().Clone()
		fieldName := makePackagePrivate(serviceType.Qualifier().Name())

		no := 2
		for {
			_, has := knownDependencies[fieldName]
			if has {
				fieldName = makePackagePrivate(serviceType.Qualifier().Name()) + strconv.Itoa(no)
				continue
			}
			break
		}

		knownDependencies[fieldName] = serviceType

		warnDepNotFoundComment := src.NewBlock()
		body.Add(warnDepNotFoundComment)
		body.Add("if a.", fieldName, ", err =", src.NewTypeDecl(serviceType.Qualifier()+"Factory"), "(")
		for i, p := range factory.factoryFunc.Params() {
			if i == 0 {
				fieldName := getUberOptionsFieldName(factory.factoryFunc.Params()[0].Decl().Qualifier())
				body.Add("options.", fieldName)
				continue
			}
			dependencyFound := false
			for field, decl := range knownDependencies {
				// we just compare qualifiers, because our dependencies are always interfaces
				if p.Decl().Qualifier() == decl.Qualifier() {
					dependencyFound = true
					body.Add(",a." + field)
					break
				}
			}

			if !dependencyFound {
				body.Add(",nil")
				warnDepNotFoundComment.AddLine("// TODO dependency to " + p.Name() + " of type " + p.Decl().String() + " cannot be resolved.")
			}
		}
		body.AddLine(");err!=nil{")
		body.AddLine("return nil,err")
		body.AddLine("}")
		body.NewLine()
	}
	body.AddLine("return a,nil")

	fun.AddBody(body)

	return fun
}

func generateUberOptions(ctx *genctx) {
	file := ctx.newFile("internal/"+appPkgName, "options", "")
	opt := src.NewStruct("Options").SetDoc("...represents the uber options and bundles all options from all\nbounded contexts and layers in one monolithic instance.")

	resetFun := src.NewFunc("Reset").SetDoc("...restores all default values.").SetPointerReceiver(true)
	resetFunBody := src.NewBlock()
	opt.AddMethods(resetFun.AddBody(resetFunBody))

	parseEnvFun := src.NewFunc("ParseEnv").SetDoc("...parses values from the environment.").SetPointerReceiver(true).AddResults(src.NewParameter("", src.NewTypeDecl("error")))
	parseEnvFunBody := src.NewBlock()
	opt.AddMethods(parseEnvFun.AddBody(parseEnvFunBody))

	configureFlagsFun := src.NewFunc("ConfigureFlags").SetDoc("...configures the flag package to receive parsed flags.").SetPointerReceiver(true)
	configureFlagsFunBody := src.NewBlock()
	opt.AddMethods(configureFlagsFun.AddBody(configureFlagsFunBody))

	for _, factory := range ctx.factorySpecs {
		serviceType := factory.factoryFunc.Params()[0].Decl().Clone()
		fieldName := getUberOptionsFieldName(serviceType.Qualifier())
		field := src.NewField(fieldName, serviceType)
		opt.AddFields(field)

		resetFunBody.AddLine(resetFun.ReceiverName(), ".", field.Name(), ".Reset()")

		parseEnvFunBody.AddLine("if err:=", resetFun.ReceiverName(), ".", field.Name(), ".ParseEnv(); err!=nil{")
		parseEnvFunBody.AddLine("return err")
		parseEnvFunBody.AddLine("}")
		parseEnvFunBody.NewLine()

		configureFlagsFunBody.AddLine(resetFun.ReceiverName(), ".", field.Name(), ".ConfigureFlags()")
	}
	parseEnvFunBody.AddLine("return nil")

	file.AddTypes(opt)
}

func getUberOptionsFieldName(optsType src.Qualifier) string {
	path := strings.Split(optsType.Path(), "/")
	return text.MakePublic(path[len(path)-2]) + text.MakePublic(path[len(path)-1]) + optsType.Name()
}

func generateAppStart(ctx *genctx) *src.FuncBuilder {
	fun := src.NewFunc("Start").
		SetDoc("...launches any blocking background processes, like e.g. an http server.").
		AddResults(src.NewParameter("", src.NewTypeDecl("error")))
	body := src.NewBlock()
	body.AddLine("return nil")
	fun.AddBody(body)

	return fun
}
