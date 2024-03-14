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
