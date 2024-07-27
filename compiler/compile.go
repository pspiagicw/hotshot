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
	Constants    []object.Object
}

type Scope struct {
	instructions []*code.Instruction
}

type Builtin struct {
	id      int
	builtin *object.Builtin
}

type Compiler struct {
	constants  []object.Object
	scopes     []Scope
	scopeIndex int
	jid        int16
	symbols    *SymbolTable
}

func (c *Compiler) enterScope() {
	scope := Scope{
		instructions: []*code.Instruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.scopeIndex++

	c.symbols = NewEnclosedSymbolTable(c.symbols)
}

func (c *Compiler) leaveScope() []*code.Instruction {
	instructions := c.currentInstructions()
	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	c.symbols = c.symbols.Outer
	return instructions
}

func NewCompiler() *Compiler {
	mainScope := Scope{
		instructions: []*code.Instruction{},
	}
	symbols := NewSymbolTable()
	return &Compiler{
		scopes:     []Scope{mainScope},
		scopeIndex: 0,
		constants:  []object.Object{},
		jid:        0,
		symbols:    symbols,
	}
}
func NewWithState(symbols *SymbolTable, constants []object.Object) *Compiler {
	mainScope := Scope{
		instructions: []*code.Instruction{},
	}
	return &Compiler{
		scopeIndex: 0,
		scopes:     []Scope{mainScope},
		constants:  constants,
		jid:        0,
		symbols:    symbols,
	}
}
func (c *Compiler) currentInstructions() []*code.Instruction {
	return c.scopes[c.scopeIndex].instructions
}

func (c *Compiler) Bytecode() *Bytecode {
	// c.constantPass()
	c.conditionalsPass()
	return &Bytecode{
		Instructions: c.currentInstructions(),
		Constants:    c.constants,
	}
}
func (c *Compiler) JumpID() int16 {
	c.jid++
	return c.jid
}

func (c *Compiler) conditionalsPass() {
}

func (c *Compiler) emit(op code.Op, arg int16) {
	ins := &code.Instruction{OpCode: op, Args: arg}
	currentScope := c.scopes[c.scopeIndex]
	currentScope.instructions = append(currentScope.instructions, ins)
	c.scopes[c.scopeIndex] = currentScope
}
func (c *Compiler) compileCallStatement(node *ast.CallStatement) error {
	for _, arg := range node.Args {
		err := c.Compile(arg)
		if err != nil {
			return err
		}
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
	case token.IDENT:
		return c.compileFunctionCall(node)
	default:
		return fmt.Errorf("Unknown operator %s", node.Op.TokenType)
	}
	return nil
}
func (c *Compiler) compileFunctionCall(node *ast.CallStatement) error {
	symbol, ok := c.symbols.Resolve(node.Op.TokenValue)

	if !ok {
		return fmt.Errorf("Undefined function %s", node.Op.TokenValue)
	}

	c.loadSymbol(symbol)

	c.emit(code.CALL, int16(len(node.Args)))

	return nil
}
func (c *Compiler) loadSymbol(symbol Symbol) {
	if symbol.Scope == Global {
		c.emit(code.GET, int16(symbol.Index))
	} else if symbol.Scope == Local {
		c.emit(code.LGET, int16(symbol.Index))
	} else {
		c.emit(code.BUILTIN, int16(symbol.Index))
	}
}
func toIdent(node ast.Statement) (*ast.IdentStatement, error) {
	val, ok := node.(*ast.IdentStatement)

	if !ok {
		return nil, fmt.Errorf("Expected ident statement, got %T", node)
	}

	return val, nil
}

func (c *Compiler) compileLambdaStatement(node *ast.LambdaStatement) error {
	c.enterScope()

	for _, p := range node.Args {
		ident, err := toIdent(p)

		if err != nil {
			return err
		}

		c.symbols.Define(ident.Value.TokenValue)
	}

	for _, statement := range node.Body {
		err := c.Compile(statement)

		if err != nil {
			return err
		}
	}
	instructions := c.leaveScope()
	compiledFn := &object.CompiledFunction{Instructions: instructions}
	c.emit(code.PUSH, c.addConstant(compiledFn))

	return nil
}
func (c *Compiler) compileFunctionStatement(node *ast.FunctionStatement) error {
	// For recursion this should be defined before the body, and obvously before entering scope
	symbol := c.symbols.Define(node.Name.TokenValue)

	c.enterScope()

	for _, p := range node.Args {
		ident, err := toIdent(p)

		if err != nil {
			return err
		}

		c.symbols.Define(ident.Value.TokenValue)
	}

	for _, statement := range node.Body {
		err := c.Compile(statement)

		if err != nil {
			return err
		}
	}
	instructions := c.leaveScope()
	compiledFn := &object.CompiledFunction{Instructions: instructions}
	c.emit(code.PUSH, c.addConstant(compiledFn))

	c.emit(code.SET, int16(symbol.Index))

	return nil

}
func (c *Compiler) compileAssertStatement(node *ast.AssertStatement) error {
	return nil
}
func (c *Compiler) compileTableStatement(node *ast.TableStatement) error {
	for _, element := range node.Elements {
		err := c.Compile(element)

		if err != nil {
			return err
		}
	}
	c.emit(code.TABLE, int16(len(node.Elements)))
	return nil
}

func (c *Compiler) Compile(node ast.Statement) error {
	switch node := node.(type) {
	case *ast.AssertStatement:
		return c.compileAssertStatement(node)
	case *ast.FunctionStatement:
		return c.compileFunctionStatement(node)
	case *ast.LambdaStatement:
		return c.compileLambdaStatement(node)
	case *ast.Program:
		for _, statement := range node.Statements {
			err := c.Compile(statement)
			if err != nil {
				return err
			}
		}
	case *ast.TableStatement:
		return c.compileTableStatement(node)
	case *ast.StringStatement:
		constId := c.addConstant(&object.String{Value: node.Value})
		c.emit(code.PUSH, constId)
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
	case *ast.IfStatement:
		return c.compileIfStatement(node)
	case *ast.AssignmentStatement:
		return c.compileAssignmentStatement(node)
	case *ast.IdentStatement:
		return c.compileIdentStatement(node)
	case *ast.WhileStatement:
		return c.compileWhileStatement(node)
	case *ast.CondStatement:
		return c.compileCondStatement(node)
	case *ast.EmptyStatement:
	default:
		return fmt.Errorf("Unknown node type %T", node)
	}
	return nil
}
func (c *Compiler) compileCondStatement(node *ast.CondStatement) error {
	endID := c.JumpID()
	for _, cond := range node.Expressions {
		condID := c.JumpID()

		err := c.Compile(cond.Condition)
		if err != nil {
			return err
		}
		c.emit(code.JCMP, condID)
		err = c.Compile(cond.Body)
		if err != nil {
			return err
		}
		c.emit(code.JMP, endID)
		c.emit(code.JT, condID)
	}
	c.emit(code.JT, endID)
	return nil
}
func (c *Compiler) compileWhileStatement(node *ast.WhileStatement) error {
	condID := c.JumpID()
	c.emit(code.JT, condID)
	err := c.Compile(node.Condition)

	if err != nil {
		return err
	}

	bodyID := c.JumpID()
	c.emit(code.JCMP, bodyID)

	err = c.Compile(node.Body)

	if err != nil {
		return err
	}

	c.emit(code.JMP, condID)
	c.emit(code.JT, bodyID)
	return nil

}
func (c *Compiler) compileIdentStatement(node *ast.IdentStatement) error {
	symbol, ok := c.symbols.Resolve(node.Value.TokenValue)
	if !ok {
		return fmt.Errorf("Undefined variable %s", node.Value.TokenValue)
	}

	c.loadSymbol(symbol)
	return nil
}
func (c *Compiler) compileAssignmentStatement(node *ast.AssignmentStatement) error {
	err := c.Compile(node.Value)
	if err != nil {
		return err
	}
	symbol := c.symbols.Define(node.Name.TokenValue)
	if symbol.Scope == Global {
		c.emit(code.SET, int16(symbol.Index))
	} else {
		c.emit(code.LSET, int16(symbol.Index))
	}
	return nil
}
func (c *Compiler) compileIfStatement(node *ast.IfStatement) error {

	err := c.Compile(node.Condition)

	if err != nil {
		return err
	}

	consequenceID := c.JumpID()

	c.emit(code.JCMP, consequenceID)

	err = c.Compile(node.Body)

	if err != nil {
		return err
	}

	if node.Else != nil {
		elseID := c.JumpID()
		c.emit(code.JMP, elseID)
		c.emit(code.JT, consequenceID)
		err = c.Compile(node.Else)
		if err != nil {
			return err
		}
		c.emit(code.JT, elseID)
	} else {
		c.emit(code.JT, consequenceID)
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
