package tree

type ExprNot struct {
	expression Expression
}

func (this *ExprNot) Expression() Expression {
	return this.expression
}

func (this *ExprNot) SetExpression(expression Expression) {
	this.expression = expression
}

var _ Expression = &ExprNot{}

func (this *ExprNot) String() string {
	return "not " + this.expression.String()
}
