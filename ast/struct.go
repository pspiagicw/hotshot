package ast

import (
	"github.com/shivamMg/ppds/tree"
)

type TableStatement struct {
	Elements []Statement
}

func (t TableStatement) String() string {
	return tree.SprintHrn(t)
}
func (t TableStatement) Data() interface{} {
	return "{}"
}
func (t TableStatement) Children() []tree.Node {
	nodes := []tree.Node{}
	for _, statement := range t.Elements {
		nodes = append(nodes, statement)
	}
	return nodes
}
