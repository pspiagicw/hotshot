package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pspiagicw/hotshot/eval"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/parser"
)

func main() {
	for true {
		fmt.Printf(">>> ")
		buffer := bufio.NewReader(os.Stdin)
		prompt, err := buffer.ReadString('\n')
		if err != nil {
			log.Fatalf("Error scanning input: %v", err)
		}

		prompt = strings.TrimSpace(prompt)

		lexer := lexer.NewLexer(prompt)
		p := parser.NewParser(lexer)

		program := p.Parse()

		program.Statements = program.Statements[:len(program.Statements)-1]

		env := object.NewEnvironment()

		if len(p.Errors()) == 0 {
			result := eval.Eval(program, env)
			fmt.Print("=> ")
			fmt.Print(result)
			fmt.Println()
		} else {
			fmt.Println("Error found during parsing!")
			for _, err := range p.Errors() {
				fmt.Println(err.Error())
			}
		}
	}
}
