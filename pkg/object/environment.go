package object

import (
	"fmt"
	"yapsi/pkg/types"
)

type Environment struct {
	types *types.Registry
	vars  map[VarName]*Variable
	super *Environment
}

func NewEnvironment(super *Environment) *Environment {
	return &Environment{
		types: types.NewRegistry(),
		vars:  make(map[VarName]*Variable),
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

func (e *Environment) LookupVar(name VarName) (Any, bool) {
	ptr := e
	for ptr != nil {
		if lookup, ok := ptr.vars[name]; ok {
			return lookup.Any, ok
		}
		ptr = ptr.super
	}
	return nil, false
}

func (e *Environment) DeclareVar(name VarName, t *types.Type) error {
	if _, ok := e.vars[name]; ok {
		return fmt.Errorf("Duplicate variable declaration: %s", name)
	}
	e.vars[name] = NewVariable(name, t, nil)
	return nil
}

func (e *Environment) AssignVar(name VarName, value Any) error {
	if _, ok := e.vars[name]; !ok {
		return fmt.Errorf("Undefined variable: %s", name)
	}
	if err := e.vars[name].SetValue(value); err != nil {
		return err
	}
	return nil
}

func (e *Environment) DeclareAssignVar(name VarName, value Any) error {
	if err := e.DeclareVar(name, value.Type()); err != nil {
		return err
	}
	return e.AssignVar(name, value)
}
