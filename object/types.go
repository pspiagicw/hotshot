package object

import "fmt"

func typeFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("TYPE function expects 1 argument.")
	}

	value := args[0]

	return &String{
		Value: fmt.Sprintf("%s", value.Type()),
	}
}
func numberpFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("NUMBERP function expects 2 arguments.")
	}

	value := args[0]

	return &Boolean{
		Value: value.Type() == INTEGER_OBJ,
	}
}

func stringpFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("STRINGP function expects 2 arguments.")
	}

	value := args[0]

	return &Boolean{
		Value: value.Type() == STRING_OBJ,
	}
}

func tablepFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("TABLEP function expects 2 arguments.")
	}

	value := args[0]

	return &Boolean{
		Value: value.Type() == TABLE_OBJ,
	}
}
func functionpFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("FUNCTIONP function expects 2 arguments.")
	}

	value := args[0]

	return &Boolean{
		Value: value.Type() == FUNCTION_OBJ,
	}
}
