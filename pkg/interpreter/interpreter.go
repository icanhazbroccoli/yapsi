package interpreter

import (
	"fmt"
	"yapsi/pkg/ast"
	"yapsi/pkg/builtin"
	"yapsi/pkg/object"
	"yapsi/pkg/types"
)

type Interpreter struct {
	env *object.Environment
}

var _ ast.NodeVisitor = (*Interpreter)(nil)

func (i *Interpreter) Evaluate(p *ast.ProgramStmt) error {
	env := object.NewEnvironment(nil)
	declareBaseTypes(env)
	declareBuiltins(env)
	i.env = env
	_, err := p.Visit(i)
	return err
}

func (i *Interpreter) VisitNumericLiteral(node *ast.NumericLiteral) (ast.VisitorResult, error) {
	if object.LooksLikeInt(node.Value) {
		intval, err := object.ParseInt(node.Value)
		if err != nil {
			return nil, err
		}
		return &object.Integer{
			Value: intval,
		}, nil
	}
	realval, err := object.ParseReal(node.Value)
	if err != nil {
		return nil, err
	}
	return &object.Real{
		Value: realval,
	}, nil
}

func (i *Interpreter) VisitStringLiteral(node *ast.StringLiteral) (ast.VisitorResult, error) {
	return &object.String{
		Value: node.Value,
	}, nil
}

func (i *Interpreter) VisitBoolLiteral(node *ast.BoolLiteral) (ast.VisitorResult, error) {
	return &object.Boolean{
		Value: node.Value,
	}, nil
}

func (i *Interpreter) VisitCharLiteral(node *ast.CharLiteral) (ast.VisitorResult, error) {
	return &object.Character{
		Value: node.Value,
	}, nil
}

func (i *Interpreter) VisitIdentifierExpr(node *ast.IdentifierExpr) (ast.VisitorResult, error) {
	lookup, ok := i.env.LookupVar(object.VarName(node.Value))
	if !ok {
		return nil, undefinedIdentErr(node.Value)
	}
	return lookup, nil
}

func (i *Interpreter) VisitUnaryExpr(node *ast.UnaryExpr) (ast.VisitorResult, error) {
	expr, err := node.Expr.Visit(i)
	if err != nil {
		return nil, err
	}
	switch node.Operator.Literal {
	case "not":
		if lg, ok := expr.(object.Logical); ok {
			return lg.OpNot()
		}
	case "-":
		if ar, ok := expr.(object.Arithmetic); ok {
			return ar.OpUnMinus()
		}
	case "+":
		if ar, ok := expr.(object.Arithmetic); ok {
			return ar.OpUnPlus()
		}
	}
	return nil, unsupportedUnaryOpErr(node.Operator.Literal)
}

func (i *Interpreter) VisitBinaryExpr(node *ast.BinaryExpr) (ast.VisitorResult, error) {
	lr, err := node.Left.Visit(i)
	if err != nil {
		return nil, err
	}
	rr, err := node.Right.Visit(i)
	if err != nil {
		return nil, err
	}
	left, right := lr.(object.Any), rr.(object.Any)
	switch node.Operator.Literal {
	case "+":
		if ar, ok := left.(object.Arithmetic); ok {
			return ar.OpPlus(right.(object.Any))
		}
	case "-":
		if ar, ok := left.(object.Arithmetic); ok {
			return ar.OpMinus(right.(object.Any))
		}
	case "*":
		if ar, ok := left.(object.Arithmetic); ok {
			return ar.OpAsterisk(right.(object.Any))
		}
	case "/":
		if ar, ok := left.(object.Arithmetic); ok {
			return ar.OpSlash(right.(object.Any))
		}
	case "%":
		if ar, ok := left.(object.Arithmetic); ok {
			return ar.OpPercent(right.(object.Any))
		}
	case "=":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpEql(right.(object.Any))
		}
	case "<>":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpNeql(right.(object.Any))
		}
	case ">":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpGt(right.(object.Any))
		}
	case ">=":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpGte(right.(object.Any))
		}
	case "<":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpLt(right.(object.Any))
		}
	case "<=":
		if cp, ok := left.(object.Comparable); ok {
			return cp.OpLte(right.(object.Any))
		}
	case "and":
		if lg, ok := left.(object.Logical); ok {
			return lg.OpAnd(right.(object.Any))
		}
	case "or":
		if lg, ok := left.(object.Logical); ok {
			return lg.OpOr(right.(object.Any))
		}
	}
	return nil, unsupportedBinaryOpErr(node.Operator.Literal, left.Type(), right.Type())
}

func (i *Interpreter) VisitElementExpr(*ast.ElementExpr) (ast.VisitorResult, error) {
	//TODO
	panic("not implemented")
}

