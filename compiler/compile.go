package compiler

import (
	"fmt"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/object"
	"github.com/pspiagicw/hotshot/token"
)

type Bytecode struct {
	Instructions []*code.Instruction
}

type Compiler struct {
	instructions []*code.Instruction
	constants    []object.Object
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
	}
}

func NewCompiler() *Compiler {
	return &Compiler{}
}
func (c *Compiler) emit(op code.Op, arg int16) {
	c.instructions = append(c.instructions, &code.Instruction{OpCode: op, Args: arg})
}
func (c *Compiler) compileCallStatement(node *ast.CallStatement) error {
	for _, arg := range node.Args {
		c.Compile(arg)
	}
	argCount := int16(len(node.Args))

	switch node.Op.TokenType {
	case token.PLUS:
		c.emit(code.ADD, argCount)
	case token.MINUS:
		c.emit(code.SUB, argCount)
	case token.MULTIPLY:
		c.emit(code.MUL, argCount)
	case token.SLASH:
		c.emit(code.DIV, argCount)
	case token.EQ:
		c.emit(code.EQ, argCount)
	case token.GREATERTHAN:
		c.emit(code.GT, argCount)
	case token.LESSTHAN:
		c.emit(code.LT, argCount)
	default:
		return fmt.Errorf("Unknown operator %s", node.Op.TokenType)
	}
	return nil
}

func (c *Compiler) Compile(node ast.Statement) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, statement := range node.Statements {
			err := c.Compile(statement)
			if err != nil {
				return err
			}
		}
	case *ast.IntStatement:
		constId := c.addConstant(toInteger(node))
		c.emit(code.PUSH, constId)
	case *ast.CallStatement:
		return c.compileCallStatement(node)
	case *ast.BoolStatement:
		if node.Value {
			c.emit(code.TRUE, -1)
		} else {
			c.emit(code.FALSE, -1)
		}
	case *ast.EmptyStatement:
	default:
		return fmt.Errorf("Unknown node type %T", node)
	}
	return nil
}
func (c *Compiler) addConstant(obj object.Object) int16 {
	c.constants = append(c.constants, obj)
	return int16(len(c.constants) - 1)
}

func toInteger(node *ast.IntStatement) object.Object {
	return &object.Integer{Value: node.Value}
}
