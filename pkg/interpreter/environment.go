package interpreter

import "yapsi/pkg/object"

type Environment struct {
	types *object.TypeRegistry
	vars  map[string]object.Any
	super *Environment
}

func NewEnvironment(super *Environment) *Environment {
	return &Environment{
		types: object.NewTypeRegistry(),
		vars:  make(map[string]object.Any),
		super: super,
	}
}

func (e *Environment) Lookup(ident string) (object.Any, bool) {
	ptr := e
	for ptr != nil {
		lookup, ok := ptr.vars[ident]
		if ok {
			return lookup, ok
		}
		ptr = ptr.super
	}
	return nil, false
}

func (e *Environment) Assign(ident string, value object.Any) {
	e.vars[ident] = value
}
