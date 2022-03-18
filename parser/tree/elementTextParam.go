package tree

type ElementTextParam struct {
	ElementText
}

var _ ElementOpFactory = &ElementTextParam{}

func (this *ElementTextParam) String() string {
	return ""
}
