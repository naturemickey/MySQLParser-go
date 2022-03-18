package tree

type TableNameAndAliases struct {
	tableNameAndAliases []*TableNameAndAlias
}

func (this *TableNameAndAliases) TableNameAndAliases() []*TableNameAndAlias {
	return this.tableNameAndAliases
}

func (this *TableNameAndAliases) SetTableNameAndAliases(tableNameAndAliases []*TableNameAndAlias) {
	this.tableNameAndAliases = tableNameAndAliases
}

var _ Stat = &TableNameAndAliases{}

func (this *TableNameAndAliases) String() string {
	sql := NewStringBuilder()
	for _, alias := range this.tableNameAndAliases {
		sql.AppendStat(alias).Append(", ")
	}
	sql.deleteLast()
	return sql.String()
}
