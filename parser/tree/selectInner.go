package tree

type SelectInner struct {
	BaseStat
	selectPrefix *SelectPrefix
	selectSuffix *SelectSuffix
}

func (this *SelectInner) Children() []Stat {
	return Children(this)
}

func (this *SelectInner) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *SelectInner) SelectPrefix() *SelectPrefix {
	return this.selectPrefix
}

func (this *SelectInner) SetSelectPrefix(selectPrefix *SelectPrefix) {
	this.selectPrefix = selectPrefix
}

func (this *SelectInner) SelectSuffix() *SelectSuffix {
	return this.selectSuffix
}

func (this *SelectInner) SetSelectSuffix(selectSuffix *SelectSuffix) {
	this.selectSuffix = selectSuffix
}

var _ Stat = (*SelectInner)(nil)

func (this *SelectInner) String() string {
	prefix := this.selectPrefix.String()
	suffix := this.selectSuffix.String()
	//if suffix == "" {
	//	return prefix
	//}
	return prefix + suffix
}
