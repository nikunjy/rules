package parser

type BoolOperation struct {
}

func (o *BoolOperation) get(left Operand, right Operand) (bool, bool, bool) {
	if left == nil {
		return false, false, false
	}
	leftVal, ok := left.(bool)
	rightVal, ok1 := right.(bool)
	return leftVal, rightVal, (ok && ok1)
}

func (o *BoolOperation) EQ(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l == r
}

func (o *BoolOperation) NE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l != r
}

func (o *BoolOperation) GT(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) LT(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) GE(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) LE(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) CO(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) SW(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *BoolOperation) EW(left Operand, right Operand) bool {
	panic("not supported")
}
