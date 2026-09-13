package parser

import (
	"fmt"
	"strings"
	"unicode"

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

// normalizeLogicalKeywords rewrites standalone AND/OR (any case) to and/or so
// the generated lexer, which only accepts lowercase logical operators until
// the grammar is regenerated, still honors the documented capitalization.
// Tokens inside quoted strings and attribute names (ANDROID, order) are left
// unchanged.
func normalizeLogicalKeywords(rule string) string {
	var b strings.Builder
	b.Grow(len(rule))
	inString := false
	escape := false
	i := 0
	for i < len(rule) {
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
			i++
			continue
		}
		if c == '"' {
			inString = true
			b.WriteByte(c)
			i++
			continue
		}
		if matchStandaloneKeyword(rule, i, "and") {
			b.WriteString("and")
			i += 3
			continue
		}
		if matchStandaloneKeyword(rule, i, "or") {
			b.WriteString("or")
			i += 2
			continue
		}
		b.WriteByte(c)
		i++
	}
	return b.String()
}

func matchStandaloneKeyword(rule string, i int, keyword string) bool {
	end := i + len(keyword)
	if end > len(rule) {
		return false
	}
	if !strings.EqualFold(rule[i:end], keyword) {
		return false
	}
	if i > 0 && isAttrNameChar(rune(rule[i-1])) {
		return false
	}
	if end < len(rule) && isAttrNameChar(rune(rule[end])) {
		return false
	}
	return true
}

func isAttrNameChar(r rune) bool {
	return r == '-' || r == '_' || r == ':' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
