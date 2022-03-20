package tree

type Commit struct {
	BaseStat
}

func (this *Commit) Children() []Stat {
	return Children(this)
}

func (this *Commit) _TranscationStat() {
	//TODO implement me
	panic("implement me")
}

func (this *Commit) Assemble(stats []Stat) {
	Assemble(this, stats)
}

var _ TranscationStat = (*Commit)(nil)

func (this *Commit) String() string {
	return "commit"
}
