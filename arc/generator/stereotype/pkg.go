package stereotype

import "github.com/golangee/src/ast"

// Pkg contains all stereotype annotations for a package instance.
type Pkg struct {
	obj *ast.Pkg
}

func PkgFrom(pkg *ast.Pkg) Pkg {
	return Pkg{obj: pkg}
}

func (c Pkg) Unwrap() *ast.Pkg {
	return c.obj
}

// SetIsCMDPkg marks this package a main package.
func (c Pkg) SetIsCMDPkg(isPublicConfig bool) Pkg {
	c.obj.PutValue(kCMDPkg, isPublicConfig)
	return c
}

// IsCMDPkg returns only true, if the package shall be main package.
func (c Pkg) IsCMDPkg() bool {
	v := c.obj.Value(kCMDPkg)
	if f, ok := v.(bool); ok {
		return f
	}

	return false
}

// WithPkg visits recursively all Pkg elements.
func WithPkg(n ast.Node, f func(pkg Pkg) error) error {
	return ast.ForEach(n, func(n ast.Node) error {
		if pkg, ok := n.(*ast.Pkg); ok {
			if err := f(PkgFrom(pkg)); err != nil {
				return err
			}
		}

		return nil
	})
}

// FindCMDPkgs returns all annotated CMDPkg.
func FindCMDPkgs(n ast.Node) []*ast.Pkg {
	var r []*ast.Pkg
	if err := WithPkg(n, func(pkg Pkg) error {
		if pkg.IsCMDPkg() {
			r = append(r, pkg.Unwrap())
		}

		return nil
	}); err != nil {
		panic(err)
	}

	return r
}
