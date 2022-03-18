package tree

type Tables struct {
}

var _ Stat = &Tables{}

func (this *Tables) String() string {
	return ""
}
