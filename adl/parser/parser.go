package parser

import (
	"bytes"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

func Parse(src string) (*File, error) {
	lexer := stateful.MustSimple([]stateful.Rule{
		{"Keyword", "claim", nil},
		{"Ident", `[a-zA-Z](\w|\.|/|:)*`, nil},
	//	{"Ident", `[a-zA-Z]\w*`, nil},
		{"String",`"(\\"|[^"])*"`,nil},
		//{"String2","`(\\`|[^`])*`",nil},
		//{"MultilineString",`[.]{4}([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*[.]{4}`,nil},
		//{"StopML", `[.]{4}\n`, nil},
		{"Term", `[{}@]`, nil},
		{"whitespace", `\s+`, nil},
	})
	_ = lexer

	parser := participle.MustBuild(&File{},
		participle.Lexer(lexer),
			participle.Unquote("String"),
		//	participle.UseLookahead(2),
	)

	ast := &File{}
	buf := bytes.NewReader([]byte(src))
	return ast, parser.Parse("test", buf, ast)
}
