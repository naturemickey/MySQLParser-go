package tree

type ExprInSelect struct {
	BaseStat
	element    Element
	not        string
	selectStat *SelectStat
}

func (this *ExprInSelect) _Expression() {
	//TODO implement me
	panic("implement me")
}

func (this *ExprInSelect) Children() []Stat {
	return Children(this)
}

func (this *ExprInSelect) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprInSelect) Element() Element {
	return this.element
}

func (this *ExprInSelect) SetElement(element Element) {
	this.element = element
}

func (this *ExprInSelect) Not() string {
	return this.not
}

func (this *ExprInSelect) SetNot(not string) {
	this.not = not
}

func (this *ExprInSelect) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *ExprInSelect) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

var _ Expression = (*ExprInSelect)(nil)

func (this *ExprInSelect) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.element).Append(" ")
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append("in (").AppendStat(this.selectStat).Append(")")
	return sql.String()
}
