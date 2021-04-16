package parser

import "github.com/alecthomas/participle/v2/lexer"

// Presentation contains the driver adapters.
type Presentation struct {
	Rest *Rest `("rest" "{" @@ "}")?`
}

type Rest struct {
	Version RestMajorVersion `@@*`
}

type RestMajorVersion struct {
	Version SemVer `@@ "{" `

	Types []JsonObject `@@*`

	Endpoints []RestEndpoint `@@* "}"`
}

type RestEndpoint struct {
	Usecases []DocSee  `@@*`
	Path     URLPath   `@@ "{" `
	Head     *HttpVerb `("HEAD" @@)? `
	Options  *HttpVerb `("OPTIONS" @@)? `
	Get      *HttpVerb `("GET" @@)? `
	Post     *HttpVerb `("POST" @@)? `
	Put      *HttpVerb `("PUT" @@)? `
	Patch    *HttpVerb `("PATCH" @@)? `
	Delete   *HttpVerb `("DELETE" @@)? "}"`
}

type URLPath struct {
	Elements []IdentOrVar ` @@ ("/" @@)* `
}

type IdentOrVar struct {
	Ident *Ident `parser:"@@" json:",omitempty"`
	Var   *Ident `parser:"|\":\" @@" json:",omitempty"`
}

type HttpVerb struct {
	// ContentType indicates the content type which the
	// client accepts and the server has to produce to
	// fulfill the request.
	ContentType String         `@@ "{" `
	In          []HttpInParam  `"in" "{" @@* "}" `
	Out         []HttpOutParam `"out" "{" @@*  `
	Errors      []HttpOutError `"errors" "{" @@* "}" "}" "}"`
}

type JsonObject struct {
	Pos lexer.Position
	// Doc contains a summary, arbitrary text lines, captions, sections and more.
	Doc    DocTypeBlock `parser:"@@"`
	Name   Ident        `"json" @@`
	Fields []*JsonField `"{" @@* "}"`
}

// A JsonField does not have a documentation, because it must "borrow"
// that from existing
type JsonField struct {
	Pos  lexer.Position
	Name String        `@@`
	Type TypeWithField `@@`
}

type HttpInParam struct {
	CopyDoc DocSee    `@@`
	Name    Ident     `@@`
	Type    Type      `@@ "="`
	From    HttpParam `@@`
}

type HttpParam struct {
	Header    *String `("HEADER" "[" @@ "]")`
	Path      *String `|("PATH" "[" @@ "]")`
	Query     *String `|("QUERY" "[" @@ "]")`
	IsBody    bool    `parser:"|@\"BODY\"" json:",omitempty"`
	IsRequest bool    `parser:"|@\"REQUEST\"" json:",omitempty"`
}

type HttpOutParam struct {
	CopyDoc  DocSee            `@@`
	Header   *HttpHeaderAssign `parser:"( (\"HEADER\" @@ )" json:",omitempty"`
	Body     *HttpAssign       `parser:"|(\"BODY\" \"=\" @@)" json:",omitempty"`
	Response *HttpAssign       `parser:"|(\"RESPONSE\" \"=\" @@))" json:",omitempty"`
}

type HttpHeaderAssign struct {
	Key   String `"[" @@ "]" "="`
	Ident Ident  `@@`
	Type  Type   `@@`
}

type HttpAssign struct {
	Ident Ident `@@`
	Type  Type  `@@`
}

type HttpOutError struct {
	Status Int                    `@@ "for"`
	Match  PathWithMemberAndParam `@@`
}
