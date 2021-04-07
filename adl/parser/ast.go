package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"regexp"
)

var unqualifiedIdentifier = regexp.MustCompile("[a-zA-Z]\\w*")

type Qualifier string

type Identifier string

func (id Identifier) Capture(values []string) error {
	if unqualifiedIdentifier.FindString(values[0]) != values[0] {
		return fmt.Errorf("identifier '" + values[0] + "' must be unqualified")
	}

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
	Pos      lexer.Position
	Claim *Annotation `("claim" @@)?`
	Name     Identifier  `"struct" @Ident`
	Fields   []*Field    `"{" @@* "}"`
}

type Annotation struct {
	Name string `@Ident`
}

type Interface struct {
	Pos     lexer.Position
	Name    Identifier `"interface" @Ident`
	Methods []*Method  `"{" @@* "}"`
}

type Method struct {
	Pos  lexer.Position
	Name Identifier `@Ident`
}

type Field struct {
	Pos  lexer.Position
	Name Identifier `@Ident`
	Type string     `@Ident`
}

type Requirement struct {
	Pos   lexer.Position
	Name  Qualifier `"requirements" @Ident`
	Epics []*Epic   `@@*`
}

type Epic struct {
	Name        Identifier `"epic" @Ident`
	Description string     `@String`
	Stories []*Story   `@@*`
}


type Story struct {
	Name        Identifier `"story" @Ident`
	Description string     `@String`
}