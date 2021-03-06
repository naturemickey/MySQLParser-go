package tree

type InsertStat struct {
	BaseStat
	tableName   string
	columnNames *ColumnNames
	valueList   *ValueList
	selectStat  *SelectStat
}

func (this *InsertStat) Children() []Stat {
	return Children(this)
}

func (this *InsertStat) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *InsertStat) TableName() string {
	return this.tableName
}

func (this *InsertStat) SetTableName(tableName string) {
	this.tableName = tableName
}

func (this *InsertStat) ColumnNames() *ColumnNames {
	return this.columnNames
}

func (this *InsertStat) SetColumnNames(columnNames *ColumnNames) {
	this.columnNames = columnNames
}

func (this *InsertStat) ValueList() *ValueList {
	return this.valueList
}

func (this *InsertStat) SetValueList(valueList *ValueList) {
	this.valueList = valueList
}

func (this *InsertStat) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *InsertStat) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

var _ Stat = (*InsertStat)(nil)

func (this *InsertStat) String() string {
	sql := NewStringBuilder().Append("insert into ").Append(this.tableName)
	if this.columnNames != nil {
		sql.Append("(").AppendStat(this.columnNames).Append(")")
	}
	if this.valueList != nil {
		sql.Append(" values(").AppendStat(this.valueList).Append(")")
	} else {
		sql.Append(" ").AppendStat(this.selectStat)
	}
	return sql.String()
}
