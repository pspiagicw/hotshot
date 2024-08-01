package vm

import (
	"github.com/pspiagicw/hotshot/code"
	"github.com/pspiagicw/hotshot/object"
)

type Frame struct {
	fn *object.CompiledFunction
	ip int
}

func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

func (f *Frame) Instructions() []*code.Instruction {
	return f.fn.Instructions
}
