package object

func multiplyFunc(args ...Object) Object {
	if len(args) == 0 {
		return createError("No arguments given to MULTIPLY function!")
	}

	result := 1
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Arguments given to MULTIPLY are not Integer")
		}

		result *= v.Value
	}

	return &Integer{
		Value: result,
	}

}
func divideFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments given to DIVIDE function!")
	}

	f := args[0]

	v, ok := f.(*Integer)

	if !ok {
		return createError("Arguments given to DIVIDE are not Integer")
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Arguments given to DIVIDE are not Integer")
		}

		result /= v.Value
	}

	return &Integer{
		Value: result,
	}
}
func minusFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments given to MINUS function!")
	}

	f := args[0]

	v, ok := f.(*Integer)

	if !ok {
		return createError("Arguments given to MINUS are not Integer")
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Arguments given to MINUS are not Integer")
		}

		result -= v.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args ...Object) Object {

	if len(args) == 0 {
		return createError("No arguments given to ADD function!")
	}

	result := 0
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("Arguments given to ADD are not Integer")
		}

		result += v.Value
	}

	return &Integer{
		Value: result,
	}

}
func gtFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("GT function needs 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", args[0].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value > s.Value,
	}

}
func ltFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("LT function needs 2 arguments!")
	}

	f, ok := args[0].(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", args[1].Type())
	}

	s, ok := args[1].(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", args[1].Type())
	}

	return &Boolean{
		Value: f.Value < s.Value,
	}

}
func equalFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("EQ function expects 2 arguments!")
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
