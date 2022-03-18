package tree

type SelectInner struct {
	selectPrefix *SelectPrefix
	selectSuffix *SelectSuffix
}

var _ Stat = &SelectInner{}

func (this *SelectInner) String() string {
	return this.selectPrefix.String() + " " + this.selectSuffix.String()
}
