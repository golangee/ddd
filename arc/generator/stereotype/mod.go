package stereotype

import (
	"github.com/golangee/architecture/arc/generator/doc"
	"github.com/golangee/src/ast"
)

const (
	docPrefix  = "docs/content/docs/"
	DocDevelop = docPrefix + "getting-started/development.md"
)

// Mod contains all stereotype annotations for a module instance.
type Mod struct {
	obj *ast.Mod
}

func ModFrom(mod *ast.Mod) Mod {
	return Mod{obj: mod}
}

func (c Mod) Unwrap() *ast.Mod {
	return c.obj
}

func (c Mod) Docs() *Docs {
	v := c.obj.Value(kModuleDocs)
	if f, ok := v.(*Docs); ok {
		return f
	}

	d := &Docs{}
	c.obj.PutValue(kModuleDocs, d)

	return d
}

type Docs struct {
	Files map[string][]doc.Node
}

func (d *Docs) Append(path string, node ...doc.Node) *Docs {
	if d.Files == nil {
		d.Files = map[string][]doc.Node{}
	}

	nodes := d.Files[path]
	nodes = append(nodes, node...)
	d.Files[path] = nodes

	return d
}
