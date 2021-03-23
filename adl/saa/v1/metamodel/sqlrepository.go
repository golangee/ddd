package metamodel

// A SQLRepository declares a repository and methods with data bindings.
type SQLRepository struct {
	Name    Identifier
	Methods []SQLMethod
}

// Method declares a method name, the according query and prepare and map bindings. These only make sense in
// the given context.
type SQLMethod struct {
	Name    Identifier
	Query   StrLit
	Prepare []Identifier
	Map     []Identifier
}
