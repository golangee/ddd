package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type mdFile struct {
	relName  string
	filename string
	md       *Markdown
}

type genctx struct {
	spec    *ddd.AppSpec
	mdFiles []mdFile
}

// markdown returns or creates a new named markdown file
func (g *genctx) markdown(filename string) *Markdown {
	for _, file := range g.mdFiles {
		if file.relName == filename {
			return file.md
		}
	}

	return g.newMarkdown(filename)
}

func (g *genctx) newMarkdown(filename string) *Markdown {
	f := mdFile{
		relName:  filename,
		filename: filename,
		md:       NewMarkdown(),
	}
	g.mdFiles = append(g.mdFiles, f)

	return f.md
}

func (g *genctx) emit(targetDir string) error {
	for _, f := range g.mdFiles {
		fname := filepath.Join(targetDir, f.filename)
		_ = os.MkdirAll(filepath.Dir(fname), os.ModePerm)
		log.Printf("write: %s\n", f.filename)
		if err := ioutil.WriteFile(fname, []byte(f.md.String()), os.ModePerm); err != nil {
			return err
		}

		if err := f.md.EmitGraphics(filepath.Dir(fname)); err != nil {
			return err
		}

	}

	return nil
}
