package tree

type SelectStat struct {
	BaseStat
	selectInner       *SelectInner
	selectUnionSuffix *SelectUnionSuffix
}

var _ Stat = &SelectStat{}

func (this *SelectStat) SelectInner() *SelectInner {
	return this.selectInner
}

func (this *SelectStat) SetSelectInner(selectInner *SelectInner) {
	this.selectInner = selectInner
}

func (this *SelectStat) SelectUnionSuffix() *SelectUnionSuffix {
	return this.selectUnionSuffix
}

func (this *SelectStat) SetSelectUnionSuffix(selectUnionSuffix *SelectUnionSuffix) {
	this.selectUnionSuffix = selectUnionSuffix
}

func (this *SelectStat) String() string {
	if this.selectUnionSuffix == nil {
		return this.selectInner.String()
	} else {
		return "(" + this.selectInner.String() + ") " + this.selectUnionSuffix.String()
	}
}
