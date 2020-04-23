package parser

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"yapsi/pkg/ast"
	"yapsi/pkg/lexer"
	"yapsi/pkg/token"
)

func TestParsingProgramDeclarationStatement(t *testing.T) {
	input := `PROGRAM foo;`
	reader := newReader(input)
	lex := lexer.New(reader)
	p := New(lex)
	stmt := p.parseProgramDeclarationStatement()
	checkParserErrors(t, p)
	expectStatement(t, stmt, &ast.ProgramDeclarationStatement{
		Token: token.Token{
			Type:    token.PROGRAM,
			Literal: "PROGRAM",
		},
		Name: "foo",
	})
}

func TestParsingLabelStatement(t *testing.T) {
	input := `
LABEL
	foo1,
	bar2,
	baz3;
	`
	reader := newReader(input)
	lex := lexer.New(reader)
	p := New(lex)
	stmt := p.parseLabelStatement()
	expectStatement(t, stmt, &ast.LabelStatement{
		Token: token.Token{
			Type:    token.LABEL,
			Literal: "LABEL",
		},
		Labels: []*ast.Identifier{
			newIdent("foo1"),
			newIdent("bar2"),
			newIdent("baz3"),
		},
	})
}

func newIdent(name string) *ast.Identifier {
	return &ast.Identifier{
		Token: token.Token{
			Type:    token.IDENT,
			Literal: name,
		},
		Value: name,
	}
}

func newReader(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}

func expectStatement(t *testing.T, got, want ast.Statement) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected statement:\ngot=%s\nwant=%s",
			got.String(), want.String())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	if err := p.Error(); err != nil {
		t.Fatalf(err.Error())
	}
}
