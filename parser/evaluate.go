package parser

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/antlr4-go/antlr/v4"
)

type Evaluator struct {
	// lastDebugErr is atomic so concurrent Process calls do not race on it;
	// with concurrent callers it holds the error of whichever call finished last.
	lastDebugErr atomic.Pointer[error]

	rule string

	// tree is kept until the first Process compiles it into root, after
	// which it is dropped. If compiling fails, tree stays and Process falls
	// back to walking it with the visitor.
	tree antlr.ParseTree
	once sync.Once
	root *node

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
	tree, err := parseRule(rule)
	if err != nil {
		return nil, err
	}

	return &Evaluator{
		rule: rule,
		tree: tree,
	}, nil
}

func parseRule(rule string) (antlr.ParseTree, error) {
	input := antlr.NewInputStream(normalizeRuleInput(rule))
	lex := NewJsonQueryLexer(input)
	errListener := &collectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lex.RemoveErrorListeners()
	lex.AddErrorListener(errListener)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := NewJsonQueryParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	tree := p.Query()
	if len(errListener.errs) > 0 {
		return nil, fmt.Errorf("invalid rule %q: %s", rule, strings.Join(errListener.errs, "; "))
	}
	consumeTrailingWhitespace(tokens)
	if tokens.LA(1) != antlr.TokenEOF {
		return nil, fmt.Errorf("invalid rule %q: unexpected input %q", rule, tokens.LT(1).GetText())
	}
	return tree, nil
}

func (e *Evaluator) Reset() error {
	e.lastDebugErr.Store(nil)
	return nil
}

func (e *Evaluator) LastDebugErr() error {
	if p := e.lastDebugErr.Load(); p != nil {
		return *p
	}
	return nil
}

func (e *Evaluator) setDebugErr(err error) {
	if err == nil {
		// load first so concurrent clean calls on a shared evaluator only
		// read the shared cache line instead of all writing to it
		if e.lastDebugErr.Load() != nil {
			e.lastDebugErr.Store(nil)
		}
		return
	}
	box := err
	e.lastDebugErr.Store(&box)
}

// compile turns the parse tree into the node tree used by Process. It runs
// once; a panic leaves the tree in place so Process uses the visitor instead.
func (e *Evaluator) compile() {
	defer func() {
		// Deliberately silent: compiling is only an optimization. A panic
		// here (e.g. a tree shape compileQuery does not expect) must not
		// change what Process returns, so drop any partial result and keep
		// the tree; Process then walks it with the visitor, which reports
		// whatever error or result the uncompiled evaluator always did.
		if info := recover(); info != nil {
			e.root = nil
		}
	}()
	if e.tree == nil {
		return
	}
	e.root = compileQuery(e.tree)
	e.tree = nil
}

func (e *Evaluator) Process(items map[string]interface{}) (bool, error) {
	return e.process(items, true)
}

// process evaluates items against the rule. With compiled set the rule is
// compiled on first use; otherwise it is walked with the visitor, which avoids
// the compile cost for one-shot evaluations.
func (e *Evaluator) process(items map[string]interface{}, compiled bool) (ret bool, retErr error) {
	e.setDebugErr(nil)
	// antlr lib has panics for exceptions so we have to put a recover here
	// in the unlikely case there is an exception
	defer func() {
		info := recover()
		if info != nil {
			retErr = fmt.Errorf("%q", info)
			ret = false
		}
	}()

	if compiled {
		e.once.Do(e.compile)
	}
	var result bool
	var err error
	if e.root != nil {
		st := evalState{item: items}
		result = e.root.eval(&st)
		err = st.err
		e.setDebugErr(st.debugErr)
	} else if e.tree != nil {
		result, err = e.visit(items)
	} else {
		err = errors.New("Invalid rule")
	}
	if e.testHookPanic != nil {
		defer e.testHookPanic()
	}
	if err != nil {
		return false, err
	}

	return result, nil
}

func (e *Evaluator) visit(items map[string]interface{}) (bool, error) {
	visitor := NewJsonQueryVisitorImpl(items)
	result := visitor.Visit(e.tree)
	e.setDebugErr(visitor.debugErr)
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
	result, _ := ev.process(items, false)
	return result
}
