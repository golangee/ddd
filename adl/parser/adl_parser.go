// Code generated from ADL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // ADL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 27, 221,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 54, 10, 2, 12, 2, 14,
	2, 57, 11, 2, 3, 2, 3, 2, 3, 2, 7, 2, 62, 10, 2, 12, 2, 14, 2, 65, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 78, 10, 4, 12, 4, 14, 4, 81, 11, 4, 3, 4, 5, 4, 84, 10, 4, 3, 5, 5,
	5, 87, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 103, 10, 9, 12, 9, 14, 9, 106, 11, 9,
	3, 9, 5, 9, 109, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	117, 10, 10, 3, 11, 3, 11, 5, 11, 121, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 7, 12, 128, 10, 12, 12, 12, 14, 12, 131, 11, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 140, 10, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 7, 14, 147, 10, 14, 12, 14, 14, 14, 150, 11, 14, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 161, 10,
	15, 3, 16, 3, 16, 5, 16, 165, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17,
	171, 10, 17, 12, 17, 14, 17, 174, 11, 17, 3, 17, 5, 17, 177, 10, 17, 5,
	17, 179, 10, 17, 3, 17, 3, 17, 3, 18, 5, 18, 184, 10, 18, 3, 18, 5, 18,
	187, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 194, 10, 19, 12,
	19, 14, 19, 197, 11, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 204,
	10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	214, 10, 23, 3, 24, 3, 24, 3, 24, 5, 24, 219, 10, 24, 3, 24, 2, 2, 25,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 2, 3, 4, 2, 3, 3, 13, 13, 2, 224, 2, 48, 3, 2, 2, 2, 4,
	68, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 90, 3, 2, 2,
	2, 12, 92, 3, 2, 2, 2, 14, 94, 3, 2, 2, 2, 16, 96, 3, 2, 2, 2, 18, 116,
	3, 2, 2, 2, 20, 120, 3, 2, 2, 2, 22, 122, 3, 2, 2, 2, 24, 134, 3, 2, 2,
	2, 26, 141, 3, 2, 2, 2, 28, 160, 3, 2, 2, 2, 30, 164, 3, 2, 2, 2, 32, 166,
	3, 2, 2, 2, 34, 183, 3, 2, 2, 2, 36, 190, 3, 2, 2, 2, 38, 198, 3, 2, 2,
	2, 40, 203, 3, 2, 2, 2, 42, 205, 3, 2, 2, 2, 44, 213, 3, 2, 2, 2, 46, 218,
	3, 2, 2, 2, 48, 49, 5, 4, 3, 2, 49, 55, 5, 46, 24, 2, 50, 51, 5, 6, 4,
	2, 51, 52, 5, 44, 23, 2, 52, 54, 3, 2, 2, 2, 53, 50, 3, 2, 2, 2, 54, 57,
	3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 63, 3, 2, 2, 2,
	57, 55, 3, 2, 2, 2, 58, 59, 5, 14, 8, 2, 59, 60, 5, 44, 23, 2, 60, 62,
	3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2,
	63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7,
	2, 2, 3, 67, 3, 3, 2, 2, 2, 68, 69, 7, 11, 2, 2, 69, 70, 7, 13, 2, 2, 70,
	5, 3, 2, 2, 2, 71, 83, 7, 14, 2, 2, 72, 84, 5, 8, 5, 2, 73, 79, 7, 26,
	2, 2, 74, 75, 5, 8, 5, 2, 75, 76, 5, 44, 23, 2, 76, 78, 3, 2, 2, 2, 77,
	74, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2,
	2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 84, 7, 27, 2, 2, 83, 72,
	3, 2, 2, 2, 83, 73, 3, 2, 2, 2, 84, 7, 3, 2, 2, 2, 85, 87, 9, 2, 2, 2,
	86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 5,
	12, 7, 2, 89, 9, 3, 2, 2, 2, 90, 91, 7, 22, 2, 2, 91, 11, 3, 2, 2, 2, 92,
	93, 5, 10, 6, 2, 93, 13, 3, 2, 2, 2, 94, 95, 5, 16, 9, 2, 95, 15, 3, 2,
	2, 2, 96, 108, 7, 15, 2, 2, 97, 109, 5, 38, 20, 2, 98, 104, 7, 26, 2, 2,
	99, 100, 5, 38, 20, 2, 100, 101, 5, 44, 23, 2, 101, 103, 3, 2, 2, 2, 102,
	99, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3,
	2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 109, 7, 27, 2,
	2, 108, 97, 3, 2, 2, 2, 108, 98, 3, 2, 2, 2, 109, 17, 3, 2, 2, 2, 110,
	117, 5, 40, 21, 2, 111, 117, 5, 20, 11, 2, 112, 113, 7, 26, 2, 2, 113,
	114, 5, 18, 10, 2, 114, 115, 7, 27, 2, 2, 115, 117, 3, 2, 2, 2, 116, 110,
	3, 2, 2, 2, 116, 111, 3, 2, 2, 2, 116, 112, 3, 2, 2, 2, 117, 19, 3, 2,
	2, 2, 118, 121, 5, 22, 12, 2, 119, 121, 5, 26, 14, 2, 120, 118, 3, 2, 2,
	2, 120, 119, 3, 2, 2, 2, 121, 21, 3, 2, 2, 2, 122, 123, 7, 10, 2, 2, 123,
	129, 7, 4, 2, 2, 124, 125, 5, 24, 13, 2, 125, 126, 5, 44, 23, 2, 126, 128,
	3, 2, 2, 2, 127, 124, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2,
	2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2,
	132, 133, 7, 5, 2, 2, 133, 23, 3, 2, 2, 2, 134, 135, 6, 13, 2, 2, 135,
	136, 5, 36, 19, 2, 136, 137, 5, 18, 10, 2, 137, 139, 3, 2, 2, 2, 138, 140,
	5, 10, 6, 2, 139, 138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 25, 3, 2,
	2, 2, 141, 142, 7, 9, 2, 2, 142, 148, 7, 4, 2, 2, 143, 144, 5, 28, 15,
	2, 144, 145, 5, 44, 23, 2, 145, 147, 3, 2, 2, 2, 146, 143, 3, 2, 2, 2,
	147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149,
	151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151, 152, 7, 5, 2, 2, 152, 27, 3,
	2, 2, 2, 153, 154, 6, 15, 3, 2, 154, 155, 7, 13, 2, 2, 155, 156, 5, 32,
	17, 2, 156, 157, 5, 30, 16, 2, 157, 161, 3, 2, 2, 2, 158, 159, 7, 13, 2,
	2, 159, 161, 5, 32, 17, 2, 160, 153, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2,
	161, 29, 3, 2, 2, 2, 162, 165, 5, 32, 17, 2, 163, 165, 5, 18, 10, 2, 164,
	162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 165, 31, 3, 2, 2, 2, 166, 178, 7,
	26, 2, 2, 167, 172, 5, 34, 18, 2, 168, 169, 7, 25, 2, 2, 169, 171, 5, 34,
	18, 2, 170, 168, 3, 2, 2, 2, 171, 174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	172, 173, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 175,
	177, 7, 25, 2, 2, 176, 175, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179,
	3, 2, 2, 2, 178, 167, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 3, 2,
	2, 2, 180, 181, 7, 27, 2, 2, 181, 33, 3, 2, 2, 2, 182, 184, 5, 36, 19,
	2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 185,
	187, 7, 6, 2, 2, 186, 185, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 189, 5, 18, 10, 2, 189, 35, 3, 2, 2, 2, 190, 195, 7, 13,
	2, 2, 191, 192, 7, 25, 2, 2, 192, 194, 7, 13, 2, 2, 193, 191, 3, 2, 2,
	2, 194, 197, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	37, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 198, 199, 7, 13, 2, 2, 199, 200,
	5, 18, 10, 2, 200, 39, 3, 2, 2, 2, 201, 204, 7, 13, 2, 2, 202, 204, 5,
	42, 22, 2, 203, 201, 3, 2, 2, 2, 203, 202, 3, 2, 2, 2, 204, 41, 3, 2, 2,
	2, 205, 206, 7, 13, 2, 2, 206, 207, 7, 3, 2, 2, 207, 208, 7, 13, 2, 2,
	208, 43, 3, 2, 2, 2, 209, 214, 7, 7, 2, 2, 210, 214, 7, 2, 2, 3, 211, 214,
	6, 23, 4, 2, 212, 214, 6, 23, 5, 2, 213, 209, 3, 2, 2, 2, 213, 210, 3,
	2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214, 45, 3, 2, 2,
	2, 215, 219, 7, 7, 2, 2, 216, 219, 6, 24, 6, 2, 217, 219, 6, 24, 7, 2,
	218, 215, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 217, 3, 2, 2, 2, 219,
	47, 3, 2, 2, 2, 25, 55, 63, 79, 83, 86, 104, 108, 116, 120, 129, 139, 148,
	160, 164, 172, 176, 178, 183, 186, 195, 203, 213, 218,
}
var literalNames = []string{
	"", "'.'", "'{'", "'}'", "'...'", "';'", "'func'", "'interface'", "'struct'",
	"'package'", "'nil'", "", "'import'", "'type'", "'*'", "'&'", "", "", "",
	"", "", "", "", "','", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "FUNC", "INTERFACE", "STRUCT", "PACKAGE", "NIL_LIT",
	"IDENTIFIER", "IMPORT", "TYPE", "STAR", "AMPERSAND", "WS", "COMMENT", "TERMINATOR",
	"LINE_COMMENT", "INTERPRETED_STRING_LIT", "OCTAL_LIT", "HEX_LIT", "COMMA",
	"L_PAREN", "R_PAREN",
}

