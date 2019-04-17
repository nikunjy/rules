package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nikunjy/rules/parser"
)

func main() {
	input := antlr.NewInputStream(`x.y.a eq 1 and x.y.b eq 2 and x.z eq "c"`)
	//input := antlr.NewInputStream(`x.y.a eq 1 and x.y.b eq 2`)
	lex := parser.NewJsonQueryLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewJsonQueryParser(tokens)
	items := map[string]interface{}{
		"name": "nikunj",
		"age":  10,
		"x": map[string]interface{}{
			"y": map[string]interface{}{
				"a": 1,
				"b": 2,
			},
			"z": "c",
		},
	}
	visitor := parser.NewJsonQueryVisitorImpl(items)
	tree := p.Query()
	fmt.Println(visitor.Visit(tree))
}
