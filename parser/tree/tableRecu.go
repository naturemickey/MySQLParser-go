package tree

type TableRecu struct {
	BaseStat
	tableRel TableRel
}

func (this *TableRecu) Children() []Stat {
	return Children(this)
}

func (this *TableRecu) _TableRel() {
	//TODO implement me
	panic("implement me")
}

func (this *TableRecu) _TableFactor() {
	//TODO implement me
	panic("implement me")
}

func (this *TableRecu) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *TableRecu) TableRel() TableRel {
	return this.tableRel
}

func (this *TableRecu) SetTableRel(tableRel TableRel) {
	this.tableRel = tableRel
}

var _ TableFactor = (*TableRecu)(nil)

func (this *TableRecu) String() string {
	return NewStringBuilder().Append("(").AppendStat(this.tableRel).Append(")").String()
}
