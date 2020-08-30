package golang

import (
	"fmt"
	"github.com/golangee/plantuml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const tocMarker = "[[_TOC_]]"

type Markdown struct {
	sb  *strings.Builder
	toc *strings.Builder
	uml map[string]*plantuml.Diagram
}

func NewMarkdown() *Markdown {
	return &Markdown{sb: &strings.Builder{}, toc: &strings.Builder{}, uml: map[string]*plantuml.Diagram{}}
}

func (m *Markdown) Print(s string) *Markdown {
	m.sb.WriteString(s)
	return m
}

func (m *Markdown) Println(s string) *Markdown {
	m.sb.WriteString(s)
	m.sb.WriteString("\n")
	return m
}

func (m *Markdown) Printf(format string, args ...interface{}) *Markdown {
	m.sb.WriteString(fmt.Sprintf(format, args...))
	return m
}

func (m *Markdown) TOC() *Markdown {
	m.P(tocMarker)
	return m
}

func (m *Markdown) Header(level int, h string) *Markdown {
	for i := 0; i < level; i++ {
		m.Print("#")
	}
	m.Print(" ")
	m.Print(h)
	m.Print("\n\n")

	for i := 0; i < (level-1)*2; i++ {
		m.toc.WriteString(" ")
	}
	m.toc.WriteString("* ")
	m.toc.WriteString(fmt.Sprintf("[%s](%s)\n", h, m.escapeAnchorTitle(h)))
	return m
}

func (m *Markdown) H1(h string) *Markdown {
	return m.Header(1, h)
}

func (m *Markdown) P(str string) *Markdown {
	m.Print(str)
	m.Print("\n\n")
	return m
}

func (m *Markdown) UML(title string) *plantuml.Diagram {
	fname := "uml-" + m.escapeAnchorTitle(title) + ".gen.svg"
	m.Printf("![%s](%s?raw=true)\n\n", title, fname)
	d := plantuml.NewDiagram().Include(plantuml.ThemeCerulean)
	m.uml[fname] = d

	return d
}

func (m *Markdown) EmitGraphics(dir string) error {
	for fname, diagram := range m.uml {
		exec.Command("plantuml")
		filename := filepath.Join(dir, fname)
		buf, err := plantuml.RenderLocal("svg", diagram)
		if err != nil {
			return err
		}

		log.Printf("write: %s\n", filename)
		if err := ioutil.WriteFile(filename, buf, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func (m *Markdown) H2(h string) *Markdown {
	return m.Header(2, h)
}

func (m *Markdown) H3(h string) *Markdown {
	return m.Header(3, h)
}

func (m *Markdown) H4(h string) *Markdown {
	return m.Header(4, h)
}

func (m *Markdown) H5(h string) *Markdown {
	return m.Header(5, h)
}

func (m *Markdown) String() string {
	buf := m.sb.String()
	return strings.ReplaceAll(buf, tocMarker, m.toc.String())
}

func (m *Markdown) escapeAnchorTitle(title string) string {
	buf := &strings.Builder{}
	for _, r := range strings.ReplaceAll(strings.ToLower(title), " ", "-") {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
