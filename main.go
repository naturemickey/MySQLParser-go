package main

import (
	"MySQLParser-go/parser"
	"MySQLParser-go/parser/tree"
)

func main() {
	stat := parser.ParseSQL("select * from tab")
	println(stat.String())

	for _, s := range tree.Children(stat) {
		if s != nil {
			println(s.String())
		}
	}
}
