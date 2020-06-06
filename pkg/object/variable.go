package object

import (
	"fmt"
	"yapsi/pkg/types"
)

type VarName string

type Variable struct {
	Any
	name VarName
	typ  *types.Type
}

func NewVariable(name VarName, typ *types.Type, val Any) *Variable {
	return &Variable{val, name, typ}
}

func (v *Variable) Name() VarName { return v.name }
func (v *Variable) Value() Any    { return v.Any }
func (v *Variable) SetValue(val Any) error {
	if v.typ != val.Type() {
		return fmt.Errorf("Type mismatch for variable `%s`: want: %s, got: %s",
			v.name, v.typ.Name(), val.Type().Name())
	}
	v.Any = val
	return nil
}
