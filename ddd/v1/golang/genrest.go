package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"strconv"
)

// StringParser is hacky way to allow custom serialization codes
type StringParser func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error)

// NotSupportedErr is only used by StringParser implementations.
var NotSupportedErr = fmt.Errorf("string parsing not supported")
var restParamParsers []StringParser

func init() {
	AddStringParser(func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error) {
		if targetType.Qualifier() == "int64" {
			return src.NewBlock().Add(src.NewTypeDecl("strconv.ParseInt"), "(", fromStrCode, ", 10, 64)"), nil
		}
		return nil, NotSupportedErr
	})

	AddStringParser(func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error) {
		if targetType.Qualifier() == "float64" {
			return src.NewBlock().Add(src.NewTypeDecl("strconv.ParseFloat"), "(", fromStrCode, ", 64)"), nil
		}
		return nil, NotSupportedErr
	})

	AddStringParser(func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error) {
		if targetType.Qualifier() == "github.com/golangee/uuid.UUID" {
			return src.NewBlock().Add(src.NewTypeDecl("github.com/golangee/uuid.Parse"), "(", fromStrCode, ")"), nil
		}
		return nil, NotSupportedErr
	})

	AddStringParser(func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error) {
		if targetType.Qualifier() == "bool" {
			return src.NewBlock().Add(src.NewTypeDecl("strconv.ParseBool"), "(", fromStrCode, ")"), nil
		}
		return nil, NotSupportedErr
	})

	AddStringParser(func(fromStrCode string, targetType *src.TypeDecl) (*src.Block, error) {
		if targetType.Qualifier() == "time.Duration" {
			return src.NewBlock().Add(src.NewTypeDecl("time.ParseDuration"), "(", fromStrCode, ")"), nil
		}
		return nil, NotSupportedErr
	})

}

// AddStringParser registers another StringParser for the REST parameter deserialization.
// The following types are supported out of the box:
//  * int64
//  * github.com/golangee/uuid.UUID
func AddStringParser(parser StringParser) {
	restParamParsers = append(restParamParsers, parser)
}

// createRestLayer emits the interfaces and types in the <bc>/rest/vX package according to the REST specification.
func createRestLayer(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, rest *ddd.RestLayerSpec) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, pkgNameRest)

	ctx.newFile(layerPath, "doc", pkgNameRest).
		SetPackageDoc("Package " + pkgNameRest + " contains subpackages for each major REST API version.")

	for _, version := range rest.Versions() {
		if err := createRestVersion(ctx, layerPath, rslv, bc, rest, version); err != nil {
			return err
		}
	}

	return nil
}

