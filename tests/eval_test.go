package tests

import (
	"reflect"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func TestEval(t *testing.T) {

	// input := "()"
	tt := map[string]object.Object{
		"()":                           createNull(),
		"; this is a simple comment ;": createNull(),

		"1":               createInt(1),
		"23234":           createInt(23234),
		"(+ 1 2)":         createInt(3),
		"(+ 1 (+ 1 2) 2)": createInt(6),
		"(- 2 1 )":        createInt(1),
		"(* 5 6)":         createInt(30),
		"(+ 5 (- 4 5))":   createInt(4),
		"(+ -5 (* 4 5))":  createInt(15),
		// "(% 4 3)":             createInt(1),
		`($ number 5) number`: createInt(5),

		`"someString"`:            createString("someString"),
		`($ name "hotshot") name`: createString("hotshot"),

		`(= 1 1)`:           createBool(true),
		`(= 1 2)`:           createBool(false),
		`(= "some" "some")`: createBool(true),
		`(= "some" "else")`: createBool(false),

		`(= true true)`:   createBool(true),
		`(= false true)`:  createBool(false),
		`(= false false)`: createBool(true),

		`(> 1 1)`: createBool(false),
		`(< 1 1)`: createBool(false),
		`(< 2 1)`: createBool(false),
		`(> 2 1)`: createBool(true),

		`(not true)`:        createBool(false),
		`(not false)`:       createBool(true),
		`(and false true)`:  createBool(false),
		`(and true true)`:   createBool(true),
		`(and false false)`: createBool(false),
		`(or true false)`:   createBool(true),
		`(or false false)`:  createBool(false),
		`(or true true)`:    createBool(true),

		`(if true 5 2)`:    createInt(5),
		`(if false 5 2)`:   createInt(2),
		`(if (= 1 2) 5 2)`: createInt(2),
		`(if (= 1 1) 7 2)`: createInt(7),

		`(fn hello () "Hello, World")`:         createNull(),
		`(fn hello () "Hello, World") (hello)`: createString("Hello, World"),
		`(fn add (x y) (+ x y)) (add 2 1)`:     createInt(3),
	}

	for input, expectedResult := range tt {
		t.Run(input, func(t *testing.T) {
			checkResult(t, input, expectedResult)
		})
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
func createBool(val bool) *object.Boolean {
	return &object.Boolean{
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

	snaps.MatchSnapshot(t, ast)

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
