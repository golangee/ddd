package golang

import (
	"encoding/json"
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/architecture/arc/token"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/stdlib"
	"github.com/golangee/src/stdlib/lang"
	"strings"
)

const ifParseEnv = "if err:={{.Get `rec`}}.{{.Get `field`}}.ParseEnv();err!=nil{\nreturn {{.Use `fmt.Errorf`}}(\"cannot parse '{{.Get `field`}}': %w\",err)}\n"

func renderConfigs(dst *ast.Mod, src *adl.Module) error {
	if len(src.Executables) == 0 {
		return nil
	}

	for _, executable := range src.Executables {
		appPkg := astutil.MkPkg(dst, getApplicationPath(dst, executable))

		uberCfg := ast.NewStruct("Configuration").
			SetComment("...contains all aggregated configurations for the entire application and all contained bounded contexts.").
			SetDefaultRecName("c")

		uberResetBody := ast.NewBlock()
		uberParseEnvBody := ast.NewBlock()
		uberConfigureFlagsBody := ast.NewBlock()

		uberCfg.AddMethods(
			ast.NewFunc("Reset").
				SetComment("...restores this instance to the default state.").
				SetPtrReceiver(true).
				SetRecName(uberCfg.DefaultRecName).
				SetBody(uberResetBody),

			ast.NewFunc("ConfigureFlags").
				SetComment("...configures the flags to be ready to get evaluated.\nYou can only use it once, otherwise the flag package will panic.").
				SetPtrReceiver(true).
				SetRecName(uberCfg.DefaultRecName).
				SetBody(uberConfigureFlagsBody),

			ast.NewFunc("ParseEnv").
				SetPtrReceiver(true).
				SetRecName(uberCfg.DefaultRecName).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetComment("...tries to parse the environment variables into this instance.").
				SetBody(uberParseEnvBody),

			ast.NewFunc("ParseFile").
				SetPtrReceiver(true).
				SetRecName(uberCfg.DefaultRecName).
				AddParams(ast.NewParam("filename", ast.NewSimpleTypeDecl(stdlib.String))).
				AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
				SetBody(ast.NewBlock(
					ast.NewTpl(
						`file,err := {{.Use "os.Open"}}(filename)
							if err != nil {
								return {{.Use "fmt.Errorf"}}("cannot open configuration file: %w",err)
							}

							defer file.Close() // intentionally ignoring read-only error on close

							dec := {{.Use "encoding/json.NewDecoder"}}(file)
							if err := dec.Decode({{.Get "rec"}}); err!=nil {
								return {{.Use "fmt.Errorf"}}("cannot decode json: %w",err)
							}

							return nil
					`).Put("rec", uberCfg.DefaultRecName),
				)),
		)

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

			bcCfg.SetDefaultRecName(strings.ToLower(bcCfg.TypeName[:1]))
			bcCfg.AddMethods(
				ast.NewFunc("Reset").
					SetComment("...restores this instance to the default state.").
					SetPtrReceiver(true).
					SetRecName(bcCfg.DefaultRecName).
					SetBody(ast.NewBlock()),

				ast.NewFunc("ConfigureFlags").
					SetComment("...configures the flags to be ready to get evaluated.\nYou can only use it once, otherwise the flag package will panic.").
					SetPtrReceiver(true).
					SetRecName(bcCfg.DefaultRecName).
					SetBody(ast.NewBlock()),

				ast.NewFunc("ParseEnv").
					SetPtrReceiver(true).
					SetRecName(bcCfg.DefaultRecName).
					AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
					SetComment("...tries to parse the environment variables into this instance.").
					SetBody(ast.NewBlock()),
			)

			uberResetBody.Add(astutil.CallMember(uberCfg.DefaultRecName, golang.MakePublic(bc.Name), "Reset"))
			uberConfigureFlagsBody.Add(astutil.CallMember(uberCfg.DefaultRecName, golang.MakePublic(bc.Name), "ConfigureFlags"))
			uberParseEnvBody.Add(ast.NewTpl(ifParseEnv).
				Put("field", golang.MakePublic(bc.Name)).
				Put("rec", uberCfg.DefaultRecName),
			)

			file.AddNodes(bcCfg)

			if _, err := addConfigHolder(dst, bcCfg, bc.Name, path.String(), pkgCore); err != nil {
				return fmt.Errorf("unable to render core configs: %w", err)
			}

			if _, err := addConfigHolder(dst, bcCfg, bc.Name, path.String(), pkgUsecase); err != nil {
				return fmt.Errorf("unable to render usecase configs: %w", err)
			}

			uberCfg.AddFields(ast.NewField(golang.MakePublic(bc.Name), ast.NewSimpleTypeDecl(ast.Name(appPkg.Path+"."+bcCfg.TypeName))))

			astutil.MethodByName(bcCfg, "ParseEnv").Body().Add(
				lang.Term(),
				ast.NewReturnStmt(ast.NewIdentLit("nil")),
			)
		}

		uberParseEnvBody.Add(
			lang.Term(),
			ast.NewReturnStmt(ast.NewIdentLit("nil")),
		)

		exampleJson, err := exampleJson(uberCfg)
		if err != nil {
			return fmt.Errorf("unable to encode example json: %w", err)
		}

		astutil.MethodByName(uberCfg, "ParseFile").
			SetComment("...tries to parse a json file into this instance. Only the defined values are overridden.\n\nExample JSON\n\n  " + exampleJson)
	}

	return nil
}

