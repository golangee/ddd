package metamodel

// Stereotype specifies different kinds of basic types.
type Stereotype int

// IsClass returns true, if this Stereotype denotes a class specifier.
func (s Stereotype) IsClass() bool {
	return s >= Value && s <= Stub
}

// IsPackage returns true, if this Stereotype denotes a package specifier.
func (s Stereotype) IsPackage() bool {
	return s >= UseCaseLayer && s <= DependencyLayer
}

const (

	// A Stub is a partially defined by architect and is more like an abstract class and needs to be completed by
	// human developer later.
	Stub

	UseCaseLayer
	DomainLayer
	PresentationLayer
	PersistenceLayer
	AdapterLayer
	DependencyLayer
)

// compound represents a compound data type und in the sense of UML maps also to a class but with some custom
// stereotypes. Valid Stereotypes are Value, Entity, DataTransferObject and Stub.
type compound struct {
	// Parent if always a package.
	Parent *Package

	// Name of the class.
	Name Identifier

	// Comment in various translations for this type.
	Comment Text

	// Fields of this compound type.
	Fields []*Field
}

// A Value class is defined by an architect and has no identity.
type Value struct {
	compound
}

// A DTO (data transfer object) is defined by an architect and not used for storing but instead for
// transferring information from one context or layer into another.
type DTO struct {
	compound
}

// An Entity class is defined by an architect and has an Id. Some say it should be a kind of a singleton,
// but that definition may be false guiding. Also it may be treated technically as a value (stack allocated)
// type.
type Entity struct {
	compound

	// ID field is a shortcut to this Entity.
	ID *Field

	// Injections declare required interfaces which will be reflected by a constructor.
	// This should be used with care and is only relevant for non-anemic Entities.
	Injections []*Interface

	// Methods of this Entity type. Only relevant for non-anemic Entities.
	Methods []*Method
}

// A Service is mostly a stub in the domain or use case layer and must be completed by a developer.
type Service struct {
	// Injections declare required interfaces which will be reflected by a constructor.
	Injections []*Interface

	// Methods of this service type. Only relevant for non-anemic Entities.
	Methods []*Method
}


