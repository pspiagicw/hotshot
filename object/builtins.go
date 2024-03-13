package object

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/ast"
)

func getBuiltins() map[string]*Builtin {
	return map[string]*Builtin{
		"+": &Builtin{
			Fn: addFunc,
		},
		"-": &Builtin{
			Fn: minusFunc,
		},
		"*": &Builtin{
			Fn: multiplyFunc,
		},
		"/": &Builtin{
			Fn: divideFunc,
		},
		"=": &Builtin{
			Fn: equalFunc,
		},
		"<": &Builtin{
			Fn: ltFunc,
		},
		">": &Builtin{
			Fn: gtFunc,
		},
		"not": &Builtin{
			Fn: notFunc,
		},
		"and": &Builtin{
			Fn: andFunc,
		},
		"or": &Builtin{
			Fn: orFunc,
		},
		"do": &Builtin{
			Fn: doFunc,
		},
		"echo": &Builtin{
			Fn: printFunc,
		},
		"len": &Builtin{
			Fn: lenFunc,
		},
		"mod": &Builtin{
			Fn: modFunc,
		},
		"let": &Builtin{
			Fn: letFunc,
		},
	}
}
func letFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) != 2 {
		return createError("LET function expects 2 arguments.")
	}

	nameIdent, ok := args[0].(*ast.IdentStatement)
	if !ok {
		return createError("LET expects IDENT to be first argument!")
	}

	value := evalFunc(args[1])

	env.Set(nameIdent.Value.TokenValue, value)

	return Null{}
}
func doFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	length := len(args)

	evals := []Object{}

	for _, arg := range args {
		evals = append(evals, evalFunc(arg))
	}

	return evals[length-1]
}
func lenFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	if len(args) == 0 {
		return createError("No arguments given to LEN function!")
	}

	if len(args) > 1 {
		return createError("LEN function only accepts 1 argument!")
	}

	value := evalFunc(args[0])
	switch v := value.(type) {
	case *String:
		return &Integer{
			Value: len(v.Value),
		}
	}

	return createError("LEN function can't find length of that type!")
}
func printFunc(args []ast.Statement, evalFunc func(ast.Statement) Object, env *Environment) Object {
	var output strings.Builder

	output.WriteString(" ")

	for _, arg := range args {
		output.WriteString(evalFunc(arg).String())
		output.WriteString(" ")
	}

	fmt.Println(output.String())

	return Null{}
}

func createError(format string, v ...interface{}) Object {
	return &Error{
		Message: fmt.Sprintf(format, v...),
	}
}
