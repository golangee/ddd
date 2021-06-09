package markdown

import (
	"bytes"
	"github.com/golangee/architecture/arc/generator/doc"
	"strconv"
)

const fakeComment = true

// Render tries its best to convert html-like Elements into markdown text. Unsupported nodes
// are mostly ignored or have just their text content extracted.
func Render(e doc.Node) []byte {
	var tmp bytes.Buffer
	render(e, &tmp)
	return tmp.Bytes()
}

func render(src doc.Node, dst *bytes.Buffer) {
	if t, ok := src.(*doc.Text); ok {
		dst.WriteString(t.Text)
		return
	}

	if t, ok := src.(*doc.Comment); ok {
		if fakeComment {
			//[//]: # (Code generated by golangee/eearc; DO NOT EDIT.)
			dst.WriteString("\n\n[//]: # (")
			dst.WriteString(t.Value)
			dst.WriteString(")\n\n")
		} else {
			// <!-- goldmark is crashing when used at the beginning of a file

			dst.WriteString("\n<!-- ")
			dst.WriteString(t.Value)
			dst.WriteString(" -->\n\n")
		}

		return
	}

	e := src.(*doc.Element)

	switch e.Name {
	case "h1":
		fallthrough
	case "h2":
		fallthrough
	case "h3":
		fallthrough
	case "h4":
		fallthrough
	case "h5":
		fallthrough
	case "h6":
		writeHeading(e, dst)
	default:
		for _, child := range e.Children {
			render(child, dst)
		}
	}

}

func writeHeading(e *doc.Element, dst *bytes.Buffer) {
	level, err := strconv.Atoi(e.Name[1:])
	if err != nil {
		panic(err)
	}

	for i := 0; i < level; i++ {
		dst.WriteString("#")
	}

	dst.WriteString(" ")
	dst.WriteString(e.TextContent())
	dst.WriteString("\n")
}
