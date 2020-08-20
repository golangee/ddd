package ddd

type CRUDOptions int

const (
	CreateOne CRUDOptions = 1 << iota
	ReadOne
	ReadAll
	UpdateOne
	DeleteOne
	DeleteAll
	CountAll
)

const CRUD = ReadOne | CreateOne | DeleteOne | UpdateOne | DeleteAll | CountAll | ReadAll

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

func Interface(name string, methods ...*MethodSpecification) StructOrInterfaceOrFunc {
	return &InterfaceSpec{name: name, funcs: methods}
}

func Repository(typ TypeName, opts CRUDOptions, funcs ...*MethodSpecification) *InterfaceSpec {

}
