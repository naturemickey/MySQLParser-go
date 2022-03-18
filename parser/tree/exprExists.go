package tree

type ExprExists struct {
}

var _ Expression = &ExprExists{}

func (this *ExprExists) String() string {
	return ""
}
