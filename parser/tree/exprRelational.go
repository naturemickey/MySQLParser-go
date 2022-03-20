package tree

type ExprRelational struct {
	BaseStat
	left         Element
	right        Element
	relationalOp string
}

func (this *ExprRelational) Children() []Stat {
	return Children(this)
}

func (this *ExprRelational) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprRelational) Left() Element {
	return this.left
}

func (this *ExprRelational) SetLeft(left Element) {
	this.left = left
}

func (this *ExprRelational) Right() Element {
	return this.right
}

func (this *ExprRelational) SetRight(right Element) {
	this.right = right
}

func (this *ExprRelational) RelationalOp() string {
	return this.relationalOp
}

func (this *ExprRelational) SetRelationalOp(relationalOp string) {
	this.relationalOp = relationalOp
}

var _ Expression = (*ExprRelational)(nil)

func (this *ExprRelational) String() string {
	return this.left.String() + " " + this.relationalOp + " " + this.right.String()
}
