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
	if lookup == nil {
		return nil, uninitializedIdentErr(node.Value)
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
		if lg, ok := expr.(types.Logical); ok {
			return lg.OpNot()
		}
	case "-":
		if ar, ok := expr.(types.Arithmetic); ok {
			return ar.OpUnMinus()
		}
	case "+":
		if ar, ok := expr.(types.Arithmetic); ok {
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
	left, right := lr.(types.Any), rr.(types.Any)
	switch node.Operator.Literal {
	case "+":
		if ar, ok := left.(types.Arithmetic); ok {
			return ar.OpPlus(right.(types.Any))
		}
	case "-":
		if ar, ok := left.(types.Arithmetic); ok {
			return ar.OpMinus(right.(types.Any))
		}
	case "*":
		if ar, ok := left.(types.Arithmetic); ok {
			return ar.OpAsterisk(right.(types.Any))
		}
	case "/":
		if ar, ok := left.(types.Arithmetic); ok {
			return ar.OpSlash(right.(types.Any))
		}
	case "%":
		if ar, ok := left.(types.Arithmetic); ok {
			return ar.OpPercent(right.(types.Any))
		}
	case "=":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpEql(right.(types.Any))
		}
	case "<>":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpNeql(right.(types.Any))
		}
	case ">":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpGt(right.(types.Any))
		}
	case ">=":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpGte(right.(types.Any))
		}
	case "<":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpLt(right.(types.Any))
		}
	case "<=":
		if cp, ok := left.(types.Comparable); ok {
			return cp.OpLte(right.(types.Any))
		}
	case "and":
		if lg, ok := left.(types.Logical); ok {
			return lg.OpAnd(right.(types.Any))
		}
	case "or":
		if lg, ok := left.(types.Logical); ok {
			return lg.OpOr(right.(types.Any))
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
	if node.VarDecl != nil {
		if _, err := node.VarDecl.Visit(i); err != nil {
			return nil, err
		}
	}
	for _, procedure := range node.Procedures {
		if _, err := procedure.Visit(i); err != nil {
			return nil, err
		}
	}
	for _, function := range node.Functions {
		if _, err := function.Visit(i); err != nil {
			return nil, err
		}
	}
	return node.Statement.Visit(i)
}

