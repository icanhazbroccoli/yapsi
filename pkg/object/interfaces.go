package object

type Arithmetic interface {
	OpUnPlus() (Any, error)
	OpUnMinus() (Any, error)
	OpPlus(Any) (Any, error)
	OpMinus(Any) (Any, error)
	OpAsterisk(Any) (Any, error)
	OpSlash(Any) (Any, error)
	OpPercent(Any) (Any, error)
}

type Comparable interface {
	OpEql(Any) (*Boolean, error)
	OpNeql(Any) (*Boolean, error)
	OpGt(Any) (*Boolean, error)
	OpGte(Any) (*Boolean, error)
	OpLt(Any) (*Boolean, error)
	OpLte(Any) (*Boolean, error)
}

type Logical interface {
	OpNot() (Any, error)
	OpAnd(Any) (Any, error)
	OpOr(Any) (Any, error)
}

type Any interface {
	Type() Type
}
