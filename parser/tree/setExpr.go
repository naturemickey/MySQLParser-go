package tree

type SetExpr struct {
	BaseStat
	left         Element
	right        Element
	rightDefault string
}

func (this *SetExpr) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *SetExpr) Left() Element {
	return this.left
}

func (this *SetExpr) SetLeft(left Element) {
	this.left = left
}

func (this *SetExpr) Right() Element {
	return this.right
}

func (this *SetExpr) SetRight(right Element) {
	this.right = right
}

func (this *SetExpr) RightDefault() string {
	return this.rightDefault
}

func (this *SetExpr) SetRightDefault(rightDefault string) {
	this.rightDefault = rightDefault
}

var _ Stat = &SetExpr{}

func (this *SetExpr) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.left).Append(" = ")
	if this.right != nil {
		sql.AppendStat(this.right)
	} else {
		sql.Append(this.rightDefault)
	}
	return sql.String()
}
