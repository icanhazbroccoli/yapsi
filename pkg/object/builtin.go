package object

import (
	"fmt"

	"yapsi/pkg/types"
)

type BuiltinBody func(*Environment, ...Any) (Any, error)

type Builtin struct {
	params   []*Variable
	variadic bool
	returns  *types.Type
	body     BuiltinBody
}

var _ Callable = (*Builtin)(nil)

func NewBuiltin(params []*Variable, variadic bool, returns *types.Type, body BuiltinBody) *Builtin {
	return &Builtin{
		params:   params,
		variadic: variadic,
		returns:  returns,
		body:     body,
	}
}

func (f *Builtin) Type() *types.Type    { return types.Builtin }
func (f *Builtin) Arity() int           { return len(f.params) }
func (f *Builtin) IsVariadic() bool     { return f.variadic }
func (f *Builtin) Returns() *types.Type { return f.returns }
func (f *Builtin) String() string       { return "(function)" }

func (f *Builtin) CallReturn(env *Environment, args ...Any) (Any, error) {
	if !f.variadic && len(args) != f.Arity() {
		return nil, fmt.Errorf("Wrong number of arguments for non-variadic function: got %d, want: %d",
			len(args), f.Arity())
	}
	localEnv := NewEnvironment(env)
	return f.body(localEnv, args...)
}

func (f *Builtin) Call(*Environment, ...Any) error {
	panic("Builtins calls must be returnable; use `CallReturn` instead")
}
