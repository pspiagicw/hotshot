package interpreter

import (
	"fmt"
	"os"
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
func completeInput(input string) bool {
	return strings.Count(input, "(") == strings.Count(input, ")")
}

func ExecuteFile(file string, debug bool) {

	code := readFile(file)

	program, errors := parseCode(code)

	errorHandler := func(message string) {
		goreland.LogFatal("Runtime Error: %s", message)
	}

	e := eval.NewEvaluator(errorHandler)

	handleErrors(errors, false)

	if debug {
		fmt.Println(printer.PrintAST(program))
	}

	env := object.NewEnvironment()

	_ = e.Eval(program, env)
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

func readFile(file string) string {

	contents, err := os.ReadFile(file)
	if err != nil {
		goreland.LogFatal("Error while reading file '%s'", file)
	}

	return string(contents)
}
