package aclyaml

import (
	"fmt"
	"github.com/golangee/architecture/acl"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

type YamlPkg struct {
	pos      acl.Pos
	children map[string]acl.Node
}

func NewYamlPkg(dir string) (*YamlPkg, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to ReadDir: %w", err)
	}

	n := &YamlPkg{
		children: map[string]acl.Node{},
		pos: acl.Pos{
			File: dir,
		},
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

func (n *YamlPkg) Pos() acl.Pos {
	return n.pos
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

func (n *YamlPkg) Names() []acl.Lit {
	tmp := make([]acl.Lit, 0, len(n.children))
	for k := range n.children {
		tmp = append(tmp, n.Name(k))
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].String() < tmp[j].String()
	})

	return tmp
}

func (n *YamlPkg) Name(key string) acl.Lit {
	return YamlLit{
		pos: acl.Pos{File: filepath.Join(n.pos.File, key)},
		val: key,
	}
}

func (n *YamlPkg) Get(key string) acl.Node {
	return n.children[key]
}
