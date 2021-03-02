package spec

// A Value is a simple data class without identity, like an email.
type Value struct {
	Name    Identifier // Unique identifier of this fragment.
	Comment String     // Comment should describe why this fragment exists.
}
