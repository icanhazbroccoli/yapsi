package interpreter

type Interpreter struct{}

var _ ast.NodeVisitor = (*Interpreter)(nil)

func (i *Interpreter) VisitNumericLiteral(node *ast.NumericLiteral) ast.VisitorResult {

}
