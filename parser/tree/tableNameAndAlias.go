package tree

type TableNameAndAlias struct {
}

var _ TableFactor = &TableNameAndAlias{}

func (this *TableNameAndAlias) String() string {
	return ""
}
