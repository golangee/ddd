package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
)

func renderExecs(dst *ast.Mod, src *adl.Module) error {
	if len(src.Executables) > 0 {
		cmd := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "cmd"))
		cmd.SetComment("...represents individual main executables and can simply be build using the *go build <import path> @ latest* command.")
		cmd.SetPreamble(makePreamble(src.Preamble))

		for _, executable := range src.Executables {
			cmdPkg := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "cmd", executable.Name.String()))
			cmdPkg.Name = "main"
			cmdPkg.SetComment(executable.Comment.String())
			cmdPkg.SetPreamble(makePreamble(src.Preamble))
			stereotype.PkgFrom(cmdPkg).SetIsCMDPkg(true)

			cmdPkg.AddFiles(
				ast.NewFile("main.go").
					SetPreamble(makePreamble(src.Preamble)).
					AddFuncs(
						ast.NewFunc("main").
							SetVisibility(ast.Private).
							SetBody(ast.NewBlock(
								ast.NewTpl(
									`ctx, done := {{.Use "os/signal.NotifyContext"}}({{.Use "context.Background"}}(), {{.Use "syscall.SIGINT"}}, {{.Use "syscall.SIGTERM"}})
										 
										 logger := {{.Use (.Get "logger")}}()
										 logger = log.WithFields(logger,
										 {{.Use "github.com/golangee/log/ecs.Log"}}("{{.Get "appName"}}"), 
											log.V("build_id",{{.Use (.Get "buildinfo")}}.ID()), 
											log.V("build_tag",{{.Use (.Get "buildinfo")}}.Tag()),
										 )

										 ctx = {{.Use "github.com/golangee/log.WithLogger"}}(ctx, logger)

										 err := realMain(ctx)
										 done()

										 if err != nil {
											logger.Println(ecs.Fatal(), err)
										 }

										 logger.Println(ecs.Info(), "successful shutdown")
								`).
									Put("appName", executable.Name.String()).
									Put("logger", dst.Name+"/internal/logging.NewLoggerFromEnv").
									Put("buildinfo", dst.Name+"/internal/buildinfo.Build"),
							)),

						ast.NewFunc("realMain").
							SetVisibility(ast.Private).
							AddParams(ast.NewParam("ctx", ast.NewSimpleTypeDecl("context.Context"))).
							AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
							SetBody(ast.NewBlock(
								lang.TryDefine(ast.NewIdent("a"), lang.CallStatic(ast.Name(getApplicationPath(dst, executable)+".NewApplication"), ast.NewIdent("ctx")), "cannot create application '"+executable.Name.String()+"'"),
								lang.TryDefine(nil, lang.CallIdent("a", "Run", ast.NewIdent("ctx")), "cannot run application '"+executable.Name.String()+"'"),
								ast.NewReturnStmt(ast.NewIdentLit("nil")),
							)),
					),
			)
		}
	}

	return nil
}
