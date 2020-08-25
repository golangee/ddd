package ddd

// A UseCaseLayerSpec represents a stereotyped USECASE layer.
type UseCaseLayerSpec struct {
	name        string
	description string
	api         []StructOrInterface
	factories   []FuncOrStruct
}

// Name of the layer.
func (u *UseCaseLayerSpec) Name() string {
	return u.name
}

// Description of the layer.
func (u *UseCaseLayerSpec) Description() string {
	panic("implement me")
}

// Stereotype of the layer.
func (u *UseCaseLayerSpec) Stereotype() Stereotype {
	return USECASE
}

// UseCases is a factory for a UseCaseLayerSpec. A use case can only ever import Core API.
func UseCases(api []StructOrInterface, factories []FuncOrStruct) *UseCaseLayerSpec {
	return &UseCaseLayerSpec{
		name: "usecases",
		description: "Package usecases contains all domain specific use cases for the current bounded context.\n" +
			"It contains an exposed public API to be imported by PRESENTATION layers. \n" +
			"It provides a private implementation of the use cases accessible by factory functions.",
		api:       api,
		factories: factories,
	}
}
