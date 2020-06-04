package object

import (
	"fmt"
	"strconv"

	"yapsi/pkg/ast"
)

func ParseInt(num ast.RawNumber) (int64, error) {
	if !LooksLikeInt(num) {
		return 0, fmt.Errorf("Unable to parse int value: %q", num)
	}
	v, err := strconv.Atoi(string(num))
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func LooksLikeInt(num ast.RawNumber) bool {
	res := false
	if len(num) > 0 {
		for ix, ch := range num {
			if ix == 0 && ch == '-' {
				continue
			}
			if !(ch >= '0' && ch <= '9') {
				return false
			}
			res = true
		}
	}
	return res
}

func ParseReal(num ast.RawNumber) (float64, error) {
	return strconv.ParseFloat(string(num), 64)
}
