package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"yapsi/pkg/ast"
	test_helper "yapsi/pkg/internal/test"
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
//			test_helper.NewToken(token.PROGRAM, "program"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			},
//			wantIdentifier: ast.ProgramIdentifier("foo"),
//			wantParams:     []ast.Identifier{},
//		},
//		{
//			input: []token.Token{
//			test_helper.NewToken(token.PROGRAM, "program"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.LPAREN, "("),
//			test_helper.NewToken(token.IDENT, "output"),
//			test_helper.NewToken(token.RPAREN, ")"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			},
//			wantIdentifier: ast.ProgramIdentifier("foo"),
//			wantParams: []ast.Identifier{
//				ast.Identifier("output"),
//			},
//		},
//		{
//			input: []token.Token{
//			test_helper.NewToken(token.PROGRAM, "program"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.LPAREN, "("),
//			test_helper.NewToken(token.IDENT, "infile1"),
//			test_helper.NewToken(token.COMMA, ","),
//			test_helper.NewToken(token.IDENT, "infile2"),
//			test_helper.NewToken(token.COMMA, ","),
//			test_helper.NewToken(token.IDENT, "mergedfile"),
//			test_helper.NewToken(token.RPAREN, ")"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
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

func TestParseVarDeclStmt(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantDecl *ast.VarDeclStmt
		wantErr  error
	}{
		{
			input:    []token.Token{},
			wantDecl: nil,
			wantErr:  nil,
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.VAR, "var"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantDecl: &ast.VarDeclStmt{
				Token: test_helper.NewToken(token.VAR, "var"),
				Declarations: map[string]ast.TypeDefinitionExprIntf{
					"foo": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "bar"),
						Identifier: test_helper.NewIdent("bar"),
					},
				},
			},
			wantErr: nil,
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.VAR, "var"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.NUMBER, "0"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "10"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantDecl: &ast.VarDeclStmt{
				Token: test_helper.NewToken(token.VAR, "var"),
				Declarations: map[string]ast.TypeDefinitionExprIntf{
					"foo": &ast.SubrangeTypeDefinitionExpr{
						Token: test_helper.NewToken(token.NUMBER, "0"),
						Left:  test_helper.NewNumber("0"),
						Right: test_helper.NewNumber("10"),
					},
				},
			},
			wantErr: nil,
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.VAR, "var"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "baz"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "boo"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantDecl: &ast.VarDeclStmt{
				Token: test_helper.NewToken(token.VAR, "var"),
				Declarations: map[string]ast.TypeDefinitionExprIntf{
					"foo": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "boo"),
						Identifier: test_helper.NewIdent("boo"),
					},
					"bar": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "boo"),
						Identifier: test_helper.NewIdent("boo"),
					},
					"baz": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "boo"),
						Identifier: test_helper.NewIdent("boo"),
					},
				},
			},
			wantErr: nil,
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.VAR, "var"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "Integer"),
				test_helper.NewToken(token.SEMICOLON, ";"),
				test_helper.NewToken(token.IDENT, "baz"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "boo"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "Real"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantDecl: &ast.VarDeclStmt{
				Token: test_helper.NewToken(token.VAR, "var"),
				Declarations: map[string]ast.TypeDefinitionExprIntf{
					"foo": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "Integer"),
						Identifier: test_helper.NewIdent("Integer"),
					},
					"bar": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "Integer"),
						Identifier: test_helper.NewIdent("Integer"),
					},
					"baz": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "Real"),
						Identifier: test_helper.NewIdent("Real"),
					},
					"boo": &ast.SimpleTypeDefinitionExpr{
						Token:      test_helper.NewToken(token.IDENT, "Real"),
						Identifier: test_helper.NewIdent("Real"),
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		decl, err := p.parseVarDeclStmt()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantDecl, decl)
	}
}

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
//			test_helper.NewToken(token.LABEL, "label"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			},
//			wantLabels: []ast.Label{
//				ast.Label{Identifier: "foo"},
//			},
//		},
//		{
//			input: []token.Token{
//			test_helper.NewToken(token.LABEL, "label"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.COMMA, ","),
//			test_helper.NewToken(token.NUMBER, "12345"),
//			test_helper.NewToken(token.COMMA, ","),
//			test_helper.NewToken(token.IDENT, "lbl123"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
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
//			test_helper.NewToken(token.CONST, "const"),
//			test_helper.NewToken(token.IDENT, "foo"),
//			test_helper.NewToken(token.EQUAL, "="),
//			test_helper.NewToken(token.NUMBER, "12345"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			},
//			wantConstants: []ast.Constant{
//				ast.Constant{
//					Identifier: ast.ConstIdentifier("foo"),
//					Raw:       test_helper.NewToken(token.NUMBER, "12345"),
//				},
//			},
//		},
//		{
//			input: []token.Token{
//			test_helper.NewToken(token.CONST, "const"),
//			test_helper.NewToken(token.IDENT, "const_number"),
//			test_helper.NewToken(token.EQUAL, "="),
//			test_helper.NewToken(token.NUMBER, "123.45"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			test_helper.NewToken(token.IDENT, "const_set"),
//			test_helper.NewToken(token.EQUAL, "="),
//			test_helper.NewToken(token.NIL, "nil"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
//			test_helper.NewToken(token.IDENT, "const_char"),
//			test_helper.NewToken(toke.EQUAL, "="),
//			test_helper.NewToken(token.CHAR, "'a'"),
//			test_helper.NewToken(token.SEMICOLON, ";"),
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
				test_helper.NewToken(token.NUMBER, "12345"),
			},
			wantExpr: &ast.NumericLiteral{
				Token: test_helper.NewToken(token.NUMBER, "12345"),
				Value: ast.RawNumber("12345"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.STRING, "Hello World!"),
			},
			wantExpr: &ast.StringLiteral{
				Token: test_helper.NewToken(token.STRING, "Hello World!"),
				Value: "Hello World!",
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.CHAR, "я"),
			},
			wantExpr: &ast.CharLiteral{
				Token: test_helper.NewToken(token.CHAR, "я"),
				Value: 'я',
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TRUE, "true"),
			},
			wantExpr: &ast.BoolLiteral{
				Token: test_helper.NewToken(token.TRUE, "true"),
				Value: true,
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NOT, "not"),
				test_helper.NewToken(token.FALSE, "false"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: test_helper.NewToken(token.NOT, "not"),
				Expr: &ast.BoolLiteral{
					Token: test_helper.NewToken(token.FALSE, "false"),
					Value: false,
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.LBRACKET, "["),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "10"),
				test_helper.NewToken(token.RBRACKET, "]"),
			},
			wantExpr: &ast.SetExpr{
				Elements: []ast.Expression{
					test_helper.NewNumber("1"),
					&ast.BinaryExpr{
						Left:     test_helper.NewNumber("2"),
						Operator: test_helper.NewToken(token.ASTERISK, "*"),
						Right:    test_helper.NewNumber("2"),
					},
					&ast.ElementExpr{
						Left:  test_helper.NewNumber("1"),
						Right: test_helper.NewNumber("10"),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseFactor()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantExpr, expr)
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
				test_helper.NewToken(token.NUMBER, "1"),
			},
			wantExpr: test_helper.NewNumber("1"),
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("1"),
				Operator: test_helper.NewToken(token.ASTERISK, "*"),
				Right:    test_helper.NewNumber("2"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.DIV, "div"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("2"),
				Operator: test_helper.NewToken(token.DIV, "div"),
				Right:    test_helper.NewNumber("3"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.MOD, "mod"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("2"),
				Operator: test_helper.NewToken(token.MOD, "mod"),
				Right:    test_helper.NewNumber("3"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.SLASH, "/"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("1"),
				Operator: test_helper.NewToken(token.ASTERISK, "*"),
				Right: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("2"),
					Operator: test_helper.NewToken(token.SLASH, "/"),
					Right:    test_helper.NewNumber("3"),
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TRUE, "true"),
				test_helper.NewToken(token.AND, "and"),
				test_helper.NewToken(token.FALSE, "false"),
				test_helper.NewToken(token.AND, "and"),
				test_helper.NewToken(token.TRUE, "true"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewBool(true),
				Operator: test_helper.NewToken(token.AND, "and"),
				Right: &ast.BinaryExpr{
					Left:     test_helper.NewBool(false),
					Operator: test_helper.NewToken(token.AND, "and"),
					Right:    test_helper.NewBool(true),
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseTerm()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantExpr, expr)
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
				test_helper.NewToken(token.NUMBER, "1"),
			},
			wantExpr: &ast.NumericLiteral{
				Token: test_helper.NewToken(token.NUMBER, "1"),
				Value: ast.RawNumber("1"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.MINUS, "-"),
				test_helper.NewToken(token.NUMBER, "42"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: test_helper.NewToken(token.MINUS, "-"),
				Expr:     test_helper.NewNumber("42"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "42"),
			},
			wantExpr: &ast.UnaryExpr{
				Operator: test_helper.NewToken(token.PLUS, "+"),
				Expr:     test_helper.NewNumber("42"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("1"),
				Operator: test_helper.NewToken(token.PLUS, "+"),
				Right:    test_helper.NewNumber("2"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.MINUS, "-"),
				test_helper.NewToken(token.NUMBER, "2"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("1"),
				Operator: test_helper.NewToken(token.MINUS, "-"),
				Right:    test_helper.NewNumber("2"),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TRUE, "true"),
				test_helper.NewToken(token.OR, "or"),
				test_helper.NewToken(token.FALSE, "false"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewBool(true),
				Operator: test_helper.NewToken(token.OR, "or"),
				Right:    test_helper.NewBool(false),
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left:     test_helper.NewNumber("1"),
				Operator: test_helper.NewToken(token.PLUS, "+"),
				Right: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("2"),
					Operator: test_helper.NewToken(token.ASTERISK, "*"),
					Right:    test_helper.NewNumber("3"),
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.SLASH, "/"),
				test_helper.NewToken(token.NUMBER, "4"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("1"),
					Operator: test_helper.NewToken(token.SLASH, "/"),
					Right:    test_helper.NewNumber("4"),
				},
				Operator: test_helper.NewToken(token.PLUS, "+"),
				Right: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("2"),
					Operator: test_helper.NewToken(token.ASTERISK, "*"),
					Right:    test_helper.NewNumber("3"),
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseSimpleExpression()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantExpr, expr)
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
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.GTEQL, ">="),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "3"),
			},
			wantExpr: &ast.BinaryExpr{
				Left: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("2"),
					Operator: test_helper.NewToken(token.ASTERISK, "*"),
					Right:    test_helper.NewNumber("2"),
				},
				Operator: test_helper.NewToken(token.GTEQL, ">="),
				Right: &ast.BinaryExpr{
					Left:     test_helper.NewNumber("1"),
					Operator: test_helper.NewToken(token.PLUS, "+"),
					Right:    test_helper.NewNumber("3"),
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		expr, err := p.parseExpression()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantExpr, expr)
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
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "a"),
				test_helper.NewToken(token.NAMED, ":="),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
			},
			wantStmt: &ast.LabeledStmt{
				Label: test_helper.NewIdent("foo"),
				Stmt: &ast.AssignmentStmt{
					Identifier: test_helper.NewIdent("a"),
					Expr: &ast.BinaryExpr{
						Left:     test_helper.NewNumber("2"),
						Operator: test_helper.NewToken(token.ASTERISK, "*"),
						Right:    test_helper.NewNumber("2"),
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.IDENT, "foo"),
			},
			wantStmt: &ast.ProcedureCallStmt{
				Token:      test_helper.NewToken(token.IDENT, "foo"),
				Identifier: test_helper.NewIdent("foo"),
				Args:       []ast.Expression{},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.LPAREN, "("),
				test_helper.NewToken(token.NUMBER, "42"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.RPAREN, ")"),
			},
			wantStmt: &ast.ProcedureCallStmt{
				Token:      test_helper.NewToken(token.IDENT, "foo"),
				Identifier: test_helper.NewIdent("foo"),
				Args: []ast.Expression{
					test_helper.NewNumber("42"),
					&ast.BinaryExpr{
						Left:     test_helper.NewNumber("2"),
						Operator: test_helper.NewToken(token.ASTERISK, "*"),
						Right:    test_helper.NewNumber("2"),
					},
					test_helper.NewIdent("bar"),
				},
			},
		},
	}

	for _, tt := range tests {
		l := NewTestLexer(tt.input)
		p := New(l)
		stmt, err := p.parseSimpleStmt()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantStmt, stmt)
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
				test_helper.NewToken(token.BEGIN, "begin"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.NAMED, ":="),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.ASTERISK, "*"),
				test_helper.NewToken(token.NUMBER, "2"),
				test_helper.NewToken(token.SEMICOLON, ";"),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.LPAREN, "("),
				test_helper.NewToken(token.NUMBER, "42"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.MINUS, "-"),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.RPAREN, ")"),
				test_helper.NewToken(token.END, "end"),
			},
			wantStmt: &ast.CompoundStmt{
				Token: test_helper.NewToken(token.BEGIN, "begin"),
				Statements: []ast.Statement{
					&ast.AssignmentStmt{
						Identifier: test_helper.NewIdent("foo"),
						Expr: &ast.BinaryExpr{
							Left:     test_helper.NewNumber("2"),
							Operator: test_helper.NewToken(token.ASTERISK, "*"),
							Right:    test_helper.NewNumber("2"),
						},
					},
					&ast.ProcedureCallStmt{
						Token:      test_helper.NewToken(token.IDENT, "bar"),
						Identifier: test_helper.NewIdent("bar"),
						Args: []ast.Expression{
							test_helper.NewNumber("42"),
							&ast.UnaryExpr{
								Operator: test_helper.NewToken(token.MINUS, "-"),
								Expr:     test_helper.NewNumber("1"),
							},
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.IF, "if"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.IDENT, "bar"),
				test_helper.NewToken(token.THEN, "then"),
				test_helper.NewToken(token.IDENT, "a"),
				test_helper.NewToken(token.NAMED, ":="),
				test_helper.NewToken(token.NUMBER, "1"),
			},
			wantStmt: &ast.IfStmt{
				Token: test_helper.NewToken(token.IF, "if"),
				Cond: &ast.BinaryExpr{
					Left:     test_helper.NewIdent("foo"),
					Operator: test_helper.NewToken(token.EQUAL, "="),
					Right:    test_helper.NewIdent("bar"),
				},
				Then: &ast.AssignmentStmt{
					Identifier: test_helper.NewIdent("a"),
					Expr:       test_helper.NewNumber("1"),
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.WHILE, "while"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.LTEQL, "<="),
				test_helper.NewToken(token.NUMBER, "10"),
				test_helper.NewToken(token.DO, "do"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.NAMED, ":="),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "1"),
			},
			wantStmt: &ast.WhileStmt{
				Token: test_helper.NewToken(token.WHILE, "while"),
				Invariant: &ast.BinaryExpr{
					Left:     test_helper.NewIdent("foo"),
					Operator: test_helper.NewToken(token.LTEQL, "<="),
					Right:    test_helper.NewNumber("10"),
				},
				Body: &ast.AssignmentStmt{
					Identifier: test_helper.NewIdent("foo"),
					Expr: &ast.BinaryExpr{
						Left:     test_helper.NewIdent("foo"),
						Operator: test_helper.NewToken(token.PLUS, "+"),
						Right:    test_helper.NewNumber("1"),
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.REPEAT, "repeat"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.NAMED, ":="),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.PLUS, "+"),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.SEMICOLON, ";"),
				test_helper.NewToken(token.IDENT, "writeln"),
				test_helper.NewToken(token.LPAREN, "("),
				test_helper.NewToken(token.STRING, "hello world"),
				test_helper.NewToken(token.RPAREN, ")"),
				test_helper.NewToken(token.UNTIL, "until"),
				test_helper.NewToken(token.IDENT, "foo"),
				test_helper.NewToken(token.LTEQL, "<="),
				test_helper.NewToken(token.NUMBER, "10"),
			},
			wantStmt: &ast.RepeatStmt{
				Token: test_helper.NewToken(token.REPEAT, "repeat"),
				Invariant: &ast.BinaryExpr{
					Left:     test_helper.NewIdent("foo"),
					Operator: test_helper.NewToken(token.LTEQL, "<="),
					Right:    test_helper.NewNumber("10"),
				},
				Statements: []ast.Statement{
					&ast.AssignmentStmt{
						Identifier: test_helper.NewIdent("foo"),
						Expr: &ast.BinaryExpr{
							Left:     test_helper.NewIdent("foo"),
							Operator: test_helper.NewToken(token.PLUS, "+"),
							Right:    test_helper.NewNumber("1"),
						},
					},
					&ast.ProcedureCallStmt{
						Token:      test_helper.NewToken(token.IDENT, "writeln"),
						Identifier: test_helper.NewIdent("writeln"),
						Args: []ast.Expression{
							test_helper.NewString("hello world"),
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
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantStmt, stmt)
	}
}

func TestTypeDeclStmt(t *testing.T) {
	tests := []struct {
		input    []token.Token
		wantStmt ast.Statement
		wantErr  error
	}{
		{
			// type
			//   myint = integer;
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myint"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.IDENT, "integer"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token: test_helper.NewToken(token.IDENT, "myint"),
						Identifier: &ast.IdentifierExpr{
							Token: test_helper.NewToken(token.IDENT, "myint"),
							Value: "myint",
						},
						Definition: &ast.SimpleTypeDefinitionExpr{
							Token: test_helper.NewToken(token.IDENT, "integer"),
							Identifier: &ast.IdentifierExpr{
								Token: test_helper.NewToken(token.IDENT, "integer"),
								Value: "integer",
							},
						},
					},
				},
			},
		},
		{
			// type
			//   myrange1 = 0 .. 10;
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myrange"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.NUMBER, "0"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "10"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token: test_helper.NewToken(token.IDENT, "myrange"),
						Identifier: &ast.IdentifierExpr{
							Token: test_helper.NewToken(token.IDENT, "myrange"),
							Value: "myrange",
						},
						Definition: &ast.SubrangeTypeDefinitionExpr{
							Token: test_helper.NewToken(token.NUMBER, "0"),
							Left:  test_helper.NewNumber("0"),
							Right: test_helper.NewNumber("10"),
						},
					},
				},
			},
		},
		{
			// type
			//   myrange = 'A' .. 'Z';
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myrange"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.CHAR, "A"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.CHAR, "Z"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token: test_helper.NewToken(token.IDENT, "myrange"),
						Identifier: &ast.IdentifierExpr{
							Token: test_helper.NewToken(token.IDENT, "myrange"),
							Value: "myrange",
						},
						Definition: &ast.SubrangeTypeDefinitionExpr{
							Token: test_helper.NewToken(token.CHAR, "A"),
							Left:  test_helper.NewChar('A'),
							Right: test_helper.NewChar('Z'),
						},
					},
				},
			},
		},
		{
			// type
			//   myrange = 0 .. maxint;
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myrange"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.NUMBER, "0"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.IDENT, "maxint"),
				test_helper.NewToken(token.SEMICOLON, ";"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "myrange"),
						Identifier: test_helper.NewIdent("myrange"),
						Definition: &ast.SubrangeTypeDefinitionExpr{
							Token: test_helper.NewToken(token.NUMBER, "0"),
							Left:  test_helper.NewNumber("0"),
							Right: test_helper.NewIdent("maxint"),
						},
					},
				},
			},
		},
		{
			// type
			//   myrange = array[1 .. 10] of char
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myarr"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.ARRAY, "array"),
				test_helper.NewToken(token.LBRACKET, "["),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "10"),
				test_helper.NewToken(token.RBRACKET, "]"),
				test_helper.NewToken(token.OF, "of"),
				test_helper.NewToken(token.IDENT, "char"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "myarr"),
						Identifier: test_helper.NewIdent("myarr"),
						Definition: &ast.ArrayTypeDefinitionExpr{
							Token:  test_helper.NewToken(token.ARRAY, "array"),
							Packed: false,
							IndexTypeDef: &ast.SubrangeTypeDefinitionExpr{
								Token: test_helper.NewToken(token.NUMBER, "1"),
								Left:  test_helper.NewNumber("1"),
								Right: test_helper.NewNumber("10"),
							},
							ComponentTypeDef: &ast.SimpleTypeDefinitionExpr{
								Token:      test_helper.NewToken(token.IDENT, "char"),
								Identifier: test_helper.NewIdent("char"),
							},
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myrec"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.RECORD, "record"),
				test_helper.NewToken(token.IDENT, "account"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "integer"),
				test_helper.NewToken(token.SEMICOLON, ";"),
				test_helper.NewToken(token.IDENT, "name"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.PACKED, "packed"),
				test_helper.NewToken(token.ARRAY, "array"),
				test_helper.NewToken(token.LBRACKET, "["),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "20"),
				test_helper.NewToken(token.RBRACKET, "]"),
				test_helper.NewToken(token.OF, "of"),
				test_helper.NewToken(token.IDENT, "char"),
				test_helper.NewToken(token.END, "end"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "myrec"),
						Identifier: test_helper.NewIdent("myrec"),
						Definition: &ast.RecordTypeDefinitionExpr{
							Token: test_helper.NewToken(token.RECORD, "record"),
							Fields: []ast.TypeDefinitionStmt{
								{
									Token:      test_helper.NewToken(token.IDENT, "account"),
									Identifier: test_helper.NewIdent("account"),
									Definition: &ast.SimpleTypeDefinitionExpr{
										Token:      test_helper.NewToken(token.IDENT, "integer"),
										Identifier: test_helper.NewIdent("integer"),
									},
								},
								{
									Token:      test_helper.NewToken(token.IDENT, "name"),
									Identifier: test_helper.NewIdent("name"),
									Definition: &ast.ArrayTypeDefinitionExpr{
										Token:  test_helper.NewToken(token.PACKED, "packed"),
										Packed: true,
										IndexTypeDef: &ast.SubrangeTypeDefinitionExpr{
											Token: test_helper.NewToken(token.NUMBER, "1"),
											Left:  test_helper.NewNumber("1"),
											Right: test_helper.NewNumber("20"),
										},
										ComponentTypeDef: &ast.SimpleTypeDefinitionExpr{
											Token:      test_helper.NewToken(token.IDENT, "char"),
											Identifier: test_helper.NewIdent("char"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "transaction"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.LPAREN, "("),
				test_helper.NewToken(token.IDENT, "addition"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "deletion"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "enquiry"),
				test_helper.NewToken(token.COMMA, ","),
				test_helper.NewToken(token.IDENT, "update"),
				test_helper.NewToken(token.RPAREN, ")"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "transaction"),
						Identifier: test_helper.NewIdent("transaction"),
						Definition: &ast.EnumTypeDefinitionExpr{
							Token: test_helper.NewToken(token.LPAREN, "("),
							Values: []*ast.IdentifierExpr{
								test_helper.NewIdent("addition"),
								test_helper.NewIdent("deletion"),
								test_helper.NewIdent("enquiry"),
								test_helper.NewIdent("update"),
							},
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myset"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.SET, "set"),
				test_helper.NewToken(token.OF, "of"),
				test_helper.NewToken(token.CHAR, "A"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.CHAR, "Z"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "myset"),
						Identifier: test_helper.NewIdent("myset"),
						Definition: &ast.SetTypeDefinitionExpr{
							Token: test_helper.NewToken(token.SET, "set"),
							BaseTypeDef: &ast.SubrangeTypeDefinitionExpr{
								Token: test_helper.NewToken(token.CHAR, "A"),
								Left: &ast.CharLiteral{
									Token: test_helper.NewToken(token.CHAR, "A"),
									Value: 'A',
								},
								Right: &ast.CharLiteral{
									Token: test_helper.NewToken(token.CHAR, "Z"),
									Value: 'Z',
								},
							},
						},
					},
				},
			},
		},
		{
			input: []token.Token{
				test_helper.NewToken(token.TYPE, "type"),
				test_helper.NewToken(token.IDENT, "myfile"),
				test_helper.NewToken(token.EQUAL, "="),
				test_helper.NewToken(token.FILE, "file"),
				test_helper.NewToken(token.OF, "of"),
				test_helper.NewToken(token.RECORD, "record"),
				test_helper.NewToken(token.IDENT, "id"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.IDENT, "integer"),
				test_helper.NewToken(token.SEMICOLON, ";"),
				test_helper.NewToken(token.IDENT, "name"),
				test_helper.NewToken(token.COLON, ":"),
				test_helper.NewToken(token.PACKED, "packed"),
				test_helper.NewToken(token.ARRAY, "array"),
				test_helper.NewToken(token.LBRACKET, "["),
				test_helper.NewToken(token.NUMBER, "1"),
				test_helper.NewToken(token.DOTDOT, ".."),
				test_helper.NewToken(token.NUMBER, "20"),
				test_helper.NewToken(token.RBRACKET, "]"),
				test_helper.NewToken(token.OF, "of"),
				test_helper.NewToken(token.IDENT, "char"),
				test_helper.NewToken(token.END, "end"),
			},
			wantStmt: &ast.TypeDeclStmt{
				Token: test_helper.NewToken(token.TYPE, "type"),
				Definitions: []ast.TypeDefinitionStmt{
					{
						Token:      test_helper.NewToken(token.IDENT, "myfile"),
						Identifier: test_helper.NewIdent("myfile"),
						Definition: &ast.FileTypeDefinitionExpr{
							Token: test_helper.NewToken(token.FILE, "file"),
							BaseTypeDef: &ast.RecordTypeDefinitionExpr{
								Token: test_helper.NewToken(token.RECORD, "record"),
								Fields: []ast.TypeDefinitionStmt{
									{
										Token:      test_helper.NewToken(token.IDENT, "id"),
										Identifier: test_helper.NewIdent("id"),
										Definition: &ast.SimpleTypeDefinitionExpr{
											Token:      test_helper.NewToken(token.IDENT, "integer"),
											Identifier: test_helper.NewIdent("integer"),
										},
									},
									{
										Token:      test_helper.NewToken(token.IDENT, "name"),
										Identifier: test_helper.NewIdent("name"),
										Definition: &ast.ArrayTypeDefinitionExpr{
											Token:  test_helper.NewToken(token.PACKED, "packed"),
											Packed: true,
											IndexTypeDef: &ast.SubrangeTypeDefinitionExpr{
												Token: test_helper.NewToken(token.NUMBER, "1"),
												Left:  test_helper.NewNumber("1"),
												Right: test_helper.NewNumber("20"),
											},
											ComponentTypeDef: &ast.SimpleTypeDefinitionExpr{
												Token:      test_helper.NewToken(token.IDENT, "char"),
												Identifier: test_helper.NewIdent("char"),
											},
										},
									},
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
		stmt, err := p.parseTypeDeclStmt()
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantStmt, stmt)
	}
}
