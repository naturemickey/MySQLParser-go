package tree

type TableJoin struct {
	BaseStat
	tableNameAndAlias *TableNameAndAlias
	tableJoinSuffix   *TableJoinSuffix
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

var _ TableRel = &TableJoin{}

func (this *TableJoin) String() string {
	return NewStringBuilder().AppendStat(this.tableNameAndAlias).AppendStat(this.tableJoinSuffix).String()
}
