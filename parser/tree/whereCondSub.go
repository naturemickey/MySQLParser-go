package tree

type WhereCondSub struct {
}

var _ WhereCondition = &WhereCondSub{}

func (this *WhereCondSub) String() string {
	return ""
}
