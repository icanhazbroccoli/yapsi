package parser

import (
	"errors"
	"fmt"
	"strings"
	"yapsi/pkg/ast"
	"yapsi/pkg/lexer"
	"yapsi/pkg/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer *lexer.Lexer

	readToken, peekToken token.Token
	errors               []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(lexer *lexer.Lexer) (*Parser, error) {
	p := &Parser{
		lexer:          lexer,
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
		errors:         make([]string, 0),
	}

	if err := p.nextToken(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program, err := p.parseProgramDeclaration()
	if err != nil {
		return nil, err
	}
	statements := []ast.Statement{}
	//for p.readToken.Type != token.EOF &&
	// p.readToken.Type != token.PERIOD {
	//	statement := p.parseStatement()
	//	if statement != nil {
	//		statements = append(statements, statement)
	//	}
	//	if err := p.nextToken(); err != nil {
	//		return nil, err
	//	}
	//}
	program.Statements = statements
	return program, nil
}

func (p *Parser) parseProgramDeclaration() (*ast.Program, error) {
	program := &ast.Program{}
	if !p.expectPeek(token.PROGRAM) {
		return nil, p.Errors()
	}
	if !p.expectPeek(token.IDENT) {
		return nil, p.Errors()
	}
	program.Name = p.readToken.Literal
	if !p.expectPeek(token.SEMICOLON) {
		return nil, p.Errors()
	}
	for p.peekTokenIs(token.SEMICOLON) {
		if err := p.nextToken(); err != nil {
			return nil, err
		}
	}
	return program, nil
}

func (p *Parser) nextToken() error {
	read := p.peekToken
	peek, err := p.lexer.NextToken()
	if err != nil {
		return nil
	}
	p.readToken, p.peekToken = read, peek
	return nil
}

func (p *Parser) expectPeek(tt token.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken()
		return true
	} else {
		p.peekError(tt)
		return false
	}
}

func (p *Parser) peekTokenIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) peekError(tt token.TokenType) {
	msg := fmt.Sprintf("expected next token to be: %s, got: %s",
		tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() error {
	if len(p.errors) == 0 {
		return nil
	}
	msg := strings.Join(p.errors, "\n")
	return errors.New(msg)
}
