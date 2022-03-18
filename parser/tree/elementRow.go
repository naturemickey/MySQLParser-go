package tree

type ElementRow struct {
}

var _ ElementOpFactory = &ElementRow{}

func (this *ElementRow) String() string {
	return ""
}
