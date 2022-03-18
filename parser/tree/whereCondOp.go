package tree

type WhereCondOp struct {
}

var _ WhereCondition = &WhereCondOp{}

func (this *WhereCondOp) String() string {
	return ""
}
