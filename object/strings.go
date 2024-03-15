package object

import "strings"

func concatFunc(args []Object) Object {
	err := assertArgs("CONCAT", args)
	if err != nil {
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
func getCharFunc(args []Object) Object {
	err := assertArity("GETCHAR", args, 2)
	if err != nil {
		return err
	}

	strValue, ok := args[0].(*String)

	if !ok {
		return createError("GETCHAR expected string, got %v", args[0].Type())
	}

	indexValue, ok := args[1].(*Integer)
	if !ok {
		return createError("GETCHAR expected int as index, got %v", args[1].Type())
	}

	if indexValue.Value >= len(strValue.Value) {
		return createError("Index out of bounds")
	}

	return &String{
		Value: string(strValue.Value[indexValue.Value]),
	}
}
