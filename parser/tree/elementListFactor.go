package tree

type ElementListFactor struct {
	BaseStat
	elementList *ElementList
}

func (this *ElementListFactor) Children() []Stat {
	return Children(this)
}

func (this *ElementListFactor) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementListFactor) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementListFactor) ElementList() *ElementList {
	return this.elementList
}

func (this *ElementListFactor) SetElementList(elementList *ElementList) {
	this.elementList = elementList
}

var _ Element = (*ElementListFactor)(nil)

func (this *ElementListFactor) String() string {
	return "(" + this.elementList.String() + ")"
}
