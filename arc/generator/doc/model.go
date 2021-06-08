package doc

import "sort"

type Node interface {
	assertNode()
}

type Comment struct {
	Value string
}

func NewComment(t string) *Comment {
	return &Comment{Value: t}
}

func (t *Comment) assertNode() {
}

type Text struct {
	Text string
}

func NewText(t string) *Text {
	return &Text{Text: t}
}

func (t *Text) assertNode() {
}

// Element is a generic nested node, just an oversimplification of xml. We use it here
// pretend something like a html subset (without css). It must be flexible and generic enough
// at the same time.
type Element struct {
	Name     string
	Attrs    map[string]string
	Children []Node
}

// NewElement allocates a new Element.
func NewElement(name string) *Element {
	return &Element{Name: name, Attrs: map[string]string{}}
}

func (e *Element) assertNode() {
}

// Attributes returns the sorted list of attributes keys.
func (e *Element) Attributes() []string {
	tmp := make([]string, 0, len(e.Attrs))
	for key := range e.Attrs {
		tmp = append(tmp, key)
	}

	sort.Strings(tmp)

	return tmp
}

func (e *Element) SetAttr(key, val string) *Element {
	e.Attrs[key] = val
	return e
}

func (e *Element) Append(c ...Node) *Element {
	e.Children = append(e.Children, c...)
	return e
}

// TextContent concates all text elements recursively.
func (e *Element) TextContent() string {
	tmp := ""
	for _, child := range e.Children {
		switch t := child.(type) {
		case *Text:
			tmp += t.Text
		case *Element:
			tmp += t.TextContent()
		}
	}

	return tmp
}
