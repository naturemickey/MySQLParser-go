package tree

type ExprLike struct {
}

var _ Expression = &ExprLike{}

func (this *ExprLike) String() string {
	return ""
}
