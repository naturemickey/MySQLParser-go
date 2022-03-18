package tree

type SelectExprs struct {
}

var _ Stat = &SelectExprs{}

func (this *SelectExprs) String() string {
	return ""
}
