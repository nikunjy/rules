package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type StringOperation struct {
	NullOperation
	RegexPatternMap map[string]*regexp.Regexp
}

func (o *StringOperation) getString(operand Operand) (string, error) {
	switch opVal := operand.(type) {
	case string:
		return opVal, nil
	case fmt.Stringer:
		return opVal.String(), nil
	default:
		return "", newErrInvalidOperand(operand, opVal)
	}
}

func (o *StringOperation) get(left Operand, right Operand) (string, string, error) {
	if left == nil {
		return "", "", ErrEvalOperandMissing
	}
	leftVal, err := o.getString(left)
	if err != nil {
		return "", "", err
	}
	rightVal, err := o.getString(right)
	if err != nil {
		return "", "", err
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
	leftVal, err := o.getString(left)
	if err != nil {
		return false, err
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

func (o *StringOperation) LIKE(left Operand, right Operand) (bool, error) {
	l, r, err := o.get(left, right)
	if err != nil {
		return false, err
	}

	if o.RegexPatternMap == nil {
		o.RegexPatternMap = make(map[string]*regexp.Regexp)
	}

	if regx, ok := o.RegexPatternMap[r]; ok {
		match := regx.MatchString(l)
		return match, nil
	}

	regx, err := regexp.Compile(r)
	if err != nil {
		return false, err
	}

	o.RegexPatternMap[r] = regx
	match := regx.MatchString(l)
	return match, nil
}
