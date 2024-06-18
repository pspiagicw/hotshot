package interpreter

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/decompiler"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
	"github.com/pspiagicw/regolith"
)

func printHeader() {
	fmt.Println("Welcome to hotshot!")
	fmt.Println("Use `(exit)` to exit the REPL")
}
func getPrompt() string {
	prompt := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render(">>> ")
	return prompt
}

func StartREPL(opts *argparse.Opts) {

	printHeader()

	env := object.NewEnvironment()
	e := eval.NewEvaluator(func(message string) {
		goreland.LogError("%s", message)
	})

	rg, err := regolith.New(&regolith.Config{
		StartWords: []string{"(", "["},
		EndWords:   []string{")", "]"},
	})

	if err != nil {
		goreland.LogFatal("Error initializing regolith: %v", err)
	}

	symbols := compiler.NewSymbolTable()
	constants := []object.Object{}

	for {
		input, err := rg.Input()

		if err != nil {
			goreland.LogFatal("Error reading input from prompt: %v", err)
		}

		program, errors := parseCode(input, opts)
		if handleErrors(errors, false) != 0 {
			continue
		}
		if opts.AST {
			fmt.Println(program.String())
		}
		result := e.Eval(program, env)
		if opts.Null || result.Type() != object.NULL_OBJ {
			fmt.Print("=> ")
			fmt.Println(result.String())
		}
		c := compiler.NewWithState(symbols, constants)
		err = c.Compile(program)

		if err != nil {
			goreland.LogError("Error compiling the program: %v", err)
		}
		decompiler.Print(c.Bytecode())
	}

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

func parseCode(code string, opts *argparse.Opts) (*ast.Program, []error) {
	l := lexer.NewLexer(code)
	p := parser.NewParser(l, opts.Token)

	program := p.Parse()
	program.Statements = program.Statements[:len(program.Statements)-1]
	return program, p.Errors()
}
