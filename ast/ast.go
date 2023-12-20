package ast

import "fmt"

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

func (e EmptyStatement) StringifyStatement() string {
	return "EmptyStatement()"
}

func (i IntStatement) StringifyStatement() string {
	return fmt.Sprintf("Integer(%d)", i.Value)
}
