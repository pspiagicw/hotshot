package help

import "github.com/pspiagicw/pelp"

func Help() {
	pelp.Print("A LISPy shell scripting language")
	pelp.HeaderWithDescription("Usage", []string{"hotshot [flags] [file]"})
	pelp.Flags(
		"flags",
		[]string{"help", "print-ast", "print-tokens", "null"},
		[]string{"Prints this help message", "Prints the AST of the program", "Prints the tokens of the program", "Print NULL in REPL"},
	)
}
