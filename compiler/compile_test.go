package compiler

import (
	"testing"

	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func TestLocalScopes(t *testing.T) {
	input := `(lambda () (let num 55) num)`

	constants := []interface{}{
		55,
		[]*code.Instruction{
			{OpCode: code.PUSH, Args: 0},
			{OpCode: code.LSET, Args: 0},
			{OpCode: code.LGET, Args: 0},
		},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 1},
	}

	checkBytecode(t, input, bytecode, constants)
}
func TestLocalStrict(t *testing.T) {
	input := `(lambda () (let a 55) ( let b 77) (+ a b))`

	constants := []interface{}{
		55,
		77,
		[]*code.Instruction{
			{OpCode: code.PUSH, Args: 0},
			{OpCode: code.LSET, Args: 0},
			{OpCode: code.PUSH, Args: 1},
			{OpCode: code.LSET, Args: 1},
			{OpCode: code.LGET, Args: 0},
			{OpCode: code.LGET, Args: 1},
			{OpCode: code.ADD, Args: 2},
		},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 2},
	}

	checkBytecode(t, input, bytecode, constants)
}
func TestLocals(t *testing.T) {
	input := `(let num 55) (lambda () num)`

	constants := []interface{}{
		55,
		[]*code.Instruction{
			{OpCode: code.GET, Args: 0},
		},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.SET, Args: 0},
		{OpCode: code.PUSH, Args: 1},
	}

	checkBytecode(t, input, bytecode, constants)
}

func TestCall(t *testing.T) {
	input := `(fn someFunc ()) (someFunc)`

	constants := []interface{}{
		[]*code.Instruction{},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.SET, Args: 0},
		{OpCode: code.GET, Args: 0},
		{OpCode: code.CALL, Args: -1},
	}

	checkBytecode(t, input, bytecode, constants)
}
func TestFunctionDec(t *testing.T) {
	input := `(fn someFunc () (let a 25) a) (someFunc)`

	contants := []interface{}{
		25,
		[]*code.Instruction{
			{OpCode: code.PUSH, Args: 0},
			{OpCode: code.LSET, Args: 0},
			{OpCode: code.LGET, Args: 0},
		},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.SET, Args: 0},
		{OpCode: code.GET, Args: 0},
		{OpCode: code.CALL, Args: -1},
	}

	checkBytecode(t, input, bytecode, contants)
}

func TestLambda(t *testing.T) {
	input := `(lambda () 25)`

	constants := []interface{}{
		25,
		[]*code.Instruction{
			{OpCode: code.PUSH, Args: 0},
		},
	}
	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 1},
	}

	checkBytecode(t, input, bytecode, constants)
}

func TestLambdaAssignment(t *testing.T) {
	input := `(let a (lambda () 25)) (a)`

	constants := []interface{}{
		25,
		[]*code.Instruction{
			{OpCode: code.PUSH, Args: 0}, // The 25
		},
	}

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 1},  // The compiled function
		{OpCode: code.SET, Args: 0},   // The variable
		{OpCode: code.GET, Args: 0},   // The variable
		{OpCode: code.CALL, Args: -1}, // The call
	}

	checkBytecode(t, input, bytecode, constants)
}

func TestCondStatements(t *testing.T) {
	input := `
    (cond ((> 10 20) 10) ((= 1 2) 2310) (true false))
    `

	constants := []interface{}{10, 20, 10, 1, 2, 2310}
	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.GT, Args: 2},
		{OpCode: code.JCMP, Args: 2},
		{OpCode: code.PUSH, Args: 2},
		{OpCode: code.JMP, Args: 1},
		{OpCode: code.JT, Args: 2},
		{OpCode: code.PUSH, Args: 3},
		{OpCode: code.PUSH, Args: 4},
		{OpCode: code.EQ, Args: 2},
		{OpCode: code.JCMP, Args: 3},
		{OpCode: code.PUSH, Args: 5},
		{OpCode: code.JMP, Args: 1},
		{OpCode: code.JT, Args: 3},
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.JCMP, Args: 4},
		{OpCode: code.FALSE, Args: -1},
		{OpCode: code.JMP, Args: 1},
		{OpCode: code.JT, Args: 4},
		{OpCode: code.JT, Args: 1},
	}

	checkBytecode(t, input, bytecode, constants)
}
func TestWhileStatements(t *testing.T) {
	input := `(while (< 10 20) 10)`

	bytecode := []*code.Instruction{
		{OpCode: code.JT, Args: 1},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.LT, Args: 2},
		{OpCode: code.JCMP, Args: 2},
		{OpCode: code.PUSH, Args: 2},
		{OpCode: code.JMP, Args: 1},
		{OpCode: code.JT, Args: 2},
	}
	constants := []interface{}{10, 20, 10}

	checkBytecode(t, input, bytecode, constants)
}

