package tree

type SelectPrefix struct {
	BaseStat
	distinct     bool
	selectExprs  *SelectExprs
	tables       *Tables
	where        WhereCondition
	groupByExprs *GbobExprs
	having       WhereCondition
}

func (this *SelectPrefix) Distinct() bool {
	return this.distinct
}

func (this *SelectPrefix) SetDistinct(distinct bool) {
	this.distinct = distinct
}

func (this *SelectPrefix) SelectExprs() *SelectExprs {
	return this.selectExprs
}

func (this *SelectPrefix) SetSelectExprs(selectExprs *SelectExprs) {
	this.selectExprs = selectExprs
}

func (this *SelectPrefix) Tables() *Tables {
	return this.tables
}

func (this *SelectPrefix) SetTables(tables *Tables) {
	this.tables = tables
}

func (this *SelectPrefix) Where() WhereCondition {
	return this.where
}

func (this *SelectPrefix) SetWhere(where WhereCondition) {
	this.where = where
}

func (this *SelectPrefix) GroupByExprs() *GbobExprs {
	return this.groupByExprs
}

func (this *SelectPrefix) SetGroupByExprs(groupByExprs *GbobExprs) {
	this.groupByExprs = groupByExprs
}

func (this *SelectPrefix) Having() WhereCondition {
	return this.having
}

func (this *SelectPrefix) SetHaving(having WhereCondition) {
	this.having = having
}

var _ Stat = &SelectPrefix{}

func (this *SelectPrefix) String() string {
	sql := NewStringBuilder()
	sql.Append("select ")
	if this.distinct {
		sql.Append("distinct ")
	}
	sql.AppendStat(this.selectExprs)
	if this.tables != nil {
		sql.Append(" from ").AppendStat(this.tables)
	}
	if this.where != nil {
		sql.Append(" where ").AppendStat(this.where)
	}
	if this.groupByExprs != nil {
		sql.Append(" group by ").AppendStat(this.groupByExprs)
	}
	if this.having != nil {
		sql.Append(" having ").AppendStat(this.having)
	}
	return sql.String()
}
