package golang

import (
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/sql"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"strconv"
	"strings"
)

func implementExecOne(fun *ast.Func, sql token.String, mapping sql.ExecOne) error {
	in, err := assemblePreparedStatementPlaceholders(mapping.In)
	if err != nil {
		return err
	}

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					if _, err := r.db.ExecContext(r.context(), q, {{.Get "in"}}); err!=nil {
						return {{.Use "fmt.Errorf"}}("cannot execute '%s': %w", q, err)	
					}
			
					return nil
				`).
				Put("query", strconv.Quote(sql.String())).
				Put("in", in),
		),
	)

	return nil
}

func implementExecMany(fun *ast.Func, sql token.String, mapping sql.ExecMany) error {
	in, err := assemblePreparedStatementPlaceholders(mapping.In)
	if err != nil {
		return err
	}

	slice := mapping.Slice.String()

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					c := r.context()
					x, err := r.db.Begin()
					if err != nil{
						return {{.Use "fmt.Errorf"}}("cannot begin transaction '%s': %w", q, err)	
					}
					defer x.Rollback()

					s, err := x.PrepareContext(c, q)
					if err != nil{
						return {{.Use "fmt.Errorf"}}("cannot prepare transaction '%s': %w", q, err)	
					}
					defer s.Close()

					for i := range {{.Get "slice"}}{
						if _, err := s.ExecContext(c, {{.Get "in"}}); err!=nil{
							return {{.Use "fmt.Errorf"}}("cannot execute '%s': %w", q, err)	
						}
					}

					if err := x.Commit(); err != nil{
						return {{.Use "fmt.Errorf"}}("cannot commit transaction '%s': %w", q, err)	
					}

					return nil
				`).
				Put("query", strconv.Quote(sql.String())).
				Put("in", in).
				Put("slice", slice),
		),
	)
	return nil
}

// this thing can only generate properly for a single primitive or a single struct for multiple return.
// Multiple primitives are not supported. Also the return must not be a generic in any way.
func implementFindOne(fun *ast.Func, sql token.String, mapping sql.QueryOne) error {
	in, err := assemblePreparedStatementPlaceholders(mapping.In)
	if err != nil {
		return err
	}

	out, err := assembleScan(mapping.Out, "&i")
	if err != nil {
		return err
	}

	// this another hard assumption
	simpleDecl := fun.FunResults[0].ParamTypeDecl.(*ast.SimpleTypeDecl)

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					var i {{.Use (.Get "returnType")}}
					w, err := r.db.QueryContext(r.context(), q, {{.Get "in"}})
					if err!=nil {
						return i, {{.Use "fmt.Errorf"}}("cannot query '%s': %w", q, err)	
					}
			
					defer w.Close()
					for r.Next() {
						if err:= w.Scan({{.Get "out"}}); err!=nil {
							return i, {{.Use "fmt.Errorf"}}("scan of '%s' failed: %w",q, err)
						}
					}

					if err := rows.Err(); err!=nil{
						return i, {{.Use "fmt.Errorf"}}("query of '%s' failed: %w",q, err)
					}
			
					return i, nil
				`).
				Put("query", strconv.Quote(sql.String())).
				Put("in", in).
				Put("out", out).
				Put("returnType", string(simpleDecl.SimpleName)),
		),
	)

	return nil
}

func implementFindMany(fun *ast.Func, sql token.String, mapping sql.QueryMany) error {
	in, err := assemblePreparedStatementPlaceholders(mapping.In)
	if err != nil {
		return err
	}

	out, err := assembleScan(mapping.Out, "&t")
	if err != nil {
		return err
	}

	// this another hard assumption
	slice, ok := fun.FunResults[0].ParamTypeDecl.(*ast.SliceTypeDecl)
	if !ok {
		return token.NewPosError(astutil.WrapNode(fun.FunResults[0]), fun.FunName+" result is a '"+fun.FunResults[0].String()+"' but expected a slice")
	}

	simpleDecl := slice.TypeDecl.(*ast.SimpleTypeDecl)

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					var i []{{.Use (.Get "returnType")}}
					w, err := r.db.QueryContext(r.context(), q, {{.Get "in"}})
					if err!=nil {
						return i, {{.Use "fmt.Errorf"}}("cannot query '%s': %w", q, err)	
					}
			
					defer w.Close()
					for r.Next() {
						var t {{.Use (.Get "returnType")}}
						if err:= w.Scan({{.Get "out"}}); err!=nil {
							return i, {{.Use "fmt.Errorf"}}("scan of '%s' failed: %w",q, err)
						}

						i = append(i, t)
					}

					if err := rows.Err(); err!=nil{
						return i, {{.Use "fmt.Errorf"}}("query of '%s' failed: %w",q, err)
					}
			
					return i, nil
				`).
				Put("query", strconv.Quote(sql.String())).
				Put("in", in).
				Put("out", out).
				Put("returnType", string(simpleDecl.SimpleName)),
		),
	)

	return nil
}

// assemblePreparedStatementPlaceholders expects lits to be either things like "myparam" or "myparam.field".
func assemblePreparedStatementPlaceholders(lits []token.String) (string, error) {
	placeholderList := ""
	for i, lit := range lits {
		if lit.String() == "." || lit.String() == "" {
			return "", token.NewPosError(lit, "invalid notation for in-parameter")
		}

		placeholderList += lit.String()
		if i < len(lits)-1 {
			placeholderList += ", "
		}
	}

	return placeholderList, nil
}

// assembleScan expects lits to be one of the following:
//  * a primitive: So only a single lit is valid and it must be "."
//  * a struct: so only ".field" declarations are valid.
// prefix is used to concat the variable name before.
func assembleScan(lits []token.String, prefix string) (string, error) {
	if len(lits) == 1 && lits[0].String() == "." {
		return prefix, nil
	}

	placeholderList := ""
	for i, lit := range lits {
		if (lit.String() == "." && len(lits) > 1) || lit.String() == "" {
			return "", token.NewPosError(lit, "invalid notation for scan-parameter. Only a single '.' is allowed.")
		}

		if !strings.HasPrefix(lit.String(), ".") {
			return "", token.NewPosError(lit, "invalid notation for scan-parameter. Every field must start with a '.'")
		}

		placeholderList += prefix + lit.String()
		if i < len(lits)-1 {
			placeholderList += ", "
		}
	}

	return placeholderList, nil
}
