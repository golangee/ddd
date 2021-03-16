package metamodel

// Package contains all relevant architecture elements.
type Package struct {
	// Stereotype to describe the purpose of this package.
	Stereotype Stereotype

	// Parent of a Package is always a Module.
	Parent *Module

	// Name of the Package in path notation.
	Name Path

	// Comment describes why this Package exists and what it is for.
	Comment Text

	// Classes in this package.
	Classes []*Class
}
