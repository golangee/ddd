package astutil

import (
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/src/ast"
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
