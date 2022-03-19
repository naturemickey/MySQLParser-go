package tree

type Commit struct {
	BaseStat
}

var _ TranscationStat = &Commit{}

func (this *Commit) String() string {
	return "commit"
}
