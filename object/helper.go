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
