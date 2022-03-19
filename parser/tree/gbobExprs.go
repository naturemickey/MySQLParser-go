package tree

type GbobExprs struct {
	BaseStat
	element   Element
	sc        string
	gbobExprs *GbobExprs
}

func (this *GbobExprs) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *GbobExprs) Element() Element {
	return this.element
}

func (this *GbobExprs) SetElement(element Element) {
	this.element = element
}

func (this *GbobExprs) Sc() string {
	return this.sc
}

func (this *GbobExprs) SetSc(sc string) {
	this.sc = sc
}

func (this *GbobExprs) GbobExprs() *GbobExprs {
	return this.gbobExprs
}

func (this *GbobExprs) SetGbobExprs(gbobExprs *GbobExprs) {
	this.gbobExprs = gbobExprs
}

var _ Stat = &GbobExprs{}

func (this *GbobExprs) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.element)
	if this.sc != "" {
		sql.Append(" ").Append(this.sc)
	}
	if this.gbobExprs != nil {
		sql.Append(", ").AppendStat(this.gbobExprs)
	}
	return sql.String()
}
