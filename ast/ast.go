package ast

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
)

type Visitor interface {
	Root(depth int, n *Root) bool
	Nullable(depth int, n *Nullable) bool
	Parameter(depth int, n *Parameter) bool
	Identifier(depth int, n *Identifier) bool
	ArgumentList(depth int, n *ArgumentList) bool
	Argument(depth int, n *Argument) bool

	StmtAltElse(depth int, n *StmtAltElse) bool
	StmtAltElseIf(depth int, n *StmtAltElseIf) bool
	StmtAltFor(depth int, n *StmtAltFor) bool
	StmtAltForeach(depth int, n *StmtAltForeach) bool
	StmtAltIf(depth int, n *StmtAltIf) bool
	StmtAltSwitch(depth int, n *StmtAltSwitch) bool
	StmtAltWhile(depth int, n *StmtAltWhile) bool
	StmtBreak(depth int, n *StmtBreak) bool
	StmtCase(depth int, n *StmtCase) bool
	StmtCaseList(depth int, n *StmtCaseList) bool
	StmtCatch(depth int, n *StmtCatch) bool
	StmtClass(depth int, n *StmtClass) bool
	StmtClassConstList(depth int, n *StmtClassConstList) bool
	StmtClassExtends(depth int, n *StmtClassExtends) bool
	StmtClassImplements(depth int, n *StmtClassImplements) bool
	StmtClassMethod(depth int, n *StmtClassMethod) bool
	StmtConstList(depth int, n *StmtConstList) bool
	StmtConstant(depth int, n *StmtConstant) bool
	StmtContinue(depth int, n *StmtContinue) bool
	StmtDeclare(depth int, n *StmtDeclare) bool
	StmtDefault(depth int, n *StmtDefault) bool
	StmtDo(depth int, n *StmtDo) bool
	StmtEcho(depth int, n *StmtEcho) bool
	StmtElse(depth int, n *StmtElse) bool
	StmtElseIf(depth int, n *StmtElseIf) bool
	StmtExpression(depth int, n *StmtExpression) bool
	StmtFinally(depth int, n *StmtFinally) bool
	StmtFor(depth int, n *StmtFor) bool
	StmtForeach(depth int, n *StmtForeach) bool
	StmtFunction(depth int, n *StmtFunction) bool
	StmtGlobal(depth int, n *StmtGlobal) bool
	StmtGoto(depth int, n *StmtGoto) bool
	StmtGroupUse(depth int, n *StmtGroupUse) bool
	StmtHaltCompiler(depth int, n *StmtHaltCompiler) bool
	StmtIf(depth int, n *StmtIf) bool
	StmtInlineHtml(depth int, n *StmtInlineHtml) bool
	StmtInterface(depth int, n *StmtInterface) bool
	StmtInterfaceExtends(depth int, n *StmtInterfaceExtends) bool
	StmtLabel(depth int, n *StmtLabel) bool
	StmtNamespace(depth int, n *StmtNamespace) bool
	StmtNop(depth int, n *StmtNop) bool
	StmtProperty(depth int, n *StmtProperty) bool
	StmtPropertyList(depth int, n *StmtPropertyList) bool
	StmtReturn(depth int, n *StmtReturn) bool
	StmtStatic(depth int, n *StmtStatic) bool
	StmtStaticVar(depth int, n *StmtStaticVar) bool
	StmtStmtList(depth int, n *StmtStmtList) bool
	StmtSwitch(depth int, n *StmtSwitch) bool
	StmtThrow(depth int, n *StmtThrow) bool
	StmtTrait(depth int, n *StmtTrait) bool
	StmtTraitAdaptationList(depth int, n *StmtTraitAdaptationList) bool
	StmtTraitMethodRef(depth int, n *StmtTraitMethodRef) bool
	StmtTraitUse(depth int, n *StmtTraitUse) bool
	StmtTraitUseAlias(depth int, n *StmtTraitUseAlias) bool
	StmtTraitUsePrecedence(depth int, n *StmtTraitUsePrecedence) bool
	StmtTry(depth int, n *StmtTry) bool
	StmtUnset(depth int, n *StmtUnset) bool
	StmtUse(depth int, n *StmtUse) bool
	StmtUseList(depth int, n *StmtUseList) bool
	StmtWhile(depth int, n *StmtWhile) bool

	ExprArray(depth int, n *ExprArray) bool
	ExprArrayDimFetch(depth int, n *ExprArrayDimFetch) bool
	ExprArrayItem(depth int, n *ExprArrayItem) bool
	ExprArrowFunction(depth int, n *ExprArrowFunction) bool
	ExprBitwiseNot(depth int, n *ExprBitwiseNot) bool
	ExprBooleanNot(depth int, n *ExprBooleanNot) bool
	ExprClassConstFetch(depth int, n *ExprClassConstFetch) bool
	ExprClone(depth int, n *ExprClone) bool
	ExprClosure(depth int, n *ExprClosure) bool
	ExprClosureUse(depth int, n *ExprClosureUse) bool
	ExprConstFetch(depth int, n *ExprConstFetch) bool
	ExprEmpty(depth int, n *ExprEmpty) bool
	ExprErrorSuppress(depth int, n *ExprErrorSuppress) bool
	ExprEval(depth int, n *ExprEval) bool
	ExprExit(depth int, n *ExprExit) bool
	ExprFunctionCall(depth int, n *ExprFunctionCall) bool
	ExprInclude(depth int, n *ExprInclude) bool
	ExprIncludeOnce(depth int, n *ExprIncludeOnce) bool
	ExprInstanceOf(depth int, n *ExprInstanceOf) bool
	ExprIsset(depth int, n *ExprIsset) bool
	ExprList(depth int, n *ExprList) bool
	ExprMethodCall(depth int, n *ExprMethodCall) bool
	ExprNew(depth int, n *ExprNew) bool
	ExprPostDec(depth int, n *ExprPostDec) bool
	ExprPostInc(depth int, n *ExprPostInc) bool
	ExprPreDec(depth int, n *ExprPreDec) bool
	ExprPreInc(depth int, n *ExprPreInc) bool
	ExprPrint(depth int, n *ExprPrint) bool
	ExprPropertyFetch(depth int, n *ExprPropertyFetch) bool
	ExprReference(depth int, n *ExprReference) bool
	ExprRequire(depth int, n *ExprRequire) bool
	ExprRequireOnce(depth int, n *ExprRequireOnce) bool
	ExprShellExec(depth int, n *ExprShellExec) bool
	ExprShortArray(depth int, n *ExprShortArray) bool
	ExprShortList(depth int, n *ExprShortList) bool
	ExprStaticCall(depth int, n *ExprStaticCall) bool
	ExprStaticPropertyFetch(depth int, n *ExprStaticPropertyFetch) bool
	ExprTernary(depth int, n *ExprTernary) bool
	ExprUnaryMinus(depth int, n *ExprUnaryMinus) bool
	ExprUnaryPlus(depth int, n *ExprUnaryPlus) bool
	ExprVariable(depth int, n *ExprVariable) bool
	ExprYield(depth int, n *ExprYield) bool
	ExprYieldFrom(depth int, n *ExprYieldFrom) bool

	ExprAssign(depth int, n *ExprAssign) bool
	ExprAssignReference(depth int, n *ExprAssignReference) bool
	ExprAssignBitwiseAnd(depth int, n *ExprAssignBitwiseAnd) bool
	ExprAssignBitwiseOr(depth int, n *ExprAssignBitwiseOr) bool
	ExprAssignBitwiseXor(depth int, n *ExprAssignBitwiseXor) bool
	ExprAssignCoalesce(depth int, n *ExprAssignCoalesce) bool
	ExprAssignConcat(depth int, n *ExprAssignConcat) bool
	ExprAssignDiv(depth int, n *ExprAssignDiv) bool
	ExprAssignMinus(depth int, n *ExprAssignMinus) bool
	ExprAssignMod(depth int, n *ExprAssignMod) bool
	ExprAssignMul(depth int, n *ExprAssignMul) bool
	ExprAssignPlus(depth int, n *ExprAssignPlus) bool
	ExprAssignPow(depth int, n *ExprAssignPow) bool
	ExprAssignShiftLeft(depth int, n *ExprAssignShiftLeft) bool
	ExprAssignShiftRight(depth int, n *ExprAssignShiftRight) bool

	ExprBinaryBitwiseAnd(depth int, n *ExprBinaryBitwiseAnd) bool
	ExprBinaryBitwiseOr(depth int, n *ExprBinaryBitwiseOr) bool
	ExprBinaryBitwiseXor(depth int, n *ExprBinaryBitwiseXor) bool
	ExprBinaryBooleanAnd(depth int, n *ExprBinaryBooleanAnd) bool
	ExprBinaryBooleanOr(depth int, n *ExprBinaryBooleanOr) bool
	ExprBinaryCoalesce(depth int, n *ExprBinaryCoalesce) bool
	ExprBinaryConcat(depth int, n *ExprBinaryConcat) bool
	ExprBinaryDiv(depth int, n *ExprBinaryDiv) bool
	ExprBinaryEqual(depth int, n *ExprBinaryEqual) bool
	ExprBinaryGreater(depth int, n *ExprBinaryGreater) bool
	ExprBinaryGreaterOrEqual(depth int, n *ExprBinaryGreaterOrEqual) bool
	ExprBinaryIdentical(depth int, n *ExprBinaryIdentical) bool
	ExprBinaryLogicalAnd(depth int, n *ExprBinaryLogicalAnd) bool
	ExprBinaryLogicalOr(depth int, n *ExprBinaryLogicalOr) bool
	ExprBinaryLogicalXor(depth int, n *ExprBinaryLogicalXor) bool
	ExprBinaryMinus(depth int, n *ExprBinaryMinus) bool
	ExprBinaryMod(depth int, n *ExprBinaryMod) bool
	ExprBinaryMul(depth int, n *ExprBinaryMul) bool
	ExprBinaryNotEqual(depth int, n *ExprBinaryNotEqual) bool
	ExprBinaryNotIdentical(depth int, n *ExprBinaryNotIdentical) bool
	ExprBinaryPlus(depth int, n *ExprBinaryPlus) bool
	ExprBinaryPow(depth int, n *ExprBinaryPow) bool
	ExprBinaryShiftLeft(depth int, n *ExprBinaryShiftLeft) bool
	ExprBinaryShiftRight(depth int, n *ExprBinaryShiftRight) bool
	ExprBinarySmaller(depth int, n *ExprBinarySmaller) bool
	ExprBinarySmallerOrEqual(depth int, n *ExprBinarySmallerOrEqual) bool
	ExprBinarySpaceship(depth int, n *ExprBinarySpaceship) bool

	ExprCastArray(depth int, n *ExprCastArray) bool
	ExprCastBool(depth int, n *ExprCastBool) bool
	ExprCastDouble(depth int, n *ExprCastDouble) bool
	ExprCastInt(depth int, n *ExprCastInt) bool
	ExprCastObject(depth int, n *ExprCastObject) bool
	ExprCastString(depth int, n *ExprCastString) bool
	ExprCastUnset(depth int, n *ExprCastUnset) bool

	ScalarDnumber(depth int, n *ScalarDnumber) bool
	ScalarEncapsed(depth int, n *ScalarEncapsed) bool
	ScalarEncapsedStringPart(depth int, n *ScalarEncapsedStringPart) bool
	ScalarHeredoc(depth int, n *ScalarHeredoc) bool
	ScalarLnumber(depth int, n *ScalarLnumber) bool
	ScalarMagicConstant(depth int, n *ScalarMagicConstant) bool
	ScalarString(depth int, n *ScalarString) bool
}

type Traverser interface {
	Traverse(n Vertex)
}

type Vertex interface {
	Visit(v Visitor, depth int) bool
	Children() map[string][]Vertex

	SetPosition(p *position.Position)
	GetPosition() *position.Position
	GetFreeFloating() *freefloating.Collection
}
