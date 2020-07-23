package object

import (
	"fmt"

	"yapsi/pkg/types"
)

func IncompTypOpErr(o types.Any, opName string) error {
	return fmt.Errorf("Operator %s is not defined on type: %v",
		opName, o.Type())
}

func WrongOperandTypErr(o1, o2 types.Any, opName string) error {
	return fmt.Errorf("Unsupported operator: %v %s %v",
		o1.Type(), opName, o2.Type())
}

func UnsupTypOpErr(t types.Type, op string) error {
	return fmt.Errorf("Unsupported operator for type: %v", t.Name())
}

func arrIndexOutOfBounsErr(ix types.Arithmetic) error {
	return fmt.Errorf("Array index %v is out of bounds", ix)
}

func arrWrongRangeConstraintErr(l, r types.Arithmetic) error {
	return fmt.Errorf("Array range constraint is wrong: Left: %v, Right: %v", l, r)
}

func arrWrongIndexTypeErr(gotType, wantType types.Type) error {
	return fmt.Errorf("Wrong array indexing type provided: got: %s, want: %s",
		gotType.Name(), wantType.Name())
}

func arrUnprocessableIndexTypeErr(t types.Type) error {
	return fmt.Errorf("Array indexing failed for index type: %s", t.Name())
}
