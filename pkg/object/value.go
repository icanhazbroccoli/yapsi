package object

import "yapsi/pkg/ast"

type Value interface {
	Type() Type
	Comparable() bool
}

type BoolValue struct {
	value bool
}

var _ Value = (*BoolValue)(nil)

func NewBoolValue(node *ast.BoolLiteral) (*BoolValue, error) {
	return &BoolValue{
		value: node.Value,
	}, nil
}

func (v *BoolValue) Type() Type       { return BOOL }
func (v *BoolValue) Comparable() bool { return false }
func (v *BoolValue) Value() bool      { return v.value }

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

func (v *IntValue) Type() Type       { return INT }
func (v *IntValue) Comparable() bool { return true }
func (v *IntValue) Value() int64     { return v.value }

type RealValue struct {
	value float64
}

var _ Value = (*RealValue)(nil)

func NewRealValue(node *ast.NumericLiteral) (*RealValue, error) {
	value, err := parseReal(node.Value)
	if err != nil {
		return nil, err
	}
	return &RealValue{
		value: value,
	}, nil
}

func (v *RealValue) Type() Type       { return REAL }
func (v *RealValue) Comparable() bool { return true }
func (v *RealValue) Value() float64   { return v.value }
