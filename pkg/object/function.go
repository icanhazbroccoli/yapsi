package object

import (
	"fmt"

	"yapsi/pkg/types"
)

type FuncBody func(*Environment, ...Any) (Any, error)

type Function struct {
	params   []*Variable
	variadic bool
	returns  *types.Type
	body     FuncBody
}

var _ Callable = (*Function)(nil)

func NewFunction(params []*Variable, variadic bool, returns *types.Type, body FuncBody) *Function {
	return &Function{
		params:   params,
		variadic: variadic,
		returns:  returns,
		body:     body,
	}
}

func (f *Function) Type() *types.Type    { return types.Function }
func (f *Function) Arity() int           { return len(f.params) }
func (f *Function) IsVariadic() bool     { return f.variadic }
func (f *Function) Returns() *types.Type { return f.returns }
func (f *Function) String() string       { return "(function)" }

func (f *Function) CallReturn(env *Environment, args ...Any) (Any, error) {
	if !f.variadic && len(args) != f.Arity() {
		return nil, fmt.Errorf("Wrong number of arguments for non-variadic function: got %d, want: %d",
			len(args), f.Arity())
	}
	localEnv := NewEnvironment(env)
	return f.body(localEnv, args...)
}

func (f *Function) Call(*Environment, ...Any) error {
	panic("Functions calls must be returnable; use `CallReturn` instead")
}
