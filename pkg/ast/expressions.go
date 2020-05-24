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

type ParamExpr struct {
	Token token.Token
	Value string
}

func (p *ParamExpr) expressionNode()      {}
func (p *ParamExpr) TokenLiteral() string { return p.Token.Literal }
func (p *ParamExpr) String() string       { return p.Token.Literal }
