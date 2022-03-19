package tree

import "strings"

type ColumnNames struct {
	BaseStat
	names []string
}

func (this *ColumnNames) Names() []string {
	return this.names
}

func (this *ColumnNames) SetNames(names []string) {
	this.names = names
}

var _ Stat = &ColumnNames{}

func (this *ColumnNames) String() string {
	return strings.Join(this.names, ", ")
}
