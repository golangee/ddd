package parser

type DocTypeBlock struct {
	Summary string        `@DocSummary`
	Elems   []DocTypeElem `@@*`
}

type DocTypeElem struct {
	Title     *DocTitle `parser:"@@" json:",omitempty"`
	Text      *DocText  `parser:"|@@" json:",omitempty"`
	Reference *DocSee   `parser:"|@@" json:",omitempty"`
}

type DocMethodBlock struct {
	Summary string          `@DocSummary`
	Elems   []DocMethodElem `parser:"@@*" json:",omitempty"`
}

type DocMethodElem struct {
	Title         *DocTitle      `parser:"@@" json:",omitempty"`
	DocParameters *DocParameters `parser:"|@@" json:",omitempty"`
	DocReturns    *DocReturns    `parser:"|@@" json:",omitempty"`
	DocErrors     *DocErrors     `parser:"|@@" json:",omitempty"`
	Text          *DocText       `parser:"|@@" json:",omitempty"`
	Reference     *DocSee        `parser:"|@@" json:",omitempty"`
}

// DocTitle is of the form # = <Title>
type DocTitle struct {
	//Pos   lexer.Position
	Value string `@DocTitle`
}

// DocSee is of the form # see <Path>
type DocSee struct {
	//Pos   lexer.Position
	Value string `@DocSeePrefix`
	Path  Path   `@@`
}

type DocParameters struct {
	Value  string     `@DocSubTitleParameters`
	Params []DocParam `(@@)*`
}

type DocReturns struct {
	Value string    `@DocSubTitleReturns`
	Text  []DocText `(@@)*`
}

type DocErrors struct {
	Value  string     `@DocSubTitleErrors`
	Params []DocParam `(@@)*`
}

// TODO we need more validation stuff
type DocParam struct {
	Summary string            `@DocListLevel0`
	More    []DocIndentLevel0 `parser:"(@@)*" json:",omitempty"`
}

type DocListItem0 struct {
	Value string            `@DocListLevel0`
	More  []DocIndentLevel0 `(@@)*`
}

type DocIndentLevel0 struct {
	Value string `@DocIndentLevel0`
}
