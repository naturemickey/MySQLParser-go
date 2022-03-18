package tree

type SelectPrefix struct {
	distinct bool
}

var _ Stat = &SelectPrefix{}

func (this *SelectPrefix) String() string {
	return ""
}
