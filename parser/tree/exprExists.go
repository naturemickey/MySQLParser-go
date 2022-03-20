package tree

type ExprExists struct {
	BaseStat
	not        string
	selectStat *SelectStat
}

func (this *ExprExists) Children() []Stat {
	return Children(this)
}

func (this *ExprExists) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprExists) Not() string {
	return this.not
}

func (this *ExprExists) SetNot(not string) {
	this.not = not
}

func (this *ExprExists) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *ExprExists) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

var _ Expression = (*ExprExists)(nil)

func (this *ExprExists) String() string {
	sql := NewStringBuilder()
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append("exists (").AppendStat(this.selectStat).Append(")")
	return sql.String()
}
