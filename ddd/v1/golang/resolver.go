package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/src"
	"reflect"
	"strconv"
	"strings"
)

type resolverScope uint8

const (
	rCore resolverScope = 1 << iota
	rUniverse
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

	return strings.Join(msg, "|")
}

type resolver struct {
	path     string
	ctx      *ddd.BoundedContextSpec
	layers   []typeLayer
	universe []typeDef
	core     typeLayer
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

			for _, structOrInterface := range l.API() {
				tDef := typeDef{
					typeName: ddd.TypeName(structOrInterface.Name()),
					typeDecl: src.NewTypeDecl(src.Qualifier(layerPath + "." + structOrInterface.Name())),
				}
				tlayer.typeDefs = append(tlayer.typeDefs, tDef)
			}

			r.core = tlayer

		default:
			panic("not yet implemented: " + reflect.TypeOf(l).String())
		}
	}

	return r
}

func (r *resolver) resolveTypeName(scopes resolverScope, t ddd.TypeName) (*src.TypeDecl, error) {
	baseType := removeGenericType(t)

	if scopes.Has(rCore) {
		for _, def := range r.core.typeDefs {
			if def.typeName == baseType {
				return makeGeneric(t, def.typeDecl), nil
			}
		}
	}

	if scopes.Has(rUniverse) {
		for _, def := range r.universe {
			if def.typeName == baseType {
				return makeGeneric(t, def.typeDecl), nil
			}
		}
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
