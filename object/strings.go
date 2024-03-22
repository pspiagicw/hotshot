package object

import "strings"

func concatFunc(args []Object) Object {
	if err := assertArgs("CONCAT", args); err != nil {
		return err
	}

	var out strings.Builder

	for _, arg := range args {
		out.WriteString(arg.String())
	}

	return &String{
		Value: out.String(),
	}
}
func stringFunc(args []Object) Object {
	err := assertArity("STRING", args, 1)
	if err != nil {
		return err
	}

	return &String{
		Value: args[0].String(),
	}
}
func validateString(obj Object) Object {
	if obj.Type() != STRING_OBJ {
		return createError("Object not a string: %v", obj.Type())
	}
	return nil
}
func validateInteger(obj Object) Object {
	if obj.Type() != INTEGER_OBJ {
		return createError("Object not an integer: %v", obj.Type())
	}
	return nil
}

func getCharFunc(args []Object) Object {
	if err := assertArity("GETCHAR", args, 2); err != nil {
		return err
	}

	if err := validateString(args[0]); err != nil {
		return err
	}
	if err := validateInteger(args[1]); err != nil {
		return err
	}

	strValue := args[0].(*String)

	indexValue := args[1].(*Integer)

	if indexValue.Value >= len(strValue.Value) {
		return createError("Index out of bounds")
	}

	return &String{
		Value: string(strValue.Value[indexValue.Value]),
	}
}
