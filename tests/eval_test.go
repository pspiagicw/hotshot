package tests

import (
	"reflect"
	"testing"

	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func TestEvalStatements(t *testing.T) {

	// input := "()"
	tt := map[string]object.Object{
		"()":                           createNull(),
		"; this is a simple comment ;": createNull(),
		"1":                            createInt(1),
		"23234":                        createInt(23234),
		`"someString"`:                 createString("someString"),
		"(+ 1 2)":                      createInt(3),
		"(+ 1 (+ 1 2) 2)":              createInt(6),
		"(- 2 1 )":                     createInt(1),
		"(* 5 6)":                      createInt(30),
		"(+ 5 (- 4 5))":                createInt(4),
		"(+ -5 (* 4 5))":               createInt(15),
		`($ number 5) number`:          createInt(5),
		`($ name "hotshot") name`:      createString("hotshot"),
	}

	for input, expectedResult := range tt {
		checkResult(t, input, expectedResult)
	}

}
func createNull() *object.Null {
	return &object.Null{}
}
func createString(val string) *object.String {
	return &object.String{
		Value: val,
	}
}
func createInt(val int) *object.Integer {
	return &object.Integer{
		Value: val,
	}
}
func checkResult(t *testing.T, input string, expected object.Object) {

	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)

	ast := parser.Parse()
	if len(parser.Errors()) != 0 {
		t.Fatalf("Errors while parsing")
	}

	// TODO: Last statement is always nil!
	ast.Statements = ast.Statements[:len(ast.Statements)-1]

	env := object.NewEnvironment()

	result := eval.Eval(ast, env)

	equalResult(t, result, expected)

}

func equalResult(t *testing.T, result, expected object.Object) {

	t.Helper()

	if result == nil {
		t.Fatalf("Evaluted result is nil!")
	}

	if result.Type() == object.NULL_OBJ && expected.Type() != object.NULL_OBJ {
		t.Errorf("Evaluated result not equal! got %v, expected Null", result)
	}

	if result.Type() != expected.Type() {
		t.Errorf("Evaluated result type not equal! got %v, expected %v", result.Type(), expected.Type())

	}

	if result.Type() != object.NULL_OBJ {
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Evaluated result not equal! got %v, expected %v", result, expected)
		}
	}
}
