package ast

// StmtAltElse node
type StmtAltElse struct {
	Node
	Stmt Vertex
}

func (n *StmtAltElse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmt": {n.Stmt},
	}
}

func (n *StmtAltElse) Visit(v Visitor, depth int) bool {
	return v.StmtAltElse(depth, n)
}

// StmtAltElseIf node
type StmtAltElseIf struct {
	Node
	Cond Vertex
	Stmt Vertex
}

func (n *StmtAltElseIf) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond": {n.Cond},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtAltElseIf) Visit(v Visitor, depth int) bool {
	return v.StmtAltElseIf(depth, n)
}

// StmtAltFor node
type StmtAltFor struct {
	Node
	Init []Vertex
	Cond []Vertex
	Loop []Vertex
	Stmt Vertex
}

func (n *StmtAltFor) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Init": n.Init,
		"Cond": n.Cond,
		"Loop": n.Loop,
		"Stmt": {n.Stmt},
	}
}

func (n *StmtAltFor) Visit(v Visitor, depth int) bool {
	return v.StmtAltFor(depth, n)
}

// StmtAltForeach node
type StmtAltForeach struct {
	Node
	Expr Vertex
	Key  Vertex
	Var  Vertex
	Stmt Vertex
}

func (n *StmtAltForeach) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
		"Key":  {n.Key},
		"Var":  {n.Var},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtAltForeach) Visit(v Visitor, depth int) bool {
	return v.StmtAltForeach(depth, n)
}

// StmtAltIf node
type StmtAltIf struct {
	Node
	Cond   Vertex
	Stmt   Vertex
	ElseIf []Vertex
	Else   Vertex
}

func (n *StmtAltIf) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond":   {n.Cond},
		"Stmt":   {n.Stmt},
		"ElseIf": n.ElseIf,
		"Else":   {n.Else},
	}
}

func (n *StmtAltIf) Visit(v Visitor, depth int) bool {
	return v.StmtAltIf(depth, n)
}

// StmtAltSwitch node
type StmtAltSwitch struct {
	Node
	Cond     Vertex
	CaseList *StmtCaseList
}

func (n *StmtAltSwitch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond":     {n.Cond},
		"CaseList": {n.CaseList},
	}
}

func (n *StmtAltSwitch) Visit(v Visitor, depth int) bool {
	return v.StmtAltSwitch(depth, n)
}

// StmtAltWhile node
type StmtAltWhile struct {
	Node
	Cond Vertex
	Stmt Vertex
}

func (n *StmtAltWhile) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond": {n.Cond},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtAltWhile) Visit(v Visitor, depth int) bool {
	return v.StmtAltWhile(depth, n)
}

// StmtBreak node
type StmtBreak struct {
	Node
	Expr Vertex
}

func (n *StmtBreak) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *StmtBreak) Visit(v Visitor, depth int) bool {
	return v.StmtBreak(depth, n)
}

// StmtCase node
type StmtCase struct {
	Node
	Cond  Vertex
	Stmts []Vertex
}

func (n *StmtCase) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond":  {n.Cond},
		"Stmts": n.Stmts,
	}
}

func (n *StmtCase) Visit(v Visitor, depth int) bool {
	return v.StmtCase(depth, n)
}

// StmtCaseList node
type StmtCaseList struct {
	Node
	Cases []Vertex
}

func (n *StmtCaseList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cases": n.Cases,
	}
}

func (n *StmtCaseList) Visit(v Visitor, depth int) bool {
	return v.StmtCaseList(depth, n)
}

// StmtCatch node
type StmtCatch struct {
	Node
	Types []Vertex
	Var   Vertex
	Stmts []Vertex
}

func (n *StmtCatch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Types": n.Types,
		"Var":   {n.Var},
		"Stmts": n.Stmts,
	}
}

func (n *StmtCatch) Visit(v Visitor, depth int) bool {
	return v.StmtCatch(depth, n)
}

// StmtClass node
type StmtClass struct {
	Node
	ClassName    Vertex
	Modifiers    []Vertex
	ArgumentList *ArgumentList
	Extends      *StmtClassExtends
	Implements   *StmtClassImplements
	Stmts        []Vertex
}

