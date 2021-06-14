package stereotype

import (
	"github.com/golangee/architecture/arc/generator/doc"
	"github.com/golangee/src/ast"
	"strings"
)

const (
	DocDevelop = "getting-started/development.md"
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

// Ident returns a unique identifier for the module or the empty string.
func (c Mod) Ident() string {
	v := c.obj.Value(kModShortName)
	if v == nil {
		return ""
	}

	return v.(string)
}

func (c Mod) SetIdent(ident string) {
	c.obj.PutValue(kModShortName, ident)
}

func (c Mod) Docs() *Docs {
	v := c.obj.Value(kModuleDocs)
	if f, ok := v.(*Docs); ok {
		return f
	}

	d := &Docs{mod: c}
	c.obj.PutValue(kModuleDocs, d)

	return d
}

type Docs struct {
	mod   Mod
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

func Doc(m *ast.Mod, lang, suffix string, node ...doc.Node) *Docs {
	return ModFrom(m).Docs().DocWithLang(lang, suffix, node...)
}

// Doc appends the given nodes as a default language (EN)
// boxed into a module specific documentation path, so that
// later all module documentations can easily be merged into
// an uber-documentation.
func (d *Docs) Doc(suffix string, node ...doc.Node) *Docs {
	return d.DocWithLang("", suffix, node...)
}

// Prefix returns the language and module specific prefix path.
func (d *Docs) Prefix(lang string) string {
	base := "docs/content"
	if lang != "" {
		base += "." + strings.ToLower(lang)
	}

	base += "/docs/" + strings.ToLower(d.mod.Ident()) + "/"

	return base
}

func (d *Docs) DocWithLang(lang, suffix string, node ...doc.Node) *Docs {
	return d.Append(d.Prefix(lang)+suffix, node...)
}
