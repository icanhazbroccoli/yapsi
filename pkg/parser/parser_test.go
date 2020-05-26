package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"yapsi/pkg/ast"
	"yapsi/pkg/token"
)

//func TestParseProgramHeader(t *testing.T) {
//	tests := []struct {
//		input          []token.Token
//		wantIdentifier ast.ProgramIdentifier
//		wantParams     []ast.Identifier
//		wantErr        error
//	}{
//		{
//			input: []token.Token{
//				newToken(token.PROGRAM, "program"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantIdentifier: ast.ProgramIdentifier("foo"),
//			wantParams:     []ast.Identifier{},
//		},
//		{
//			input: []token.Token{
//				newToken(token.PROGRAM, "program"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.LPAREN, "("),
//				newToken(token.IDENT, "output"),
//				newToken(token.RPAREN, ")"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantIdentifier: ast.ProgramIdentifier("foo"),
//			wantParams: []ast.Identifier{
//				ast.Identifier("output"),
//			},
//		},
//		{
//			input: []token.Token{
//				newToken(token.PROGRAM, "program"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.LPAREN, "("),
//				newToken(token.IDENT, "infile1"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "infile2"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "mergedfile"),
//				newToken(token.RPAREN, ")"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantIdentifier: ast.ProgramIdentifier("foo"),
//			wantParams: []ast.Identifier{
//				ast.Identifier("infile1"),
//				ast.Identifier("infile2"),
//				ast.Identifier("mergedfile"),
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		l := NewTestLexer(tt.input)
//		p := New(l)
//		identifier, params, err := p.parseProgramHeader()
//		assert.Equal(t, err, tt.wantErr)
//		if err != nil {
//			continue
//		}
//		assert.Equal(t, tt.wantIdentifier, identifier)
//		assert.Equal(t, tt.wantParams, params)
//	}
//}
//
//func TestParseVariables(t *testing.T) {
//	tests := []struct {
//		input    []token.Token
//		wantVars []ast.Variable
//		wantErr  error
//	}{
//		{
//			input:    []token.Token{},
//			wantVars: []ast.Variable{},
//			wantErr:  nil,
//		},
//		{
//			input: []token.Token{
//				newToken(token.VAR, "var"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.COLON, ":"),
//				newToken(token.IDENT, "bar"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantVars: []ast.Variable{
//				ast.Variable{
//					Identifier: ast.VarIdentifier("foo"),
//					Type:       ast.TypeIdentifier("bar"),
//				},
//			},
//			wantErr: nil,
//		},
//		{
//			input: []token.Token{
//				newToken(token.VAR, "var"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "bar"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "baz"),
//				newToken(token.COLON, ":"),
//				newToken(token.IDENT, "boo"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantVars: []ast.Variable{
//				ast.Variable{
//					Identifier: ast.VarIdentifier("foo"),
//					Type:       ast.TypeIdentifier("boo"),
//				},
//				ast.Variable{
//					Identifier: ast.VarIdentifier("bar"),
//					Type:       ast.TypeIdentifier("boo"),
//				},
//				ast.Variable{
//					Identifier: ast.VarIdentifier("baz"),
//					Type:       ast.TypeIdentifier("boo"),
//				},
//			},
//			wantErr: nil,
//		},
//		{
//			input: []token.Token{
//				newToken(token.VAR, "var"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "bar"),
//				newToken(token.COLON, ":"),
//				newToken(token.IDENT, "Integer"),
//				newToken(token.SEMICOLON, ";"),
//				newToken(token.IDENT, "baz"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "boo"),
//				newToken(token.COLON, ":"),
//				newToken(token.IDENT, "Real"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantVars: []ast.Variable{
//				ast.Variable{
//					Identifier: ast.VarIdentifier("foo"),
//					Type:       ast.TypeIdentifier("Integer"),
//				},
//				ast.Variable{
//					Identifier: ast.VarIdentifier("bar"),
//					Type:       ast.TypeIdentifier("Integer"),
//				},
//				ast.Variable{
//					Identifier: ast.VarIdentifier("baz"),
//					Type:       ast.TypeIdentifier("Real"),
//				},
//				ast.Variable{
//					Identifier: ast.VarIdentifier("boo"),
//					Type:       ast.TypeIdentifier("Real"),
//				},
//			},
//			wantErr: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		l := NewTestLexer(tt.input)
//		p := New(l)
//		vars, err := p.parseVariables()
//		assert.Equal(t, err, tt.wantErr)
//		if err != nil {
//			continue
//		}
//		assert.Equal(t, vars, tt.wantVars)
//	}
//}
//
//func TestParseLabels(t *testing.T) {
//	tests := []struct {
//		input      []token.Token
//		wantLabels []ast.Label
//		wantErr    error
//	}{
//		{
//			input:      []token.Token{},
//			wantLabels: []ast.Label{},
//		},
//		{
//			input: []token.Token{
//				newToken(token.LABEL, "label"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantLabels: []ast.Label{
//				ast.Label{Identifier: "foo"},
//			},
//		},
//		{
//			input: []token.Token{
//				newToken(token.LABEL, "label"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.COMMA, ","),
//				newToken(token.NUMBER, "12345"),
//				newToken(token.COMMA, ","),
//				newToken(token.IDENT, "lbl123"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantLabels: []ast.Label{
//				ast.Label{Identifier: "foo"},
//				ast.Label{Identifier: "12345"},
//				ast.Label{Identifier: "lbl123"},
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		l := NewTestLexer(tt.input)
//		p := New(l)
//		labels, err := p.parseLabels()
//		assert.Equal(t, err, tt.wantErr)
//		if err != nil {
//			continue
//		}
//		assert.Equal(t, labels, tt.wantLabels)
//	}
//}

