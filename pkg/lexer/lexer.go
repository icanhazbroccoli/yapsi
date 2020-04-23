package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"yapsi/pkg/token"
)

type Lexer struct {
	reader     *bufio.Reader
	line, col  int
	read, peek rune
}

func New(reader *bufio.Reader) *Lexer {
	lex := &Lexer{
		reader: reader,
	}
	lex.next()
	lex.next()
	lex.line, lex.col = 1, 1
	return lex
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

Begin:
	lex.eatWhitespace()

	switch lex.read {
	case '+':
		tok = newToken(token.PLUS, lex.read)
	case '-':
		tok = newToken(token.MINUS, lex.read)
	case '*':
		tok = newToken(token.ASTERISK, lex.read)
	case '/':
		tok = newToken(token.SLASH, lex.read)
	case '=':
		tok = newToken(token.EQUAL, lex.read)
	case '<':
		switch lex.peek {
		case '=':
			read := lex.read
			lex.next()
			tok.Type = token.LTEQL
			tok.Literal = string(read) + string(lex.read)
		default:
			tok = newToken(token.LESS, lex.read)
		}
	case '>':
		switch lex.peek {
		case '=':
			read := lex.read
			lex.next()
			tok.Type = token.GTEQL
			tok.Literal = string(read) + string(lex.read)
		default:
			tok = newToken(token.GREATER, lex.read)
		}
	case '[':
		tok = newToken(token.LBRACKET, lex.read)
	case ']':
		tok = newToken(token.RBRACKET, lex.read)
	case '.':
		switch lex.peek {
		case '.':
			read := lex.read
			lex.next()
			tok.Type = token.DOTDOT
			tok.Literal = string(read) + string(lex.read)
		default:
			tok = newToken(token.PERIOD, lex.read)
		}
	case ',':
		tok = newToken(token.COMMA, lex.read)
	case ':':
		switch lex.peek {
		case '=':
			read := lex.read
			lex.next()
			tok.Type = token.NAMED
			tok.Literal = string(read) + string(lex.read)
		default:
			tok = newToken(token.COLON, lex.read)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lex.read)
	case '^':
		tok = newToken(token.PTR, lex.read)
	case '(':
		switch lex.peek {
		case '*':
			lex.next()
			lex.skipComment()
			goto Begin
		default:
			tok = newToken(token.LPAREN, lex.read)
		}
	case ')':
		tok = newToken(token.RPAREN, lex.read)
	case '\'':
		tok.Type = token.STRING
		tok.Literal = lex.readString()
	case '{':
		lex.skipComment()
		goto Begin
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		return tok
	default:
		if isDigit(lex.read) {
			tok.Type = token.NUMBER
			tok.Literal = lex.readNumber()
		} else if isLetter(lex.read) {
			lit := lex.readIdentifier()
			tok.Type = token.LookupIdent(lit)
			tok.Literal = lit
		} else {
			tok.Type = token.ILLEGAL
			lex.error(fmt.Sprintf(
				"Illegal token: %c", lex.read))
		}
	}
	lex.next()

	return tok
}

func (lex *Lexer) readNumber() string {
	var buf bytes.Buffer
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !(isDigit(lex.peek) ||
			lex.peek == '.' || lex.peek == 'e' || lex.peek == 'E') {
			break
		}
		lex.next()
	}
	return buf.String()
}

func (lex *Lexer) readIdentifier() string {
	var buf bytes.Buffer
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !isAlphanumeric(lex.peek) {
			break
		}
		lex.next()
	}
	return buf.String()
}

func (lex *Lexer) readString() string {
	var buf bytes.Buffer
	for lex.hasNext() && lex.peek != '\'' {
		lex.next()
		buf.WriteRune(lex.read)
	}
	lex.next()
	return buf.String()
}

func (lex *Lexer) skipComment() {
	shouldBreak := false
	lex.next()
	for !shouldBreak {
		if !lex.hasNext() {
			break
		}
		switch lex.read {
		case '}':
			shouldBreak = true
		case '*':
			if lex.peek == ')' {
				lex.next()
				shouldBreak = true
			}
		}
		lex.next()
	}
}

func (lex *Lexer) hasNext() bool {
	return lex.read != 0
}

func (lex *Lexer) next() {
	read, peek := lex.read, lex.peek
	read = peek
	lex.col++
	if read == '\n' {
		lex.col = 0
		lex.line++
	}
	var err error
	peek, _, err = lex.reader.ReadRune()
	if err != nil {
		if err != io.EOF {
			panic(err.Error())
		}
		peek = 0
	}
	lex.read, lex.peek = read, peek
}

func (lex *Lexer) eatWhitespace() {
	for lex.hasNext() && isWhitespace(lex.read) {
		lex.next()
	}
}

func isWhitespace(r rune) bool {
	switch r {
	case ' ', '\t', '\n', '\r':
		return true
	}
	return false
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isAlphanumeric(r rune) bool {
	return isLetter(r) || isDigit(r)
}

func (lex *Lexer) error(msg string) error {
	panic(fmt.Sprintf("Syntax error on line: %d, pos: %d: %s",
		lex.line, lex.col, msg))
}

func (lex *Lexer) pos() (int, int) {
	return lex.line, lex.col
}

func newToken(t token.TokenType, l rune) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(l),
	}
}
