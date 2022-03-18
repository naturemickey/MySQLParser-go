package tree

type ElementCase struct {
	value        Element
	caseWhenPart *CaseWhenPart
	elseEl       Element
}

var _ ElementOpFactory = &ElementCase{}

func (this *ElementCase) String() string {
	sql := NewStringBuilder()
	sql.Append("case ")
	if this.value != nil {
		sql.AppendStat(this.value).Append(" ")
	}
	sql.AppendStat(this.caseWhenPart).Append(" ")
	if this.elseEl != nil {
		sql.AppendStat(this.elseEl).Append(" ")
	}
	sql.Append("end")
	return sql.String()
}
