package core

// A Text can be translated into multiple languages and is not required to serve as something else than
// to be read by a human.
type Text struct {
	Default string // Default denotes the default value which shall be english.
	Values  map[Locale]struct {
		Value string
		Pos
	}
}

// String returns the default language value.
func (s Text) String() string {
	return s.Default
}

// Localize tries to localize the value or returns the default. We intentionally do not fuzz around with all
// that complex CLDR things, because that can be normalized/checked/fixed at meta model creation time once.
func (s Text) Localize(locale Locale) string {
	if s.Values != nil {
		v, ok := s.Values[locale]
		if ok {
			return v.Value
		}
	}

	return s.Default
}
