package vm

import (
	"fmt"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

type vmTestCase struct {
	input  string
	result interface{}
}

func TestVM(t *testing.T) {
	tt := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"(+ 1 2)", 3},
		{"(- 2 1)", 1},
		{"(* 2 2)", 4},
		{"(/ 4 2)", 2},
		{"(- (+ 10 (* 2 (/ 50 2))) 5)", 55},
		{"(- (+ 5 5 5 5) 10)", 10},
		{"(* 2 2 2 2 2)", 32},
		{"(+ 10 (* 5 2))", 20},
		{"(+ 5 (* 2 10))", 25},
		{"(* 5 (+ 2 10))", 60},
		{"true", true},
		{"false", false},
		{"(< 1 2)", true},
		{"(> 1 2)", false},
		{"(> 1 1)", false},
		{"(< 1 1)", false},
		{"(= 1 1)", true},
		{"(= 1 2)", false},
		{"(= 1 2)", false},
		{"(= true true)", true},
		{"(= false true)", false},
		{"(= (< 1 2) true)", true},
		{"(= (< 1 2) false)", false},
		{"(= (> 1 2) true)", false},
		{"(= (> 1 2) false)", true},
		{"(if true 10)", 10},
		{"(if true 10 20)", 10},
		{"(if false 10 20)", 20},
		{"(if (< 1 2) 10)", 10},
		{"(if (< 1 2) 10 20)", 10},
		{"(if (> 1 2) 10 20)", 20},
		{"(let one 1) one", 1},
		{"(let one 1) (let two 2) (+ one two)", 3},
		{"(let one 1) (let two (+ one one)) (+ one two)", 3},
		{"(let fivePlusTen (lambda () (+ 5 10))) (fivePlusTen)", 15},
		{"(let one (lambda () 1)) (let two (lambda () 2)) (+ (one) (two))", 3},
		{`(let returnsOne (lambda() 1))
            (let returnsOneReturner (lambda () (returnsOne)))
            (returnsOneReturner)
            `,
			1,
		},
		// {`(let one (lambda() (let one 1) one)) (one)`, 1},
	}
	for _, test := range tt {
		t.Run(test.input, func(t *testing.T) {
			checkResult(t, test)
		})
	}
}
func checkResult(t *testing.T, test vmTestCase) {
	t.Helper()

	lexer := lexer.NewLexer(test.input)
	parser := parser.NewParser(lexer, false)

	ast := parser.Parse()
	if len(parser.Errors()) != 0 {
		for _, err := range parser.Errors() {
			t.Errorf("Error: %s", err)
		}
		t.Fatalf("Errors while parsing")
	}
	snaps.MatchSnapshot(t, ast)

	c := compiler.NewCompiler()

	err := c.Compile(ast)

	if err != nil {
		t.Fatalf("Error while compiling: %s", err)
	}

	bytecode := c.Bytecode()

	bytecode = compiler.JumpPass(bytecode)

	vm := NewVM(bytecode)

	status := vm.Run()

	if status != nil {
		t.Errorf("Error: %s", status)
		t.Fatalf("Error while running the VM")
	}

	result := vm.StackTop()
	if result == nil {
		t.Fatalf("Executed result is nil!")
	}
	checkEqual(t, result, test.result)
}
func checkEqual(t *testing.T, result object.Object, expected interface{}) {
	t.Helper()
	var err error
	switch expected := expected.(type) {
	case int:
		err = testIntegerObject(t, result, expected)
	case string:
		err = testStringObject(t, result, expected)
	case bool:
		err = testBooleanObject(t, result, expected)
	default:
		t.Fatalf("Type testing not supported")
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
}
func testIntegerObject(t *testing.T, result object.Object, expected int) error {
	t.Helper()
	if result.Type() != object.INTEGER_OBJ {
		return fmt.Errorf("Expected INTEGER_OBJ, got %s", result.Type())
	}
	if result.(*object.Integer).Value != expected {
		return fmt.Errorf("Expected %d, got %d", expected, result.(*object.Integer).Value)
	}
	return nil
}
func testStringObject(t *testing.T, result object.Object, expected string) error {
	t.Helper()
	if result.Type() != object.STRING_OBJ {
		return fmt.Errorf("Expected STRING_OBJ, got %s", result.Type())
	}
	if result.(*object.String).Value != expected {
		return fmt.Errorf("Expected %s, got %s", expected, result.(*object.String).Value)
	}
	return nil
}
func testBooleanObject(t *testing.T, result object.Object, expected bool) error {
	t.Helper()
	if result.Type() != object.BOOLEAN_OBJ {
		return fmt.Errorf("Expected BOOLEAN_OBJ, got %s", result.Type())
	}
	if result.(*object.Boolean).Value != expected {
		return fmt.Errorf("Expected %t, got %t", expected, result.(*object.Boolean).Value)
	}
	return nil
}
