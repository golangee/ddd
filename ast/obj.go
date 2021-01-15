package ast

// An Attr is a tuple of a named key/value pair (Str -> Node).
type Attr struct {
	AttrStereotype string
	AttrParent     Node
	Key            Str
	Value          Node
}

func (a *Attr) Pos() Pos {
	return a.Key.ValuePos
}

func (a *Attr) End() Pos {
	return a.Value.End()
}

func (a *Attr) Parent() Node {
	return a.AttrParent
}

func (a *Attr) Stereotype() string {
	return a.AttrStereotype
}

// Obj is a Node which contains named attributes (see Attr).
type Obj struct {
	ObjPos        Pos
	ObjEnd        Pos
	ObjParent     Node
	ObjStereotype string
	Attrs         []Attr
}

func (o *Obj) Pos() Pos {
	return o.ObjPos
}

func (o *Obj) End() Pos {
	return o.ObjEnd
}

func (o *Obj) Parent() Node {
	return o.ObjParent
}

func (o *Obj) Stereotype() string {
	return o.ObjStereotype
}

// Get returns the first attribute value associated with the given key.
func (o *Obj) Get(key string) Node {
	for _, attr := range o.Attrs {
		if attr.Key.Value == key {
			return attr.Value
		}
	}

	return nil
}

// Names returns all names in ascending order. May contain duplicates.
func (o *Obj) Names() []string {
	tmp := make([]string, 0, len(o.Attrs))
	for _, attr := range o.Attrs {
		tmp = append(tmp, attr.Key.Value)
	}

	return tmp
}

