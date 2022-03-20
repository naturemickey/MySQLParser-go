package tree

type WhereCondition interface {
	Stat

	_WhereCondition()
}
