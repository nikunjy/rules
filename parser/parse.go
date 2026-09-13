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

// normalizeLogicalKeywords rewrites standalone AND/OR (any case) to and/or so
// the generated lexer, which only accepts lowercase logical operators, still
// honors the documented capitalization. The grammar is left listing only
// 'and' | 'or': adding 'AND' | 'OR' there without regenerating the lexer
// drifts source from generated code, and regenerating those literals would
// reserve AND/OR as keywords in attribute positions such as x.AND.
//
// Quoted strings and attribute names are left unchanged, including ANDROID,
// order, dotted segments (x.AND), and AND/OR used as an attrPath at the start
// of a query (AND eq 1).
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
	// '.' separates attrPath segments and is not ATTR_NAME_CHAR, so x.AND / AND.x
	// must not be treated as logical operators.
	if i > 0 && rule[i-1] == '.' {
		return false
	}
	if end < len(rule) && rule[end] == '.' {
		return false
	}
	// Logical operators are infix. AND/OR at the start of a query is an identifier.
	if atQueryStart(rule, i) {
		return false
	}
	return true
}

func atQueryStart(rule string, i int) bool {
	j := i - 1
	for j >= 0 && (rule[j] == ' ' || rule[j] == '\n') {
		j--
	}
	if j < 0 || rule[j] == '(' {
		return true
	}
	// After a logical operator the next query starts (`x eq 1 or AND eq 2`).
	end := j + 1
	start := end
	for start > 0 && isAttrNameChar(rune(rule[start-1])) {
		start--
	}
	word := rule[start:end]
	return strings.EqualFold(word, "and") || strings.EqualFold(word, "or")
}

func isAttrNameChar(r rune) bool {
	return r == '-' || r == '_' || r == ':' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
