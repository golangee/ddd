package parser

import (
	"fmt"
	"github.com/alecthomas/participle/v2/lexer"
	"regexp"
	"strings"
)

var rIdentifier = regexp.MustCompile(sIdentifier)

// String is just anything like sString accepts.
type String struct {
	//Pos   lexer.Position
	Value string `@String`
}

// Doc is just anything like sString accepts.
type Doc struct {
	//Pos   lexer.Position
	Value string `@String`
}

func (d *Doc) Capture(values []string) error {
	d.Value = values[0]

	if !strings.HasPrefix(d.Value, "...") {
		return fmt.Errorf("a documentation must start with ellipsis (...)")
	}

	if !strings.HasSuffix(d.Value, ".") {
		return fmt.Errorf("a documentation must end with a dot")
	}

	if len(d.Value) < 6 {
		return fmt.Errorf("the documentation is to short")
	}

	return nil
}

type Path struct {
	//Local bool `@PkgSep?`
	Elements []Ident ` @@ ("::" @@)* `
}

type Type struct {
	Pointer   bool   `(@Pointer)?`
	Qualifier Path   `@@`
	Params    []Type `("<" @@ ("," @@)* ">")?`
}

type Ident struct {
	//Pos   lexer.Position
	Value string `@Ident`
}

// Qualifier is just anything like sQualifier accepts.
type Qualifier struct {
	//Pos         lexer.Position
	IsSliceType bool   `@SliceType?`
	Value       string `@Qualifier`
	IsStdType   bool   `@StdlibType?`
}

// Identifier is a subset of Qualifier which must match identifier rules which are similar but
// not equal. Without a regex lookahead this seems hard to match at the lexer level.
type Identifier struct {
	//Pos   lexer.Position
	Value string `@Qualifier`
}

func (id *Identifier) Capture(values []string) error {
	if rIdentifier.FindString(values[0]) != values[0] {
		return fmt.Errorf("invalid identifier: " + values[0])
	}

	id.Value = values[0]
	return nil
}

// File represents a source code file.
type File struct {
	//	Pos     lexer.Position
	Modules []*Module `@@*`
	//Requirements []*Requirement `@@`
}

// A Module is a distinct unit of generation. Usually it maps only to a single
// service like a Go microservice but there may be also library modules which may
// get emitted into different targets.
type Module struct {
	// Claims to requirements this element.
	Claims []*Claim `(":claim" @@)*`
	// Doc of the module.
	Doc Doc `@@`
	// Name of the module.
	Name     Ident      `"module" @@ "{" `
	Generate *Generate  `"generate" "{" @@ "}" `
	Contexts []*Context `@@* "}"`
	//Packages []*Package `"{" @@ "}"`
	//Parts *ModuleParts `"{" @@ "}"`
}

// Generate describes multiple emitter configurations.
type Generate struct {
	Go *GeneratorGo `("go" "{" @@ "}")?`
}

// GeneratorGo describes all go specific generator settings.
type GeneratorGo struct {
	// Module is the go.mod name and prefix for all packages.
	Module String `("module" "=" @@)`
	// Output is either an absolute or relative local path to emit the source.
	Output String `("output" "=" @@)`
}

// Claim is a reference to a requirement like an epic, a story, a scenario, a glossary entry or
// just a requirement id.
type Claim struct {
	//Pos  lexer.Position
	Path Path `@@`
}

// A Context describes the top-level grouping structure in DDD.
type Context struct {
	Pos lexer.Position
	// Claims to requirements of the context.
	Claims []*Claim `(":claim" @@)*`
	// Documentation of the context (usually a package comment).
	Doc Doc `@@`
	// Name of the context package
	Name Ident `"context" @@ "{" `
	// Domain with core and application layer
	Domain         *Domain         `"domain" "{" @@ "}" `
	Infrastructure *Infrastructure `"infrastructure" "{" @@ "}" "}"`
	//Presentation   *Presentation   `"presentation" "{" @@ "}" "}"`
}

// Domain contains the application (use case) and the core layer (packages).
type Domain struct {
	Pos        lexer.Position
	Core       *Core        `"core" "{" @@ "}"`
	UseCase    *Usecase     `"usecase" "{" @@ "}"`
	Subdomains []*Subdomain `@@*`
	//Types   []*TypeDef ` @@* "}"`
}

