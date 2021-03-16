package metamodel

type FlagType int

const (
	Int64Flag FlagType = iota + 1
	Float64Flag
	StringFlag
	BoolFlag
	DurationFlag
)

// Flag
type Flag struct {
	// Name is the unique identifier of this flag.
	Name Identifier

	// Type of the actual flag.
	Type FlagType

	// Default contains a serialized string literal version of the default value.
	Default StrLit

	// Comment for this flag.
	Comment Text
}

// Flags is like a DTO but can only contain flags. Flags must not be used in a domain context. They should only
// be used and thought of like feature flags or secrets which are injected by the process environment.
type Flags struct {
	// Fields contains the actual flags.
	Fields []Flag

	// Comment for this flag set.
	Comment Text
}
