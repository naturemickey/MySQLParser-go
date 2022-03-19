package tree

type Rollback struct {
	BaseStat
}

var _ TranscationStat = &Rollback{}

func (this *Rollback) String() string {
	return "rollback"
}
