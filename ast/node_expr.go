package ast

// ExprArray node
type ExprArray struct {
	Node
	Items []Vertex
}

func (n *ExprArray) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Items": n.Items,
	}
}

func (n *ExprArray) Visit(v Visitor, depth int) bool {
	return v.ExprArray(depth, n)
}

// ExprArrayDimFetch node
type ExprArrayDimFetch struct {
	Node
	Var Vertex
	Dim Vertex
}

func (n *ExprArrayDimFetch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
		"Dim": {n.Dim},
	}
}

func (n *ExprArrayDimFetch) Visit(v Visitor, depth int) bool {
	return v.ExprArrayDimFetch(depth, n)
}

// ExprArrayItem node
type ExprArrayItem struct {
	Node
	Unpack bool
	Key    Vertex
	Val    Vertex
}

func (n *ExprArrayItem) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Key": {n.Key},
		"Val": {n.Val},
	}
}

func (n *ExprArrayItem) Visit(v Visitor, depth int) bool {
	return v.ExprArrayItem(depth, n)
}

// ExprArrowFunction node
type ExprArrowFunction struct {
	Node
	ReturnsRef bool
	Static     bool
	Params     []Vertex
	ReturnType Vertex
	Expr       Vertex
}

func (n *ExprArrowFunction) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Params":     n.Params,
		"ReturnType": {n.ReturnType},
		"Expr":       {n.Expr},
	}
}

func (n *ExprArrowFunction) Visit(v Visitor, depth int) bool {
	return v.ExprArrowFunction(depth, n)
}

// ExprBitwiseNot node
type ExprBitwiseNot struct {
	Node
	Expr Vertex
}

func (n *ExprBitwiseNot) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprBitwiseNot) Visit(v Visitor, depth int) bool {
	return v.ExprBitwiseNot(depth, n)
}

// ExprBooleanNot node
type ExprBooleanNot struct {
	Node
	Expr Vertex
}

func (n *ExprBooleanNot) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprBooleanNot) Visit(v Visitor, depth int) bool {
	return v.ExprBooleanNot(depth, n)
}

// ExprClassConstFetch node
type ExprClassConstFetch struct {
	Node
	Class        Vertex
	ConstantName Vertex
}

func (n *ExprClassConstFetch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Class":        {n.Class},
		"ConstantName": {n.ConstantName},
	}
}

func (n *ExprClassConstFetch) Visit(v Visitor, depth int) bool {
	return v.ExprClassConstFetch(depth, n)
}

// ExprClone node
type ExprClone struct {
	Node
	Expr Vertex
}

func (n *ExprClone) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprClone) Visit(v Visitor, depth int) bool {
	return v.ExprClone(depth, n)
}

// ExprClosure node
type ExprClosure struct {
	Node
	ReturnsRef bool
	Static     bool
	Params     []Vertex
	ClosureUse *ExprClosureUse
	ReturnType Vertex
	Stmts      []Vertex
}

func (n *ExprClosure) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Params":     n.Params,
		"ClosureUse": {n.ClosureUse},
		"ReturnType": {n.ReturnType},
		"Stmts":      n.Stmts,
	}
}

func (n *ExprClosure) Visit(v Visitor, depth int) bool {
	return v.ExprClosure(depth, n)
}

// ExprClosureUse node
type ExprClosureUse struct {
	Node
	Uses []Vertex
}

func (n *ExprClosureUse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Uses": n.Uses,
	}
}

func (n *ExprClosureUse) Visit(v Visitor, depth int) bool {
	return v.ExprClosureUse(depth, n)
}

// ExprConstFetch node
type ExprConstFetch struct {
	Node
	Const Vertex
}

func (n *ExprConstFetch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Const": {n.Const},
	}
}

func (n *ExprConstFetch) Visit(v Visitor, depth int) bool {
	return v.ExprConstFetch(depth, n)
}

// ExprEmpty node
type ExprEmpty struct {
	Node
	Expr Vertex
}

func (n *ExprEmpty) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprEmpty) Visit(v Visitor, depth int) bool {
	return v.ExprEmpty(depth, n)
}

// ExprErrorSuppress node
type ExprErrorSuppress struct {
	Node
	Expr Vertex
}

func (n *ExprErrorSuppress) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprErrorSuppress) Visit(v Visitor, depth int) bool {
	return v.ExprErrorSuppress(depth, n)
}

// ExprEval node
type ExprEval struct {
	Node
	Expr Vertex
}

