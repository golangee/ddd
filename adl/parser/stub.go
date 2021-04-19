package parser

// Service can be many things, like any stub. It defines members for
// configuration, injection and custom fields. Afterwards methods
// are defined just like in an interface.
type Service struct {
	// Doc contains a summary, arbitrary text lines, captions, sections and more.
	Doc       DocTypeBlock `parser:"@@"`
	Name      Ident        `"service" @@ "{"`
	Configure []*Field     `("configure" "{" @@* "}")?`
	Inject    []*Field     `("inject" "{" @@* "}")?`
	Private   []*Field     `("private" "{" @@* "}")?`
	Methods   []*Method    ` @@* "}"`
}
