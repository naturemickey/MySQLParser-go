package tree

type TableRecu struct {
	BaseStat
	tableRel TableRel
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

var _ TableFactor = &TableRecu{}

func (this *TableRecu) String() string {
	return NewStringBuilder().Append("(").AppendStat(this.tableRel).Append(")").String()
}
