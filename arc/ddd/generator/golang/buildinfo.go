package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/src/ast"
)

func renderBuildInfo(dst *ast.Mod, src *adl.Module) error {
	log := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name, "internal", "buildinfo"))
	log.SetPreamble(makePreamble(src.Preamble)).
		SetComment("...provides build environment information injected at linker time during build.")

	log.AddFiles(
		ast.NewFile("buildinfo.go").
			SetPreamble(makePreamble(src.Preamble)).
			AddNodes(
				ast.NewTpl(`
					var(
						// BuildID is usually a strict monotonic increasing build number.
						BuildID string = "unknown"
					
						// BuildTag usually refers to the VCS branch or tag name.
						BuildTag string = "unknown"
					)

					// buildInfo is a private type to allow an interface based access to the environment variables.
					type buildInfo struct{}
					
					// ID returns the current BuildID.
					func (_ buildInfo) ID() string {
						return BuildID
					}
					
					// Tag returns the current BuildTag.
					func (_ buildInfo) Tag() string {
						return BuildTag
					}
					
					// Build provides the build information about the application.
					var Build buildInfo


`),
			),

	)

	return nil
}
