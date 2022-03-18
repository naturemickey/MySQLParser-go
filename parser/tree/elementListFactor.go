package tree

type ElementListFactor struct {
}

var _ Element = &ElementListFactor{}

func (this *ElementListFactor) String() string {
	return ""
}
