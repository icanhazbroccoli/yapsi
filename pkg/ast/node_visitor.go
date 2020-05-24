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
}
