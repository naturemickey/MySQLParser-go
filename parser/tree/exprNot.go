package tree

type ExprNot struct {
	BaseStat
	expression Expression
}

func (this *ExprNot) Children() []Stat {
	return Children(this)
}

func (this *ExprNot) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprNot) Expression() Expression {
	return this.expression
}

func (this *ExprNot) SetExpression(expression Expression) {
	this.expression = expression
}

var _ Expression = (*ExprNot)(nil)

func (this *ExprNot) String() string {
	return "not " + this.expression.String()
}