func createRestVersion(ctx *genctx, layerPath string, rslv *resolver, bc *ddd.BoundedContextSpec, rest *ddd.RestLayerSpec, version *ddd.VersionSpec) error {
	ctx.newFile(filepath.Join(layerPath, version.MajorVersion()), "doc", pkgNameRest).
		SetPackageDoc(rest.Description() + "\nThe current api version is " + version.Version() + ".")

	/*var generatedResourceIfaces []*src.TypeBuilder
	var resourcePaths []string
	for _, resource := range rest.Resources() {
		goResName := text2GoIdentifier(strings.ReplaceAll(resource.Path(), ":", "-By-"))
		resFile := ctx.newFile(filepath.Join(layerPath, rest.MajorVersion()), strings.ToLower(text.Safename(goResName)), pkgNameRest)

		// create a documented interface
		resIface, err := createRestResourceInterface(resFile, rslv, goResName, resource)
		if err != nil {
			return err
		}
		generatedResourceIfaces = append(generatedResourceIfaces, resIface)

		endpointPath := text.JoinSlashes(rest.Prefix(), resource.Path())
		resourcePaths = append(resourcePaths, endpointPath)
		resIface.SetDoc(resIface.Name() + " represents the REST resource " + endpointPath + ".\n" + resource.Description())
		resFile.AddTypes(resIface)
		resFile.AddTypes(src.ImplementMock(resIface))

		// create the handler factories
		for i, fun := range resIface.Methods() {
			var myParamsInGenFieldOrder []*ddd.ParamSpec
			for _, p := range resource.Params() {
				for _, spec := range p.Params() {
					myParamsInGenFieldOrder = append(myParamsInGenFieldOrder, spec)
				}
			}
			for _, p := range resource.Verbs()[i].Params() {
				for _, spec := range p.Params() {
					myParamsInGenFieldOrder = append(myParamsInGenFieldOrder, spec)
				}
			}
			factoryFunc, err := createHandlerFuncFactory(resFile, myParamsInGenFieldOrder, endpointPath, fun)
			if err != nil {
				return fmt.Errorf("unable to create handler func for %s: %w", endpointPath, err)
			}
			resFile.AddFuncs(factoryFunc)
		}

		resFile.AddFuncs(createConfigureFunc(resIface))
	}

	api := ctx.newFile(filepath.Join(layerPath, rest.MajorVersion()), "api", pkgNameRest)
	api.AddFuncs(createWrapHandlerFunc())

	// add a common interface
	commonIface := src.NewInterface("API")
	tmpComment := "... is the common interface which represents the following resources:\n"
	for i, iface := range generatedResourceIfaces {
		commonIface.AddEmbedded(src.NewTypeDecl(src.Qualifier(iface.Name())))
		tmpComment += " * " + iface.Name() + ": " + resourcePaths[i] + "\n"
	}
	commonIface.SetDoc(tmpComment)
	api.AddTypes(commonIface)

	// add the all-in-one configure method for the commonIface
	api.AddFuncs(createMainConfigureFunc(generatedResourceIfaces))*/

	return nil
}
/*
// createRestResourceInterface creates an interface which represents the http verbs as methods.
func createRestResourceInterface(resFile *src.FileBuilder, rslv *resolver, ifaceName string, resource *ddd.HttpResourceSpec) (*src.TypeBuilder, error) {
	resIface := src.NewInterface(ifaceName)

	for _, verb := range resource.Verbs() {
		paramType, err := createRestResourceVerbRequestContextStruct(ifaceName, rslv, resource, verb)
		if err != nil {
			return nil, err
		}
		resFile.AddTypes(paramType)

		verbFunc := src.NewFunc(text2GoIdentifier(strings.ToLower(verb.Method()) + ifaceName))
		verbFunc.AddParams(src.NewParameter("ctx", src.NewTypeDecl(src.Qualifier(paramType.Name())))).
			SetDoc(verbFunc.Name() + " represents the http " + verb.Method() + " request on the " + resource.Path() + " resource.\n" + verb.Description()).
			AddResults(src.NewParameter("", src.NewTypeDecl("error")))
		resIface.AddMethods(verbFunc)

	}

	return resIface, nil
}

// createRestResourceVerbRequestContextStruct assembles a type safe struct per verb request.
func createRestResourceVerbRequestContextStruct(name string, rslv *resolver, parent *ddd.HttpResourceSpec, verb *ddd.VerbSpec) (*src.TypeBuilder, error) {
	req := src.NewStruct(name + text2GoIdentifier(strings.ToLower(verb.Method())) + "Context")
	req.SetDoc(req.Name()+" provides the specific http request and response context including already parsed parameters.").
		AddFields(
			src.NewField("Request", src.NewPointerDecl(src.NewTypeDecl("net/http.Request"))).
				SetDoc("Request contains the raw http request."),
			src.NewField("Writer", src.NewTypeDecl("net/http.ResponseWriter")).
				SetDoc("Writer contains a reference to the raw http response writer."),
		)

	var allParams []*ddd.ParamSpec

	for _, params := range parent.Params() {
		for _, param := range params.Params() {
			allParams = append(allParams, param)
		}
	}

	for _, params := range verb.Params() {
		for _, param := range params.Params() {
			allParams = append(allParams, param)
		}
	}

	for _, param := range allParams {
		decl, err := rslv.resolveTypeName(rUniverse|rUsecase|rCore|rRest, param.TypeName())
		if err != nil {
			return nil, err
		}

		field := src.NewField(text2GoIdentifier(param.Name()), decl)
		var originName string
		switch param.NameValidationKind() {
		case ddd.NVHttpHeaderParameter:
			originName = "header"
		case ddd.NVHttpPathParameter:
			originName = "path"
		case ddd.NVHttpQueryParameter:
			originName = "query"
		default:
			panic("unexpected parameter kind: " + strconv.Itoa(int(param.NameValidationKind())))
		}

		field.SetDoc(field.Name() + " contains the parsed " + originName + " parameter for '" + param.Name() + "'.")
		req.AddFields(field)
	}
	return req, nil
}

func createMainConfigureFunc(ifaces []*src.TypeBuilder) *src.FuncBuilder {

	cfgFunc := src.NewFunc("Configure").
		SetDoc("...just applies the entire API to the given router.").
		AddParams(
			src.NewParameter("api", src.NewTypeDecl("API")),
			src.NewParameter("router", src.NewTypeDecl("github.com/julienschmidt/httprouter.Router")),
		)

	body := src.NewBlock()
	for _, iface := range ifaces {
		body.AddLine("Configure"+iface.Name(), "(api, router)")
	}
	cfgFunc.AddBody(body)

	return cfgFunc
}
	*/
