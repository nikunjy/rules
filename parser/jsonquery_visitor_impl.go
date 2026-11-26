package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type objStack struct {
	items []interface{}
}

func (o *objStack) empty() bool {
	return len(o.items) == 0
}

func (o *objStack) peek() interface{} {
	return o.items[len(o.items)-1]
}

func (o *objStack) pop() interface{} {
	val := o.peek()
	o.items = o.items[:len(o.items)-1]
	return val
}

func (o *objStack) push(item interface{}) {
	o.items = append(o.items, item)
}

func (o *objStack) clear() {
	o.items = nil
}

// JsonQueryVisitorImpl evaluates a parsed rule against one input object.
type JsonQueryVisitorImpl struct {
	item map[string]interface{}

	stack *objStack

	currentOperation Operation
	leftOp           Operand
	rightOp          Operand

	err      error
	debugErr error
}

func NewJsonQueryVisitorImpl(item map[string]interface{}) *JsonQueryVisitorImpl {
	return &JsonQueryVisitorImpl{
		stack: &objStack{},
		item:  item,
	}
}

func (j *JsonQueryVisitorImpl) setDebugErr(err error) {
	j.debugErr = err
}

func (j *JsonQueryVisitorImpl) setErr(err error) {
	j.err = err
}

func (j *JsonQueryVisitorImpl) hasErr() bool {
	return j.err != nil
}

// Visit evaluates a query node to a bool.
func (j *JsonQueryVisitorImpl) Visit(n QueryNode) interface{} {
	if j.hasErr() {
		return false
	}
	return j.visitQuery(n)
}

func (j *JsonQueryVisitorImpl) visitQuery(n QueryNode) bool {
	if c, ok := n.AsLogicalExp(); ok {
		return j.VisitLogicalExp(c)
	}
	if c, ok := n.AsCompareExp(); ok {
		return j.VisitCompareExp(c)
	}
	if c, ok := n.AsParenExp(); ok {
		return j.VisitParenExp(c)
	}
	if c, ok := n.AsPresentExp(); ok {
		return j.VisitPresentExp(c)
	}
	j.setErr(errors.New("Invalid rule"))
	return false
}

func (j *JsonQueryVisitorImpl) VisitParenExp(ctx ParenExpNode) bool {
	result := j.visitQuery(ctx.Query())
	if ctx.NOT().Ok() {
		return !result
	}
	return result
}

func (j *JsonQueryVisitorImpl) VisitLogicalExp(ctx LogicalExpNode) bool {
	left := j.visitQuery(ctx.Query(0))
	op := ctx.LOGICAL_OPERATOR().GetText()
	if op == "or" {
		if left {
			return left
		}
		return j.visitQuery(ctx.Query(1))
	}
	// means it is and
	if !left {
		return left
	}
	return j.visitQuery(ctx.Query(1))
}

func (j *JsonQueryVisitorImpl) VisitPresentExp(ctx PresentExpNode) bool {
	j.VisitAttrPath(ctx.AttrPath())
	return j.leftOp != nil
}

func (j *JsonQueryVisitorImpl) VisitCompareExp(ctx CompareExpNode) bool {
	j.VisitAttrPath(ctx.AttrPath())
	j.visitValue(ctx.Value())
	if j.hasErr() {
		return false
	}
	var apply func(Operand, Operand) (bool, error)
	currentOp := j.currentOperation
	op := ctx.GetOp()
	switch op.Token().Type {
	case JsonQueryEQ:
		apply = currentOp.EQ
	case JsonQueryNE:
		apply = currentOp.NE
	case JsonQueryGT:
		apply = currentOp.GT
	case JsonQueryLT:
		apply = currentOp.LT
	case JsonQueryLE:
		apply = currentOp.LE
	case JsonQueryGE:
		apply = currentOp.GE
	case JsonQueryCO:
		apply = currentOp.CO
	case JsonQuerySW:
		apply = currentOp.SW
	case JsonQueryEW:
		apply = currentOp.EW
	case JsonQueryIN:
		apply = currentOp.IN
	default:
		j.setErr(fmt.Errorf("Unknown operation %s", op.GetText()))
		return false
	}
	defer func() { j.rightOp = nil }()
	ret, err := apply(j.leftOp, j.rightOp)
	if err != nil {
		switch err {
		case ErrInvalidOperation:
			// in case of invalid operation lets rather
			// be conservative and return false because the rule doesn't even make
			// sense. It can be argued that it would be false positive if we were
			// to return true
			j.setErr(err)
			j.setDebugErr(
				newNestedError(err, "Not a valid operation for datatypes").Set(ErrVals{
					"operation":           int(op.Token().Type),
					"object_path_operand": j.leftOp,
					"rule_operand":        j.rightOp,
				}),
			)
		case ErrEvalOperandMissing:
			j.setDebugErr(
				newNestedError(err, "Eval operand missing in input object").Set(ErrVals{
					"attr_path": ctx.AttrPath().GetText(),
				}),
			)
		default:
			switch err.(type) {
			case *ErrInvalidOperand:
				j.setDebugErr(
					newNestedError(err, "operands are not the right value type").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)
			default:
				j.setDebugErr(
					newNestedError(err, "unknown error").Set(ErrVals{
						"attr_path":           ctx.AttrPath().GetText(),
						"object_path_operand": j.leftOp,
						"rule_operand":        j.rightOp,
					}),
				)

			}
		}

		return false
	}
	return ret
}

