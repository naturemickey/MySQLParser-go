package tree

type SelectSuffix struct {
	BaseStat
	orderByExprs *GbobExprs
	offset       string
	rowCount     string
	hasOffSet    bool
	hasLimit     bool
	lock         string
}

func (this *SelectSuffix) Children() []Stat {
	return Children(this)
}

func (this *SelectSuffix) Assemble(stats []Stat) {
	Assemble(this, stats)
}

func (this *SelectSuffix) HasLimit() bool {
	return this.hasLimit
}

func (this *SelectSuffix) SetHasLimit(hasLimit bool) {
	this.hasLimit = hasLimit
}

func (this *SelectSuffix) OrderByExprs() *GbobExprs {
	return this.orderByExprs
}

func (this *SelectSuffix) SetOrderByExprs(orderByExprs *GbobExprs) {
	this.orderByExprs = orderByExprs
}

func (this *SelectSuffix) Offset() string {
	return this.offset
}

func (this *SelectSuffix) SetOffset(offset string) {
	this.offset = offset
}

func (this *SelectSuffix) RowCount() string {
	return this.rowCount
}

func (this *SelectSuffix) SetRowCount(rowCount string) {
	this.rowCount = rowCount
}

func (this *SelectSuffix) HasOffSet() bool {
	return this.hasOffSet
}

func (this *SelectSuffix) SetHasOffSet(hasOffSet bool) {
	this.hasOffSet = hasOffSet
}

func (this *SelectSuffix) Lock() string {
	return this.lock
}

func (this *SelectSuffix) SetLock(lock string) {
	this.lock = lock
}

var _ Stat = (*SelectSuffix)(nil)

func (this *SelectSuffix) String() string {
	sql := NewStringBuilder()
	if this.orderByExprs != nil {
		sql.Append(" order by ").AppendStat(this.orderByExprs)
	}
	if this.hasLimit {
		sql.Append(" limit ")
		if this.hasOffSet {
			sql.Append(this.rowCount).Append(" offset ").Append(this.offset)
		} else {
			if this.offset == "" {
				sql.Append(this.rowCount)
			} else {
				sql.Append(this.offset).Append(", ").Append(this.rowCount)
			}
		}
	}
	if this.lock != "" {
		sql.Append(" ").Append(this.lock)
	}
	return sql.String()
}
