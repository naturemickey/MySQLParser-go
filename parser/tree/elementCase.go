package tree

type ElementCase struct {
	BaseStat
	value        Element
	caseWhenPart *CaseWhenPart
	elseEl       Element
}

func (this *ElementCase) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementCase) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementCase) Assemble(stats []Stat) {
	Assemble(this, stats)
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
