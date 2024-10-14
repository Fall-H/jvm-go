package heap

import "ch_07/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}
