package tree

type ExprBetweenAnd struct {
}

var _ Expression = &ExprBetweenAnd{}

func (this *ExprBetweenAnd) String() string {
	return ""
}
