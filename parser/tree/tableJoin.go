package tree

type TableJoin struct {
}

var _ TableRel = &TableJoin{}

func (this *TableJoin) String() string {
	return ""
}
