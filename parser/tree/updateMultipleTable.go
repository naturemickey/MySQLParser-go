package tree

type UpdateMultipleTable struct {
}

var _ UpdateStat = &UpdateMultipleTable{}

func (this *UpdateMultipleTable) String() string {
	return ""
}
