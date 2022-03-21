package tree

type Expression interface {
	Stat

	_Expression()
}
