package ast

type VisitorResult interface{}

type NodeVisitor interface {
	VisitNumericLiteral(*NumericLiteral) VisitorResult
	VisitStringLiteral(*StringLiteral) VisitorResult
	VisitBoolLiteral(*BoolLiteral) VisitorResult
	VisitCharLiteral(*CharLiteral) VisitorResult

	VisitIdentifierExpr(*IdentifierExpr) VisitorResult
	VisitUnaryExpr(*UnaryExpr) VisitorResult
	VisitBinaryExpr(*BinaryExpr) VisitorResult
	VisitElementExpr(*ElementExpr) VisitorResult
	VisitSetExpr(*SetExpr) VisitorResult

	VisitAssignmentStmt(*AssignmentStmt) VisitorResult
	VisitGotoStmt(*GotoStmt) VisitorResult
	VisitLabeledStmt(*LabeledStmt) VisitorResult
	VisitCallStmt(*CallStmt) VisitorResult
	VisitCompoundStmt(*CompoundStmt) VisitorResult
	VisitIfStmt(*IfStmt) VisitorResult
	VisitWhileStmt(*WhileStmt) VisitorResult
}
