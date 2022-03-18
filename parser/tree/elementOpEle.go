package tree

type ElementOpEle struct {
}

var _ Element = &ElementOpEle{}

func (this *ElementOpEle) String() string {
	return ""
}
