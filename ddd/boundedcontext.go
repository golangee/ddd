package ddd

// A BoundedContextSpec contains the information about the layers within a bounded context. It contains different
// layers which may dependencies with each other, depending on their stereotype.
type BoundedContextSpec struct {
	name        string
	description string
	layers      []Layer
}

// BoundedContexts is a factory to create a slice of BoundedContextSpec from variable amount of arguments.
func BoundedContexts(contexts ...*BoundedContextSpec) []*BoundedContextSpec {
	return contexts
}

// Context is a factory for a BoundedContextSpec.
func Context(name string, description string, layers ...Layer) *BoundedContextSpec {
	return &BoundedContextSpec{
		name:        name,
		description: description,
		layers:      layers,
	}
}
