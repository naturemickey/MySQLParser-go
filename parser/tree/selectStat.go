package tree

type SelectStat struct {
	selectInner       *SelectInner
	selectUnionSuffix *SelectUnionSuffix
}

func (this *SelectStat) String() string {
	if this.selectUnionSuffix == nil {
		return this.selectInner.String()
	} else {
		return "(" + this.selectInner.String() + ") " + this.selectUnionSuffix.String()
	}
}

func (this *SelectStat) SetSelectInner(selectInner *SelectInner) {
	this.selectInner = selectInner
}
