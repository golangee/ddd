package ddd

// AppSpec contains the information about the entire application and all its bounded contexts.
type AppSpec struct {
	name            string
	boundedContexts []*BoundedContextSpec
}

// Application is a factory to create an AppSpec.
func Application(name string, domains []*BoundedContextSpec) *AppSpec {
	return &AppSpec{
		name:            name,
		boundedContexts: domains,
	}
}
