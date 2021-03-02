package arch

import (
	"fmt"
	"github.com/golangee/architecture/modlet"
	"github.com/golangee/architecture/modlet/arch/iface"
	"github.com/golangee/architecture/objn"
	objnyaml "github.com/golangee/architecture/objn-yaml"
	"github.com/golangee/architecture/yast"
	yastyaml "github.com/golangee/architecture/yast-yaml"
	"golang.org/x/mod/semver"
	"strings"
)

type funcModlet func(prj modlet.Project, node yast.Node) error

func (f funcModlet) Apply(prj modlet.Project, node yast.Node) error {
	return f(prj, node)
}

type ModletFactory struct {
	Version string
	Modlet  modlet.Modlet
}

var modletFactories map[string][]ModletFactory

func init() {
	modletFactories = map[string][]ModletFactory{}
	modletFactories["arch/Interface"] = []ModletFactory{
		{Version: "v0.0.1", Modlet: funcModlet(iface.ApplyV0_0_1)},
	}
}

func Apply(node yast.Node) error {
	nodes, err := yast.Filter(node, func(n yast.Node) (bool, error) {
		if n, ok := path[len(path)-1].(yast.Map); ok {
			if n, ok := n.Get("apply").(objn.Seq); ok {
				for i := 0; i < n.Count(); i++ {
					lit := n.Get(i).(objn.Lit)
					plugin := lit.String()
					idx := strings.LastIndex(plugin, "@")
					version := "v0.0.0"
					if idx > 0 {
						version = plugin[idx+1:]
					}

					if !semver.IsValid(version) {
						return false, objn.NewPosErrorMark(path, lit, "invalid version syntax", version)
					}

					fmt.Println("found => " + plugin)
				}
				return true, nil
			}
		}
		return false, nil
	})

	if err != nil {
		return err
	}

	fmt.Println(nodes)
	return err
}

func Build(dir string) error {
	pkg, err := yastyaml.ParseDir(dir)
	if err != nil {
		return fmt.Errorf("unable to parse yaml dir: %w", err)
	}

	return Apply(pkg)
}
