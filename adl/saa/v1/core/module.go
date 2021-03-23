package core

// A Module is a homogenous software building block. Usually maps to things like a gradle module, a maven artifact
// or a Go module.
type Module struct {
	Packages []*Package
	Name     StrLit
}

func NewModule(name StrLit) *Module {
	return &Module{Name: name}
}
