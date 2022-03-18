package tree

type TableJoinMod struct {
}

var _ Stat = &TableJoinMod{}

func (this *TableJoinMod) String() string {
	return ""
}