// createWrapHandlerFunc emits a wrap function to convert between the julienschmidt-router handle the stdlib handlerFunc.
func createWrapHandlerFunc() *src.FuncBuilder {

	/*
		//wrap converts from the http.HandlerFunc to the httprouter.Handle interface.
		func wrap(route string, handler http.HandlerFunc) (string, httprouter.Handle) {
			return route, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
				handler(writer, request)
			}
		}
	*/

	const httpRouter = "github.com/julienschmidt/httprouter.Handle"
	return src.NewFunc("wrap").
		SetDoc("wrap converts from the http.HandlerFunc to the httprouter.Handle interface.").
		AddParams(
			src.NewParameter("route", src.NewTypeDecl("string")),
			src.NewParameter("handler", src.NewTypeDecl("net/http.HandlerFunc")),
		).
		AddResults(
			src.NewParameter("", src.NewTypeDecl("string")),
			src.NewParameter("", src.NewTypeDecl(httpRouter)),
		).AddBody(
		src.NewBlock().
			AddLine("return route, func(writer ", src.NewTypeDecl("net/http.ResponseWriter"), ", request *", src.NewTypeDecl("net/http.Request"), ",_ ", src.NewTypeDecl("github.com/julienschmidt/httprouter.Params"), ") {").
			AddLine("handler(writer, request)").
			AddLine("}"),
	)

}

// createHandlerFuncFactory assembles a public package function to create a
// standard handler and the according path. It currently only works
// with the julienschmidt httprouter
func createHandlerFuncFactory(resFile *src.FileBuilder, myParamsInGenFieldOrder []*ddd.ParamSpec, endpointPath string, fun *src.FuncBuilder) (*src.FuncBuilder, error) {
	const httprouter = "github.com/julienschmidt/httprouter."
	statusBadRequest := src.NewTypeDecl("net/http.StatusBadRequest")
	ctxType := resFile.TypeByName(fun.Params()[0].Decl().Qualifier().Name())
	hasAnyPathParams := false
	for _, p := range myParamsInGenFieldOrder {
		switch p.NameValidationKind() {
		case ddd.NVHttpPathParameter:
			hasAnyPathParams = true
			break
		}
	}

	hFun := src.NewFunc(fun.Name()).SetDoc("...returns the route to register on and the handler to execute.\nCurrently, only the httprouter.Router is supported.")
	hFun.AddParams(
		src.NewParameter("api", src.NewFuncDecl(fun.Params(), fun.Results())),
	)
	hFun.AddResults(
		src.NewParameter("route", src.NewTypeDecl("string")),
		src.NewParameter("handler", src.NewTypeDecl("net/http.HandlerFunc")),
	)

	body := src.NewBlock().
		Add("return \"", endpointPath, "\"").
		AddLine(", func(w ", src.NewTypeDecl("net/http.ResponseWriter"), ", r", src.NewPointerDecl(src.NewTypeDecl("net/http.Request")), "){")

	if hasAnyPathParams {
		body.AddLine("p := ", src.NewTypeDecl(httprouter+"ParamsFromContext"), "(r.Context())")
	}
	body.AddLine("var err error")
	body.AddLine("ctx := ", ctxType.Name(), "{")
	body.AddLine("Request: r,")
	body.AddLine("Writer: w,")
	body.AddLine("}")

	for i, field := range ctxType.Fields() {
		switch field.Name() {
		case "Request":
			fallthrough
		case "Writer":
			// intentionally ignored
		default:
			p := myParamsInGenFieldOrder[i-2]
			switch p.NameValidationKind() {
			case ddd.NVHttpHeaderParameter:
				code, err := parseVarFromStr("ctx."+field.Name(), "r.Header.Get(\""+p.Name()+"\")", field.Type(), p.Optional())
				if err != nil {
					return nil, err
				}
				body.Add(code)
			case ddd.NVHttpPathParameter:
				code, err := parseVarFromStr("ctx."+field.Name(), "p.ByName(\""+p.Name()+"\")", field.Type(), p.Optional())
				if err != nil {
					return nil, err
				}
				body.Add(code)
			case ddd.NVHttpQueryParameter:
				code, err := parseVarFromStr("ctx."+field.Name(), "r.URL.Query().Get(\""+p.Name()+"\")", field.Type(), p.Optional())
				if err != nil {
					return nil, err
				}
				body.Add(code)
			default:
				panic("kind not implemented: " + strconv.Itoa(int(p.NameValidationKind())))
			}

		}
	}

	body.AddLine("if err=api(ctx);err!=nil{")
	body.AddLine(logPrintln("r.URL.String(),", "err")...)
	body.AddLine("w.WriteHeader(", statusBadRequest.Clone(), ")")
	body.AddLine("return")
	body.AddLine("}")

	body.AddLine("}")

	hFun.AddBody(body)
	return hFun, nil
}

