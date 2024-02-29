package object

import (
	"fmt"
	"strings"
)

func getBuiltins() map[string]Object {
	return map[string]Object{
		"+": &Builtin{
			Fn: addFunc,
		},
		"-": &Builtin{
			Fn: minusFunc,
		},
		"*": &Builtin{
			Fn: multiplyFunc,
		},
		"?": &Builtin{
			Fn: printFunc,
		},
	}
}
func printFunc(args ...Object) Object {
	var output strings.Builder

	for _, arg := range args {
		output.WriteString(arg.String())
		output.WriteString(" ")
	}

	return &String{
		Value: output.String(),
	}
}

func multiplyFunc(args ...Object) Object {
	if len(args) == 0 {
		return createError("No arguments!")
	}

	result := 1
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result *= v.Value
	}

	return &Integer{
		Value: result,
	}

}
func minusFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments!")
	}

	f := args[0]

	v, ok := f.(*Integer)

	if !ok {
		return createError("Argument not integer!")
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result -= v.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments!")
	}

	result := 0
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result += v.Value
	}

	return &Integer{
		Value: result,
	}

}

func createError(format string, v ...interface{}) Object {
	return &Error{
		Message: fmt.Sprintf("ERROR: %s", fmt.Sprintf(format, v...)),
	}
}
