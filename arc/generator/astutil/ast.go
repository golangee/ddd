package astutil

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"strconv"
	"strings"
)

func CallMember(recName, fieldName, methodName string, args ...ast.Expr) *ast.CallExpr {
	return ast.NewCallExpr(ast.NewSelExpr(ast.NewSelExpr(ast.NewIdent(recName), ast.NewIdent(fieldName)), ast.NewIdent(methodName)), args...)
}

func MethodByName(f ast.Node, name string) *ast.Func {
	type Methoder interface {
		Methods() []*ast.Func
	}

	if m, ok := f.(Methoder); ok {
		for _, a := range m.Methods() {
			if a.FunName == name {
				return a
			}
		}
	}

	return nil
}

func FieldByName(f ast.Node, name string) *ast.Field {
	type Fielder interface {
		Fields() []*ast.Field
	}

	if m, ok := f.(Fielder); ok {
		for _, a := range m.Fields() {
			if a.FieldName == name {
				return a
			}
		}
	}

	return nil
}

func FindMod(name token.String, prj *ast.Prj) (*ast.Mod, error) {
	for _, mod := range prj.Mods {
		if mod.Name == name.String() {
			return mod, nil
		}
	}

	return nil, token.NewPosError(name, "module not found")
}

// MkMod returns or create the module.
func MkMod(prj *ast.Prj, modName string) *ast.Mod {
	for _, mod := range prj.Mods {
		if mod.Name == modName {
			return mod
		}
	}

	mod := ast.NewMod(modName)
	prj.AddModules(mod)

	return mod
}

// FindPkg returns an existing package or nil.
func FindPkg(mod *ast.Mod, pkgPath string) *ast.Pkg {
	for _, pkg := range mod.Pkgs {
		if pkg.Path == pkgPath {
			return pkg
		}
	}

	return nil
}

// MkPkg returns an existing or create a new package inside mod.
func MkPkg(mod *ast.Mod, pkgPath string) *ast.Pkg {
	for _, pkg := range mod.Pkgs {
		if pkg.Path == pkgPath {
			return pkg
		}
	}

	pkg := ast.NewPkg(pkgPath)
	mod.AddPackages(pkg)

	return pkg
}

// MkFile returns an existing file or creates a new file inside the package.
func MkFile(pkg *ast.Pkg, name string) *ast.File {
	for _, file := range pkg.PkgFiles {
		if file.Name == name {
			return file
		}
	}

	f := ast.NewFile(name)
	pkg.AddFiles(f)

	return f
}

func ResolveLocal(ctx ast.Node, name string) ast.Node {
	lastDot := strings.LastIndex(name, ".")
	if lastDot < 0 {
		for _, file := range Pkg(ctx).PkgFiles {
			for _, node := range file.Nodes {
				if tname, ok := node.(ast.NamedType); ok {
					if tname.Identifier() == name {
						return tname
					}
				}
			}
		}
	}

	return nil
}

// Resolve takes the name and walks until it finds whatever declares it. Returns nil
// if not found.
func Resolve(ctx ast.Node, name string) ast.Node {
	lastDot := strings.LastIndex(name, ".")

	// try local package search
	if lastDot < 0 {
		return ResolveLocal(ctx, name)
	}

	// resolve package and return type from that
	pkgName := name[:lastDot]
	typeName := name[lastDot+1:]
	for _, pkg := range Mod(ctx).Pkgs {
		if pkg.Path == pkgName {
			for _, file := range pkg.PkgFiles {
				for _, node := range file.Children() {
					if tname, ok := node.(ast.NamedType); ok {
						if tname.Identifier() == typeName {
							return tname
						}
					}
				}
			}

		}
	}

	return nil
}

func Mod(n ast.Node) *ast.Mod {
	mod := &ast.Mod{}
	if ok := ast.ParentAs(n, &mod); ok {
		return mod
	}

	return nil
}

func Pkg(n ast.Node) *ast.Pkg {
	p := &ast.Pkg{}
	if ok := ast.ParentAs(n, &p); ok {
		return p
	}

	return nil
}

func File(n ast.Node) *ast.File {
	p := &ast.File{}
	if ok := ast.ParentAs(n, &p); ok {
		return p
	}

	return nil
}

// FullQualifiedName tries to resolve and return the full qualified name (<package>.<Identifier>).
func FullQualifiedName(n ast.NamedType) string {
	return Pkg(n).Path + "." + n.Identifier()
}

