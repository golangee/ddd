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
			if err := golang.RenderModule(astPrj, module); err != nil {
				return nil, token.NewPosError(module.Name, "unable to render module")
			}
		}

		if noGenSettings {
			return nil, token.NewPosError(module.Name, "module has no generator settings details")
		}
	}

	var renderedWorkspace render.Dir
	for _, mod := range astPrj.Mods {
		if mod.Target.Lang != ast.LangGo {
			return nil, fmt.Errorf("module '%v': language unsupported: '%v'", mod.Name, mod.Target.Lang)
		}

		renderer := golang2.NewRenderer(golang2.Options{})
		a, err := renderer.Render(mod)
		if err != nil {
			return a, fmt.Errorf("unable to render module %v", mod.Name)
		}

		renderedWorkspace.Dirs = append(renderedWorkspace.Dirs, a.(*render.Dir))
	}

	return &renderedWorkspace, nil
}
