package tree

type ParamList struct {
	BaseStat
	element        Element
	exprRelational *ExprRelational
	paramList      *ParamList
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
