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

// FindStructs applies the predicate on each contained struct. However,
// only non-raw types can be inspected.
func (c Pkg) FindStructs(predicate func(s Struct) bool) []*ast.Struct {
	var r []*ast.Struct
	for _, file := range c.obj.PkgFiles {
		for _, namedType := range file.Types() {
			if s, ok := namedType.(*ast.Struct); ok {
				if predicate(StructFrom(s)) {
					r = append(r, s)
				}
			}
		}
	}

	return r
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
