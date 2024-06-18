package compiler

import (
	"testing"

	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/parser"
)

func TestIfElse(t *testing.T) {
	input := `(if true 10 20)`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.JCMP, Args: 2},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.JMP, Args: 1},
		{OpCode: code.PUSH, Args: 1},
	}
	checkBytecode(t, input, bytecode)
}

func TestIf(t *testing.T) {
	input := `(if true 10)`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.JCMP, Args: 1},
		{OpCode: code.PUSH, Args: 0},
	}

	checkBytecode(t, input, bytecode)
}

func TestComparison(t *testing.T) {
	input := `(> 1 2) (< 1 2) (= 1 2)`

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.GT, Args: 2},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.LT, Args: 2},
		{OpCode: code.PUSH, Args: 0},
		{OpCode: code.PUSH, Args: 1},
		{OpCode: code.EQ, Args: 2},
	}
	checkBytecode(t, input, bytecode)
}
func TestBoolean(t *testing.T) {
	input := `true false`

	bytecode := []*code.Instruction{
		{OpCode: code.TRUE, Args: -1},
		{OpCode: code.FALSE, Args: -1},
	}

	checkBytecode(t, input, bytecode)

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
	checkBytecode(t, input, bytecode)
}

func TestPush(t *testing.T) {
	input := `1`

	bytecode := []*code.Instruction{
		{OpCode: code.PUSH, Args: 0},
	}

	checkBytecode(t, input, bytecode)
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

	checkBytecode(t, input, bytecode)
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

	checkBytecode(t, input, bytecode)
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

	checkBytecode(t, input, bytecode)

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

	checkBytecode(t, input, bytecode)

}

func checkBytecode(t *testing.T, input string, expected []*code.Instruction) {
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

	for i, instr := range bytecode.Instructions {
		if instr.OpCode != expected[i].OpCode {
			t.Fatalf("Expected OpCode %s, got %s", expected[i].OpCode, instr.OpCode)
		}
		if instr.Args != expected[i].Args {
			t.Fatalf("Expected Args %d, got %d", expected[i].Args, instr.Args)
		}
	}
}