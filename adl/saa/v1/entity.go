package spec

// An Entity is a complex data class which has an ID like an account.
type Entity struct {
	Name    Identifier // Unique identifier of this fragment.
	Comment String     // Comment should describe why this fragment exists.
}
