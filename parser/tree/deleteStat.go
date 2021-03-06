package tree

type DeleteStat struct {
	BaseStat
	tableNameAndAlias *TableNameAndAlias
	whereCondition    WhereCondition
	rowCount          string
}

func (this *DeleteStat) Children() []Stat {
	return Children(this)
}

func (this *DeleteStat) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *DeleteStat) TableNameAndAlias() *TableNameAndAlias {
	return this.tableNameAndAlias
}

func (this *DeleteStat) SetTableNameAndAlias(tableNameAndAlias *TableNameAndAlias) {
	this.tableNameAndAlias = tableNameAndAlias
}

func (this *DeleteStat) WhereCondition() WhereCondition {
	return this.whereCondition
}

func (this *DeleteStat) SetWhereCondition(whereCondition WhereCondition) {
	this.whereCondition = whereCondition
}

func (this *DeleteStat) RowCount() string {
	return this.rowCount
}

func (this *DeleteStat) SetRowCount(rowCount string) {
	this.rowCount = rowCount
}

var _ Stat = (*DeleteStat)(nil)

func (this *DeleteStat) String() string {
	sql := NewStringBuilder()
	sql.Append("delete from ").AppendStat(this.tableNameAndAlias)
	if this.whereCondition != nil {
		sql.Append(" where ").AppendStat(this.whereCondition)
	}
	if this.rowCount != "" {
		sql.Append(" limit ").Append(this.rowCount)
	}
	return sql.String()
}
