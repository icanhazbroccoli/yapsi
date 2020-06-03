package ast

type Node interface {
	Visit(NodeVisitor) (VisitorResult, error)
}

type Literal interface {
	Node
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type RawNumber string

type Callable interface{}

type TypeDefinition interface{}

type Identifier string
type VarIdentifier Identifier
type TypeIdentifier Identifier
type LabelIdentifier Identifier
type ConstIdentifier Identifier
type ProgramIdentifier Identifier
type CallableIdentifier Identifier

type Label struct {
	Identifier LabelIdentifier
}

type Value interface{}

type Constant struct {
	Identifier ConstIdentifier
	Raw        Expression
	Type       TypeIdentifier
	Value      Value
}

type Type struct {
	Identifier TypeIdentifier
	Definition TypeDefinition
}

type Variable struct {
	Identifier VarIdentifier
	Type       TypeIdentifier
}

type Program struct {
	Identifier ProgramIdentifier
	Params     []Identifier
	Block
}

type Block struct {
	Labels    []Label
	Constants []Constant
	Types     []Type
	Variables []Variable
	Callables []Callable
	Body      []Statement
}

type Procedure struct {
	Identifier CallableIdentifier
	Variables  []Variable
	Arguments  []Variable
	Body       []Statement
}

type Function struct {
	Identifier CallableIdentifier
	Variables  []Variable
	Arguments  []Variable
	ReturnType TypeIdentifier
	Body       []Statement
}
