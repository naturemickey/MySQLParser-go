package tree

import "strings"

type ValueList struct {
	elements []Element
}

func (this *ValueList) Elements() []Element {
	return this.elements
}

func (this *ValueList) SetElements(elements []Element) {
	this.elements = elements
}

var _ Stat = &ValueList{}

func (this *ValueList) String() string {
	elementsStr := []string{}
	for _, element := range this.elements {
		elementsStr = append(elementsStr, element.String())
	}
	return strings.Join(elementsStr, ", ")
}
