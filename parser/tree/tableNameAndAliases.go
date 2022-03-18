package tree

type TableNameAndAliases struct {
}

var _ Stat = &TableNameAndAliases{}

func (this *TableNameAndAliases) String() string {
	return ""
}
