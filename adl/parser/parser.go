package parser

import (
	"bytes"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer/stateful"
)

const (
	// sLocalSelector selects something from an unknown local variable like .name or .Field.Other.Member.
	sLocalSelector = `\.(\w|\.)+`

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

func Parse(fname, src string) (*File, error) {
	lexer := stateful.MustSimple([]stateful.Rule{
		{"comment", `//.*|/\*.*?\*/`, nil},

		// parseable documentation style
		{"DocTitle", `^[[:blank:]]*# =[^=].*`, nil},
		{"DocSubTitleParameters", `^[[:blank:]]*# == Parameters`, nil},
		{"DocSubTitleReturns", `^[[:blank:]]*# == Returns`, nil},
		{"DocSubTitleErrors", `^[[:blank:]]*# == Errors`, nil},
		{"DocSection", `^[[:blank:]]*# ==.*`, nil},
		{"DocSeePrefix", `^[[:blank:]]*# see\s`, nil},
		{"DocSummary", `^[[:blank:]]*# \.\.\.[a-zA-Z]+.*\.`, nil},
		{"DocListLevel0", `^[[:blank:]]*# \* [a-zA-Z].*`, nil},
		{"DocIndentLevel0", `^[[:blank:]]*#\s{3}[a-zA-Z].*`, nil},
		{"DocText", `^[[:blank:]]*#.*`, nil},

		{"BoolLit","true|false",nil},

		// dots is ambiguous in Go and weired in Java, so using rusts :: seems like a good idea
		{"PkgSep", "::", nil},
		{"UrlSep", "/", nil},
		{"UrlVarSep", ":", nil},
		{"MacroSep", "!", nil},
		{"Optional", `\?`, nil},
		{"ParamSep", `\$`, nil},
		{"Sel", `\.`, nil},
		{"Keyword", ":claim|=>|->", nil},
		{"SumType", `\|`, nil},
		{"TypeParam", `<|>`, nil},
		{"Pointer", `\*`, nil},
		//{"LocalSelector", sLocalSelector, nil},
		//{"Qualifier", sQualifier, nil},
		{"Ident", `([a-zA-Z_][a-zA-Z0-9_]*)`, nil},
		{"SliceLooper", sLocalSlice, nil},
		//{"SliceType", `\[\]`, nil},
		{"OpenSlice", `\[`, nil},
		{"CloseSlice", `\]`, nil},
		{"String", sString, nil},
		{"IntLiteral", `[0-9]+`, nil},

		{"Term", `[=,(){}@]`, nil},
		{"whitespace", `\s+`, nil},
	})
	_ = lexer

	parser := participle.MustBuild(&File{},
		participle.Lexer(lexer),
		participle.Unquote("String"),
		participle.UseLookahead(3),
	)

	//fmt.Println(parser.String())

	ast := &File{}
	buf := bytes.NewReader([]byte(src))
	return ast, parser.Parse(fname, buf, ast)
}
