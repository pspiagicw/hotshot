package interpreter

import (
	"fmt"
	"os"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/printer"
)

func ExecuteFile(file string, opts *argparse.Opts) {

	code := readFile(file)

	program, errors := parseCode(code, opts)

	errorHandler := func(message string) {
		goreland.LogFatal("Runtime Error: %s", message)
	}

	e := eval.NewEvaluator(errorHandler)

	handleErrors(errors, true)

	if opts.AST {
		fmt.Println(printer.PrintAST(program))
	}

	env := object.NewEnvironment()

	_ = e.Eval(program, env)

	// c := compiler.NewCompiler()
	// err := c.Compile(program)
	// if err != nil {
	// 	goreland.LogFatal("Error while compiling program: %v", err)
	// }
	// decompiler.Print(c.Bytecode())
}
func readFile(file string) string {

	contents, err := os.ReadFile(file)
	if err != nil {
		goreland.LogFatal("Error while reading file '%s'", file)
	}

	return string(contents)
}
