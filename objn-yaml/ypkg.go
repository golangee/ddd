package objnyaml

import (
	"fmt"
	"github.com/golangee/architecture/objn"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

type YamlPkg struct {
	pos      objn.Pos
	children map[string]objn.Node
	comment  string
}

func NewYamlPkg(dir string) (*YamlPkg, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to ReadDir: %w", err)
	}

	n := &YamlPkg{
		children: map[string]objn.Node{},
		pos: objn.Pos{
			File: dir,
		},
	}

	// try a doc.txt for package documentation
	buf, err := ioutil.ReadFile(filepath.Join(dir, "doc.txt"))
	if err == nil && len(buf) > 0 {
		n.comment = string(buf)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			childPkg, err := NewYamlPkg(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("unable to create child pkg: '%s': %w", file.Name(), err)
			}

			n.children[file.Name()] = childPkg
		} else {
			fname := strings.ToLower(file.Name())
			if strings.HasSuffix(fname, ".yaml") || strings.HasSuffix(fname, ".yml") {
				doc, err := NewYamlDoc(filepath.Join(dir, file.Name()))
				if err != nil {
					return nil, fmt.Errorf("unable to parse document: '%s': %w", file.Name(), err)
				}

				if err := doc.Validate(); err != nil {
					return nil, fmt.Errorf("invalid yml: %w", err)
				}

				n.children[file.Name()] = doc
			}
		}
	}

	return n, nil
}

func (n *YamlPkg) Pos() objn.Pos {
	return n.pos
}

func (n *YamlPkg) Comment() string {
	return n.comment
}

func (n *YamlPkg) Validate() error {
	for _, node := range n.children {
		if v, ok := node.(validateable); ok {
			if err := v.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (n *YamlPkg) Names() []string {
	tmp := make([]string, 0, len(n.children))
	for k := range n.children {
		tmp = append(tmp, k)
	}

	sort.Strings(tmp)

	return tmp
}

func (n *YamlPkg) Get(key string) objn.Node {
	return n.children[key]
}
