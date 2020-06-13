package builtin

import (
	"fmt"
	"strings"

	"yapsi/pkg/object"
	"yapsi/pkg/types"
)

var WriteLn *object.Builtin = object.NewBuiltin(
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
		bs, _ := fmt.Println(strings.Join(chunks, " "))
		return &object.Integer{Value: int64(bs)}, nil
	},
)

var Len *object.Builtin = object.NewBuiltin(
	[]*object.Variable{
		object.NewVariable(
			object.VarName("v"),
			nil,
			nil,
		),
	},
	false,
	types.Int,
	func(env *object.Environment, args ...object.Any) (object.Any, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("wrong number of arguments on len() call: given: %d, expected: %d",
				len(args), 1)
		}
		arg := args[0]
		switch ixbl := arg.(type) {
		case object.Indexable:
			return &object.Integer{Value: int64(ixbl.Len())}, nil
		default:
			return nil, fmt.Errorf("object of non-indexable type %s does not support len()",
				arg.Type().Name())
		}
	},
)
