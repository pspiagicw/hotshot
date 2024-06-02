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
	if name == "math" {
		return "math"
	}
	return name
}

func (e *Evaluator) getImportContent(name string) string {
	if name == "math" {
		content, err := STD_LIB.ReadFile("stdlib/math.ht")
		if err != nil {
			e.ErrorHandler("Error reading file: , " + name)
		}
		return string(content)
	}
	contents, err := os.ReadFile(name)

	if err != nil {
		e.ErrorHandler("Error reading file: " + name)
	}

	return string(contents)
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
