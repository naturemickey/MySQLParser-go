package tree

type SetExprs struct {
}

var _ Stat = &SetExprs{}

func (this *SetExprs) String() string {
	return ""
}
