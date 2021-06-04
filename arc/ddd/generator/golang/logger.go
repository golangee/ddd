package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/src/ast"
)

func renderLogger(dst *ast.Mod, src *adl.Module) error {
	log := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "internal", "logging"))
	log.SetPreamble(makePreamble(src.Preamble)).
		SetComment("...configures the modules logger.")

	log.AddFiles(
		ast.NewFile("logger.go").
			AddFuncs(
				ast.NewFunc("NewLoggerFromEnv").
					SetComment("...creates a new logger based on the current environment.").
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl("github.com/golangee/log.Logger"))).
					SetBody(ast.NewBlock(
						ast.NewTpl(
							`logger := log.NewLogger()
								logger = log.WithFields(logger, log.V("build_id","tbd"), log.V("build_tag","tbd"))
								
								return logger
`,
						),
					)),
			),
	)

	return nil
}
