package tree

type CaseWhenPart struct {
	BaseStat
	whenEl Element
	whenRe *ExprRelational
	then   Element
	suffix *CaseWhenPart
}

func (this *CaseWhenPart) Children() []Stat {
	return Children(this)
}

func (this *CaseWhenPart) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *CaseWhenPart) WhenEl() Element {
	return this.whenEl
}

func (this *CaseWhenPart) SetWhenEl(whenEl Element) {
	this.whenEl = whenEl
}

func (this *CaseWhenPart) WhenRe() *ExprRelational {
	return this.whenRe
}

func (this *CaseWhenPart) SetWhenRe(whenRe *ExprRelational) {
	this.whenRe = whenRe
}

func (this *CaseWhenPart) Then() Element {
	return this.then
}

func (this *CaseWhenPart) SetThen(then Element) {
	this.then = then
}

func (this *CaseWhenPart) Suffix() *CaseWhenPart {
	return this.suffix
}

func (this *CaseWhenPart) SetSuffix(suffix *CaseWhenPart) {
	this.suffix = suffix
}

var _ Stat = (*CaseWhenPart)(nil)

func (this *CaseWhenPart) String() string {
	sql := NewStringBuilder().Append("when ")
	if this.whenEl != nil {
		sql.AppendStat(this.whenEl)
	} else {
		sql.AppendStat(this.whenRe)
	}
	sql.Append(" then ").AppendStat(this.then)
	if this.suffix != nil {
		sql.Append(" ").AppendStat(this.suffix)
	}
	return sql.String()
}
