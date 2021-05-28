package token

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode/utf8"
)

// A String holds a positional value.
type String struct {
	Position
	Val string
}

// NewString just creates a new string.
func NewString(val string) String {
	return String{Val: val}
}

// Lines parses the given bytes and returns the correct located string for each line.
func Lines(filename string, r io.ByteReader) ([]String, error) {
	var tmp bytes.Buffer
	beginPos := Pos{
		File:   filename,
		Offset: 0,
		Line:   1,
		Col:    1,
	}

	endPos := Pos{
		File:   filename,
		Offset: 0,
		Line:   1,
		Col:    1,
	}

	var res []String

	for {
		b, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		endPos.Offset++
		endPos.Col++

		if b == '\n' {
			str := String{
				Position: Position{
					BeginPos: beginPos,
					EndPos:   endPos,
				},
				Val: tmp.String(),
			}

			// remove newline
			str.EndPos.Col--
			str.EndPos.Offset--

			res = append(res, str)

			tmp.Reset()
			endPos.Line++
			endPos.Col = 1
			beginPos = endPos
		} else {
			if err := tmp.WriteByte(b); err != nil {
				return nil, err
			}
		}

	}

	if tmp.Len() > 0 {
		res = append(res, String{
			Position: Position{
				BeginPos: beginPos,
				EndPos:   endPos,
			},
			Val: tmp.String(),
		})
	}

	return res, nil
}

// Value of this string.
func (s String) Value() string {
	return s.Val
}

// String just returns the value.
func (s String) String() string {
	return s.Value()
}

// GoString returns a positional information with the string.
func (s String) GoString() string {
	return s.BeginPos.String() + ":" + s.Value()
}

// Locate creates a new string with the applied position, assuming this string is ever in a single line.
func (s String) Locate(filename string, offset, line, col int) String {
	s.BeginPos.File = filename
	s.BeginPos.Line = line
	s.BeginPos.Col = col
	s.BeginPos.Offset = offset

	s.EndPos.File = filename
	s.EndPos.Line = line
	s.EndPos.Col = s.BeginPos.Col + utf8.RuneCountInString(s.Val)
	s.EndPos.Offset = s.BeginPos.Offset + len(s.Val)

	return s
}

// TrimSpace returns a new string with the retro-fitted positions.
func (s String) TrimSpace() String {
	v := strings.TrimSpace(s.Val)
	idx := strings.Index(s.Val, v)
	str := s
	str.Val = v
	str.BeginPos.Col += utf8.RuneCountInString(s.Val[:idx])
	str.BeginPos.Offset += idx

	str.EndPos.Col = str.BeginPos.Col + utf8.RuneCountInString(v)
	str.EndPos.Offset = str.BeginPos.Offset + len(v)

	return str
}
