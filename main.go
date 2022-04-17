package main

import (
	"MySQLParser-go/parser"
)

func main() {
	stat := parser.ParseSQL("select * from tab a left outer join tab b on a.x = b.x")
	// update aaaa set x = ?, v = 1 where id = ? and v = 1
	// insert tab(a, b) values(?, ?) -->  insert tab(a, b, v) values(?, ? ,1)
	// select * from tab a left outer join tab b on a.x = b.x
	// select * from tab a left outer join tab b on a.x = b.x where a.v = 1 and (b.v = 1 or b.v is null)
	// select * from tab a left outer join tab b on a.x = b.x
	// select tab from tab_p a left outer join tab_p b on a.x = b.x
	println(stat.String())

	for _, s := range stat.Children() {
		if s != nil {
			println(s.String())
		}
	}
}
