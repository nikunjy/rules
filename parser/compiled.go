package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/blang/semver/v4"
)

type nodeKind int

const (
	nodeInvalid nodeKind = iota
	nodeParen
	nodeLogical
	nodePresent
	nodeCompare
)

type valueKind int

const (
	valueConst valueKind = iota
	valueInts
	valueDoubles
	valueStrings
)

// node is one immutable node of a rule compiled from the ANTLR parse tree.
// Everything that does not depend on the input object (operators, constants,
// attribute paths, IN lists) is resolved once by compileQuery so that eval
// never touches the parse tree. Paren and logical nodes only use the fields
// below; present and compare nodes are allocated as a compareNode holding
// the rest, so the common inner nodes stay small.
type node struct {
	kind nodeKind

	// paren / logical
	not  bool // paren: negate the result
	isOr bool // logical: or (true) / and (false)

	left, right *node

	cmp *compareNode // present / compare: the compareNode this node is part of
}

// compareNode is a present or compare node.
type compareNode struct {
	node

	path []string

	// compare
	attrPath string
	opType   int
	opText   string
	apply    func(Operand, Operand) (bool, error)
	version  *versionCompare // set instead of using apply for a pre-parsed rule version
	valKind  valueKind
	rightOp  Operand // valueConst: operand set for the rule value
	boolErr  error   // valueConst: failed boolean parse, reported on eval
	valErr   error   // valueConst: error set on eval (invalid integer)
	listErr  error   // lists: parse error of the first bad element; the list then holds the elements before it
	listOp   Operand // lists: the parsed list ([]int, []float64 or []string)
}

func newCompareNode(kind nodeKind, path []string) *compareNode {
	c := &compareNode{path: path}
	c.kind = kind
	c.cmp = c
	return c
}

// evalState mirrors the mutable state the visitor keeps while walking a
// rule, including its quirks (stale leftOp/rightOp, the attribute stack).
type evalState struct {
	item map[string]interface{}

	stack attrStack

	leftOp  Operand
	rightOp Operand

	err      error
	debugErr error
}

// attrStack is the visitor's attribute stack. The first attrStackInline
// entries live in the struct itself so that, with evalState on the
// goroutine stack, nested paths do not allocate; deeper entries spill to
// more. Nothing takes the address of inline, which would move the whole
// evalState to the heap.
type attrStack struct {
	n      int
	inline [attrStackInline]interface{}
	more   []interface{}
}

const attrStackInline = 4

func (s *attrStack) push(v interface{}) {
	if s.n < attrStackInline {
		s.inline[s.n] = v
	} else {
		s.more = append(s.more, v)
	}
	s.n++
}

// top returns the last pushed entry; s must not be empty.
func (s *attrStack) top() interface{} {
	if s.n <= attrStackInline {
		return s.inline[s.n-1]
	}
	return s.more[s.n-1-attrStackInline]
}

// pop drops the last pushed entry; s must not be empty.
func (s *attrStack) pop() {
	s.n--
	if s.n >= attrStackInline {
		s.more = s.more[:s.n-attrStackInline]
	}
}

func (s *attrStack) clear() {
	s.n = 0
	s.more = s.more[:0]
}

func (st *evalState) resolve(path []string) {
	last := len(path) - 1
	for _, name := range path[:last] {
		var item interface{}
		if st.stack.n > 0 {
			item = st.stack.top()
		} else {
			item = st.item
		}
		if item == nil {
			return
		}
		m := item.(map[string]interface{})
		st.stack.push(m[name])
	}
	var item interface{}
	if st.stack.n > 0 {
		item = st.stack.top()
		st.stack.pop()
	} else {
		item = st.item
	}
	if item == nil {
		return
	}
	m := item.(map[string]interface{})
	st.leftOp = m[path[last]]
	st.stack.clear()
}

func (n *node) eval(st *evalState) bool {
	switch n.kind {
	case nodeParen:
		result := n.left.eval(st)
		if n.not {
			return !result
		}
		return result
	case nodeLogical:
		left := n.left.eval(st)
		if n.isOr {
			if left {
				return left
			}
			return n.right.eval(st)
		}
		if !left {
			return left
		}
		return n.right.eval(st)
	case nodePresent:
		st.resolve(n.cmp.path)
		return st.leftOp != nil
	case nodeCompare:
		return n.cmp.evalCompare(st)
	default:
		st.err = errors.New("Invalid rule")
		return false
	}
}

