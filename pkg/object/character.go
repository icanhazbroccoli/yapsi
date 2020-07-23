package object

import "yapsi/pkg/types"

type Character struct {
	Value rune
}

var _ (types.Arithmetic) = (*Character)(nil)
var _ (types.Comparable) = (*Character)(nil)

func (c *Character) Type() types.Type { return types.Char }

func (c *Character) String() string { return string([]rune{c.Value}) }

func (c *Character) OpUnPlus() (types.Any, error) {
	return c, nil
}

func (c *Character) OpPlus(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Character{
			Value: c.Value + c2.Value,
		}, nil
	case *Integer:
		return &Character{
			Value: c.Value + rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "+")
	}
}

func (c *Character) OpMinus(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Character{
			Value: c.Value - rune(c2.Value),
		}, nil
	case *Integer:
		return &Character{
			Value: c.Value - rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "-")
	}
}

func (c *Character) OpEql(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value == c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value == rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "=")
	}
}

func (c *Character) OpNeql(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value != c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value != rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "<>")
	}
}

func (c *Character) OpGt(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value > c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value > rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, ">")
	}
}

func (c *Character) OpGte(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value >= c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value >= rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, ">=")
	}
}

func (c *Character) OpLt(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value < c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value < rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "<")
	}
}

func (c *Character) OpLte(o types.Any) (types.Any, error) {
	switch c2 := o.(type) {
	case *Character:
		return &Boolean{
			Value: c.Value <= c2.Value,
		}, nil
	case *Integer:
		return &Boolean{
			Value: c.Value <= rune(c2.Value),
		}, nil
	default:
		return nil, WrongOperandTypErr(c, o, "<=")
	}
}

func (c *Character) OpUnMinus() (types.Any, error) { return nil, UnsupTypOpErr(c.Type(), "-") }
func (c *Character) OpAsterisk(types.Any) (types.Any, error) {
	return nil, UnsupTypOpErr(c.Type(), "*")
}
func (c *Character) OpSlash(types.Any) (types.Any, error)   { return nil, UnsupTypOpErr(c.Type(), "/") }
func (c *Character) OpPercent(types.Any) (types.Any, error) { return nil, UnsupTypOpErr(c.Type(), "%") }
