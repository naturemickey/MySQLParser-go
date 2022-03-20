package tree

type ElementRow struct {
	BaseStat
	paramList *ParamList
}

func (this *ElementRow) Children() []Stat {
	return Children(this)
}

func (this *ElementRow) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementRow) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementRow) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementRow) ParamList() *ParamList {
	return this.paramList
}

func (this *ElementRow) SetParamList(paramList *ParamList) {
	this.paramList = paramList
}

var _ ElementOpFactory = (*ElementRow)(nil)

func (this *ElementRow) String() string {
	return "row(" + this.paramList.String() + ")"
}
