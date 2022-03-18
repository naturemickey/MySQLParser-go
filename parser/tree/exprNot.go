package tree

type ExprNot struct {
}

var _ Expression = &ExprNot{}

func (this *ExprNot) String() string {
	return ""
}
