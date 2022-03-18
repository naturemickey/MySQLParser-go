package tree

type ExprRelational struct {
}

var _ Expression = &ExprRelational{}

func (this *ExprRelational) String() string {
	return ""
}
