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
	Content() string
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
func (f Function) Content() string {
	return "function"
}
func (f Function) String() string {
	args := len(f.Args)
	return fmt.Sprintf("(fn with %d args at %p)", args, &f)
}

type Null struct {
}

func (n Null) Type() ObjectType {
	return NULL_OBJ
}
func (n Null) String() string {
	return "null"
}
func (n Null) Content() string {
	return "null"
}

type Integer struct {
	Value int
}

func (i Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i Integer) String() string {
	return fmt.Sprintf("(int %d)", i.Value)
}
func (i Integer) Content() string {
	return fmt.Sprintf("%d", i.Value)
}

type String struct {
	Value string
}

func (s String) Type() ObjectType {
	return STRING_OBJ
}

func (s String) String() string {
	return fmt.Sprintf("(str %s)", s.Value)
}
func (s String) Content() string {
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
func (s Builtin) Content() string {
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
func (e Error) Content() string {
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
		return "(bool true)"
	}
	return "(bool false)"
}
func (b Boolean) Content() string {
	return fmt.Sprintf("%t", b.Value)
}

type Table struct {
	Elements  []Object
	Hash      map[string]Object
	arrLength int
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
	output.WriteString(" + ")
	output.WriteString("{")
	elements = []string{}
	for key, value := range t.Hash {
		elements = append(elements, fmt.Sprintf("%s: %s", key, value.Content()))
	}
	output.WriteString(strings.Join(elements, ", "))
	output.WriteString("}")
	return output.String()
}
func (t Table) Content() string {
	var output strings.Builder

	output.WriteString("[")

	elements := []string{}
	for _, element := range t.Elements {
		elements = append(elements, element.Content())
	}

	output.WriteString(strings.Join(elements, " "))
	output.WriteString("]")
	return output.String()
}
func (t *Table) Index(key Object) Object {
	switch key := key.(type) {
	case *Integer:
		if key.Value < 0 || key.Value >= len(t.Elements) {
			val, ok := t.Hash[key.String()]
			if !ok {
				return Error{Message: "Index out of range!"}
			}
			return val
		}
		return t.Elements[key.Value]
	case *String:
		if value, ok := t.Hash[key.String()]; ok {
			return value
		}
		return Error{Message: fmt.Sprintf("Key not found: %s", key.Value)}
	default:
		return Error{Message: fmt.Sprintf("Invalid type of key received %T", key)}
	}
}
func (t *Table) Set(key Object, value Object) Object {
	t.Hash[key.String()] = value
	return Null{}
}

type Return struct {
}

func (r Return) Type() ObjectType {
	return RETURN_OBJ
}
func (r Return) String() string {
	return "RETURN"
}
func (r Return) Content() string {
	return "RETURN"
}
