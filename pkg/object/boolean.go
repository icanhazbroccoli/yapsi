package object

type Boolean struct {
	Value bool
}

var _ (Any) = (*Boolean)(nil)
var _ (Logical) = (*Boolean)(nil)

func (b *Boolean) Type() Type { return BOOL }

func (b *Boolean) OpNot() (Any, error) {
	return &Boolean{
		Value: !b.Value,
	}, nil
}

func (b *Boolean) OpAnd(o Any) (Any, error) {
	switch b2 := o.(type) {
	case *Boolean:
		return &Boolean{
			Value: b.Value && b2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(b, o, "and")
	}
}

func (b *Boolean) OpOr(o Any) (Any, error) {
	switch b2 := o.(type) {
	case *Boolean:
		return &Boolean{
			Value: b.Value || b2.Value,
		}, nil
	default:
		return nil, WrongOperandTypErr(b, o, "or")
	}
}
