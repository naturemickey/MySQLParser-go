package tree

import "strings"

type StringBuilder struct {
	strings []string
}

func NewStringBuilder() *StringBuilder {
	return new(StringBuilder)
}

func (this *StringBuilder) Append(s string) *StringBuilder {
	this.strings = append(this.strings, s)
	return this
}

func (this *StringBuilder) AppendStat(stat Stat) *StringBuilder {
	this.strings = append(this.strings, stat.String())
	return this
}

func (this *StringBuilder) String() string {
	return strings.Join(this.strings, "")
}

func (this *StringBuilder) deleteLast() *StringBuilder {
	this.strings = this.strings[:len(this.strings)-1]
	return this
}
