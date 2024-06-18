package decompiler

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/hotshot/compiler"
)

func Print(bytecode *compiler.Bytecode) {
	line := 0
	for _, instruction := range bytecode.Instructions {
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
