package decompiler

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/object"
)

var compiledStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
var constantStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Padding(1).MarginRight(1)
var codeStyle lipgloss.Style = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder())

func Print(bytecode *compiler.Bytecode) {
	constants := constantStyle.Render(printConstants(bytecode))
	code := codeStyle.Render(printInstructions(bytecode.Instructions))

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, constants, code))
}
func printConstants(bytecode *compiler.Bytecode) string {
	var buffer strings.Builder
	for i, constant := range bytecode.Constants {
		switch constant := constant.(type) {
		case *object.CompiledFunction:
			if len(constant.Instructions) != 0 {
				instructions := printInstructions(constant.Instructions)
				content := compiledStyle.Render(fmt.Sprintf("--- %s ---\n%s", constant.Name, instructions))
				buffer.WriteString(content)
				buffer.WriteString("\n")
			} else {
				buffer.WriteString(compiledStyle.Render("VOID"))
			}
		default:
			buffer.WriteString(fmt.Sprintf("%02d. %s\n", i, constant))
		}
	}
	return buffer.String()
}
func printInstructions(instructions []*code.Instruction) string {
	var buffer strings.Builder
	line := 0
	for _, instruction := range instructions {
		op := instruction.OpCode.String()

		lineNumber := getLineNumber(line)
		argString := ""
		if instruction.Operand >= 0 {
			argString = fmt.Sprintf("%d", instruction.Operand)
		}
		buffer.WriteString(fmt.Sprintf("%s %s %s\t%s\n", lineNumber, op, argString, instruction.Comment))
		line++
	}
	return strings.TrimSpace(buffer.String())
}

func getLineNumber(line int) string {
	return lipgloss.NewStyle().Faint(true).Render(fmt.Sprintf("%05d", line))
}
