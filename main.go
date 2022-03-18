package main

import (
	"MySQLParser-go/parser"
	"MySQLParser-go/parser/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//type MyMySQLListener struct {
//	*antlr4.BaseMySQLListener
//}
//
//var _ antlr4.MySQLListener = &MyMySQLListener{}
//
//func NewMySQLListener() *MyMySQLListener {
//	return new(MyMySQLListener)
//}
//
//func (this *MyMySQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	fmt.Println(ctx.GetText())
//}

func main() {
	input := antlr.NewInputStream("select * from tab")
	lexer := antlr4.NewMySQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := antlr4.NewMySQLParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// p.BuildParseTrees = true
	tree := p.Stat()
	antlr.ParseTreeWalkerDefault.Walk(parser.NewMySQLListener(), tree)
}
