package tree

type TableSubQuery struct {
	BaseStat
	selectStat *SelectStat
	alias      string
}

func (this *TableSubQuery) Children() []Stat {
	return Children(this)
}

func (this *TableSubQuery) _TableRel() {
	//TODO implement me
	panic("implement me")
}

func (this *TableSubQuery) _TableFactor() {
	//TODO implement me
	panic("implement me")
}

func (this *TableSubQuery) Assemble(stats []Stat) {
	Assemble(this, stats)
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

var _ TableFactor = (*TableSubQuery)(nil)

func (this *TableSubQuery) String() string {
	sql := NewStringBuilder()
	sql.Append("(").AppendStat(this.selectStat).Append(") ").Append(this.alias)
	return sql.String()
}
