package compiler

import (
	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/object"
)

func createTable(instructions []*code.Instruction) map[int16]int {
	table := make(map[int16]int)

	for i := 0; i < len(instructions); i++ {
		cur := instructions[i]

		if cur.OpCode == code.JT {
			table[cur.Operand] = i
		}
	}

	return table
}
func editInstructions(instructions []*code.Instruction, table map[int16]int) []*code.Instruction {
	for i := 0; i < len(instructions); i++ {
		cur := instructions[i]

		if cur.OpCode == code.JCMP || cur.OpCode == code.JMP {
			cur.Operand = int16(table[cur.Operand] - i)
		}

		instructions[i] = cur
	}

	return instructions
}
func updateFunction(instructions []*code.Instruction) []*code.Instruction {
	table := createTable(instructions)
	instructions = editInstructions(instructions, table)
	return instructions
}

func JumpPass(bytecode *Bytecode) *Bytecode {

	instructions := bytecode.Instructions

	table := createTable(instructions)
	instructions = editInstructions(instructions, table)

	for _, constant := range bytecode.Constants {
		switch obj := constant.(type) {
		case *object.CompiledFunction:
			instructions := updateFunction(obj.Instructions)
			obj.Instructions = instructions
		}
	}

	return bytecode
}
