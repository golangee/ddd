package ddd

type DomainSpec struct {
}

func Context(name string, comment string, domainCore ...*LayerSpec) *DomainSpec {
	return &DomainSpec{}
}
