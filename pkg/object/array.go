package object

type ArrayTypeDefinition struct {
	IndexType     *Type
	ComponentType *Type
	Left, Right   Arithmetic
}

var _ TypeDefinition = (*ArrayTypeDefinition)(nil)

type Array struct {
	typ    *Type
	values []Any
}

var _ Indexable = (*Array)(nil)

func (a *Array) Type() *Type { return a.typ }

func NewArray(l, r Arithmetic, it *Type, ct *Type) (*Array, error) {
	def := &ArrayTypeDefinition{
		IndexType:     it,
		ComponentType: ct,
		Left:          l,
		Right:         r,
	}
	a := &Array{
		typ: NewType(ARRAY, def, nil),
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
	def := a.typ.def.(*ArrayTypeDefinition)
	if ix.Type() != def.IndexType {
		return -1, arrWrongIndexTypeErr(ix.Type(), def.IndexType)
	}
	diff, err := ix.OpMinus(def.Left)
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
