package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/src"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type resolverScope uint8

const (
	rCore resolverScope = 1 << iota
	rUniverse
	rUsecase
	rRest
)

func (r resolverScope) Has(flag resolverScope) bool {
	return r&flag != 0
}

func (r resolverScope) String() string {
	var msg []string
	if r.Has(rCore) {
		msg = append(msg, "core")
	}

	if r.Has(rUniverse) {
		msg = append(msg, "universe")
	}

	if r.Has(rRest) {
		msg = append(msg, "rest")
	}

	return strings.Join(msg, "|")
}

type resolver struct {
	path       string
	ctx        *ddd.BoundedContextSpec
	layers     []typeLayer
	universe   []typeDef
	core       typeLayer
	usecase    typeLayer
	restLayers []typeLayer
}

type typeLayer struct {
	layer    ddd.Layer
	path     string
	typeDefs []typeDef
}

type typeDef struct {
	typeName ddd.TypeName
	typeDecl *src.TypeDecl
}

func (t typeDef) getDecl() *src.TypeDecl {
	return t.typeDecl.Clone()
}

func newResolver(modPath string, ctx *ddd.BoundedContextSpec) *resolver {
	r := &resolver{
		path: modPath,
		ctx:  ctx,
		universe: []typeDef{
			{
				typeName: ddd.UUID,
				typeDecl: src.NewTypeDecl("github.com/golangee/uuid.UUID"),
			},
			{
				typeName: ddd.String,
				typeDecl: src.NewTypeDecl("string"),
			},
			{
				typeName: ddd.Int64,
				typeDecl: src.NewTypeDecl("int64"),
			},
			{
				typeName: ddd.Float32,
				typeDecl: src.NewTypeDecl("float32"),
			},
			{
				typeName: ddd.Bool,
				typeDecl: src.NewTypeDecl("bool"),
			},
			{
				typeName: ddd.Ctx,
				typeDecl: src.NewTypeDecl("context.Context"),
			},
			{
				typeName: ddd.Error,
				typeDecl: src.NewTypeDecl("error"),
			},
		},
	}

	for _, layer := range ctx.Layers() {
		switch l := layer.(type) {
		case *ddd.CoreLayerSpec:
			layerPath := modPath + "/internal/" + safename(ctx.Name()) + "/" + pkgNameCore
			tlayer := typeLayer{
				layer: layer,
				path:  layerPath,
			}

			appendStructOrInterfaces(&tlayer, l.API())
			appendFuncOrStructs(&tlayer, l.Factories())
			r.core = tlayer

		case *ddd.UseCaseLayerSpec:
			layerPath := modPath + "/internal/" + safename(ctx.Name()) + "/" + pkgNameUseCase
			tlayer := typeLayer{
				layer: layer,
				path:  layerPath,
			}

			for _, useCase := range l.UseCases() {
				for _, story := range useCase.Stories() {
					appendFuncOrStructs(&tlayer, []ddd.FuncOrStruct{story.Func()})
					for _, strct := range story.Structs() {
						appendFuncOrStructs(&tlayer, []ddd.FuncOrStruct{strct})
					}
				}
			}

			r.usecase = tlayer

		case *ddd.RestLayerSpec:
			layerPath := modPath + "/internal/" + safename(ctx.Name()) + "/" + l.Name()
			tlayer := typeLayer{
				layer: layer,
				path:  layerPath,
			}

			r.restLayers = append(r.restLayers, tlayer)
		default:
			panic("not yet implemented: " + reflect.TypeOf(l).String())
		}
	}

	return r
}

func appendStructOrInterfaces(dst *typeLayer, structOrInterfaces []ddd.StructOrInterface) {
	for _, structOrInterface := range structOrInterfaces {
		tDef := typeDef{
			typeName: ddd.TypeName(structOrInterface.Name()),
			typeDecl: src.NewTypeDecl(src.Qualifier(dst.path + "." + structOrInterface.Name())),
		}
		dst.typeDefs = append(dst.typeDefs, tDef)
	}
}

func appendFuncOrStructs(dst *typeLayer, funcOrStructs []ddd.FuncOrStruct) {
	for _, funcOrStruct := range funcOrStructs {
		if strct, ok := funcOrStruct.(*ddd.StructSpec); ok {
			tDef := typeDef{
				typeName: ddd.TypeName(strct.Name()),
				typeDecl: src.NewTypeDecl(src.Qualifier(dst.path + "." + strct.Name())),
			}
			dst.typeDefs = append(dst.typeDefs, tDef)
		}
	}
}

// looksLikeFullQualifier returns true for strings like abc/xyz.Def
func looksLikeFullQualifier(t ddd.TypeName) bool {
	aDot := strings.LastIndex(string(t), ".")
	quiteOk := aDot > 0
	for _, r := range src.Qualifier(t).Name() {
		return unicode.IsUpper(r) && quiteOk
	}

	return false
}

func (r *resolver) resolveTypeName(scopes resolverScope, t ddd.TypeName) (*src.TypeDecl, error) {
	baseType := removeGenericType(t)

	if looksLikeFullQualifier(baseType) {
		return makeGeneric(t, src.NewTypeDecl(src.Qualifier(baseType))), nil
	}

	if scopes.Has(rCore) {
		for _, def := range r.core.typeDefs {
			if def.typeName == baseType {
				return makeGeneric(t, def.getDecl()), nil
			}
		}
	}

	if scopes.Has(rUniverse) {
		for _, def := range r.universe {
			if def.typeName == baseType {
				return makeGeneric(t, def.getDecl()), nil
			}
		}
	}

	if scopes.Has(rUsecase) {
		for _, def := range r.usecase.typeDefs {
			if def.typeName == baseType {
				return makeGeneric(t, def.getDecl()), nil
			}
		}
	}

	if scopes.Has(rRest) {
		//TODO
	}

	return nil, fmt.Errorf("type '%s' cannot be resolved from layers '%s'", t, scopes.String())
}

// Removes * or [] or [x]
func removeGenericType(t ddd.TypeName) ddd.TypeName {
	if strings.HasPrefix(string(t), "*") {
		return t[1:]
	}

	if strings.HasPrefix(string(t), "[]") {
		return t[2:]
	}

	if strings.HasPrefix(string(t), "[") {
		for i, r := range t {
			if r == ']' {
				return t[:i]
			}
		}
		panic("illegal state")
	}

	return t
}

func makeGeneric(t ddd.TypeName, decl *src.TypeDecl) *src.TypeDecl {
	if strings.HasPrefix(string(t), "*") {
		return src.NewPointerDecl(decl)
	}

	if strings.HasPrefix(string(t), "[]") {
		return src.NewSliceDecl(decl)
	}

	if strings.HasPrefix(string(t), "[") {
		for i, r := range t {
			if r == ']' {
				size, err := strconv.ParseInt(string(t[0:i]), 10, 32)
				if err != nil {
					panic("failed to parse " + string(t))
				}
				src.NewArrayDecl(size, decl)
			}
		}
		panic("illegal state")
	}

	return decl
}
