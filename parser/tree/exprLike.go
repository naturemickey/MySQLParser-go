package tree

type ExprLike struct {
	BaseStat
	left  Element
	right Element
	not   string
}

func (this *ExprLike) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprLike) Left() Element {
	return this.left
}

func (this *ExprLike) SetLeft(left Element) {
	this.left = left
}

func (this *ExprLike) Right() Element {
	return this.right
}

func (this *ExprLike) SetRight(right Element) {
	this.right = right
}

func (this *ExprLike) Not() string {
	return this.not
}

func (this *ExprLike) SetNot(not string) {
	this.not = not
}

var _ Expression = &ExprLike{}

func (this *ExprLike) String() string {
	sql := NewStringBuilder().AppendStat(this.left).Append(" ")
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append("like ").AppendStat(this.right)
	return sql.String()
}
