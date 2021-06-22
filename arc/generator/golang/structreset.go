package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"strconv"
	"strings"
	"time"
)

// SimulateDefaultJson inspects the fields default values and estimates how a potential json serialization may look like.
// Better to call AddResetFunc before, to ensure more correctly defined default values.
func SimulateDefaultJson(node *ast.Struct) (interface{}, error) {
	obj := map[string]interface{}{}
	for _, field := range node.Fields() {
		var rawLiteral string
		if field.FieldDefault != nil {
			rawLiteral = field.FieldDefault.Val
		}
		switch t := field.FieldType.(type) {

		case *ast.SimpleTypeDecl:
			switch t.SimpleName {
			case stdlib.Bool:
				v, err := strconv.ParseBool(rawLiteral)
				if err != nil {
					return "", fmt.Errorf("default value of '%s' not a bool: %w", field.FieldName, err)
				}

				obj[field.FieldName] = v
			case stdlib.String:
				v, err := strconv.Unquote(rawLiteral)
				if err != nil {
					return "", fmt.Errorf("default value of '%s' not a quoted string: %w", field.FieldName, err)
				}

				obj[field.FieldName] = v
			case stdlib.Int64:
				fallthrough
			case stdlib.Int32:
				fallthrough
			case stdlib.Int16:
				fallthrough
			case stdlib.Float64:
				fallthrough
			case stdlib.Float32:
				fallthrough
			case stdlib.Int:
				v, err := strconv.ParseFloat(rawLiteral, 64)
				if err != nil {
					return "", fmt.Errorf("default value of '%s' not a number: %w", field.FieldName, err)
				}

				obj[field.FieldName] = v
			case stdlib.Duration:
				v, err := time.ParseDuration(rawLiteral)
				if err != nil {
					return "", fmt.Errorf("default value of '%s' not a duration: %w", field.FieldName, err)
				}

				obj[field.FieldName] = v
			default:
				// try to resolve nested structs
				node := astutil.Resolve(node, t.String())
				if node == nil {
					return "", fmt.Errorf("default value of '%s' not resolvable: '%s'", field.FieldName, t.String())
				}

				if s, ok := node.(*ast.Struct); ok {
					v, err := SimulateDefaultJson(s)
					if err != nil {
						return "", fmt.Errorf("default value of nested struct '%s' not simulatable: %w", field.FieldName, err)
					}

					obj[field.FieldName] = v
				} else {
					return "", fmt.Errorf("default value of '%s' not supported: '%s'", field.FieldName, t.String())
				}
			}
		default:
			// don't know what to do, but maps and slices may be possible but yet not supported as a configuration.
			obj[field.FieldName] = nil
		}

	}

	return obj, nil
}

// AddResetFunc appends a method named "Reset" which has the given struct as a pointer receiver and sets all
// literals back to default.
func AddResetFunc(node *ast.Struct) (*ast.Func, error) {
	fun := ast.NewFunc("Reset").
		SetPtrReceiver(true).
		SetRecName(node.DefaultRecName)

	comment := "...restores this instance to the default state.\n"
	node.AddMethods(fun)
	body := ast.NewBlock()
	fun.SetBody(body)

	for _, field := range node.Fields() {
		var rawLiteral string
		if field.FieldDefault != nil {
			rawLiteral = field.FieldDefault.Val
		}

		switch t := field.FieldType.(type) {
		case *ast.SimpleTypeDecl:
			switch t.SimpleName {
			case stdlib.Bool:
				if rawLiteral == "" {
					rawLiteral = "false"
					field.FieldDefault = ast.NewBasicLit(ast.TokenIdent, rawLiteral)
				}

				body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(fun.RecName()), ast.NewIdent(field.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewBasicLit(ast.TokenIdent, rawLiteral))))
				body.Add(ast.NewSym(ast.SymNewline))
			case stdlib.String:
				if rawLiteral == "" {
					rawLiteral = strconv.Quote("")
					field.FieldDefault = ast.NewBasicLit(ast.TokenString, rawLiteral)
				}

				body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(fun.RecName()), ast.NewIdent(field.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewBasicLit(ast.TokenString, rawLiteral))))
				body.Add(ast.NewSym(ast.SymNewline))
			case stdlib.Int64:
				fallthrough
			case stdlib.Int32:
				fallthrough
			case stdlib.Int16:
				fallthrough
			case stdlib.Float64:
				fallthrough
			case stdlib.Float32:
				fallthrough
			case stdlib.Int:
				if rawLiteral == "" {
					rawLiteral = "0"
					field.FieldDefault = ast.NewBasicLit(ast.TokenInt, rawLiteral)
				}

				body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(fun.RecName()), ast.NewIdent(field.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewBasicLit(ast.TokenInt, rawLiteral))))
				body.Add(ast.NewSym(ast.SymNewline))

			case stdlib.Duration:
				if rawLiteral == "" {
					rawLiteral = "0"
					field.FieldDefault = ast.NewBasicLit(ast.TokenInt, rawLiteral)
					body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(fun.RecName()), ast.NewIdent(field.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewBasicLit(ast.TokenInt, rawLiteral))))
				} else {
					d, err := time.ParseDuration(rawLiteral)
					if err != nil {
						return fun, token.NewPosError(astutil.WrapNode(field), "invalid default duration literal").SetCause(err)
					}

					body.Add(ast.NewAssign(ast.Exprs(ast.NewSelExpr(ast.NewIdent(fun.RecName()), ast.NewIdent(field.FieldName))), ast.AssignSimple, ast.Exprs(ast.NewCallExpr(ast.NewSelExpr(ast.NewQualIdent("time"), ast.NewIdent("Duration")), ast.NewInt64Lit(int64(d))))))
				}

				body.Add(ast.NewSym(ast.SymNewline))
			default:
				return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct reset function")
			}

		default:
			return fun, token.NewPosError(astutil.WrapNode(field), field.FieldName+" "+field.FieldType.String()+": unsupported field type for struct reset function")
		}

		commentLit := rawLiteral
		if strings.HasPrefix(commentLit, `"`) && strings.HasSuffix(commentLit, `"`) {
			commentLit = commentLit[1 : len(commentLit)-1]
		}

		comment += " * The default value of " + field.FieldName + " is '" + commentLit + "'\n"

	}

	fun.SetComment(comment)

	return fun, nil
}
