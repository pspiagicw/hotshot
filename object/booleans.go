package object

func orFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("OR function expects 2 arguments")
	}

	f, ok := args[0].(*Boolean)
	if !ok {
		return createError("OR function expects Boolean, found %v", args[0].Type())
	}

	s, ok := args[1].(*Boolean)
	if !ok {
		return createError("OR function expects Boolean, found %v", args[0].Type())
	}

	return &Boolean{
		Value: s.Value || f.Value,
	}

}
func andFunc(args ...Object) Object {
	if len(args) != 2 {
		return createError("AND function expects 2 arguments")
	}

	f, ok := args[0].(*Boolean)
	if !ok {
		return createError("AND function expects Boolean, found %v", args[0].Type())
	}

	s, ok := args[1].(*Boolean)
	if !ok {
		return createError("AND function expects Boolean, found %v", args[0].Type())
	}

	return &Boolean{
		Value: s.Value && f.Value,
	}

}
func notFunc(args ...Object) Object {
	if len(args) != 1 {
		return createError("NOT function expects 1 argument")
	}

	v, ok := args[0].(*Boolean)
	if !ok {
		return createError("NOT function expects Boolean, found %v", args[0].Type())
	}

	return &Boolean{
		Value: !v.Value,
	}
}