func TestLetStatement(t *testing.T) {
	input := `(let a 3) (let b 4) (+ a b)`

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.SET, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.SET, Args: 1},
		{OpCode: code.GET, Args: 0},
		{OpCode: code.GET, Args: 1},
		{OpCode: code.ADD, Args: 2},
	}
	constants := []interface{}{3, 4}

	checkBytecode(t, input, bytecode, constants)
}

func TestIfElse(t *testing.T) {
	input := `(if true 10 20)`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.JCMP, Args: 1},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.JMP, Args: 2},
		{OpCode: code.JT, Args: 1},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.JT, Args: 2},
	}
	constants := []interface{}{10, 20}

	checkBytecode(t, input, bytecode, constants)
}

func TestIf(t *testing.T) {
	input := `(if true 10)`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.JCMP, Args: 1},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.JT, Args: 1},
	}
	constants := []interface{}{10}

	checkBytecode(t, input, bytecode, constants)
}

func TestComparison(t *testing.T) {
	input := `(> 1 2) (< 1 2) (= 1 2)`

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.GT, Args: 2},
		{OpCode: code.PUSH, Args: 2},
		{OpCode: code.PUSH, Args: 3},
		{OpCode: code.LT, Args: 2},
		{OpCode: code.PUSH, Args: 4},
		{OpCode: code.PUSH, Args: 5},
		{OpCode: code.EQ, Args: 2},
	}
	constants := []interface{}{1, 2, 1, 2, 1, 2}
	checkBytecode(t, input, bytecode, constants)
}
func TestBoolean(t *testing.T) {
	input := `true false`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.FALSE, Args: -1},
	}
	constants := []interface{}{}

	checkBytecode(t, input, bytecode, constants)

}

func TestArithmetic(t *testing.T) {
	input := `(+ 1 2 (- 3 4) (* 5 (/ 6 3)))`

	// 00000 PUSH 0
	// 00001 PUSH 1
	// 00002 PUSH 2
	// 00003 PUSH 3
	// 00004 SUB 2
	// 00005 PUSH 4
	// 00006 PUSH 5
	// 00007 PUSH 6
	// 00008 DIV 2
	// 00009 MUL 2
	// 00010 ADD 4
	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.PUSH, Args: 2},
		{OpCode: code.PUSH, Args: 3},
		{OpCode: code.SUB, Args: 2},
		{OpCode: code.PUSH, Args: 4},
		{OpCode: code.PUSH, Args: 5},
		{OpCode: code.PUSH, Args: 6},
		{OpCode: code.DIV, Args: 2},
		{OpCode: code.MUL, Args: 2},
		{OpCode: code.ADD, Args: 4},
	}

	constants := []interface{}{1, 2, 3, 4, 5, 6, 3}
	checkBytecode(t, input, bytecode, constants)
}

