package main

import (
	"MySQLParser-go/parser"
)

func main() {
	println(parser.ParseSQL("select * from tab").String())
}
