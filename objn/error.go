package objn

import (
	"fmt"
	"strconv"
	"strings"
)

type PosErrorDetail struct {
	pos      Pos
	lineText string
	message  string
	markText string // if not empty, overrides the position by the given found index
}

func NewPosErrorDetailFromDoc(d Doc, pos Pos, msg string) PosErrorDetail {
	lines := strings.Split(d.String(), "\n")
	no := pos.Line - 1

	if no > len(lines) {
		no = len(lines) - 1
	}

	ltext := ""
	if no < len(lines) && no >= 0 {
		ltext = lines[no]
	}

	return PosErrorDetail{
		pos:      pos,
		lineText: ltext,
		message:  msg,
	}
}

// PosError represents a very specific positional error with a lot of explaining noise.
type PosError struct {
	pos     Pos
	message string
	cause   error
	details []PosErrorDetail
}

func NewPosError(path []Node, pos Node, msg string, details ...PosErrorDetail) PosError {
	return NewPosErrorMark(path, pos, msg, "", details...)
}

func NewPosErrorMark(path []Node, pos Node, msg, mark string, details ...PosErrorDetail) PosError {
	err := PosError{
		pos:     pos.Pos(),
		message: msg,
	}

	if len(details) == 0 {
		doc := DocFromPath(path)
		if doc != nil {
			err.details = append(err.details, NewPosErrorDetailFromDoc(doc, pos.Pos(), msg))
			err.details[len(err.details)-1].markText = mark
		}
	} else {
		err.details = append(err.details, details...)
	}

	return err
}

func (p PosError) Unwrap() error {
	return p.cause
}

func (p PosError) Error() string {
	return p.message
}

// Explain returns a multi-line text suited to be printed into the console.
func (p PosError) Explain() string {
	sb := &strings.Builder{}
	sb.WriteString("error: ")
	sb.WriteString(p.message)
	sb.WriteString("\n")
	sb.WriteString("--> ")
	sb.WriteString(p.pos.String())
	sb.WriteString("\n")

	indent := 0
	for _, detail := range p.details {
		l := len(strconv.Itoa(detail.pos.Line))
		if l > indent {
			indent = l
		}
	}
	for i, detail := range p.details {
		if detail.pos.File != p.pos.File {
			sb.WriteString(p.pos.String())
			sb.WriteString("\n")
		}

		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s|\n", ""))
		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"d|", detail.pos.Line))
		sb.WriteString(detail.lineText)
		sb.WriteString("\n")

		markIdx := -1
		if detail.markText != "" {
			markIdx = strings.Index(detail.lineText, detail.markText)
		}

		sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s|", ""))
		if markIdx == -1 {
			sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(detail.pos.Col-1)+"s", ""))
			sb.WriteString("^~~~ ")
		} else {
			sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(markIdx)+"s", ""))
			for i := 0; i < len(detail.markText); i++ {
				sb.WriteRune('^')
			}
			sb.WriteRune(' ')
		}

		sb.WriteString(detail.message)
		sb.WriteString("\n")

		if i < len(p.details)-1 {
			sb.WriteString(fmt.Sprintf("%"+strconv.Itoa(indent)+"s\n", "..."))
		}
	}

	return sb.String()
}
