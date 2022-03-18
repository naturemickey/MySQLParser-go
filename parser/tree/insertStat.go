package tree

type InsertStat struct {
}

var _ Stat = &InsertStat{}

func (this *InsertStat) String() string {
	return ""
}
