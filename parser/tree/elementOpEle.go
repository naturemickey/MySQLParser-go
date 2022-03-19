package tree

type ElementOpEle struct {
	BaseStat
	elementOpFactory   ElementOpFactory
	elementOpEleSuffix *ElementOpEleSuffix
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

var _ Element = &ElementOpEle{}

func (this *ElementOpEle) String() string {
	sql := NewStringBuilder()
	sql.AppendStat(this.elementOpFactory)
	if this.elementOpEleSuffix != nil {
		sql.Append(" ").AppendStat(this.elementOpEleSuffix)
	}
	return sql.String()
}
