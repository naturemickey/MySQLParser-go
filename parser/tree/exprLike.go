package tree

type ExprLike struct {
	left  Element
	right Element
	not   string
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
