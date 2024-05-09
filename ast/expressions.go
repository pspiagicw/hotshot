package ast

import (
	"fmt"

	"github.com/pspiagicw/hotshot/token"
	"github.com/shivamMg/ppds/tree"
)

type IntStatement struct {
	Value int
}

func (i IntStatement) String() string {
	return fmt.Sprintf("Integer(%d)", i.Value)
}
func (i IntStatement) Data() interface{} {
	return i.String()
}
func (i IntStatement) Children() []tree.Node {
	return []tree.Node{}
}

type EmptyStatement struct {
}

func (e EmptyStatement) String() string {
	return "()"
}
func (e EmptyStatement) Data() interface{} {
	return e.String()
}
func (e EmptyStatement) Children() []tree.Node {
	return []tree.Node{}
}

type StringStatement struct {
	Value string
}

func (s StringStatement) String() string {
	return fmt.Sprintf("String(%s)", s.Value)
}
func (s StringStatement) Data() interface{} {
	return s.String()
}
func (e StringStatement) Children() []tree.Node {
	return []tree.Node{}
}

type BoolStatement struct {
	Value bool
}

func (b BoolStatement) String() string {
	return fmt.Sprintf("Bool(%t)", b.Value)
}
func (b BoolStatement) Data() interface{} {
	return b.String()
}
func (b BoolStatement) Children() []tree.Node {
	return []tree.Node{}
}

type IdentStatement struct {
	Value *token.Token
}

func (i IdentStatement) String() string {
	return fmt.Sprintf("Ident(%s)", i.Value.TokenValue)
}

func (i IdentStatement) Data() interface{} {
	return i.String()
}

func (i IdentStatement) Children() []tree.Node {
	return []tree.Node{}
}
