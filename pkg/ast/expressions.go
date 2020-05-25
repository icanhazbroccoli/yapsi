package ast

import "yapsi/pkg/token"

type IdentifierExpr struct {
	Token token.Token
	Value string
}

var _ Expression = (*IdentifierExpr)(nil)

func (i *IdentifierExpr) expressionNode()                   {}
func (i *IdentifierExpr) Visit(v NodeVisitor) VisitorResult { return v.VisitIdentifierExpr(i) }

type UnaryExpr struct {
	Operator token.Token
	Expr     Expression
}

var _ Expression = (*UnaryExpr)(nil)

func (u *UnaryExpr) expressionNode()                   {}
func (u *UnaryExpr) Visit(v NodeVisitor) VisitorResult { return v.VisitUnaryExpr(u) }

type BinaryExpr struct {
	Left, Right Expression
	Operator    token.Token
}

var _ Expression = (*BinaryExpr)(nil)

func (b *BinaryExpr) expressionNode()                   {}
func (b *BinaryExpr) Visit(v NodeVisitor) VisitorResult { return v.VisitBinaryExpr(b) }

type ElementExpr struct {
	Left, Right Expression
}

var _ Expression = (*ElementExpr)(nil)

func (e *ElementExpr) expressionNode()                   {}
func (e *ElementExpr) Visit(v NodeVisitor) VisitorResult { return v.VisitElementExpr(e) }

type SetExpr struct {
	Elements []Expression
}

var _ Expression = (*SetExpr)(nil)

func (s *SetExpr) expressionNode()                   {}
func (s *SetExpr) Visit(v NodeVisitor) VisitorResult { return v.VisitSetExpr(s) }
