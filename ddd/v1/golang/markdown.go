package golang

import (
	"fmt"
	"strings"
)

type Markdown struct {
	sb *strings.Builder
}

func NewMarkdown() *Markdown {
	return &Markdown{sb: &strings.Builder{}}
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

func (m *Markdown) Header(level int, h string) *Markdown {
	for i := 0; i < level; i++ {
		m.Print("#")
	}
	m.Print(" ")
	m.Print(h)
	m.Print("\n\n")
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
	return m.sb.String()
}
