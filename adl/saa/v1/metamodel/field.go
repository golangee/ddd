package metamodel

// A Field is part of type definition. The parent is one of Stub|Value|Entity|DTO.
type Field struct {
	// Parent declares the parent of this of field.
	Parent interface{}

	// Name is the unique identifier of this field.
	Name Identifier

	// Optional is mostly an indicator for a nullable field.
	Optional bool

	// Property signals that this is not only a simple field but should provide things like listeners for
	// data binding etc.
	Property bool

	// Comment for this field.
	Comment Text
}
