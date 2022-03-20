package tree

type ElementTextParam struct {
	ElementText
}

var _ ElementOpFactory = (*ElementTextParam)(nil)
