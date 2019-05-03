package parser

type BoolOperation struct {
	NullOperation
}

func (o *BoolOperation) get(left Operand, right Operand) (bool, bool, bool) {
	if left == nil {
		return false, false, false
	}
	leftVal, ok := left.(bool)
	rightVal, ok1 := right.(bool)
	return leftVal, rightVal, (ok && ok1)
}

func (o *BoolOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l == r, nil
}

func (o *BoolOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
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