var ruleNames = []string{
	"sourceFile", "packageClause", "importDecl", "importSpec", "string_", "importPath",
	"declaration", "typeDecl", "type_", "typeLit", "structType", "fieldDecl",
	"interfaceType", "methodSpec", "result", "parameters", "parameterDecl",
	"identifierList", "typeSpec", "typeName", "qualifiedIdent", "eos", "eos2",
}

type ADLParser struct {
	*antlr.BaseParser
}

// NewADLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ADLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewADLParser(input antlr.TokenStream) *ADLParser {
	this := new(ADLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ADL.g4"

	return this
}

// ADLParser tokens.
const (
	ADLParserEOF                    = antlr.TokenEOF
	ADLParserT__0                   = 1
	ADLParserT__1                   = 2
	ADLParserT__2                   = 3
	ADLParserT__3                   = 4
	ADLParserT__4                   = 5
	ADLParserFUNC                   = 6
	ADLParserINTERFACE              = 7
	ADLParserSTRUCT                 = 8
	ADLParserPACKAGE                = 9
	ADLParserNIL_LIT                = 10
	ADLParserIDENTIFIER             = 11
	ADLParserIMPORT                 = 12
	ADLParserTYPE                   = 13
	ADLParserSTAR                   = 14
	ADLParserAMPERSAND              = 15
	ADLParserWS                     = 16
	ADLParserCOMMENT                = 17
	ADLParserTERMINATOR             = 18
	ADLParserLINE_COMMENT           = 19
	ADLParserINTERPRETED_STRING_LIT = 20
	ADLParserOCTAL_LIT              = 21
	ADLParserHEX_LIT                = 22
	ADLParserCOMMA                  = 23
	ADLParserL_PAREN                = 24
	ADLParserR_PAREN                = 25
)

