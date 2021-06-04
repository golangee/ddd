package golang

import (
	"github.com/golangee/architecture/arc/adl"
	"strings"
)

func makePreamble(p adl.Preamble) string {
	tmp := ""
	if p.Generator != "" {
		tmp = p.Generator
	}

	if tmp != "" && p.License != "" {
		tmp += "\n\n"
	}

	if p.License != "" {
		tmp += p.License
	}

	return tmp
}

func makeEscapedPreamble(p adl.Preamble, escapeWith string) string {
	tmp := makePreamble(p)
	buf := strings.Builder{}
	for _, s := range strings.Split(tmp, "\n") {
		buf.WriteString(escapeWith)
		buf.WriteString(s)
		buf.WriteString("\n")
	}

	return buf.String()
}
