package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io"
	"reflect"
)

type TreeShapeListener struct {
	*BaseADLListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(reflect.TypeOf(ctx).String() + " ->" + ctx.GetText())
}

type errorAdapter struct {
	err error
	antlr.DefaultErrorListener
}

func (err *errorAdapter) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err.err = fmt.Errorf("syntax error: %v -> %s: %d:%d", offendingSymbol, msg, line, column)
}

func ParseText(in string) error {
	lexer := NewADLLexer(antlr.NewInputStream(in))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewADLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	err := &errorAdapter{}
	p.AddErrorListener(err)
	tree := p.SourceFile()


	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	return err.err
}

func Parse(in io.Reader) error {
	buf, err := io.ReadAll(in)
	if err != nil {
		return err
	}

	return ParseText(string(buf))

}
