package yast

import (
	"reflect"
	"strconv"
)

type Predicate = func(n Node) (bool, error)

// Walk recursively walks over all nodes until everything has been
// visited or an error is returned by f.
func Walk(root Node, f func(node Node) error) error {
	if root == nil {
		return nil
	}

	if err := f(root); err != nil {
		return err
	}

	if pnode, ok := root.(Parent); ok {
		for _, node := range pnode.Children() {
			if err := Walk(node, f); err != nil {
				return err
			}
		}
	}

	return nil
}

// ByStereotype is a predicate and returns a function which always returns true if the given stereotype matches.
func ByStereotype(stereotype Stereotype) Predicate {
	return func(n Node) (bool, error) {
		return n.Stereotype() == stereotype, nil
	}
}

func ByKey(key string) Predicate {
	return func(n Node) (bool, error) {
		if x, ok := n.(*Obj); ok {
			if x.Get(key) != nil {
				return true, nil
			}
		}
		return false, nil
	}
}

// Root walks up the parent relation and returns that node whose parent is nil.
func Root(node Node) Node {
	if node == nil {
		return nil
	}

	for node.Parent() != nil {
		node = node.Parent()
	}

	return node
}

// First returns the first node from the slice or nil if empty.
func First(nodes []Node, err error) (Node, error) {
	if len(nodes) == 0 {
		return nil, err
	}

	return nodes[0], err
}

// Filter returns only all those nodes which matches the predicate.
func Filter(root Node, predicate Predicate) ([]Node, error) {
	var res []Node
	err := Walk(root, func(n Node) error {
		pred, err := predicate(n)
		if err != nil {
			return err
		}

		if pred {
			res = append(res, n)
		}

		return nil
	})

	return res, err
}

// SelectParent walks up the node hierarchy and returns the first node which conforms to the predicate. If no such
// node is found, returns nil.
func SelectParent(node Node, p Predicate) (Node, error) {
	if node == nil {
		return nil, nil
	}

	pred, err := p(node)
	if err != nil {
		return nil, err
	}

	if pred {
		return node, nil
	}

	return SelectParent(node.Parent(), p)
}

// ParentName returns the name of the given node in the context of its parent. If the parent is a Seq the iota of
// the index is returned. If there is no parent, the empty string is returned.
func ParentName(node Node) (string, error) {
	p := node.Parent()
	if p == nil {
		return "", NewPosError(node, "expected a parent, but is the root")
	}
	switch t := p.(type) {
	case *Obj:
		for _, attr := range t.Attrs {
			if attr.Value == node {
				return attr.Key.Value, nil
			}
		}
	case *Seq:
		for i, value := range t.Values {
			if value == node {
				return strconv.Itoa(i), nil
			}
		}
	default:
		return "", NewPosError(node, "undefined parent relation: "+reflect.TypeOf(t).String())
	}

	return "", NewPosError(node, "no parent-child match found")
}
