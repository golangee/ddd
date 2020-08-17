package ddd

type DomainSpec struct {
}

func Context(name string, comment string,requires *InterfaceSpecs, usecases *UseCaseSpecs, types *TypeSpecs, presentationSpec *PresentationSpecs, service *ServiceSpec, persistence *PersistenceSpec) *DomainSpec {
	return &DomainSpec{}
}