// ADLParser rules.
const (
	ADLParserRULE_sourceFile     = 0
	ADLParserRULE_packageClause  = 1
	ADLParserRULE_importDecl     = 2
	ADLParserRULE_importSpec     = 3
	ADLParserRULE_string_        = 4
	ADLParserRULE_importPath     = 5
	ADLParserRULE_declaration    = 6
	ADLParserRULE_typeDecl       = 7
	ADLParserRULE_type_          = 8
	ADLParserRULE_typeLit        = 9
	ADLParserRULE_structType     = 10
	ADLParserRULE_fieldDecl      = 11
	ADLParserRULE_interfaceType  = 12
	ADLParserRULE_methodSpec     = 13
	ADLParserRULE_result         = 14
	ADLParserRULE_parameters     = 15
	ADLParserRULE_parameterDecl  = 16
	ADLParserRULE_identifierList = 17
	ADLParserRULE_typeSpec       = 18
	ADLParserRULE_typeName       = 19
	ADLParserRULE_qualifiedIdent = 20
	ADLParserRULE_eos            = 21
	ADLParserRULE_eos2           = 22
)

// ISourceFileContext is an interface to support dynamic dispatch.
type ISourceFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceFileContext differentiates from other interfaces.
	IsSourceFileContext()
}

type SourceFileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceFileContext() *SourceFileContext {
	var p = new(SourceFileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_sourceFile
	return p
}

func (*SourceFileContext) IsSourceFileContext() {}

func NewSourceFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceFileContext {
	var p = new(SourceFileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_sourceFile

	return p
}

func (s *SourceFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceFileContext) PackageClause() IPackageClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageClauseContext)
}

