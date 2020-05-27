package object

type Type uint16

const (
	NONE Type = 0

	BOOL Type = 1 << iota
	INT
	REAL
	STRING
	CHAR
	PTR
)

func (t Type) Compatible(t2 Type) bool {
	t1, t2 := t, t2
	if t1 > t2 {
		t2, t1 = t1, t2
	}
	if t1 == t2 {
		return true
	}
	switch t1 {
	case INT:
		return t2 == REAL
	default:
		return false
	}
}
