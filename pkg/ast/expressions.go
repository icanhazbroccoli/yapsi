package ast

import "yapsi/pkg/token"

type IdentifierExpr struct {
	Token token.Token
	Value string
}

var _ Expression = (*IdentifierExpr)(nil)

func (e *IdentifierExpr) expressionNode() {}
func (e *IdentifierExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitIdentifierExpr(e)
}

type UnaryExpr struct {
	Operator token.Token
	Expr     Expression
}

var _ Expression = (*UnaryExpr)(nil)

func (e *UnaryExpr) expressionNode() {}
func (e *UnaryExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitUnaryExpr(e)
}

type BinaryExpr struct {
	Left, Right Expression
	Operator    token.Token
}

var _ Expression = (*BinaryExpr)(nil)

func (e *BinaryExpr) expressionNode() {}
func (e *BinaryExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitBinaryExpr(e)
}

type ElementExpr struct {
	Left, Right Expression
}

var _ Expression = (*ElementExpr)(nil)

func (e *ElementExpr) expressionNode() {}
func (e *ElementExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitElementExpr(e)
}

type SetExpr struct {
	Elements []Expression
}

var _ Expression = (*SetExpr)(nil)

func (e *SetExpr) expressionNode() {}
func (e *SetExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitSetExpr(e)
}

type FunctionCallExpr struct {
	Token      token.Token
	Identifier *IdentifierExpr
	Args       []Expression
}

var _ Expression = (*FunctionCallExpr)(nil)

func (e *FunctionCallExpr) expressionNode() {}
func (e *FunctionCallExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitFunctionCallExpr(e)
}

type TypeDefinitionExprIntf interface{}

type SimpleTypeDefinitionExpr struct {
	Token      token.Token
	Identifier Expression
}

var _ Expression = (*SimpleTypeDefinitionExpr)(nil)
var _ TypeDefinitionExprIntf = (*SimpleTypeDefinitionExpr)(nil)

func (e *SimpleTypeDefinitionExpr) expressionNode() {}
func (e *SimpleTypeDefinitionExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitSimpleTypeDefinitionExpr(e)
}

type SubrangeTypeDefinitionExpr struct {
	Token token.Token
	Left  Expression
	Right Expression
}

var _ Expression = (*SubrangeTypeDefinitionExpr)(nil)
var _ TypeDefinitionExprIntf = (*SubrangeTypeDefinitionExpr)(nil)

func (e *SubrangeTypeDefinitionExpr) expressionNode() {}
func (e *SubrangeTypeDefinitionExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitSubrangeTypeDefinitionExpr(e)
}

type ArrayTypeDefinitionExpr struct {
	Token            token.Token
	Packed           bool
	IndexTypeDef     TypeDefinitionExprIntf
	ComponentTypeDef TypeDefinitionExprIntf
}

var _ Expression = (*ArrayTypeDefinitionExpr)(nil)
var _ TypeDefinitionExprIntf = (*ArrayTypeDefinitionExpr)(nil)

func (e *ArrayTypeDefinitionExpr) expressionNode() {}
func (e *ArrayTypeDefinitionExpr) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitArrayTypeDefinitionExpr(e)
}
