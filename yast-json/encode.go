package yastjson

import (
	"encoding/json"
	"github.com/golangee/architecture/yast"
	"reflect"
)

// Options influences the rendering or conversion.
type Options struct {
	Beautify      bool     // if true, use MarshalIndent
	PosKey        string   // if not empty, put a position Object behind that key. Sequences will be wrapped in another Obj.
	TypeKey       string   // if not empty, put a string type behind that key. Sequences will be wrapped in another Obj.
	SeqKey        string   // if a sequence is wrapped, use that key. By default the empty key.
	StereotypeKey string   // if not empty, put the stereotype behind it.
	OmitKeys      []string // obj key names, which should be ignored.
}

// Marshal takes the given node and serializes it into json.
func Marshal(node yast.Node, opts Options) ([]byte, error) {
	tmp := Convert(node, opts)
	if opts.Beautify {
		return json.MarshalIndent(tmp, " ", " ")
	}

	return json.Marshal(tmp)
}

// Convert allocates new stdlib data structures to represent the given node as json.
func Convert(node yast.Node, opts Options) interface{} {
	switch n := node.(type) {
	case *yast.Obj:
		obj := map[string]interface{}{}
		if opts.PosKey != "" {
			obj[opts.PosKey] = newPosMap(n.Pos())
		}

		if opts.TypeKey != "" {
			obj[opts.TypeKey] = "obj"
		}

		if opts.StereotypeKey != "" {
			if n.Stereotype() != "" {
				obj[opts.StereotypeKey] = n.Stereotype()
			}
		}

		for _, lit := range n.Names() {
			ignore := false
			for _, key := range opts.OmitKeys {
				if key == lit {
					ignore = true
					break
				}
			}

			if !ignore {
				obj[lit] = Convert(n.Get(lit), opts)
			}
		}

		return obj

	case *yast.Seq:
		seq := make([]interface{}, 0, len(n.Values))
		for _, value := range n.Values {
			seq = append(seq, Convert(value, opts))
		}

		// straight forward json conversion
		if opts.PosKey == "" && opts.TypeKey == "" && opts.StereotypeKey == "" {
			return seq
		}

		// otherwise perform map boxing
		obj := map[string]interface{}{}
		if opts.PosKey != "" {
			obj[opts.PosKey] = newPosMap(n.Pos())
		}

		if opts.TypeKey != "" {
			obj[opts.TypeKey] = "seq"
		}

		if opts.StereotypeKey != "" {
			if n.Stereotype() != "" {
				obj[opts.StereotypeKey] = n.Stereotype()
			}
		}

		obj[opts.SeqKey] = seq

		return obj
	case *yast.Str:
		return n.Value
	case *yast.Null:
		return nil
	case *yast.Int:
		return n.Value
	case *yast.Bool:
		return n.Value
	case *yast.Float:
		return n.Value
	case nil:
		return nil
	default:
		panic("not implemented: " + reflect.TypeOf(node).String())
	}
}

func newPosMap(pos yast.Pos) map[string]interface{} {
	obj := map[string]interface{}{}
	if pos.File != "" {
		obj["file"] = pos.File
	}

	if pos.Line != 0 {
		obj["line"] = pos.Line
	}

	if pos.Col != 0 {
		obj["col"] = pos.Col
	}

	return obj
}

// Sprint just serializes the given node into a human readable string.
func Sprint(node yast.Node) string {
	buf, err := Marshal(node, Options{
		Beautify:      true,
		OmitKeys:      []string{"src"},
		StereotypeKey: ".stereotype",
	})

	if err != nil {
		return err.Error()
	}

	return string(buf)
}
