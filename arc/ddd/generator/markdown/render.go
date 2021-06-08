package markdown

import (
	"bytes"
	"embed"
	"github.com/golangee/architecture/arc/adl"
	"github.com/golangee/architecture/arc/generator/astutil"
	"github.com/golangee/architecture/arc/generator/doc"
	"github.com/golangee/architecture/arc/generator/doc/markdown"
	"github.com/golangee/architecture/arc/generator/golang"
	"github.com/golangee/architecture/arc/generator/stereotype"
	"github.com/golangee/src/ast"
	"io/fs"
)

import _ "embed"

//go:embed hugo
var hugo embed.FS

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

			pkg := astutil.MkPkg(mod, golang.MakePkgPath(mod.Name, "docs"))
			if err := copyFS(hugo, "hugo", pkg); err != nil {
				return err
			}
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
