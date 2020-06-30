package ast

type VisitorResult interface{}

type NodeVisitor interface {
	VisitNumericLiteral(*NumericLiteral) (VisitorResult, error)
	VisitStringLiteral(*StringLiteral) (VisitorResult, error)
	VisitBoolLiteral(*BoolLiteral) (VisitorResult, error)
	VisitCharLiteral(*CharLiteral) (VisitorResult, error)

	VisitIdentifierExpr(*IdentifierExpr) (VisitorResult, error)
	VisitUnaryExpr(*UnaryExpr) (VisitorResult, error)
	VisitBinaryExpr(*BinaryExpr) (VisitorResult, error)
	VisitElementExpr(*ElementExpr) (VisitorResult, error)
	VisitSetExpr(*SetExpr) (VisitorResult, error)
	VisitFunctionCallExpr(*FunctionCallExpr) (VisitorResult, error)
	VisitSimpleTypeDefinitionExpr(*SimpleTypeDefinitionExpr) (VisitorResult, error)
	VisitSubrangeTypeDefinitionExpr(*SubrangeTypeDefinitionExpr) (VisitorResult, error)
	VisitArrayTypeDefinitionExpr(*ArrayTypeDefinitionExpr) (VisitorResult, error)
	VisitRecordTypeDefinitionExpr(*RecordTypeDefinitionExpr) (VisitorResult, error)

	VisitProgramStmt(*ProgramStmt) (VisitorResult, error)
	VisitBlockStmt(*BlockStmt) (VisitorResult, error)
	VisitCompoundStmt(*CompoundStmt) (VisitorResult, error)
	VisitAssignmentStmt(*AssignmentStmt) (VisitorResult, error)
	VisitGotoStmt(*GotoStmt) (VisitorResult, error)
	VisitLabeledStmt(*LabeledStmt) (VisitorResult, error)
	VisitProcedureDeclStmt(*ProcedureDeclStmt) (VisitorResult, error)
	VisitFunctionDeclStmt(*FunctionDeclStmt) (VisitorResult, error)
	VisitIfStmt(*IfStmt) (VisitorResult, error)
	VisitWhileStmt(*WhileStmt) (VisitorResult, error)
	VisitRepeatStmt(*RepeatStmt) (VisitorResult, error)
	VisitVarDeclStmt(*VarDeclStmt) (VisitorResult, error)
	VisitProcedureCallStmt(*ProcedureCallStmt) (VisitorResult, error)
	VisitReturnStmt(*ReturnStmt) (VisitorResult, error)
	VisitTypeDeclStmt(*TypeDeclStmt) (VisitorResult, error)
	VisitTypeDefinitionStmt(*TypeDefinitionStmt) (VisitorResult, error)
}