// Core is together with any subdomain packages self-containing and creates the
// base for any another layer.
type Core struct {
	Pos   lexer.Position
	Types []*TypeDef `@@*`
}

type Usecase struct {
	Pos lexer.Position
	//Types []*TypeDef ` "{" "}"`
	Bla []*String `@@*`
}

// Subdomain contains the application (use case) and the core layer (packages).
type Subdomain struct {
	Pos lexer.Position
	// Claims to requirements of the subdomain.
	Claims []*Claim `(":claim" @@)*`
	// Documentation of the subdomain (usually a package comment).
	Doc *String `@@`

	Name    Ident    `"subdomain" @@ "{"`
	Core    *Core    `"core" "{" @@ "}"`
	UseCase *Usecase `"usecase" "{" @@ "}" "}"`
}

// Infrastructure helps with additional hints to generate stuff like SQL or Event adapter.
type Infrastructure struct {
	MySQL *SQL `("mysql" "{" @@ "}")?`
	//MySQL *SQL `"mysql" "{" @@ "}"`
}

// SQL contains sql independent declaration, but the generators are dialect specific.
type SQL struct {
	Database   String               `"database" "=" @@`
	Implements []*SQLImplementation `@@+`
}

type SQLImplementation struct {
	Type    Path       `"impl" @@`
	SQLFunc []*SQLFunc `"{" @@* "}"`
}

type SQLFunc struct {
	Name Ident           `@@`
	SQL  String          `@@`
	In   []SQLFuncInVar  `( "(" @@ ("," @@)* ")" )?`
	Out  []SQLFuncOutVar `( "=>" "(" @@ ("," @@)* ")" )?`
}

type SQLFuncInVar struct {
	Selector    []LooperIdent `( @@ ("." @@)* )?`
	IsLooper bool  `(@SliceLooper)?`
}

type LooperIdent struct{
	Ident Ident `@@`
	IsLooper bool  `(@SliceLooper)?`
}

type SQLFuncOutVar struct {
	Ident Ident `"." @@?`
}

type TypeDef struct {
	Pos       lexer.Position
	Struct    *Struct    ` @@`
	Repository *Repository `| @@`
	Interface *Interface `| @@`
}

type Interface struct {
	Pos     lexer.Position
	// Claims to requirements of the subdomain.
	Claims []*Claim `(":claim" @@)*`
	Doc     String    `@@`
	Name    Ident     `"interface" @@`
	Methods []*Method `"{" @@* "}"`
}

type Repository struct {
	Pos     lexer.Position
	// Claims to requirements.
	Claims []*Claim `(":claim" @@)*`
	Doc     Doc    `@@`
	Name    Ident     `"repository" @@`
	Methods []*Method `"{" @@* "}"`
}

type Method struct {
	Pos    lexer.Position
	// Claims to requirements.
	Claims []*Claim `(":claim" @@)*`
	Doc    Doc     `@@`
	Name   Ident    `@@`
	Params []*Param `"(" @@? | ("," @@)* ")"`
	Return *Type `"->" "(" @@? `
	Error *Error ` ("," @@)?  ")"`
}

type Error struct{
	Kinds []Ident `"error" "<" @@ ("|" @@)* ">"`
}

type Param struct {
	Name Ident `@@`
	Type Type  `@@`
}

type Struct struct {
	Pos lexer.Position
	// Claims points to requirements of this element.
	Claims []*Claim `(":claim" @@)*`
	Doc    Doc      `@@`
	Name   Ident    `"struct" @@`
	Fields []*Field `"{" @@* "}"`
}

type Field struct {
	Pos lexer.Position
	// Claims points to requirements of this element.
	Claims []*Claim `(":claim" @@)*`
	Doc    Doc      `@@`
	Name   Ident    `@@`
	Type   Type     `@@`
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

/*
type Subdomain struct {
	Pos     lexer.Position
	Name    Qualifier  `"subdomain"  "{"`
	Core    *Core      `"core" "{" @@ "}"`
	UseCase *UseCase   `"usecase" "{" @@ "}"`
	Types   []*TypeDef ` @@* "}"`
}*/

type UseCase struct {
	Types []*TypeDef ` @@*`
}

type Presentation struct {
	Types []*TypeDef ` @@*`
}

type Property struct {
	Name  Identifier `@Ident "="`
	Value string     `@String`
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
