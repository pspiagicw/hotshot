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
func gtFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("lt expects 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value > s.Value,
	}

}
func ltFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("lt expects 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("Integer expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value < s.Value,
	}

}
func equalFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("eq expects 2 arguments!")
	}

	if args[0].Type() == args[1].Type() && args[0].String() == args[1].String() {
		return &Boolean{
			Value: true,
		}
	}

	return &Boolean{
		Value: false,
	}

}
