package tree

type SelectUnionSuffix struct {
}

var _ Stat = &SelectUnionSuffix{}

func (this *SelectUnionSuffix) String() string {
	return ""
}