func (s *SourceFileContext) Eos2() IEos2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEos2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEos2Context)
}

func (s *SourceFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLParserEOF, 0)
}

func (s *SourceFileContext) AllImportDecl() []IImportDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclContext)(nil)).Elem())
	var tst = make([]IImportDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclContext)
		}
	}

	return tst
}

func (s *SourceFileContext) ImportDecl(i int) IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *SourceFileContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SourceFileContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *SourceFileContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *SourceFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterSourceFile(s)
	}
}

func (s *SourceFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitSourceFile(s)
	}
}

func (s *SourceFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitSourceFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) SourceFile() (localctx ISourceFileContext) {
	localctx = NewSourceFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ADLParserRULE_sourceFile)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.PackageClause()
	}
	{
		p.SetState(47)
		p.Eos2()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserIMPORT {
		{
			p.SetState(48)
			p.ImportDecl()
		}
		{
			p.SetState(49)
			p.Eos()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserTYPE {
		{
			p.SetState(56)
			p.Declaration()
		}

		{
			p.SetState(57)
			p.Eos()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
		p.Match(ADLParserEOF)
	}

	return localctx
}

// IPackageClauseContext is an interface to support dynamic dispatch.
type IPackageClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageClauseContext differentiates from other interfaces.
	IsPackageClauseContext()
}

type PackageClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageClauseContext() *PackageClauseContext {
	var p = new(PackageClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_packageClause
	return p
}

func (*PackageClauseContext) IsPackageClauseContext() {}

func NewPackageClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageClauseContext {
	var p = new(PackageClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_packageClause

	return p
}

func (s *PackageClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageClauseContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(ADLParserPACKAGE, 0)
}

func (s *PackageClauseContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, 0)
}

func (s *PackageClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterPackageClause(s)
	}
}

func (s *PackageClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitPackageClause(s)
	}
}

func (s *PackageClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitPackageClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) PackageClause() (localctx IPackageClauseContext) {
	localctx = NewPackageClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ADLParserRULE_packageClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(ADLParserPACKAGE)
	}
	{
		p.SetState(67)
		p.Match(ADLParserIDENTIFIER)
	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(ADLParserIMPORT, 0)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportSpecContext)(nil)).Elem())
	var tst = make([]IImportSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportSpecContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserL_PAREN, 0)
}

func (s *ImportDeclContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserR_PAREN, 0)
}

func (s *ImportDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) ImportDecl() (localctx IImportDeclContext) {
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ADLParserRULE_importDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(ADLParserIMPORT)
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserT__0, ADLParserIDENTIFIER, ADLParserINTERPRETED_STRING_LIT:
		{
			p.SetState(70)
			p.ImportSpec()
		}

	case ADLParserL_PAREN:
		{
			p.SetState(71)
			p.Match(ADLParserL_PAREN)
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserT__0)|(1<<ADLParserIDENTIFIER)|(1<<ADLParserINTERPRETED_STRING_LIT))) != 0 {
			{
				p.SetState(72)
				p.ImportSpec()
			}
			{
				p.SetState(73)
				p.Eos()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)
			p.Match(ADLParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) ImportPath() IImportPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportPathContext)
}

func (s *ImportSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ADLParserRULE_importSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserT__0 || _la == ADLParserIDENTIFIER {
		{
			p.SetState(83)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ADLParserT__0 || _la == ADLParserIDENTIFIER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(86)
		p.ImportPath()
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ADLParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitString_(s)
	}
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ADLParserRULE_string_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ADLParserINTERPRETED_STRING_LIT)
	}

	return localctx
}

// IImportPathContext is an interface to support dynamic dispatch.
type IImportPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportPathContext differentiates from other interfaces.
	IsImportPathContext()
}

type ImportPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportPathContext() *ImportPathContext {
	var p = new(ImportPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_importPath
	return p
}

func (*ImportPathContext) IsImportPathContext() {}

func NewImportPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportPathContext {
	var p = new(ImportPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_importPath

	return p
}

func (s *ImportPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportPathContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ImportPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterImportPath(s)
	}
}

func (s *ImportPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitImportPath(s)
	}
}

func (s *ImportPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitImportPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) ImportPath() (localctx IImportPathContext) {
	localctx = NewImportPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ADLParserRULE_importPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.String_()
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ADLParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.TypeDecl()
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ADLParserTYPE, 0)
}

func (s *TypeDeclContext) AllTypeSpec() []ITypeSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem())
	var tst = make([]ITypeSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecContext)
		}
	}

	return tst
}

func (s *TypeDeclContext) TypeSpec(i int) ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeDeclContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserL_PAREN, 0)
}

func (s *TypeDeclContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserR_PAREN, 0)
}

func (s *TypeDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *TypeDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ADLParserRULE_typeDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(ADLParserTYPE)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserIDENTIFIER:
		{
			p.SetState(95)
			p.TypeSpec()
		}

	case ADLParserL_PAREN:
		{
			p.SetState(96)
			p.Match(ADLParserL_PAREN)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ADLParserIDENTIFIER {
			{
				p.SetState(97)
				p.TypeSpec()
			}
			{
				p.SetState(98)
				p.Eos()
			}

			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(ADLParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *Type_Context) TypeLit() ITypeLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *Type_Context) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserL_PAREN, 0)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserR_PAREN, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitType_(s)
	}
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ADLParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.TypeName()
		}

	case ADLParserINTERFACE, ADLParserSTRUCT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.TypeLit()
		}

	case ADLParserL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(ADLParserL_PAREN)
		}
		{
			p.SetState(111)
			p.Type_()
		}
		{
			p.SetState(112)
			p.Match(ADLParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeLit
	return p
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) StructType() IStructTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *TypeLitContext) InterfaceType() IInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeContext)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterTypeLit(s)
	}
}

func (s *TypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitTypeLit(s)
	}
}

func (s *TypeLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitTypeLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ADLParserRULE_typeLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ADLParserSTRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.StructType()
		}

	case ADLParserINTERFACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.InterfaceType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStructTypeContext is an interface to support dynamic dispatch.
type IStructTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTypeContext differentiates from other interfaces.
	IsStructTypeContext()
}

type StructTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeContext() *StructTypeContext {
	var p = new(StructTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_structType
	return p
}

func (*StructTypeContext) IsStructTypeContext() {}

func NewStructTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeContext {
	var p = new(StructTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_structType

	return p
}

func (s *StructTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(ADLParserSTRUCT, 0)
}

func (s *StructTypeContext) AllFieldDecl() []IFieldDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem())
	var tst = make([]IFieldDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDeclContext)
		}
	}

	return tst
}

func (s *StructTypeContext) FieldDecl(i int) IFieldDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclContext)
}

func (s *StructTypeContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *StructTypeContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterStructType(s)
	}
}

func (s *StructTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitStructType(s)
	}
}

func (s *StructTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitStructType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) StructType() (localctx IStructTypeContext) {
	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ADLParserRULE_structType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(ADLParserSTRUCT)
	}
	{
		p.SetState(121)
		p.Match(ADLParserT__1)
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.FieldDecl()
			}
			{
				p.SetState(123)
				p.Eos()
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	{
		p.SetState(130)
		p.Match(ADLParserT__2)
	}

	return localctx
}

// IFieldDeclContext is an interface to support dynamic dispatch.
type IFieldDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclContext differentiates from other interfaces.
	IsFieldDeclContext()
}

type FieldDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclContext() *FieldDeclContext {
	var p = new(FieldDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_fieldDecl
	return p
}

func (*FieldDeclContext) IsFieldDeclContext() {}

func NewFieldDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclContext {
	var p = new(FieldDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_fieldDecl

	return p
}

func (s *FieldDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *FieldDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDeclContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *FieldDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterFieldDecl(s)
	}
}

func (s *FieldDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitFieldDecl(s)
	}
}

