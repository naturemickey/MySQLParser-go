package tree

import "strings"

type ColumnNames struct {
	BaseStat
	names []string
}

func (this *ColumnNames) Children() []Stat {
	return Children(this)
}

func (this *ColumnNames) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ColumnNames) Names() []string {
	return this.names
}

func (this *ColumnNames) SetNames(names []string) {
	this.names = names
}

var _ Stat = (*ColumnNames)(nil)

func (this *ColumnNames) String() string {
	return strings.Join(this.names, ", ")
}
