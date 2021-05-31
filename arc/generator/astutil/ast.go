package astutil

import (
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"strings"
)

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
