package object

type TypeRegistry struct {
	types map[TypeName]*Type
}

func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types: make(map[TypeName]*Type),
	}
}

func (tr *TypeRegistry) Register(t *Type) error {
	if _, ok := tr.types[t.Name()]; ok {
		return duplicateTypeDefErr(t)
	}
	tr.types[t.Name()] = t
	return nil
}

func (tr *TypeRegistry) Define(name TypeName, def TypeDefinition, base *Type) error {
	t := NewType(name, def, base)
	return tr.Register(t)
}

func (tr *TypeRegistry) Lookup(name TypeName) (*Type, bool) {
	t, ok := tr.types[name]
	return t, ok
}
