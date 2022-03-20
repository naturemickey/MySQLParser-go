package tree

type JoinCondition struct {
	BaseStat
	whereCondition WhereCondition
	columnNames    *ColumnNames
}

func (this *JoinCondition) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *JoinCondition) WhereCondition() WhereCondition {
	return this.whereCondition
}

func (this *JoinCondition) SetWhereCondition(whereCondition WhereCondition) {
	this.whereCondition = whereCondition
}

func (this *JoinCondition) ColumnNames() *ColumnNames {
	return this.columnNames
}

func (this *JoinCondition) SetColumnNames(columnNames *ColumnNames) {
	this.columnNames = columnNames
}

var _ Stat = &JoinCondition{}

func (this *JoinCondition) String() string {
	sql := NewStringBuilder()
	if this.whereCondition != nil {
		sql.Append("on ").AppendStat(this.whereCondition)
	} else {
		sql.Append("using(").AppendStat(this.columnNames).Append(")")
	}
	return sql.String()
}
