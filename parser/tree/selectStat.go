package tree

type SelectStat struct {
	BaseStat
	selectInner       *SelectInner
	selectUnionSuffix *SelectUnionSuffix
}

func (this *SelectStat) Children() []Stat {
	return Children(this)
}

func (this *SelectStat) Assemble(stats []Stat) {
	Assemble(this, stats)
}

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

var _ Stat = (*SelectStat)(nil)

func (this *SelectStat) String() string {
	if this.selectUnionSuffix == nil {
		return this.selectInner.String()
	} else {
		return "(" + this.selectInner.String() + ") " + this.selectUnionSuffix.String()
	}
}
