package ddd



type InterfaceSpec struct {
	name  string
	funcs []*MethodSpecification
}

func (s *InterfaceSpec) structOrInterface() {
	panic("implement me")
}

func (s *InterfaceSpec) Name() string {
	return s.name
}

func Interface(name,comment string, methods ...*MethodSpecification) StructOrInterface {
	return &InterfaceSpec{name: name, funcs: methods}
}

