package tree

type TableJoinMod struct {
	BaseStat
	mod string
}

func (this *TableJoinMod) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *TableJoinMod) Mod() string {
	return this.mod
}

func (this *TableJoinMod) SetMod(mod string) {
	this.mod = mod
}

var _ Stat = &TableJoinMod{}

func (this *TableJoinMod) String() string {
	return this.mod
}
