package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"strconv"
	"strings"
)

// createRestLayer emits the interfaces and types in the <bc>/rest/vX package according to the REST specification.
func createRestLayer(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, rest *ddd.RestLayerSpec) error {
	bcPath := filepath.Join("internal", safename(bc.Name()))
	layerPath := filepath.Join(bcPath, pkgNameRest)

	ctx.newFile(layerPath, "doc", pkgNameRest).
		SetPackageDoc("Package " + pkgNameRest + " contains subpackages for each major REST API version.")
	ctx.newFile(filepath.Join(layerPath, rest.MajorVersion()), "doc", pkgNameRest).
		SetPackageDoc(rest.Description() + "\nThe current api version is " + rest.Version() + ".")

	for _, resource := range rest.Resources() {
		goResName := text2GoIdentifier(strings.ReplaceAll(resource.Path(), ":", "-By-"))
		resFile := ctx.newFile(filepath.Join(layerPath, rest.MajorVersion()), strings.ToLower(safename(goResName)), pkgNameRest)

		resIface := src.NewInterface(goResName)
		resIface.SetDoc(resIface.Name() + " represents the REST resource " + text.JoinSlashes(rest.Prefix(), resource.Path()) + ".\n" + resource.Description())
		for _, verb := range resource.Verbs() {
			paramType, err := createRestResourceVerbRequest(goResName, rslv, resource, verb)
			if err != nil {
				return err
			}
			resFile.AddTypes(paramType)

			verbFunc := src.NewFunc(text2GoIdentifier(strings.ToLower(verb.Method())))
			verbFunc.AddParams(src.NewParameter("ctx", src.NewTypeDecl(src.Qualifier(paramType.Name())))).
				SetDoc(verbFunc.Name() + " represents the http " + verb.Method() + " request on the " + resource.Path() + " resource.\n" + verb.Description()).
				AddResults(src.NewParameter("", src.NewTypeDecl("error")))
			resIface.AddMethods(verbFunc)

		}
		resFile.AddTypes(resIface)
	}

	return nil
}

// createRestResourceVerbRequest assembles a type safe struct per verb request.
func createRestResourceVerbRequest(name string, rslv *resolver, parent *ddd.HttpResourceSpec, verb *ddd.VerbSpec) (*src.TypeBuilder, error) {
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