func (j *JsonQueryVisitorImpl) VisitAttrPath(ctx AttrPathNode) {
	var item interface{}
	if !ctx.SubAttr().Ok() {
		if !j.stack.empty() {
			item = j.stack.pop()
		} else {
			item = j.item
		}
		if item == nil {
			return
		}
		m := item.(map[string]interface{})
		j.leftOp = m[ctx.ATTRNAME().GetText()]
		j.stack.clear()
		return
	}
	if !j.stack.empty() {
		item = j.stack.peek()
	} else {
		item = j.item
	}
	if item == nil {
		return
	}
	m := item.(map[string]interface{})
	j.stack.push(m[ctx.ATTRNAME().GetText()])
	j.VisitAttrPath(ctx.SubAttr().AttrPath())
}

func (j *JsonQueryVisitorImpl) visitValue(n ValueNode) {
	if c, ok := n.AsBoolean(); ok {
		j.VisitBoolean(c)
	} else if c, ok := n.AsNull(); ok {
		j.VisitNull(c)
	} else if c, ok := n.AsVersion(); ok {
		j.VisitVersion(c)
	} else if c, ok := n.AsString(); ok {
		j.VisitString(c)
	} else if c, ok := n.AsDouble(); ok {
		j.VisitDouble(c)
	} else if c, ok := n.AsLong(); ok {
		j.VisitLong(c)
	} else if c, ok := n.AsListOfInts(); ok {
		j.currentOperation = &IntOperation{}
		j.VisitSubListOfInts(c.ListInts().SubListOfInts())
	} else if c, ok := n.AsListOfDoubles(); ok {
		j.currentOperation = &FloatOperation{}
		j.VisitSubListOfDoubles(c.ListDoubles().SubListOfDoubles())
	} else if c, ok := n.AsListOfStrings(); ok {
		j.currentOperation = &StringOperation{}
		j.VisitSubListOfStrings(c.ListStrings().SubListOfStrings())
	}
}

func (j *JsonQueryVisitorImpl) VisitBoolean(ctx BooleanNode) {
	j.currentOperation = &BoolOperation{}

	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		j.rightOp = nil
		j.setErr(newNestedError(err, "Error converting boolean"))
		return
	}
	j.rightOp = val
}

func (j *JsonQueryVisitorImpl) VisitNull(ctx NullNode) {
	j.currentOperation = &NullOperation{}
	j.rightOp = nil
}

func getString(s string) string {
	if len(s) > 2 {
		return s[1 : len(s)-1]
	}
	return ""
}

func (j *JsonQueryVisitorImpl) VisitString(ctx StringNode) {
	j.currentOperation = &StringOperation{}
	j.rightOp = getString(ctx.GetText())
}

func (j *JsonQueryVisitorImpl) VisitDouble(ctx DoubleNode) {
	j.currentOperation = &FloatOperation{}
	val, err := strconv.ParseFloat(ctx.GetText(), 10)
	if err != nil {
		// TODO set err somewhere
		j.rightOp = nil
		return
	}
	j.rightOp = val
}

func (j *JsonQueryVisitorImpl) VisitVersion(ctx VersionNode) {
	j.currentOperation = &VersionOperation{}
	j.rightOp = ctx.VERSION().GetText()
}

func (j *JsonQueryVisitorImpl) VisitLong(ctx LongNode) {
	j.currentOperation = &IntOperation{}
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		j.rightOp = nil
		j.setErr(err)
		return
	}
	j.rightOp = int(val)
}

func (j *JsonQueryVisitorImpl) VisitSubListOfInts(ctx SubListOfIntsNode) {
	if j.rightOp == nil {
		j.rightOp = make([]int, 0)
	}
	list := j.rightOp.([]int)
	val, err := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	if err != nil {
		j.setErr(err)
		return
	}
	j.rightOp = append(list, int(val))
	if next := ctx.SubListOfInts(); next.Ok() {
		j.VisitSubListOfInts(next)
	}
}

func (j *JsonQueryVisitorImpl) VisitSubListOfDoubles(ctx SubListOfDoublesNode) {
	if j.rightOp == nil {
		j.rightOp = make([]float64, 0)
	}
	list := j.rightOp.([]float64)
	val, err := strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)
	if err != nil {
		j.setErr(err)
		return
	}
	j.rightOp = append(list, val)
	if next := ctx.SubListOfDoubles(); next.Ok() {
		j.VisitSubListOfDoubles(next)
	}
}

func (j *JsonQueryVisitorImpl) VisitSubListOfStrings(ctx SubListOfStringsNode) {
	if j.rightOp == nil {
		j.rightOp = make([]string, 0)
	}
	val := getString(ctx.STRING().GetText())
	list := j.rightOp.([]string)
	j.rightOp = append(list, val)
	if next := ctx.SubListOfStrings(); next.Ok() {
		j.VisitSubListOfStrings(next)
	}
}
