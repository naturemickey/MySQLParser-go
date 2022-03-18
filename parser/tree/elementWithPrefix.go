package tree

type ElementWithPrefix struct {
	prefix           string
	elementOpFactory ElementOpFactory
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

var _ ElementOpFactory = &ElementWithPrefix{}

func (this *ElementWithPrefix) String() string {
	return this.prefix + " " + this.elementOpFactory.String()
}
