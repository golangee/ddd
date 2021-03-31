package astutil

import (
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/src/ast"
	"strings"
)

func FindMod(name core.StrLit, prj *ast.Prj) (*ast.Mod, error) {
	for _, mod := range prj.Mods {
		if mod.Name == name.String() {
			return mod, nil
		}
	}

	return nil, core.NewPosError(name, "module not found")
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

// Resolve takes the name and walks until it finds whatever declares it. Returns nil
// if not found.
func Resolve(n ast.Node, name string) ast.Node {
	lastDot := strings.LastIndex(name, ".")
	lastSlash := strings.LastIndex(name, "/")

	// try local package search
	if lastDot < 0 && lastSlash < 0 {
		for _, node := range Pkg(n).Children() {
			if tname, ok := node.(ast.NamedType); ok {
				if tname.Identifier() == name {
					return tname
				}
			}
		}
	}

	// resolve package and return type from that
	if lastDot > lastSlash {
		pkgName := name[:lastDot]
		typeName := name[lastDot+1:]
		for _, pkg := range Mod(n).Pkgs {
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
