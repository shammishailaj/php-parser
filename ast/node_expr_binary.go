package ast

// ExprBinary node
type ExprBinary struct {
	Node
	Left  Vertex
	Right Vertex
}

func (n *ExprBinary) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Left":  {n.Left},
		"Right": {n.Right},
	}
}

// ExprBinaryBitwiseAnd node
type ExprBinaryBitwiseAnd struct {
	ExprBinary
}

func (n *ExprBinaryBitwiseAnd) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryBitwiseAnd(depth, n)
}

// ExprBinaryBitwiseOr node
type ExprBinaryBitwiseOr struct {
	ExprBinary
}

func (n *ExprBinaryBitwiseOr) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryBitwiseOr(depth, n)
}

// ExprBinaryBitwiseXor node
type ExprBinaryBitwiseXor struct {
	ExprBinary
}

func (n *ExprBinaryBitwiseXor) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryBitwiseXor(depth, n)
}

// ExprBinaryBooleanAnd node
type ExprBinaryBooleanAnd struct {
	ExprBinary
}

func (n *ExprBinaryBooleanAnd) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryBooleanAnd(depth, n)
}

// ExprBinaryBooleanOr node
type ExprBinaryBooleanOr struct {
	ExprBinary
}

func (n *ExprBinaryBooleanOr) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryBooleanOr(depth, n)
}

// ExprBinaryCoalesce node
type ExprBinaryCoalesce struct {
	ExprBinary
}

func (n *ExprBinaryCoalesce) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryCoalesce(depth, n)
}

// ExprBinaryConcat node
type ExprBinaryConcat struct {
	ExprBinary
}

func (n *ExprBinaryConcat) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryConcat(depth, n)
}

// ExprBinaryDiv node
type ExprBinaryDiv struct {
	ExprBinary
}

func (n *ExprBinaryDiv) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryDiv(depth, n)
}

// ExprBinaryEqual node
type ExprBinaryEqual struct {
	ExprBinary
}

func (n *ExprBinaryEqual) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryEqual(depth, n)
}

// ExprBinaryGreater node
type ExprBinaryGreater struct {
	ExprBinary
}

func (n *ExprBinaryGreater) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryGreater(depth, n)
}

// ExprBinaryGreaterOrEqual node
type ExprBinaryGreaterOrEqual struct {
	ExprBinary
}

func (n *ExprBinaryGreaterOrEqual) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryGreaterOrEqual(depth, n)
}

// ExprBinaryIdentical node
type ExprBinaryIdentical struct {
	ExprBinary
}

func (n *ExprBinaryIdentical) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryIdentical(depth, n)
}

// ExprBinaryLogicalAnd node
type ExprBinaryLogicalAnd struct {
	ExprBinary
}

func (n *ExprBinaryLogicalAnd) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryLogicalAnd(depth, n)
}

// ExprBinaryLogicalOr node
type ExprBinaryLogicalOr struct {
	ExprBinary
}

func (n *ExprBinaryLogicalOr) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryLogicalOr(depth, n)
}

// ExprBinaryLogicalXor node
type ExprBinaryLogicalXor struct {
	ExprBinary
}

func (n *ExprBinaryLogicalXor) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryLogicalXor(depth, n)
}

// ExprBinaryMinus node
type ExprBinaryMinus struct {
	ExprBinary
}

func (n *ExprBinaryMinus) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryMinus(depth, n)
}

// ExprBinaryMod node
type ExprBinaryMod struct {
	ExprBinary
}

func (n *ExprBinaryMod) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryMod(depth, n)
}

// ExprBinaryMul node
type ExprBinaryMul struct {
	ExprBinary
}

func (n *ExprBinaryMul) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryMul(depth, n)
}

// ExprBinaryNotEqual node
type ExprBinaryNotEqual struct {
	ExprBinary
}

func (n *ExprBinaryNotEqual) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryNotEqual(depth, n)
}

// ExprBinaryNotIdentical node
type ExprBinaryNotIdentical struct {
	ExprBinary
}

func (n *ExprBinaryNotIdentical) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryNotIdentical(depth, n)
}

// ExprBinaryPlus node
type ExprBinaryPlus struct {
	ExprBinary
}

func (n *ExprBinaryPlus) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryPlus(depth, n)
}

// ExprBinaryPow node
type ExprBinaryPow struct {
	ExprBinary
}

func (n *ExprBinaryPow) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryPow(depth, n)
}

// ExprBinaryShiftLeft node
type ExprBinaryShiftLeft struct {
	ExprBinary
}

func (n *ExprBinaryShiftLeft) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryShiftLeft(depth, n)
}

// ExprBinaryShiftRight node
type ExprBinaryShiftRight struct {
	ExprBinary
}

func (n *ExprBinaryShiftRight) Visit(v Visitor, depth int) bool {
	return v.ExprBinaryShiftRight(depth, n)
}

// ExprBinarySmaller node
type ExprBinarySmaller struct {
	ExprBinary
}

func (n *ExprBinarySmaller) Visit(v Visitor, depth int) bool {
	return v.ExprBinarySmaller(depth, n)
}

// ExprBinarySmallerOrEqual node
type ExprBinarySmallerOrEqual struct {
	ExprBinary
}

func (n *ExprBinarySmallerOrEqual) Visit(v Visitor, depth int) bool {
	return v.ExprBinarySmallerOrEqual(depth, n)
}

// ExprBinarySpaceship node
type ExprBinarySpaceship struct {
	ExprBinary
}

func (n *ExprBinarySpaceship) Visit(v Visitor, depth int) bool {
	return v.ExprBinarySpaceship(depth, n)
}
