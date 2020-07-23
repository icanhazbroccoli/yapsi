package object

import (
	"bytes"
	"strings"

	"yapsi/pkg/ast"
	"yapsi/pkg/types"
)

type Procedure struct {
	Identifier string
	Params     []*Variable
	Body       *ast.BlockStmt
}

var _ (types.Any) = (*Function)(nil)

func (p *Procedure) Type() types.Type { return types.Procedure }
func (p *Procedure) Arity() int       { return len(p.Params) }

func (p *Procedure) String() string {
	var out bytes.Buffer
	out.WriteString("procedure ")
	out.WriteString(p.Identifier)
	out.WriteRune('(')
	chunks := make([]string, 0, len(p.Params))
	for _, param := range p.Params {
		chunks = append(chunks, string(param.Type().Name()))
	}
	out.WriteString(strings.Join(chunks, ", "))
	out.WriteRune(')')
	return out.String()
}
