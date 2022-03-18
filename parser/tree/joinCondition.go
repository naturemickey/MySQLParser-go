package tree

type JoinCondition struct {
}

var _ Stat = &JoinCondition{}

func (this *JoinCondition) String() string {
	return ""
}