func (n *ExprEval) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprEval) Visit(v Visitor, depth int) bool {
	return v.ExprEval(depth, n)
}

// ExprExit node
type ExprExit struct {
	Node
	Die  bool
	Expr Vertex
}

func (n *ExprExit) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprExit) Visit(v Visitor, depth int) bool {
	return v.ExprExit(depth, n)
}

// ExprFunctionCall node
type ExprFunctionCall struct {
	Node
	Function     Vertex
	ArgumentList *ArgumentList
}

func (n *ExprFunctionCall) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Function":     {n.Function},
		"ArgumentList": {n.ArgumentList},
	}
}

func (n *ExprFunctionCall) Visit(v Visitor, depth int) bool {
	return v.ExprFunctionCall(depth, n)
}

// ExprInclude node
type ExprInclude struct {
	Node
	Expr Vertex
}

func (n *ExprInclude) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprInclude) Visit(v Visitor, depth int) bool {
	return v.ExprInclude(depth, n)
}

// ExprIncludeOnce node
type ExprIncludeOnce struct {
	Node
	Expr Vertex
}

func (n *ExprIncludeOnce) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprIncludeOnce) Visit(v Visitor, depth int) bool {
	return v.ExprIncludeOnce(depth, n)
}

// ExprInstanceOf node
type ExprInstanceOf struct {
	Node
	Expr  Vertex
	Class Vertex
}

func (n *ExprInstanceOf) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr":  {n.Expr},
		"Class": {n.Class},
	}
}

func (n *ExprInstanceOf) Visit(v Visitor, depth int) bool {
	return v.ExprInstanceOf(depth, n)
}

// ExprIsset node
type ExprIsset struct {
	Node
	Vars []Vertex
}

func (n *ExprIsset) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Vars": n.Vars,
	}
}

func (n *ExprIsset) Visit(v Visitor, depth int) bool {
	return v.ExprIsset(depth, n)
}

// ExprList node
type ExprList struct {
	Node
	Items []Vertex
}

func (n *ExprList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Items": n.Items,
	}
}

func (n *ExprList) Visit(v Visitor, depth int) bool {
	return v.ExprList(depth, n)
}

// ExprMethodCall node
type ExprMethodCall struct {
	Node
	Var          Vertex
	Method       Vertex
	ArgumentList *ArgumentList
}

func (n *ExprMethodCall) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var":          {n.Var},
		"Method":       {n.Method},
		"ArgumentList": {n.ArgumentList},
	}
}

func (n *ExprMethodCall) Visit(v Visitor, depth int) bool {
	return v.ExprMethodCall(depth, n)
}

// ExprNew node
type ExprNew struct {
	Node
	Class        Vertex
	ArgumentList *ArgumentList
}

func (n *ExprNew) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Class":        {n.Class},
		"ArgumentList": {n.ArgumentList},
	}
}

func (n *ExprNew) Visit(v Visitor, depth int) bool {
	return v.ExprNew(depth, n)
}

// ExprPostDec node
type ExprPostDec struct {
	Node
	Var Vertex
}

func (n *ExprPostDec) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
	}
}

func (n *ExprPostDec) Visit(v Visitor, depth int) bool {
	return v.ExprPostDec(depth, n)
}

// ExprPostInc node
type ExprPostInc struct {
	Node
	Var Vertex
}

func (n *ExprPostInc) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
	}
}

func (n *ExprPostInc) Visit(v Visitor, depth int) bool {
	return v.ExprPostInc(depth, n)
}

// ExprPreDec node
type ExprPreDec struct {
	Node
	Var Vertex
}

func (n *ExprPreDec) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
	}
}

func (n *ExprPreDec) Visit(v Visitor, depth int) bool {
	return v.ExprPreDec(depth, n)
}

// ExprPreInc node
type ExprPreInc struct {
	Node
	Var Vertex
}

func (n *ExprPreInc) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
	}
}

func (n *ExprPreInc) Visit(v Visitor, depth int) bool {
	return v.ExprPreInc(depth, n)
}

// ExprPrint node
type ExprPrint struct {
	Node
	Expr Vertex
}

func (n *ExprPrint) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprPrint) Visit(v Visitor, depth int) bool {
	return v.ExprPrint(depth, n)
}

// ExprPropertyFetch node
type ExprPropertyFetch struct {
	Node
	Var      Vertex
	Property Vertex
}

func (n *ExprPropertyFetch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var":      {n.Var},
		"Property": {n.Property},
	}
}

