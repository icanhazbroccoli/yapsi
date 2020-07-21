package test

import (
	"yapsi/pkg/ast"
	"yapsi/pkg/token"
)

func NewToken(t token.TokenType, l string) token.Token {
	return token.Token{
		Type:    t,
		Literal: l,
	}
}

func NewNumber(num string) *ast.NumericLiteral {
	return &ast.NumericLiteral{
		Token: NewToken(token.NUMBER, num),
		Value: ast.RawNumber(num),
	}
}

func NewIdent(ident string) *ast.IdentifierExpr {
	return &ast.IdentifierExpr{
		Token: NewToken(token.IDENT, ident),
		Value: ident,
	}
}

func NewBool(val bool) *ast.BoolLiteral {
	if val {
		return &ast.BoolLiteral{
			Token: NewToken(token.TRUE, "true"),
			Value: true,
		}
	}
	return &ast.BoolLiteral{
		Token: NewToken(token.FALSE, "false"),
		Value: false,
	}
}

func NewString(s string) *ast.StringLiteral {
	return &ast.StringLiteral{
		Token: NewToken(token.STRING, s),
		Value: s,
	}
}

func NewChar(c rune) *ast.CharLiteral {
	return &ast.CharLiteral{
		Token: NewToken(token.CHAR, string(c)),
		Value: c,
	}
}