//func TestParseConstants(t *testing.T) {
//	tests := []struct {
//		input         []token.Token
//		wantConstants []ast.Constant
//		wantErr       error
//	}{
//		{
//			input:         []token.Token{},
//			wantConstants: []ast.Constant{},
//		},
//		{
//			input: []token.Token{
//				newToken(token.CONST, "const"),
//				newToken(token.IDENT, "foo"),
//				newToken(token.EQUAL, "="),
//				newToken(token.NUMBER, "12345"),
//				newToken(token.SEMICOLON, ";"),
//			},
//			wantConstants: []ast.Constant{
//				ast.Constant{
//					Identifier: ast.ConstIdentifier("foo"),
//					Raw:        newToken(token.NUMBER, "12345"),
//				},
//			},
//		},
//		{
//			input: []token.Token{
//				newToken(token.CONST, "const"),
//				newToken(token.IDENT, "const_number"),
//				newToken(token.EQUAL, "="),
//				newToken(token.NUMBER, "123.45"),
//				newToken(token.SEMICOLON, ";"),
//				newToken(token.IDENT, "const_set"),
//				newToken(token.EQUAL, "="),
//				newToken(token.NIL, "nil"),
//				newToken(token.SEMICOLON, ";"),
//				newToken(token.IDENT, "const_char"),
//				newToken(toke.EQUAL, "="),
//				newToken(token.CHAR, "'a'"),
//				newToken(token.SEMICOLON, ";"),
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		l := NewTestLexer(tt.input)
//		p := New(l)
//		constants, err := p.parseConstants()
//		assert.Equal(t, err, tt.wantErr)
//		if err != nil {
//			continue
//		}
//		assert.Equal(t, constants, tt.wantConstants)
//	}
//}

func TestParseFactor(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantExpr ast.Expression
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.NUMBER, "12345"),
			},
			wantExpr: &ast.NumericLiteral{
				Token: newToken(token.NUMBER, "12345"),
				Value: ast.RawNumber("12345"),
			},
		},
		{
			input: []token.Token{
				newToken(token.STRING, "Hello World!"),
			},
			wantExpr: &ast.StringLiteral{
				Token: newToken(token.STRING, "Hello World!"),
				Value: "Hello World!",
			},
		},
		{
			input: []token.Token{
				newToken(token.CHAR, "я"),
			},
			wantExpr: &ast.CharLiteral{
				Token: newToken(token.CHAR, "я"),
				Value: 'я',
			},
		},
		{
			input: []token.Token{
				newToken(token.TRUE, "true"),
			},
			wantExpr: &ast.BoolLiteral{
				Token: newToken(token.TRUE, "true"),
				Value: true,
			},
		},
		{
			input: []token.Token{
				newToken(token.NOT, "not"),
				newToken(token.FALSE, "false"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: newToken(token.NOT, "not"),
				Expr: &ast.BoolLiteral{
					Token: newToken(token.FALSE, "false"),
					Value: false,
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.LBRACKET, "["),
				newToken(token.NUMBER, "1"),
				newToken(token.COMMA, ","),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
				newToken(token.COMMA, ","),
				newToken(token.NUMBER, "1"),
				newToken(token.DOTDOT, ".."),
				newToken(token.NUMBER, "10"),
				newToken(token.RBRACKET, "]"),
			},
			wantExpr: &ast.SetExpr{
				Elements: []ast.Expression{
					&ast.NumericLiteral{
						Token: newToken(token.NUMBER, "1"),
						Value: ast.RawNumber("1"),
					},
					&ast.BinaryExpr{
						Left: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
						Operator: newToken(token.ASTERISK, "*"),
						Right: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
					},
					&ast.ElementExpr{
						Left: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "1"),
							Value: ast.RawNumber("1"),
						},
						Right: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "10"),
							Value: ast.RawNumber("10"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseFactor()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, expr, tt.wantExpr)
	}
}

