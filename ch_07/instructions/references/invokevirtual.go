package references

import (
	"ch_07/instructions/base"
	"ch_07/rtda"
	"ch_07/rtda/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct{ base.Index16Instruction }

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if ref == nil {
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func _println(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Println("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Println("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Println("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Println("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Println("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Println("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Println("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Println("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}
