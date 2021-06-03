package golang

import "github.com/golangee/architecture/arc/adl"

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
