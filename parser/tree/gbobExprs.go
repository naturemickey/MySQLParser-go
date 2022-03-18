package tree

type GbobExprs struct {
}

var _ Stat = &GbobExprs{}

func (this *GbobExprs) String() string {
	return ""
}
