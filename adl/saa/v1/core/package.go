package core

// Package contains all relevant architecture elements.
type Package struct {
	Parent *Module
	Path   StrLit
}

func NewPackage(parent *Module, path StrLit) *Package {
	return &Package{
		Parent: parent,
		Path:   path,
	}
}
