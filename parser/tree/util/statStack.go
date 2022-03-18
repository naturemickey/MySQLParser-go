package util

import "MySQLParser-go/parser/tree"

type StatStack struct {
	stats []tree.Stat
}

func (this *StatStack) Top() tree.Stat {
	if len(this.stats) > 0 {
		return this.stats[len(this.stats)-1]
	} else {
		return nil
	}
}

func (this *StatStack) Push(stat tree.Stat) {
	this.stats = append(this.stats, stat)
}

func (this *StatStack) Pop() tree.Stat {
	if len(this.stats) > 0 {
		stat := this.stats[len(this.stats)-1]
		this.stats = this.stats[:len(this.stats)-1]
		return stat
	} else {
		return nil
	}
}
