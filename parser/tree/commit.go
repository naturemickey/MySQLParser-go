package tree

type Commit struct {
	BaseStat
}

func (this *Commit) _TranscationStat() {
	//TODO implement me
	panic("implement me")
}

func (this *Commit) Assemble(stats []Stat) {
	Assemble(this, stats)
}

var _ TranscationStat = &Commit{}

func (this *Commit) String() string {
	return "commit"
}
