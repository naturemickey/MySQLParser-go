package tree

type TableSubQuery struct {
}

var _ TableFactor = &TableSubQuery{}

func (this *TableSubQuery) String() string {
	return ""
}
