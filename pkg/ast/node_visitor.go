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

	VisitProgramStmt(*ProgramStmt) (VisitorResult, error)
	VisitBlockStmt(*BlockStmt) (VisitorResult, error)
	VisitCompoundStmt(*CompoundStmt) (VisitorResult, error)
	VisitAssignmentStmt(*AssignmentStmt) (VisitorResult, error)
	VisitGotoStmt(*GotoStmt) (VisitorResult, error)
	VisitLabeledStmt(*LabeledStmt) (VisitorResult, error)
	VisitCallStmt(*CallStmt) (VisitorResult, error)
	VisitIfStmt(*IfStmt) (VisitorResult, error)
	VisitWhileStmt(*WhileStmt) (VisitorResult, error)
	VisitRepeatStmt(*RepeatStmt) (VisitorResult, error)
}
