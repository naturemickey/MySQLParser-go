package tree

type FunCall struct {
}

var _ ElementOpFactory = &FunCall{}

func (this *FunCall) String() string {
	return ""
}
