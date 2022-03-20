package tree

type UpdateMultipleTable struct {
	BaseStat
	tableNameAndAliases *TableNameAndAliases
	setExprs            *SetExprs
	whereCondition      WhereCondition
}

func (this *UpdateMultipleTable) _UpdateStat() {
	//TODO implement me
	panic("implement me")
}

func (this *UpdateMultipleTable) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *UpdateMultipleTable) TableNameAndAliases() *TableNameAndAliases {
	return this.tableNameAndAliases
}

func (this *UpdateMultipleTable) SetTableNameAndAliases(tableNameAndAliases *TableNameAndAliases) {
	this.tableNameAndAliases = tableNameAndAliases
}

func (this *UpdateMultipleTable) SetExprs() *SetExprs {
	return this.setExprs
}

func (this *UpdateMultipleTable) SetSetExprs(setExprs *SetExprs) {
	this.setExprs = setExprs
}

func (this *UpdateMultipleTable) WhereCondition() WhereCondition {
	return this.whereCondition
}

func (this *UpdateMultipleTable) SetWhereCondition(whereCondition WhereCondition) {
	this.whereCondition = whereCondition
}

var _ UpdateStat = &UpdateMultipleTable{}

func (this *UpdateMultipleTable) String() string {
	sql := NewStringBuilder()
	sql.Append("update ").AppendStat(this.tableNameAndAliases).Append(" set ").AppendStat(this.setExprs)
	if this.whereCondition != nil {
		sql.Append(" where ").AppendStat(this.whereCondition)
	}
	return sql.String()
}
