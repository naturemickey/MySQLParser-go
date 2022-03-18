package tree

type ElementCase struct {
}

var _ ElementOpFactory = &ElementCase{}

func (this *ElementCase) String() string {
	return ""
}
