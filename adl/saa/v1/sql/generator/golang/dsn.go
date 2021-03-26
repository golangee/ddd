package golang

import (
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"github.com/golangee/src/stdlib/strings"
)

func addDSNFunc(opt *ast.Struct) {
	opt.AddMethods(
		ast.NewFunc("DSN").
			SetComment("...returns the options as a fully serialized datasource name.\n" +
				"The returned string is of the form:\n" +
				"  username:password@protocol(address)/dbname?param=value").
			SetPtrReceiver(true).
			SetRecName(opt.DefaultRecName).
			AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.String)).SetComment("... the DSN value.")).
			SetBody(ast.NewBlock(
				strings.NewStrBuilder("sb",
					urlEscape("User"),
					ast.NewStrLit(":"),
					urlEscape("Password"),
					ast.NewStrLit("@"),
					lang.Attr("Protocol"),
					ast.NewStrLit("("),
					lang.Attr("Address"),
					ast.NewStrLit(")/"),
					lang.Attr("Database"),
					ast.NewStrLit("?"),

					ast.NewStrLit("charset="),
					urlEscape("Charset"),

					ast.NewStrLit("&collation="),
					urlEscape("Collation"),

					ast.NewStrLit("&maxAllowedPacket="),
					lang.Itoa(ast.NewIdent("MaxAllowedPacket")),

					ast.NewStrLit("&sql_mode="),
					lang.Attr("SqlMode"),

					ast.NewStrLit("&tls="),
					lang.ToString(ast.NewIdent("tls")),

					ast.NewStrLit("&timeout="),
					lang.Itoa(ast.NewIdent("Timeout")),

					ast.NewStrLit("&writeTimeout="),
					lang.Itoa(ast.NewIdent("WriteTimeout")),

				),
				ast.NewReturnStmt(ast.NewCallExpr(ast.NewSelExpr(ast.NewIdent("sb"),ast.NewIdent("String")))),
		)),
	)
}

func urlEscape(attrName string) ast.Expr {
	return lang.CallStatic("net/url.QueryEscape", lang.Attr(attrName))
}
