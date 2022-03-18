package tree

type ParamList struct {
}

var _ Stat = &ParamList{}

func (this *ParamList) String() string {
	return ""
}