func (n *compareNode) setRight(st *evalState) {
	if n.valKind != valueConst && n.listErr == nil && st.rightOp == nil {
		st.rightOp = n.listOp
		return
	}
	switch n.valKind {
	case valueInts:
		var list []int
		if st.rightOp == nil {
			list = []int{}
		} else {
			list = st.rightOp.([]int)
		}
		st.rightOp = append(list, n.listOp.([]int)...)
		if n.listErr != nil {
			st.err = cloneErr(n.listErr)
		}
	case valueDoubles:
		var list []float64
		if st.rightOp == nil {
			list = []float64{}
		} else {
			list = st.rightOp.([]float64)
		}
		st.rightOp = append(list, n.listOp.([]float64)...)
		if n.listErr != nil {
			st.err = cloneErr(n.listErr)
		}
	case valueStrings:
		var list []string
		if st.rightOp == nil {
			list = []string{}
		} else {
			list = st.rightOp.([]string)
		}
		st.rightOp = append(list, n.listOp.([]string)...)
	default:
		st.rightOp = n.rightOp
		if n.boolErr != nil {
			st.err = newNestedError(n.boolErr, "Error converting boolean")
		}
		if n.valErr != nil {
			st.err = cloneErr(n.valErr)
		}
	}
}

func (n *compareNode) evalCompare(st *evalState) bool {
	st.resolve(n.path)
	n.setRight(st)
	if st.err != nil {
		return false
	}
	if n.apply == nil {
		st.err = fmt.Errorf("Unknown operation %s", n.opText)
		return false
	}
	var ret bool
	var err error
	if n.version != nil {
		ret, err = n.version.compare(st.leftOp)
	} else {
		ret, err = n.apply(st.leftOp, st.rightOp)
	}
	if err != nil {
		if eio, ok := err.(*ErrInvalidOperand); ok && n.valKind != valueConst {
			// the operand may be the shared compiled list
			c := *eio
			c.Val = copyOperand(c.Val)
			c.typeObj = copyOperand(c.typeObj)
			err = &c
		}
		n.setDebugErr(st, err)
		ret = false
	}
	st.rightOp = nil
	return ret
}

func (n *compareNode) setDebugErr(st *evalState, err error) {
	switch err {
	case ErrInvalidOperation:
		// in case of invalid operation lets rather
		// be conservative and return false because the rule doesn't even make
		// sense. It can be argued that it would be false positive if we were
		// to return true
		st.err = err
		st.debugErr = newNestedError(err, "Not a valid operation for datatypes").Set(ErrVals{
			"operation":           n.opType,
			"object_path_operand": st.leftOp,
			"rule_operand":        copyOperand(st.rightOp),
		})
	case ErrEvalOperandMissing:
		st.debugErr = newNestedError(err, "Eval operand missing in input object").Set(ErrVals{
			"attr_path": n.attrPath,
		})
	default:
		msg := "unknown error"
		if _, ok := err.(*ErrInvalidOperand); ok {
			msg = "operands are not the right value type"
		}
		st.debugErr = newNestedError(err, msg).Set(ErrVals{
			"attr_path":           n.attrPath,
			"object_path_operand": st.leftOp,
			"rule_operand":        copyOperand(st.rightOp),
		})
	}
}

// copyOperand copies compiled IN lists before they are exposed through a
// debug error so callers cannot mutate the shared rule constants.
func copyOperand(op Operand) Operand {
	switch val := op.(type) {
	case []int:
		return append(make([]int, 0, len(val)), val...)
	case []float64:
		return append(make([]float64, 0, len(val)), val...)
	case []string:
		return append(make([]string, 0, len(val)), val...)
	}
	return op
}

// cloneErr returns a fresh copy of a pre-computed parse error so each call
// surfaces its own error value.
func cloneErr(err error) error {
	if ne, ok := err.(*strconv.NumError); ok {
		c := *ne
		return &c
	}
	return err
}

