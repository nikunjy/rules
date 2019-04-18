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

func (o *StringOperation) EQ(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l == r
}

func (o *StringOperation) NE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l != r
}

func (o *StringOperation) GT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l > r
}

func (o *StringOperation) LT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l < r
}

func (o *StringOperation) GE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l >= r
}

func (o *StringOperation) LE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l <= r
}

func (o *StringOperation) CO(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return false
	}
	return strings.Contains(l, r)
}

func (o *StringOperation) SW(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return false
	}
	return strings.HasPrefix(l, r)
}

func (o *StringOperation) EW(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return false
	}
	return strings.HasSuffix(l, r)
}

func (o *StringOperation) IN(left Operand, right Operand) bool {
	leftVal, ok := left.(string)
	if !ok {
		return ok
	}

	rightVal, ok := right.([]string)
	if !ok {
		return ok
	}
	for _, val := range rightVal {
		if leftVal == val {
			return true
		}
	}
	return false
}
