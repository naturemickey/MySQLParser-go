package tree

type WhereCondOp struct {
	BaseStat
	expression         Expression
	expressionOperator string
	whereCondition     WhereCondition
}

func (this *WhereCondOp) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *WhereCondOp) Expression() Expression {
	return this.expression
}

func (this *WhereCondOp) SetExpression(expression Expression) {
	this.expression = expression
}

func (this *WhereCondOp) ExpressionOperator() string {
	return this.expressionOperator
}

func (this *WhereCondOp) SetExpressionOperator(expressionOperator string) {
	this.expressionOperator = expressionOperator
}

func (this *WhereCondOp) WhereCondition() WhereCondition {
	return this.whereCondition
}

func (this *WhereCondOp) SetWhereCondition(whereCondition WhereCondition) {
	this.whereCondition = whereCondition
}

var _ WhereCondition = &WhereCondOp{}

func (this *WhereCondOp) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.expression)
	if this.whereCondition != nil {
		sql.Append(" ").Append(this.expressionOperator).Append(" ").AppendStat(this.whereCondition)
	}
	return sql.String()
}
