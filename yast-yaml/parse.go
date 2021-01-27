package yastyaml

import (
	"fmt"
	"github.com/golangee/architecture/yast"
	"github.com/golangee/architecture/yast/validate"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// ParseDir parses all .yml/.yaml files from the given path and returns them in a nested hierarchy.
func ParseDir(path string) (*yast.Obj, error) {
	obj, err := newYamlPkg(nil, path)
	if err != nil {
		return nil, err
	}

	if err := validate.Recursive(validate.NoDuplicateKeys())(obj); err != nil {
		return obj, err
	}

	return obj, nil
}

func newYamlPkg(parent yast.Node, dir string) (*yast.Obj, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to ReadDir: %w", err)
	}

	n := &yast.Obj{
		ObjParent:     parent,
		ObjStereotype: yast.Package,
		ObjPos:        yast.Pos{File: dir},
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			childPkg, err := newYamlPkg(n, filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("unable to create child pkg: '%s': %w", file.Name(), err)
			}

			attr := &yast.Attr{
				AttrStereotype: yast.Directory,
				AttrParent:     n,
			}

			attr.Key = &yast.Str{
				ValuePos:    yast.Pos{File: filepath.Join(dir, file.Name())},
				ValueParent: attr,
				Value:       file.Name(),
			}

			attr.Value = childPkg

			n.Attrs = append(n.Attrs, attr)

		} else {
			fname := strings.ToLower(file.Name())
			if strings.HasSuffix(fname, ".yaml") || strings.HasSuffix(fname, ".yml") {
				doc, err := parseYaml(n, filepath.Join(dir, file.Name()))
				if err != nil {
					return nil, fmt.Errorf("unable to parse document: '%s': %w", file.Name(), err)
				}

				attr := &yast.Attr{
					AttrStereotype: yast.File,
					AttrParent:     n,
				}

				attr.Key = &yast.Str{
					ValuePos:    yast.Pos{File: filepath.Join(dir, file.Name())},
					ValueParent: attr,
					Value:       file.Name(),
				}

				attr.Value = doc

				n.Attrs = append(n.Attrs, attr)

			}
		}
	}

	return n, nil
}

func parseYaml(parent yast.Node, fname string) (yast.Node, error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	node := &yaml.Node{}
	err = yaml.Unmarshal(buf, node)
	if err != nil {
		return nil, fmt.Errorf("unable to parse yaml file: %w", err)
	}

	if len(buf) > 0 {
		if node.Kind != yaml.DocumentNode {
			return nil, fmt.Errorf("invalid node kind: " + strconv.Itoa(int(node.Kind)))
		}
	}

	if len(node.Content) > 1 {
		return nil, fmt.Errorf("invalid content length")
	}

	n := doc2yast(parent, fname, node)
	attr := &yast.Attr{AttrParent: n}

	attr.Key = &yast.Str{
		ValuePos:    posFromNode(fname, node),
		ValueEnd:    yast.Pos{},
		ValueParent: attr,
		Value:       "src", // src is a magic document value
	}

	attr.Value = &yast.Str{
		ValuePos:    posFromNode(fname, node),
		ValueEnd:    yast.Pos{},
		ValueParent: attr,
		Value:       string(buf),
	}

	n.Attrs = append(n.Attrs, attr)

	return n, nil
}

func yamlNode2yast(parent yast.Node, fname string, yn *yaml.Node) yast.Node {
	switch yn.Kind {
	case yaml.MappingNode:
		return map2yast(parent, fname, yn)
	case yaml.SequenceNode:
		return seq2yast(parent, fname, yn)
	case yaml.ScalarNode:
		switch yn.Tag {
		case "!!str":
			return str2yast(parent, fname, yn)
		case "!!null":
			return null2yast(parent, fname, yn)
		case "!!int":
			return int2yast(parent, fname, yn)
		default:
			panic("scalar node Tag not implemented: " + yn.Tag)
		}
	default:
		panic("node type not implemented: " + strconv.Itoa(int(yn.Kind)))
	}
}

func doc2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Obj {
	doc := &yast.Obj{
		ObjStereotype: yast.Document,
		ObjPos:        posFromNode(fname, yn),
		ObjParent:     parent,
	}

	if len(yn.Content) == 0 {
		return doc
	}

	root := yamlNode2yast(doc, fname, yn.Content[0])

	attr := &yast.Attr{AttrParent: doc}

	attr.Key = &yast.Str{
		ValuePos:    posFromNode(fname, yn),
		ValueEnd:    yast.Pos{},
		ValueParent: attr,
		Value:       "root", // root is a magic document value
	}

	attr.Value = root
	doc.Attrs = append(doc.Attrs, attr)

	return doc
}

func map2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Obj {
	obj := &yast.Obj{
		ObjPos:    posFromNode(fname, yn),
		ObjParent: parent,
	}

	for i := 0; i < len(yn.Content); i += 2 {
		nodeKey := yn.Content[i]
		nodeVal := yn.Content[i+1]

		attr := &yast.Attr{AttrParent: obj}
		attr.Key = str2yast(attr, fname, nodeKey)
		attr.Value = yamlNode2yast(attr, fname, nodeVal)

		obj.Attrs = append(obj.Attrs, attr)
	}

	return obj
}

func seq2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Seq {
	seq := &yast.Seq{
		SeqPos:    posFromNode(fname, yn),
		SeqParent: parent,
	}

	for _, node := range yn.Content {
		seq.Values = append(seq.Values, yamlNode2yast(seq, fname, node))
	}

	return seq
}

func str2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Str {
	str := &yast.Str{
		ValuePos:    posFromNode(fname, yn),
		ValueParent: parent,
		Value:       yn.Value,
	}

	str.ValueEnd = str.ValuePos
	str.ValueEnd.Col += len(str.Value)

	return str
}

func int2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Int {
	num, err := strconv.ParseInt(yn.Value, 10, 64)
	if err != nil {
		panic(fmt.Errorf("illegal state: unparseable int in int-node: %w", err))
	}

	str := &yast.Int{
		ValuePos:    posFromNode(fname, yn),
		ValueParent: parent,
		Value:       num,
	}

	str.ValueEnd = str.ValuePos
	str.ValueEnd.Col += len(strconv.Itoa(int(num)))

	return str
}

func null2yast(parent yast.Node, fname string, yn *yaml.Node) *yast.Null {
	return &yast.Null{
		ValuePos:    posFromNode(fname, yn),
		ValueParent: parent,
	}
}

func posFromNode(fname string, n *yaml.Node) yast.Pos {
	return yast.Pos{
		File: fname,
		Line: n.Line,
		Col:  n.Column,
	}
}
