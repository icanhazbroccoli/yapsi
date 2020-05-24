package ast

import "yapsi/pkg/token"

type NumericLiteral struct {
	Token token.Token
	Value RawNumber
}

var _ Literal = (*NumericLiteral)(nil)
var _ Expression = (*NumericLiteral)(nil)

func (i *NumericLiteral) expressionNode()                   {}
func (i *NumericLiteral) Visit(v NodeVisitor) VisitorResult { return v.VisitNumericLiteral(i) }
func (i *NumericLiteral) TokenLiteral() string              { return i.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

var _ Literal = (*StringLiteral)(nil)
var _ Expression = (*StringLiteral)(nil)

func (s *StringLiteral) expressionNode()                   {}
func (s *StringLiteral) Visit(v NodeVisitor) VisitorResult { return v.VisitStringLiteral(s) }
func (s *StringLiteral) TokenLiteral() string              { return s.Token.Literal }

type BoolLiteral struct {
	Token token.Token
	Value bool
}

var _ Literal = (*BoolLiteral)(nil)
var _ Expression = (*BoolLiteral)(nil)

func (b *BoolLiteral) expressionNode()                   {}
func (b *BoolLiteral) Visit(v NodeVisitor) VisitorResult { return v.VisitBoolLiteral(b) }
func (b *BoolLiteral) TokenLiteral() string              { return b.Token.Literal }

type CharLiteral struct {
	Token token.Token
	Value rune
}

var _ Literal = (*CharLiteral)(nil)
var _ Expression = (*CharLiteral)(nil)

func (c *CharLiteral) expressionNode()                   {}
func (c *CharLiteral) Visit(v NodeVisitor) VisitorResult { return v.VisitCharLiteral(c) }
func (c *CharLiteral) TokenLiteral() string              { return c.Token.Literal }
