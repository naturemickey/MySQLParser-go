package tree

type TableNameAndAlias struct {
	BaseStat
	name  string
	alias string
}

func (this *TableNameAndAlias) _TableRel() {
	//TODO implement me
	panic("implement me")
}

func (this *TableNameAndAlias) _TableFactor() {
	//TODO implement me
	panic("implement me")
}

func (this *TableNameAndAlias) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *TableNameAndAlias) Name() string {
	return this.name
}

func (this *TableNameAndAlias) SetName(name string) {
	this.name = name
}

func (this *TableNameAndAlias) Alias() string {
	return this.alias
}

func (this *TableNameAndAlias) SetAlias(alias string) {
	this.alias = alias
}

var _ TableFactor = &TableNameAndAlias{}

func (this *TableNameAndAlias) String() string {
	if this.alias == "" {
		return this.name
	}
	return this.name + " " + this.alias
}
