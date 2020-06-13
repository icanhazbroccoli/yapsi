package object

import (
	"bytes"
	"strings"

	"yapsi/pkg/ast"
	"yapsi/pkg/types"
)

type Function struct {
	Identifier string
	ReturnType *types.Type
	Params     []*Variable
	Body       *ast.BlockStmt
}

var _ (Any) = (*Function)(nil)

//var _ (Callable) = (*Function)(nil)

func (f *Function) Type() *types.Type { return types.Function }
func (f *Function) Arity() int        { return len(f.Params) }

//func (f *Function) IsVariadic() bool     { return false }
//func (f *Function) Returns() *types.Type { return f.ReturnType }
//
func (f *Function) String() string {
	var out bytes.Buffer
	out.WriteString("function ")
	out.WriteString(f.Identifier)
	out.WriteRune('(')
	chunks := make([]string, 0, len(f.Params))
	for _, p := range f.Params {
		chunks = append(chunks, string(p.Type().Name()))
	}
	out.WriteString(strings.Join(chunks, ", "))
	out.WriteRune(')')
	return out.String()
}

//
//func (f *Function) CallReturn(env *Environment, args ...Any) (Any, error) {
//	if len(args) != f.Arity() {
//		return nil, fmt.Errorf("%s: Wrong number of arguments: got: %d, want: %d",
//			f.Identifier, len(args), f.Arity())
//	}
//	callEnv := NewEnvironment(env)
//	for i := 0; i < len(f.Params); i++ {
//		callEnv.DeclareVar()
//	}
//}
//
//func (f *Function) Call(*Environment, ...Any) error {
//	panic("Function calls must be returnable")
//}
