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
		"/": &Builtin{
			Fn: divideFunc,
		},
		"?": &Builtin{
			Fn: printFunc,
		},
		"#": &Builtin{
			Fn: lenFunc,
		},
		"=": &Builtin{
			Fn: equalFunc,
		},
		"<": &Builtin{
			Fn: ltFunc,
		},
		">": &Builtin{
			Fn: gtFunc,
		},
	}
}
func gtFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("lt expects 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value > s.Value,
	}

}
func ltFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("lt expects 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value < s.Value,
	}

}
func equalFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("eq expects 2 arguments!")
	}

	if args[0].Type() == args[1].Type() && args[0].String() == args[1].String() {
		return &Boolean{
			Value: true,
		}
	}

	return &Boolean{
		Value: false,
	}

}
func lenFunc(args ...Object) Object {
	if len(args) == 0 {
		return createError("No arguments!")
	}

	if len(args) > 1 {
		return createError("Length function expects only 1 argument!")
	}

	switch v := args[0].(type) {
	case *String:
		return &Integer{
			Value: len(v.Value),
		}
	}

	return createError("Can't find length of that type")
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

func createError(format string, v ...interface{}) Object {
	return &Error{
		Message: fmt.Sprintf("ERROR: %s", fmt.Sprintf(format, v...)),
	}
}
