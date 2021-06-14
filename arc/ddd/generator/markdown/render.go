package markdown

import (
	"bytes"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/doc"
	"github.com/golangee/architecture/arc/generator/doc/markdown"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/src/ast"
	"io/fs"
	"strconv"
	"strings"
)

import _ "embed"

//go:embed hugo/config.toml
var configToml string

func RenderModule(dst *ast.Prj, prj *adl.Project, src *adl.Module) error {
	for _, mod := range dst.Mods {
		docs := stereotype.ModFrom(mod)
		if len(docs.Docs().Files) > 0 {
			for path, nodes := range docs.Docs().Files {
				var buf bytes.Buffer
				pkg := astutil.MkPkg(mod, golang.MakePkgPath(mod.Name, golang.PkgPathDir(path)))
				buf.Write(markdown.Render(doc.NewComment(src.Preamble.Generator)))
				for _, node := range nodes {
					buf.Write(markdown.Render(node))
				}

				pkg.AddRawFiles(ast.NewRawFile(golang.PkgPathBase(path), "text/markdown", buf.Bytes()))
			}

			modIdent := stereotype.ModFrom(mod).Ident()
			pkg := astutil.MkPkg(mod, golang.MakePkgPath(mod.Name, "docs"))
			pkg.AddRawFiles(ast.NewRawFile("config.toml", "text/x-toml", []byte(strings.ReplaceAll(configToml, "{{.Title}}", modIdent))))
		}
	}

	return nil
}

func renderIndex(dst *ast.Prj, prj *adl.Project, src *adl.Module) error {
	for _, mod := range dst.Mods {
		menuIndexBuf := bytes.Buffer{}
		menuIndexBuf.WriteString("---\nheadless: true\n---\n\n")
		modIdent := stereotype.ModFrom(mod).Ident()

		docs := stereotype.ModFrom(mod)

		if len(docs.Docs().Files) > 0 {
			for path, nodes := range docs.Docs().Files {
				hierarchyPart := path[strings.LastIndex(path, modIdent):]
				for level, _ := range strings.Split(hierarchyPart, "/") {
					for i := 0; i < level; i++ {
						menuIndexBuf.WriteString("  ")
					}
					menuIndexBuf.WriteString("- [")
					menuIndexBuf.WriteString(caption(nodes...))
					menuIndexBuf.WriteString("]({{< relref ")
					menuIndexBuf.WriteString(strconv.Quote(hugoBookHtmlIndexPath(path)))
					menuIndexBuf.WriteString(" >}})\n")

					//menuIndexBuf.WriteString(`  - [With ToC]({{< relref "/docs/example/table-of-contents/with-toc" >}})`)
				}

			}

			pkg := astutil.MkPkg(mod, golang.MakePkgPath(mod.Name, golang.MakePkgPath("docs/content/menu/")))
			pkg.AddRawFiles(ast.NewRawFile("index.md", "text/markdown", menuIndexBuf.Bytes()))

		}
	}

	return nil
}

func copyFS(f fs.FS, src string, dst *ast.Pkg) error {
	files, err := fs.ReadDir(f, src)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			cPkg := astutil.MkPkg(astutil.Mod(dst), dst.Path+"/"+file.Name())
			if err := copyFS(f, src+"/"+file.Name(), cPkg); err != nil {
				return err
			}
		} else {
			buf, err := fs.ReadFile(f, src+"/"+file.Name())
			if err != nil {
				return err
			}
			dst.AddRawFiles(ast.NewRawFile(file.Name(), "application/octet-stream", buf))
		}
	}

	return nil
}

func caption(nodes ...doc.Node) string {
	for _, n := range nodes {
		if el, ok := n.(*doc.Element); ok {
			if strings.HasPrefix(el.Name, "h") {
				return el.TextContent()
			}

			for _, child := range el.Children {
				if c := caption(child); c != "" {
					return c
				}
			}
		}
	}

	return ""
}

// hugoBookHtmlIndexPath takes something like docs/content/docs/server/getting-started/development.md
// and returns /docs/server/getting-started/development from it.
func hugoBookHtmlIndexPath(path string) string {
	const prefix = "docs/content"
	if strings.HasPrefix(path, prefix) {
		path = path[len(prefix):]
	}

	if strings.HasSuffix(path, ".md") {
		path = path[:len(path)-2]
	}

	return path
}
