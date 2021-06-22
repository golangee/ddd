package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"strings"
)

func renderConfigs(dst *ast.Mod, src *adl.Module) error {
	if len(src.Executables) > 0 {

		for _, executable := range src.Executables {
			appPkg := astutil.MkPkg(dst, getApplicationPath(dst, executable))

			uberCfg := ast.NewStruct("Configuration").
				SetComment("...contains all aggregated configurations for the entire application and all contained bounded contexts.")

			file := ast.NewFile("config.go").
				SetPreamble(makePreamble(src.Preamble)).
				AddNodes(uberCfg)

			appPkg.AddFiles(
				file,
			)

			for _, path := range executable.BoundedContextPaths {
				bc := astutil.FindPkg(dst, path.String())
				if bc == nil {
					return token.NewPosError(path, "invalid bounded context import path: "+path.String())
				}

				bcCfg := ast.NewStruct(golang.MakePublic(bc.Name) + "Config").
					SetComment("...contains all aggregated configurations for the entire bounded context '" + bc.Name + "'.")

				file.AddNodes(bcCfg)

				if _, err := addConfigHolder(dst, bcCfg, bc.Name, path.String(), pkgCore); err != nil {
					return fmt.Errorf("unable to render core configs: %w", err)
				}

				if _, err := addConfigHolder(dst, bcCfg, bc.Name, path.String(), pkgUsecase); err != nil {
					return fmt.Errorf("unable to render usecase configs: %w", err)
				}

				uberCfg.AddFields(ast.NewField(bcCfg.TypeName, ast.NewSimpleTypeDecl(ast.Name(appPkg.Path+"."+bcCfg.TypeName))))
			}

		}
	}

	return nil
}

func addConfigHolder(mod *ast.Mod, dst *ast.Struct, bcName, bcPath, pathSuffix string) (*ast.Struct, error) {
	holder := ast.NewStruct(golang.MakePublic(bcName) + golang.MakePublic(pathSuffix) + "Config").
		SetComment("...contains all configurations for the '" + pathSuffix + "' layer of '" + bcName + "'.")

	cfgs := findTypes(findPrefixPkgs(mod, golang.MakePkgPath(bcPath, pathSuffix)), func(s stereotype.Struct) bool {
		return s.IsConfiguration()
	})

	if len(cfgs) == 0 {
		return holder, nil
	}

	for _, cfg := range cfgs {
		if err := addConfigUtil(bcName, pathSuffix, cfg); err != nil {
			return nil, fmt.Errorf("unable to add config utility to %s: %w", cfg.TypeName, err)
		}
		holder.AddFields(ast.NewField(cfg.TypeName, ast.NewSimpleTypeDecl(ast.Name(astutil.FullQualifiedName(cfg)))))
	}

	astutil.File(dst).AddNodes(holder)
	dst.AddFields(ast.NewField(holder.TypeName, ast.NewSimpleTypeDecl(ast.Name(astutil.Pkg(dst).Path+"."+holder.TypeName))))

	return holder, nil
}

func addConfigUtil(bcName, pathSuffix string, cfg *ast.Struct) error {
	if cfg.DefaultRecName == "" {
		cfg.DefaultRecName = strings.ToLower(cfg.TypeName[:1])
	}

	if _, err := golang.AddResetFunc(cfg); err != nil {
		return fmt.Errorf("unable to add reset func: %w", err)
	}

	if _, err := golang.AddParseEnvFunc(strings.ToLower(bcName+"_"+pathSuffix), cfg); err != nil {
		return fmt.Errorf("unable to add parse-env func: %w", err)
	}

	if _, err := golang.AddParseFlagFunc(strings.ToLower(bcName+"-"+pathSuffix), cfg); err != nil {
		return fmt.Errorf("unable to add parse-flag func: %w", err)
	}

	return nil
}
