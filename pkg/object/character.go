package object

import "yapsi/pkg/types"

type Character struct {
	Value rune
}

var _ (Arithmetic) = (*Character)(nil)
var _ (Comparable) = (*Character)(nil)

func (c *Character) Type() *types.Type { return types.Char }

func (c *Character) OpUnPlus() (Any, error) {
	return c, nil
}

func (c *Character) OpPlus(o Any) (Any, error) {
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

func (c *Character) OpMinus(o Any) (Any, error) {
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

func (c *Character) OpEql(o Any) (*Boolean, error) {
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

func (c *Character) OpNeql(o Any) (*Boolean, error) {
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

func (c *Character) OpGt(o Any) (*Boolean, error) {
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

func (c *Character) OpGte(o Any) (*Boolean, error) {
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

func (c *Character) OpLt(o Any) (*Boolean, error) {
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

func (c *Character) OpLte(o Any) (*Boolean, error) {
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

func (c *Character) OpUnMinus() (Any, error)     { return nil, UnsupTypOpErr(c.Type(), "-") }
func (c *Character) OpAsterisk(Any) (Any, error) { return nil, UnsupTypOpErr(c.Type(), "*") }
func (c *Character) OpSlash(Any) (Any, error)    { return nil, UnsupTypOpErr(c.Type(), "/") }
func (c *Character) OpPercent(Any) (Any, error)  { return nil, UnsupTypOpErr(c.Type(), "%") }
