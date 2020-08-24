package ddd

type LayerSpec struct {
	name      string
	comment   string
	api       []StructOrInterface
	factories []FuncOrStruct
}

// UseCases can only ever import DomainCore API
func UseCases(api []StructOrInterface, factories []FuncOrStruct) *LayerSpec {
	return Layer("UseCases", "...all the use cases of the domain.", api, factories...)
}

// DomainCore has never any dependencies to any other layer.
func DomainCore(api []StructOrInterface, factories []FuncOrStruct) *LayerSpec {
	return Layer("DomainCore", "...all the core domain API of the domain.", api, factories...)
}

func Layer(name, comment string, api []StructOrInterface, factories ...FuncOrStruct) *LayerSpec {
	return &LayerSpec{
		name:      name,
		comment:   comment,
		api:       api,
		factories: factories,
	}
}
