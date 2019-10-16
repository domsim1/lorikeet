package vm

import (
	"lorikeet/code"
	"lorikeet/object"
)

// Frame holds execution-relevant information
type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

// NewFrame inits frame
func NewFrame(cl *object.Closure, basePointer int) *Frame {
	f := &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}

	return f
}

// Instructions gets compiled function in the frame
func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
