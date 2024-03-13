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

// type AssignmentStatement struct {
// 	Name  *token.Token
// 	Value Statement
// }
//
// func (a AssignmentStatement) String() string {
// 	return tree.SprintHrn(a)
// }
// func (a AssignmentStatement) Data() interface{} {
// 	return fmt.Sprintf("$(%s)", a.Name.TokenValue)
// }
// func (a AssignmentStatement) Children() []tree.Node {
// 	return []tree.Node{
// 		a.Value,
// 	}
// }

type FunctionStatement struct {
	Name *token.Token
	Args []Statement
	Body Statement
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
	return []tree.Node{
		f.Body,
	}
}

type LambdaStatement struct {
	Args []Statement
	Body Statement
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
	return []tree.Node{
		l.Body,
	}
}
