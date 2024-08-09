package htmlparser

import (
	"github.com/Gabrieltrinidad0101/html-parser/lexer"
	"github.com/Gabrieltrinidad0101/html-parser/parser"
)

func Init() []*parser.Element {
	lexer_, err := lexer.NewLexer("/home/gabriel/Desktop/go/simpleWebBrowser/htmlParser/index.html")
	if err != nil {
		panic(err)
	}
	lexer_.Tokens()
	parser_ := parser.NewParser(lexer_.Targets)
	dom := parser_.Parser()
	tagets := dom.Children
	return tagets
}
