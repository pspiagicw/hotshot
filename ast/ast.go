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

func (p *Program) StringifyStatement() string {
	var output strings.Builder

	for _, statement := range p.Statements {
		output.WriteString(statement.StringifyStatement())
		output.WriteString("\n")
	}
	output.WriteString("\n")
	return output.String()
}
func (p *Program) Data() interface{} {
	return "[HT]"
}
func (p *Program) Children() []tree.Node {
	nodes := []tree.Node{}
	for _, statement := range p.Statements {
		nodes = append(nodes, statement)
	}
	return nodes
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

type CallStatement struct {
	Op   *token.Token
	Args []Statement
}

type IdentStatement struct {
	Value *token.Token
}

type AssignmentStatement struct {
	Name  *token.Token
	Value Statement
}

type FunctionStatement struct {
	Name *token.Token
	Args []Statement
	Body Statement
}

type IfStatement struct {
	Condition Statement
	Body      Statement
	Else      Statement
}

func (i IfStatement) StringifyStatement() string {
	return fmt.Sprintf("If statement")
}
func (i IfStatement) String() string {
	return i.StringifyStatement()
}
func (i IfStatement) Data() interface{} {
	return "if"
}
func (i IfStatement) Children() []tree.Node {
	if i.Else != nil {
		return []tree.Node{
			i.Body,
			i.Else,
		}
	}
	return []tree.Node{
		i.Body,
	}
}

func (f FunctionStatement) StringifyStatement() string {
	return fmt.Sprintf("Function(%s)", f.Name.TokenValue)
}
func (f FunctionStatement) String() string {
	return f.StringifyStatement()
}
func (f FunctionStatement) Data() interface{} {
	strArgs := []string{}

	for _, arg := range f.Args {
		strArgs = append(strArgs, arg.StringifyStatement())
	}

	return fmt.Sprintf("fn(%s[%s])", f.Name.TokenValue, strings.Join(strArgs, ","))
}
func (f FunctionStatement) Children() []tree.Node {
	return []tree.Node{
		f.Body,
	}

}

func (a AssignmentStatement) StringifyStatement() string {
	return fmt.Sprintf("($ %s %s)", a.Name.String(), a.Value.StringifyStatement())
}
func (a AssignmentStatement) String() string {
	return a.StringifyStatement()
}
func (a AssignmentStatement) Data() interface{} {
	return fmt.Sprintf("$(%s)", a.Name.TokenValue)
}
func (a AssignmentStatement) Children() []tree.Node {
	return []tree.Node{
		a.Value,
	}
}

func (i IdentStatement) StringifyStatement() string {
	return fmt.Sprintf("Ident(%s)", i.Value.TokenValue)
}

func (i IdentStatement) String() string {
	return i.StringifyStatement()
}

func (i IdentStatement) Data() interface{} {
	return i.StringifyStatement()
}

func (i IdentStatement) Children() []tree.Node {
	return []tree.Node{}
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

func (f CallStatement) StringifyStatement() string {
	var output strings.Builder
	output.WriteString(tree.SprintHrn(f))
	return output.String()
}

func (f CallStatement) String() string {
	return f.StringifyStatement()
}
func (f CallStatement) Data() interface{} {
	return f.Op.TokenValue
}
func (f CallStatement) Children() []tree.Node {
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
