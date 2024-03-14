package object

import "github.com/pspiagicw/hotshot/ast"

type EvalFunc func(ast.Statement) Object

func multiplyFunc(args []Object) Object {

	err := assertArgs("MULTIPLY", args)
	if err != nil {
		return err
	}
	result := 1
	for _, arg := range args {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("MULTIPLY function expects Integer, found %v", arg.Type())
		}

		result *= v.Value
	}

	return &Integer{
		Value: result,
	}

}
func divideFunc(args []Object) Object {

	err := assertArgs("MULTIPLY", args)
	if err != nil {
		return err
	}

	first := args[0]

	v, ok := first.(*Integer)

	if !ok {
		return createError("DIVIDE function expects Integer, found %v", first.Type())
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("DIVIDE function expects Integer, found %v", arg.Type())
		}

		result /= v.Value
	}

	return &Integer{
		Value: result,
	}
}
func minusFunc(args []Object) Object {

	if len(args) == 0 {
		return createError("No arguments given to MINUS function!")
	}

	first := args[0]

	v, ok := first.(*Integer)

	if !ok {
		return createError("MINUS function expects Integer, found %v", first.Type())
	}

	result := v.Value
	for _, arg := range args[1:] {
		v, ok := arg.(*Integer)

		if !ok {
			return createError("MINUS function expects Integer, found %v", arg.Type())
		}

		result -= v.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args []Object) Object {

	if len(args) == 0 {
		return createError("No arguments given to ADD function!")
	}

	result := 0
	for _, arg := range args {

		v, ok := arg.(*Integer)

		if !ok {
			return createError("ADD function expects Integer, found %v", arg.Type())
		}

		result += v.Value
	}

	return &Integer{
		Value: result,
	}

}
func gtFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("GT function needs 2 arguments!")
	}

	left := args[0]
	f, ok := left.(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", left.Type())
	}

	right := args[1]
	s, ok := right.(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", right.Type())
	}

	return &Boolean{
		Value: f.Value > s.Value,
	}

}
func ltFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("LT function needs 2 arguments!")
	}

	left := args[0]
	f, ok := left.(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", left.Type())
	}

	right := args[1]
	s, ok := right.(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", right.Type())
	}

	return &Boolean{
		Value: f.Value < s.Value,
	}

}
func equalFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("EQ function expects 2 arguments!")
	}

	left := args[0]
	right := args[1]

	if left.Type() == right.Type() && left.String() == right.String() {
		return &Boolean{
			Value: true,
		}
	}

	return &Boolean{
		Value: false,
	}

}
func modFunc(args []Object) Object {
	if len(args) != 2 {
		return createError("MOD function expects 2 arguments")
	}

	left := args[0]
	f, ok := left.(*Integer)
	if !ok {
		return createError("MOD function expects Integer, found %v", left.Type())
	}

	right := args[1]
	s, ok := right.(*Integer)
	if !ok {
		return createError("MOD function expects Integer, found %v", right.Type())
	}

	return &Integer{
		Value: f.Value % s.Value,
	}
}
