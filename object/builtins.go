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
		"not": &Builtin{
			Fn: notFunc,
		},
		"and": &Builtin{
			Fn: andFunc,
		},
		"or": &Builtin{
			Fn: orFunc,
		},
		"do": &Builtin{
			Fn: doFunc,
		},
	}
}
func doFunc(args ...Object) Object {
	length := len(args)

	return args[length-1]
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

	output.WriteString(" ")

	for _, arg := range args {
		output.WriteString(arg.String())
		output.WriteString(" ")
	}

	fmt.Println(output.String())

	return Null{}
}

func createError(format string, v ...interface{}) Object {
	return &Error{
		Message: fmt.Sprintf("ERROR: %s", fmt.Sprintf(format, v...)),
	}
}
