package main

import (
	"MySQLParser-go/parser"
)

func main() {
	stat := parser.ParseSQL("insert into aaa(a,b,c) values(?,?,?)")
	println(stat.String())

	for _, s := range stat.Children() {
		if s != nil {
			println(s.String())
		}
	}
}
