package ast

// ExprAssign node
type ExprAssign struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *ExprAssign) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var":  {n.Var},
		"Expr": {n.Expr},
	}
}

func (n *ExprAssign) Visit(v Visitor, depth int) bool {
	return v.ExprAssign(depth, n)
}

// ExprAssignReference node
type ExprAssignReference struct {
	ExprAssign
}

func (n *ExprAssignReference) Visit(v Visitor, depth int) bool {
	return v.ExprAssignReference(depth, n)
}

// ExprAssignBitwiseAnd node
type ExprAssignBitwiseAnd struct {
	ExprAssign
}

func (n *ExprAssignBitwiseAnd) Visit(v Visitor, depth int) bool {
	return v.ExprAssignBitwiseAnd(depth, n)
}

// ExprAssignBitwiseOr node
type ExprAssignBitwiseOr struct {
	ExprAssign
}

func (n *ExprAssignBitwiseOr) Visit(v Visitor, depth int) bool {
	return v.ExprAssignBitwiseOr(depth, n)
}

// ExprAssignBitwiseXor node
type ExprAssignBitwiseXor struct {
	ExprAssign
}

func (n *ExprAssignBitwiseXor) Visit(v Visitor, depth int) bool {
	return v.ExprAssignBitwiseXor(depth, n)
}

// ExprAssignCoalesce node
type ExprAssignCoalesce struct {
	ExprAssign
}

func (n *ExprAssignCoalesce) Visit(v Visitor, depth int) bool {
	return v.ExprAssignCoalesce(depth, n)
}

// ExprAssignConcat node
type ExprAssignConcat struct {
	ExprAssign
}

func (n *ExprAssignConcat) Visit(v Visitor, depth int) bool {
	return v.ExprAssignConcat(depth, n)
}

// ExprAssignDiv node
type ExprAssignDiv struct {
	ExprAssign
}

func (n *ExprAssignDiv) Visit(v Visitor, depth int) bool {
	return v.ExprAssignDiv(depth, n)
}

// ExprAssignMinus node
type ExprAssignMinus struct {
	ExprAssign
}

func (n *ExprAssignMinus) Visit(v Visitor, depth int) bool {
	return v.ExprAssignMinus(depth, n)
}

// ExprAssignMod node
type ExprAssignMod struct {
	ExprAssign
}

func (n *ExprAssignMod) Visit(v Visitor, depth int) bool {
	return v.ExprAssignMod(depth, n)
}

// ExprAssignMul node
type ExprAssignMul struct {
	ExprAssign
}

func (n *ExprAssignMul) Visit(v Visitor, depth int) bool {
	return v.ExprAssignMul(depth, n)
}

// ExprAssignPlus node
type ExprAssignPlus struct {
	ExprAssign
}

func (n *ExprAssignPlus) Visit(v Visitor, depth int) bool {
	return v.ExprAssignPlus(depth, n)
}

// ExprAssignPow node
type ExprAssignPow struct {
	ExprAssign
}

func (n *ExprAssignPow) Visit(v Visitor, depth int) bool {
	return v.ExprAssignPow(depth, n)
}

// ExprAssignShiftLeft node
type ExprAssignShiftLeft struct {
	ExprAssign
}

func (n *ExprAssignShiftLeft) Visit(v Visitor, depth int) bool {
	return v.ExprAssignShiftLeft(depth, n)
}

// ExprAssignShiftRight node
type ExprAssignShiftRight struct {
	ExprAssign
}

func (n *ExprAssignShiftRight) Visit(v Visitor, depth int) bool {
	return v.ExprAssignShiftRight(depth, n)
}
