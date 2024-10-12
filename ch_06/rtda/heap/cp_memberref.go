package heap

import "ch_06/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}

func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}

func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}

func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}

func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
