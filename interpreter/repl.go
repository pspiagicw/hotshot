package interpreter

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	for true {
		fmt.Printf(">>> ")

		prompt := getPrompt()

		program, errors := parseCode(prompt)

		if handleErrors(errors, false) != 0 {
			continue
		}

		// A extra empty statement is added by the parser at the end.
		program.Statements = program.Statements[:len(program.Statements)-1]

		fmt.Println(printer.PrintAST(program))

		result := eval.Eval(program, env)
		fmt.Print("=> ")
		fmt.Print(result)
		fmt.Println()
	}
}
func getPrompt() string {
	buffer := bufio.NewReader(os.Stdin)
	prompt, err := buffer.ReadString('\n')

	if err != nil {
		goreland.LogFatal("Error scanning input: %v", err)
	}

	prompt = strings.TrimSpace(prompt)

	return prompt
}

func ExecuteFile(file string, debug bool) {

	code := readFile(file)

	program, errors := parseCode(code)

	handleErrors(errors, false)

	if debug {
		fmt.Println(printer.PrintAST(program))
	}

	env := object.NewEnvironment()

	_ = eval.Eval(program, env)
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
