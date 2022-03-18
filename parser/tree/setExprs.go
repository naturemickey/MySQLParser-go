package tree

type SetExprs struct {
	setExprs []*SetExpr
}

func (this *SetExprs) SetExprs() []*SetExpr {
	return this.setExprs
}

func (this *SetExprs) SetSetExprs(setExprs []*SetExpr) {
	this.setExprs = setExprs
}

var _ Stat = &SetExprs{}

func (this *SetExprs) String() string {
	sql := NewStringBuilder()
	for _, expr := range this.setExprs {
		sql.AppendStat(expr).Append(", ")
	}
	sql.deleteLast()
	return sql.String()
}
