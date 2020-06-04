package object

type TypeName string

type TypeDefinition interface{}

const (
	BOOL   TypeName = "bool"
	INT             = "int"
	REAL            = "real"
	STRING          = "string"
	CHAR            = "char"
	ARRAY           = "array"
	RECORD          = "record"
)

type Type struct {
	name TypeName
	def  TypeDefinition
	base *Type
}

func NewType(name TypeName, def TypeDefinition, base *Type) *Type {
	return &Type{
		name: name,
		def:  def,
		base: base,
	}
}

func (t *Type) Name() TypeName             { return t.name }
func (t *Type) Definition() TypeDefinition { return t.def }
func (t *Type) BaseType() *Type            { return t.base }
