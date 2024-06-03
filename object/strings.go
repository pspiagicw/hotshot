package object

import "strings"

func concatFunc(args []Object) Object {
	if err := assertArgs("CONCAT", args); err != nil {
		return err
	}

	var out strings.Builder

	for _, arg := range args {
		s, ok := arg.(*String)

		if !ok {
			return createError("Argument is not a string")
		}
		out.WriteString(s.Value)
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
		Value: args[0].Content(),
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
func subStringFunc(args []Object) Object {
	if err := assertArity("SUBSTRING", args, 3); err != nil {
		return err
	}

	if err := validateString(args[0]); err != nil {
		return err
	}
	if err := validateInteger(args[1]); err != nil {
		return err
	}
	if err := validateInteger(args[2]); err != nil {
		return err
	}
	strValue := args[0].(*String)
	startValue := args[1].(*Integer)
	endValue := args[2].(*Integer)

	if startValue.Value > endValue.Value {
		return createError("Start index greater than end index")
	}
	if endValue.Value > len(strValue.Value) {
		return createError("End index out of bounds")
	}

	if startValue.Value < 0 {
		return createError("Start index out of bounds")
	}

	return &String{
		Value: strValue.Value[startValue.Value:endValue.Value],
	}
}
func lowerFunc(args []Object) Object {
	if err := assertArity("LOWER", args, 1); err != nil {
		return err
	}

	if err := validateString(args[0]); err != nil {
		return err
	}

	strValue := args[0].(*String)

	return &String{
		Value: strings.ToLower(strValue.Value),
	}
}
func upperFunc(args []Object) Object {
	if err := assertArity("UPPER", args, 1); err != nil {
		return err
	}

	if err := validateString(args[0]); err != nil {
		return err
	}

	strValue := args[0].(*String)

	return &String{
		Value: strings.ToUpper(strValue.Value),
	}
}
