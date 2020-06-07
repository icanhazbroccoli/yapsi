package interpreter

import (
	"fmt"
	"yapsi/pkg/types"
)

func unexpectedVisitorResultTypeErr(v interface{}, wantType string) error {
	return fmt.Errorf("Unexpected visitor result type: got: %T, want: %s",
		v, wantType)
}

func unsupportedUnaryOpErr(op string) error {
	return fmt.Errorf("Unsupported operator: %q", op)
}

func unsupportedBinaryOpErr(op string, t1, t2 *types.Type) error {
	return fmt.Errorf("Unsupported operator: %s %s %s",
		t1.Name(), op, t2.Name())
}

func undefinedIdentErr(ident string) error {
	return fmt.Errorf("Variable %q is not defined", ident)
}

func uncallableEntityErr(ident string) error {
	return fmt.Errorf("Variable %q is not callable", ident)
}

func wrongIfCondTypErr(cond interface{}) error {
	return fmt.Errorf("Condition type must be a boolean expression, got: %T", cond)
}
