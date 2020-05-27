package parser

import (
	"yapsi/pkg/ast"
	"yapsi/pkg/lexer"
	"yapsi/pkg/token"
)

type TestLexer struct {
	tokens []token.Token
	cur    int
}

var _ lexer.Interface = (*TestLexer)(nil)

func NewTestLexer(tokens []token.Token) *TestLexer {
	return &TestLexer{
		tokens: tokens,
	}
}

func (l *TestLexer) NextToken() token.Token {
	if l.cur >= len(l.tokens) {
		return token.Token{
			Type: token.EOF,
		}
	}
	peek := l.tokens[l.cur]
	l.cur++
	return peek
}

func (l *TestLexer) Pos() (int, int) {
	return 0, l.cur
}

func newToken(t token.TokenType, l string) token.Token {
	return token.Token{
		Type:    t,
		Literal: l,
	}
}

func newNumber(num string) *ast.NumericLiteral {
	return &ast.NumericLiteral{
		Token: newToken(token.NUMBER, num),
		Value: ast.RawNumber(num),
	}
}

func newIdent(ident string) *ast.IdentifierExpr {
	return &ast.IdentifierExpr{
		Token: newToken(token.IDENT, ident),
		Value: ident,
	}
}

func newBool(val bool) *ast.BoolLiteral {
	if val {
		return &ast.BoolLiteral{
			Token: newToken(token.TRUE, "true"),
			Value: true,
		}
	}
	return &ast.BoolLiteral{
		Token: newToken(token.FALSE, "false"),
		Value: false,
	}
}
