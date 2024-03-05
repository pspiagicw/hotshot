package ast

import (
	"strings"

	"github.com/shivamMg/ppds/tree"
)

type Statement interface {
	String() string
	// For `ppds` package
	Data() interface{}
	Children() []tree.Node
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var output strings.Builder

	for _, statement := range p.Statements {
		output.WriteString(statement.String())
		output.WriteString("\n")
	}
	output.WriteString("\n")
	return output.String()
}

func (p *Program) Data() interface{} {
	return "[hotshot]"
}
func (p *Program) Children() []tree.Node {
	nodes := []tree.Node{}
	for _, statement := range p.Statements {
		nodes = append(nodes, statement)
	}
	return nodes
}
