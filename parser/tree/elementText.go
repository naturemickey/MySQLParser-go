package tree

type ElementText struct {
}

var _ ElementOpFactory = &ElementText{}

func (this *ElementText) String() string {
	return ""
}
