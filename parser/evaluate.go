package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Evaluate(rule string, items map[string]interface{}) bool {
	input := antlr.NewInputStream(rule)
	lex := NewJsonQueryLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := NewJsonQueryParser(tokens)
	visitor := NewJsonQueryVisitorImpl(items)
	tree := p.Query()
	result := visitor.Visit(tree)
	if result == nil {
		return false
	}
	return result.(bool)
}
