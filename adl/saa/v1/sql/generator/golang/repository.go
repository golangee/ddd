package golang

import (
	"fmt"
	"github.com/golangee/architecture/adl/saa/v1/astutil"
	"github.com/golangee/architecture/adl/saa/v1/core"
	"github.com/golangee/architecture/adl/saa/v1/core/generator/corego"
	"github.com/golangee/architecture/adl/saa/v1/sql"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"reflect"
	"strconv"
	"strings"
)

func renderRepositories(dst *ast.Prj, src *sql.Ctx) error {
	modName := src.Mod.String()
	pkgName := src.Pkg.String()

	for _, repository := range src.Repositories {
		repoTypeName := repository.Implements
		file := corego.MkFile(dst, modName, pkgName, "tmp.go")
		rNode := astutil.Resolve(file, repoTypeName.String())
		repo, ok := rNode.(*ast.Interface)
		if !ok {
			return fmt.Errorf("unable to resolve required interface type " + repoTypeName.String())
		}

		file.Name = simpleLowerCaseName(repo.TypeName) + ".go"

		implTypeName := corego.MakePublic(string(src.Dialect) + repo.TypeName + "Impl")

		stub := ast.NewStruct(corego.MakePrivate("Abstract" + repo.TypeName)).
			SetVisibility(ast.PackagePrivate).
			SetComment("...provides function stubs for the interface\n" + repoTypeName.String() + ".")

		stub.AddMethods(
			ast.NewFunc("init").
				SetVisibility(ast.PackagePrivate).
				SetComment("...is called after construction from the factory method.").
				SetBody(ast.NewBlock()),

			ast.NewFunc("context").
				SetVisibility(ast.PackagePrivate).
				SetComment("...is called to get the significant context.\n"+
					"This is not idiomatic, but it is not fine to clutter domain APIs with I/O details either.\n"+
					"This is a compromise to actually customize timeouts a bit, but yes, this could be better.").
				SetBody(ast.NewBlock(ast.NewTpl(`return {{.Use "context.Background"}}()`))),
		)
		corego.ImplementFunctions(repo, stub)

		for _, f := range stub.Methods() {
			f.SetComment(f.CommentText() + "\nOverride this method in the embedding type in another file.")
		}

		for _, m := range repository.Methods {
			var method *ast.Func
			for _, f := range stub.Methods() {
				if f.FunName == m.Name.String() {
					method = f
					break
				}
			}

			if method == nil {
				return core.NewPosError(m.Name, "refers to a non existing interface method")
			}

			method.SetRecName("r")
			if err := implementBody(method, m); err != nil {
				return fmt.Errorf("cannot implement method %s: %w", m.Name, err)
			}
		}

		file.AddNodes(
			ast.NewTpl("// document and assert interface compatibility.\n// Entities must be imported anyway, so we won't loose modularity.\n"),
			ast.NewTpl(`var _ {{.Use (.Get "repoTypeName")}} = ({{.Get "iface"}})(nil)`).
				Put("repoTypeName", repoTypeName.String()).
				Put("iface", implTypeName),
			lang.Term(),
		)

		file.AddTypes(
			ast.NewStruct(implTypeName).
				SetComment(repo.CommentText()).
				AddFields(
					ast.NewField("db", ast.NewSimpleTypeDecl("DBTX")).SetVisibility(ast.PackagePrivate),
				).
				AddEmbedded(ast.NewSimpleTypeDecl(ast.Name(stub.TypeName))),
		)

		file.AddFuncs(
			ast.NewFunc("New" + implTypeName).
				SetComment("...creates a new repository instance.").
				AddParams(
					ast.NewParam("db", ast.NewSimpleTypeDecl("DBTX")),
				).
				AddResults(
					ast.NewParam("", ast.NewTypeDeclPtr(ast.NewSimpleTypeDecl(ast.Name(implTypeName)))),
				).
				SetBody(ast.NewBlock(
					ast.NewTpl(`r := &{{.Get "typename"}}{
										db:db,
									}
									r.{{.Get "stubname"}}.db = db
									r.init()

									return r`).
						Put("typename", implTypeName).
						Put("stubname", stub.TypeName),
				)),

		)

		file.AddNodes(stub)
	}

	return nil
}

func implementBody(fun *ast.Func, method sql.Method) error {
	if len(fun.FunResults) == 0 {
		return core.NewPosError(core.NewNodeFromAst(fun), "method must have at least error! as return value")
	}

	if dec, ok := fun.FunResults[len(fun.FunResults)-1].ParamTypeDecl.(*ast.SimpleTypeDecl); ok {
		if dec.SimpleName != stdlib.Error {
			return core.NewPosError(core.NewNodeFromAst(dec), "expected error! as return type")
		}
	}

	if len(method.Result) == 0 {
		return implementExec(fun, method)
	}

	return implementQuery(fun, method)
}

func implementExec(fun *ast.Func, method sql.Method) error {
	if len(method.Result) == 0 {
		if len(method.Prepare) == 0 {
			return implementExecOne(fun, method)
		}

		switch t := method.Prepare[0].(type) {
		case sql.MapSelOne:
			return implementExecOne(fun, method)
		case sql.MapSelMany:
			return implementExeMany(fun, method)
		default:
			panic("unsupported: " + reflect.TypeOf(t).String())
		}
	}

	return implementQuery(fun, method)
}

