package metamodel

// Interface represents a contract for behavior.
type Interface struct {
	// Parent is always a package.
	Parent *Package

	// Name of the interface.
	Name Identifier

	// Comment in various translations for this type.
	Comment Text

	// Methods of this interface type.
	Methods []*Method
}

// Method defines an interface method specification.
type Method struct {
	// Parent is a bit arbitrary?
	Parent interface{}

	// Name of the interface.
	Name Identifier

	// Comment in various translations for this type.
	Comment Text


	In  []Param
	Out []Param
}

// Param defines a method parameter.
type Param struct {
	// Parent declares the parent of this of parameter.
	Parent interface{}

	// Name is the unique identifier of this parameter.
	Name Identifier

	// Optional is mostly an indicator for a nullable parameter.
	Optional bool

	// Comment for this parameter.
	Comment Text
}
