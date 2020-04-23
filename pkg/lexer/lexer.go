package lexer

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"yapsi/pkg/token"
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
		switch lex.peek {
		case '.':
			read := lex.read
			if err := lex.next(); err != nil {
				return tok, err
			}
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
		switch lex.peek {
		case '*':
			if err := lex.next(); err != nil {
				return tok, err
			}
			if err := lex.skipComment(); err != nil {
				return tok, err
			}
			goto Begin
		default:
			tok = newToken(token.LPAREN, lex.read)
		}
	case ')':
		tok = newToken(token.RPAREN, lex.read)
	case '\'':
		lit, err := lex.readString()
		if err != nil {
			return tok, err
		}
		tok.Type = token.STRING
		tok.Literal = lit
	case '{':
		if err := lex.skipComment(); err != nil {
			return tok, err
		}
		goto Begin
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		return tok, nil
	default:
		if isDigit(lex.read) {
			lit, err := lex.readNumber()
			if err != nil {
				return tok, err
			}
			tok.Type = token.NUMBER
			tok.Literal = lit
		} else if isLetter(lex.read) {
			lit, err := lex.readIdentifier()
			if err != nil {
				return tok, err
			}
			tok.Type = token.LookupIdent(lit)
			tok.Literal = lit
		} else {
			tok.Type = token.ILLEGAL
			return tok, lex.error(fmt.Sprintf(
				"Illegal token: %c", lex.read))
		}
	}
	if err := lex.next(); err != nil {
		return tok, err
	}
	return tok, nil
}

func (lex *Lexer) readNumber() (string, error) {
	var buf bytes.Buffer
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !(isDigit(lex.peek) ||
			lex.peek == '.' || lex.peek == 'e' || lex.peek == 'E') {
			break
		}
		if err := lex.next(); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (lex *Lexer) readIdentifier() (string, error) {
	var buf bytes.Buffer
	for {
		buf.WriteRune(lex.read)
		if !lex.hasNext() || !isAlphanumeric(lex.peek) {
			break
		}
		if err := lex.next(); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (lex *Lexer) readString() (string, error) {
	var buf bytes.Buffer
	for lex.hasNext() && lex.peek != '\'' {
		if err := lex.next(); err != nil {
			return "", err
		}
		buf.WriteRune(lex.read)
	}
	if err := lex.next(); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (lex *Lexer) skipComment() error {
	shouldBreak := false
	if err := lex.next(); err != nil {
		return err
	}
	for !shouldBreak {
		if !lex.hasNext() {
			return lex.error("No comment closing token found")
		}
		switch lex.read {
		case '}':
			shouldBreak = true
		case '*':
			if lex.peek == ')' {
				if err := lex.next(); err != nil {
					return err
				}
				shouldBreak = true
			}
		}
		if err := lex.next(); err != nil {
			return err
		}
	}
	return nil
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

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isAlphanumeric(r rune) bool {
	return isLetter(r) || isDigit(r)
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
