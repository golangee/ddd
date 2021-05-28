package token

import (
	"bytes"
	"errors"
	"io"
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

// Lines parses the given bytes, optionally trims each line and creates the according strings for each line.
func Lines(filename string, r io.ByteReader, trim bool) ([]String, error) {
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
