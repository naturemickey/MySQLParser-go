package tree

type ExprIsOrIsNotNull struct {
}

var _ Expression = &ExprIsOrIsNotNull{}

func (this *ExprIsOrIsNotNull) String() string {
	return ""
}
