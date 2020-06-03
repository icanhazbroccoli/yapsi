package ast

import "yapsi/pkg/token"

type NumericLiteral struct {
	Token token.Token
	Value RawNumber
}

var _ Literal = (*NumericLiteral)(nil)
var _ Expression = (*NumericLiteral)(nil)

func (i *NumericLiteral) expressionNode()      {}
func (i *NumericLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *NumericLiteral) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitNumericLiteral(i)
}

type StringLiteral struct {
	Token token.Token
	Value string
}

var _ Literal = (*StringLiteral)(nil)
var _ Expression = (*StringLiteral)(nil)

func (s *StringLiteral) expressionNode()      {}
func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *StringLiteral) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitStringLiteral(s)
}

type BoolLiteral struct {
	Token token.Token
	Value bool
}

var _ Literal = (*BoolLiteral)(nil)
var _ Expression = (*BoolLiteral)(nil)

func (b *BoolLiteral) expressionNode()      {}
func (b *BoolLiteral) TokenLiteral() string { return b.Token.Literal }
func (b *BoolLiteral) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitBoolLiteral(b)
}

type CharLiteral struct {
	Token token.Token
	Value rune
}

var _ Literal = (*CharLiteral)(nil)
var _ Expression = (*CharLiteral)(nil)

func (c *CharLiteral) expressionNode()      {}
func (c *CharLiteral) TokenLiteral() string { return c.Token.Literal }
func (c *CharLiteral) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitCharLiteral(c)
}
