package builtin

import (
	"fmt"
	"strings"

	"yapsi/pkg/object"
	"yapsi/pkg/types"
)

var WriteLn *object.Function = object.NewFunction(
	[]*object.Variable{
		object.NewVariable(
			object.VarName("args"),
			types.Array,
			nil,
		),
	},
	true,
	types.Int,
	func(env *object.Environment, args ...object.Any) (object.Any, error) {
		chunks := make([]string, 0, len(args))
		for _, arg := range args {
			chunks = append(chunks, arg.String())
		}
		fmt.Println(strings.Join(chunks, " "))
		return nil, nil
	},
)
