package interpreter

import "yapsi/pkg/object"

type Environment struct {
	mappings map[string]object.Any
	super    *Environment
}

func NewEnvironment(super *Environment) *Environment {
	return &Environment{
		mappings: make(map[string]object.Any),
		super:    super,
	}
}

func (e *Environment) Lookup(ident string) (object.Any, bool) {
	ptr := e
	for ptr != nil {
		lookup, ok := ptr.mappings[ident]
		if ok {
			return lookup, ok
		}
		ptr = ptr.super
	}
	return nil, false
}

func (e *Environment) Assign(ident string, value object.Any) {
	e.mappings[ident] = value
}
