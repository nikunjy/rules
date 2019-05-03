package parser

type BoolOperation struct {
	NullOperation
}

func (o *BoolOperation) get(left Operand, right Operand) (bool, bool, error) {
	if left == nil {
		return false, false, ErrEvalOperandMissing
	}
	leftVal, ok := left.(bool)
	if !ok {
		return false, false, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.(bool)
	if !ok {
		return false, false, newErrInvalidOperand(right, rightVal)
	}
	return leftVal, rightVal, nil
}

func (o *BoolOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *BoolOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *BoolOperation) GT(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) LT(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) GE(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) LE(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) CO(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) SW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *BoolOperation) EW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}
