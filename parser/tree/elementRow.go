package tree

type ElementRow struct {
	BaseStat
	paramList *ParamList
}

func (this *ElementRow) ParamList() *ParamList {
	return this.paramList
}

func (this *ElementRow) SetParamList(paramList *ParamList) {
	this.paramList = paramList
}

var _ ElementOpFactory = &ElementRow{}

func (this *ElementRow) String() string {
	return "row(" + this.paramList.String() + ")"
}
