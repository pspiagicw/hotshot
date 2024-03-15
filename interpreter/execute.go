package interpreter

import (
	"fmt"
	"os"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/printer"
)

func ExecuteFile(file string, debug bool) {

	code := readFile(file)

	program, errors := parseCode(code)

	errorHandler := func(message string) {
		goreland.LogFatal("Runtime Error: %s", message)
	}

	e := eval.NewEvaluator(errorHandler)

	handleErrors(errors, true)

	if debug {
		fmt.Println(printer.PrintAST(program))
	}

	env := object.NewEnvironment()

	_ = e.Eval(program, env)
}
func readFile(file string) string {

	contents, err := os.ReadFile(file)
	if err != nil {
		goreland.LogFatal("Error while reading file '%s'", file)
	}

	return string(contents)
}
