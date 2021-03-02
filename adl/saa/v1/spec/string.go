package spec

// A String can be translated into multiple languages.
type String struct {
	defaultValue string
	values       map[Locale]struct {
		Value string
		Pos
	}
}

func (s String) String() string {
	return s.defaultValue
}
