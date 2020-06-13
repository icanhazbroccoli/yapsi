package types

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

type Type struct {
	name TypeName
}

func (t *Type) Name() TypeName { return t.name }

var (
	Array     = &Type{name: ARRAY}
	Bool      = &Type{name: BOOL}
	Builtin   = &Type{name: BUILTIN}
	Char      = &Type{name: CHAR}
	Function  = &Type{name: FUNCTION}
	Int       = &Type{name: INT}
	Pointer   = &Type{name: POINTER}
	Procedure = &Type{name: PROCEDURE}
	Real      = &Type{name: REAL}
	Record    = &Type{name: RECORD}
	String    = &Type{name: STRING}
)
