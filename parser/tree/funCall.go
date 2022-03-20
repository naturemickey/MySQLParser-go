package tree

type FunCall struct {
	BaseStat
	funName   string
	paramList *ParamList
}

func (this *FunCall) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *FunCall) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *FunCall) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *FunCall) FunName() string {
	return this.funName
}

func (this *FunCall) SetFunName(funName string) {
	this.funName = funName
}

func (this *FunCall) ParamList() *ParamList {
	return this.paramList
}

func (this *FunCall) SetParamList(paramList *ParamList) {
	this.paramList = paramList
}

var _ ElementOpFactory = &FunCall{}

func (this *FunCall) String() string {
	sql := NewStringBuilder().Append(this.funName).Append("(")
	if this.paramList != nil {
		sql.AppendStat(this.paramList)
	}
	sql.Append(")")
	return sql.String()
}
