package tree

type TableRecu struct {
	tableRel TableRel
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
