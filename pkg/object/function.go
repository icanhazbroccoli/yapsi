package object

import (
	"bytes"
	"strings"

	"yapsi/pkg/ast"
	"yapsi/pkg/types"
)

type Function struct {
	Identifier string
	ReturnType types.Type
	Params     []*Variable
	Body       *ast.BlockStmt
}

var _ (types.Any) = (*Function)(nil)

func (f *Function) Type() types.Type { return types.Function }
func (f *Function) Arity() int       { return len(f.Params) }

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
