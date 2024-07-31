package vm

import (
	"fmt"

	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/compiler"
	"github.com/pspiagicw/hotshot/object"
)

const (
	STATUS_OK = iota
	STATUS_ERROR
)

var TRUE = &object.Boolean{Value: true}
var FALSE = &object.Boolean{Value: false}

const StackSize = 2048
const GlobalsSize = 65536

type VM struct {
	instructions []*code.Instruction
	constants    []object.Object

	builtins     []object.BuiltinIndex
	essentials   map[string]*object.Builtin
	stack        []object.Object
	stackPointer int

	globals []object.Object
}

func NewVM(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack:        make([]object.Object, StackSize),
		stackPointer: 0,
		essentials:   object.Essentials(),
		globals:      make([]object.Object, GlobalsSize),
	}
}
func (vm *VM) executePush(instr *code.Instruction) error {
	constant := vm.constants[instr.Operand]
	return vm.push(constant)
}
func (vm *VM) executeAdd(instr *code.Instruction) error {
	elements := vm.popElements(instr.Operand)

	fn := vm.essentials["+"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeSub(instr *code.Instruction) error {
	elements := vm.popElements(instr.Operand)

	fn := vm.essentials["-"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeMul(instr *code.Instruction) error {
	elements := vm.popElements(instr.Operand)

	fn := vm.essentials["*"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeDiv(instr *code.Instruction) error {
	elements := vm.popElements(instr.Operand)

	fn := vm.essentials["/"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeLT() error {
	elements := vm.popElements(2)

	fn := vm.essentials["<"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeGT() error {
	elements := vm.popElements(2)

	fn := vm.essentials[">"]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeEQ() error {
	elements := vm.popElements(2)

	fn := vm.essentials["="]

	result := fn.Fn(elements)

	return vm.push(result)
}
func (vm *VM) executeTrue() error {
	return vm.push(TRUE)
}
func (vm *VM) executeFalse() error {
	return vm.push(FALSE)
}
func (vm *VM) executeSet(instr *code.Instruction) error {
	globalIndex := instr.Operand
	vm.globals[globalIndex] = vm.pop()
	return nil
}
func (vm *VM) executeGet(instr *code.Instruction) error {
	globalIndex := instr.Operand
	return vm.push(vm.globals[globalIndex])
}
func (vm *VM) Run() error {
	ip := 0
	var err error
	for ip < len(vm.instructions) {
		instr := vm.instructions[ip]

		switch instr.OpCode {
		case code.JT:
			// Do nothing
		case code.PUSH:
			err = vm.executePush(instr)
		case code.ADD:
			err = vm.executeAdd(instr)
		case code.SUB:
			err = vm.executeSub(instr)
		case code.MUL:
			err = vm.executeMul(instr)
		case code.DIV:
			err = vm.executeDiv(instr)
		case code.TRUE:
			err = vm.executeTrue()
		case code.FALSE:
			err = vm.executeFalse()
		case code.LT:
			err = vm.executeLT()
		case code.GT:
			err = vm.executeGT()
		case code.EQ:
			err = vm.executeEQ()
		case code.JMP:
			ip += int(instr.Operand)
		case code.JCMP:
			if !vm.stackTrue() {
				ip += int(instr.Operand)
			}
		case code.SET:
			err = vm.executeSet(instr)
		case code.GET:
			err = vm.executeGet(instr)
		default:
			err = fmt.Errorf("Unknown opcode %s", instr.OpCode.String())
		}

		if err != nil {
			return err
		}
		ip++
	}
	return nil
}
func (vm *VM) popElements(number int16) []object.Object {
	result := make([]object.Object, number)

	var i int16
	for i = 0; i < number; i++ {
		result[number-i-1] = vm.pop()
	}

	return result
}
func (vm *VM) pop() object.Object {
	obj := vm.stack[vm.stackPointer-1]
	vm.stackPointer--
	return obj
}
func (vm *VM) push(value object.Object) error {
	if vm.stackPointer >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.stackPointer] = value
	vm.stackPointer++

	return nil
}
func (vm *VM) stackTrue() bool {
	if vm.StackTop().Type() != object.BOOLEAN_OBJ {
		return false
	}

	b, ok := vm.StackTop().(*object.Boolean)

	if !ok {
		return false
	}

	return b.Value
}
func (vm *VM) StackTop() object.Object {
	if vm.stackPointer == 0 {
		return nil
	}
	return vm.stack[vm.stackPointer-1]
}
