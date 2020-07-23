package object

import (
	"fmt"
	"yapsi/pkg/types"
)

type String struct {
	Value string
}

var _ types.Arithmetic = (*String)(nil)
var _ types.Comparable = (*String)(nil)
var _ types.Indexable = (*String)(nil)

func (s *String) Type() types.Type { return types.String }
func (s *String) String() string   { return s.Value }

func (s *String) OpPlus(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &String{
			Value: s.Value + s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "+")
	}
}

func (s *String) OpEql(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value == s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "=")
	}
}

func (s *String) OpNeql(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value != s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<>")
	}
}

func (s *String) OpGt(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value > s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, ">")
	}
}

func (s *String) OpGte(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value >= s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, ">=")
	}
}

func (s *String) OpLt(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value < s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<")
	}
}

func (s *String) OpLte(o types.Any) (types.Any, error) {
	switch s2 := o.(type) {
	case *String:
		return &Boolean{
			s.Value <= s2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(s, o, "<=")
	}
}

func (s *String) Len() int {
	return len(s.Value)
}

func (s *String) OpSubscrGet(arg types.Arithmetic) (types.Any, error) {
	ix, ok := arg.(*Integer)
	if !ok {
		return nil, fmt.Errorf("string is int-indexable, %s is provided",
			arg.Type().Name())
	}
	if ix.Value < 0 || int(ix.Value) >= len(s.Value) {
		return nil, fmt.Errorf("the index is out of bounds")
	}
	r := []rune(s.Value)[ix.Value]
	return &Character{Value: r}, nil
}

func (s *String) OpUnPlus() (types.Any, error)            { return nil, UnsupTypOpErr(s.Type(), "+") }
func (s *String) OpUnMinus() (types.Any, error)           { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpMinus(types.Any) (types.Any, error)    { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpAsterisk(types.Any) (types.Any, error) { return nil, UnsupTypOpErr(s.Type(), "*") }
func (s *String) OpSlash(types.Any) (types.Any, error)    { return nil, UnsupTypOpErr(s.Type(), "/") }
func (s *String) OpPercent(types.Any) (types.Any, error)  { return nil, UnsupTypOpErr(s.Type(), "%") }
func (s *String) OpSubscrSet(types.Arithmetic, types.Any) error {
	return UnsupTypOpErr(s.Type(), "[]=")
}
