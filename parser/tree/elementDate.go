package tree

type ElementDate struct {
	BaseStat
	dt  string
	str string
}

func (this *ElementDate) Dt() string {
	return this.dt
}

func (this *ElementDate) SetDt(dt string) {
	this.dt = dt
}

func (this *ElementDate) Str() string {
	return this.str
}

func (this *ElementDate) SetStr(str string) {
	this.str = str
}

var _ ElementOpFactory = &ElementDate{}

func (this *ElementDate) String() string {
	return NewStringBuilder().Append(this.dt).Append(" ").Append(this.str).String()
}
