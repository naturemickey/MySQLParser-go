package tree

type ElementText struct {
	BaseStat
	text string
}

func (this *ElementText) Children() []Stat {
	return Children(this)
}

func (this *ElementText) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementText) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementText) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementText) Text() string {
	return this.text
}

func (this *ElementText) SetText(text string) {
	this.text = text
}

var _ ElementOpFactory = (*ElementText)(nil)

func (this *ElementText) String() string {
	return this.text
}
