package tree

type TableRecu struct {
}

var _ TableFactor = &TableRecu{}

func (this *TableRecu) String() string {
	return ""
}
