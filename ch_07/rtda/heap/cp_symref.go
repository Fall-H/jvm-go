package heap

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolvedClassRef()
	}

	return self.class
}

func (self *SymRef) resolvedClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)

	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
