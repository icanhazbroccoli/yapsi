package object

import (
	"fmt"

	"yapsi/pkg/types"
)

type Integer struct {
	Value int64
}

var _ types.Arithmetic = (*Integer)(nil)
var _ types.Comparable = (*Integer)(nil)

func (i *Integer) Type() types.Type { return types.Int }

func (i *Integer) String() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) OpUnPlus() (types.Any, error) {
	return i, nil
}

func (i *Integer) OpUnMinus() (types.Any, error) {
	return &Integer{
		Value: -i.Value,
	}, nil
}

func (i *Integer) OpPlus(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value + i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "+")
	}
}

func (i *Integer) OpMinus(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value - i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "-")
	}
}

func (i *Integer) OpAsterisk(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value * i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "*")
	}
}

func (i *Integer) OpSlash(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value / i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "/")
	}
}

func (i *Integer) OpPercent(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value % i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "%")
	}
}

func (i *Integer) OpEql(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value == i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "=")
	}
}

func (i *Integer) OpNeql(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value != i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<>")
	}
}

func (i *Integer) OpGt(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value > i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, ">")
	}
}

func (i *Integer) OpGte(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value >= i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, ">=")
	}
}

func (i *Integer) OpLt(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value < i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<")
	}
}

func (i *Integer) OpLte(o types.Any) (types.Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value <= i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<=")
	}
}