func TestParseTermExpr(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantExpr ast.Expression
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
			},
			wantExpr: &ast.NumericLiteral{
				Token: newToken(token.NUMBER, "1"),
				Value: ast.RawNumber("1"),
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "1"),
					Value: ast.RawNumber("1"),
				},
				Operator: newToken(token.ASTERISK, "*"),
				Right: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "2"),
					Value: ast.RawNumber("2"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "2"),
				newToken(token.DIV, "div"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "2"),
					Value: ast.RawNumber("2"),
				},
				Operator: newToken(token.DIV, "div"),
				Right: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "3"),
					Value: ast.RawNumber("3"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "2"),
				newToken(token.MOD, "mod"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "2"),
					Value: ast.RawNumber("2"),
				},
				Operator: newToken(token.MOD, "mod"),
				Right: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "3"),
					Value: ast.RawNumber("3"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
				newToken(token.SLASH, "/"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "1"),
					Value: ast.RawNumber("1"),
				},
				Operator: newToken(token.ASTERISK, "*"),
				Right: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "2"),
						Value: ast.RawNumber("2"),
					},
					Operator: newToken(token.SLASH, "/"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "3"),
						Value: ast.RawNumber("3"),
					},
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.TRUE, "true"),
				newToken(token.AND, "and"),
				newToken(token.FALSE, "false"),
				newToken(token.AND, "and"),
				newToken(token.TRUE, "true"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Token: newToken(token.TRUE, "true"),
					Value: true,
				},
				Operator: newToken(token.AND, "and"),
				Right: &ast.BinaryExpr{
					Left: &ast.BoolLiteral{
						Token: newToken(token.FALSE, "false"),
						Value: false,
					},
					Operator: newToken(token.AND, "and"),
					Right: &ast.BoolLiteral{
						Token: newToken(token.TRUE, "true"),
						Value: true,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseTerm()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, expr, tt.wantExpr)
	}
}

func TestParseSimpleExpression(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantExpr ast.Expression
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
			},
			wantExpr: &ast.NumericLiteral{
				Token: newToken(token.NUMBER, "1"),
				Value: ast.RawNumber("1"),
			},
		},
		{
			input: []token.Token{
				newToken(token.MINUS, "-"),
				newToken(token.NUMBER, "42"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: newToken(token.MINUS, "-"),
				Expr: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "42"),
					Value: ast.RawNumber("42"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.PLUS, "+"),
				newToken(token.NUMBER, "42"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: newToken(token.PLUS, "+"),
				Expr: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "42"),
					Value: ast.RawNumber("42"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.PLUS, "+"),
				newToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "1"),
					Value: ast.RawNumber("1"),
				},
				Operator: newToken(token.PLUS, "+"),
				Right: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "2"),
					Value: ast.RawNumber("2"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.MINUS, "-"),
				newToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "1"),
					Value: ast.RawNumber("1"),
				},
				Operator: newToken(token.MINUS, "-"),
				Right: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "2"),
					Value: ast.RawNumber("2"),
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.TRUE, "true"),
				newToken(token.OR, "or"),
				newToken(token.FALSE, "false"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BoolLiteral{
					Token: newToken(token.TRUE, "true"),
					Value: true,
				},
				Operator: newToken(token.OR, "or"),
				Right: &ast.BoolLiteral{
					Token: newToken(token.FALSE, "false"),
					Value: false,
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.PLUS, "+"),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.NumericLiteral{
					Token: newToken(token.NUMBER, "1"),
					Value: ast.RawNumber("1"),
				},
				Operator: newToken(token.PLUS, "+"),
				Right: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "2"),
						Value: ast.RawNumber("2"),
					},
					Operator: newToken(token.ASTERISK, "*"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "3"),
						Value: ast.RawNumber("3"),
					},
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.NUMBER, "1"),
				newToken(token.SLASH, "/"),
				newToken(token.NUMBER, "4"),
				newToken(token.PLUS, "+"),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "1"),
						Value: ast.RawNumber("1"),
					},
					Operator: newToken(token.SLASH, "/"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "4"),
						Value: ast.RawNumber("4"),
					},
				},
				Operator: newToken(token.PLUS, "+"),
				Right: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "2"),
						Value: ast.RawNumber("2"),
					},
					Operator: newToken(token.ASTERISK, "*"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "3"),
						Value: ast.RawNumber("3"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseSimpleExpression()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, expr, tt.wantExpr)
	}
}

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantExpr ast.Expression
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
				newToken(token.GTEQL, ">="),
				newToken(token.NUMBER, "1"),
				newToken(token.PLUS, "+"),
				newToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "2"),
						Value: ast.RawNumber("2"),
					},
					Operator: newToken(token.ASTERISK, "*"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "2"),
						Value: ast.RawNumber("2"),
					},
				},
				Operator: newToken(token.GTEQL, ">="),
				Right: &ast.BinaryExpr{
					Left: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "1"),
						Value: ast.RawNumber("1"),
					},
					Operator: newToken(token.PLUS, "+"),
					Right: &ast.NumericLiteral{
						Token: newToken(token.NUMBER, "3"),
						Value: ast.RawNumber("3"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseExpression()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, expr, tt.wantExpr)
	}
}

