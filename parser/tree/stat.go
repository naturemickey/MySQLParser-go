package tree

type Stat interface {
	String() string
	Assemble([]Stat)
}
