package tree

type SelectInner struct {
	BaseStat
	selectPrefix *SelectPrefix
	selectSuffix *SelectSuffix
}

var _ Stat = &SelectInner{}

func (this *SelectInner) String() string {
	prefix := this.selectPrefix.String()
	suffix := this.selectSuffix.String()
	//if suffix == "" {
	//	return prefix
	//}
	return prefix + suffix
}
