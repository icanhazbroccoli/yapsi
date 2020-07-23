package types

import "fmt"

type TypeName string

const (
	ARRAY     TypeName = "array"
	BOOL               = "bool"
	BUILTIN            = "builtin"
	CHAR               = "char"
	FUNCTION           = "function"
	INT                = "integer"
	POINTER            = "pointer"
	PROCEDURE          = "procedure"
	REAL               = "real"
	RECORD             = "record"
	STRING             = "string"
)

type Type interface {
	Name() TypeName
	Base() Type
}

type TypeDef struct {
	name TypeName
	base Type
}

var _ Type = (*TypeDef)(nil)

func (t *TypeDef) Name() TypeName { return t.name }
func (t *TypeDef) Base() Type     { return t.base }

var (
	Array     = &TypeDef{name: ARRAY}
	Bool      = &TypeDef{name: BOOL}
	Builtin   = &TypeDef{name: BUILTIN}
	Char      = &TypeDef{name: CHAR}
	Function  = &TypeDef{name: FUNCTION}
	Int       = &TypeDef{name: INT}
	Pointer   = &TypeDef{name: POINTER}
	Procedure = &TypeDef{name: PROCEDURE}
	Real      = &TypeDef{name: REAL}
	Record    = &TypeDef{name: RECORD}
	String    = &TypeDef{name: STRING}
)

type SubrangeType struct {
	left, right Any
}

var _ Type = (*SubrangeType)(nil)

func NewSubrangeType(left, right Any) *SubrangeType {
	return &SubrangeType{left, right}
}

func (t *SubrangeType) Name() TypeName {
	base := t.Base()
	return TypeName(
		fmt.Sprintf("(%s)%v .. (%s)%v",
			base, t.left,
			base, t.right))
}

func (t *SubrangeType) Base() Type {
	return t.left.Type()
}
