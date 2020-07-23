package types

type Any interface {
	Type() Type
	String() string
}

type Arithmetic interface {
	Any
	OpUnPlus() (Any, error)
	OpUnMinus() (Any, error)
	OpPlus(Any) (Any, error)
	OpMinus(Any) (Any, error)
	OpAsterisk(Any) (Any, error)
	OpSlash(Any) (Any, error)
	OpPercent(Any) (Any, error)
}

type Comparable interface {
	Any
	OpEql(Any) (Any, error)
	OpNeql(Any) (Any, error)
	OpGt(Any) (Any, error)
	OpGte(Any) (Any, error)
	OpLt(Any) (Any, error)
	OpLte(Any) (Any, error)
}

/*
type Callable interface {
	Any
	Arity() int
	IsVariadic() bool
	Returns() types.Type
	Call(*Environment, ...Any) error
	CallReturn(*Environment, ...Any) (Any, error)
}
*/

type Logical interface {
	Any
	OpNot() (Any, error)
	OpAnd(Any) (Any, error)
	OpOr(Any) (Any, error)
}

type Indexable interface {
	Any
	Len() int
	OpSubscrGet(Arithmetic) (Any, error)
	OpSubscrSet(Arithmetic, Any) error
}
