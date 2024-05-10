package eval

import (
	"os"

	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func resolveImport(name string) string {
	return name
}

func resolveEnvironment(file string, errorHandler func(string)) *object.Environment {

	env := object.NewEnvironment()

	contents, err := os.ReadFile(file)

	if err != nil {
		errorHandler("Error reading file: " + file)
	}

	l := lexer.NewLexer(string(contents))
	p := parser.NewParser(l, false)

	program := p.Parse()

	if len(p.Errors()) > 0 {
		errorHandler("Error parsing file: " + file)
	}

	e := NewEvaluator(errorHandler)

	e.Eval(program, env)

	return env
}
func applyEnvironment(parent, child *object.Environment) {
	for child.Outer != nil {
		child = child.Outer
	}

	child.Outer = parent
}
