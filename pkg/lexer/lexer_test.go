package lexer

import (
	"bufio"
	"strings"
	"testing"
	"yapsi/pkg/token"
)

func TestLexerNextToken(t *testing.T) {
	input := `
PROGRAM HelloWorld;

LABEL
	foobar;

VAR
	X := 1;
	time := 2.1;
	readinteger := 3.5e18;
	WG4 := 0e1;
	AltrerHeatSetting := 5E42;
	FooBar := 'Hello world!';
{ this is a comment }
(* this is a comment too *)
(*
	multiline
	comment
	*)
{ comment with distinct separators *)
(* comment with another combination of distinct separators }

BEGIN
	WHILE Y >= -0.5 DO
	BEGIN
		writeln();
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
		{token.LABEL, "LABEL"}, {token.IDENT, "foobar"}, {token.SEMICOLON, ";"},
		{token.VAR, "VAR"},
		{token.IDENT, "X"}, {token.NAMED, ":="}, {token.NUMBER, "1"}, {token.SEMICOLON, ";"},
		{token.IDENT, "time"}, {token.NAMED, ":="}, {token.NUMBER, "2.1"}, {token.SEMICOLON, ";"},
		{token.IDENT, "readinteger"}, {token.NAMED, ":="}, {token.NUMBER, "3.5e18"}, {token.SEMICOLON, ";"},
		{token.IDENT, "WG4"}, {token.NAMED, ":="}, {token.NUMBER, "0e1"}, {token.SEMICOLON, ";"},
		{token.IDENT, "AltrerHeatSetting"}, {token.NAMED, ":="}, {token.NUMBER, "5E42"}, {token.SEMICOLON, ";"},

		{token.IDENT, "FooBar"}, {token.NAMED, ":="}, {token.STRING, "Hello world!"}, {token.SEMICOLON, ";"},
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

	lex := New(bufio.NewReader(reader))

	for _, tt := range tests {
		tok := lex.NextToken()
		line, col := lex.pos()
		if tok.Type != tt.expType {
			t.Fatalf("Unexpected token type around pos: line: %d, col: %d: got=%#v, want: %#v",
				line, col, tok.Type, tt.expType)
		}

		if tok.Literal != tt.expLiteral {
			t.Fatalf("Unexpected token literal around pos: line: %d, col: %d: got=%s, want: %s",
				line, col, tok.Literal, tt.expLiteral)
		}
	}
}
