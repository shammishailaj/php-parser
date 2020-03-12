package ast

// ExprCast node
type ExprCast struct {
	Node
	Expr Vertex
}

func (n *ExprCast) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

// ExprCastArray node
type ExprCastArray struct {
	ExprCast
}

func (n *ExprCastArray) Visit(v Visitor, depth int) bool {
	return v.ExprCastArray(depth, n)
}

// ExprCastBool node
type ExprCastBool struct {
	ExprCast
}

func (n *ExprCastBool) Visit(v Visitor, depth int) bool {
	return v.ExprCastBool(depth, n)
}

// ExprCastDouble node
type ExprCastDouble struct {
	ExprCast
}

func (n *ExprCastDouble) Visit(v Visitor, depth int) bool {
	return v.ExprCastDouble(depth, n)
}

// ExprCastInt node
type ExprCastInt struct {
	ExprCast
}

func (n *ExprCastInt) Visit(v Visitor, depth int) bool {
	return v.ExprCastInt(depth, n)
}

// ExprCastObject node
type ExprCastObject struct {
	ExprCast
}

func (n *ExprCastObject) Visit(v Visitor, depth int) bool {
	return v.ExprCastObject(depth, n)
}

// ExprCastString node
type ExprCastString struct {
	ExprCast
}

func (n *ExprCastString) Visit(v Visitor, depth int) bool {
	return v.ExprCastString(depth, n)
}

// ExprCastUnset node
type ExprCastUnset struct {
	ExprCast
}

func (n *ExprCastUnset) Visit(v Visitor, depth int) bool {
	return v.ExprCastUnset(depth, n)
}