func TestParseSimpleStmt(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantStmt ast.Statement
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.IDENT, "foo"),
				newToken(token.COLON, ":"),
				newToken(token.IDENT, "a"),
				newToken(token.NAMED, ":="),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
			},
			wantStmt: &ast.LabeledStmt{
				Label: &ast.IdentifierExpr{
					Token: newToken(token.IDENT, "foo"),
					Value: "foo",
				},
				Stmt: &ast.AssignmentStmt{
					Identifier: &ast.IdentifierExpr{
						Token: newToken(token.IDENT, "a"),
						Value: "a",
					},
					Expr: &ast.BinaryExpr{
						Left: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
						Operator: newToken(token.ASTERISK, "*"),
						Right: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				newToken(token.IDENT, "foo"),
			},
			wantStmt: &ast.CallStmt{
				Identifier: &ast.IdentifierExpr{
					Token: newToken(token.IDENT, "foo"),
					Value: "foo",
				},
				Args: []ast.Expression{},
			},
		},
		{
			input: []token.Token{
				newToken(token.IDENT, "foo"),
				newToken(token.LPAREN, "("),
				newToken(token.NUMBER, "42"),
				newToken(token.COMMA, ","),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
				newToken(token.COMMA, ","),
				newToken(token.IDENT, "bar"),
				newToken(token.RPAREN, ")"),
			},
			wantStmt: &ast.CallStmt{
				Identifier: &ast.IdentifierExpr{
					Token: newToken(token.IDENT, "foo"),
					Value: "foo",
				},
				Args: []ast.Expression{
					&ast.NumericLiteral{
						Token: newToken(token.NUMBER, "42"),
						Value: ast.RawNumber("42"),
					},
					&ast.BinaryExpr{
						Left: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
						Operator: newToken(token.ASTERISK, "*"),
						Right: &ast.NumericLiteral{
							Token: newToken(token.NUMBER, "2"),
							Value: ast.RawNumber("2"),
						},
					},
					&ast.IdentifierExpr{
						Token: newToken(token.IDENT, "bar"),
						Value: "bar",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		stmt, err := p.parseSimpleStmt()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, stmt, tt.wantStmt)
	}
}

func TestParseStmt(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantStmt ast.Statement
		wantErr  error
	}{
		{
			input: []token.Token{
				newToken(token.BEGIN, "begin"),
				newToken(token.IDENT, "foo"),
				newToken(token.NAMED, ":="),
				newToken(token.NUMBER, "2"),
				newToken(token.ASTERISK, "*"),
				newToken(token.NUMBER, "2"),
				newToken(token.SEMICOLON, ";"),
				newToken(token.IDENT, "bar"),
				newToken(token.LPAREN, "("),
				newToken(token.NUMBER, "42"),
				newToken(token.COMMA, ","),
				newToken(token.MINUS, "-"),
				newToken(token.NUMBER, "1"),
				newToken(token.RPAREN, ")"),
				newToken(token.END, "end"),
			},
			wantStmt: &ast.CompoundStmt{
				Token: newToken(token.BEGIN, "begin"),
				Statements: []ast.Statement{
					&ast.AssignmentStmt{
						Identifier: &ast.IdentifierExpr{
							Token: newToken(token.IDENT, "foo"),
							Value: "foo",
						},
						Expr: &ast.BinaryExpr{
							Left: &ast.NumericLiteral{
								Token: newToken(token.NUMBER, "2"),
								Value: ast.RawNumber("2"),
							},
							Operator: newToken(token.ASTERISK, "*"),
							Right: &ast.NumericLiteral{
								Token: newToken(token.NUMBER, "2"),
								Value: ast.RawNumber("2"),
							},
						},
					},
					&ast.CallStmt{
						Identifier: &ast.IdentifierExpr{
							Token: newToken(token.IDENT, "bar"),
							Value: "bar",
						},
						Args: []ast.Expression{
							&ast.NumericLiteral{
								Token: newToken(token.NUMBER, "42"),
								Value: ast.RawNumber("42"),
							},
							&ast.UnaryExpr{
								Operator: newToken(token.MINUS, "-"),
								Expr: &ast.NumericLiteral{
									Token: newToken(token.NUMBER, "1"),
									Value: ast.RawNumber("1"),
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		stmt, err := p.parseStmt()
		assert.Equal(t, err, tt.wantErr)
		if err != nil {
			continue
		}
		assert.Equal(t, stmt, tt.wantStmt)
	}
}
