package object

import (
	"strings"

	"yapsi/pkg/types"
)

type Array struct {
	left, right Arithmetic
	it, ct      *types.Type
	values      []Any
}

var _ Indexable = (*Array)(nil)

func (a *Array) Type() *types.Type { return types.Array }
func (a *Array) Len() int          { return len(a.values) }

func (a *Array) String() string {
	chunks := make([]string, 0, len(a.values))
	for _, v := range a.values {
		chunks = append(chunks, v.String())
	}
	return "[" + strings.Join(chunks, ", ") + "]"
}

func NewArray(l, r Arithmetic, it, ct *types.Type) (*Array, error) {
	a := &Array{
		left:  l,
		right: r,
		it:    it,
		ct:    ct,
	}
	i0, err := a.convIdx(l)
	if err != nil {
		return nil, err
	}
	i1, err := a.convIdx(r)
	if err != nil {
		return nil, err
	}
	length := i1 - i0
	if length < 0 {
		return nil, arrWrongRangeConstraintErr(l, r)
	}
	a.values = make([]Any, length)
	return a, nil
}

func (a *Array) OpSubscrGet(ix Arithmetic) (Any, error) {
	i, err := a.convIdx(ix)
	if err != nil {
		return nil, err
	}
	if i < 0 || i >= len(a.values) {
		return nil, arrIndexOutOfBounsErr(ix)
	}
	return a.values[i], nil
}

func (a *Array) OpSubscrSet(ix Arithmetic, v Any) error {
	i, err := a.convIdx(ix)
	if err != nil {
		return err
	}
	if i < 0 || i >= len(a.values) {
		return arrIndexOutOfBounsErr(ix)
	}
	a.values[i] = v
	return nil
}

func (a *Array) convIdx(ix Arithmetic) (int, error) {
	if ix.Type() != a.it {
		return -1, arrWrongIndexTypeErr(ix.Type(), a.it)
	}
	diff, err := ix.OpMinus(a.left)
	if err != nil {
		return -1, err
	}
	switch i := diff.(type) {
	case *Integer:
		return int(i.Value), nil
	case *Character:
		return int(i.Value), nil
	default:
		return -1, arrUnprocessableIndexTypeErr(diff.Type())
	}
}
