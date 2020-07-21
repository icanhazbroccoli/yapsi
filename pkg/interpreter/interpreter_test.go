package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"yapsi/pkg/ast"
	test_helper "yapsi/pkg/internal/test"
	"yapsi/pkg/object"
	"yapsi/pkg/token"
)

func TestVisitNumericLiteral(t *testing.T) {
	tests := []struct {
		input   *ast.NumericLiteral
		wantRes ast.VisitorResult
		wantErr error
	}{
		{
			input: &ast.NumericLiteral{
				Token: token.Token{
					Type:    token.NUMBER,
					Literal: "0",
				},
				Value: ast.RawNumber("0"),
			},
			wantRes: &object.Integer{
				Value: 0,
			},
		},
		{
			input: &ast.NumericLiteral{
				Token: token.Token{
					Type:    token.NUMBER,
					Literal: "0.1",
				},
				Value: ast.RawNumber("0.1"),
			},
			wantRes: &object.Real{
				Value: 0.1,
			},
		},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		i := &Interpreter{env}
		res, err := i.VisitNumericLiteral(tt.input)
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantRes, res)
	}
}

func TestVisitStringLiteral(t *testing.T) {
	tests := []struct {
		input   *ast.StringLiteral
		wantRes ast.VisitorResult
		wantErr error
	}{
		{
			input: &ast.StringLiteral{
				Token: token.Token{
					Type:    token.STRING,
					Literal: "hello world",
				},
				Value: "hello world",
			},
			wantRes: &object.String{
				Value: "hello world",
			},
		},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		i := &Interpreter{env}
		res, err := i.VisitStringLiteral(tt.input)
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantRes, res)
	}
}

func TestVisitBoolLiteral(t *testing.T) {
	tests := []struct {
		input   *ast.BoolLiteral
		wantRes ast.VisitorResult
		wantErr error
	}{
		{
			input: &ast.BoolLiteral{
				Token: token.Token{
					Type:    token.BOOL,
					Literal: "true",
				},
				Value: true,
			},
			wantRes: &object.Boolean{
				Value: true,
			},
		},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		i := &Interpreter{env}
		res, err := i.VisitBoolLiteral(tt.input)
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantRes, res)
	}
}

func TestVisitSimpleTypeDefinitionExpr(t *testing.T) {
	tests := []struct {
		input   *ast.SimpleTypeDefinitionExpr
		wantRes ast.VisitorResult
		wantErr error
	}{
		{
			input: &ast.SimpleTypeDefinitionExpr{
				Token: test_helper.NewToken(token.IDENT, "integer"),
				Identifier: &ast.IdentifierExpr{
					Token: test_helper.NewToken(token.IDENT, "integer"),
					Value: "integer",
				},
			},
			wantRes: nil,
		},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		i := &Interpreter{env}
		res, err := i.VisitSimpleTypeDefinitionExpr(tt.input)
		assert.Equal(t, tt.wantErr, err)
		if err != nil {
			continue
		}
		assert.Equal(t, tt.wantRes, res)
		//TODO: check env
	}
}
