package tree

type ElementTextParam struct {
}

var _ ElementOpFactory = &ElementTextParam{}

func (this *ElementTextParam) String() string {
	return ""
}
