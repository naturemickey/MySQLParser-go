package tree

type WhereCondSub struct {
	BaseStat
	wc1                WhereCondition
	wc2                WhereCondition
	expressionOperator string
}

func (this *WhereCondSub) Children() []Stat {
	return Children(this)
}

func (this *WhereCondSub) _WhereCondition() {
	//TODO implement me
	panic("implement me")
}

func (this *WhereCondSub) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *WhereCondSub) Wc1() WhereCondition {
	return this.wc1
}

func (this *WhereCondSub) SetWc1(wc1 WhereCondition) {
	this.wc1 = wc1
}

func (this *WhereCondSub) Wc2() WhereCondition {
	return this.wc2
}

func (this *WhereCondSub) SetWc2(wc2 WhereCondition) {
	this.wc2 = wc2
}

func (this *WhereCondSub) ExpressionOperator() string {
	return this.expressionOperator
}

func (this *WhereCondSub) SetExpressionOperator(expressionOperator string) {
	this.expressionOperator = expressionOperator
}

var _ WhereCondition = (*WhereCondSub)(nil)

func (this *WhereCondSub) String() string {
	sql := NewStringBuilder()
	sql.Append("(").AppendStat(this.wc1).Append(")")
	if this.wc2 != nil {
		sql.Append(" ").Append(this.expressionOperator).Append(" ").AppendStat(this.wc2)
	}
	return sql.String()
}
