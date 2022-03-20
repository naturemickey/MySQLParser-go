package tree

import "strings"

type ValueList struct {
	BaseStat
	elements []Element
}

func (this *ValueList) Children() []Stat {
	return Children(this)
}

func (this *ValueList) Elements() []Element {
	return this.elements
}

func (this *ValueList) SetElements(elements []Element) {
	this.elements = elements
}
func (this *ValueList) Assemble(stats []Stat) {
	for _, stat := range stats {
		element := stat.(Element)
		this.elements = append(this.elements, element)
	}
}

var _ Stat = (*ValueList)(nil)

func (this *ValueList) String() string {
	elementsStr := []string{}
	for _, element := range this.elements {
		elementsStr = append(elementsStr, element.String())
	}
	return strings.Join(elementsStr, ", ")
}
