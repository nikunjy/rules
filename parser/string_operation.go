package parser

import (
	"strings"
)

type StringOperation struct {
	NullOperation
}

func (o *StringOperation) get(left Operand, right Operand) (string, string, bool) {
	if left == nil {
		return "", "", false
	}
	leftVal, ok := left.(string)
	rightVal, ok1 := right.(string)
	return strings.ToLower(leftVal), strings.ToLower(rightVal), (ok1 && ok)

}

func (o *StringOperation) EQ(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l == r, nil
}

func (o *StringOperation) NE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l != r, nil
}

func (o *StringOperation) GT(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l > r, nil
}

func (o *StringOperation) LT(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l < r, nil
}

func (o *StringOperation) GE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l >= r, nil
}

func (o *StringOperation) LE(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok, nil
	}
	return l <= r, nil
}

func (o *StringOperation) CO(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return false, nil
	}
	return strings.Contains(l, r), nil
}

func (o *StringOperation) SW(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return false, nil
	}
	return strings.HasPrefix(l, r), nil
}

func (o *StringOperation) EW(left Operand, right Operand) (bool, error) {
	l, r, ok := o.get(left, right)
	if !ok {
		return false, nil
	}
	return strings.HasSuffix(l, r), nil
}

func (o *StringOperation) IN(left Operand, right Operand) (bool, error) {
	leftVal, ok := left.(string)
	if !ok {
		return ok, nil
	}

	rightVal, ok := right.([]string)
	if !ok {
		return ok, nil
	}
	for _, val := range rightVal {
		if leftVal == val {
			return true, nil
		}
	}
	return false, nil
}
