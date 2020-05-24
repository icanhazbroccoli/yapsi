package parser

import (
	"fmt"
	"unicode/utf8"
	"yapsi/pkg/ast"
	"yapsi/pkg/lexer"
	"yapsi/pkg/token"
)

// grammar rules are taken from:
// https://condor.depaul.edu/ichu/csc447/notes/wk2/pascal.html

type Parser struct {
	lexer            lexer.Interface
	prev, curr, next token.Token
}

func New(lexer lexer.Interface) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	p.advance()
	p.advance()

	return p
}

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	identifier, params, err := p.parseProgramHeader()
	if err != nil {
		return nil, err
	}

	block, err := p.parseBlock()
	if err != nil {
		return nil, err
	}

	program.Identifier = identifier
	program.Params = params
	program.Block = *block

	return program, nil
}

// <factor> ::= <variable> | <unsigned constant> | ( <expression> ) |
//				<function designator> | <set> | not <factor>
func (p *Parser) parseFactor() (ast.Expression, error) {
	if p.match(token.IDENT) {
		return &ast.IdentifierExpr{
			Token: p.previous(),
			Value: p.previous().Literal,
		}, nil
	}
	if p.match(token.NUMBER) {
		return &ast.NumericLiteral{
			Token: p.previous(),
			Value: ast.RawNumber(p.previous().Literal),
		}, nil
	}
	if p.match(token.STRING) {
		return &ast.StringLiteral{
			Token: p.previous(),
			Value: p.previous().Literal,
		}, nil
	}
	if p.match(token.CHAR) {
		r, _ := utf8.DecodeRune([]byte(p.previous().Literal))
		return &ast.CharLiteral{
			Token: p.previous(),
			Value: r,
		}, nil
	}
	if p.match(token.TRUE, token.FALSE) {
		var v bool
		switch p.previous().Literal {
		case "true", "TRUE", "True":
			v = true
		}
		return &ast.BoolLiteral{
			Token: p.previous(),
			Value: v,
		}, nil
	}
	if p.match(token.NOT) {
		op := p.previous()
		expr, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{
			Operator: op,
			Expr:     expr,
		}, nil
	}
	if p.match(token.LPAREN) {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.RPAREN); err != nil {
			return nil, err
		}
		return expr, nil
	}
	return nil, fmt.Errorf("Unknown literal token: %s (%s)",
		p.peek().Literal, p.peek().Type)
}

// <term> ::= <factor> | <term> <multiplying operator> <factor>
func (p *Parser) parseTerm() (ast.Expression, error) {
	expr, err := p.parseFactor()
	if err != nil {
		return nil, err
	}
	for p.match(
		token.ASTERISK,
		token.SLASH,
		token.DIV,
		token.MOD,
		token.AND,
	) {
		op := p.previous()
		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: op,
			Right:    right,
		}
	}
	return expr, nil
}

// <simple expression> ::= <term> | <sign> <term>| <simple expression> <adding operator> <term>
func (p *Parser) parseSimpleExpression() (ast.Expression, error) {
	if p.match(
		token.PLUS,
		token.MINUS,
	) {
		op := p.previous()
		expr, err := p.parseTerm()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{
			Operator: op,
			Expr:     expr,
		}, nil
	}
	expr, err := p.parseTerm()
	if err != nil {
		return nil, err
	}
	for p.match(
		token.PLUS,
		token.MINUS,
		token.OR,
	) {
		op := p.previous()
		right, err := p.parseSimpleExpression()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: op,
			Right:    right,
		}
	}
	return expr, nil
}

func (p *Parser) parseProgramHeader() (ast.ProgramIdentifier, []ast.Identifier, error) {
	if _, err := p.consume(token.PROGRAM); err != nil {
		return "", nil, err
	}
	identifier, err := p.consume(token.IDENT)
	if err != nil {
		return "", nil, err
	}
	params := []ast.Identifier{}
	if p.match(token.LPAREN) {
		for p.check(token.IDENT) {
			param, err := p.consume(token.IDENT)
			if err != nil {
				return "", nil, err
			}
			params = append(params, ast.Identifier(param.Literal))
			if !p.match(token.COMMA) {
				break
			}
		}
		if _, err := p.consume(token.RPAREN); err != nil {
			return "", nil, err
		}
	}
	if _, err := p.consume(token.SEMICOLON); err != nil {
		return "", nil, err
	}
	return ast.ProgramIdentifier(identifier.Literal), params, nil
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
		for p.hasNext() {
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
	labels := []ast.Label{}
	if !p.match(token.LABEL) {
		return labels, nil
	}
	for p.hasNext() {
		ident, err := p.consume(token.IDENT, token.NUMBER)
		if err != nil {
			return nil, err
		}
		labels = append(labels, ast.Label{
			Identifier: ast.LabelIdentifier(ident.Literal),
		})
		if !p.match(token.COMMA) {
			break
		}
	}
	if _, err := p.consume(token.SEMICOLON); err != nil {
		return nil, err
	}
	return labels, nil
}

func (p *Parser) parseConstants() ([]ast.Constant, error) {
	constants := []ast.Constant{}
	if !p.match(token.CONST) {
		return constants, nil
	}
	for p.hasNext() {
		ident, err := p.consume(token.IDENT)
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.EQUAL); err != nil {
			return nil, err
		}
		value, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		constants = append(constants, ast.Constant{
			Identifier: ast.ConstIdentifier(ident.Literal),
			Raw:        value,
		})
		if _, err := p.consume(token.SEMICOLON); err != nil {
			return nil, err
		}
		if !p.match(token.IDENT) {
			break
		}
	}

	return constants, nil
}

// <expression> ::= <simple expression> |
//					<simple expression> <relational operator> <simple expression>
func (p *Parser) parseExpression() (ast.Expression, error) {
	expr, err := p.parseSimpleExpression()
	if err != nil {
		return nil, err
	}
	if p.match(
		token.EQUAL,
		token.NOTEQUAL,
		token.LESS,
		token.LTEQL,
		token.GTEQL,
		token.GREATER,
		token.IN,
	) {
		op := p.previous()
		right, err := p.parseSimpleExpression()
		if err != nil {
			return nil, err
		}
		expr = &ast.BinaryExpr{
			Left:     expr,
			Operator: op,
			Right:    right,
		}
	}
	return expr, nil
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

func (p *Parser) previous() token.Token {
	return p.prev
}

func (p *Parser) hasNext() bool {
	return p.peek().Type != token.EOF
}

func (p *Parser) consume(tt ...token.TokenType) (token.Token, error) {
	for _, t := range tt {
		if p.check(t) {
			return p.advance(), nil
		}
	}
	line, pos := p.lexer.Pos()
	err := fmt.Errorf("Unexpected token at line: %d, pos: %d; got token: %s, expected: %s",
		line, pos, p.curr.Type, tt)
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
