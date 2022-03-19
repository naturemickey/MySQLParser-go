package tree

type Tables struct {
	BaseStat
	tableRels []TableRel
}

func (this *Tables) Assemble(stats []Stat) {
	for _, stat := range stats {
		this.tableRels = append(this.tableRels, stat.(TableRel))
	}
}

func (this *Tables) TableRels() []TableRel {
	return this.tableRels
}

func (this *Tables) SetTableRels(tableRels []TableRel) {
	this.tableRels = tableRels
}

var _ Stat = &Tables{}

func (this *Tables) String() string {
	sql := NewStringBuilder()
	for _, rel := range this.tableRels {
		sql.AppendStat(rel).Append(", ")
	}
	sql.deleteLast()
	return sql.String()
}