func exampleJson(cfg *ast.Struct) (string, error) {
	v, err := golang.SimulateDefaultJson(cfg)
	if err != nil {
		return "", fmt.Errorf("cannot simulate default json encoding: %w", err)
	}

	buf, err := json.MarshalIndent(v, "  ", "  ")
	if err != nil {
		return "", fmt.Errorf("cannot marshal default json encoding: %w", err)
	}

	return string(buf), nil
}

func addConfigHolder(mod *ast.Mod, dst *ast.Struct, bcName, bcPath, pathSuffix string) (*ast.Struct, error) {
	holder := ast.NewStruct(golang.MakePublic(bcName) + golang.MakePublic(pathSuffix) + "Config").
		SetComment("...contains all configurations for the '" + pathSuffix + "' layer of '" + bcName + "'.").
		SetDefaultRecName(strings.ToLower(bcName[:1]))

	resetBody := ast.NewBlock()
	parseEnvBody := ast.NewBlock()
	configureFlagsBody := ast.NewBlock()
	holder.AddMethods(
		ast.NewFunc("Reset").
			SetPtrReceiver(true).
			SetRecName(holder.DefaultRecName).
			SetComment("...restores this instance to the default state.").
			SetBody(resetBody),

		ast.NewFunc("ParseEnv").
			SetPtrReceiver(true).
			SetRecName(holder.DefaultRecName).
			AddResults(ast.NewParam("", ast.NewSimpleTypeDecl(stdlib.Error))).
			SetComment("...tries to parse the environment variables into this instance.").
			SetBody(parseEnvBody),

		ast.NewFunc("ConfigureFlags").
			SetPtrReceiver(true).
			SetRecName(holder.DefaultRecName).
			SetComment("...configures the flags to be ready to get evaluated.\nYou can only use it once, otherwise the flag package will panic.").
			SetBody(configureFlagsBody),
	)

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

		f := ast.NewField(cfg.TypeName, ast.NewSimpleTypeDecl(ast.Name(astutil.FullQualifiedName(cfg))))
		resetBody.Add(astutil.CallMember(holder.DefaultRecName, f.FieldName, "Reset"))
		configureFlagsBody.Add(astutil.CallMember(holder.DefaultRecName, f.FieldName, "ConfigureFlags"))
		parseEnvBody.Add(ast.NewTpl(ifParseEnv).
			Put("field", f.FieldName).
			Put("rec", holder.DefaultRecName),
		)
		holder.AddFields(f)
	}

	parseEnvBody.Add(
		lang.Term(),
		ast.NewReturnStmt(ast.NewIdentLit("nil")),
	)

	astutil.File(dst).AddNodes(holder)
	field := ast.NewField(golang.MakePublic(pathSuffix), ast.NewSimpleTypeDecl(ast.Name(astutil.Pkg(dst).Path+"."+holder.TypeName)))
	dst.AddFields(field)

	parentReset := astutil.MethodByName(dst, "Reset")
	parentReset.Body().Add(
		astutil.CallMember(dst.DefaultRecName, field.FieldName, "Reset"),
	)

	parentConfigureFlags := astutil.MethodByName(dst, "ConfigureFlags")
	parentConfigureFlags.Body().Add(
		astutil.CallMember(dst.DefaultRecName, field.FieldName, "ConfigureFlags"),
	)

	parentParseEnv := astutil.MethodByName(dst, "ParseEnv")
	parentParseEnv.Body().Add(ast.NewTpl(ifParseEnv).
		Put("field", field.FieldName).
		Put("rec", dst.DefaultRecName),
	)

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
