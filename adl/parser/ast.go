package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"regexp"
)

var unqualifiedIdentifier = regexp.MustCompile("[a-zA-Z]\\w*")

type Qualifier string

type Identifier string

func (id *Identifier) Capture(values []string) error {
	if unqualifiedIdentifier.FindString(values[0]) != values[0] {
		return fmt.Errorf("identifier '" + values[0] + "' must be unqualified")
	}

	*id = Identifier(values[0])
	return nil
}

type File struct {
	Pos          lexer.Position
	Modules      []*Module      `@@`
	Requirements []*Requirement `@@`
}

type Module struct {
	Name     Qualifier  `"module" @Ident`
	Packages []*Package `"{" @@ "}"`
}

type ModuleParts struct{

}

type Package struct {
	Pos   lexer.Position
	Name  Qualifier  `"package" @Ident`
	Types []*TypeDef `"{" @@* "}"`
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