func (s *FieldDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitFieldDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) FieldDecl() (localctx IFieldDeclContext) {
	localctx = NewFieldDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ADLParserRULE_fieldDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(132)

	if !(p.noTerminatorBetween(2)) {
		panic(antlr.NewFailedPredicateException(p, "p.noTerminatorBetween(2)", ""))
	}
	{
		p.SetState(133)
		p.IdentifierList()
	}
	{
		p.SetState(134)
		p.Type_()
	}

	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(136)
			p.String_()
		}

	}

	return localctx
}

// IInterfaceTypeContext is an interface to support dynamic dispatch.
type IInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeContext differentiates from other interfaces.
	IsInterfaceTypeContext()
}

type InterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeContext() *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_interfaceType
	return p
}

func (*InterfaceTypeContext) IsInterfaceTypeContext() {}

func NewInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeContext {
	var p = new(InterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_interfaceType

	return p
}

func (s *InterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ADLParserINTERFACE, 0)
}

func (s *InterfaceTypeContext) AllMethodSpec() []IMethodSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem())
	var tst = make([]IMethodSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodSpecContext)
		}
	}

	return tst
}

func (s *InterfaceTypeContext) MethodSpec(i int) IMethodSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodSpecContext)
}

func (s *InterfaceTypeContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *InterfaceTypeContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *InterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterInterfaceType(s)
	}
}

func (s *InterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitInterfaceType(s)
	}
}

func (s *InterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) InterfaceType() (localctx IInterfaceTypeContext) {
	localctx = NewInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ADLParserRULE_interfaceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(ADLParserINTERFACE)
	}
	{
		p.SetState(140)
		p.Match(ADLParserT__1)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(141)
				p.MethodSpec()
			}
			{
				p.SetState(142)
				p.Eos()
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	{
		p.SetState(149)
		p.Match(ADLParserT__2)
	}

	return localctx
}

// IMethodSpecContext is an interface to support dynamic dispatch.
type IMethodSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodSpecContext differentiates from other interfaces.
	IsMethodSpecContext()
}

type MethodSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodSpecContext() *MethodSpecContext {
	var p = new(MethodSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_methodSpec
	return p
}

func (*MethodSpecContext) IsMethodSpecContext() {}

func NewMethodSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodSpecContext {
	var p = new(MethodSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_methodSpec

	return p
}

func (s *MethodSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, 0)
}

func (s *MethodSpecContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *MethodSpecContext) Result() IResultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IResultContext)
}

func (s *MethodSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterMethodSpec(s)
	}
}

func (s *MethodSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitMethodSpec(s)
	}
}

func (s *MethodSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitMethodSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) MethodSpec() (localctx IMethodSpecContext) {
	localctx = NewMethodSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ADLParserRULE_methodSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(151)

		if !(p.noTerminatorAfterParams(2)) {
			panic(antlr.NewFailedPredicateException(p, "p.noTerminatorAfterParams(2)", ""))
		}
		{
			p.SetState(152)
			p.Match(ADLParserIDENTIFIER)
		}
		{
			p.SetState(153)
			p.Parameters()
		}
		{
			p.SetState(154)
			p.Result()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(ADLParserIDENTIFIER)
		}
		{
			p.SetState(157)
			p.Parameters()
		}

	}

	return localctx
}

// IResultContext is an interface to support dynamic dispatch.
type IResultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResultContext differentiates from other interfaces.
	IsResultContext()
}

type ResultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultContext() *ResultContext {
	var p = new(ResultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_result
	return p
}

func (*ResultContext) IsResultContext() {}

func NewResultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultContext {
	var p = new(ResultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_result

	return p
}

func (s *ResultContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *ResultContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ResultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterResult(s)
	}
}

func (s *ResultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitResult(s)
	}
}

func (s *ResultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitResult(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Result() (localctx IResultContext) {
	localctx = NewResultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ADLParserRULE_result)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Parameters()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Type_()
		}

	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserL_PAREN, 0)
}

func (s *ParametersContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(ADLParserR_PAREN, 0)
}

func (s *ParametersContext) AllParameterDecl() []IParameterDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem())
	var tst = make([]IParameterDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterDeclContext)
		}
	}

	return tst
}

