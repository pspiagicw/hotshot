package interpreter

import (
	"fmt"
	"os"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/printer"
	"github.com/pspiagicw/hotshot/vm"
)

func ExecuteFile(file string, opts *argparse.Opts) {

	code := readFile(file)

	program, errors := parseCode(code, opts)

	// errorHandler := func(message string) {
	// 	goreland.LogFatal("Runtime Error: %s", message)
	// }

	// e := eval.NewEvaluator(errorHandler)

	handleErrors(errors, true)

	if opts.AST {
		fmt.Println(printer.PrintAST(program))
	}

	// env := object.NewEnvironment()

	// _ = e.Eval(program, env)

	c := compiler.NewCompiler()
	err := c.Compile(program)
	if err != nil {
		goreland.LogFatal("Error while compiling program: %v", err)
	}
	bytecode := c.Bytecode()
	bytecode = compiler.JumpPass(bytecode)

	// decompiler.Print(bytecode)
	vm := vm.NewVM(bytecode)

	error := vm.Run()

	if error != nil {
		goreland.LogFatal("Error while running the VM: %v", error)
	}
	// fmt.Println(vm.StackTop())
}
func readFile(file string) string {

	contents, err := os.ReadFile(file)
	if err != nil {
		goreland.LogFatal("Error while reading file '%s'", file)
	}

	return string(contents)
}
