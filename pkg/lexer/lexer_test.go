package lexer

import (
	"strings"
	"testing"
	"yapsi/pkg/token"
)

func TestLexerNextToken(t *testing.T) {
	input := `
PROGRAM HelloWorld;

USES crt;

CONST
	LOG10 = 2.3025851;

VAR
	i: INTEGER;
	X, Y, A: REAL;

PROCEDURE foo(x: INTEGER;
			  y: REAL);
BEGIN
FOR i := 1 TO 42 DO
	BEGIN
		X := 0.2 * i - 0.1;
		A := POWER(X);
	END

BEGIN
	WHILE Y >= -0.5 DO
	BEGIN
		writeln('Hello World!');
	END
	readkey;
END.
	`
	tests := []struct {
		expType    token.TokenType
		expLiteral string
	}{
		{token.PROGRAM, "PROGRAM"},
		{token.IDENT, "HelloWorld"},
		{token.SEMICOLON, ";"},
		{token.USES, "USES"},
		{token.IDENT, "crt"},
		{token.SEMICOLON, ";"},
		{token.CONST, "CONST"},
		{token.IDENT, "LOG10"},
		{token.EQUAL, "="},
		{token.NUMBER, "2.3025851"},
		{token.SEMICOLON, ";"},
		{token.VAR, "VAR"},
		{token.IDENT, "i"},
		{token.COLON, ":"},
		{token.INTEGER, "INTEGER"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "X"},
		{token.COMMA, ","},
		{token.IDENT, "Y"},
		{token.COMMA, ","},
		{token.IDENT, "A"},
		{token.COLON, ":"},
		{token.REAL, "REAL"},
		{token.SEMICOLON, ";"},
		{token.PROCEDURE, "PROCEDURE"},
		{token.IDENT, "foo"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COLON, ":"},
		{token.INTEGER, "INTEGER"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "y"},
		{token.COLON, ":"},
		{token.REAL, "REAL"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BEGIN, "BEGIN"},
		{token.FOR, "FOR"},
		{token.IDENT, "i"},
		{token.NAMED, ":="},
		{token.NUMBER, "1"},
		{token.TO, "TO"},
		{token.NUMBER, "42"},
		{token.DO, "DO"},
		{token.BEGIN, "BEGIN"},
		{token.IDENT, "X"},
		{token.NAMED, ":="},
		{token.NUMBER, "0.2"},
		{token.ASTERISK, "*"},
		{token.IDENT, "i"},
		{token.MINUS, "-"},
		{token.NUMBER, "0.1"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "A"},
		{token.NAMED, ":="},
		{token.IDENT, "POWER"},
		{token.LPAREN, "("},
		{token.IDENT, "X"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.END, "END"},
		{token.BEGIN, "BEGIN"},
		{token.WHILE, "WHILE"},
		{token.IDENT, "Y"},
		{token.GTEQL, ">="},
		{token.MINUS, "-"},
		{token.NUMBER, "0.5"},
		{token.DO, "DO"},
		{token.BEGIN, "BEGIN"},
		{token.IDENT, "writeln"},
		{token.LPAREN, "("},
		{token.STRING, "Hello World!"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.END, "END"},
		{token.IDENT, "readkey"},
		{token.SEMICOLON, ";"},
		{token.END, "END"},
		{token.PERIOD, "."},
		{token.EOF, ""},
	}

	reader := strings.NewReader(input)

	lex, err := New(reader)
	if err != nil {
		t.Fatalf("Unexpected New() error: %s", err)
	}

	for _, tt := range tests {
		tok, err := lex.NextToken()
		line, col := lex.pos()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if tok.Type != tt.expType {
			t.Errorf("Unexpected token type around pos: line: %d, col: %d: got=%#v, want: %#v",
				line, col, tok.Type, tt.expType)
		}

		if tok.Literal != tt.expLiteral {
			t.Errorf("Unexpected token literal around pos: line: %d, col: %d: got=%s, want: %s",
				line, col, tok.Literal, tt.expLiteral)
		}
	}
}