func (n *ExprPropertyFetch) Visit(v Visitor, depth int) bool {
	return v.ExprPropertyFetch(depth, n)
}

// ExprReference node
type ExprReference struct {
	Node
	Var Vertex
}

func (n *ExprReference) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var": {n.Var},
	}
}

func (n *ExprReference) Visit(v Visitor, depth int) bool {
	return v.ExprReference(depth, n)
}

// ExprRequire node
type ExprRequire struct {
	Node
	Expr Vertex
}

func (n *ExprRequire) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprRequire) Visit(v Visitor, depth int) bool {
	return v.ExprRequire(depth, n)
}

// ExprRequireOnce node
type ExprRequireOnce struct {
	Node
	Expr Vertex
}

func (n *ExprRequireOnce) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprRequireOnce) Visit(v Visitor, depth int) bool {
	return v.ExprRequireOnce(depth, n)
}

// ExprShellExec node
type ExprShellExec struct {
	Node
	Parts []Vertex
}

func (n *ExprShellExec) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Parts": n.Parts,
	}
}

func (n *ExprShellExec) Visit(v Visitor, depth int) bool {
	return v.ExprShellExec(depth, n)
}

// ExprShortArray node
type ExprShortArray struct {
	Node
	Items []Vertex
}

func (n *ExprShortArray) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Items": n.Items,
	}
}

func (n *ExprShortArray) Visit(v Visitor, depth int) bool {
	return v.ExprShortArray(depth, n)
}

// ExprShortList node
type ExprShortList struct {
	Node
	Items []Vertex
}

func (n *ExprShortList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Items": n.Items,
	}
}

func (n *ExprShortList) Visit(v Visitor, depth int) bool {
	return v.ExprShortList(depth, n)
}

// ExprStaticCall node
type ExprStaticCall struct {
	Node
	Class        Vertex
	Call         Vertex
	ArgumentList *ArgumentList
}

func (n *ExprStaticCall) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Class":        {n.Class},
		"Call":         {n.Call},
		"ArgumentList": {n.ArgumentList},
	}
}

func (n *ExprStaticCall) Visit(v Visitor, depth int) bool {
	return v.ExprStaticCall(depth, n)
}

// ExprStaticPropertyFetch node
type ExprStaticPropertyFetch struct {
	Node
	Class    Vertex
	Property Vertex
}

func (n *ExprStaticPropertyFetch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Class":    {n.Class},
		"Property": {n.Property},
	}
}

func (n *ExprStaticPropertyFetch) Visit(v Visitor, depth int) bool {
	return v.ExprStaticPropertyFetch(depth, n)
}

// ExprTernary node
type ExprTernary struct {
	Node
	Condition Vertex
	IfTrue    Vertex
	IfFalse   Vertex
}

func (n *ExprTernary) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Condition": {n.Condition},
		"IfTrue":    {n.IfTrue},
		"IfFalse":   {n.IfFalse},
	}
}

func (n *ExprTernary) Visit(v Visitor, depth int) bool {
	return v.ExprTernary(depth, n)
}

// ExprUnaryMinus node
type ExprUnaryMinus struct {
	Node
	Expr Vertex
}

func (n *ExprUnaryMinus) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprUnaryMinus) Visit(v Visitor, depth int) bool {
	return v.ExprUnaryMinus(depth, n)
}

// ExprUnaryPlus node
type ExprUnaryPlus struct {
	Node
	Expr Vertex
}

func (n *ExprUnaryPlus) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprUnaryPlus) Visit(v Visitor, depth int) bool {
	return v.ExprUnaryPlus(depth, n)
}

// ExprVariable node
type ExprVariable struct {
	Node
	VarName Vertex
}

func (n *ExprVariable) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"VarName": {n.VarName},
	}
}

func (n *ExprVariable) Visit(v Visitor, depth int) bool {
	return v.ExprVariable(depth, n)
}

// ExprYield node
type ExprYield struct {
	Node
	Key   Vertex
	Value Vertex
}

func (n *ExprYield) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Key":   {n.Key},
		"Value": {n.Value},
	}
}

func (n *ExprYield) Visit(v Visitor, depth int) bool {
	return v.ExprYield(depth, n)
}

// ExprYieldFrom node
type ExprYieldFrom struct {
	Node
	Expr Vertex
}

func (n *ExprYieldFrom) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *ExprYieldFrom) Visit(v Visitor, depth int) bool {
	return v.ExprYieldFrom(depth, n)
}
