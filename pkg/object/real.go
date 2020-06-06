package object

import "yapsi/pkg/types"

type Real struct {
	Value float64
}

var _ Arithmetic = (*Real)(nil)
var _ Comparable = (*Real)(nil)

func (r *Real) Type() *types.Type { return types.Real }

func (r *Real) OpUnPlus() (Any, error) { return r, nil }

func (r *Real) OpUnMinus() (Any, error) {
	return &Real{
		Value: -r.Value,
	}, nil
}

func (r *Real) OpPlus(o Any) (Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value + r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "+")
	}
}

func (r *Real) OpMinus(o Any) (Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value - r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "-")
	}
}

func (r *Real) OpAsterisk(o Any) (Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value * r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "*")
	}
}

func (r *Real) OpSlash(o Any) (Any, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Real{
			Value: r.Value / r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "/")
	}
}

func (r *Real) OpEql(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value == r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "=")
	}
}

func (r *Real) OpNeql(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value != r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<>")
	}
}

func (r *Real) OpGt(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value > r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, ">")
	}
}

func (r *Real) OpGte(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value >= r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, ">=")
	}
}

func (r *Real) OpLt(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value < r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<")
	}
}

func (r *Real) OpLte(o Any) (*Boolean, error) {
	switch r2 := o.(type) {
	case *Real:
		return &Boolean{
			Value: r.Value <= r2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(r, o, "<=")
	}
}

func (r *Real) OpPercent(o Any) (Any, error) { return nil, UnsupTypOpErr(r.Type(), "%") }
