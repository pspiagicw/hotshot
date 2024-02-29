package object

func multiplyFunc(args ...Object) Object {
	if len(args) == 0 {
		return createError("No arguments!")
	}

	result := 1
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result *= v.Value
	}

	return &Integer{
		Value: result,
	}

}
func divideFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments!")
	}

	f := args[0]

	v, ok := f.(*Integer)

	if !ok {
		return createError("Argument not integer!")
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result /= v.Value
	}

	return &Integer{
		Value: result,
	}
}
func minusFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments!")
	}

	f := args[0]

	v, ok := f.(*Integer)

	if !ok {
		return createError("Argument not integer!")
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result -= v.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments!")
	}

	result := 0
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Argument not integer!")
		}

		result += v.Value
	}

	return &Integer{
		Value: result,
	}

}
