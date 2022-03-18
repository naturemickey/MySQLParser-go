package tree

type SelectSuffix struct {
}

var _ Stat = &SelectSuffix{}

func (this *SelectSuffix) String() string {
	return ""
}
