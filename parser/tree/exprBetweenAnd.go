package tree

type ExprBetweenAnd struct {
	el    Element
	not   string
	left  Element
	right Element
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
