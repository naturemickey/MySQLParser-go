package tree

type ElementListFactor struct {
	elementList *ElementList
}

func (this *ElementListFactor) ElementList() *ElementList {
	return this.elementList
}

func (this *ElementListFactor) SetElementList(elementList *ElementList) {
	this.elementList = elementList
}

var _ Element = &ElementListFactor{}

func (this *ElementListFactor) String() string {
	return "(" + this.elementList.String() + ")"
}
