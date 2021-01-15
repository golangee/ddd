package objn

import "reflect"

// Walk recursively walks over all nodes until everything has been
// visited or an error is returned by f.
func Walk(root Node, f func(path []Node) error) error {
	return walk([]Node{root}, f)
}

// DocFromPath returns the last Doc or nil.
func DocFromPath(path []Node) Doc {
	for i := len(path) - 1; i >= 0; i-- {
		if d, ok := path[i].(Doc); ok {
			return d
		}
	}

	return nil
}

// Collect returns all those paths which matches the predicate.
func Collect(root Node, predicate func(path []Node) (bool, error)) ([][]Node, error) {
	var res [][]Node
	err := Walk(root, func(path []Node) error {
		pred, err := predicate(path)
		if err != nil {
			return err
		}

		if pred {
			res = append(res, path)
		}

		return nil
	})

	return res, err
}

func walk(path []Node, f func(path []Node) error) error {
	if len(path) == 0 {
		return nil
	}

	if err := f(path); err != nil {
		return err
	}

	leaf := path[len(path)-1]

	switch n := leaf.(type) {
	case Pkg:
		for _, lit := range n.Names() {
			if err := walk(addPath(path, n.Get(lit)), f); err != nil {
				return err
			}
		}

		return nil
	case Doc:
		if err := walk(addPath(path, n.Root()), f); err != nil {
			return err
		}

		return nil
	case Map:
		for _, lit := range n.Names() {
			if err := walk(addPath(path, n.Get(lit.String())), f); err != nil {
				return err
			}
		}

		return nil

	case Seq:
		for i := 0; i < n.Count(); i++ {
			if err := walk(addPath(path, n.Get(i)), f); err != nil {
				return err
			}
		}

		return nil
	case Lit:
		return nil
	case nil:
		return nil
	default:
		panic("not implemented: " + reflect.TypeOf(leaf).String())
	}

}

func addPath(path []Node, leaf Node) []Node {
	t := append([]Node{}, path...)
	t = append(t, leaf)
	return t
}
