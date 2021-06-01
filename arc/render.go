package arc

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/ddd/generator/golang"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	golang2 "github.com/golangee/src/golang"
	"github.com/golangee/src/render"
)

func Render(prj *adl.Project) (render.Artifact, error) {
	astPrj := ast.NewPrj(prj.Name.String())

	for _, module := range prj.Modules {
		if module.Generator == nil {
			return nil, token.NewPosError(module.Name, "module has no generator settings")
		}

		noGenSettings := true
		if module.Generator.Go != nil {
			noGenSettings = false
			if err := golang.RenderModule(astPrj, prj, module); err != nil {
				return nil, token.NewPosError(module.Name, "unable to render module").SetCause(err)
			}
		}

		if noGenSettings {
			return nil, token.NewPosError(module.Name, "module has no generator settings details")
		}
	}

	renderer := golang2.NewRenderer(golang2.Options{})
	a, err := renderer.Render(astPrj)
	if err != nil {
		return a, fmt.Errorf("unable to render prj %v: %w", astPrj.Name, err)
	}

	return a, nil
}
