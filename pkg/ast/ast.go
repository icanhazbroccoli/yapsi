package ast

type Callable interface{}

type TypeDefinition interface{}

type VarIdentifier string
type TypeIdentifier string
type LabelIdentifier string
type ConstIdentifier string
type ProgramIdentifier string
type CallableIdentifier string

type Label struct {
	Identifier LabelIdentifier
}

type Value interface{}

type Constant struct {
	Identifier ConstIdentifier
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

//======

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
