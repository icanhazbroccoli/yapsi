package object

import (
	"fmt"

	"yapsi/pkg/types"
)

type Real struct {
	Value float64
}

var _ types.Arithmetic = (*Real)(nil)
var _ types.Comparable = (*Real)(nil)

func (r *Real) Type() types.Type { return types.Real }
func (r *Real) String() string   { return fmt.Sprintf("%f", r.Value) }

func (r *Real) OpUnPlus() (types.Any, error) { return r, nil }

func (r *Real) OpUnMinus() (types.Any, error) {
	return &Real{
		Value: -r.Value,
	}, nil
}

func (r *Real) OpPlus(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value + r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "+")
	}
}

func (r *Real) OpMinus(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value - r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "-")
	}
}

func (r *Real) OpAsterisk(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value * r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "*")
	}
}

func (r *Real) OpSlash(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value / r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "/")
	}
}

func (r *Real) OpEql(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value == r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "=")
	}
}

func (r *Real) OpNeql(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value != r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<>")
	}
}

func (r *Real) OpGt(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value > r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, ">")
	}
}

func (r *Real) OpGte(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value >= r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, ">=")
	}
}

func (r *Real) OpLt(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value < r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<")
	}
}

func (r *Real) OpLte(o types.Any) (types.Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value <= r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<=")
	}
}

func (r *Real) OpPercent(o types.Any) (types.Any, error) {
	return nil, UnsupTypOpErr(r.Type(), "%")
}
