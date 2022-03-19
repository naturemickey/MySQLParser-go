package tree

type StatStack struct {
	stats []*statAndStatus
}

type StatStatus int8

const (
	_materiel StatStatus = 0
	_product  StatStatus = 1
)

type statAndStatus struct {
	stat   Stat
	status StatStatus
}

func NewStatStack() *StatStack {
	return new(StatStack)
}

func (this *StatStack) PushMateriel(stat Stat) {
	this.stats = append(this.stats, &statAndStatus{stat, _materiel})
}

func (this *StatStack) MakeProduct() {
	var product *statAndStatus
	i := len(this.stats) - 1
	for ; i >= 0; i-- {
		s := this.stats[i]
		if s.status == _materiel {
			product = s
			break
		}
	}
	var ss []Stat
	for ; i < len(this.stats); i++ {
		ss = append(ss, this.stats[i].stat)
	}
	product.stat.Assemble(ss)
	product.status = _product
}

func (this *StatStack) Head() Stat {
	return this.stats[0].stat
}
