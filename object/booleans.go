package object

func orFunc(args []Object) Object {
	err, values := assertArgsBoolean("OR", args)
	if err != nil {
		return err
	}

	result := false
	for _, val := range values {
		result = result || val.Value
	}

	return &Boolean{
		Value: result,
	}

}
func andFunc(args []Object) Object {
	err, values := assertArgsBoolean("AND", args)
	if err != nil {
		return err
	}

	result := true
	for _, val := range values {
		result = result && val.Value
	}

	return &Boolean{
		Value: result,
	}

}
func notFunc(args []Object) Object {

	err, values := assertArgsBoolean("NOT", args)
	if err != nil {
		return err
	}

	result := !values[0].Value
	for _, val := range values[1:] {
		result = result && !val.Value
	}

	return &Boolean{
		Value: result,
	}
}