func compileQuery(tree antlr.ParseTree) *node {
	switch val := tree.(type) {
	case *ParenExpContext:
		return &node{kind: nodeParen, left: compileQuery(val.Query()), not: val.NOT() != nil}
	case *LogicalExpContext:
		return &node{
			kind:  nodeLogical,
			left:  compileQuery(val.Query(0)),
			right: compileQuery(val.Query(1)),
			isOr:  val.LOGICAL_OPERATOR().GetText() == "or",
		}
	case *PresentExpContext:
		return &newCompareNode(nodePresent, compileAttrPath(val.AttrPath())).node
	case *CompareExpContext:
		n := newCompareNode(nodeCompare, compileAttrPath(val.AttrPath()))
		n.attrPath = val.AttrPath().GetText()
		n.opType = val.op.GetTokenType()
		n.opText = val.op.GetText()
		op := compileValue(n, val.Value())
		n.apply = applyFor(op, n.opType)
		if _, ok := op.(*VersionOperation); ok {
			n.version = compileVersionCompare(n.rightOp.(string), n.opType)
		}
		return &n.node
	default:
		return &node{kind: nodeInvalid}
	}
}

func compileAttrPath(ctx IAttrPathContext) []string {
	var path []string
	for {
		path = append(path, ctx.ATTRNAME().GetText())
		sub := ctx.SubAttr()
		if sub == nil || sub.IsEmpty() {
			return path
		}
		ctx = sub.AttrPath()
	}
}

func applyFor(op Operation, opType int) func(Operand, Operand) (bool, error) {
	switch opType {
	case JsonQueryParserEQ:
		return op.EQ
	case JsonQueryParserNE:
		return op.NE
	case JsonQueryParserGT:
		return op.GT
	case JsonQueryParserLT:
		return op.LT
	case JsonQueryParserLE:
		return op.LE
	case JsonQueryParserGE:
		return op.GE
	case JsonQueryParserCO:
		return op.CO
	case JsonQueryParserSW:
		return op.SW
	case JsonQueryParserEW:
		return op.EW
	case JsonQueryParserIN:
		return op.IN
	}
	return nil
}

// versionCompare is a VersionOperation comparison whose rule version was
// parsed at compile time.
type versionCompare struct {
	right semver.Version
	// want reports whether left.Compare(right) satisfies the operator.
	want func(cmp int) bool
}

// compileVersionCompare returns a comparison against the pre-parsed version
// right for the six comparison operators, or nil to keep VersionOperation
// (other operators, or a rule version semver rejects: its error is then
// produced on every Process exactly as before).
func compileVersionCompare(right string, opType int) *versionCompare {
	var want func(int) bool
	switch opType {
	case JsonQueryParserEQ:
		want = func(c int) bool { return c == 0 }
	case JsonQueryParserNE:
		want = func(c int) bool { return c != 0 }
	case JsonQueryParserGT:
		want = func(c int) bool { return c > 0 }
	case JsonQueryParserLT:
		want = func(c int) bool { return c < 0 }
	case JsonQueryParserGE:
		want = func(c int) bool { return c >= 0 }
	case JsonQueryParserLE:
		want = func(c int) bool { return c <= 0 }
	default:
		return nil
	}
	ver, err := semver.Make(right)
	if err != nil {
		return nil
	}
	return &versionCompare{right: ver, want: want}
}

// compare mirrors VersionOperation.get followed by the comparison; the rule
// operand is always the compiled version string, so only left is checked.
func (c *versionCompare) compare(left Operand) (bool, error) {
	leftVal, ok := left.(string)
	if !ok {
		return false, newErrInvalidOperand(left, leftVal)
	}
	leftVer, ok := parsePlainVersion(leftVal)
	if !ok {
		var err error
		if leftVer, err = semver.Make(leftVal); err != nil {
			return false, err
		}
	}
	return c.want(leftVer.Compare(c.right)), nil
}

