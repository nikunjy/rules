package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// collectingErrorListener records ANTLR syntax errors instead of writing them
// to stderr. NewEvaluator uses this so invalid rules fail at construction time.
type collectingErrorListener struct {
	*antlr.DefaultErrorListener
	errs []string
}

func (l *collectingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func consumeTrailingWhitespace(tokens *antlr.CommonTokenStream) {
	for {
		tok := tokens.LT(1)
		if tok == nil {
			return
		}
		switch tok.GetTokenType() {
		case JsonQueryLexerSP, JsonQueryLexerNEWLINE:
			tokens.Consume()
		default:
			return
		}
	}
}
