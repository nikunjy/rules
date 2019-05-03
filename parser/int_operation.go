package parser

type IntOperation struct {
}

func (o *IntOperation) get(left Operand, right Operand) (int, int, error) {
	if left == nil {
		return 0, 0, ErrEvalOperandMissing
	}
	leftVal, ok := left.(int)
	if !ok {
		return 0, 0, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.(int)
	if !ok {
		return 0, 0, newErrInvalidOperand(left, leftVal)
	}
	return leftVal, rightVal, nil

}

func (o *IntOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *IntOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *IntOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l > r, nil
}

func (o *IntOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l < r, nil
}

func (o *IntOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l >= r, nil
}

func (o *IntOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
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
		return false, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.([]int)
	if !ok {
		return false, newErrInvalidOperand(right, rightVal)
	}
	for _, num := range rightVal {
		if num == leftVal {
			return true, nil
		}
	}
	return false, nil
}
