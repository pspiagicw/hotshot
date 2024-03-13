package object

import "github.com/pspiagicw/hotshot/ast"

type EvalFunc func(ast.Statement) Object

func multiplyFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) == 0 {
		return createError("No arguments given to MULTIPLY function!")
	}

	result := 1
	for _, arg := range args {
		value := evalFunc(arg)
		v, ok := value.(*Integer)

		if !ok {
			return createError("MULTIPLY function expects Integer, found %v", value.Type())
		}

		result *= v.Value
	}

	return &Integer{
		Value: result,
	}

}
func divideFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {

	if len(args) == 0 {
		return createError("No arguments given to DIVIDE function!")
	}

	f := evalFunc(args[0])

	v, ok := f.(*Integer)

	if !ok {
		return createError("DIVIDE function expects Integer, found %v", f.Type())
	}

	result := v.Value
	for _, arg := range args[1:] {
		value := evalFunc(arg)
		v, ok := value.(*Integer)

		if !ok {
			return createError("DIVIDE function expects Integer, found %v", value.Type())
		}

		result /= v.Value
	}

	return &Integer{
		Value: result,
	}
}
func minusFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {

	if len(args) == 0 {
		return createError("No arguments given to MINUS function!")
	}

	f := evalFunc(args[0])

	v, ok := f.(*Integer)

	if !ok {
		return createError("MINUS function expects Integer, found %v", f.Type())
	}

	result := v.Value
	for _, arg := range args[1:] {
		value := evalFunc(arg)
		v, ok := value.(*Integer)

		if !ok {
			return createError("MINUS function expects Integer, found %v", value.Type())
		}

		result -= v.Value
	}

	return &Integer{
		Value: result,
	}
}

func addFunc(args []ast.Statement, evalFunc func(node ast.Statement) Object, env *Environment) Object {

	if len(args) == 0 {
		return createError("No arguments given to ADD function!")
	}

	result := 0
	for _, arg := range args {

		value := evalFunc(arg)
		v, ok := value.(*Integer)

		if !ok {
			return createError("ADD function expects Integer, found %v", value.Type())
		}

		result += v.Value
	}

	return &Integer{
		Value: result,
	}

}
func gtFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) != 2 {
		return createError("GT function needs 2 arguments!")
	}

	left := evalFunc(args[0])
	f, ok := left.(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", left.Type())
	}

	right := evalFunc(args[1])
	s, ok := right.(*Integer)
	if !ok {
		return createError("GT function expects Integer, found %v", right.Type())
	}

	return &Boolean{
		Value: f.Value > s.Value,
	}

}
func ltFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) != 2 {
		return createError("LT function needs 2 arguments!")
	}

	left := evalFunc(args[0])
	f, ok := left.(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", left.Type())
	}

	right := evalFunc(args[1])
	s, ok := right.(*Integer)
	if !ok {
		return createError("LT function expects Integer, found %v", right.Type())
	}

	return &Boolean{
		Value: f.Value < s.Value,
	}

}
func equalFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) != 2 {
		return createError("EQ function expects 2 arguments!")
	}

	left := evalFunc(args[0])
	right := evalFunc(args[1])

	if left.Type() == right.Type() && left.String() == right.String() {
		return &Boolean{
			Value: true,
		}
	}

	return &Boolean{
		Value: false,
	}

}
func modFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) != 2 {
		return createError("MOD function expects 2 arguments")
	}

	left := evalFunc(args[0])
	f, ok := left.(*Integer)
	if !ok {
		return createError("MOD function expects Integer, found %v", left.Type())
	}

	right := evalFunc(args[1])
	s, ok := right.(*Integer)
	if !ok {
		return createError("MOD function expects Integer, found %v", right.Type())
	}

	return &Integer{
		Value: f.Value % s.Value,
	}
}
