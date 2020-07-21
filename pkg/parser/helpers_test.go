package parser

import (
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
