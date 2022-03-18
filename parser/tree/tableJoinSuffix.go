package tree

type TableJoinSuffix struct {
}

var _ Stat = &TableJoinSuffix{}

func (this *TableJoinSuffix) String() string {
	return ""
}
