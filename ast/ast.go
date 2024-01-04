package ast

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/token"
	"github.com/shivamMg/ppds/tree"
)

type Program struct {
	Statements []Statement
}

type Statement interface {
	StringifyStatement() string
	// For `ppds` package
	Data() interface{}
	Children() []tree.Node
}

type IntStatement struct {
	Value int
}

type EmptyStatement struct {
}

type StringStatement struct {
	Value string
}

type BoolStatement struct {
	Value bool
}

type FunctionalStatement struct {
	Op   *token.Token
	Args []Statement
}

func (b BoolStatement) StringifyStatement() string {
	return fmt.Sprintf("Bool(%t)", b.Value)
}
func (b BoolStatement) String() string {
	return b.StringifyStatement()
}
func (b BoolStatement) Data() interface{} {
	return b.Value
}
func (b BoolStatement) Children() []tree.Node {
	return []tree.Node{}
}

func (f FunctionalStatement) StringifyStatement() string {
	var output strings.Builder
	output.WriteString(tree.SprintHrn(f))
	return output.String()
}

func (f FunctionalStatement) String() string {
	return f.StringifyStatement()
}
func (f FunctionalStatement) Data() interface{} {
	return f.Op.TokenValue
}
func (f FunctionalStatement) Children() []tree.Node {
	nodes := []tree.Node{}
	for _, statement := range f.Args {
		nodes = append(nodes, statement)
	}
	return nodes
}

func (s StringStatement) StringifyStatement() string {
	return fmt.Sprintf("String(%s)", s.Value)
}
func (s StringStatement) String() string {
	return s.StringifyStatement()
}
func (s StringStatement) Data() interface{} {
	return s.StringifyStatement()
}
func (e StringStatement) Children() []tree.Node {
	return []tree.Node{}
}

func (e EmptyStatement) StringifyStatement() string {
	return "EmptyStatement()"
}
func (e EmptyStatement) String() string {
	return e.StringifyStatement()
}
func (e EmptyStatement) Data() interface{} {
	return e.StringifyStatement()
}
func (e EmptyStatement) Children() []tree.Node {
	return []tree.Node{}
}

func (i IntStatement) StringifyStatement() string {
	return fmt.Sprintf("Integer(%d)", i.Value)
}
func (i IntStatement) Data() interface{} {
	return i.StringifyStatement()
}
func (i IntStatement) Children() []tree.Node {
	return []tree.Node{}
}
