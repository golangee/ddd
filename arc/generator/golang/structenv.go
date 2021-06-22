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

// AddParseEnvFunc appends a method named "ParseEnv" which has the given struct as a pointer
// receiver and sets all defined environment variables to the environment one.
func AddParseEnvFunc(fieldPrefix string, node *ast.Struct) (*ast.Func, error) {
	fun := ast.NewFunc("ParseEnv").
		SetPtrReceiver(true).
		SetRecName(node.DefaultRecName).
		AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error)))

	comment := "... tries to parse the environment variables into this instance.\nIt will only set those values, which have been actually defined.\nIf values cannot be parsed, an error is returned.\n"
	node.AddMethods(fun)
	body := ast.NewBlock()
	fun.SetBody(body)

	envNamePrefix := strings.ReplaceAll(fieldPrefix, "/", "_")

	for _, field := range node.Fields() {
		flagName := strings.ToLower(envNamePrefix + "_" + field.FieldName)
		comment += " * " + field.FieldName + " is parsed from variable '" + flagName + "' if it has been set.\n"

		stereotype.FieldFrom(field).SetEnvironmentVariable(flagName)

		var parseBody []ast.Node

		switch t := field.FieldType.(type) {
		case *ast.SimpleTypeDecl:
			switch t.SimpleName {
			case stdlib.Bool:
				parseBody = append(parseBody,
					lang.TryDefine(ast.NewIdent("parsed"), lang.CallStatic("strconv.ParseBool", ast.NewIdent("value")), "unable to parse flag '"+flagName+"'"),
					ast.NewAssign(ast.Exprs(lang.Attr(field.FieldName)), ast.AssignSimple, ast.Exprs(ast.NewIdent("parsed"))),
				)
			case stdlib.String:
				parseBody = append(parseBody, ast.NewAssign(ast.Exprs(lang.Attr(field.FieldName)), ast.AssignSimple, ast.Exprs(ast.NewIdent("value"))))
			case stdlib.Int64:
				parseBody = tryParseInt("int64", flagName, field.FieldName)
			case stdlib.Int32:
				parseBody = tryParseInt("int32", flagName, field.FieldName)
			case stdlib.Int16:
				parseBody = tryParseInt("int16", flagName, field.FieldName)
			case stdlib.Int:
				parseBody = tryParseInt("int", flagName, field.FieldName)
			case stdlib.Float32:
				parseBody = tryParseInt("float32", flagName, field.FieldName)
			case stdlib.Float64:
				parseBody = tryParseInt("float64", flagName, field.FieldName)
			case stdlib.Duration:
				parseBody = append(parseBody,
					lang.TryDefine(ast.NewIdent("parsed"), lang.CallStatic("time.ParseDuration", ast.NewIdent("value")), "unable to parse flag '"+flagName+"'"),
					ast.NewAssign(ast.Exprs(lang.Attr(field.FieldName)), ast.AssignSimple, ast.Exprs(ast.NewIdent("parsed"))),
				)
			default:
				return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct parse env function")
			}

		default:
			return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct parse env function")
		}

		body.Add(
			tryParseEnv(flagName, parseBody...),
		)

	}

	body.Add(
		lang.Term(),
		ast.NewReturnStmt(ast.NewIdentLit("nil")),
	)
	fun.SetComment(comment)
	return fun, nil
}

func tryParseInt(intType ast.Name, flagName, fieldName string) []ast.Node {
	dst := ast.NewIdent("value")
	var conv *ast.Ident
	var call ast.Expr
	switch intType {
	case "int":
		conv = ast.NewIdent("int")
		call = lang.CallStatic("strconv.Atoi", dst)
	case "int32":
		conv = ast.NewIdent("int32")
		call = lang.CallStatic("strconv.Atoi", dst)
	case "int16":
		conv = ast.NewIdent("int16")
		call = lang.CallStatic("strconv.Atoi", dst)
	case "int64":
		conv = ast.NewIdent("int64")
		call = lang.CallStatic("strconv.ParseInt", dst, ast.NewIntLit(10), ast.NewIntLit(64))
	case "float32":
		conv = ast.NewIdent("float32")
		call = lang.CallStatic("strconv.ParseFloat", dst, ast.NewIntLit(32))
	case "float64":
		conv = ast.NewIdent("float64")
		call = lang.CallStatic("strconv.ParseFloat", dst, ast.NewIntLit(64))
	default:
		panic("not implemented " + intType)
	}
	return []ast.Node{
		lang.TryDefine(ast.NewIdent("parsed"), call, "unable to parse flag '"+flagName+"'"),
		ast.NewAssign(ast.Exprs(lang.Attr(fieldName)), ast.AssignSimple, ast.Exprs(ast.NewCallExpr(conv, ast.NewIdent("parsed")))),
	}
}

func tryParseEnv(flag string, body ...ast.Node) ast.Node {
	return ast.NewIfStmt(ast.NewIdent("ok"), ast.NewBlock(body...)).
		SetInit(ast.NewAssign(ast.Exprs(ast.NewIdent("value"), ast.NewIdent("ok")), ast.AssignDefine, ast.Exprs(lang.CallStatic("os.LookupEnv", ast.NewStrLit(flag)))))
}
