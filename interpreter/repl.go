package interpreter

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/hotshot/argparse"
	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func printHeader() {
	fmt.Println("Welcome to hotshot!")
	fmt.Println("Use `(exit)` to exit the REPL")
}

func StartREPL(opts *argparse.Opts) {

	printHeader()

	env := object.NewEnvironment()
	e := eval.NewEvaluator(func(message string) {
		goreland.LogError("Runtime Error: %s\n", message)
	})

	rl := initREPL()
	defer rl.Close()

	for {
		input := getInput(rl)

		program, errors := parseCode(input, opts)
		handleErrors(errors, false)
		result := e.Eval(program, env)
		if opts.Null || result.Type() != object.NULL_OBJ {
			fmt.Print("=> ")
			fmt.Print(result.String())
			fmt.Println()
		}
	}

}

func initREPL() *readline.Instance {
	r, err := readline.NewEx(&readline.Config{
		Prompt:          ">>> ",
		HistoryFile:     "/tmp/readline.tmp",
		InterruptPrompt: "^D",
	})

	if err != nil {
		goreland.LogFatal("Error initalizing readline: %v", err)
	}

	return r
}

func getInput(r *readline.Instance) string {
	var cmds []string

	parenScore := 0
	for {
		line, err := r.Readline()
		if err != nil {
			goreland.LogFatal("Error reading input from prompt: %v", err)
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cmds = append(cmds, line)
		parenScore += strings.Count(line, "(") - strings.Count(line, ")")
		if parenScore == 0 {
			break
		}
		r.SetPrompt("... ")
	}
	r.SetPrompt(">>> ")

	return strings.Join(cmds, "\n")
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
