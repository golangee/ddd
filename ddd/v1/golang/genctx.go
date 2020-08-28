package golang

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/reflectplus/mod"
	"github.com/golangee/src"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type genctx struct {
	mod   mod.Modules
	spec  *ddd.AppSpec
	files []genfile
}

type genfile struct {
	filename    string
	path        string
	packagename string
	file        *src.FileBuilder
}

func (g *genctx) newFile(path, fname, pkgname string) *src.FileBuilder {
	if pkgname == "" {
		pkgname = filepath.Base(strings.ToLower(path))
	}

	f := genfile{
		filename:    strings.ToLower(fname + ".gen.go"),
		path:        filepath.Join(g.mod.Main().Dir, strings.ToLower(path)),
		packagename: pkgname,
	}

	f.file = src.NewFile(f.packagename)
	f.file.SetGeneratorName("golangee/architecture")
	f.file.SetImportPath(g.mod.Main().Path + "/" + strings.ToLower(path))
	g.files = append(g.files, f)

	return f.file
}

func (g *genctx) emit() error {
	for _, f := range g.files {
		_ = os.MkdirAll(f.path, os.ModePerm)
		fname := filepath.Join(f.path, f.filename)

		w := &src.BufferedWriter{}
		f.file.Emit(w)
		str, err := w.Format()
		if err != nil {
			return fmt.Errorf("%s\n%w", str, err)
		}

		log.Printf("write: %s\n", fname)
		if err := ioutil.WriteFile(fname, []byte(str), os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
