package object

import "yapsi/pkg/ast"

type Value interface {
	Type() Type
	Comparable() bool
	Value() interface{}
}

type IntValue struct {
	value int64
}

var _ Value = (*IntValue)(nil)

func NewIntValue(node *ast.NumericLiteral) (*IntValue, error) {
	value, err := parseInt(node.Value)
	if err != nil {
		return nil, err
	}
	return &IntValue{
		value: value,
	}, nil
}

func (v *IntValue) Type() Type         { return INT }
func (v *IntValue) Comparable() bool   { return true }
func (v *IntValue) Value() interface{} { return v.value }
