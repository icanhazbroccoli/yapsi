package object

import (
	"fmt"

	"yapsi/pkg/types"
)

type Boolean struct {
	Value bool
}

var _ types.Logical = (*Boolean)(nil)

func (b *Boolean) Type() types.Type { return types.Bool }

func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) OpNot() (types.Any, error) {
	return &Boolean{
		Value: !b.Value,
	}, nil
}

func (b *Boolean) OpAnd(o types.Any) (types.Any, error) {
	switch b2 := o.(type) {
	case *Boolean:
		return &Boolean{
			Value: b.Value && b2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(b, o, "and")
	}
}

func (b *Boolean) OpOr(o types.Any) (types.Any, error) {
	switch b2 := o.(type) {
	case *Boolean:
		return &Boolean{
			Value: b.Value || b2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(b, o, "or")
	}
}
