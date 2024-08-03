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
	err := assertArgs("NUMBERP", args)
	if err != nil {
		return err
	}

	result := true
	for _, arg := range args {
		if arg.Type() != INTEGER_OBJ {
			result = false
		}
	}

	return &Boolean{
		Value: result,
	}
}

func stringpFunc(args []Object) Object {
	err := assertArgs("STRINGP", args)
	if err != nil {
		return err
	}

	result := true
	for _, arg := range args {
		if arg.Type() != STRING_OBJ {
			result = false
		}
	}

	return &Boolean{
		Value: result,
	}
}

func tablepFunc(args []Object) Object {
	err := assertArgs("TABLEP", args)
	if err != nil {
		return err
	}

	result := true
	for _, arg := range args {
		if arg.Type() != TABLE_OBJ {
			result = false
		}
	}

	return &Boolean{Value: result}

}
func functionpFunc(args []Object) Object {
	err := assertArgs("FUNCTIONP", args)
	if err != nil {
		return err
	}

	result := true
	for _, arg := range args {
		if arg.Type() != FUNCTION_OBJ || arg.Type() != COMPILED_FUNCTION_OBJ {
			result = false
		}
	}

	return &Boolean{Value: result}
}