func (i *Interpreter) VisitCompoundStmt(node *ast.CompoundStmt) (ast.VisitorResult, error) {
	for _, stmt := range node.Statements {
		res, err := stmt.Visit(i)
		if err != nil {
			return nil, err
		}
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return res, nil
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
	value, ok := r.(types.Any)
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

func (i *Interpreter) evaluateArgs(in []ast.Expression) ([]types.Any, error) {
	out := make([]types.Any, 0, len(in))
	for _, arg := range in {
		value, err := arg.Visit(i)
		if err != nil {
			return nil, err
		}
		out = append(out, value.(types.Any))
	}
	return out, nil
}

func (i *Interpreter) newEnvWithParams(formal []*object.Variable, actual []types.Any) (*object.Environment, error) {
	env := object.NewEnvironment(i.env)
	if len(formal) != len(actual) {
		return nil, wrongCallableArgLenErr(len(formal), len(actual))
	}
	for ix := 0; ix < len(formal); ix++ {
		if err := env.DeclareVar(formal[ix].Name(), formal[ix].Type()); err != nil {
			return nil, err
		}
		if err := env.AssignVar(formal[ix].Name(), actual[ix]); err != nil {
			return nil, err
		}
	}
	return env, nil
}

func (i *Interpreter) VisitProcedureCallStmt(node *ast.ProcedureCallStmt) (ast.VisitorResult, error) {
	ident := object.VarName(node.Identifier.Value)
	p, ok := i.env.LookupVar(ident)
	if !ok {
		return nil, undefinedIdentErr(string(ident))
	}
	args, err := i.evaluateArgs(node.Args)
	if err != nil {
		return nil, err
	}
	if builtin, ok := p.(*object.Builtin); ok {
		return builtin.CallReturn(i.env, args...)
	}
	procedure, ok := p.(*object.Procedure)
	if !ok {
		return nil, uncallableEntityErr(string(ident))
	}
	oldEnv := i.env
	newEnv, err := i.newEnvWithParams(procedure.Params, args)
	if err != nil {
		return nil, err
	}

	i.env = newEnv

	if _, err := procedure.Body.Visit(i); err != nil {
		return nil, err
	}

	for ix, arg := range node.Args {
		extIdent, ok := arg.(*ast.IdentifierExpr)
		if !ok {
			continue
		}
		fVar := procedure.Params[ix]
		fVal, _ := newEnv.LookupVar(fVar.Name())
		if !ok {
			// Should never happen but I want to keep it here in case variable
			// removal would be supported in the future.
			continue
		}
		oldEnv.AssignVar(object.VarName(extIdent.Value), fVal)
	}

	i.env = oldEnv

	return nil, nil
}

func (i *Interpreter) VisitFunctionCallExpr(node *ast.FunctionCallExpr) (ast.VisitorResult, error) {
	ident := object.VarName(node.Token.Literal)
	f, ok := i.env.LookupVar(ident)
	if !ok {
		return nil, undefinedIdentErr(string(ident))
	}
	args, err := i.evaluateArgs(node.Args)
	if err != nil {
		return nil, err
	}
	// TODO: refactor me
	if builtin, ok := f.(*object.Builtin); ok {
		return builtin.CallReturn(i.env, args...)
	}
	function, ok := f.(*object.Function)
	if !ok {
		return nil, uncallableEntityErr(string(ident))
	}
	oldEnv := i.env
	newEnv, err := i.newEnvWithParams(function.Params, args)
	if err != nil {
		return nil, callableErr(string(ident), err)
	}

	ret := object.NewVariable(object.VarName("result"), function.ReturnType, nil)
	if err := newEnv.DeclareVar(ret.Name(), ret.Type()); err != nil {
		return nil, err
	}

	i.env = newEnv

	if _, err := function.Body.Visit(i); err != nil {
		return nil, err
	}

	i.env = oldEnv

	res, _ := newEnv.LookupVar(object.VarName("result"))

	return res, nil
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
		if _, err := node.Body.Visit(i); err != nil {
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
		t, ok := i.env.LookupType(types.TypeName(ident))
		if !ok {
			return nil, fmt.Errorf("Undeclared type: %s", typ)
		}
		if err := i.env.DeclareVar(object.VarName(ident), t); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (i *Interpreter) VisitFunctionDeclStmt(node *ast.FunctionDeclStmt) (ast.VisitorResult, error) {
	ident := node.Identifier.Value
	retTypeIdent := node.ReturnType.Value
	retType, ok := i.env.LookupType(types.TypeName(retTypeIdent))
	if !ok {
		return nil, fmt.Errorf("Undeclared type: %s", retTypeIdent)
	}
	params := []*object.Variable{}
	for _, arg := range node.Args {
		ident := arg.Identifer.Value
		typIdent := arg.Type.Value
		typ, ok := i.env.LookupType(types.TypeName(typIdent))
		if !ok {
			return nil, fmt.Errorf("Undeclared type: %s", typIdent)
		}
		params = append(params, object.NewVariable(object.VarName(ident), typ, nil))
	}
	function := &object.Function{
		Identifier: ident,
		ReturnType: retType,
		Params:     params,
		Body:       node.Body,
	}
	if err := i.env.DeclareAssignVar(object.VarName(ident), function); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *Interpreter) VisitReturnStmt(node *ast.ReturnStmt) (ast.VisitorResult, error) {
	res, err := node.Expression.Visit(i)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *Interpreter) VisitProcedureDeclStmt(node *ast.ProcedureDeclStmt) (ast.VisitorResult, error) {
	ident := node.Identifier.Value
	params := []*object.Variable{}
	for _, arg := range node.Args {
		ident := arg.Identifer.Value
		typIdent := arg.Type.Value
		typ, ok := i.env.LookupType(types.TypeName(typIdent))
		if !ok {
			return nil, fmt.Errorf("Undeclared type: %s", typIdent)
		}
		params = append(params, object.NewVariable(object.VarName(ident), typ, nil))
	}
	procedure := &object.Procedure{
		Identifier: ident,
		Params:     params,
		Body:       node.Body,
	}
	if err := i.env.DeclareAssignVar(object.VarName(ident), procedure); err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *Interpreter) VisitSimpleTypeDefinitionExpr(node *ast.SimpleTypeDefinitionExpr) (ast.VisitorResult, error) {
	ident := node.Identifier
	typ, ok := i.env.LookupType(types.TypeName(ident.Value))
	if !ok {
		return nil, undefinedRefTypeErr(ident.Value)
	}
	return typ, nil
}

func (i *Interpreter) VisitSubrangeTypeDefinitionExpr(node *ast.SubrangeTypeDefinitionExpr) (ast.VisitorResult, error) {
	lr, err := node.Left.Visit(i)
	if err != nil {
		return nil, err
	}
	rr, err := node.Right.Visit(i)
	if err != nil {
		return nil, err
	}
	left := lr.(types.Any)
	right := rr.(types.Any)
	if err := assertTypeEqual(left.Type(), right.Type()); err != nil {
		return nil, err
	}
	return types.NewSubrangeType(left, right), nil
}

func (i *Interpreter) VisitArrayTypeDefinitionExpr(*ast.ArrayTypeDefinitionExpr) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitRecordTypeDefinitionExpr(*ast.RecordTypeDefinitionExpr) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitEnumTypeDefinitionExpr(*ast.EnumTypeDefinitionExpr) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitSetTypeDefinitionExpr(*ast.SetTypeDefinitionExpr) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitFileTypeDefinitionExpr(*ast.FileTypeDefinitionExpr) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitTypeDeclStmt(*ast.TypeDeclStmt) (ast.VisitorResult, error) {
	panic("not implemented")
}

func (i *Interpreter) VisitTypeDefinitionStmt(*ast.TypeDefinitionStmt) (ast.VisitorResult, error) {
	panic("not implemented")
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
	env.DeclareAssignVar(object.VarName("len"), builtin.Len)
}
