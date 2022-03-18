package tree

type StatStack struct {
	stats []Stat
}

func NewStatStack() *StatStack {
	return new(StatStack)
}

func (this *StatStack) Top() Stat {
	if len(this.stats) > 0 {
		return this.stats[len(this.stats)-1]
	} else {
		return nil
	}
}

func (this *StatStack) Push(stat Stat) {
	this.stats = append(this.stats, stat)
}

func (this *StatStack) Pop() Stat {
	if len(this.stats) > 0 {
		stat := this.stats[len(this.stats)-1]
		this.stats = this.stats[:len(this.stats)-1]
		return stat
	} else {
		return nil
	}
}
