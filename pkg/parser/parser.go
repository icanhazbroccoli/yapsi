package parser

import (
	"fmt"
	"yapsi/pkg/ast"
	"yapsi/pkg/lexer"
	"yapsi/pkg/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer lexer.Interface

	prev, curr, next token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(lexer lexer.Interface) *Parser {
	p := &Parser{
		lexer:          lexer,
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	p.advance()
	p.advance()

	return p
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	if _, err := p.consume(token.PROGRAM); err != nil {
		return nil, err
	}
	identifier, err := p.consume(token.IDENT)
	if err != nil {
		return nil, err
	}
	if _, err := p.consume(token.SEMICOLON); err != nil {
		return nil, err
	}
	block, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	program.Identifier = ast.ProgramIdentifier(identifier.Literal)
	program.Block = *block

	return program, nil
}

func (p *Parser) parseBlock() (*ast.Block, error) {
	block := &ast.Block{}
	labels, err := p.parseLabels()
	if err != nil {
		return nil, err
	}
	constants, err := p.parseConstants()
	if err != nil {
		return nil, err
	}
	types, err := p.parseTypes()
	if err != nil {
		return nil, err
	}
	variables, err := p.parseVariables()
	if err != nil {
		return nil, err
	}
	callables, err := p.parseCallables()
	if err != nil {
		return nil, err
	}
	body, err := p.parseBody()
	if err != nil {
		return nil, err
	}

	block.Labels = labels
	block.Constants = constants
	block.Types = types
	block.Variables = variables
	block.Callables = callables
	block.Body = body

	return block, nil
}

func (p *Parser) parseVariables() ([]ast.Variable, error) {
	variables := []ast.Variable{}
	if !p.match(token.VAR) {
		return variables, nil
	}

	for {
		identifiers := []ast.VarIdentifier{}
		for {
			identifier, err := p.consume(token.IDENT)
			if err != nil {
				return nil, err
			}
			identifiers = append(identifiers, ast.VarIdentifier(identifier.Literal))
			if !p.check(token.COMMA) {
				break
			}
			if _, err := p.consume(token.COMMA); err != nil {
				return nil, err
			}
		}
		if _, err := p.consume(token.COLON); err != nil {
			return nil, err
		}
		typeIdentifier, err := p.consume(token.IDENT)
		if err != nil {
			return nil, err
		}
		for _, varident := range identifiers {
			variables = append(variables, ast.Variable{
				Identifier: varident,
				Type:       ast.TypeIdentifier(typeIdentifier.Literal),
			})
		}
		if _, err := p.consume(token.SEMICOLON); err != nil {
			return nil, err
		}
		if !p.check(token.IDENT) {
			break
		}
	}

	return variables, nil
}

func (p *Parser) parseLabels() ([]ast.Label, error) {
	//TODO
	panic("not implemented")
}

func (p *Parser) parseConstants() ([]ast.Constant, error) {
	//TODO
	panic("not implemented")
}

func (p *Parser) parseTypes() ([]ast.Type, error) {
	//TODO
	panic("not implemented")
}

func (p *Parser) parseCallables() ([]ast.Callable, error) {
	//TODO
	panic("not implemented")
}

func (p *Parser) parseBody() ([]ast.Statement, error) {
	//TODO
	panic("not implemented")
}

func (p *Parser) match(tt ...token.TokenType) bool {
	for _, t := range tt {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) peek() token.Token {
	return p.curr
}

func (p *Parser) hasNext() bool {
	return p.peek().Type != token.EOF
}

func (p *Parser) previous() token.Token {
	return p.prev
}

func (p *Parser) consume(t token.TokenType) (token.Token, error) {
	if p.check(t) {
		return p.advance(), nil
	}
	line, pos := p.lexer.Pos()
	err := fmt.Errorf("Unexpected token at line: %d, pos: %d; got token: %s, expected: %s",
		line, pos, p.curr.Type, t)
	//TODO
	//p.recover()
	return token.Token{}, err
}

func (p *Parser) check(t token.TokenType) bool {
	if !p.hasNext() {
		return false
	}
	return p.peek().Type == t
}

func (p *Parser) advance() token.Token {
	if p.hasNext() {
		p.prev, p.curr, p.next = p.curr, p.next, p.lexer.NextToken()
	}
	return p.prev
}
