package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	String() string
}

const (
	INTEGER_OBJ = "INTEGER"
	STRING_OBJ  = "STRING"
	BOOLEAN_OBJ = "BOOLEAN"

	NULL_OBJ     = "NULL"
	FUNCTION_OBJ = "FUNCTION"
	BUILTIN_OBJ  = "BUILTIN"
	ERROR_OBJ    = "ERROR"
)

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

type BuiltinFunc func(args ...Object) Object

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
