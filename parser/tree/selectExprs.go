package tree

type SelectExprs struct {
	element     Element
	alias       string
	selectExprs *SelectExprs
}

func (this *SelectExprs) Element() Element {
	return this.element
}

func (this *SelectExprs) SetElement(element Element) {
	this.element = element
}

func (this *SelectExprs) Alias() string {
	return this.alias
}

func (this *SelectExprs) SetAlias(alias string) {
	this.alias = alias
}

func (this *SelectExprs) SelectExprs() *SelectExprs {
	return this.selectExprs
}

func (this *SelectExprs) SetSelectExprs(selectExprs *SelectExprs) {
	this.selectExprs = selectExprs
}

var _ Stat = &SelectExprs{}

func (this *SelectExprs) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.element)
	if this.alias != "" {
		sql.Append(" ").Append(this.alias)
	}
	if this.selectExprs != nil {
		sql.Append(", ").AppendStat(this.selectExprs)
	}
	return sql.String()
}
