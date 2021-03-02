package modlet

import (
	"fmt"
	validate2 "github.com/golangee/architecture/modlet/validate"
	"github.com/golangee/architecture/yast"
	yastyaml "github.com/golangee/architecture/yast-yaml"
	"github.com/golangee/architecture/yast/validate"
	"github.com/golangee/src/ast"
	"github.com/golangee/src/golang"
	"github.com/golangee/src/render"
)

// A Workspace coordinates multiple modlets to generate all relevant projects or modules.
type Workspace struct {
	Dir string
	Src *ast.Prj  // Src contains the generated source code for the entire project: multi-language-multi-module
	ADL *yast.Obj // ADL contains the architecture description language: language-neutral-multi-module
}

// ParseWorkspace parses the directory for the adl.
func ParseWorkspace(dir string) (*Workspace, error) {
	pkg, err := yastyaml.ParseDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ADL dir: %w", err)
	}

	var wsNode *yast.Obj
	if err := validate.ExpectObj(&wsNode, "workspace.yml/root")(pkg); err != nil {
		return nil, fmt.Errorf("unable to find workspace node: %w", err)
	}

	prjName, err := validate2.YamlDirName(wsNode)
	if err != nil {
		return nil, fmt.Errorf("unable to get project name: %w", err)
	}

	if _, err := validate2.Doc(wsNode.Get(prjName)); err != nil {
		return nil, err
	}

	ws := &Workspace{
		Dir: dir,
		ADL: pkg,
		Src: ast.NewPrj(prjName),
	}

	return ws, nil
}

func (ws *Workspace) Render() (render.Artifact, error) {
	renderer := golang.NewRenderer(golang.Options{})
	artifact, err := renderer.Render(ws.Src)
	if err != nil {
		return nil, fmt.Errorf("unable to render go modules: %w", err)
	}

	return artifact, nil
}

type Modlet interface {
	Apply(prj Workspace, node yast.Node) error
}
