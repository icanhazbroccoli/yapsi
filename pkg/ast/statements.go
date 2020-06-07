package ast

import "yapsi/pkg/token"

// <assignment statement> ::= <variable> := <expression> | <function identifier> := <expression>
type AssignmentStmt struct {
	Identifier *IdentifierExpr
	Expr       Expression
}

var _ (Statement) = (*AssignmentStmt)(nil)

func (s *AssignmentStmt) statementNode() {}
func (s *AssignmentStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitAssignmentStmt(s)
}

type GotoStmt struct {
	Label IdentifierExpr
}

var _ (Statement) = (*GotoStmt)(nil)

func (s *GotoStmt) statementNode() {}
func (s *GotoStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitGotoStmt(s)
}

type LabeledStmt struct {
	Label *IdentifierExpr
	Stmt  Statement
}

var _ (Statement) = (*LabeledStmt)(nil)

func (s *LabeledStmt) statementNode() {}
func (s *LabeledStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitLabeledStmt(s)
}

type ProcedureStmt struct {
	Identifier *IdentifierExpr
	Args       []Expression
}

var _ (Statement) = (*ProcedureStmt)(nil)

func (s *ProcedureStmt) statementNode() {}
func (s *ProcedureStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitProcedureStmt(s)
}

type CompoundStmt struct {
	Token      token.Token
	Statements []Statement
}

var _ (Statement) = (*CompoundStmt)(nil)

func (s *CompoundStmt) statementNode() {}
func (s *CompoundStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitCompoundStmt(s)
}

type IfStmt struct {
	Token token.Token
	Cond  Expression
	Then  Statement
	Else  Statement
}

var _ (Statement) = (*IfStmt)(nil)

func (s *IfStmt) statementNode() {}
func (s *IfStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitIfStmt(s)
}

type WhileStmt struct {
	Token     token.Token
	Invariant Expression
	Body      Statement
}

var _ Statement = (*WhileStmt)(nil)

func (s *WhileStmt) statementNode() {}
func (s *WhileStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitWhileStmt(s)
}

type RepeatStmt struct {
	Token      token.Token
	Invariant  Expression
	Statements []Statement
}

var _ Statement = (*RepeatStmt)(nil)

func (s *RepeatStmt) statementNode() {}
func (s *RepeatStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitRepeatStmt(s)
}

type ProgramStmt struct {
	Token      token.Token
	Identifier *IdentifierExpr
	Block      *BlockStmt
}

var _ Statement = (*ProgramStmt)(nil)

func (s *ProgramStmt) statementNode() {}
func (s *ProgramStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitProgramStmt(s)
}

type BlockStmt struct {
	Token     token.Token
	VarDecl   *VarDeclStmt
	Statement *CompoundStmt
}

var _ Statement = (*BlockStmt)(nil)

func (s *BlockStmt) statementNode() {}
func (s *BlockStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitBlockStmt(s)
}

type VarDeclStmt struct {
	Token        token.Token
	Declarations map[string]string
}

func (s *VarDeclStmt) statementNode() {}
func (s *VarDeclStmt) Visit(v NodeVisitor) (VisitorResult, error) {
	return v.VisitVarDeclStmt(s)
}