func (n *StmtClass) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"ClassName":    {n.ClassName},
		"Modifiers":    n.Modifiers,
		"ArgumentList": {n.ArgumentList},
		"Extends":      {n.Extends},
		"Implements":   {n.Implements},
		"Stmts":        n.Stmts,
	}
}

func (n *StmtClass) Visit(v Visitor, depth int) bool {
	return v.StmtClass(depth, n)
}

// StmtClassConstList node
type StmtClassConstList struct {
	Node
	Modifiers []Vertex
	Consts    []Vertex
}

func (n *StmtClassConstList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Modifiers": n.Modifiers,
		"Consts":    n.Consts,
	}
}

func (n *StmtClassConstList) Visit(v Visitor, depth int) bool {
	return v.StmtClassConstList(depth, n)
}

// StmtClassExtends node
type StmtClassExtends struct {
	Node
	ClassName Vertex
}

func (n *StmtClassExtends) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"ClassName": {n.ClassName},
	}
}

func (n *StmtClassExtends) Visit(v Visitor, depth int) bool {
	return v.StmtClassExtends(depth, n)
}

// StmtClassImplements node
type StmtClassImplements struct {
	Node
	InterfaceNames []Vertex
}

func (n *StmtClassImplements) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"InterfaceNames": n.InterfaceNames,
	}
}

func (n *StmtClassImplements) Visit(v Visitor, depth int) bool {
	return v.StmtClassImplements(depth, n)
}

// StmtClassMethod node
type StmtClassMethod struct {
	Node
	ReturnsRef bool
	MethodName Vertex
	Modifiers  []Vertex
	Params     []Vertex
	ReturnType Vertex
	Stmt       Vertex
}

func (n *StmtClassMethod) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"MethodName": {n.MethodName},
		"Modifiers":  n.Modifiers,
		"Params":     n.Params,
		"ReturnType": {n.ReturnType},
		"Stmt":       {n.Stmt},
	}
}

func (n *StmtClassMethod) Visit(v Visitor, depth int) bool {
	return v.StmtClassMethod(depth, n)
}

// StmtConstList node
type StmtConstList struct {
	Node
	Consts []Vertex
}

func (n *StmtConstList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Consts": n.Consts,
	}
}

func (n *StmtConstList) Visit(v Visitor, depth int) bool {
	return v.StmtConstList(depth, n)
}

// StmtConstant node
type StmtConstant struct {
	Node
	ConstantName Vertex
	Expr         Vertex
}

func (n *StmtConstant) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"ConstantName": {n.ConstantName},
		"Expr":         {n.Expr},
	}
}

func (n *StmtConstant) Visit(v Visitor, depth int) bool {
	return v.StmtConstant(depth, n)
}

// StmtContinue node
type StmtContinue struct {
	Node
	Expr Vertex
}

func (n *StmtContinue) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *StmtContinue) Visit(v Visitor, depth int) bool {
	return v.StmtContinue(depth, n)
}

// StmtDeclare node
type StmtDeclare struct {
	Node
	Alt    bool
	Consts []Vertex
	Stmt   Vertex
}

func (n *StmtDeclare) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Consts": n.Consts,
		"Stmt":   {n.Stmt},
	}
}

func (n *StmtDeclare) Visit(v Visitor, depth int) bool {
	return v.StmtDeclare(depth, n)
}

// StmtDefault node
type StmtDefault struct {
	Node
	Stmts []Vertex
}

func (n *StmtDefault) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmts": n.Stmts,
	}
}

func (n *StmtDefault) Visit(v Visitor, depth int) bool {
	return v.StmtDefault(depth, n)
}

// StmtDo node
type StmtDo struct {
	Node
	Stmt Vertex
	Cond Vertex
}

func (n *StmtDo) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmt": {n.Stmt},
		"Cond": {n.Cond},
	}
}

func (n *StmtDo) Visit(v Visitor, depth int) bool {
	return v.StmtDo(depth, n)
}

// StmtEcho node
type StmtEcho struct {
	Node
	Exprs []Vertex
}

func (n *StmtEcho) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Exprs": n.Exprs,
	}
}

func (n *StmtEcho) Visit(v Visitor, depth int) bool {
	return v.StmtEcho(depth, n)
}

// StmtElse node
type StmtElse struct {
	Node
	Stmt Vertex
}

func (n *StmtElse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmt": {n.Stmt},
	}
}

