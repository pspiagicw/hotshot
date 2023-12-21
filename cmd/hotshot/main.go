package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
		prompt = strings.TrimSpace(prompt)
		lexer := lexer.NewLexer(prompt)
		p := parser.NewParser(lexer)
		fmt.Println(printer.PrintAST(p.Parse()))
		if len(p.Errors()) == 0 {
			fmt.Println("No parsing errors!")
		} else {
			fmt.Println("Error found during parsing!")
			for _, err := range p.Errors() {
				fmt.Println(err.Error())
			}
		}
	}
}
