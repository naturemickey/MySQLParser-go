package tree

type ExprBetweenAnd struct {
	BaseStat
	el    Element
	not   string
	left  Element
	right Element
}

func (this *ExprBetweenAnd) El() Element {
	return this.el
}

func (this *ExprBetweenAnd) SetEl(el Element) {
	this.el = el
}

func (this *ExprBetweenAnd) Not() string {
	return this.not
}

func (this *ExprBetweenAnd) SetNot(not string) {
	this.not = not
}

func (this *ExprBetweenAnd) Left() Element {
	return this.left
}

func (this *ExprBetweenAnd) SetLeft(left Element) {
	this.left = left
}

func (this *ExprBetweenAnd) Right() Element {
	return this.right
}

func (this *ExprBetweenAnd) SetRight(right Element) {
	this.right = right
}

var _ Expression = &ExprBetweenAnd{}

func (this *ExprBetweenAnd) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.el)
	if this.not != "" {
		sql.Append(" ").Append(this.not)
	}
	sql.Append(" between ").AppendStat(this.left).Append(" and ").AppendStat(this.right)
	return sql.String()
}
