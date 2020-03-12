package visitor

import (
	"github.com/z7zmey/php-parser/ast"
)

type Null struct {
}

func (v *Null) Root(_ int, _ *ast.Root) bool {
	return true
}

func (v *Null) Nullable(_ int, _ *ast.Nullable) bool {
	return true
}

func (v *Null) Parameter(_ int, _ *ast.Parameter) bool {
	return true
}

func (v *Null) Identifier(_ int, _ *ast.Identifier) bool {
	return true
}

func (v *Null) ArgumentList(_ int, _ *ast.ArgumentList) bool {
	return true
}

func (v *Null) Argument(_ int, _ *ast.Argument) bool {
	return true
}

func (v *Null) StmtAltElse(_ int, _ *ast.StmtAltElse) bool {
	return true
}

func (v *Null) StmtAltElseIf(_ int, _ *ast.StmtAltElseIf) bool {
	return true
}

func (v *Null) StmtAltFor(_ int, _ *ast.StmtAltFor) bool {
	return true
}

func (v *Null) StmtAltForeach(_ int, _ *ast.StmtAltForeach) bool {
	return true
}

func (v *Null) StmtAltIf(_ int, _ *ast.StmtAltIf) bool {
	return true
}

func (v *Null) StmtAltSwitch(_ int, _ *ast.StmtAltSwitch) bool {
	return true
}

func (v *Null) StmtAltWhile(_ int, _ *ast.StmtAltWhile) bool {
	return true
}

func (v *Null) StmtBreak(_ int, _ *ast.StmtBreak) bool {
	return true
}

func (v *Null) StmtCase(_ int, _ *ast.StmtCase) bool {
	return true
}

func (v *Null) StmtCaseList(_ int, _ *ast.StmtCaseList) bool {
	return true
}

func (v *Null) StmtCatch(_ int, _ *ast.StmtCatch) bool {
	return true
}

func (v *Null) StmtClass(_ int, _ *ast.StmtClass) bool {
	return true
}

func (v *Null) StmtClassConstList(_ int, _ *ast.StmtClassConstList) bool {
	return true
}

func (v *Null) StmtClassExtends(_ int, _ *ast.StmtClassExtends) bool {
	return true
}

func (v *Null) StmtClassImplements(_ int, _ *ast.StmtClassImplements) bool {
	return true
}

func (v *Null) StmtClassMethod(_ int, _ *ast.StmtClassMethod) bool {
	return true
}

func (v *Null) StmtConstList(_ int, _ *ast.StmtConstList) bool {
	return true
}

func (v *Null) StmtConstant(_ int, _ *ast.StmtConstant) bool {
	return true
}

func (v *Null) StmtContinue(_ int, _ *ast.StmtContinue) bool {
	return true
}

func (v *Null) StmtDeclare(_ int, _ *ast.StmtDeclare) bool {
	return true
}

func (v *Null) StmtDefault(_ int, _ *ast.StmtDefault) bool {
	return true
}

func (v *Null) StmtDo(_ int, _ *ast.StmtDo) bool {
	return true
}

func (v *Null) StmtEcho(_ int, _ *ast.StmtEcho) bool {
	return true
}

func (v *Null) StmtElse(_ int, _ *ast.StmtElse) bool {
	return true
}

func (v *Null) StmtElseIf(_ int, _ *ast.StmtElseIf) bool {
	return true
}

func (v *Null) StmtExpression(_ int, _ *ast.StmtExpression) bool {
	return true
}

func (v *Null) StmtFinally(_ int, _ *ast.StmtFinally) bool {
	return true
}

func (v *Null) StmtFor(_ int, _ *ast.StmtFor) bool {
	return true
}

func (v *Null) StmtForeach(_ int, _ *ast.StmtForeach) bool {
	return true
}

func (v *Null) StmtFunction(_ int, _ *ast.StmtFunction) bool {
	return true
}

func (v *Null) StmtGlobal(_ int, _ *ast.StmtGlobal) bool {
	return true
}

func (v *Null) StmtGoto(_ int, _ *ast.StmtGoto) bool {
	return true
}

func (v *Null) StmtGroupUse(_ int, _ *ast.StmtGroupUse) bool {
	return true
}

func (v *Null) StmtHaltCompiler(_ int, _ *ast.StmtHaltCompiler) bool {
	return true
}

func (v *Null) StmtIf(_ int, _ *ast.StmtIf) bool {
	return true
}

func (v *Null) StmtInlineHtml(_ int, _ *ast.StmtInlineHtml) bool {
	return true
}

func (v *Null) StmtInterface(_ int, _ *ast.StmtInterface) bool {
	return true
}

func (v *Null) StmtInterfaceExtends(_ int, _ *ast.StmtInterfaceExtends) bool {
	return true
}