func TestPush(t *testing.T) {
	input := `1`

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
	}

	constants := []interface{}{1}
	checkBytecode(t, input, bytecode, constants)
}
func TestAdd(t *testing.T) {
	input := `(+ 1 2)`

	bytecode := []*code.Instruction{
		{
			OpCode: code.PUSH,
			Args:   0,
		},
		{
			OpCode: code.PUSH,
			Args:   1,
		},
		{
			OpCode: code.ADD,
			Args:   2,
		},
	}

	constants := []interface{}{1, 2}
	checkBytecode(t, input, bytecode, constants)
}
func TestSubtract(t *testing.T) {
	input := `(- 1 2)`

	bytecode := []*code.Instruction{
		{
			OpCode: code.PUSH,
			Args:   0,
		},
		{
			OpCode: code.PUSH,
			Args:   1,
		},
		{
			OpCode: code.SUB,
			Args:   2,
		},
	}
	constants := []interface{}{1, 2}

	checkBytecode(t, input, bytecode, constants)
}
func TestMultiply(t *testing.T) {
	input := `(* 1 2)`

	bytecode := []*code.Instruction{
		{
			OpCode: code.PUSH,
			Args:   0,
		},
		{
			OpCode: code.PUSH,
			Args:   1,
		},
		{
			OpCode: code.MUL,
			Args:   2,
		},
	}
	constants := []interface{}{1, 2}

	checkBytecode(t, input, bytecode, constants)

}
func TestDivide(t *testing.T) {
	input := `(/ 1 2 3)`

	bytecode := []*code.Instruction{
		{
			OpCode: code.PUSH,
			Args:   0,
		},
		{
			OpCode: code.PUSH,
			Args:   1,
		},
		{
			OpCode: code.PUSH,
			Args:   2,
		},
		{
			OpCode: code.DIV,
			Args:   3,
		},
	}
	constants := []interface{}{1, 2, 3}

	checkBytecode(t, input, bytecode, constants)

}

func checkBytecode(t *testing.T, input string, expected []*code.Instruction, constants []interface{}) {
	t.Helper()

	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer, false)
	compiler := NewCompiler()

	ast := parser.Parse()
	if len(parser.Errors()) != 0 {
		t.Fatalf("Errors while parsing")
	}
	err := compiler.Compile(ast)

	if err != nil {
		t.Fatalf("Erorr when compiling: %s", err)
	}

	bytecode := compiler.Bytecode()

	checkInstructions(t, bytecode.Instructions, expected)
	checkConstants(t, bytecode.Constants, constants)
}
func checkConstants(t *testing.T, constants []object.Object, expected []interface{}) {
	t.Helper()

	if len(constants) != len(expected) {
		t.Fatalf("Expected %d constants, got %d", len(expected), len(constants))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			checkInt(t, constants[i], constant)
		case bool:
			checkBool(t, constants[i], constant)
		case []*code.Instruction:
			checkCompiledFunction(t, constants[i], constant)
		default:
			t.Fatalf("Unknown type %T", constant)
		}
	}
}
func checkCompiledFunction(t *testing.T, obj object.Object, expected []*code.Instruction) {
	t.Helper()

	if obj.Type() != object.COMPILED_FUNCTION_OBJ {
		t.Fatalf("Expected COMPILED_FUNCTION_OBJ, got %s", obj.Type())
	}

	fn, ok := obj.(*object.CompiledFunction)

	if !ok {
		t.Fatalf("Expected COMPILED_FUNCTION_OBJ, got %s", obj.Type())
	}

	checkInstructions(t, fn.Instructions, expected)
}
func checkBool(t *testing.T, obj object.Object, expected bool) {
	t.Helper()

	if obj.Type() != object.BOOLEAN_OBJ {
		t.Fatalf("Expected BOOLEAN_OBJ, got %s", obj.Type())
	}

	value := obj.(*object.Boolean).Value

	if value != expected {
		t.Fatalf("Expected %t, got %t", expected, value)
	}
}
func checkInt(t *testing.T, obj object.Object, expected int) {
	t.Helper()

	if obj.Type() != object.INTEGER_OBJ {
		t.Fatalf("Expected INTEGER_OBJ, got %s", obj.Type())
	}

	value := obj.(*object.Integer).Value
	if value != expected {
		t.Fatalf("Expected %d, got %d", expected, value)
	}
}

func checkInstructions(t *testing.T, bytecode []*code.Instruction, expected []*code.Instruction) {
	t.Helper()

	if len(bytecode) != len(expected) {
		t.Fatalf("Expected %d instructions, got %d", len(expected), len(bytecode))
	}

	for i, instr := range bytecode {
		if instr.OpCode != expected[i].OpCode {
			t.Fatalf("Expected OpCode %s, got %s", expected[i].OpCode, instr.OpCode)
		}
		if instr.Args != expected[i].Args {
			t.Fatalf("Expected Args %d, got %d for instruction %s", expected[i].Args, instr.Args, instr.OpCode)
		}
	}
}
