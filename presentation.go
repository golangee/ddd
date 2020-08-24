package ddd

type PresentationSpec struct {
	version string
	restSpecs []*HttpResourceSpec
}

type PresentationSpecs struct {
	presentations []*PresentationSpec
}

