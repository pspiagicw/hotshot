package object

func orFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("or expects 2 arguments!")
	}

	f, ok := args[0].(*Boolean)
	if !ok {
		return createError("Boolean expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Boolean)
	if !ok {
		return createError("Boolean expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: s.Value || f.Value,
	}

}
func andFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("and expects 2 arguments!")
	}

	f, ok := args[0].(*Boolean)
	if !ok {
		return createError("Boolean expected, got %v", args[0].Type())
	}

	s, ok := args[1].(*Boolean)
	if !ok {
		return createError("Boolean expected, got %v", args[1].Type())
	}

	return &Boolean{
		Value: s.Value && f.Value,
	}

}
func notFunc(args ...Object) Object {
	if len(args) != 1 {
		return createError("not expects 1 arguments!")
	}

	v, ok := args[0].(*Boolean)
	if !ok {
		return createError("Boolean expected, got %v", args[0].Type())
	}

	return &Boolean{
		Value: !v.Value,
	}
}
