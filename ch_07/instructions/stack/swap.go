package stack

import (
	"ch_07/instructions/base"
	"ch_07/rtda"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

/*
bottom -> top
[...][c][b][a]

	 \/
	 /\
	V  V

[...][c][a][b]
*/
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
