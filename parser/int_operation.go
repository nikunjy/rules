package parser

type IntOperation struct {
}

func (o *IntOperation) get(left Operand, right Operand) (int, int, bool) {
	if left == nil {
		return 0, 0, false
	}
	leftVal, ok := left.(int)
	rightVal, ok1 := right.(int)
	return leftVal, rightVal, (ok1 && ok)

}

func (o *IntOperation) EQ(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l == r
}

func (o *IntOperation) NE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l != r
}

func (o *IntOperation) GT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l > r
}

func (o *IntOperation) LT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l < r
}

func (o *IntOperation) GE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l >= r
}

func (o *IntOperation) LE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l <= r
}

func (o *IntOperation) CO(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *IntOperation) SW(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *IntOperation) EW(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *IntOperation) IN(left Operand, right Operand) bool {
	leftVal, ok := left.(int)
	if !ok {
		return ok
	}
	rightVal, ok := right.([]int)
	if !ok {
		return ok
	}
	for _, num := range rightVal {
		if num == leftVal {
			return true
		}
	}
	return false
}
