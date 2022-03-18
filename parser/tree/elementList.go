package tree

type ElementList struct {
}

var _ Stat = &ElementList{}

func (this *ElementList) String() string {
	return ""
}
