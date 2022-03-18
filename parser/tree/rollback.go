package tree

type Rollback struct {
}

var _ TranscationStat = &Rollback{}

func (this *Rollback) String() string {
	return "rollback"
}
