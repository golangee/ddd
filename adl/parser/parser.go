package parser

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

const (
	// sLocalSelector selects something from an unknown local variable like .name or .Field.Other.Member.
	sLocalSelector = `\.(\w|.)+`

	// sQualifier selects things like "identifier", "github.com/so_me-thing.else" or "rust::like.qualifier".
	sQualifier = `[a-zA-Z](\w|\.|/|:|-)*`

	// sLocalSlice is a placeholder for an unknown local slice or array loop index variable.
	sLocalSlice = `\[i\]`

	// sString denotes an arbitrary string with quoted ", e.g. 'hello world' or 'hello\"world\"'
	sString = `"(\\"|[^"])*"`

	// sIdentifier is not part of the lexer because it is already a subset of sQualifier and ambiguities cannot
	// be matched properly. This happens directly at Identifier.
	sIdentifier = `[a-zA-Z]\w*`
)

func Parse(src string) (*File, error) {
	lexer := stateful.MustSimple([]stateful.Rule{
		{"Keyword", "claim|<-|->", nil},
		{"LocalSelector", sLocalSelector, nil},
		{"Qualifier", sQualifier, nil},
		{"LocalSlice", sLocalSlice, nil},
		{"String", sString, nil},

		{"Term", `[=,(){}@]`, nil},
		{"whitespace", `\s+`, nil},
	})
	_ = lexer

	parser := participle.MustBuild(&File{},
		participle.Lexer(lexer),
		participle.Unquote("String"),
		participle.UseLookahead(2),
	)

	fmt.Println(parser.String())

	ast := &File{}
	buf := bytes.NewReader([]byte(src))
	return ast, parser.Parse("test", buf, ast)
}