// parsePlainVersion parses "MAJOR.MINOR.PATCH" with optional "+build"
// metadata without allocating. It only accepts strings that semver.Make
// accepts, and returns the same Major, Minor and Patch with no prerelease;
// build metadata is validated but dropped since Compare ignores it. Anything
// else (prerelease, errors) reports false so the caller uses semver.Make.
func parsePlainVersion(s string) (semver.Version, bool) {
	var v semver.Version
	core := s
	if i := strings.IndexByte(s, '+'); i >= 0 {
		if !validBuild(s[i+1:]) {
			return v, false
		}
		core = s[:i]
	}
	var nums [3]uint64
	for i := range nums {
		end := len(core)
		if i < 2 {
			end = strings.IndexByte(core, '.')
			if end < 0 {
				return v, false
			}
		}
		n, ok := parseVersionNumber(core[:end])
		if !ok {
			return v, false
		}
		nums[i] = n
		if i < 2 {
			core = core[end+1:]
		}
	}
	v.Major, v.Minor, v.Patch = nums[0], nums[1], nums[2]
	return v, true
}

// parseVersionNumber accepts a non-empty run of digits without leading
// zeroes that fits in a uint64, like semver.Parse does for a version number.
func parseVersionNumber(s string) (uint64, bool) {
	if s == "" || (len(s) > 1 && s[0] == '0') {
		return 0, false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
	}
	n, err := strconv.ParseUint(s, 10, 64)
	return n, err == nil
}

// validBuild reports whether s is dot separated, non-empty identifiers of
// [0-9A-Za-z-], the build metadata semver.Parse accepts.
func validBuild(s string) bool {
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '.' {
			if i == start {
				return false
			}
			start = i + 1
			continue
		}
		c := s[i]
		if !(c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '-') {
			return false
		}
	}
	return true
}

// compileValue pre-parses the right-hand side of a comparison into n and
// returns the Operation the visitor would have selected for it.
func compileValue(n *compareNode, v IValueContext) Operation {
	switch ctx := v.(type) {
	case *BooleanContext:
		val, err := strconv.ParseBool(ctx.GetText())
		if err != nil {
			n.boolErr = err
		} else {
			n.rightOp = val
		}
		return &BoolOperation{}
	case *NullContext:
		return &NullOperation{}
	case *StringContext:
		n.rightOp = getString(ctx.GetText())
		return &StringOperation{}
	case *DoubleContext:
		// an unparsable double leaves the operand nil without an error
		if val, err := strconv.ParseFloat(ctx.GetText(), 10); err == nil {
			n.rightOp = val
		}
		return &FloatOperation{}
	case *VersionContext:
		n.rightOp = ctx.VERSION().GetText()
		return &VersionOperation{}
	case *LongContext:
		val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
		if err != nil {
			n.valErr = err
		} else {
			n.rightOp = int(val)
		}
		return &IntOperation{}
	case *ListOfIntsContext:
		n.valKind = valueInts
		ints := []int{}
		for sub := ctx.ListInts().SubListOfInts(); ; {
			val, err := strconv.ParseInt(sub.INT().GetText(), 10, 64)
			if err != nil {
				n.listErr = err
				break
			}
			ints = append(ints, int(val))
			next := sub.SubListOfInts()
			if next == nil || next.IsEmpty() {
				break
			}
			sub = next
		}
		n.listOp = ints[:len(ints):len(ints)]
		return &IntOperation{}
	case *ListOfDoublesContext:
		n.valKind = valueDoubles
		doubles := []float64{}
		for sub := ctx.ListDoubles().SubListOfDoubles(); ; {
			val, err := strconv.ParseFloat(sub.DOUBLE().GetText(), 10)
			if err != nil {
				n.listErr = err
				break
			}
			doubles = append(doubles, val)
			next := sub.SubListOfDoubles()
			if next == nil || next.IsEmpty() {
				break
			}
			sub = next
		}
		n.listOp = doubles[:len(doubles):len(doubles)]
		return &FloatOperation{}
	case *ListOfStringsContext:
		n.valKind = valueStrings
		strs := []string{}
		for sub := ctx.ListStrings().SubListOfStrings(); ; {
			strs = append(strs, getString(sub.STRING().GetText()))
			next := sub.SubListOfStrings()
			if next == nil || next.IsEmpty() {
				break
			}
			sub = next
		}
		n.listOp = strs[:len(strs):len(strs)]
		return &StringOperation{}
	}
	return &NullOperation{}
}
