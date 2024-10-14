package control

import (
	"ch_07/instructions/base"
	"ch_07/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }  // Return void from method
type ARETURN struct{ base.NoOperandsInstruction } // Return reference from met
type DRETURN struct{ base.NoOperandsInstruction } // Return double from method
type FRETURN struct{ base.NoOperandsInstruction } // Return float from method
type IRETURN struct{ base.NoOperandsInstruction } // Return int from method
type LRETURN struct{ base.NoOperandsInstruction } // Return long from method

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}
