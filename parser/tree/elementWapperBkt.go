package tree

type ElementWapperBkt struct {
	element Element
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
