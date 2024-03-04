package eval

import (
	"fmt"

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
	case *ast.CallStatement:
		return evalFunction(node, env)
	case *ast.EmptyStatement:
		return object.Null{}
	case *ast.BoolStatement:
		return &object.Boolean{Value: node.Value}
	case *ast.IdentStatement:
		return evalIdent(node, env)
	case *ast.AssignmentStatement:
		return applyAssignment(node, env)
	case *ast.IfStatement:
		return evalIfStatement(node, env)
	case *ast.WhileStatement:
		return evalWhileStatement(node, env)
	case *ast.FunctionStatement:
		return evalFunctionStatement(node, env)
	}
	return createError("Evaluation for statement can't be done!")
}

func evalFunctionStatement(node *ast.FunctionStatement, env *object.Environment) object.Object {
	name := node.Name.TokenValue

	fn := &object.Function{}

	args := []*ast.IdentStatement{}

	for _, arg := range node.Args {
		v, ok := arg.(*ast.IdentStatement)
		if !ok {
			return createError("Argument is not a identifier!")
		}
		args = append(args, v)
	}

	fn.Args = args

	fn.Body = &node.Body

	env.Functions[name] = fn

	return object.Null{}
}
func evalWhileStatement(node *ast.WhileStatement, env *object.Environment) object.Object {
	var result object.Object

	result = object.Null{}

	for true {
		result := Eval(node.Condition, env)

		if result.Type() != object.BOOLEAN_OBJ {
			return createError("Condition for WHILE doesn't evaluate to true/false!")
		}

		if result.String() == "true" {
			result = Eval(node.Body, env)
		} else {
			break
		}
	}
	return result
}
func evalIfStatement(node *ast.IfStatement, env *object.Environment) object.Object {

	result := Eval(node.Condition, env)

	if result.Type() != object.BOOLEAN_OBJ {
		return createError("Condition for IF doesn't evaluate to true/false!")
	}

	if result.String() == "true" {
		return Eval(node.Body, env)
	} else {
		if node.Else != nil {
			return Eval(node.Else, env)
		}
	}
	return object.Null{}

}
func evalIdent(node *ast.IdentStatement, env *object.Environment) object.Object {
	value, ok := env.Vars[node.Value.TokenValue]

	if !ok {
		return createError("No variable '%s' defined!", node.Value.TokenValue)
	}

	return value
}
func evalProgram(statements []ast.Statement, env *object.Environment) object.Object {
	var result object.Object
	result = &object.Null{}

	for _, statement := range statements {
		result = Eval(statement, env)
	}

	return result
}
func applyAssignment(node *ast.AssignmentStatement, env *object.Environment) object.Object {

	value := Eval(node.Value, env)

	env.Vars[node.Name.TokenValue] = value

	return object.Null{}
}
func evalFunction(node *ast.CallStatement, env *object.Environment) object.Object {
	fn, ok := env.Functions[node.Op.TokenValue]

	args := []object.Object{}

	for _, arg := range node.Args {
		args = append(args, Eval(arg, env))
	}

	if !ok {
		return createError("No function '%s' found!", node.Op.TokenValue)
	}

	v, ok := fn.(*object.Builtin)

	if !ok {
		v, ok := fn.(*object.Function)
		if !ok {
			return createError("INTERNAL! Builtin or user function failed to initalize!")
		} else {
			// Apply user defined function!
			return applyFunction(v, args, env)
		}
	}

	return v.Fn(args...)
}
func applyFunction(v *object.Function, args []object.Object, env *object.Environment) object.Object {
	return Eval(*v.Body, env)
}
func createError(message string, v ...interface{}) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf("ERROR: %s\n", fmt.Sprintf(message, v...)),
	}

}
