package tree

type ElementWapperBkt struct {
	BaseStat
	element Element
}

func (this *ElementWapperBkt) Children() []Stat {
	return Children(this)
}

func (this *ElementWapperBkt) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementWapperBkt) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementWapperBkt) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementWapperBkt) Element() Element {
	return this.element
}

func (this *ElementWapperBkt) SetElement(element Element) {
	this.element = element
}

var _ ElementOpFactory = &ElementWapperBkt{}

func (this *ElementWapperBkt) String() string {
	return "(" + this.element.String() + ")"
}
