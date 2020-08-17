package ddd

type UseCaseName string

type UseCaseSpec struct {
	name     UseCaseName
	extends  []UseCaseName
	includes []UseCaseName
	actor    string
}

type UseCaseSpecs struct {
	specs []*MethodSpecification
}

func UseCases(usesCases ...*MethodSpecification) *UseCaseSpecs {
	return &UseCaseSpecs{specs: usesCases}
}

func UseCase(actor string, name UseCaseName) *UseCaseSpec {
	return &UseCaseSpec{actor: actor, name: name}
}
