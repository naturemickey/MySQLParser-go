package tree

type UpdateSingleTable struct {
	BaseStat
	tableNameAndAlias *TableNameAndAlias
	setExprs          *SetExprs
	whereCondition    WhereCondition
	rowCount          string
}

func (this *UpdateSingleTable) _UpdateStat() {
	//TODO implement me
	panic("implement me")
}

func (this *UpdateSingleTable) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *UpdateSingleTable) TableNameAndAlias() *TableNameAndAlias {
	return this.tableNameAndAlias
}

func (this *UpdateSingleTable) SetTableNameAndAlias(tableNameAndAlias *TableNameAndAlias) {
	this.tableNameAndAlias = tableNameAndAlias
}

func (this *UpdateSingleTable) SetExprs() *SetExprs {
	return this.setExprs
}

func (this *UpdateSingleTable) SetSetExprs(setExprs *SetExprs) {
	this.setExprs = setExprs
}

func (this *UpdateSingleTable) WhereCondition() WhereCondition {
	return this.whereCondition
}

func (this *UpdateSingleTable) SetWhereCondition(whereCondition WhereCondition) {
	this.whereCondition = whereCondition
}

func (this *UpdateSingleTable) RowCount() string {
	return this.rowCount
}

func (this *UpdateSingleTable) SetRowCount(rowCount string) {
	this.rowCount = rowCount
}

var _ UpdateStat = &UpdateSingleTable{}

func (this *UpdateSingleTable) String() string {
	sql := NewStringBuilder()
	sql.Append("update ").AppendStat(this.tableNameAndAlias).Append(" set ").AppendStat(this.setExprs)
	if this.whereCondition != nil {
		sql.Append(" where ").AppendStat(this.whereCondition)
	}
	if this.rowCount != "" {
		sql.Append(" limit ").Append(this.rowCount)
	}
	return sql.String()
}
