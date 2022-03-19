package tree

type ElementSubQuery struct {
	BaseStat
	sqWith     string
	selectStat *SelectStat
}

func (this *ElementSubQuery) SqWith() string {
	return this.sqWith
}

func (this *ElementSubQuery) SetSqWith(sqWith string) {
	this.sqWith = sqWith
}

func (this *ElementSubQuery) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *ElementSubQuery) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

var _ ElementOpFactory = &ElementSubQuery{}

func (this *ElementSubQuery) String() string {
	sql := NewStringBuilder()
	if this.sqWith != "" {
		sql.Append(this.SqWith()).Append(" ")
	}
	sql.Append("(").AppendStat(this.selectStat).Append(")")
	return sql.String()
}