func (n *StmtElse) Visit(v Visitor, depth int) bool {
	return v.StmtElse(depth, n)
}

// StmtElseIf node
type StmtElseIf struct {
	Node
	Cond Vertex
	Stmt Vertex
}

func (n *StmtElseIf) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond": {n.Cond},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtElseIf) Visit(v Visitor, depth int) bool {
	return v.StmtElseIf(depth, n)
}

// StmtExpression node
type StmtExpression struct {
	Node
	Expr Vertex
}

func (n *StmtExpression) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *StmtExpression) Visit(v Visitor, depth int) bool {
	return v.StmtExpression(depth, n)
}

// StmtFinally node
type StmtFinally struct {
	Node
	Stmts []Vertex
}

func (n *StmtFinally) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmts": n.Stmts,
	}
}

func (n *StmtFinally) Visit(v Visitor, depth int) bool {
	return v.StmtFinally(depth, n)
}

// StmtFor node
type StmtFor struct {
	Node
	Init []Vertex
	Cond []Vertex
	Loop []Vertex
	Stmt Vertex
}

func (n *StmtFor) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Init": n.Init,
		"Cond": n.Cond,
		"Loop": n.Loop,
		"Stmt": {n.Stmt},
	}
}

func (n *StmtFor) Visit(v Visitor, depth int) bool {
	return v.StmtFor(depth, n)
}

// StmtForeach node
type StmtForeach struct {
	Node
	Expr Vertex
	Key  Vertex
	Var  Vertex
	Stmt Vertex
}

func (n *StmtForeach) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
		"Key":  {n.Key},
		"Var":  {n.Var},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtForeach) Visit(v Visitor, depth int) bool {
	return v.StmtForeach(depth, n)
}

// StmtFunction node
type StmtFunction struct {
	Node
	ReturnsRef   bool
	FunctionName Vertex
	Params       []Vertex
	ReturnType   Vertex
	Stmts        []Vertex
}

func (n *StmtFunction) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"FunctionName": {n.FunctionName},
		"Params":       n.Params,
		"ReturnType":   {n.ReturnType},
		"Stmts":        n.Stmts,
	}
}

func (n *StmtFunction) Visit(v Visitor, depth int) bool {
	return v.StmtFunction(depth, n)
}

// StmtGlobal node
type StmtGlobal struct {
	Node
	Vars []Vertex
}

func (n *StmtGlobal) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Vars": n.Vars,
	}
}

func (n *StmtGlobal) Visit(v Visitor, depth int) bool {
	return v.StmtGlobal(depth, n)
}

// StmtGoto node
type StmtGoto struct {
	Node
	Label Vertex
}

func (n *StmtGoto) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Label": {n.Label},
	}
}

func (n *StmtGoto) Visit(v Visitor, depth int) bool {
	return v.StmtGoto(depth, n)
}

// StmtGroupUse node
type StmtGroupUse struct {
	Node
	UseType Vertex
	Prefix  Vertex
	UseList []Vertex
}

func (n *StmtGroupUse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"UseType": {n.UseType},
		"Prefix":  {n.Prefix},
		"UseList": n.UseList,
	}
}

func (n *StmtGroupUse) Visit(v Visitor, depth int) bool {
	return v.StmtGroupUse(depth, n)
}

// StmtHaltCompiler node
type StmtHaltCompiler struct {
	Node
}

func (n *StmtHaltCompiler) Children() map[string][]Vertex {
	return nil
}

func (n *StmtHaltCompiler) Visit(v Visitor, depth int) bool {
	return v.StmtHaltCompiler(depth, n)
}

// StmtIf node
type StmtIf struct {
	Node
	Cond   Vertex
	Stmt   Vertex
	ElseIf []Vertex
	Else   Vertex
}

func (n *StmtIf) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond":   {n.Cond},
		"Stmt":   {n.Stmt},
		"ElseIf": n.ElseIf,
		"Else":   {n.Else},
	}
}

func (n *StmtIf) Visit(v Visitor, depth int) bool {
	return v.StmtIf(depth, n)
}

// StmtInlineHtml node
type StmtInlineHtml struct {
	Node
	Value string
}

func (n *StmtInlineHtml) Children() map[string][]Vertex {
	return nil
}

