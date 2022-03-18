package tree

type Commit struct {
}

var _ TranscationStat = &Commit{}

func (this *Commit) String() string {
	return "commit"
}
