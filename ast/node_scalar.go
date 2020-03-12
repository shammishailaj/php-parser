package ast

// ScalarDnumber node
type ScalarDnumber struct {
	Node
	Value string
}

func (n *ScalarDnumber) Children() map[string][]Vertex {
	return nil
}

func (n *ScalarDnumber) Visit(v Visitor, depth int) bool {
	return v.ScalarDnumber(depth, n)
}

// ScalarEncapsed node
type ScalarEncapsed struct {
	Node
	Parts []Vertex
}

func (n *ScalarEncapsed) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Parts": n.Parts,
	}
}

func (n *ScalarEncapsed) Visit(v Visitor, depth int) bool {
	return v.ScalarEncapsed(depth, n)
}

// ScalarEncapsedStringPart node
type ScalarEncapsedStringPart struct {
	Node
	Value string
}

func (n *ScalarEncapsedStringPart) Children() map[string][]Vertex {
	return nil
}

func (n *ScalarEncapsedStringPart) Visit(v Visitor, depth int) bool {
	return v.ScalarEncapsedStringPart(depth, n)
}

// ScalarHeredoc node
type ScalarHeredoc struct {
	Node
	Label string
	Parts []Vertex
}

func (n *ScalarHeredoc) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Parts": n.Parts,
	}
}

func (n *ScalarHeredoc) Visit(v Visitor, depth int) bool {
	return v.ScalarHeredoc(depth, n)
}

// ScalarLnumber node
type ScalarLnumber struct {
	Node
	Value string
}

func (n *ScalarLnumber) Children() map[string][]Vertex {
	return nil
}

func (n *ScalarLnumber) Visit(v Visitor, depth int) bool {
	return v.ScalarLnumber(depth, n)
}

// ScalarMagicConstant node
type ScalarMagicConstant struct {
	Node
	Value string
}

func (n *ScalarMagicConstant) Children() map[string][]Vertex {
	return nil
}

func (n *ScalarMagicConstant) Visit(v Visitor, depth int) bool {
	return v.ScalarMagicConstant(depth, n)
}

// ScalarString node
type ScalarString struct {
	Node
	Value string
}

func (n *ScalarString) Children() map[string][]Vertex {
	return nil
}

func (n *ScalarString) Visit(v Visitor, depth int) bool {
	return v.ScalarString(depth, n)
}
