package printer

import (
	"bytes"
	"fmt"
	"strings"

	"yapsi/pkg/ast"
)

type AstPrinter struct{}

var _ ast.NodeVisitor = (*AstPrinter)(nil)

func (p *AstPrinter) VisitNumericLiteral(node *ast.NumericLiteral) (ast.VisitorResult, error) {
	return fmt.Sprintf("%s", node.Value), nil
}

func (p *AstPrinter) VisitStringLiteral(node *ast.StringLiteral) (ast.VisitorResult, error) {
	return "'" + node.Value + "'", nil
}

func (p *AstPrinter) VisitBoolLiteral(node *ast.BoolLiteral) (ast.VisitorResult, error) {
	return fmt.Sprintf("%t", node.Value), nil
}

func (p *AstPrinter) VisitCharLiteral(node *ast.CharLiteral) (ast.VisitorResult, error) {
	return fmt.Sprintf("%c", node.Value), nil
}

func (p *AstPrinter) VisitIdentifierExpr(node *ast.IdentifierExpr) (ast.VisitorResult, error) {
	return node.Value, nil
}

func (p *AstPrinter) VisitUnaryExpr(node *ast.UnaryExpr) (ast.VisitorResult, error) {
	expr, _ := node.Expr.Visit(p)
	return node.Operator.Literal + expr.(string), nil
}

func (p *AstPrinter) VisitBinaryExpr(node *ast.BinaryExpr) (ast.VisitorResult, error) {
	var out bytes.Buffer
	left, _ := node.Left.Visit(p)
	right, _ := node.Right.Visit(p)
	out.WriteString(left.(string))
	out.WriteString(" " + node.Operator.Literal + " ")
	out.WriteString(right.(string))
	return out.String(), nil
}

func (p *AstPrinter) VisitElementExpr(node *ast.ElementExpr) (ast.VisitorResult, error) {
	left, _ := node.Left.Visit(p)
	right, _ := node.Right.Visit(p)
	return left.(string) + " .. " + right.(string), nil
}

func (p *AstPrinter) VisitSetExpr(node *ast.SetExpr) (ast.VisitorResult, error) {
	chunks := make([]string, 0, len(node.Elements))
	for _, el := range node.Elements {
		chunk, _ := el.Visit(p)
		chunks = append(chunks, chunk.(string))
	}
	return "[" + strings.Join(chunks, ", ") + "]", nil
}

func (p *AstPrinter) VisitAssignmentStmt(node *ast.AssignmentStmt) (ast.VisitorResult, error) {
	ident, _ := node.Identifier.Visit(p)
	expr, _ := node.Expr.Visit(p)
	return ident.(string) + " := " + expr.(string), nil
}

func (p *AstPrinter) VisitGotoStmt(node *ast.GotoStmt) (ast.VisitorResult, error) {
	label, _ := node.Label.Visit(p)
	return "goto " + label.(string), nil
}

func (p *AstPrinter) VisitLabeledStmt(node *ast.LabeledStmt) (ast.VisitorResult, error) {
	label, _ := node.Label.Visit(p)
	stmt, _ := node.Stmt.Visit(p)
	return label.(string) + " : " + stmt.(string), nil
}

func (p *AstPrinter) VisitCallStmt(node *ast.CallStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	ident, _ := node.Identifier.Visit(p)
	out.WriteString(ident.(string))
	if len(node.Args) > 0 {
		out.WriteRune('(')
		args := make([]string, 0, len(node.Args))
		for _, arg := range node.Args {
			s, _ := arg.Visit(p)
			args = append(args, s.(string))
		}
		out.WriteString(strings.Join(args, ", "))
		out.WriteRune(')')
	}
	return out.String(), nil
}

func (p *AstPrinter) VisitCompoundStmt(node *ast.CompoundStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	out.WriteString("begin\n")
	chunks := make([]string, 0, len(node.Statements))
	for _, stmt := range node.Statements {
		chunk, _ := stmt.Visit(p)
		chunks = append(chunks, chunk.(string))
	}
	out.WriteString(strings.Join(chunks, ";\n"))
	out.WriteString("\nend")
	return out.String(), nil
}

func (p *AstPrinter) VisitIfStmt(node *ast.IfStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	expr, _ := node.Cond.Visit(p)
	out.WriteString("if " + expr.(string) + " then\n")
	thenstr, _ := node.Then.Visit(p)
	out.WriteString(thenstr.(string))
	if node.Else != nil {
		out.WriteString("\nelse\n")
		elsestr, _ := node.Else.Visit(p)
		out.WriteString(elsestr.(string))
	}
	return out.String(), nil
}

func (p *AstPrinter) VisitWhileStmt(node *ast.WhileStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	invariant, _ := node.Invariant.Visit(p)
	body, _ := node.Body.Visit(p)
	out.WriteString("while " + invariant.(string) + " do\n")
	out.WriteString(body.(string))
	return out.String(), nil
}

func (p *AstPrinter) VisitRepeatStmt(node *ast.RepeatStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	invariant, _ := node.Invariant.Visit(p)
	chunks := make([]string, 0, len(node.Statements))
	for _, stmt := range node.Statements {
		ch, _ := stmt.Visit(p)
		chunks = append(chunks, ch.(string))
	}
	out.WriteString("repeat\n")
	out.WriteString(strings.Join(chunks, ";\n"))
	out.WriteString(" until " + invariant.(string))
	return out.String(), nil
}

func (p *AstPrinter) VisitProgramStmt(node *ast.ProgramStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	ident, _ := node.Identifier.Visit(p)
	out.WriteString("program " + ident.(string) + ";\n\n")
	stmt, _ := node.Block.Visit(p)
	out.WriteString(stmt.(string) + ".")
	return out.String(), nil
}

func (p *AstPrinter) VisitBlockStmt(node *ast.BlockStmt) (ast.VisitorResult, error) {
	var out bytes.Buffer
	res, _ := node.Statement.Visit(p)
	out.WriteString(res.(string))
	return out.String(), nil
}
