package object

func orFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("OR function expects 2 arguments")
	}

	left := args[0]
	f, ok := left.(*Boolean)
	if !ok {
		return createError("OR function expects Boolean, found %v", left.Type())
	}

	right := args[1]
	s, ok := right.(*Boolean)
	if !ok {
		return createError("OR function expects Boolean, found %v", right.Type())
	}

	return &Boolean{
		Value: s.Value || f.Value,
	}

}
func andFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("AND function expects 2 arguments")
	}

	left := args[0]
	f, ok := left.(*Boolean)
	if !ok {
		return createError("AND function expects Boolean, found %v", left.Type())
	}

	right := args[1]
	s, ok := right.(*Boolean)
	if !ok {
		return createError("AND function expects Boolean, found %v", right.Type())
	}

	return &Boolean{
		Value: s.Value && f.Value,
	}

}
func notFunc(args []Object) Object {
	if len(args) != 1 {
		return createError("NOT function expects 1 argument")
	}

	value := args[0]
	v, ok := value.(*Boolean)
	if !ok {
		return createError("NOT function expects Boolean, found %v", value.Type())
	}

	return &Boolean{
		Value: !v.Value,
	}
}
