package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/src/ast"
)

func RenderModule(dst *ast.Prj, src *adl.Module) error {
	if src.Generator == nil {
		return fmt.Errorf("cannot render a non-target project: %s", src.Name)
	}

	if src.Generator.Go == nil {
		return fmt.Errorf("cannot render a non-go module: %s -> %s", src.Name, src.Generator.OutDir)
	}

	mod := astutil.MkMod(dst, src.Generator.Go.Module.String())
	mod.SetLang(ast.LangGo)

	return nil
}