func (s *ParametersContext) ParameterDecl(i int) IParameterDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ADLParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(ADLParserL_PAREN)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ADLParserT__3)|(1<<ADLParserINTERFACE)|(1<<ADLParserSTRUCT)|(1<<ADLParserIDENTIFIER)|(1<<ADLParserL_PAREN))) != 0 {
		{
			p.SetState(165)
			p.ParameterDecl()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(166)
					p.Match(ADLParserCOMMA)
				}
				{
					p.SetState(167)
					p.ParameterDecl()
				}

			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ADLParserCOMMA {
			{
				p.SetState(173)
				p.Match(ADLParserCOMMA)
			}

		}

	}
	{
		p.SetState(178)
		p.Match(ADLParserR_PAREN)
	}

	return localctx
}

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_parameterDecl
	return p
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ParameterDeclContext) IdentifierList() IIdentifierListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterParameterDecl(s)
	}
}

func (s *ParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitParameterDecl(s)
	}
}

func (s *ParameterDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitParameterDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) ParameterDecl() (localctx IParameterDeclContext) {
	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ADLParserRULE_parameterDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(180)
			p.IdentifierList()
		}

	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ADLParserT__3 {
		{
			p.SetState(183)
			p.Match(ADLParserT__3)
		}

	}
	{
		p.SetState(186)
		p.Type_()
	}

	return localctx
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_identifierList
	return p
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ADLParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ADLParserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ADLParserRULE_identifierList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ADLParserIDENTIFIER)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ADLParserCOMMA {
		{
			p.SetState(189)
			p.Match(ADLParserCOMMA)
		}
		{
			p.SetState(190)
			p.Match(ADLParserIDENTIFIER)
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, 0)
}

func (s *TypeSpecContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ADLParserRULE_typeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ADLParserIDENTIFIER)
	}
	{
		p.SetState(197)
		p.Type_()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, 0)
}

func (s *TypeNameContext) QualifiedIdent() IQualifiedIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ADLParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(ADLParserIDENTIFIER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.QualifiedIdent()
		}

	}

	return localctx
}

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_qualifiedIdent
	return p
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ADLParserIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ADLParserIDENTIFIER, i)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ADLParserRULE_qualifiedIdent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(ADLParserIDENTIFIER)
	}
	{
		p.SetState(204)
		p.Match(ADLParserT__0)
	}
	{
		p.SetState(205)
		p.Match(ADLParserIDENTIFIER)
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(ADLParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ADLParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(ADLParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(ADLParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(209)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(210)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "p.checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

// IEos2Context is an interface to support dynamic dispatch.
type IEos2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEos2Context differentiates from other interfaces.
	IsEos2Context()
}

type Eos2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEos2Context() *Eos2Context {
	var p = new(Eos2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ADLParserRULE_eos2
	return p
}

func (*Eos2Context) IsEos2Context() {}

func NewEos2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Eos2Context {
	var p = new(Eos2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ADLParserRULE_eos2

	return p
}

func (s *Eos2Context) GetParser() antlr.Parser { return s.parser }
func (s *Eos2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Eos2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Eos2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.EnterEos2(s)
	}
}

func (s *Eos2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ADLListener); ok {
		listenerT.ExitEos2(s)
	}
}

func (s *Eos2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ADLVisitor:
		return t.VisitEos2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ADLParser) Eos2() (localctx IEos2Context) {
	localctx = NewEos2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ADLParserRULE_eos2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.Match(ADLParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(214)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(215)

		if !(p.checkPreviousTokenText("}")) {
			panic(antlr.NewFailedPredicateException(p, "p.checkPreviousTokenText(\"}\")", ""))
		}

	}

	return localctx
}

func (p *ADLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *FieldDeclContext = nil
		if localctx != nil {
			t = localctx.(*FieldDeclContext)
		}
		return p.FieldDecl_Sempred(t, predIndex)

	case 13:
		var t *MethodSpecContext = nil
		if localctx != nil {
			t = localctx.(*MethodSpecContext)
		}
		return p.MethodSpec_Sempred(t, predIndex)

	case 21:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	case 22:
		var t *Eos2Context = nil
		if localctx != nil {
			t = localctx.(*Eos2Context)
		}
		return p.Eos2_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ADLParser) FieldDecl_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.noTerminatorBetween(2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ADLParser) MethodSpec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.noTerminatorAfterParams(2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ADLParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.lineTerminatorAhead()

	case 3:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ADLParser) Eos2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.lineTerminatorAhead()

	case 5:
		return p.checkPreviousTokenText("}")

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
