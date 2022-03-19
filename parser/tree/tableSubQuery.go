package tree

type TableSubQuery struct {
	BaseStat
	selectStat *SelectStat
	alias      string
}

func (this *TableSubQuery) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *TableSubQuery) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

func (this *TableSubQuery) Alias() string {
	return this.alias
}

func (this *TableSubQuery) SetAlias(alias string) {
	this.alias = alias
}

var _ TableFactor = &TableSubQuery{}

func (this *TableSubQuery) String() string {
	sql := NewStringBuilder()
	sql.Append("(").AppendStat(this.selectStat).Append(")").Append(this.alias)
	return sql.String()
}
