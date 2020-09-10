package ddd

// EnvParams is just a StructSpec and to provide nicer autocompletion.
type EnvParams StructSpec

// InterfaceName is just a string naming an interface
type InterfaceName string

// ServiceImplSpec represents a factory for a service interface
type ServiceImplSpec struct {
	refersTo string
	requires []string
	options  *StructSpec
}

// Implementation specifies the dependencies for a service implementation of the given interface.
// Only interfaces of the same layer can be implemented and a package global factory function is generated for
// it, which you need to set as you like, e.g. by an init() method in a separate file.
func Implementation(of string, require []InterfaceName, params *EnvParams) *ServiceImplSpec {
	cfg := &ServiceImplSpec{
		refersTo: of,
		options:  (*StructSpec)(params),
	}
	for _, name := range require {
		cfg.requires = append(cfg.requires, string(name))
	}

	cfg.options.name = of + "Opts"
	cfg.options.comment = "... provides the options for creating a new instance of " + of + "."
	return cfg
}

// Of returns the interface for which this is a factory.
func (s *ServiceImplSpec) Of() string {
	return s.refersTo
}

// Requires returns the required contracts to inject.
func (s *ServiceImplSpec) Requires() []string {
	return s.requires
}

// Options returns the struct which represents the serializable options.
func (s *ServiceImplSpec) Options() *StructSpec {
	return s.options
}

// Requires assembles a slice of strings into InterfaceNames.
func Requires(interfaceNames ...string) []InterfaceName {
	var r []InterfaceName
	for _, name := range interfaceNames {
		r = append(r, InterfaceName(name))
	}
	return r
}

// Options is actually a factory for *StructSpec to declare your configuration. It will automatically emit
// json tags and adds helper methods to parse those parameters.
func Options(params ...*FieldSpec) *EnvParams {
	p := &EnvParams{
		fields: params,
		pos:    capturePos("Options", 1),
	}

	return p
}