func (v *Null) StmtLabel(_ int, _ *ast.StmtLabel) bool {
	return true
}

func (v *Null) StmtNamespace(_ int, _ *ast.StmtNamespace) bool {
	return true
}

func (v *Null) StmtNop(_ int, _ *ast.StmtNop) bool {
	return true
}

func (v *Null) StmtProperty(_ int, _ *ast.StmtProperty) bool {
	return true
}

func (v *Null) StmtPropertyList(_ int, _ *ast.StmtPropertyList) bool {
	return true
}

func (v *Null) StmtReturn(_ int, _ *ast.StmtReturn) bool {
	return true
}

func (v *Null) StmtStatic(_ int, _ *ast.StmtStatic) bool {
	return true
}

func (v *Null) StmtStaticVar(_ int, _ *ast.StmtStaticVar) bool {
	return true
}

func (v *Null) StmtStmtList(_ int, _ *ast.StmtStmtList) bool {
	return true
}

func (v *Null) StmtSwitch(_ int, _ *ast.StmtSwitch) bool {
	return true
}

func (v *Null) StmtThrow(_ int, _ *ast.StmtThrow) bool {
	return true
}

func (v *Null) StmtTrait(_ int, _ *ast.StmtTrait) bool {
	return true
}

func (v *Null) StmtTraitAdaptationList(_ int, _ *ast.StmtTraitAdaptationList) bool {
	return true
}

func (v *Null) StmtTraitMethodRef(_ int, _ *ast.StmtTraitMethodRef) bool {
	return true
}

func (v *Null) StmtTraitUse(_ int, _ *ast.StmtTraitUse) bool {
	return true
}

func (v *Null) StmtTraitUseAlias(_ int, _ *ast.StmtTraitUseAlias) bool {
	return true
}

func (v *Null) StmtTraitUsePrecedence(_ int, _ *ast.StmtTraitUsePrecedence) bool {
	return true
}

func (v *Null) StmtTry(_ int, _ *ast.StmtTry) bool {
	return true
}

func (v *Null) StmtUnset(_ int, _ *ast.StmtUnset) bool {
	return true
}

func (v *Null) StmtUse(_ int, _ *ast.StmtUse) bool {
	return true
}

func (v *Null) StmtUseList(_ int, _ *ast.StmtUseList) bool {
	return true
}

func (v *Null) StmtWhile(_ int, _ *ast.StmtWhile) bool {
	return true
}

func (v *Null) ExprArray(_ int, _ *ast.ExprArray) bool {
	return true
}

func (v *Null) ExprArrayDimFetch(_ int, _ *ast.ExprArrayDimFetch) bool {
	return true
}

func (v *Null) ExprArrayItem(_ int, _ *ast.ExprArrayItem) bool {
	return true
}

func (v *Null) ExprArrowFunction(_ int, _ *ast.ExprArrowFunction) bool {
	return true
}

func (v *Null) ExprBitwiseNot(_ int, _ *ast.ExprBitwiseNot) bool {
	return true
}

func (v *Null) ExprBooleanNot(_ int, _ *ast.ExprBooleanNot) bool {
	return true
}

func (v *Null) ExprClassConstFetch(_ int, _ *ast.ExprClassConstFetch) bool {
	return true
}

func (v *Null) ExprClone(_ int, _ *ast.ExprClone) bool {
	return true
}

func (v *Null) ExprClosure(_ int, _ *ast.ExprClosure) bool {
	return true
}

func (v *Null) ExprClosureUse(_ int, _ *ast.ExprClosureUse) bool {
	return true
}

func (v *Null) ExprConstFetch(_ int, _ *ast.ExprConstFetch) bool {
	return true
}

func (v *Null) ExprEmpty(_ int, _ *ast.ExprEmpty) bool {
	return true
}

func (v *Null) ExprErrorSuppress(_ int, _ *ast.ExprErrorSuppress) bool {
	return true
}

func (v *Null) ExprEval(_ int, _ *ast.ExprEval) bool {
	return true
}

func (v *Null) ExprExit(_ int, _ *ast.ExprExit) bool {
	return true
}

func (v *Null) ExprFunctionCall(_ int, _ *ast.ExprFunctionCall) bool {
	return true
}

func (v *Null) ExprInclude(_ int, _ *ast.ExprInclude) bool {
	return true
}

func (v *Null) ExprIncludeOnce(_ int, _ *ast.ExprIncludeOnce) bool {
	return true
}

func (v *Null) ExprInstanceOf(_ int, _ *ast.ExprInstanceOf) bool {
	return true
}

func (v *Null) ExprIsset(_ int, _ *ast.ExprIsset) bool {
	return true
}

func (v *Null) ExprList(_ int, _ *ast.ExprList) bool {
	return true
}

