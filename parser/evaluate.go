package parser

import (
	"fmt"
	"strings"
)

type Evaluator struct {
	lastDebugErr error

	rule string
	tree QueryNode

	testHookPanic func()
}

func NewEvaluator(rule string) (ret *Evaluator, retErr error) {
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
		}
	}()
	ts := JsonQueryLex(normalizeRuleInput(rule))
	tree, parseErrs := ParseQueryWithRecovery(ts)
	if errs := syntaxErrors(ts, tree.Tree(), parseErrs); len(errs) > 0 {
		return nil, fmt.Errorf("invalid rule %q: %s", rule, strings.Join(errs, "; "))
	}
	if next := trailingInput(tree.Tree()); tree.Tree().Tokens[next].Type != JsonQueryEOF {
		return nil, fmt.Errorf("invalid rule %q: unexpected input %q", rule, ts.Text(next))
	}

	return &Evaluator{
		rule: rule,
		tree: tree,
	}, nil
}

func (e *Evaluator) Reset() error {
	e.lastDebugErr = nil
	return nil
}

func (e *Evaluator) LastDebugErr() error {
	return e.lastDebugErr
}

func (e *Evaluator) Process(items map[string]interface{}) (ret bool, retErr error) {
	e.lastDebugErr = nil
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
			ret = false
		}
	}()

	visitor := NewJsonQueryVisitorImpl(items)
	result := visitor.Visit(e.tree)
	e.lastDebugErr = visitor.debugErr
	if e.testHookPanic != nil {
		defer e.testHookPanic()
	}
	if result == nil || visitor.err != nil {
		return false, visitor.err
	}

	return result.(bool), visitor.err
}

func Evaluate(rule string, items map[string]interface{}) bool {
	ev, err := NewEvaluator(rule)
	if err != nil {
		return false
	}
	result, _ := ev.Process(items)
	return result
}
