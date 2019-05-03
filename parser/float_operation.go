package parser

type FloatOperation struct {
	NullOperation
}

func toNum(op Operand) (float64, error) {
	switch val := op.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	case int64:
		return float64(val), nil
	}
	var exp float64
	return 0, newErrInvalidOperand(op, exp)
}

func (o *FloatOperation) get(left Operand, right Operand) (float64, float64, error) {
	if left == nil {
		return 0, 0, ErrEvalOperandMissing
	}
	leftVal, err := toNum(left)
	if err != nil {
		return 0, 0, err
	}
	rightVal, err := toNum(right)
	if err != nil {
		return 0, 0, err
	}
	return leftVal, rightVal, nil

}

func (o *FloatOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *FloatOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *FloatOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l > r, nil
}

func (o *FloatOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l < r, nil
}

func (o *FloatOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l >= r, nil
}

func (o *FloatOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, nil
	}
	return l <= r, nil
}

func (o *FloatOperation) IN(left Operand, right Operand) (bool, error) {
	leftVal, err := toNum(left)
	if err != nil {
		return false, err
	}
	rightVal, ok := right.([]float64)
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
