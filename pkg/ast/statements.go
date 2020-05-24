package ast

/*
type ProgramStmt struct {
	Header *ProgramHeaderStmt
	Block  *BlockStmt
}

func (p *ProgramStmt) statementNode() {}
func (p *ProgramStmt) TokenLiteral() string {
	if p.Header != nil {
		return p.Header.TokenLiteral()
	}
	return ""
}
func (p *ProgramStmt) String() string {
	var out bytes.Buffer
	if p.Header != nil {
		out.WriteString(p.Header.String())
	}
	if p.Block != nil {
		out.WriteString(p.Block.String())
	}
	out.WriteRune('.')
	return out.String()
}

type ProgramHeaderStmt struct {
	Identifier *IdentifierExpr
	Params     []ParamExpr
}

func (p *ProgramHeaderStmt) statementNode() {}
func (p *ProgramHeaderStmt) TokenLiteral() string {
	if p.Identifier != nil {
		return p.Identifier.TokenLiteral
	}
	return ""
}
func (p *ProgramHeaderStmt) String() string {
	var out bytes.Buffer
	out.WriteString("Program")
	if p.Identifier != nil {
		out.WriteRune(' ')
		out.WriteString(p.Identifier.String())
	}
	if len(p.Params) > 0 {
		params := make([]string, 0, len(p.Params))
		for _, param := range p.Params {
			params = append(params, param.String())
		}
		out.WriteRune('(')
		out.WriteString(string.Join(params, ", "))
		out.WriteRune(')')
	}
	out.WriteRune(';')
}

type BlockStmt struct {
	Labels    []LabelDeclStmt
	Constants []ConstDeclStmt
	Types     []TypeDeclStmt
	Variables []VarDeclStmt
	Callables []CallDeclStmt
	Body      []BodyStatement
}

func (b *BlockStmt) statementNode() {}
func (b *BlockStmt) TokenLiteral() string {
	if len(b.Body) > 0 {
		return b.Body[0].TokenLiteral()
	}
}
func (b *BlockStmt) String() string {
	var out bytes.Buffer
	out.WriteString("Begin\n")
	for _, stmt := range b.Body {
		out.WriteString(stmt.String())
	}
	out.WriteString("End")
}

type LabelDeclStmt struct {
	Identifier *IdentifierExpr
}

func (l *LabelDeclStmt) statementNode()       {}
func (l *LabelDeclStmt) TokenLiteral() string { return l.Identifier.TokenLiteral() }
func (l *LabelDeclStmt) String() string       { return l.Identifier.String() }

type ConstDeclStmt struct {
	Identifier *IdentifierExpr
	Value      *ConstValExpr
}

func (c *ConstDeclStmt) statementNode() {}
func (c *ConstDeclStmt) TokenLiteral() string {
	if c.Identifier != nil {
		return c.Identifier.TokenLiteral()
	}
	return ""
}
func (c *ConstDeclStmt) String() string {
	var out bytes.Buffer
	if c.Identifier == nil || c.Value == nil {
		return out.String()
	}
	c.WriteString(c.Identifier.String())
	c.WriteString(" = ")
}
*/
