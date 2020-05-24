package ast

import (
	"bytes"
	"fmt"
)

type AstPrinter struct{}

var _ NodeVisitor = (*AstPrinter)(nil)

func (p *AstPrinter) VisitNumericLiteral(node *NumericLiteral) VisitorResult {
	return fmt.Sprintf("%s", node.Value)
}

func (p *AstPrinter) VisitStringLiteral(node *StringLiteral) VisitorResult {
	return node.Value
}

func (p *AstPrinter) VisitBoolLiteral(node *BoolLiteral) VisitorResult {
	return fmt.Sprintf("%t", node.Value)
}

func (p *AstPrinter) VisitCharLiteral(node *CharLiteral) VisitorResult {
	return fmt.Sprintf("%c", node.Value)
}

func (p *AstPrinter) VisitIdentifierExpr(node *IdentifierExpr) VisitorResult {
	return node.Value
}

func (p *AstPrinter) VisitUnaryExpr(node *UnaryExpr) VisitorResult {
	var out bytes.Buffer
	right := node.Expr.Visit(p)
	out.WriteString(node.Operator.Literal + " ")
	out.WriteString(right.(string))
	return out.String()
}

func (p *AstPrinter) VisitBinaryExpr(node *BinaryExpr) VisitorResult {
	var out bytes.Buffer
	left := node.Left.Visit(p)
	right := node.Right.Visit(p)
	out.WriteString(left.(string))
	out.WriteString(" " + node.Operator.Literal + " ")
	out.WriteString(right.(string))
	return out.String()
}
