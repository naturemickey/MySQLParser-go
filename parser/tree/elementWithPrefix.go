package tree

type ElementWithPrefix struct {
	BaseStat
	prefix           string
	elementOpFactory ElementOpFactory
}

func (this *ElementWithPrefix) Children() []Stat {
	return Children(this)
}

func (this *ElementWithPrefix) _Element() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementWithPrefix) _ElementOpFactory() {
	//TODO implement me
	panic("implement me")
}

func (this *ElementWithPrefix) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *ElementWithPrefix) Prefix() string {
	return this.prefix
}

func (this *ElementWithPrefix) SetPrefix(prefix string) {
	this.prefix = prefix
}

func (this *ElementWithPrefix) ElementOpFactory() ElementOpFactory {
	return this.elementOpFactory
}

func (this *ElementWithPrefix) SetElementOpFactory(elementOpFactory ElementOpFactory) {
	this.elementOpFactory = elementOpFactory
}

var _ ElementOpFactory = (*ElementWithPrefix)(nil)

func (this *ElementWithPrefix) String() string {
	return this.prefix + " " + this.elementOpFactory.String()
}
