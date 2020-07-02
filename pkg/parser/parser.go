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

func (p *Parser) ParseProgram() (*ast.ProgramStmt, error) {
	if _, err := p.consume(token.PROGRAM); err != nil {
		return nil, err
	}
	tok := p.previous()
	ident, err := p.consume(token.IDENT)
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
	if _, err := p.consume(token.PERIOD); err != nil {
		return nil, err
	}
	return &ast.ProgramStmt{
		Token: tok,
		Identifier: &ast.IdentifierExpr{
			Token: ident,
			Value: ident.Literal,
		},
		Block: block,
	}, nil
}

func (p *Parser) parseBlock() (*ast.BlockStmt, error) {
	tok := p.previous() // FIXME
	declStmt, err := p.parseVarDeclStmt()
	if err != nil {
		return nil, err
	}
	procedures, functions, err := p.parseProcedureAndFunctionDeclStmts()
	if err != nil {
		return nil, err
	}
	stmt, err := p.parseCompoundStmt()
	if err != nil {
		return nil, err
	}
	return &ast.BlockStmt{
		Token:      tok,
		VarDecl:    declStmt,
		Procedures: procedures,
		Functions:  functions,
		Statement:  stmt,
	}, nil
}

// <type definition> ::= <identifier> = <type>
// <type> ::= <simple type> | <structured type> | <pointer type>
// <simple type> ::= <scalar type> | <subrange type> | <type identifier>
// <scalar type> ::= (<identifier> {,<identifier>})
// <subrange type> ::= <constant> .. <constant>
// <type identifier> ::= <identifier>
func (p *Parser) parseTypeDeclStmt() (*ast.TypeDeclStmt, error) {
	if !p.match(token.TYPE) {
		return nil, nil
	}
	stmt := &ast.TypeDeclStmt{
		Token: p.previous(),
	}
	defs, err := p.parseTypeDefinitionStmts()
	if err != nil {
		return nil, err
	}
	stmt.Definitions = defs
	return stmt, nil
}

func (p *Parser) parseTypeDefinitionStmts() ([]ast.TypeDefinitionStmt, error) {
	defs := []ast.TypeDefinitionStmt{}
	for p.match(token.IDENT) {
		tok := p.previous()
		if _, err := p.consume(token.EQUAL); err != nil {
			return nil, err
		}
		def, err := p.parseTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		defs = append(defs, ast.TypeDefinitionStmt{
			Token: tok,
			Identifier: &ast.IdentifierExpr{
				Token: tok,
				Value: tok.Literal,
			},
			Definition: def,
		})
		if !p.match(token.SEMICOLON) {
			break
		}
	}
	return defs, nil
}

func (p *Parser) parseTypeDefinitionExpr() (ast.TypeDefinitionExprIntf, error) {
	if p.match(token.PACKED, token.ARRAY) {
		// <array type> ::= array [<index type>{,<index type>}] of <component type>
		packed := false
		tok := p.previous()
		if p.previous().Type == token.PACKED {
			packed = true
			if _, err := p.consume(token.ARRAY); err != nil {
				return nil, err
			}
		}
		if _, err := p.consume(token.LBRACKET); err != nil {
			return nil, err
		}
		indexTypeDef, err := p.parseSimpleTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.RBRACKET); err != nil {
			return nil, err
		}
		if _, err := p.consume(token.OF); err != nil {
			return nil, err
		}
		componentTypeDef, err := p.parseTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		return &ast.ArrayTypeDefinitionExpr{
			Token:            tok,
			Packed:           packed,
			IndexTypeDef:     indexTypeDef,
			ComponentTypeDef: componentTypeDef,
		}, nil
	} else if p.match(token.RECORD) {
		// <record type> ::= record <field list> end
		// <field list> ::= <fixed part> | <fixed part> ; <variant part> | <variant part>
		// <fixed part> ::= <record section> {;<record section>}
		// <record section> ::= <field identifier> {, <field identifier>} : <type> | <empty>
		tok := p.previous()
		defs := []ast.TypeDefinitionStmt{}
		for p.match(token.IDENT) {
			tok := p.previous()
			if _, err := p.consume(token.COLON); err != nil {
				return nil, err
			}
			def, err := p.parseTypeDefinitionExpr()
			if err != nil {
				return nil, err
			}
			defs = append(defs, ast.TypeDefinitionStmt{
				Token: tok,
				Identifier: &ast.IdentifierExpr{
					Token: tok,
					Value: tok.Literal,
				},
				Definition: def,
			})
			if !p.match(token.SEMICOLON) {
				break
			}
		}
		if _, err := p.consume(token.END); err != nil {
			return nil, err
		}
		return &ast.RecordTypeDefinitionExpr{
			Token:  tok,
			Fields: defs,
		}, nil
	} else if p.match(token.SET) {
		// <set type> ::=set of <base type>
		// <base type> ::= <simple type>
		tok := p.previous()
		if _, err := p.consume(token.OF); err != nil {
			return nil, err
		}
		baseType, err := p.parseSimpleTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		return &ast.SetTypeDefinitionExpr{
			Token:       tok,
			BaseTypeDef: baseType,
		}, nil
	} else if p.match(token.FILE) {
		// <file type> ::= file of <type>
		tok := p.previous()
		if _, err := p.consume(token.OF); err != nil {
			return nil, err
		}
		baseType, err := p.parseTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		return &ast.FileTypeDefinitionExpr{
			Token:       tok,
			BaseTypeDef: baseType,
		}, nil
	}

	return p.parseSimpleTypeDefinitionExpr()
}

