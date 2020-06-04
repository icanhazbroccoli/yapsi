package object

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
	OpEql(Any) (*Boolean, error)
	OpNeql(Any) (*Boolean, error)
	OpGt(Any) (*Boolean, error)
	OpGte(Any) (*Boolean, error)
	OpLt(Any) (*Boolean, error)
	OpLte(Any) (*Boolean, error)
}

type Logical interface {
	Any
	OpNot() (Any, error)
	OpAnd(Any) (Any, error)
	OpOr(Any) (Any, error)
}

type Indexable interface {
	Any
	OpSubscrGet(Arithmetic) (Any, error)
	OpSubscrSet(Arithmetic, Any) error
}

type Any interface {
	Type() *Type
}