func genParseStr(fromName string, targetType *src.TypeDecl) (*src.Block, error) {
	for _, parser := range restParamParsers {
		if code, err := parser(fromName, targetType); err == nil {
			return code, nil
		}
	}
	return nil, fmt.Errorf("REST: don't know how to parse '" + targetType.String() + "' from string")
}

// parseVarFromStr emits a helper block to parse various types from a string.
func parseVarFromStr(toName, fromName string, targetType *src.TypeDecl, optional bool) (*src.Block, error) {
	if targetType.Qualifier() == "string" {
		return src.NewBlock(toName, " = ", fromName), nil
	}

	statusBadRequest := src.NewTypeDecl("net/http.StatusBadRequest")

	block := src.NewBlock()
	if optional {
		block.Add(toName, ",_ = ")
	} else {
		block.Add("if ", toName, ", err = ")
	}
	code, err := genParseStr(fromName, targetType)
	if err != nil {
		return nil, err
	}

	block.Add(code)

	if !optional {
		block.AddLine("; err != nil {")
		block.AddLine(logPrintln("r.URL.String(),", "err")...)
		block.AddLine("w.WriteHeader(", statusBadRequest.Clone(), ")")
		block.AddLine("return")
		block.AddLine("}")
	}
	block.NewLine()

	return block, nil
}

func createConfigureFunc(resIface *src.TypeBuilder) *src.FuncBuilder {
	cfgFunc := src.NewFunc("Configure"+resIface.Name()).
		SetDoc("...just applies the package wide endpoints into the given router without any other middleware.").
		AddParams(
			src.NewParameter("api", src.NewTypeDecl(src.Qualifier(resIface.Name()))),
			src.NewParameter("router", src.NewTypeDecl("github.com/julienschmidt/httprouter.Router")),
		)
	body := src.NewBlock()
	for _, fun := range resIface.Methods() {
		verb := text.ParseVerb(fun.Name())
		if verb == "" {
			panic("verb routing not yet implemented: " + fun.Name())
		}

		body.AddLine("router.", verb, "(wrap(", fun.Name(), "(api.", fun.Name(), ")))")
	}

	cfgFunc.AddBody(body)
	return cfgFunc
}

// logPrintln emits a row to write into the golang logger
func logPrintln(codes ...interface{}) []interface{} {
	tmp := []interface{}{
		src.NewTypeDecl("log.Println"),
		"(",
	}
	tmp = append(tmp, codes...)
	tmp = append(tmp, ")")
	return tmp
}
