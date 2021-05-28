package token

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestLines(t *testing.T) {
	type args struct {
		filename string
		buf      string
		trim     bool
	}
	tests := []struct {
		name string
		args args
		want []String
	}{
		{
			name: "single line",
			args: args{
				filename: "test.txt",
				buf:      "hello world",
				trim:     false,
			},
			want: []String{
				newStr("hello world", "test.txt", 0, 1, 1, 11, 1, 12),
			},
		},
		{
			name: "double line",
			args: args{
				filename: "test.txt",
				buf:      "a\nb",
				trim:     false,
			},
			want: []String{
				newStr("a", "test.txt", 0, 1, 1, 1, 1, 2),
				newStr("b", "test.txt", 2, 2, 1, 3, 2, 2),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lines(tt.args.filename, bytes.NewBuffer([]byte(tt.args.buf)), tt.args.trim)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lines() = %v, want %v", deepString(got), deepString(tt.want))
			}
		})
	}
}

func deepString(i interface{}) string {
	buf, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func newStr(val, file string, beginOffset, beginLine, beginCol int, endOffset, endLine, endCol int) String {
	return String{
		Position: Position{
			BeginPos: Pos{
				File:   file,
				Offset: beginOffset,
				Line:   beginLine,
				Col:    beginCol,
			},
			EndPos: Pos{
				File:   file,
				Offset: endOffset,
				Line:   endLine,
				Col:    endCol,
			},
		},
		Val: val,
	}
}
