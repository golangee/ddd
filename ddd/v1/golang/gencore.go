package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/src"
	"path/filepath"
	"reflect"
)

const (
	pkgNameCore = "core"
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
					default:
						panic("not yet implemented: " + reflect.TypeOf(t).String())
					}
				}
			default:
				panic("not yet implemented: " + reflect.TypeOf(l).String())
			}

		}

	}

	return nil
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

		s.AddMethods(f)

		f.SetDoc(myComment)
	}

	return s, nil
}
