package tree

type UpdateSingleTable struct {
}

var _ UpdateStat = &UpdateSingleTable{}

func (this *UpdateSingleTable) String() string {
	return ""
}