func implementExecOne(fun *ast.Func, method sql.Method) error {
	inParamList, err := getMapSelOnePlaceholderList(method.Prepare)
	if err != nil {
		return err
	}

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					if _, err := r.db.ExecContext(r.context(), q, {{.Get "params"}}); err!=nil{
						return {{.Use "fmt.Errorf"}}("cannot execute '%s': %w", q, err)	
					}
			
					return nil
				`).
				Put("query", strconv.Quote(method.Query.String())).
				Put("params", inParamList),
		),
	)

	return nil
}

func implementExeMany(fun *ast.Func, method sql.Method) error {
	many := method.Prepare[0].(sql.MapSelMany)
	paramName := fun.FunParams[0].ParamName
	placeholderList := ""
	for i, lit := range many.Sel {
		if lit.String() == "." {
			placeholderList += paramName + "[i]"
		} else {
			placeholderList += paramName + "[i]" + lit.String()
		}
		if i < len(many.Sel)-1 {
			placeholderList += ", "
		}
	}

	fun.SetBody(
		ast.NewBlock(
			ast.NewTpl(`const q = {{.Get "query"}}
					c := r.context()
					s, err := r.db.PrepareContext(c, q)
					if err != nil{
						return {{.Use "fmt.Errorf"}}("cannot prepare '%s': %w", q, err)	
					}
					defer s.Close()

					for i := range {{.Get "slice"}}{
						if _, err := s.ExecContext(c, {{.Get "params"}}); err!=nil{
							return {{.Use "fmt.Errorf"}}("cannot execute '%s': %w", q, err)	
						}
					}

					return nil
				`).
				Put("query", strconv.Quote(method.Query.String())).
				Put("params", placeholderList).
				Put("slice", paramName),
		),
	)
	return nil
}

func implementQuery(fun *ast.Func, method sql.Method) error {
	if len(method.Result) == 1 {
		switch t := method.Result[0].(type) {
		case sql.MapSelMany:
			return implementFindMany(fun, method)
		case sql.MapSelOne:
			return implementFindOne(fun, method)
		default:
			panic("unsupported: " + reflect.TypeOf(t).String())
		}
	}

	return implementFindMany(fun, method)
}

func getMapSelOnePlaceholderList(mappings []sql.Mapping) (string, error) {
	placeholderList := ""
	for i, mapping := range mappings {
		if sel, ok := mapping.(sql.MapSelOne); ok {
			placeholderList += sel.Sel.String()
			if i < len(mappings)-1 {
				placeholderList += ","
			}
		} else {
			return "", core.NewPosError(sel.Sel, "expected a *MapSelOne* type, you cannot mix")
		}
	}

	return placeholderList, nil
}

func getMapSelOneRefPlaceholderListR(mappings []sql.Mapping, varName string) (string, error) {
	placeholderList := ""
	for i, mapping := range mappings {
		if sel, ok := mapping.(sql.MapSelOne); ok {
			placeholderList += "&"
			if sel.Sel.String() == "." {
				placeholderList += varName
			} else {
				placeholderList += sel.Sel.String()
			}
			if i < len(mappings)-1 {
				placeholderList += ","
			}
		} else {
			return "", core.NewPosError(sel.Sel, "expected a *MapSelOne* type, you cannot mix")
		}
	}

	return placeholderList, nil
}

func getMapSelManyRefPlaceholderListR(mappings []sql.Mapping, varName string) (string, error) {
	placeholderList := ""
	many := mappings[0].(sql.MapSelMany).Sel
	for i, sel := range many {
		placeholderList += "&"
		placeholderList += varName
		if sel.String() == "." {
			// omit
		} else {
			placeholderList += sel.String()
		}
		if i < len(many)-1 {
			placeholderList += ","
		}
	}

	return placeholderList, nil
}

// this thing can only generate properly for a single primitive or a single struct for multiple return.
// Multiple primitives are not supported. Also the return must not be a generic in any way.
func implementFindOne(fun *ast.Func, method sql.Method) error {
	in, err := getMapSelOnePlaceholderList(method.Prepare)
	if err != nil {
		return err
	}

	out, err := getMapSelOneRefPlaceholderListR(method.Result, "i")
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
				Put("query", strconv.Quote(method.Query.String())).
				Put("in", in).
				Put("out", out).
				Put("returnType", string(simpleDecl.SimpleName)),
		),
	)

	return nil
}

func implementFindMany(fun *ast.Func, method sql.Method) error {
	in, err := getMapSelOnePlaceholderList(method.Prepare)
	if err != nil {
		return err
	}

	out, err := getMapSelManyRefPlaceholderListR(method.Result, "t")
	if err != nil {
		return err
	}

	// this another hard assumption
	slice := fun.FunResults[0].ParamTypeDecl.(*ast.SliceTypeDecl)
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
				Put("query", strconv.Quote(method.Query.String())).
				Put("in", in).
				Put("out", out).
				Put("returnType", string(simpleDecl.SimpleName)),
		),
	)

	return nil
}

// simpleLowerCaseName returns a lowercase name which just contains a..z, nothing else.
func simpleLowerCaseName(str string) string {
	str = strings.ToLower(str)
	sb := &strings.Builder{}
	for _, r := range str {
		if r >= 'a' && r <= 'z' {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
