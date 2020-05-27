package printer

import (
	"bytes"
	"fmt"
	"strings"

	"yapsi/pkg/ast"
)

type AstPrinter struct{}

var _ ast.NodeVisitor = (*AstPrinter)(nil)

func (p *AstPrinter) VisitNumericLiteral(node *ast.NumericLiteral) ast.VisitorResult {
	return fmt.Sprintf("%s", node.Value)
}

func (p *AstPrinter) VisitStringLiteral(node *ast.StringLiteral) ast.VisitorResult {
	return "'" + node.Value + "'"
}

func (p *AstPrinter) VisitBoolLiteral(node *ast.BoolLiteral) ast.VisitorResult {
	return fmt.Sprintf("%t", node.Value)
}

func (p *AstPrinter) VisitCharLiteral(node *ast.CharLiteral) ast.VisitorResult {
	return fmt.Sprintf("%c", node.Value)
}

func (p *AstPrinter) VisitIdentifierExpr(node *ast.IdentifierExpr) ast.VisitorResult {
	return node.Value
}

func (p *AstPrinter) VisitUnaryExpr(node *ast.UnaryExpr) ast.VisitorResult {
	return node.Operator.Literal + node.Expr.Visit(p).(string)
}

func (p *AstPrinter) VisitBinaryExpr(node *ast.BinaryExpr) ast.VisitorResult {
	var out bytes.Buffer
	left := node.Left.Visit(p)
	right := node.Right.Visit(p)
	out.WriteString(left.(string))
	out.WriteString(" " + node.Operator.Literal + " ")
	out.WriteString(right.(string))
	return out.String()
}

func (p *AstPrinter) VisitElementExpr(node *ast.ElementExpr) ast.VisitorResult {
	left := node.Left.Visit(p)
	right := node.Right.Visit(p)
	return left.(string) + " .. " + right.(string)
}

func (p *AstPrinter) VisitSetExpr(node *ast.SetExpr) ast.VisitorResult {
	chunks := make([]string, 0, len(node.Elements))
	for _, el := range node.Elements {
		chunk := el.Visit(p)
		chunks = append(chunks, chunk.(string))
	}
	return "[" + strings.Join(chunks, ", ") + "]"
}

func (p *AstPrinter) VisitAssignmentStmt(node *ast.AssignmentStmt) ast.VisitorResult {
	ident := node.Identifier.Visit(p)
	expr := node.Expr.Visit(p)
	return ident.(string) + " := " + expr.(string)
}

func (p *AstPrinter) VisitGotoStmt(node *ast.GotoStmt) ast.VisitorResult {
	label := node.Label.Visit(p)
	return "goto " + label.(string)
}

func (p *AstPrinter) VisitLabeledStmt(node *ast.LabeledStmt) ast.VisitorResult {
	label := node.Label.Visit(p)
	stmt := node.Stmt.Visit(p)
	return label.(string) + " : " + stmt.(string)
}

func (p *AstPrinter) VisitCallStmt(node *ast.CallStmt) ast.VisitorResult {
	var out bytes.Buffer
	ident := node.Identifier.Visit(p)
	out.WriteString(ident.(string))
	if len(node.Args) > 0 {
		out.WriteRune('(')
		args := make([]string, 0, len(node.Args))
		for _, arg := range node.Args {
			s := arg.Visit(p)
			args = append(args, s.(string))
		}
		out.WriteString(strings.Join(args, ", "))
		out.WriteRune(')')
	}
	return out.String()
}

func (p *AstPrinter) VisitCompoundStmt(node *ast.CompoundStmt) ast.VisitorResult {
	var out bytes.Buffer
	out.WriteString("begin\n")
	chunks := make([]string, 0, len(node.Statements))
	for _, stmt := range node.Statements {
		chunk := stmt.Visit(p)
		chunks = append(chunks, chunk.(string))
	}
	out.WriteString(strings.Join(chunks, ";\n"))
	out.WriteString("\nend")
	return out.String()
}

func (p *AstPrinter) VisitIfStmt(node *ast.IfStmt) ast.VisitorResult {
	var out bytes.Buffer
	expr := node.Expr.Visit(p)
	out.WriteString("if " + expr.(string) + " then\n")
	thenstr := node.Then.Visit(p)
	out.WriteString(thenstr.(string))
	if node.Else != nil {
		out.WriteString("\nelse\n")
		elsestr := node.Else.Visit(p)
		out.WriteString(elsestr.(string))
	}
	return out.String()
}

func (p *AstPrinter) VisitWhileStmt(node *ast.WhileStmt) ast.VisitorResult {
	var out bytes.Buffer
	invariant := node.Invariant.Visit(p)
	body := node.Body.Visit(p)
	out.WriteString("while " + invariant.(string) + " do\n")
	out.WriteString(body.(string))
	return out.String()
}

func (p *AstPrinter) VisitRepeatStmt(node *ast.RepeatStmt) ast.VisitorResult {
	var out bytes.Buffer
	invariant := node.Invariant.Visit(p)
	chunks := make([]string, 0, len(node.Statements))
	for _, stmt := range node.Statements {
		chunks = append(chunks, stmt.Visit(p).(string))
	}
	out.WriteString("repeat\n")
	out.WriteString(strings.Join(chunks, ";\n"))
	out.WriteString(" until " + invariant.(string))
	return out.String()
}

func (p *AstPrinter) VisitProgramStmt(node *ast.ProgramStmt) ast.VisitorResult {
	var out bytes.Buffer
	out.WriteString("program " + node.Identifier.Visit(p).(string) + ";\n\n")
	out.WriteString(node.Block.Visit(p).(string) + ".")
	return out.String()
}

func (p *AstPrinter) VisitBlockStmt(node *ast.BlockStmt) ast.VisitorResult {
	var out bytes.Buffer
	out.WriteString(node.Statement.Visit(p).(string))
	return out.String()
}
