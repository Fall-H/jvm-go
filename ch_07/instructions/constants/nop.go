package constants

import (
	"ch_07/instructions/base"
	"ch_07/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
