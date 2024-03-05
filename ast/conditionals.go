package ast

import (
	"strings"

	"github.com/shivamMg/ppds/tree"
)

type IfStatement struct {
	Condition Statement
	Body      Statement
	Else      Statement
}

func (i IfStatement) String() string {
	var output strings.Builder
	output.WriteString(tree.SprintHrn(i))
	return output.String()
}
func (i IfStatement) Data() interface{} {
	return "if"
}

func (i IfStatement) Children() []tree.Node {
	if i.Else != nil {
		return []tree.Node{
			i.Condition,
			i.Body,
			i.Else,
		}
	}
	return []tree.Node{
		i.Condition,
		i.Body,
	}
}

type WhileStatement struct {
	Condition Statement
	Body      Statement
}

func (w WhileStatement) String() string {
	var output strings.Builder
	output.WriteString(tree.SprintHrn(w))
	return output.String()
}
func (w WhileStatement) Data() interface{} {
	return "while"
}
func (w WhileStatement) Children() []tree.Node {
	return []tree.Node{
		w.Condition,
		w.Body,
	}
}
