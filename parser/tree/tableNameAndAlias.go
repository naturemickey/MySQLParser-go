package tree

type TableNameAndAlias struct {
	name  string
	alias string
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
	return this.name + " " + this.alias
}
