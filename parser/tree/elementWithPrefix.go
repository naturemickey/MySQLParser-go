package tree

type ElementWithPrefix struct {
}

var _ ElementOpFactory = &ElementWithPrefix{}

func (this *ElementWithPrefix) String() string {
	return ""
}
