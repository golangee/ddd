package spec

// A Stub provides method declarations without implementations, attributes and constructors.
// The actual implementation will be implemented, tested and debugged by a developer any time
// later.
type Stub struct {
	Name         Identifier   // Unique identifier of this fragment.
	Comment      String       // Comment should describe why this fragment exists.
	Dependencies []Identifier // Identifiers of declared types from the BoundedContext. You cannot depend on others.
}
