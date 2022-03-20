package tree

type TableJoinMod struct {
	BaseStat
	mod string
}

func (this *TableJoinMod) Children() []Stat {
	return Children(this)
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

var _ Stat = (*TableJoinMod)(nil)

func (this *TableJoinMod) String() string {
	return this.mod
}
