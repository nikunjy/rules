package parser

import (
	"strings"
)

// syntaxErrors are the errors ANTLR's parser and lexer report for the rule in
// ts, in the order they report them: a token recognition error comes before
// the first syntax error reported once the parse has read past it, and one
// the parse never reads up to is not reported. tree is the parse's tree.
func syntaxErrors(ts *JsonQueryTokens, tree *JsonQueryTree, parseErrs []JsonQueryParseError) []string {
	var errs []string
	lex := 0
	lexUpTo := func(read int) {
		for ; lex < len(ts.Errors) && int(ts.ErrorAt[lex]) <= read; lex++ {
			errs = append(errs, ts.Errors[lex].Error())
		}
	}
	for _, e := range parseErrs {
		lexUpTo(e.Read)
		errs = append(errs, e.Error())
	}
	lexUpTo(tree.Read)
	return errs
}

// trailingInput is the first token after the rule's input other than spaces
// and newlines, or EOF.
func trailingInput(tree *JsonQueryTree) int {
	i := tree.Next
	for {
		switch tree.Tokens[i].Type {
		case JsonQuerySP, JsonQueryNEWLINE:
			i++
		default:
			return i
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
