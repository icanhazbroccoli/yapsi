package object

type Variable struct {
	Any
	name string
}

func NewVariable(name string, value Any) *Variable {
	return &Variable{value, name}
}

func (v *Variable) Name() string { return v.name }
