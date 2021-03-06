package golang

import (
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"strings"
)

// AddParseFlagFunc appends a method named "ConfigureFlags" which has the given struct as a pointer
// receiver and sets all defined struct variables to the flag ones. The go flags package to be
// parsed by the according struct instance.
// The naming is <a>-<b>-<c> for the flags. On unix, camel case is discouraged, so we have only the alternatives
// of . _ or - and we decided for now, to use -.
func AddParseFlagFunc(fieldPrefix string, node *ast.Struct) (*ast.Func, error) {
	fun := ast.NewFunc("ConfigureFlags").
		SetPtrReceiver(true).
		AddParams(
			ast.NewParam("flags", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl("flag.FlagSet"))),
		).
		SetRecName(node.DefaultRecName)

	node.AddMethods(fun)
	body := ast.NewBlock()
	fun.SetBody(body)

	const sep = "-"
	comment := "... configures the flags to be ready to get evaluated. The default values are taken from the struct at calling time.\nAfter calling, use flags.Parse() to load the values.\nThe default values are the field values at calling time.\n" +
		"Example:\n  cfg := " + node.TypeName + "{}\n  cfg.Reset()\n  flags := flag.NewFlagSet(`my app`, flag.ExitOnError)\n  cfg." + fun.FunName + "(flags)\n  flags.Parse(os.Args[1:])\n\n" +
		"The following flags will be tied to this instance:\n"
	envNamePrefix := strings.ReplaceAll(fieldPrefix, "/", sep)

	numFlags := 0
	for _, field := range node.Fields() {
		if !stereotype.FieldFrom(field).IsProgramFlag() {
			continue
		}

		numFlags++

		flagName := strings.ToLower(envNamePrefix + sep + field.FieldName)
		comment += " * " + field.FieldName + " is parsed from flag '" + flagName + "' if it has been set.\n"

		stereotype.FieldFrom(field).SetProgramFlagVariable(flagName)

		var parseBody ast.Node

		switch t := field.FieldType.(type) {
		case *ast.SimpleTypeDecl:
			dst := ast.NewUnaryExpr(ast.NewSelExpr(ast.NewIdent(fun.FunReceiverName), ast.NewIdent(field.FieldName)), ast.OpAnd)
			def := ast.NewSelExpr(ast.NewIdent(fun.FunReceiverName), ast.NewIdent(field.FieldName))
			comment := ""
			if field.Comment() != nil {
				comment = field.Comment().Text
			}

			switch t.SimpleName {
			case stdlib.Bool:
				parseBody = lang.CallIdent("flags", "BoolVar", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))
			case stdlib.Int:
				parseBody = lang.CallIdent("flags", "IntVar", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))
			case stdlib.String:
				parseBody = lang.CallIdent("flags", "StringVar", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))
			case stdlib.Int64:
				parseBody = lang.CallIdent("flags", "Int64Var", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))
			case stdlib.Duration:
				parseBody = lang.CallIdent("flags", "DurationVar", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))
			case stdlib.Float64:
				parseBody = lang.CallIdent("flags", "Float64Var", dst, ast.NewStrLit(flagName), def, ast.NewStrLit(comment))

			default:
				return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct parse flag function")
			}
		default:
			return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct parse flag function")
		}

		body.Add(
			parseBody, lang.Term(),
		)
	}

	if numFlags == 0{
		comment+="  - there are no flags available."
	}
	fun.SetComment(comment)
	return fun, nil
}
