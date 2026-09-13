package parser

import (
	"fmt"
	"strings"

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
		if tok == nil || tok.GetTokenType() == antlr.TokenEOF {
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

// normalizeRuleInput keeps pretty-printed and Windows rules parseable against
// a grammar where SP is a single space and NEWLINE is only '\n':
// CRLF becomes '\n', runs of spaces/tabs outside quotes collapse to one space,
// and leading/trailing whitespace is trimmed. Quoted strings are unchanged.
func normalizeRuleInput(rule string) string {
	rule = strings.ReplaceAll(rule, "\r\n", "\n")
	rule = strings.ReplaceAll(rule, "\r", "\n")

	var b strings.Builder
	b.Grow(len(rule))
	inString := false
	escape := false
	lastWasSpace := false
	for i := 0; i < len(rule); i++ {
		c := rule[i]
		if inString {
			b.WriteByte(c)
			if escape {
				escape = false
			} else if c == '\\' {
				escape = true
			} else if c == '"' {
				inString = false
			}
			lastWasSpace = false
			continue
		}
		if c == '"' {
			inString = true
			b.WriteByte(c)
			lastWasSpace = false
			continue
		}
		if c == ' ' || c == '\t' {
			if !lastWasSpace {
				b.WriteByte(' ')
				lastWasSpace = true
			}
			continue
		}
		lastWasSpace = false
		b.WriteByte(c)
	}
	return strings.TrimSpace(b.String())
}