func (p *Parser) parseSimpleTypeDefinitionExpr() (ast.TypeDefinitionExprIntf, error) {
	if p.match(token.LPAREN) {
		// scalar type
		tok := p.previous()
		values := []*ast.IdentifierExpr{}
		for p.match(token.IDENT) {
			vtok := p.previous()
			val := &ast.IdentifierExpr{
				Token: vtok,
				Value: vtok.Literal,
			}
			values = append(values, val)
			if !p.match(token.COMMA) {
				break
			}
		}
		return &ast.EnumTypeDefinitionExpr{
			Token:  tok,
			Values: values,
		}, nil
	}
	// ident or subrange type
	left, err := p.parseSimpleExpression()
	tok := p.previous()
	if err != nil {
		return nil, err
	}
	if p.match(token.DOTDOT) {
		right, err := p.parseSimpleExpression()
		if err != nil {
			return nil, err
		}
		return &ast.SubrangeTypeDefinitionExpr{
			Token: tok,
			Left:  left,
			Right: right,
		}, nil
	}
	return &ast.SimpleTypeDefinitionExpr{
		Token:      tok,
		Identifier: left,
	}, nil
}

func (p *Parser) parseProcedureAndFunctionDeclStmts() ([]*ast.ProcedureDeclStmt, []*ast.FunctionDeclStmt, error) {
	procedures := []*ast.ProcedureDeclStmt{}
	functions := []*ast.FunctionDeclStmt{}
	for {
		if p.match(token.PROCEDURE) {
			procedure, err := p.parseProcedureDecl()
			if err != nil {
				return nil, nil, err
			}
			procedures = append(procedures, procedure)
		} else if p.match(token.FUNCTION) {
			function, err := p.parseFunctionDecl()
			if err != nil {
				return nil, nil, err
			}
			functions = append(functions, function)
		} else {
			break
		}
	}
	return procedures, functions, nil
}

// <procedure declaration> ::= <procedure heading> <block>
//
// <procedure heading> ::= procedure <identifier> ; |
// 		procedure <identifier> ( <formal parameter section> {;<formal parameter section>} );
//
// <formal parameter section> ::= <parameter group> | var <parameter group> |
// 		function <parameter group> | procedure <identifier> { , <identifier>}
//
// <parameter group> ::= <identifier> {, <identifier>} : <type identifier>
func (p *Parser) parseProcedureDecl() (*ast.ProcedureDeclStmt, error) {
	tok := p.previous()
	ident, err := p.consume(token.IDENT)
	if err != nil {
		return nil, err
	}
	args, err := p.parseArgumentListWithTypes()
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
	return &ast.ProcedureDeclStmt{
		Token: tok,
		Identifier: &ast.IdentifierExpr{
			Token: ident,
			Value: ident.Literal,
		},
		Args: args,
		Body: block,
	}, nil
}

// <function declaration> ::= <function heading> <block>
// <function heading> ::= function <identifier> : <result type> ; |
// 		function <identifier> ( <formal parameter section> {;<formal parameter section>} ) : <result type> ;
// <result type> ::= <type identifier>
func (p *Parser) parseFunctionDecl() (*ast.FunctionDeclStmt, error) {
	tok := p.previous()
	ident, err := p.consume(token.IDENT)
	if err != nil {
		return nil, err
	}
	args, err := p.parseArgumentListWithTypes()
	if err != nil {
		return nil, err
	}
	if _, err := p.consume(token.COLON); err != nil {
		return nil, err
	}
	typident, err := p.consume(token.IDENT)
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
	if _, err := p.consume(token.SEMICOLON); err != nil {
		return nil, err
	}
	return &ast.FunctionDeclStmt{
		Token: tok,
		Identifier: &ast.IdentifierExpr{
			Token: ident,
			Value: ident.Literal,
		},
		ReturnType: &ast.IdentifierExpr{
			Token: typident,
			Value: typident.Literal,
		},
		Args: args,
		Body: block,
	}, nil
}