// TypeDecl tries to resolve and returns a type declaration from the given named type.
func TypeDecl(n ast.NamedType) ast.TypeDecl {
	return ast.NewSimpleTypeDecl(ast.Name(FullQualifiedName(n)))
}

// UseTypeDeclIn does a quite complex job. It looks up the different parts
// of 'what' and creates a new TypeDecl to be used in another package ('where').
// In languages with cycles, import definitions (and qualified names) may
// even interchange - per generic declaration!
func UseTypeDeclIn(what ast.TypeDecl, where *ast.Pkg) ast.TypeDecl {
	newTypeDecl := what.Clone()
	err := ast.ForEach(newTypeDecl, func(n ast.Node) error {
		switch t := n.(type) {
		case *ast.SimpleTypeDecl:
			// only rewrite local simple names
			if foundType := ResolveLocal(what, string(t.SimpleName)); foundType != nil {
				foundPkg := Pkg(foundType)
				t.SimpleName = ast.Name(foundPkg.Path) + "." + t.SimpleName
			}
		}
		return nil
	})

	if err != nil {
		panic(err) // cannot happen
	}

	return newTypeDecl
}

// WrapNode envelopes the given ast.Node into a token.Node.
func WrapNode(n ast.Node) token.Position {
	return token.Position{BeginPos: token.Pos{
		File:   n.Pos().File,
		Offset: -1,
		Line:   n.Pos().Line,
		Col:    n.Pos().Col,
	},
		EndPos: token.Pos{
			File:   n.End().File,
			Offset: -1,
			Line:   n.End().Line,
			Col:    n.End().Col,
		},
	}
}

// MakeTypeDecl converts the more unspecific and generic adl type description into the ast model.
func MakeTypeDecl(t *adl.TypeDecl) ast.TypeDecl {
	if t.IsMap() {
		return ast.NewGenericDecl(ast.NewSimpleTypeDecl(stdlib.Map), MakeTypeDecl(t.TypeParams[0]), MakeTypeDecl(t.TypeParams[1]))
	}

	if t.IsPtr() {
		return ast.NewTypeDeclPtr(MakeTypeDecl(t.TypeParams[0]))
	}

	if t.IsSlice() {
		return ast.NewSliceTypeDecl(MakeTypeDecl(t.TypeParams[1]))
	}

	if t.IsArray() {
		l, err := strconv.Atoi(t.TypeParams[0].Name.String())
		if err != nil {
			panic(token.Explain(token.NewPosError(t.TypeParams[0].Name, "invalid array length")))
		}

		return ast.NewArrayTypeDecl(l, MakeTypeDecl(t.TypeParams[1]))
	}

	if len(t.TypeParams) == 0 {
		return ast.NewSimpleTypeDecl(ast.Name(t.Name.String()))
	}

	var typeParams []ast.TypeDecl
	for _, param := range t.TypeParams {
		typeParams = append(typeParams, MakeTypeDecl(param))
	}

	return ast.NewGenericDecl(ast.NewSimpleTypeDecl(ast.Name(t.Name.String())), typeParams...)
}

func LastPathSegment(path string) string {
	segments := strings.Split(path, "/")
	if len(segments) == 0 {
		return ""
	}

	return segments[len(segments)-1]
}

func CloneFuncSig(f *ast.Func) *ast.Func {
	c := ast.NewFunc(f.FunName).SetVisibility(f.FunVisibility)
	if f.Comment() != nil {
		c.SetComment(f.CommentText())
	}

	for _, param := range f.FunParams {
		p := ast.NewParam(param.ParamName, param.TypeDecl().Clone())
		if param.Comment() != nil {
			p.SetComment(param.CommentText())
		}

		c.AddParams(p)
	}

	for _, param := range f.FunResults {
		p := ast.NewParam(param.ParamName, param.TypeDecl().Clone())
		if param.Comment() != nil {
			p.SetComment(param.CommentText())
		}

		c.AddResults(p)
	}

	return c
}

// FindImplementations traverses the entire module to find structs which implement the given interface.
// Note that other types implementing an interface are (not yet) supported.
func FindImplementations(ctx ast.Node, iface ast.Name) []*ast.Struct {
	var r []*ast.Struct
	for _, pkg := range Mod(ctx).Pkgs {
		for _, file := range pkg.PkgFiles {
			for _, node := range file.Nodes {
				if s, ok := node.(*ast.Struct); ok {
					for _, implement := range s.Implements {
						if implement == iface {
							r = append(r, s)
						}
					}
				}
			}
		}
	}

	return r
}