func (i *Interpreter) VisitSetExpr(*ast.SetExpr) (ast.VisitorResult, error) {
	//TODO
	panic("not implemented")
}

func (i *Interpreter) VisitProgramStmt(node *ast.ProgramStmt) (ast.VisitorResult, error) {
	return node.Block.Visit(i)
}

func (i *Interpreter) VisitBlockStmt(node *ast.BlockStmt) (ast.VisitorResult, error) {
	if _, err := node.VarDecl.Visit(i); err != nil {
		return nil, err
	}
	return node.Statement.Visit(i)
}

func (i *Interpreter) VisitCompoundStmt(node *ast.CompoundStmt) (ast.VisitorResult, error) {
	for _, stmt := range node.Statements {
		_, err := stmt.Visit(i)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitAssignmentStmt(node *ast.AssignmentStmt) (ast.VisitorResult, error) {
	ident := node.Identifier.Value
	r, err := node.Expr.Visit(i)
	if err != nil {
		return nil, err
	}
	value, ok := r.(object.Any)
	if !ok {
		return nil, unexpectedVisitorResultTypeErr(r, "object.Any")
	}
	if err := i.env.AssignVar(object.VarName(ident), value); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *Interpreter) VisitGotoStmt(*ast.GotoStmt) (ast.VisitorResult, error) {
	//TODO
	panic("not implemented")
}

func (i *Interpreter) VisitLabeledStmt(node *ast.LabeledStmt) (ast.VisitorResult, error) {
	//TODO
	return node.Stmt.Visit(i)
}

func (i *Interpreter) VisitCallStmt(node *ast.CallStmt) (ast.VisitorResult, error) {
	ident := node.Identifier.Value
	lr, ok := i.env.LookupVar(object.VarName(ident))
	if !ok {
		return nil, undefinedIdentErr(ident)
	}
	callable, ok := lr.(object.Callable)
	if !ok {
		return nil, uncallableEntityErr(ident)
	}
	args := make([]object.Any, 0, len(node.Args))
	for _, arg := range node.Args {
		ar, err := arg.Visit(i)
		if err != nil {
			return nil, err
		}
		args = append(args, ar.(object.Any))
	}
	if callable.Returns() != nil {
		//function
		return callable.CallReturn(i.env, args...)
	}
	//procedure
	err := callable.Call(i.env, args...)
	return nil, err
}

func (i *Interpreter) VisitIfStmt(node *ast.IfStmt) (ast.VisitorResult, error) {
	cond, err := node.Cond.Visit(i)
	if err != nil {
		return nil, err
	}
	b, ok := cond.(*object.Boolean)
	if !ok {
		return nil, wrongIfCondTypErr(cond)
	}
	if b.Value {
		return node.Then.Visit(i)
	} else if node.Else != nil {
		return node.Else.Visit(i)
	}
	return nil, nil
}

func (i *Interpreter) VisitWhileStmt(node *ast.WhileStmt) (ast.VisitorResult, error) {
	for {
		inv, err := node.Invariant.Visit(i)
		if err != nil {
			return nil, err
		}
		b, ok := inv.(*object.Boolean)
		if !ok {
			return nil, wrongIfCondTypErr(inv)
		}
		if !b.Value {
			break
		}
		if _, err := node.Visit(i); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitRepeatStmt(node *ast.RepeatStmt) (ast.VisitorResult, error) {
	for {
		for _, stmt := range node.Statements {
			if _, err := stmt.Visit(i); err != nil {
				return nil, err
			}
		}
		inv, err := node.Invariant.Visit(i)
		if err != nil {
			return nil, err
		}
		b, ok := inv.(*object.Boolean)
		if !ok {
			return nil, wrongIfCondTypErr(inv)
		}
		if !b.Value {
			break
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitVarDeclStmt(node *ast.VarDeclStmt) (ast.VisitorResult, error) {
	for ident, typ := range node.Declarations {
		t, ok := i.env.LookupType(types.TypeName(typ))
		if !ok {
			return nil, fmt.Errorf("Undeclared type: %s", typ)
		}
		if err := i.env.DeclareVar(object.VarName(ident), t); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func declareBaseTypes(env *object.Environment) {
	env.DeclareType(types.BOOL, types.Bool)
	env.DeclareType(types.INT, types.Int)
	env.DeclareType(types.REAL, types.Real)
	env.DeclareType(types.CHAR, types.Char)
	env.DeclareType(types.STRING, types.String)
}

func declareBuiltins(env *object.Environment) {
	env.DeclareAssignVar(object.VarName("writeln"), builtin.WriteLn)
}
