package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestParseVariables(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantVars []ast.Variable
		wantErr  error
	}{
		{
			input:    []token.Token{},
			wantVars: []ast.Variable{},
			wantErr:  nil,
		},
		{
			input: []token.Token{
				newToken(token.VAR, "var"),
				newToken(token.IDENT, "foo"),
				newToken(token.COLON, ":"),
				newToken(token.IDENT, "bar"),
				newToken(token.SEMICOLON, ";"),
			},
			wantVars: []ast.Variable{
				ast.Variable{
					Identifier: ast.VarIdentifier("foo"),
					Type:       ast.TypeIdentifier("bar"),
				},
			},
			wantErr: nil,
		},
		{
			input: []token.Token{
				newToken(token.VAR, "var"),
				newToken(token.IDENT, "foo"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "bar"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "baz"),
				newToken(token.COLON, ":"),
				newToken(token.IDENT, "boo"),
				newToken(token.SEMICOLON, ";"),
			},
			wantVars: []ast.Variable{
				ast.Variable{
					Identifier: ast.VarIdentifier("foo"),
					Type:       ast.TypeIdentifier("boo"),
				},
				ast.Variable{
					Identifier: ast.VarIdentifier("bar"),
					Type:       ast.TypeIdentifier("boo"),
				},
				ast.Variable{
					Identifier: ast.VarIdentifier("baz"),
					Type:       ast.TypeIdentifier("boo"),
				},
			},
			wantErr: nil,
		},
		{
			input: []token.Token{
				newToken(token.VAR, "var"),
				newToken(token.IDENT, "foo"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "bar"),
				newToken(token.COLON, ":"),
				newToken(token.IDENT, "Integer"),
				newToken(token.SEMICOLON, ";"),
				newToken(token.IDENT, "baz"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "boo"),
				newToken(token.COLON, ":"),
				newToken(token.IDENT, "Real"),
				newToken(token.SEMICOLON, ";"),
			},
			wantVars: []ast.Variable{
				ast.Variable{
					Identifier: ast.VarIdentifier("foo"),
					Type:       ast.TypeIdentifier("Integer"),
				},
				ast.Variable{
					Identifier: ast.VarIdentifier("bar"),
					Type:       ast.TypeIdentifier("Integer"),
				},
				ast.Variable{
					Identifier: ast.VarIdentifier("baz"),
					Type:       ast.TypeIdentifier("Real"),
				},
				ast.Variable{
					Identifier: ast.VarIdentifier("boo"),
					Type:       ast.TypeIdentifier("Real"),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		vars, err := p.parseVariables()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, vars, tt.wantVars)
	}
}