func (v *Null) ExprMethodCall(_ int, _ *ast.ExprMethodCall) bool {
	return true
}

func (v *Null) ExprNew(_ int, _ *ast.ExprNew) bool {
	return true
}

func (v *Null) ExprPostDec(_ int, _ *ast.ExprPostDec) bool {
	return true
}

func (v *Null) ExprPostInc(_ int, _ *ast.ExprPostInc) bool {
	return true
}

func (v *Null) ExprPreDec(_ int, _ *ast.ExprPreDec) bool {
	return true
}

func (v *Null) ExprPreInc(_ int, _ *ast.ExprPreInc) bool {
	return true
}

func (v *Null) ExprPrint(_ int, _ *ast.ExprPrint) bool {
	return true
}

func (v *Null) ExprPropertyFetch(_ int, _ *ast.ExprPropertyFetch) bool {
	return true
}

func (v *Null) ExprReference(_ int, _ *ast.ExprReference) bool {
	return true
}

func (v *Null) ExprRequire(_ int, _ *ast.ExprRequire) bool {
	return true
}

func (v *Null) ExprRequireOnce(_ int, _ *ast.ExprRequireOnce) bool {
	return true
}

func (v *Null) ExprShellExec(_ int, _ *ast.ExprShellExec) bool {
	return true
}

func (v *Null) ExprShortArray(_ int, _ *ast.ExprShortArray) bool {
	return true
}

func (v *Null) ExprShortList(_ int, _ *ast.ExprShortList) bool {
	return true
}

func (v *Null) ExprStaticCall(_ int, _ *ast.ExprStaticCall) bool {
	return true
}

func (v *Null) ExprStaticPropertyFetch(_ int, _ *ast.ExprStaticPropertyFetch) bool {
	return true
}

func (v *Null) ExprTernary(_ int, _ *ast.ExprTernary) bool {
	return true
}

func (v *Null) ExprUnaryMinus(_ int, _ *ast.ExprUnaryMinus) bool {
	return true
}

func (v *Null) ExprUnaryPlus(_ int, _ *ast.ExprUnaryPlus) bool {
	return true
}

func (v *Null) ExprVariable(_ int, _ *ast.ExprVariable) bool {
	return true
}

func (v *Null) ExprYield(_ int, _ *ast.ExprYield) bool {
	return true
}

func (v *Null) ExprYieldFrom(_ int, _ *ast.ExprYieldFrom) bool {
	return true
}

func (v *Null) ExprAssign(_ int, _ *ast.ExprAssign) bool {
	return true
}


func (v *Null) ExprAssignReference(_ int, _ *ast.ExprAssignReference) bool {
	return true
}


func (v *Null) ExprAssignBitwiseAnd(_ int, _ *ast.ExprAssignBitwiseAnd) bool {
	return true
}


func (v *Null) ExprAssignBitwiseOr(_ int, _ *ast.ExprAssignBitwiseOr) bool {
	return true
}


func (v *Null) ExprAssignBitwiseXor(_ int, _ *ast.ExprAssignBitwiseXor) bool {
	return true
}


func (v *Null) ExprAssignCoalesce(_ int, _ *ast.ExprAssignCoalesce) bool {
	return true
}


func (v *Null) ExprAssignConcat(_ int, _ *ast.ExprAssignConcat) bool {
	return true
}


func (v *Null) ExprAssignDiv(_ int, _ *ast.ExprAssignDiv) bool {
	return true
}


func (v *Null) ExprAssignMinus(_ int, _ *ast.ExprAssignMinus) bool {
	return true
}


func (v *Null) ExprAssignMod(_ int, _ *ast.ExprAssignMod) bool {
	return true
}


func (v *Null) ExprAssignMul(_ int, _ *ast.ExprAssignMul) bool {
	return true
}


func (v *Null) ExprAssignPlus(_ int, _ *ast.ExprAssignPlus) bool {
	return true
}


func (v *Null) ExprAssignPow(_ int, _ *ast.ExprAssignPow) bool {
	return true
}


func (v *Null) ExprAssignShiftLeft(_ int, _ *ast.ExprAssignShiftLeft) bool {
	return true
}


func (v *Null) ExprAssignShiftRight(_ int, _ *ast.ExprAssignShiftRight) bool {
	return true
}

func (v *Null) ExprBinaryBitwiseAnd(_ int, _ *ast.ExprBinaryBitwiseAnd) bool {
	return true
}


func (v *Null) ExprBinaryBitwiseOr(_ int, _ *ast.ExprBinaryBitwiseOr) bool {
	return true
}


func (v *Null) ExprBinaryBitwiseXor(_ int, _ *ast.ExprBinaryBitwiseXor) bool {
	return true
}


