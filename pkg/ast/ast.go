package ast

import (
	"bytes"
	"strings"

	"yapsi/pkg/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Name       string
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var buf bytes.Buffer
	for _, stmt := range p.Statements {
		buf.WriteString(stmt.String() + "\n")
	}
	return buf.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ProgramDeclarationStatement struct {
	Token token.Token
	Name  string
}

func (pds *ProgramDeclarationStatement) statementNode()       {}
func (pds *ProgramDeclarationStatement) TokenLiteral() string { return pds.Token.Literal }
func (pds *ProgramDeclarationStatement) String() string {
	return "PROGRAM " + pds.Name + ";"
}

type LabelStatement struct {
	Token  token.Token
	Labels []*Identifier
}

func (ls *LabelStatement) statementNode()       {}
func (ls *LabelStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LabelStatement) String() string {
	var buf bytes.Buffer
	tab := "    "
	if len(ls.Labels) > 0 {
		buf.WriteString("LABEL\n")
		labels := make([]string, 0, len(ls.Labels))
		for _, label := range ls.Labels {
			labels = append(labels, tab+label.String())
		}
		buf.WriteString(strings.Join(labels, ",\n"))
		buf.WriteString(";")
	}
	return buf.String()
}
