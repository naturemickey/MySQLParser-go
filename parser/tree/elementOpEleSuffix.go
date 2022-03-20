package tree

type ElementOpEleSuffix struct {
	BaseStat
	op           string
	elementOpEle *ElementOpEle
}

func (this *ElementOpEleSuffix) Children() []Stat {
	return Children(this)
}

func (this *ElementOpEleSuffix) Assemble(stats []Stat) {
	if this.op == "" {
		// 处理elementOpEle中包含op的情况
		element := stats[0].(*ElementOpEle).ElementOpFactory().(*ElementTextParam)
		text := []rune(element.Text())
		this.op = string([]rune{text[0]})
		text = text[1:]
		element.SetText(string(text))
	}
	Assemble(this, stats)
}

func (this *ElementOpEleSuffix) Op() string {
	return this.op
}

func (this *ElementOpEleSuffix) SetOp(op string) {
	this.op = op
}

func (this *ElementOpEleSuffix) ElementOpEle() *ElementOpEle {
	return this.elementOpEle
}

func (this *ElementOpEleSuffix) SetElementOpEle(elementOpEle *ElementOpEle) {
	this.elementOpEle = elementOpEle
}

var _ Stat = (*ElementOpEleSuffix)(nil)

func (this *ElementOpEleSuffix) String() string {
	sql := NewStringBuilder()
	if this.op != "" {
		sql.Append(this.op).Append(" ")
	}
	sql.AppendStat(this.elementOpEle)
	return sql.String()
}
