package objnjson

import (
	"encoding/json"
	"github.com/golangee/architecture/objn"
	"reflect"
)

func Debug(node objn.Node) string {
	tmp := debug(node)
	buf, err := json.MarshalIndent(tmp, " ", " ")
	if err != nil {
		panic(err)
	}

	return string(buf)
}

func debug(node objn.Node) interface{} {
	switch n := node.(type) {
	case objn.Pkg:
		obj := map[string]interface{}{}
		//	obj[".pos"] = debugPos(n.Pos())
		//	obj[".type"] = "pkg"
		for _, lit := range n.Names() {
			obj[lit] = debug(n.Get(lit))
		}

		return obj

	case objn.Doc:
		obj := map[string]interface{}{}
		//obj[".pos"] = debugPos(n.Pos())
		//	obj[".type"] = "doc"
		obj[".root"] = debug(n.Root())

		return obj
	case objn.Map:
		obj := map[string]interface{}{}
		//obj[".pos"] = debugPos(n.Pos())
		//	obj[".type"] = "map"
		for _, lit := range n.Names() {
			obj[lit.String()] = debug(n.Get(lit.String()))
		}

		return obj

	case objn.Seq:
		arr := make([]interface{}, 0, n.Count())
		for i := 0; i < n.Count(); i++ {
			arr = append(arr, debug(n.Get(i)))
		}

		return arr
	case objn.Lit:
		return n.String()
	case nil:
		return nil
	default:
		panic("not implemented: " + reflect.TypeOf(node).String())
	}
}

func debugPos(pos objn.Pos) map[string]interface{} {
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
