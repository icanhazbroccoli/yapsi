package object

import "yapsi/pkg/types"

type String struct {
	Value string
}

var _ Any = (*String)(nil)
var _ Arithmetic = (*String)(nil)
var _ Comparable = (*String)(nil)

func (s *String) Type() *types.Type { return types.String }

func (s *String) OpPlus(o Any) (Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &String{
			Value: s.Value + s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "+")
	}
}

func (s *String) OpEql(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value == s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "=")
	}
}

func (s *String) OpNeql(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value != s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<>")
	}
}

func (s *String) OpGt(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value > s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, ">")
	}
}

func (s *String) OpGte(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value >= s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, ">=")
	}
}

func (s *String) OpLt(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value < s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<")
	}
}

func (s *String) OpLte(o Any) (*Boolean, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value <= s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<=")
	}
}

func (s *String) OpUnPlus() (Any, error)      { return nil, UnsupTypOpErr(s.Type(), "+") }
func (s *String) OpUnMinus() (Any, error)     { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpMinus(Any) (Any, error)    { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpAsterisk(Any) (Any, error) { return nil, UnsupTypOpErr(s.Type(), "*") }
func (s *String) OpSlash(Any) (Any, error)    { return nil, UnsupTypOpErr(s.Type(), "/") }
func (s *String) OpPercent(Any) (Any, error)  { return nil, UnsupTypOpErr(s.Type(), "%") }
