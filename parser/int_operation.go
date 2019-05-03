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

func (o *IntOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l == r, nil
}

func (o *IntOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l != r, nil
}

func (o *IntOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l > r, nil
}

func (o *IntOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l < r, nil
}

func (o *IntOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l >= r, nil
}

func (o *IntOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l <= r, nil
}

func (o *IntOperation) CO(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *IntOperation) SW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *IntOperation) EW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *IntOperation) IN(left Operand, right Operand) (bool, error) {
	leftVal, ok := left.(int)
	if !ok {
		return ok, nil
	}
	rightVal, ok := right.([]int)
	if !ok {
		return ok, nil
	}
	for _, num := range rightVal {
		if num == leftVal {
			return true, nil
		}
	}
	return false, nil
}
