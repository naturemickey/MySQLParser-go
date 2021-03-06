package tree

type TableJoin struct {
	BaseStat
	tableNameAndAlias *TableNameAndAlias
	tableJoinSuffix   *TableJoinSuffix
}

func (this *TableJoin) Children() []Stat {
	return Children(this)
}

func (this *TableJoin) _TableRel() {
	//TODO implement me
	panic("implement me")
}

func (this *TableJoin) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *TableJoin) TableNameAndAlias() *TableNameAndAlias {
	return this.tableNameAndAlias
}

func (this *TableJoin) SetTableNameAndAlias(tableNameAndAlias *TableNameAndAlias) {
	this.tableNameAndAlias = tableNameAndAlias
}

func (this *TableJoin) TableJoinSuffix() *TableJoinSuffix {
	return this.tableJoinSuffix
}

func (this *TableJoin) SetTableJoinSuffix(tableJoinSuffix *TableJoinSuffix) {
	this.tableJoinSuffix = tableJoinSuffix
}

var _ TableRel = (*TableJoin)(nil)

func (this *TableJoin) String() string {
	return NewStringBuilder().AppendStat(this.tableNameAndAlias).AppendStat(this.tableJoinSuffix).String()
}
