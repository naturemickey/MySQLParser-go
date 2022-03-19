package tree

// todo 处理elementOpEle中包含op的情况
type ElementOpEleSuffix struct {
	BaseStat
	op           string
	elementOpEle *ElementOpEle
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

var _ Stat = &ElementOpEleSuffix{}

func (this *ElementOpEleSuffix) String() string {
	sql := NewStringBuilder()
	if this.op != "" {
		sql.Append(this.op).Append(" ")
	}
	sql.AppendStat(this.elementOpEle)
	return sql.String()
}
