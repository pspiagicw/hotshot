package object

import (
	"fmt"
)

func assertArity(name string, args []Object, expected int) Object {
	if len(args) != expected {
		return createError("%s function expects %d arguments, given %d", name, expected, len(args))
	}
	return nil
}
func assertArgs(name string, args []Object) Object {
	if len(args) == 0 {
		return createError("%s function not given any arguments", name)
	}
	return nil
}

func createError(format string, v ...interface{}) Object {
	return &Error{
		Message: fmt.Sprintf(format, v...),
	}
}
func assertArgsInteger(name string, args []Object) (Object, []*Integer) {
	if len(args) == 0 {
		return createError("%s function not given any arguments", name), nil
	}

	results := []*Integer{}
	for _, arg := range args {
		v, ok := arg.(*Integer)
		if !ok {
			return createError("%s function expects Integer, found %v", name, arg.Type()), nil
		}
		results = append(results, v)
	}
	return nil, results
}
func assertArgsBoolean(name string, args []Object) (Object, []*Boolean) {
	if len(args) == 0 {
		return createError("%s function not given any arguments", name), nil
	}

	results := []*Boolean{}
	for _, arg := range args {
		v, ok := arg.(*Boolean)
		if !ok {
			return createError("%s function expects Boolean, found %v", name, arg.Type()), nil
		}
		results = append(results, v)
	}
	return nil, results
}
