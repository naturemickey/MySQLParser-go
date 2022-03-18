package tree

type SetExpr struct {
}

var _ Stat = &SetExpr{}

func (this *SetExpr) String() string {
	return ""
}
