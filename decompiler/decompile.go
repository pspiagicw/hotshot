package decompiler

import (
	"fmt"

	"github.com/pspiagicw/hotshot/compiler"
)

func Print(bytecode *compiler.Bytecode) {
	line := 0
	for _, instruction := range bytecode.Instructions {
		op := instruction.OpCode.String()

		// if instruction.Args != nil {
		// 	showArgument = instruction.Args.String()
		// }
		fmt.Printf("%05d %s %d\n", line, op, instruction.Args)
		line++
	}
}
