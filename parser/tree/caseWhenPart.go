package tree

type CaseWhenPart struct {
}

var _ Stat = &CaseWhenPart{}

func (this *CaseWhenPart) String() string {
	return ""
}
