package tree

type ExprInValues struct {
	element   Element
	not       string
	valueList *ValueList
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

var _ Expression = &ExprInValues{}

func (this *ExprInValues) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.element).Append(" ")
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append(" in (").AppendStat(this.valueList).Append(")")
	return sql.String()
}