func (p *Parser) parseArgumentListWithTypes() ([]ast.FormalArg, error) {
	args := []ast.FormalArg{}
	if p.match(token.LPAREN) {
		for {
			idents := []*ast.IdentifierExpr{}
			for {
				varident, err := p.consume(token.IDENT)
				if err != nil {
					return nil, err
				}
				idents = append(idents, &ast.IdentifierExpr{
					Token: varident,
					Value: varident.Literal,
				})
				if !p.match(token.COMMA) {
					break
				}
			}
			if _, err := p.consume(token.COLON); err != nil {
				return nil, err
			}
			typident, err := p.consume(token.IDENT)
			if err != nil {
				return nil, err
			}
			for _, ident := range idents {
				args = append(args, ast.FormalArg{
					Identifer: ident,
					Type: &ast.IdentifierExpr{
						Token: typident,
						Value: typident.Literal,
					},
				})
			}
			if !p.match(token.COMMA) {
				break
			}
		}
		if _, err := p.consume(token.RPAREN); err != nil {
			return nil, err
		}
	}
	return args, nil
}

func (p *Parser) parseVarDeclStmt() (*ast.VarDeclStmt, error) {
	if !p.match(token.VAR) {
		return nil, nil
	}
	tok := p.previous()
	mappings := make(map[string]ast.TypeDefinitionExprIntf)
Decl:
	for {
		idents := []string{}
		for {
			if !p.match(token.IDENT) {
				break Decl
			}
			ident := p.previous()
			idents = append(idents, ident.Literal)
			if p.match(token.COLON) {
				break
			}
			if _, err := p.consume(token.COMMA); err != nil {
				return nil, err
			}
		}
		typ, err := p.parseTypeDefinitionExpr()
		if err != nil {
			return nil, err
		}
		for _, ident := range idents {
			if _, ok := mappings[ident]; ok {
				return nil, fmt.Errorf("Duplicate variable declaration: %s", ident)
			}
			mappings[ident] = typ
		}
		if !p.match(token.SEMICOLON) {
			break
		}
	}
	return &ast.VarDeclStmt{
		Token:        tok,
		Declarations: mappings,
	}, nil
}

func (p *Parser) parseCompoundStmt() (*ast.CompoundStmt, error) {
	tok, err := p.consume(token.BEGIN)
	if err != nil {
		return nil, err
	}
	stmts := []ast.Statement{}
	for {
		stmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		stmts = append(stmts, stmt)
		if !p.match(token.SEMICOLON) {
			break
		}
	}
	if _, err := p.consume(token.END); err != nil {
		return nil, err
	}
	return &ast.CompoundStmt{
		Token:      tok,
		Statements: stmts,
	}, nil
}

// <statement> ::= <unlabelled statement> | <label> : <unlabelled statement>
// <unlabelled statement> ::= <simple statement> | <structured statement>
// <structured statement> ::= <compound statement> | <conditional statement> |
//							  <repetitive statement> | <with statement>
// <compound statement> ::= begin <statement> {; <statement> } end;
// <conditional statement> ::= <if statement> | <case statement>
// <if statement> ::= if <expression> then <statement> |
//					  if <expression> then <statement> else <statement>
// <case statement> ::= case <expression> of <case list element> {; <case list element> } end
// <case list element> ::= <case label list> : <statement> | <empty>
// <case label list> ::= <case label> {, <case label> }
// <repetitive statement> ::= <while statement> | <repeat statemant> | <for statement>
// <while statement> ::= while <expression> do <statement>
// <repeat statement> ::= repeat <statement> {; <statement>} until <expression>
// <for statement> ::= for <control variable> := <for list> do <statement>
// <with statement> ::= with <record variable list> do <statement>
//
func (p *Parser) parseStmt() (ast.Statement, error) {
	if p.check(token.BEGIN) {
		return p.parseCompoundStmt()
	}
	if p.match(token.IF) {
		tok := p.previous()
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.THEN); err != nil {
			return nil, err
		}
		thenstmt, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		ifstmt := &ast.IfStmt{
			Token: tok,
			Cond:  expr,
			Then:  thenstmt,
		}
		if p.match(token.ELSE) {
			elsestmt, err := p.parseStmt()
			if err != nil {
				return nil, err
			}
			ifstmt.Else = elsestmt
		}
		return ifstmt, nil
	}
	if p.match(token.WHILE) {
		tok := p.previous()
		invariant, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(token.DO); err != nil {
			return nil, err
		}
		body, err := p.parseStmt()
		if err != nil {
			return nil, err
		}
		return &ast.WhileStmt{
			Token:     tok,
			Invariant: invariant,
			Body:      body,
		}, nil
	}
	if p.match(token.REPEAT) {
		tok := p.previous()
		stmts := []ast.Statement{}
		for {
			stmt, err := p.parseStmt()
			if err != nil {
				return nil, err
			}
			stmts = append(stmts, stmt)
			if !p.match(token.SEMICOLON) {
				break
			}
		}
		if _, err := p.consume(token.UNTIL); err != nil {
			return nil, err
		}
		invariant, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		return &ast.RepeatStmt{
			Token:      tok,
			Invariant:  invariant,
			Statements: stmts,
		}, nil
	}
	//TODO: for stmt
	//TODO: case stmt
	//TODO: with stmt
	return p.parseSimpleStmt()
}

