package parser

import (
	"MySQLParser-go/parser/antlr4"
	"MySQLParser-go/parser/tree"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseSQL(sql string) tree.Stat {
	input := antlr.NewInputStream(sql)
	lexer := antlr4.NewMySQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := antlr4.NewMySQLParser(stream)
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	// p.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	// p.BuildParseTrees = true
	tree := p.Stat()
	listener := NewMySQLListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.Result()
}
