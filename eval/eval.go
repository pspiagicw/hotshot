package eval

import (
	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/object"
)

func Eval(node ast.Statement, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntStatement:
		return &object.Integer{Value: node.Value}
	case *ast.StringStatement:
		return &object.String{Value: node.Value}
	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.FunctionalStatement:
		return evalFunction(node, env)
	case *ast.EmptyStatement:
		return object.Null{}
	}
	return &object.Error{
		Message: "ERROR: Evaluation for statement can't be done!\n",
	}
}
func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, env)
	}

	return result
}
func evalFunction(node *ast.FunctionalStatement, env *object.Environment) object.Object {
	fn, ok := env.Functions[node.Op.TokenValue]

	if !ok {
		return object.Error{
			Message: "ERROR: No function with that signature found!\n",
		}
	}

	v, ok := fn.(*object.Builtin)

	if !ok {
		return object.Error{
			Message: "ERROR: INTERNAL!!! Builtin can't be initialized!\n",
		}
	}

	args := []object.Object{}

	for _, arg := range node.Args {
		args = append(args, Eval(arg, env))
	}

	return v.Fn(args...)
}
