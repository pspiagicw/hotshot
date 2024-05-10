package object

import (
	"fmt"
	"os"
)

func exitFunc(args []Object) Object {
	os.Exit(0)
	return &Null{}
}
func inputFunc(args []Object) Object {
	e := assertArity("INPUT", args, 1)

	if e != nil {
		return e
	}

	v, ok := args[0].(*String)

	if !ok {
		return createError("INPUT function expects String, found %v", args[0].Type())
	}

	var input string

	fmt.Print(v.Value)
	_, err := fmt.Scanln(&input)

	if err != nil {
		return createError("Error reading input, %v", e)
	}

	return &String{
		Value: input,
	}

}
