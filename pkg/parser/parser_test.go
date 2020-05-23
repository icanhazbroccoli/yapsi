package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"yapsi/pkg/ast"
	"yapsi/pkg/token"
)

func TestParseProgramHeader(t *testing.T) {
	tests := []struct {
		input          []token.Token
		wantIdentifier ast.ProgramIdentifier
		wantParams     []ast.Identifier
		wantErr        error
	}{
		{
			input: []token.Token{
				newToken(token.PROGRAM, "program"),
				newToken(token.IDENT, "foo"),
				newToken(token.SEMICOLON, ";"),
			},
			wantIdentifier: ast.ProgramIdentifier("foo"),
			wantParams:     []ast.Identifier{},
		},
		{
			input: []token.Token{
				newToken(token.PROGRAM, "program"),
				newToken(token.IDENT, "foo"),
				newToken(token.LPAREN, "("),
				newToken(token.IDENT, "output"),
				newToken(token.RPAREN, ")"),
				newToken(token.SEMICOLON, ";"),
			},
			wantIdentifier: ast.ProgramIdentifier("foo"),
			wantParams: []ast.Identifier{
				ast.Identifier("output"),
			},
		},
		{
			input: []token.Token{
				newToken(token.PROGRAM, "program"),
				newToken(token.IDENT, "foo"),
				newToken(token.LPAREN, "("),
				newToken(token.IDENT, "infile1"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "infile2"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "mergedfile"),
				newToken(token.RPAREN, ")"),
				newToken(token.SEMICOLON, ";"),
			},
			wantIdentifier: ast.ProgramIdentifier("foo"),
			wantParams: []ast.Identifier{
				ast.Identifier("infile1"),
				ast.Identifier("infile2"),
				ast.Identifier("mergedfile"),
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		identifier, params, err := p.parseProgramHeader()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantIdentifier, identifier)
		assert.Equal(t, tt.wantParams, params)
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

func TestParseLabels(t *testing.T) {
	tests := []struct {
		input      []token.Token
		wantLabels []ast.Label
		wantErr    error
	}{
		{
			input:      []token.Token{},
			wantLabels: []ast.Label{},
		},
		{
			input: []token.Token{
				newToken(token.LABEL, "label"),
				newToken(token.IDENT, "foo"),
				newToken(token.SEMICOLON, ";"),
			},
			wantLabels: []ast.Label{
				ast.Label{Identifier: "foo"},
			},
		},
		{
			input: []token.Token{
				newToken(token.LABEL, "label"),
				newToken(token.IDENT, "foo"),
				newToken(token.COMMA, ","),
				newToken(token.NUMBER, "12345"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "lbl123"),
				newToken(token.SEMICOLON, ";"),
			},
			wantLabels: []ast.Label{
				ast.Label{Identifier: "foo"},
				ast.Label{Identifier: "12345"},
				ast.Label{Identifier: "lbl123"},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		labels, err := p.parseLabels()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, labels, tt.wantLabels)
	}
}
