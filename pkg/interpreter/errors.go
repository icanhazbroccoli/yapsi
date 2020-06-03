package interpreter

import (
	"fmt"
)

func unexpectedVisitorResultTypeErr(v interface{}, wantType string) error {
	return fmt.Errorf("Unexpected visitor result type: got: %T, want: %s",
		v, wantType)
}

func unsupportedUnaryOpErr(op string) error {
	return fmt.Errorf("Unsupported operator: %q", op)
}

func undefinedIdentErr(ident string) error {
	return fmt.Errorf("Variable %q is not defined", ident)
}

func wrongIfCondTypErr(cond interface{}) error {
	return fmt.Errorf("Condition type must be a boolean expression, got: %T", cond)
}
