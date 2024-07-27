package decompiler

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/object"
)

func Print(bytecode *compiler.Bytecode) {
	printConstants(bytecode)
	fmt.Println("=== BYTECODE ===")
	printInstructions(bytecode.Instructions)
}
func printConstants(bytecode *compiler.Bytecode) {
	if len(bytecode.Constants) != 0 {
		fmt.Println("=== CONSTANT ===")
	}
	for _, constant := range bytecode.Constants {
		switch constant := constant.(type) {
		case *object.CompiledFunction:
			fmt.Println("-- COMPILED START ---")
			printInstructions(constant.Instructions)
			fmt.Println("--- COMPILED END ---")
		default:
			fmt.Printf("%s\n", constant)
		}
	}
}
func printInstructions(instructions []*code.Instruction) {
	line := 0
	for _, instruction := range instructions {
		op := instruction.OpCode.String()

		lineNumber := getLineNumber(line)
		argString := ""
		if instruction.Args >= 0 {
			argString = fmt.Sprintf("%d", instruction.Args)
		}
		fmt.Printf("%s %s %s\t%s\n", lineNumber, op, argString, instruction.Comment)
		line++
	}
}

func getLineNumber(line int) string {
	return lipgloss.NewStyle().Faint(true).Render(fmt.Sprintf("%05d", line))
}
