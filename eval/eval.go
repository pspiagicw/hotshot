package eval

import (
	"fmt"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/object"
)

type Evaluator struct {
	ErrorHandler func(message string)
}

func NewEvaluator(handler func(string)) *Evaluator {
	return &Evaluator{
		ErrorHandler: handler,
	}
}

func (e *Evaluator) Eval(node ast.Statement, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IntStatement:
		return &object.Integer{Value: node.Value}
	case *ast.StringStatement:
		return &object.String{Value: node.Value}
	case *ast.Program:
		return e.evalProgram(node, env)
	case *ast.CallStatement:
		return e.evalFunction(node, env)
	case *ast.EmptyStatement:
		return object.Null{}
	case *ast.AssignmentStatement:
		return e.applyAssignment(node, env)
	case *ast.BoolStatement:
		return &object.Boolean{Value: node.Value}
	case *ast.IdentStatement:
		return evalIdent(node, env)
	case *ast.IfStatement:
		return e.evalIfStatement(node, env)
	case *ast.WhileStatement:
		return e.evalWhileStatement(node, env)
	case *ast.FunctionStatement:
		return e.evalFunctionStatement(node, env)
	case *ast.LambdaStatement:
		return e.evalLambdaStatement(node, env)
	case *ast.TableStatement:
		return e.evalTableStatement(node, env)
	case *ast.CondStatement:
		return e.evalCondStatement(node, env)
	case *ast.AssertStatement:
		return e.evalAssertStatement(node, env)
	}
	return e.createError("Evaluation for statement can't be done!")
}

func (e *Evaluator) evalAssertStatement(node *ast.AssertStatement, env *object.Environment) object.Object {
	left := e.Eval(node.Left, env)
	right := e.Eval(node.Right, env)

	if left.Type() != right.Type() || left.String() != right.String() {
		return e.createError("Assertion Error: %s", node.Message.TokenValue)
	}

	return left
}
func (e *Evaluator) evalCondStatement(node *ast.CondStatement, env *object.Environment) object.Object {

	for _, exp := range node.Expressions {
		if e.isTrue(e.Eval(exp.Condition, env)) {
			return e.Eval(exp.Body, env)
		}
	}

	return object.Null{}
}
func (e *Evaluator) evalTableStatement(node *ast.TableStatement, env *object.Environment) object.Object {
	fn := &object.Table{}

	fn.Elements = e.evalStatements(node.Elements, env)

	return fn
}
func (e *Evaluator) evalLambdaStatement(node *ast.LambdaStatement, env *object.Environment) object.Object {
	fn := &object.Function{}

	args := []*ast.IdentStatement{}

	for _, arg := range node.Args {
		v, ok := arg.(*ast.IdentStatement)
		if !ok {
			return e.createError("Argument is not a identifier!")
		}
		args = append(args, v)
	}

	fn.Args = args

	fn.Body = &node.Body

	return fn
}

func (e *Evaluator) evalFunctionStatement(node *ast.FunctionStatement, env *object.Environment) object.Object {
	name := node.Name.TokenValue

	fn := &object.Function{}

	args := []*ast.IdentStatement{}

	for _, arg := range node.Args {
		v, ok := arg.(*ast.IdentStatement)
		if !ok {
			return e.createError("Argument is not a identifier!")
		}
		args = append(args, v)
	}

	fn.Args = args

	fn.Body = &node.Body

	env.Set(name, fn)

	return object.Null{}
}
func (e Evaluator) evalWhileStatement(node *ast.WhileStatement, env *object.Environment) object.Object {
	var result object.Object

	result = object.Null{}

	for true {
		result := e.Eval(node.Condition, env)

		if e.isTrue(result) {
			result = e.Eval(node.Body, env)
		} else {
			break
		}
	}
	return result
}
func (e *Evaluator) evalIfStatement(node *ast.IfStatement, env *object.Environment) object.Object {

	result := e.Eval(node.Condition, env)

	if e.isTrue(result) {
		return e.Eval(node.Body, env)
	} else {
		if node.Else != nil {
			return e.Eval(node.Else, env)
		}
	}

	return object.Null{}

}
func (e *Evaluator) isTrue(value object.Object) bool {
	if value.Type() != object.BOOLEAN_OBJ {
		e.createError("Condition doesn't evaluate to true/false!")
	}

	return value.String() == "true"
}
func evalIdent(node *ast.IdentStatement, env *object.Environment) object.Object {
	val := env.Get(node.Value.TokenValue)

	return val
}
func (e *Evaluator) evalProgram(statements *ast.Program, env *object.Environment) object.Object {
	results := e.evalStatements(statements.Statements, env)

	if len(results) == 0 {
		return object.Null{}
	}

	return results[len(results)-1]
}

func (e *Evaluator) applyAssignment(node *ast.AssignmentStatement, env *object.Environment) object.Object {

	value := e.Eval(node.Value, env)

	env.Set(node.Name.TokenValue, value)

	return object.Null{}
}
func (e *Evaluator) evalStatements(args []ast.Statement, env *object.Environment) []object.Object {
	evals := []object.Object{}

	for _, arg := range args {
		evals = append(evals, e.Eval(arg, env))
	}
	return evals
}
func (e *Evaluator) evalUserFunc(node *ast.CallStatement, env *object.Environment) object.Object {
	fn := env.Get(node.Op.TokenValue)

	if fn.Type() != object.FUNCTION_OBJ {
		return e.createError("No function named '%s'", node.Op.TokenValue)
	}

	v, ok := fn.(*object.Function)
	if !ok {
		return e.createError("INTERNAL: Couldn't initialize user function!")
	}

	evals := e.evalStatements(node.Args, env)

	return e.applyFunction(v, evals, env)
}
func (e *Evaluator) evalFunction(node *ast.CallStatement, env *object.Environment) object.Object {

	fn := env.GetBuiltin(node.Op.TokenValue)

	if fn == nil {
		return e.evalUserFunc(node, env)
	}

	evals := e.evalStatements(node.Args, env)

	v, ok := fn.(*object.Builtin)

	if !ok {
		return e.createError("INTERNAL: Could't initalize builtin function")
	}

	value := v.Fn(evals)

	if value.Type() == object.ERROR_OBJ {
		v, ok := value.(*object.Error)
		if !ok {
			return e.createError("INTERNAL: Couldn't cast error value")

		}
		e.createError(v.Message)
	}

	return value

}
func (e *Evaluator) applyFunction(v *object.Function, args []object.Object, env *object.Environment) object.Object {
	if len(v.Args) != len(args) {
		return e.createError("Function expects %d argument, given %d", len(v.Args), len(args))
	}

	newEnv := extendEnvironment(v.Args, args, env)
	return e.Eval(*v.Body, newEnv)
}
func extendEnvironment(declaredArgs []*ast.IdentStatement, givenArgs []object.Object, env *object.Environment) *object.Environment {
	newEnv := object.NewEnvironment()
	newEnv.Outer = env
	for i, darg := range declaredArgs {
		newEnv.Set(darg.Value.TokenValue, givenArgs[i])
	}
	return newEnv
}
func (e *Evaluator) createError(message string, v ...interface{}) *object.Error {
	message = fmt.Sprintf("ERROR: %s\n", fmt.Sprintf(message, v...))
	e.ErrorHandler(message)
	return &object.Error{
		Message: message,
	}

}