func (v *Null) ExprBinaryBooleanAnd(_ int, _ *ast.ExprBinaryBooleanAnd) bool {
	return true
}


func (v *Null) ExprBinaryBooleanOr(_ int, _ *ast.ExprBinaryBooleanOr) bool {
	return true
}


func (v *Null) ExprBinaryCoalesce(_ int, _ *ast.ExprBinaryCoalesce) bool {
	return true
}


func (v *Null) ExprBinaryConcat(_ int, _ *ast.ExprBinaryConcat) bool {
	return true
}


func (v *Null) ExprBinaryDiv(_ int, _ *ast.ExprBinaryDiv) bool {
	return true
}


func (v *Null) ExprBinaryEqual(_ int, _ *ast.ExprBinaryEqual) bool {
	return true
}


func (v *Null) ExprBinaryGreater(_ int, _ *ast.ExprBinaryGreater) bool {
	return true
}


func (v *Null) ExprBinaryGreaterOrEqual(_ int, _ *ast.ExprBinaryGreaterOrEqual) bool {
	return true
}


func (v *Null) ExprBinaryIdentical(_ int, _ *ast.ExprBinaryIdentical) bool {
	return true
}


func (v *Null) ExprBinaryLogicalAnd(_ int, _ *ast.ExprBinaryLogicalAnd) bool {
	return true
}


func (v *Null) ExprBinaryLogicalOr(_ int, _ *ast.ExprBinaryLogicalOr) bool {
	return true
}


func (v *Null) ExprBinaryLogicalXor(_ int, _ *ast.ExprBinaryLogicalXor) bool {
	return true
}


func (v *Null) ExprBinaryMinus(_ int, _ *ast.ExprBinaryMinus) bool {
	return true
}


func (v *Null) ExprBinaryMod(_ int, _ *ast.ExprBinaryMod) bool {
	return true
}


func (v *Null) ExprBinaryMul(_ int, _ *ast.ExprBinaryMul) bool {
	return true
}


func (v *Null) ExprBinaryNotEqual(_ int, _ *ast.ExprBinaryNotEqual) bool {
	return true
}


func (v *Null) ExprBinaryNotIdentical(_ int, _ *ast.ExprBinaryNotIdentical) bool {
	return true
}


func (v *Null) ExprBinaryPlus(_ int, _ *ast.ExprBinaryPlus) bool {
	return true
}


func (v *Null) ExprBinaryPow(_ int, _ *ast.ExprBinaryPow) bool {
	return true
}


func (v *Null) ExprBinaryShiftLeft(_ int, _ *ast.ExprBinaryShiftLeft) bool {
	return true
}


func (v *Null) ExprBinaryShiftRight(_ int, _ *ast.ExprBinaryShiftRight) bool {
	return true
}


func (v *Null) ExprBinarySmaller(_ int, _ *ast.ExprBinarySmaller) bool {
	return true
}


func (v *Null) ExprBinarySmallerOrEqual(_ int, _ *ast.ExprBinarySmallerOrEqual) bool {
	return true
}


func (v *Null) ExprBinarySpaceship(_ int, _ *ast.ExprBinarySpaceship) bool {
	return true
}


func (v *Null) ExprCastArray(_ int, _ *ast.ExprCastArray) bool {
	return true
}


func (v *Null) ExprCastBool(_ int, _ *ast.ExprCastBool) bool {
	return true
}


func (v *Null) ExprCastDouble(_ int, _ *ast.ExprCastDouble) bool {
	return true
}


func (v *Null) ExprCastInt(_ int, _ *ast.ExprCastInt) bool {
	return true
}


func (v *Null) ExprCastObject(_ int, _ *ast.ExprCastObject) bool {
	return true
}


func (v *Null) ExprCastString(_ int, _ *ast.ExprCastString) bool {
	return true
}


func (v *Null) ExprCastUnset(_ int, _ *ast.ExprCastUnset) bool {
	return true
}

func (v *Null) ScalarDnumber(_ int, _ *ast.ScalarDnumber) bool {
	return true
}


func (v *Null) ScalarEncapsed(_ int, _ *ast.ScalarEncapsed) bool {
	return true
}


func (v *Null) ScalarEncapsedStringPart(_ int, _ *ast.ScalarEncapsedStringPart) bool {
	return true
}


func (v *Null) ScalarHeredoc(_ int, _ *ast.ScalarHeredoc) bool {
	return true
}


func (v *Null) ScalarLnumber(_ int, _ *ast.ScalarLnumber) bool {
	return true
}


func (v *Null) ScalarMagicConstant(_ int, _ *ast.ScalarMagicConstant) bool {
	return true
}


func (v *Null) ScalarString(_ int, _ *ast.ScalarString) bool {
	return true
}


