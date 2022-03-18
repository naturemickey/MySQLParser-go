package tree

type ExprInSelect struct {
}

var _ Expression = &ExprInSelect{}

func (this *ExprInSelect) String() string {
	return ""
}