func (n *StmtInlineHtml) Visit(v Visitor, depth int) bool {
	return v.StmtInlineHtml(depth, n)
}

// StmtInterface node
type StmtInterface struct {
	Node
	InterfaceName Vertex
	Extends       *StmtInterfaceExtends
	Stmts         []Vertex
}

func (n *StmtInterface) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"InterfaceName": {n.InterfaceName},
		"Extends":       {n.Extends},
		"Stmts":         n.Stmts,
	}
}

func (n *StmtInterface) Visit(v Visitor, depth int) bool {
	return v.StmtInterface(depth, n)
}

// StmtInterfaceExtends node
type StmtInterfaceExtends struct {
	Node
	InterfaceNames []Vertex
}

func (n *StmtInterfaceExtends) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"InterfaceNames": n.InterfaceNames,
	}
}

func (n *StmtInterfaceExtends) Visit(v Visitor, depth int) bool {
	return v.StmtInterfaceExtends(depth, n)
}

// StmtLabel node
type StmtLabel struct {
	Node
	LabelName Vertex
}

func (n *StmtLabel) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"LabelName": {n.LabelName},
	}
}

func (n *StmtLabel) Visit(v Visitor, depth int) bool {
	return v.StmtLabel(depth, n)
}

// StmtNamespace node
type StmtNamespace struct {
	Node
	NamespaceName Vertex
	Stmts         []Vertex
}

func (n *StmtNamespace) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"NamespaceName": {n.NamespaceName},
		"Stmts":         n.Stmts,
	}
}

func (n *StmtNamespace) Visit(v Visitor, depth int) bool {
	return v.StmtNamespace(depth, n)
}

// StmtNop node
type StmtNop struct {
	Node
}

func (n *StmtNop) Children() map[string][]Vertex {
	return nil
}

func (n *StmtNop) Visit(v Visitor, depth int) bool {
	return v.StmtNop(depth, n)
}

// StmtProperty node
type StmtProperty struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *StmtProperty) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var":  {n.Var},
		"Expr": {n.Expr},
	}
}

func (n *StmtProperty) Visit(v Visitor, depth int) bool {
	return v.StmtProperty(depth, n)
}

// StmtPropertyList node
type StmtPropertyList struct {
	Node
	Modifiers  []Vertex
	Type       Vertex
	Properties []Vertex
}

func (n *StmtPropertyList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Modifiers":  n.Modifiers,
		"Type":       {n.Type},
		"Properties": n.Properties,
	}
}

func (n *StmtPropertyList) Visit(v Visitor, depth int) bool {
	return v.StmtPropertyList(depth, n)
}

// StmtReturn node
type StmtReturn struct {
	Node
	Expr Vertex
}

func (n *StmtReturn) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *StmtReturn) Visit(v Visitor, depth int) bool {
	return v.StmtReturn(depth, n)
}

// StmtStatic node
type StmtStatic struct {
	Node
	Vars []Vertex
}

func (n *StmtStatic) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Vars": n.Vars,
	}
}

func (n *StmtStatic) Visit(v Visitor, depth int) bool {
	return v.StmtStatic(depth, n)
}

// StmtStaticVar node
type StmtStaticVar struct {
	Node
	Var  Vertex
	Expr Vertex
}

func (n *StmtStaticVar) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Var":  {n.Var},
		"Expr": {n.Expr},
	}
}

func (n *StmtStaticVar) Visit(v Visitor, depth int) bool {
	return v.StmtStaticVar(depth, n)
}

// StmtStmtList node
type StmtStmtList struct {
	Node
	Stmts []Vertex
}

func (n *StmtStmtList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmts": n.Stmts,
	}
}

func (n *StmtStmtList) Visit(v Visitor, depth int) bool {
	return v.StmtStmtList(depth, n)
}

// StmtSwitch node
type StmtSwitch struct {
	Node
	Cond     Vertex
	CaseList *StmtCaseList
}

func (n *StmtSwitch) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond":     {n.Cond},
		"CaseList": {n.CaseList},
	}
}

func (n *StmtSwitch) Visit(v Visitor, depth int) bool {
	return v.StmtSwitch(depth, n)
}

// StmtThrow node
type StmtThrow struct {
	Node
	Expr Vertex
}

func (n *StmtThrow) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *StmtThrow) Visit(v Visitor, depth int) bool {
	return v.StmtThrow(depth, n)
}

