package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"path/filepath"
	"strings"
)

func createSQLLayer(ctx *genctx, rslv *resolver, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())

	if err := createSQLUtil(ctx, bc, sql); err != nil {
		return err
	}

	for _, repo := range sql.Repositories() {
		file := ctx.newFile(layerPath, text.Safename(repo.InterfaceName()), "").
			SetPackageDoc(sql.Description())

		iface := bc.SPIServiceByName(repo.InterfaceName())
		if iface == nil {
			panic("illegal state: validate the model first")
		}

		repoSpec := ctx.repoSpecByName(repo.InterfaceName())
		if repoSpec == nil {
			panic("illegal state: validate the model first and define core layer before sql layer")
		}

		impl := src.Implement(repoSpec.iface, true)
		impl.SetName(sql.Name() + repo.InterfaceName())
		impl.AddFields(
			src.NewField("db", src.NewTypeDecl("DBTX")),
		)
		impl.SetDoc("...is an implementation of the " + pkgNameCore + "." + repo.InterfaceName() + " defined as SPI/driven port in the domain/core layer.\nThe queries are specific for the " + strings.ToLower(sql.Name()) + " dialect.")
		file.AddTypes(impl)

		for _, method := range impl.Methods() {
			sqlSpec := repo.ImplementationByName(method.Name())
			if sqlSpec == nil {
				panic("illegal state: undefined method: validate the model first")
			}
			body := src.NewBlock()
			method.AddBody(body)
			if len(sqlSpec.Row()) == 0 {
				createSQLExec(sqlSpec, method, body)
			} else {
				if method.Results()[0].Decl().Qualifier() == "[]" {
					createSQLQueryMany(sqlSpec, method, body)
				} else {
					createSQLQueryOne(sqlSpec, method, body)
				}
			}

		}
	}

	return nil
}

func createSQLQueryMany(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")

	body.Add("r, err := ", method.ReceiverName(), ".db.QueryContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")
	body.If("err!=nil", src.NewBlock(
		"return nil, ", src.NewTypeDecl("fmt.Errorf"), "(\"QueryContext failed: %w\",err)",
	))
	body.AddLine("defer r.Close()")

	body.AddLine("var l ", method.Results()[0].Decl().Clone())
	body.AddLine("for r.Next() {")
	body.AddLine("var i ", method.Results()[0].Decl().Params()[0].Clone())
	body.Add("if err := r.Scan(")
	for _, row := range sqlSpec.Row() {
		body.Add(makeSqlVarAccess("i", string(row)), ",")
	}
	body.AddLine(");err!=nil{")
	body.AddLine("return nil, ", src.NewTypeDecl("fmt.Errorf"), "(\"scan failed: %w\",err)")
	body.AddLine("}")
	body.AddLine("l = append(l, i)")
	body.AddLine("}")

	body.AddLine("err = r.Close()")
	body.Check("err", "cannot close rows", "l")
	body.NewLine()

	body.AddLine("err = r.Err()")
	body.Check("err", "query failed", "l")
	body.NewLine()

	body.AddLine("return l, nil")
}

func makeSqlVarAccess(name, accessor string) string {
	if accessor == "&." {
		return "&" + name
	}

	if strings.HasPrefix(accessor, "&.") {
		return "&" + name + "." + accessor[2:]
	}

	if accessor == "." {
		return name
	}

	if strings.HasPrefix(accessor, ".") {
		return name + accessor
	}

	return accessor
}

func createSQLQueryOne(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")

	body.AddLine("var i ", method.Results()[0].Decl().Clone())
	body.Add("r, err := ", method.ReceiverName(), ".db.QueryContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")
	body.If("err!=nil", src.NewBlock(
		"return i, ", src.NewTypeDecl("fmt.Errorf"), "(\"QueryContext failed: %w\",err)",
	))
	body.AddLine("defer r.Close()")

	body.AddLine("for r.Next() {")
	body.Add("if err := r.Scan(")
	for _, row := range sqlSpec.Row() {
		body.Add(makeSqlVarAccess("i", string(row)), ",")
	}
	body.AddLine(");err!=nil{")
	body.AddLine("return i, ", src.NewTypeDecl("fmt.Errorf"), "(\"scan failed: %w\",err)")
	body.AddLine("}")

	body.AddLine("err = r.Close()")
	body.Check("err", "cannot close rows", "i")
	body.NewLine()

	body.AddLine("err = r.Err()")
	body.Check("err", "query failed", "i")
	body.NewLine()

	body.AddLine("return i, err")
	body.AddLine("}")

	body.AddLine("return i, ", src.NewTypeDecl("fmt.Errorf(\"empty result set\")"))
}

func createSQLExec(sqlSpec *ddd.GenFuncSpec, method *src.FuncBuilder, body *src.Block) {
	body.AddLine("const q = \"", string(sqlSpec.RawStatement())+"\"")
	body.Add("_, err := ", method.ReceiverName(), ".db.ExecContext(ctx, q ")
	for _, p := range sqlSpec.Params() {
		body.Add(",", p)
	}
	body.AddLine(")")

	body.Check("err", "ExecContext failed")
	body.NewLine()

	body.AddLine("return nil")
}

func createSQLUtil(ctx *genctx, bc *ddd.BoundedContextSpec, sql ddd.SQLLayer) error {
	bcPath := filepath.Join("internal", text.Safename(bc.Name()))
	layerPath := filepath.Join(bcPath, sql.Name())

	file := ctx.newFile(layerPath, "db", "").
		SetPackageDoc(sql.Description())

	file.AddTypes(
		src.NewInterface("DBTX").
			SetDoc("...abstracts from a concrete sql.DB or sql.Tx dependency.").
			AddMethods(
				src.NewFunc("ExecContext").SetDoc("...represents an according call to sql.DB or sql.Tx").
					AddParams(
						src.NewParameter("ctx", src.NewTypeDecl("context.Context")),
						src.NewParameter("query", src.NewTypeDecl("string")),
						src.NewParameter("args", src.NewTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						src.NewParameter("", src.NewTypeDecl("database/sql.Result")),
						src.NewParameter("", src.NewTypeDecl("error")),
					),

				src.NewFunc("QueryContext").SetDoc("...represents an according call to sql.DB or sql.Tx").
					AddParams(
						src.NewParameter("ctx", src.NewTypeDecl("context.Context")),
						src.NewParameter("query", src.NewTypeDecl("string")),
						src.NewParameter("args", src.NewTypeDecl("interface{}")),
					).
					SetVariadic(true).
					AddResults(
						src.NewParameter("", src.NewTypeDecl("database/sql.Rows")),
						src.NewParameter("", src.NewTypeDecl("error")),
					),
			),
	)

	return nil
}
