package tree

type ElementList struct {
	BaseStat
	elements []Element
}

func (this *ElementList) Children() []Stat {
	return Children(this)
}

func (this *ElementList) Assemble(stats []Stat) {
	for _, stat := range stats {
		this.elements = append(this.elements, stat.(Element))
	}
}

func (this *ElementList) Elements() []Element {
	return this.elements
}

func (this *ElementList) SetElements(elements []Element) {
	this.elements = elements
}

var _ Stat = (*ElementList)(nil)

func (this *ElementList) String() string {
	sql := NewStringBuilder()
	for _, element := range this.elements {
		sql.AppendStat(element).Append(", ")
	}
	sql.deleteLast()
	return sql.String()
}
