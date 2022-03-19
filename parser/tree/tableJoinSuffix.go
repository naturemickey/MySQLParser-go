package tree

type TableJoinSuffix struct {
	BaseStat
	tableJoinMod        *TableJoinMod
	tableNameAndAliases *TableNameAndAliases
	tableRecu           *TableRecu
	joinCondition       *JoinCondition
	tableJoinSuffix     *TableJoinSuffix
}

func (this *TableJoinSuffix) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *TableJoinSuffix) TableJoinMod() *TableJoinMod {
	return this.tableJoinMod
}

func (this *TableJoinSuffix) SetTableJoinMod(tableJoinMod *TableJoinMod) {
	this.tableJoinMod = tableJoinMod
}

func (this *TableJoinSuffix) TableNameAndAliases() *TableNameAndAliases {
	return this.tableNameAndAliases
}

func (this *TableJoinSuffix) SetTableNameAndAliases(tableNameAndAliases *TableNameAndAliases) {
	this.tableNameAndAliases = tableNameAndAliases
}

func (this *TableJoinSuffix) TableRecu() *TableRecu {
	return this.tableRecu
}

func (this *TableJoinSuffix) SetTableRecu(tableRecu *TableRecu) {
	this.tableRecu = tableRecu
}

func (this *TableJoinSuffix) JoinCondition() *JoinCondition {
	return this.joinCondition
}

func (this *TableJoinSuffix) SetJoinCondition(joinCondition *JoinCondition) {
	this.joinCondition = joinCondition
}

func (this *TableJoinSuffix) TableJoinSuffix() *TableJoinSuffix {
	return this.tableJoinSuffix
}

func (this *TableJoinSuffix) SetTableJoinSuffix(tableJoinSuffix *TableJoinSuffix) {
	this.tableJoinSuffix = tableJoinSuffix
}

var _ Stat = &TableJoinSuffix{}

func (this *TableJoinSuffix) String() string {
	sql := NewStringBuilder()
	sql.Append(" ").AppendStat(this.tableJoinMod).Append(" join ")
	if this.tableNameAndAliases != nil {
		sql.AppendStat(this.tableNameAndAliases)
	} else {
		sql.AppendStat(this.tableRecu)
	}
	if this.joinCondition != nil {
		sql.Append(" ").AppendStat(this.joinCondition)
	}
	if this.tableJoinSuffix != nil {
		sql.Append(" ").AppendStat(this.tableJoinSuffix)
	}
	return sql.String()
}
