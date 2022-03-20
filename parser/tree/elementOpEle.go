package tree

type ElementOpEle struct {
	BaseStat
	elementOpFactory   ElementOpFactory
	elementOpEleSuffix *ElementOpEleSuffix
}

func (this *ElementOpEle) Children() []Stat {
	return Children(this)
}

func (this *ElementOpEle) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementOpEle) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementOpEle) ElementOpFactory() ElementOpFactory {
	return this.elementOpFactory
}

func (this *ElementOpEle) SetElementOpFactory(elementOpFactory ElementOpFactory) {
	this.elementOpFactory = elementOpFactory
}

func (this *ElementOpEle) ElementOpEleSuffix() *ElementOpEleSuffix {
	return this.elementOpEleSuffix
}

func (this *ElementOpEle) SetElementOpEleSuffix(elementOpEleSuffix *ElementOpEleSuffix) {
	this.elementOpEleSuffix = elementOpEleSuffix
}

var _ Element = (*ElementOpEle)(nil)

func (this *ElementOpEle) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.elementOpFactory)
	if this.elementOpEleSuffix != nil {
		sql.Append(" ").AppendStat(this.elementOpEleSuffix)
	}
	return sql.String()
}
