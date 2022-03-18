package tree

type ElementSubQuery struct {
}

var _ ElementOpFactory = &ElementSubQuery{}

func (this *ElementSubQuery) String() string {
	return ""
}
