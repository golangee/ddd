package metamodel

// An Actor is a system or role.
type Actor struct {
	// Name of the Actor.
	Name Identifier

	// Description of the Actor.
	Comment Text
}

// A UseCase connects an actor, a method
type UseCase struct {
	// Actor of this use case.
	Actor *Actor

	// Includes is a list of other included use cases. These are straight forward dependencies.
	Includes []*UseCase

	// Extends use cases actually require (usually) a condition.
	Extends []*UseCase
}
