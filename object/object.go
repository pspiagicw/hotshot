package object

import (
	"fmt"
	"strings"

	"github.com/pspiagicw/hotshot/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	String() string
}

const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"
	TABLE_OBJ   = "TABLE"

	NULL_OBJ     = "NULL"
	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"
	ERROR_OBJ    = "ERROR"
	RETURN_OBJ   = "RETURN"
)

type Function struct {
	Args []*ast.IdentStatement
	Body []ast.Statement
	Env  *Environment
}

func (f Function) Type() ObjectType {
	return FUNCTION_OBJ
}
func (f Function) String() string {
	return "fn()"
}

type Null struct {
}

func (n Null) Type() ObjectType {
	return NULL_OBJ
}
func (n Null) String() string {
	return "NULL"
}

type Integer struct {
	Value int
}

func (i Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type String struct {
	Value string
}

func (s String) Type() ObjectType {
	return STRING_OBJ
}

func (s String) String() string {
	return s.Value
}

type BuiltinFunc func(args []Object) Object

type Builtin struct {
	Fn BuiltinFunc
}

func (s Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
func (s Builtin) String() string {
	return "builtin function"
}

type Error struct {
	Message string
}

func (e Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e Error) String() string {
	return e.Message
}

type Boolean struct {
	Value bool
}

func (b Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b Boolean) String() string {
	if b.Value {
		return "true"
	}
	return "false"
}

type Table struct {
	Elements []Object
}

func (t Table) Type() ObjectType {
	return TABLE_OBJ
}
func (t Table) String() string {
	var output strings.Builder

	output.WriteString("[")

	elements := []string{}
	for _, element := range t.Elements {
		elements = append(elements, element.String())
	}

	output.WriteString(strings.Join(elements, " "))
	output.WriteString("]")
	return output.String()
}

type Return struct {
}

func (r Return) Type() ObjectType {
	return RETURN_OBJ
}
func (r Return) String() string {
	return "RETURN"
}
