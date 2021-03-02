package spec

// FlagType declares an enum of different Flag kinds.
type FlagType int

const (
	FlagInt64 FlagType = iota + 1
	FlagFloat64
	FlagBool
	FlagDuration
	FlagString
)

type Flag struct {
	Name    Identifier
	Type    FlagType
	Comment String // Comment should describe why this fragment exists.
}

// A Configuration describes a few named primitive key/value flags.
type Configuration struct {
	Flags []*Flag
}
