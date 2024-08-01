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
const FramesSize = 1024

type VM struct {
	constants []object.Object

	builtins     []object.BuiltinIndex
	essentials   map[string]*object.Builtin
	stack        []object.Object
	stackPointer int

	globals    []object.Object
	frames     []*Frame
	frameIndex int
}

func NewVM(bytecode *compiler.Bytecode) *VM {
	mainFn := &object.CompiledFunction{Instructions: bytecode.Instructions}
	mainFrame := NewFrame(mainFn, 0)

	frames := make([]*Frame, FramesSize)
	frames[0] = mainFrame

	return &VM{
		constants: bytecode.Constants,

		stack:        make([]object.Object, StackSize),
		stackPointer: 0,
		essentials:   object.Essentials(),
		globals:      make([]object.Object, GlobalsSize),
		frames:       frames,
		frameIndex:   1,
	}
}
func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.frameIndex-1]
}
func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.frameIndex] = f
	vm.frameIndex++
}
func (vm *VM) popFrame() *Frame {
	vm.frameIndex--
	return vm.frames[vm.frameIndex]
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
func (vm *VM) executeCall(instr *code.Instruction) error {
	top := vm.pop()
	fn, ok := top.(*object.CompiledFunction)

	if !ok {
		return fmt.Errorf("calling non-function")
	}
	frame := NewFrame(fn, vm.stackPointer)
	vm.pushFrame(frame)
	vm.stackPointer = frame.basePointer + int(fn.NumLocals)
	return nil
}
func (vm *VM) executeReturn() error {
	result := vm.pop()
	prevFrame := vm.popFrame()
	vm.stackPointer = prevFrame.basePointer
	return vm.push(result)
}
func (vm *VM) executeLocalSet(instr *code.Instruction) error {
	frame := vm.currentFrame()

	vm.stack[frame.basePointer+int(instr.Operand)] = vm.pop()

	return nil
}
func (vm *VM) executeLocalGet(instr *code.Instruction) error {
	frame := vm.currentFrame()

	value := vm.stack[frame.basePointer+int(instr.Operand)]

	return vm.push(value)
}
func (vm *VM) Run() error {
	var ip int
	var ins []*code.Instruction
	var err error
	for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
		vm.currentFrame().ip++
		ip = vm.currentFrame().ip
		ins = vm.currentFrame().Instructions()

		instr := ins[ip]

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
			vm.currentFrame().ip += int(instr.Operand)
		case code.JCMP:
			if !vm.stackTrue() {
				vm.currentFrame().ip += int(instr.Operand)
			}
		case code.SET:
			err = vm.executeSet(instr)
		case code.GET:
			err = vm.executeGet(instr)
		case code.CALL:
			err = vm.executeCall(instr)
		case code.RETURN:
			err = vm.executeReturn()
		case code.LSET:
			err = vm.executeLocalSet(instr)
		case code.LGET:
			err = vm.executeLocalGet(instr)
		default:
			err = fmt.Errorf("Unknown opcode %s", instr.OpCode.String())
		}

		if err != nil {
			return err
		}
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
