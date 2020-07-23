package interpreter

import (
	"fmt"

	"yapsi/pkg/types"
)

func assertTypeEqual(t1, t2 types.Type) error {
	if t1.Name() != t2.Name() {
		return fmt.Errorf("Types do not match: %s Vs %s", t1.Name(), t2.Name())
	}
	return nil
}
