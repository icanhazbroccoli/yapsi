package types

import "fmt"

type Registry struct {
	types map[TypeName]Type
}

func NewRegistry() *Registry {
	return &Registry{
		types: make(map[TypeName]Type),
	}
}

func (tr *Registry) Register(name TypeName, t Type) error {
	if _, ok := tr.types[name]; ok {
		return duplicateTypeDefErr(name)
	}
	tr.types[name] = t
	return nil
}

func (tr *Registry) Lookup(name TypeName) (Type, bool) {
	t, ok := tr.types[name]
	return t, ok
}

func duplicateTypeDefErr(name TypeName) error {
	return fmt.Errorf("Type named %q has already been defined", name)
}
