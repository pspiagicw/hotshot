package interpreter

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
	"github.com/pspiagicw/hotshot/printer"
)

func StartREPL() {
	env := object.NewEnvironment()

	errorHandler := func(message string) {
		goreland.LogError("Runtime Error: %s\n", message)
	}

	e := eval.NewEvaluator(errorHandler)
	for true {
		prompt := getInput(">>> ")

		program, errors := parseCode(prompt)

		if handleErrors(errors, false) != 0 {
			continue
		}

		// A extra empty statement is added by the parser at the end.
		program.Statements = program.Statements[:len(program.Statements)-1]

		fmt.Println(printer.PrintAST(program))

		result := e.Eval(program, env)
		fmt.Print("=> ")
		fmt.Print(result)
		fmt.Println()
	}
}
func getInput(prompt string) string {
	input, err := readline.Line(prompt)
	if err != nil {
		goreland.LogFatal("Error scanning input: %v", err)
	}
	return input

}

func handleErrors(errors []error, exit bool) int {
	if len(errors) != 0 {
		goreland.LogError("Error parsing the file.")
		for _, err := range errors {
			fmt.Println(err.Error())
		}
		if exit {
			goreland.LogFatal("Not executing till errors resolved!")
		} else {
			return len(errors)
		}
	}
	return 0
}

func parseCode(code string) (*ast.Program, []error) {
	l := lexer.NewLexer(code)
	p := parser.NewParser(l)

	program := p.Parse()
	return program, p.Errors()
}
