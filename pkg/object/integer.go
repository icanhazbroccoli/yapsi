package object

type Integer struct {
	Value int64
}

var _ Any = (*Integer)(nil)
var _ Arithmetic = (*Integer)(nil)
var _ Comparable = (*Integer)(nil)

func (i *Integer) Type() Type { return INT }

func (i *Integer) OpUnPlus() (Any, error) {
	return i, nil
}

func (i *Integer) OpUnMinus() (Any, error) {
	return &Integer{
		Value: -i.Value,
	}, nil
}

func (i *Integer) OpPlus(o Any) (Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value + i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "+")
	}
}

func (i *Integer) OpMinus(o Any) (Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value - i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "-")
	}
}

func (i *Integer) OpAsterisk(o Any) (Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value * i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "*")
	}
}

func (i *Integer) OpSlash(o Any) (Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value / i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "/")
	}
}

func (i *Integer) OpPercent(o Any) (Any, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Integer{
			Value: i.Value % i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "%")
	}
}

func (i *Integer) OpEql(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value == i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "=")
	}
}

func (i *Integer) OpNeql(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value != i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<>")
	}
}

func (i *Integer) OpGt(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value > i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, ">")
	}
}

func (i *Integer) OpGte(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value >= i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, ">=")
	}
}

func (i *Integer) OpLt(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value < i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<")
	}
}

func (i *Integer) OpLte(o Any) (*Boolean, error) {
	switch i2 := o.(type) {
	case *Integer:
		return &Boolean{
			Value: i.Value <= i2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(i, o, "<=")
	}
}
