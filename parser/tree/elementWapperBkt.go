package tree

type ElementWapperBkt struct {
}

var _ ElementOpFactory = &ElementWapperBkt{}

func (this *ElementWapperBkt) String() string {
	return ""
}
