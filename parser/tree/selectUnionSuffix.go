package tree

type SelectUnionSuffix struct {
	BaseStat
	method       string
	selectStat   *SelectStat
	selectSuffix *SelectSuffix
}

func (this *SelectUnionSuffix) Method() string {
	return this.method
}

func (this *SelectUnionSuffix) SetMethod(method string) {
	this.method = method
}

func (this *SelectUnionSuffix) SelectStat() *SelectStat {
	return this.selectStat
}

func (this *SelectUnionSuffix) SetSelectStat(selectStat *SelectStat) {
	this.selectStat = selectStat
}

func (this *SelectUnionSuffix) SelectSuffix() *SelectSuffix {
	return this.selectSuffix
}

func (this *SelectUnionSuffix) SetSelectSuffix(selectSuffix *SelectSuffix) {
	this.selectSuffix = selectSuffix
}

var _ Stat = &SelectUnionSuffix{}

func (this *SelectUnionSuffix) String() string {
	sql := NewStringBuilder()
	sql.Append(" union ")
	if this.method != "" {
		sql.Append(this.method).Append(" ")
	}
	sql.Append("(").AppendStat(this.selectStat).Append(")").AppendStat(this.selectSuffix)
	return sql.String()
}
