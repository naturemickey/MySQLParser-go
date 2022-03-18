package tree

type ElementText struct {
	text string
}

func (this *ElementText) Text() string {
	return this.text
}

func (this *ElementText) SetText(text string) {
	this.text = text
}

var _ ElementOpFactory = &ElementText{}

func (this *ElementText) String() string {
	return this.text
}
