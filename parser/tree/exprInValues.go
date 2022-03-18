package tree

type ExprInValues struct {
}

var _ Expression = &ExprInValues{}

func (this *ExprInValues) String() string {
	return ""
}
