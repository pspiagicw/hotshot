package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/parser"
	"github.com/pspiagicw/hotshot/printer"
)

func main() {
	for true {
		fmt.Printf(">>> ")
		buffer := bufio.NewReader(os.Stdin)
		prompt, err := buffer.ReadString('\n')
		if err != nil {
			log.Fatalf("Error scanning input: %v", err)
		}
		lexer := lexer.NewLexer(prompt)
		p := parser.NewParser(lexer)
		for !lexer.EOF {
			// fmt.Println(lexer.Next().String())
			// fmt.Println(p.Parse())
			fmt.Println(printer.PrintAST(p.Parse()))
			fmt.Println(p.Errors)
		}
	}
}
