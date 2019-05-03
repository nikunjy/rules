package parser

import (
	"strings"
)

type StringOperation struct {
	NullOperation
}

func (o *StringOperation) get(left Operand, right Operand) (string, string, error) {
	if left == nil {
		return "", "", ErrEvalOperandMissing
	}
	leftVal, ok := left.(string)
	if !ok {
		return "", "", newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.(string)
	if !ok {
		return "", "", newErrInvalidOperand(right, rightVal)
	}
	return strings.ToLower(leftVal), strings.ToLower(rightVal), nil

}

func (o *StringOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l == r, nil
}

func (o *StringOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l != r, nil
}

func (o *StringOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l > r, nil
}

func (o *StringOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l < r, nil
}

func (o *StringOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l >= r, nil
}

func (o *StringOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return l <= r, nil
}

func (o *StringOperation) CO(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return strings.Contains(l, r), nil
}

func (o *StringOperation) SW(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(l, r), nil
}

func (o *StringOperation) EW(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(l, r), nil
}

func (o *StringOperation) IN(left Operand, right Operand) (bool, error) {
	leftVal, ok := left.(string)
	if !ok {
		return false, newErrInvalidOperand(left, leftVal)
	}

	rightVal, ok := right.([]string)
	if !ok {
		return ok, newErrInvalidOperand(right, rightVal)
	}
	for _, val := range rightVal {
		if leftVal == val {
			return true, nil
		}
	}
	return false, nil
}
