package tree

type SetExpr struct {
	left         Element
	right        Element
	rightDefault string
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
