package tree

type DeleteStat struct {
}

var _ Stat = &DeleteStat{}

func (this *DeleteStat) String() string {
	return ""
}
