package lexer

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"yapsi/pkg/token"
)

type StringMode uint8

const (
	SINGLELINE StringMode = iota
	MULTILINE
)

type Lexer struct {
	reader     *strings.Reader
	line, col  int
	read, peek rune
}

func New(reader *strings.Reader) (*Lexer, error) {
	lex := &Lexer{
		reader: reader,
	}
	lex.next()
	lex.next()
	lex.line, lex.col = 1, 1
	return lex, nil
}

func (lex *Lexer) NextToken() (token.Token, error) {
	var tok token.Token
	var err error

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
			if err := lex.next(); err != nil {
				return tok, err
			}
			tok.Type = token.LTEQL
			tok.Literal = string(read) + string(lex.read)
		default:
			tok = newToken(token.LESS, lex.read)
		}
	case '>':
		switch lex.peek {
		case '=':
			read := lex.read
			if err := lex.next(); err != nil {
				return tok, err
			}
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
		tok = newToken(token.PERIOD, lex.read)
	case ',':
		tok = newToken(token.COMMA, lex.read)
	case ':':
		switch lex.peek {
		case '=':
			read := lex.read
			if err := lex.next(); err != nil {
				return tok, err
			}
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
		tok = newToken(token.LPAREN, lex.read)
	case ')':
		tok = newToken(token.RPAREN, lex.read)
	case '\'':
		tok, err = lex.readString(SINGLELINE)
		if err != nil {
			return tok, err
		}
	case '"':
		tok, err = lex.readString(MULTILINE)
		if err != nil {
			return tok, err
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		return tok, nil
	default:
		tok, err = lex.readIdentOrNum()
		if err != nil {
			return tok, err
		}
	}
	if err := lex.next(); err != nil {
		return tok, err
	}
	return tok, nil
}

func (lex *Lexer) readIdentOrNum() (token.Token, error) {
	if isDigit(lex.read) {
		return lex.readNumber()
	} else {
		return lex.readIdent()
	}
}

func (lex *Lexer) readString(sm StringMode) (token.Token, error) {
	var buf bytes.Buffer
	var tok token.Token
	var sentinel rune
	switch sm {
	case SINGLELINE:
		sentinel = '\''
	case MULTILINE:
		sentinel = '"'
	default:
		return tok, lex.error(fmt.Sprintf(
			"unknown string mode: %v", sm))
	}
	if err := lex.next(); err != nil {
		return tok, err
	}
	for {
		if lex.read == '\n' && sm != MULTILINE {
			return tok, lex.error(
				"unexpected newline in a non-multiline string literal")
		}
		buf.WriteRune(lex.read)
		if !lex.hasNext() || lex.peek == sentinel {
			break
		}
		if err := lex.next(); err != nil {
			return tok, err
		}
	}

	if lex.peek != sentinel {
		return tok, lex.error("missing closing string quote")
	}
	if err := lex.next(); err != nil {
		return tok, err
	}
	tok.Type = token.STRING
	tok.Literal = buf.String()
	return tok, nil
}

func (lex *Lexer) readNumber() (token.Token, error) {
	var buf bytes.Buffer
	var tok token.Token
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !(isDigit(lex.peek) ||
			lex.peek == '.' || lex.peek == 'b' || lex.peek == 'x') {
			break
		}
		if err := lex.next(); err != nil {
			return tok, err
		}
	}

	tok.Type = token.NUMBER
	tok.Literal = buf.String()
	return tok, nil
}

func (lex *Lexer) readIdent() (token.Token, error) {
	var tok token.Token
	var buf bytes.Buffer
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !isAlphaNum(lex.peek) {
			break
		}
		if err := lex.next(); err != nil {
			return tok, err
		}
	}

	tok.Literal = buf.String()
	tok.Type = token.LookupIdent(tok.Literal)

	return tok, nil
}

func (lex *Lexer) hasNext() bool {
	return lex.read != 0
}

func (lex *Lexer) next() error {
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
			return err
		}
		peek = 0
	}
	lex.read, lex.peek = read, peek
	return nil
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

func isAlphaNum(r rune) bool {
	return isDigit(r) || (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') || (r == '_')
}

func (lex *Lexer) error(msg string) error {
	return fmt.Errorf("Syntax error on line: %d, pos: %d: %s",
		lex.line, lex.col, msg)
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
