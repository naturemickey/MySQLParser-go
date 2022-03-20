package tree

type ParamList struct {
	BaseStat
	element        Element
	exprRelational *ExprRelational
	paramList      *ParamList
}

func (this *ParamList) Element() Element {
	return this.element
}

func (this *ParamList) SetElement(element Element) {
	this.element = element
}

func (this *ParamList) ExprRelational() *ExprRelational {
	return this.exprRelational
}

func (this *ParamList) SetExprRelational(exprRelational *ExprRelational) {
	this.exprRelational = exprRelational
}

func (this *ParamList) ParamList() *ParamList {
	return this.paramList
}

func (this *ParamList) SetParamList(paramList *ParamList) {
	this.paramList = paramList
}

func (this *ParamList) Assemble(stats []Stat) {
	Assemble(this, stats)
}

var _ Stat = &ParamList{}

func (this *ParamList) String() string {
	sql := NewStringBuilder()
	if this.element != nil {
		sql.AppendStat(this.element)
	} else {
		sql.AppendStat(this.exprRelational)
	}
	if this.paramList != nil {
		sql.Append(", ").AppendStat(this.paramList)
	}
	return sql.String()
}
