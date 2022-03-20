package tree

type Rollback struct {
	BaseStat
}

func (this *Rollback) _TranscationStat() {
	//TODO implement me
	panic("implement me")
}

func (this *Rollback) Assemble(stats []Stat) {
	Assemble(this, stats)
}

var _ TranscationStat = &Rollback{}

func (this *Rollback) String() string {
	return "rollback"
}
