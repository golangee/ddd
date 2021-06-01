package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/src/ast"
	"strings"
)

func RenderModule(dst *ast.Prj, src *adl.Module) error {
	if src.Generator == nil {
		return fmt.Errorf("cannot render a non-target project: %s", src.Name)
	}

	if src.Generator.Go == nil {
		return fmt.Errorf("cannot render a non-go module: %s -> %s", src.Name, src.Generator.OutDir)
	}


	astutil.MkMod(dst, src.Generator.Go.Module.String()).
		SetLang(ast.LangGo).
		SetOutputDirectory(src.Generator.OutDir.String())

	domainPkg := src.Generator.Go.Module.String()+"/"+src.Domain.Name
	return nil
}