// <simple statement> ::= <assignment statement> | <procedure statement> |
//						  <go to statement> | <empty statement>
// <assignment statement> ::= <variable> := <expression> |
// 							  <function identifier> := <expression>
// <procedure statement> ::= <procedure identifier> |
//							 <procedure identifier> (<actual parameter> {, <actual parameter> })
// <actual parameter> ::= <expression> | <variable> | <procedure identifier> |
//						  <function identifier>
//
func (p *Parser) parseSimpleStmt() (ast.Statement, error) {
	if p.match(token.IDENT) {
		ident := p.previous()
		if p.match(token.COLON) {
			// labeled stmt
			stmt, err := p.parseSimpleStmt()
			if err != nil {
				return nil, err
			}
			return &ast.LabeledStmt{
				Label: &ast.IdentifierExpr{
					Token: ident,
					Value: ident.Literal,
				},
				Stmt: stmt,
			}, nil
		} else if p.match(token.NAMED) {
			expr, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			return &ast.AssignmentStmt{
				Identifier: &ast.IdentifierExpr{
					Token: ident,
					Value: ident.Literal,
				},
				Expr: expr,
			}, nil
		}
		args := []ast.Expression{}
		// procedure call
		if p.match(token.LPAREN) {
			for {
				expr, err := p.parseExpression()
				if err != nil {
					return nil, err
				}
				args = append(args, expr)
				if !p.match(token.COMMA) {
					break
				}
			}
			if _, err := p.consume(token.RPAREN); err != nil {
				return nil, err
			}
		}
		return &ast.ProcedureCallStmt{
			Token: ident,
			Identifier: &ast.IdentifierExpr{
				Token: ident,
				Value: ident.Literal,
			},
			Args: args,
		}, nil
	}
	if p.match(token.GOTO) {
		label, err := p.consume(token.IDENT)
		if err != nil {
			return nil, err
		}
		return &ast.GotoStmt{
			Label: ast.IdentifierExpr{
				Token: label,
				Value: label.Literal,
			},
		}, nil
	}
	return nil, fmt.Errorf("Unable to parse statement around: %s", p.peek().Literal)
}

// <factor> ::= <variable> | <unsigned constant> | ( <expression> ) |
//				<function designator> | <set> | not <factor>
func (p *Parser) parseFactor() (ast.Expression, error) {
	if p.match(token.IDENT) {
		ident := &ast.IdentifierExpr{
			Token: p.previous(),
			Value: p.previous().Literal,
		}
		if p.match(token.LPAREN) {
			// function call
			args := []ast.Expression{}
			for {
				expr, err := p.parseExpression()
				if err != nil {
					return nil, err
				}
				args = append(args, expr)
				if !p.match(token.COMMA) {
					break
				}
			}
			if _, err := p.consume(token.RPAREN); err != nil {
				return nil, err
			}
			return &ast.FunctionCallExpr{
				Token:      ident.Token,
				Identifier: ident,
				Args:       args,
			}, nil
		}
		return ident, nil
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
	if p.match(token.LBRACKET) {
		elements, err := p.parseElementList()
		if err != nil {
			return nil, err
		}
		set := &ast.SetExpr{
			Elements: elements,
		}
		if _, err := p.consume(token.RBRACKET); err != nil {
			return nil, err
		}
		return set, nil
	}
	return nil, fmt.Errorf("Unknown literal token: %s (%s)",
		p.peek().Literal, p.peek().Type)
}

// <element list> ::= <element> {, <element> } | <empty>
// 		<element> ::= <expression> | <expression> .. <expression>
func (p *Parser) parseElementList() ([]ast.Expression, error) {
	elements := []ast.Expression{}
	expr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}
	for {
		if p.match(token.COMMA) {
			elements = append(elements, expr)
			expr, err = p.parseExpression()
			if err != nil {
				return nil, err
			}
		} else if p.match(token.DOTDOT) {
			element := &ast.ElementExpr{
				Left: expr,
			}
			expr, err = p.parseExpression()
			if err != nil {
				return nil, err
			}
			element.Right = expr
			expr = element
		} else {
			break
		}
	}
	elements = append(elements, expr)
	return elements, nil
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
