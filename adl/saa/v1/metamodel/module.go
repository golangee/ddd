package metamodel

// A Module is a homogenous software building block. Usually maps to things like a gradle module, a maven artifact
// or a Go module.
type Module struct {
	// Name of the module in path notation.
	Name Path

	// Comment describes why this module exists and what it is for.
	Comment Text
}
