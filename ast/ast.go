package ast

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/token"
)

type Program struct {
	Statements []Statement
}

type Statement interface {
	StringifyStatement() string
}

type IntStatement struct {
	Value int
}

type EmptyStatement struct {
}

type StringStatement struct {
	Value string
}

type FunctionalStatement struct {
	Op   *token.Token
	Args []Statement
}

func (f FunctionalStatement) StringifyStatement() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("Func(%s) [\n", f.Op.TokenValue))
	for _, arg := range f.Args {
		if arg != nil {
			output.WriteString(fmt.Sprintf("  %s", arg.StringifyStatement()))
		} else {
			output.WriteString("  NIL Statement")
		}
		output.WriteString("  \n")
	}
	output.WriteString("]\n")
	return output.String()
}

func (f FunctionalStatement) String() string {
	return f.StringifyStatement()
}

func (s StringStatement) StringifyStatement() string {
	return fmt.Sprintf("String(%s)", s.Value)
}
func (s StringStatement) String() string {
	return s.StringifyStatement()
}

func (e EmptyStatement) StringifyStatement() string {
	return "EmptyStatement()"
}
func (e EmptyStatement) String() string {
	return e.StringifyStatement()
}

func (i IntStatement) StringifyStatement() string {
	return fmt.Sprintf("Integer(%d)", i.Value)
}
