package validate

import (
	"fmt"
	"github.com/golangee/architecture/yast"
	"path/filepath"
	"strings"
	"unicode"
)

// YamlIdentifier validates and returns the type name or identifier.
// It enforce that only one upper case obj definition is placed
// in root and that the file is named properly.
func YamlIdentifier(n yast.Node) (string, error) {
	doc, err := yast.SelectParent(n, yast.ByStereotype(yast.Document))
	if err != nil {
		return "", fmt.Errorf("parent must be a document: %w", err)
	}

	name, err := yast.ParentName(doc)
	if err != nil {
		return "", fmt.Errorf("cannot get my name in parent: %w", err)
	}

	idx := strings.LastIndex(name, ".yml")
	if idx < 0 {
		return "", yast.NewPosError(doc, "file name is invalid: '"+name+"'").SetHint("set extension to .yml")
	}

	realName := name[0:idx]
	root, err := ExpectRoot(doc)
	if err != nil {
		return "", fmt.Errorf("invalid document: %w", err)
	}

	if len(root.Attrs) != 1 {
		return "", yast.NewPosError(root, "file must define exactly one upper case identifier")
	}

	if _, ok := root.Attrs[0].Value.(*yast.Obj); ok {
		if err := IsExportedIdentifier(root.Attrs[0].Key.Value); err != nil {
			return "", yast.NewPosError(root.Attrs[0].Key, "invalid identifier").SetHint("must be a public identifier as defined by https://golang.org/ref/spec#Identifiers")
		}

		if root.Attrs[0].Key.Value != realName {
			return "", yast.NewPosError(root.Attrs[0].Key, "unmatched identifier").SetHint("must be named as the file (" + realName + ")")
		}

		return realName, nil
	} else {
		return "", yast.NewPosError(root.Attrs[0].Value, "attribute value must be an Object")
	}

}

// YamlDirName is like YamlName but it proofs the name against the containing directory name.
func YamlDirName(n yast.Node) (string, error) {
	doc, err := yast.SelectParent(n, yast.ByStereotype(yast.Document))
	if err != nil {
		return "", fmt.Errorf("parent must be a document: %w", err)
	}

	// this is a bit hacky, but the actual real filename must be equal to the Pos. For the root package we have no other choice at all
	name := filepath.Base(doc.Parent().Pos().File)

	root, err := ExpectRoot(doc)
	if err != nil {
		return "", fmt.Errorf("invalid document: %w", err)
	}

	if len(root.Attrs) != 1 {
		return "", yast.NewPosError(root, "file must define exactly one upper case identifier")
	}

	if _, ok := root.Attrs[0].Value.(*yast.Obj); ok {
		nm := root.Attrs[0].Key.Value
		if err := IsExportedIdentifier(nm); err != nil {
			return "", yast.NewPosError(root.Attrs[0].Key, "invalid identifier").SetHint("must be a public identifier as defined by https://golang.org/ref/spec#Identifiers")
		}

		if nm != name {
			return "", yast.NewPosError(root.Attrs[0].Key, "'"+nm+"' is an unmatched identifier").SetHint("must be named as the containing directory which is '" + name + "'")
		}

		return name, nil
	} else {
		return "", yast.NewPosError(root.Attrs[0].Value, "attribute value must be an Object")
	}
}

// Expect root returns the root object or fails.
func ExpectRoot(n yast.Node) (*yast.Obj, error) {
	if obj, ok := n.(*yast.Obj); ok {
		if robj, ok := obj.Get("root").(*yast.Obj); ok {
			return robj, nil
		} else {
			return nil, yast.NewPosError(obj, "must contain an object named 'root'").SetHint("ensure that file is not empty and properly indented.")
		}
	} else {
		return nil, yast.NewPosError(n, "expected a document object")
	}
}

func filterChildrenByUppercaseKey(p *yast.Obj) []yast.Node {
	var r []yast.Node
	for _, attr := range p.Attrs {
		if len(attr.Key.Value) > 0 && unicode.IsUpper(getRune(attr.Key.Value, 0)) {
			r = append(r, attr)
		}
	}

	return r
}

func getRune(str string, idx int) rune {
	for i, r := range str {
		if i == idx {
			return r
		}
	}

	panic("rune index out of bounds")
}

// Doc asserts and returns the documentation attribute. This must be a string.
func Doc(n yast.Node) (string, error) {
	if obj, ok := n.(*yast.Obj); ok {
		str, err := ValidComment(obj.Get("doc"))
		if err != nil {
			return "", fmt.Errorf("not a valid comment: %w", err)
		}

		return str, nil
	} else {
		return "", yast.NewPosError(n, "node must be an object and should contain a 'doc' string attribute")
	}
}

// ValidComment asserts that the Node is a yast.Str and fulfills the minimum quality for a comment.
func ValidComment(n yast.Node) (string, error) {
	const hintEmpty = "should look like 'doc: ... is a thing which is required for something.'"
	if str, ok := n.(*yast.Str); ok {
		if strings.HasPrefix(str.Value, "...") && strings.HasSuffix(str.Value, ".") && len(str.Value) > 8 {
			return str.Value, nil
		} else {
			return "", yast.NewPosError(n, "invalid documentation content").SetHint(hintEmpty)
		}
	} else {
		if n == nil {
			return "", yast.NewPosError(n, "must have a 'doc' string attribute").SetHint(hintEmpty)
		}
		return "", yast.NewPosError(n, "must be a string value")
	}
}
