package ddd

type InterfaceSpec struct {
	name  string
	funcs []*MethodSpecification
}

func Interface(name string, methods ...*MethodSpecification) *InterfaceSpec {
	return &InterfaceSpec{name: name, funcs: methods}
}
