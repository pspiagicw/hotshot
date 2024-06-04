package eval

import (
	"embed"
	"os"

	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

//go:embed stdlib
var STD_LIB embed.FS

func (e *Evaluator) getImportPath(name string) string {
	return name
}
func readStdFile(name string) string {
	contents, err := STD_LIB.ReadFile(name)

	if err != nil {
		panic("Error reading file: " + name)
	}

	return string(contents)
}
func readFile(name string) string {
	contents, err := os.ReadFile(name)

	if err != nil {
		panic("Error reading file: " + name)
	}

	return string(contents)
}

func (e *Evaluator) getImportContent(name string) string {
	switch name {
	case "math":
		return readStdFile("stdlib/math.ht")
	default:
		return readFile(name)
	}
}

func (e *Evaluator) resolveImport(contents string, env *object.Environment) *object.Environment {

	l := lexer.NewLexer(contents)
	p := parser.NewParser(l, false)

	program := p.Parse()

	if len(p.Errors()) > 0 {
		e.ErrorHandler("Error parsing file! ")
		return nil
	}

	newEval := NewEvaluator(e.ErrorHandler)

	newEval.Eval(program, env)

	return env
}
func applyEnvironment(parent, child *object.Environment) {
	for child.Outer != nil {
		child = child.Outer
	}

	child.Outer = parent
}