// StmtTrait node
type StmtTrait struct {
	Node
	TraitName Vertex
	Stmts     []Vertex
}

func (n *StmtTrait) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"TraitName": {n.TraitName},
		"Stmts":     n.Stmts,
	}
}

func (n *StmtTrait) Visit(v Visitor, depth int) bool {
	return v.StmtTrait(depth, n)
}

// StmtTraitAdaptationList node
type StmtTraitAdaptationList struct {
	Node
	Adaptations []Vertex
}

func (n *StmtTraitAdaptationList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Adaptations": n.Adaptations,
	}
}

func (n *StmtTraitAdaptationList) Visit(v Visitor, depth int) bool {
	return v.StmtTraitAdaptationList(depth, n)
}

// StmtTraitMethodRef node
type StmtTraitMethodRef struct {
	Node
	Trait  Vertex
	Method Vertex
}

func (n *StmtTraitMethodRef) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Trait":  {n.Trait},
		"Method": {n.Method},
	}
}

func (n *StmtTraitMethodRef) Visit(v Visitor, depth int) bool {
	return v.StmtTraitMethodRef(depth, n)
}

// StmtTraitUse node
type StmtTraitUse struct {
	Node
	Traits              []Vertex
	TraitAdaptationList Vertex
}

func (n *StmtTraitUse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Traits":              n.Traits,
		"TraitAdaptationList": {n.TraitAdaptationList},
	}
}

func (n *StmtTraitUse) Visit(v Visitor, depth int) bool {
	return v.StmtTraitUse(depth, n)
}

// StmtTraitUseAlias node
type StmtTraitUseAlias struct {
	Node
	Ref      Vertex
	Modifier Vertex
	Alias    Vertex
}

func (n *StmtTraitUseAlias) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Ref":      {n.Ref},
		"Modifier": {n.Modifier},
		"Alias":    {n.Alias},
	}
}

func (n *StmtTraitUseAlias) Visit(v Visitor, depth int) bool {
	return v.StmtTraitUseAlias(depth, n)
}

// StmtTraitUsePrecedence node
type StmtTraitUsePrecedence struct {
	Node
	Ref       Vertex
	Insteadof []Vertex
}

func (n *StmtTraitUsePrecedence) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Ref":       {n.Ref},
		"Insteadof": n.Insteadof,
	}
}

func (n *StmtTraitUsePrecedence) Visit(v Visitor, depth int) bool {
	return v.StmtTraitUsePrecedence(depth, n)
}

// StmtTry node
type StmtTry struct {
	Node
	Stmts   []Vertex
	Catches []Vertex
	Finally Vertex
}

func (n *StmtTry) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmts":   n.Stmts,
		"Catches": n.Catches,
		"Finally": {n.Finally},
	}
}

func (n *StmtTry) Visit(v Visitor, depth int) bool {
	return v.StmtTry(depth, n)
}

// StmtUnset node
type StmtUnset struct {
	Node
	Vars []Vertex
}

func (n *StmtUnset) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Vars": n.Vars,
	}
}

func (n *StmtUnset) Visit(v Visitor, depth int) bool {
	return v.StmtUnset(depth, n)
}

// StmtUse node
type StmtUse struct {
	Node
	UseType Vertex
	Use     Vertex
	Alias   Vertex
}

func (n *StmtUse) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"UseType": {n.UseType},
		"Use":     {n.Use},
		"Alias":   {n.Alias},
	}
}

func (n *StmtUse) Visit(v Visitor, depth int) bool {
	return v.StmtUse(depth, n)
}

// StmtUseList node
type StmtUseList struct {
	Node
	UseType Vertex
	Uses    []Vertex
}

func (n *StmtUseList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"UseType": {n.UseType},
		"Uses":    n.Uses,
	}
}

func (n *StmtUseList) Visit(v Visitor, depth int) bool {
	return v.StmtUseList(depth, n)
}

// StmtWhile node
type StmtWhile struct {
	Node
	Cond Vertex
	Stmt Vertex
}

func (n *StmtWhile) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Cond": {n.Cond},
		"Stmt": {n.Stmt},
	}
}

func (n *StmtWhile) Visit(v Visitor, depth int) bool {
	return v.StmtWhile(depth, n)
}
