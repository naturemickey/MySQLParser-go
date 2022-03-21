package tree

type ExprInValues struct {
	BaseStat
	element   Element
	not       string
	valueList *ValueList
}

func (this *ExprInValues) _Expression() {
	//TODO implement me
	panic("implement me")
}

func (this *ExprInValues) Children() []Stat {
	return Children(this)
}

func (this *ExprInValues) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprInValues) Element() Element {
	return this.element
}

func (this *ExprInValues) SetElement(element Element) {
	this.element = element
}

func (this *ExprInValues) Not() string {
	return this.not
}

func (this *ExprInValues) SetNot(not string) {
	this.not = not
}

func (this *ExprInValues) ValueList() *ValueList {
	return this.valueList
}

func (this *ExprInValues) SetValueList(valueList *ValueList) {
	this.valueList = valueList
}

var _ Expression = (*ExprInValues)(nil)

func (this *ExprInValues) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.element).Append(" ")
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append("in (").AppendStat(this.valueList).Append(")")
	return sql.String()
}
