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

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:          lexer,
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
		errors:         make([]string, 0),
	}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) hasNextToken() bool {
	if len(p.errors) > 0 {
		return false
	}
	return p.readToken.Type != token.EOF
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	statements := []ast.Statement{}
	for p.hasNextToken() {
		statement := p.parseStatement()
		if statement != nil {
			statements = append(statements, statement)
		}
	}
	program.Statements = statements
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.readToken.Type {
	case token.PROGRAM:
		return p.parseProgramDeclarationStatement()
	case token.LABEL:
		return p.parseLabelStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLabelStatement() *ast.LabelStatement {
	ls := &ast.LabelStatement{}
	ls.Token = p.readToken
	labels := []*ast.Identifier{}
	for {
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		label := &ast.Identifier{
			Token: p.readToken,
			Value: p.readToken.Literal,
		}
		labels = append(labels, label)
		if p.peekTokenIs(token.SEMICOLON) {
			break
		}
		if !p.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	p.nextToken()
	ls.Labels = labels

	return ls
}

func (p *Parser) parseExpressionStatement() ast.Statement {
	p.errors = append(p.errors, "not implemented: "+p.readToken.Literal)
	return nil
}

func (p *Parser) parseProgramDeclarationStatement() *ast.ProgramDeclarationStatement {
	decl := &ast.ProgramDeclarationStatement{}
	decl.Token = p.readToken
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	decl.Name = p.readToken.Literal
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	for p.readTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return decl
}

func (p *Parser) nextToken() {
	if p.readTokenIs(token.EOF) {
		return
	}
	read := p.peekToken
	peek := p.lexer.NextToken()
	p.readToken, p.peekToken = read, peek
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

func (p *Parser) readTokenIs(tt token.TokenType) bool {
	return p.readToken.Type == tt
}

func (p *Parser) peekError(tt token.TokenType) {
	msg := fmt.Sprintf("expected next token to be: %s, got: %s",
		tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Error() error {
	if len(p.errors) == 0 {
		return nil
	}
	msg := strings.Join(p.errors, "\n")
	return errors.New(msg)
}
