package object

import (
	"fmt"
	"yapsi/pkg/types"
)

type String struct {
	Value string
}

var _ Arithmetic = (*String)(nil)
var _ Comparable = (*String)(nil)
var _ Indexable = (*String)(nil)

func (s *String) Type() *types.Type { return types.String }
func (s *String) String() string    { return s.Value }

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

func (s *String) Len() int {
	return len(s.Value)
}

func (s *String) OpSubscrGet(arg Arithmetic) (Any, error) {
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

func (s *String) OpUnPlus() (Any, error)      { return nil, UnsupTypOpErr(s.Type(), "+") }
func (s *String) OpUnMinus() (Any, error)     { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpMinus(Any) (Any, error)    { return nil, UnsupTypOpErr(s.Type(), "-") }
func (s *String) OpAsterisk(Any) (Any, error) { return nil, UnsupTypOpErr(s.Type(), "*") }
func (s *String) OpSlash(Any) (Any, error)    { return nil, UnsupTypOpErr(s.Type(), "/") }
func (s *String) OpPercent(Any) (Any, error)  { return nil, UnsupTypOpErr(s.Type(), "%") }
func (s *String) OpSubscrSet(Arithmetic, Any) error {
	return UnsupTypOpErr(s.Type(), "[]=")
}
