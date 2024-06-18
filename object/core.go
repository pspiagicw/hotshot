package object

import (
	"fmt"
	"strings"
)

func returnFunc(args []Object) Object {
	err := assertArity("RETURN", args, 0)
	if err != nil {
		return err
	}

	return &Return{}
}

func doFunc(args []Object) Object {
	length := len(args)

	return args[length-1]
}
func lenFunc(args []Object) Object {
	err := assertArity("LEN", args, 1)
	if err != nil {
		return err
	}

	value := args[0]
	switch v := value.(type) {
	case *String:
		return &Integer{
			Value: len(v.Value),
		}
	case *Table:
		return &Integer{
			Value: len(v.Elements),
		}
	}

	return createError("LEN function can't find length of that type!")
}
func printFunc(args []Object) Object {
	var output strings.Builder

	output.WriteString(" ")

	for _, arg := range args {
		output.WriteString(arg.Content())
		output.WriteString(" ")
	}

	fmt.Println(output.String())

	return Null{}
}
