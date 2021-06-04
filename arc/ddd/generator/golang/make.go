package golang

import (
	"fmt"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/src/ast"
)

func renderGolangciyaml(dst *ast.Mod, src *adl.Module) error {
	const cfg = `
run:
  timeout: '5m'

  build-tags:
  - 'all'

  skip-dirs:

  skip-dirs-use-default: false

  modules-download-mode: 'readonly'

  allow-parallel-runners: true

linters:
  enable:
  - 'asciicheck'
  - 'bodyclose'
  - 'deadcode'
  - 'depguard'
  - 'dogsled'
  - 'errcheck'
  - 'errorlint'
  - 'exportloopref'
  - 'gofmt'
  - 'gofumpt'
  - 'goheader'
  - 'goimports'
  - 'revive'
  - 'gomodguard'
  - 'goprintffuncname'
  - 'gosec'
  - 'gosimple'
  - 'govet'
  - 'ineffassign'
  - 'makezero'
  - 'misspell'
  - 'noctx'
  - 'paralleltest'
  - 'prealloc'
  - 'predeclared'
  - 'sqlclosecheck'
  - 'staticcheck'
  - 'structcheck'
  - 'stylecheck'
  - 'typecheck'
  - 'unconvert'
  - 'unused'
  - 'varcheck'
  - 'whitespace'

issues:
  exclude:
  - '^SA3000:' # staticcheck: redundant, fixed in #34129
  - '^Range statement' # paralleltest: usually false positives

  max-issues-per-linter: 0

  max-same-issues: 0

severity:
  default-severity: error
`

	pkg := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name))
	pkg.AddRawFiles(ast.NewRawTpl("golangci.yaml", "text/x-yaml", ast.NewTpl(
		makeEscapedPreamble(src.Preamble, "# ")+cfg,
	)))

	return nil
}

func renderMakefile(dst *ast.Mod, src *adl.Module) error {
	const mkfile = `

test: ## Executes all tests.
	@go test -race -timeout=5m ./...
.PHONY: test

lint: ## Runs the linter. See golangci.yaml for details.
	@command -v golangci-lint > /dev/null 2>&1 || (cd $${TMPDIR} && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1)
	golangci-lint run --config golangci.yaml
.PHONY: lint

help: ## Shows this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help


# required by GNU standard

dist: ## builds for all major platforms
{{- range .Get "dists"}}
	GOOS={{.Os}} GOARCH={{.Arch}} go build -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/{{.Os}}_{{.Arch}}/app.wasm ${MAIN_PATH}
{{- end}}
.PHONY: dist

all: generate check build ## Compiles all programs
.PHONY: all

.DEFAULT_GOAL := all
`

	pkg := astutil.MkPkg(dst, golang.MakePkgPath(dst.Name))
	pkg.AddRawFiles(ast.NewRawTpl("Makefile", "text/x-makefile", ast.NewTpl(
		makeEscapedPreamble(src.Preamble, "# ")+mkfile,
	).Put("dists",src.Generator.Go.GoDist)))

	if err := renderGolangciyaml(dst, src); err != nil {
		return fmt.Errorf("cannot create golangci.yml: %w", err)
	}

	return nil
}
