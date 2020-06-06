package interpreter

import (
	"fmt"
	"yapsi/pkg/object"
	"yapsi/pkg/types"
)

type Environment struct {
	types *types.Registry
	vars  map[object.VarName]*object.Variable
	super *Environment
}

func NewEnvironment(super *Environment) *Environment {
	return &Environment{
		types: types.NewRegistry(),
		vars:  make(map[object.VarName]*object.Variable),
		super: super,
	}
}

func (e *Environment) LookupType(name types.TypeName) (*types.Type, bool) {
	ptr := e
	for ptr != nil {
		if typ, ok := e.types.Lookup(name); ok {
			return typ, true
		}
		ptr = ptr.super
	}
	return nil, false
}

func (e *Environment) DeclareType(name types.TypeName, t *types.Type) error {
	if err := e.types.Register(name, t); err != nil {
		return err
	}
	return nil
}

func (e *Environment) LookupVar(name object.VarName) (object.Any, bool) {
	ptr := e
	for ptr != nil {
		if lookup, ok := ptr.vars[name]; ok {
			return lookup, ok
		}
		ptr = ptr.super
	}
	return nil, false
}

func (e *Environment) DeclareVar(name object.VarName, t *types.Type) error {
	if _, ok := e.vars[name]; ok {
		return fmt.Errorf("Duplicate variable declaration: %s", name)
	}
	e.vars[name] = object.NewVariable(name, t, nil)
	return nil
}

func (e *Environment) AssignVar(name object.VarName, value object.Any) error {
	if _, ok := e.vars[name]; !ok {
		return fmt.Errorf("Undefined variable: %s", name)
	}
	if err := e.vars[name].SetValue(value); err != nil {
		return err
	}
	return nil
}
