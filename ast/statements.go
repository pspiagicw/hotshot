package ast

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/token"
	"github.com/shivamMg/ppds/tree"
)

type CallStatement struct {
	Op   *token.Token
	Args []Statement
}

func (f CallStatement) String() string {
	return tree.SprintHrn(f)
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

type AssertStatement struct {
	Left    Statement
	Right   Statement
	Message *token.Token
}

func (a AssertStatement) String() string {
	return tree.SprintHrn(a)
}
func (a AssertStatement) Data() interface{} {
	return fmt.Sprintf("assert")
}
func (a AssertStatement) Children() []tree.Node {
	return []tree.Node{
		a.Left,
		a.Right,
		&StringStatement{
			Value: a.Message.String(),
		},
	}
}

type AssignmentStatement struct {
	Name  *token.Token
	Value Statement
}

func (a AssignmentStatement) String() string {
	return tree.SprintHrn(a)
}
func (a AssignmentStatement) Data() interface{} {
	return fmt.Sprintf("$(%s)", a.Name.TokenValue)
}
func (a AssignmentStatement) Children() []tree.Node {
	return []tree.Node{
		a.Value,
	}
}

type ImportStatement struct {
	Package string
}

func (i ImportStatement) String() string {
	return tree.SprintHrn(i)
}

func (i ImportStatement) Data() interface{} {
	return fmt.Sprintf("import(%s)", i.Package)
}

func (i ImportStatement) Children() []tree.Node {
	return []tree.Node{}
}

type FunctionStatement struct {
	Name *token.Token
	Args []Statement
	Body []Statement
}

func (f FunctionStatement) String() string {
	return tree.SprintHrn(f)
}
func (f FunctionStatement) Data() interface{} {
	strArgs := []string{}

	for _, arg := range f.Args {
		strArgs = append(strArgs, arg.String())
	}

	return fmt.Sprintf("fn(%s[%s])", f.Name.TokenValue, strings.Join(strArgs, ","))
}
func (f FunctionStatement) Children() []tree.Node {
	return []tree.Node{}
}

type LambdaStatement struct {
	Args []Statement
	Body []Statement
}

func (l LambdaStatement) String() string {
	return tree.SprintHrn(l)
}
func (l LambdaStatement) Data() interface{} {
	strArgs := []string{}

	for _, arg := range l.Args {
		strArgs = append(strArgs, arg.String())
	}
	return fmt.Sprintf("lambda([%s])", strings.Join(strArgs, ","))
}
func (l LambdaStatement) Children() []tree.Node {
	return []tree.Node{}
}

type QuoteStatement struct {
	Body *token.Token
}

func (q QuoteStatement) String() string {
	return fmt.Sprintf("quote(%s)", q.Body.TokenValue)
}
func (q QuoteStatement) Data() interface{} {
	return q.Body.TokenValue
}
func (q QuoteStatement) Children() []tree.Node {
	return []tree.Node{}
}

type IndexStatement struct {
	Key    Statement
	Target Statement
}

func (s IndexStatement) String() string {
	return tree.SprintHrn(s)
}
func (s IndexStatement) Data() interface{} {
	return "slice"
}
func (s IndexStatement) Children() []tree.Node {
	return []tree.Node{
		s.Key,
		s.Target,
	}
}

type SetStatement struct {
	Target *IndexStatement
	Value  Statement
}

func (s SetStatement) String() string {
	return tree.SprintHrn(s)
}
func (s SetStatement) Data() interface{} {
	return "set"
}
func (s SetStatement) Children() []tree.Node {
	return []tree.Node{
		s.Target,
		s.Value,
	}
}
