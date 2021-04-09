package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"regexp"
)

var rIdentifier = regexp.MustCompile(sIdentifier)

// Qualifier is just anything like sQualifier accepts.
type Qualifier struct {
	Pos   lexer.Position
	Value string `@Qualifier`
}

// Identifier is a subset of Qualifier which must match identifier rules which are similar but
// not equal. Without a regex lookahead this seems hard to match at the lexer level.
type Identifier struct {
	Pos   lexer.Position
	Value string `@Qualifier`
}

func (id *Identifier) Capture(values []string) error {
	if rIdentifier.FindString(values[0]) != values[0] {
		return fmt.Errorf("invalid identifier: " + values[0])
	}

	id.Value = values[0]
	return nil
}

type File struct {
	Pos     lexer.Position
	Modules []*Module `@@*`
	//Requirements []*Requirement `@@`
}

type Module struct {
	Name     Qualifier  `"module" @@ "{" `
	Contexts []*Context `@@* "}"`
	//Packages []*Package `"{" @@ "}"`
	//Parts *ModuleParts `"{" @@ "}"`
}

type ModuleParts struct {
	Context  []*Context `@@*`
	Packages []*Package `@@*`
}

type Package struct {
	Pos   lexer.Position
	Name  Qualifier  `"package" @Ident`
	Types []*TypeDef `"{" @@* "}"`
}

type Context struct {
	Pos  lexer.Position
	Name Identifier `"context" @@ "{" "}"`
	Domain         *Domain         ` @@ `
	//Infrastructure *Infrastructure `"infrastructure" "{" @@ "}" `
	//Presentation   *Presentation   `"presentation" "{" @@ "}" "}"`
}

type Domain struct {
	Pos     lexer.Position
	//Core    *Core      `"domain" "{" "}"`
	//UseCase *UseCase   `"usecase" "{" @@ "}"`
	//Subdomains []*Subdomain `@@*`
	//Types   []*TypeDef ` @@* "}"`
}

type Subdomain struct {
	Pos     lexer.Position
	Name    Qualifier  `"subdomain"  "{"`
	Core    *Core      `"core" "{" @@ "}"`
	UseCase *UseCase   `"usecase" "{" @@ "}"`
	Types   []*TypeDef ` @@* "}"`
}

type Core struct {
	Types []*TypeDef ` @@*`
}

type UseCase struct {
	Types []*TypeDef ` @@*`
}

type Presentation struct {
	Types []*TypeDef ` @@*`
}

type Infrastructure struct {
	Types []*TypeDef ` @@*`
	MySQL *MySQL     `("mysql" "{" @@ "}")?`
}

type MySQL struct {
	DefaultDatabase Identifier `("db" "=" @String)?`
	//Properties []*Property `@@*`
	Implements []*SQLImplementation `@@+`
}

type SQLImplementation struct {
	Type    Qualifier  `"implements" @Ident`
	SQLFunc []*SQLFunc `"{" @@* "}"`
}

type SQLFunc struct {
	Name Identifier `@Ident`
	SQL  string     `@String`
	In   []Sel      `("->" @@* ("," @@)*  )?`
	Out  []Sel      `("<-" @@* ("," @@)*  )?`
}

type Property struct {
	Name  Identifier `@Ident "="`
	Value string     `@String`
}

type Sel struct {
	Pos   lexer.Position
	Value string `@Selector | @Ident`
}

type Ident struct {
	Pos   lexer.Position
	Value string `@Ident`
}

type TypeDef struct {
	Pos       lexer.Position
	Struct    *Struct    ` @@`
	Interface *Interface `| @@`
}

type Struct struct {
	Pos    lexer.Position
	Claim  *Claim     `("claim" @@)?`
	Doc    *Doc       `"..." @@`
	Name   Identifier `"struct" @Ident`
	Fields []*Field   `"{" @@* "}"`
}

type Claim struct {
	Pos   lexer.Position
	Ident string `@Ident`
}

type Doc struct {
	Pos   lexer.Position
	Value string `@String`
}

type Interface struct {
	Pos     lexer.Position
	Claim   *Claim     `("claim" @@)?`
	Doc     *Doc       `"..." @@`
	Name    Identifier `"interface" @Ident`
	Methods []*Method  `"{" @@* "}"`
}

type Method struct {
	Pos    lexer.Position
	Doc    *Doc       `"..." @@`
	Name   Identifier `@Ident`
	Params []*Param   `"(" @@* ("," @@)*  ")"`
}

type Param struct {
	Name Identifier `@Ident`
	Type Qualifier  `@Ident`
}

type Field struct {
	Pos  lexer.Position
	Doc  *Doc       `"..." @@`
	Name Identifier `@Ident`
	Type Qualifier  `@Ident`
}

type Requirement struct {
	Pos   lexer.Position
	Name  Qualifier `"requirements" @Ident`
	Epics []*Epic   `@@*`
}

type Epic struct {
	Name        Qualifier `"epic" @Ident`
	Description string    `@String`
	Stories     []*Story  `@@*`
}

type Story struct {
	Name        Qualifier   `"story" @Ident`
	Description string      `@String`
	Scenarios   []*Scenario `@@*`
}

type Scenario struct {
	Name        Qualifier `"scenario" @Ident`
	Description string    `@String`
}
