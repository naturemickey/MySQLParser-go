package tree

type ExprIsOrIsNotNull struct {
	BaseStat
	element Element
	not     string
	what    string
}

func (this *ExprIsOrIsNotNull) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ExprIsOrIsNotNull) Element() Element {
	return this.element
}

func (this *ExprIsOrIsNotNull) SetElement(element Element) {
	this.element = element
}

func (this *ExprIsOrIsNotNull) Not() string {
	return this.not
}

func (this *ExprIsOrIsNotNull) SetNot(not string) {
	this.not = not
}

func (this *ExprIsOrIsNotNull) What() string {
	return this.what
}

func (this *ExprIsOrIsNotNull) SetWhat(what string) {
	this.what = what
}

var _ Expression = &ExprIsOrIsNotNull{}

func (this *ExprIsOrIsNotNull) String() string {
	sql := NewStringBuilder().AppendStat(this.element).Append(" is ")
	if this.not != "" {
		sql.Append(this.not).Append(" ")
	}
	sql.Append(this.what)
	return sql.String()
}
