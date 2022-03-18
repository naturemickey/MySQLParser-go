package tree

type ElementDate struct {
}

var _ ElementOpFactory = &ElementDate{}

func (this *ElementDate) String() string {
	return ""
}
