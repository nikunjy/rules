package parser

type FloatOperation struct {
	NullOperation
}

func toNum(op Operand) (float64, bool) {
	switch val := op.(type) {
	case int:
		return float64(val), true
	case float64:
		return val, true
	case int64:
		return float64(val), true
	}
	return 0, false
}

func (o *FloatOperation) get(left Operand, right Operand) (float64, float64, bool) {
	if left == nil {
		return 0, 0, false
	}
	leftVal, ok := toNum(left)
	rightVal, ok1 := toNum(right)
	return leftVal, rightVal, (ok1 && ok)

}

func (o *FloatOperation) EQ(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l == r
}

func (o *FloatOperation) NE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l != r
}

func (o *FloatOperation) GT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l > r
}

func (o *FloatOperation) LT(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l < r
}

func (o *FloatOperation) GE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l >= r
}

func (o *FloatOperation) LE(left Operand, right Operand) bool {
	l, r, ok := o.get(left, right)
	if !ok {
		return ok
	}
	return l <= r
}

func (o *FloatOperation) IN(left Operand, right Operand) bool {
	leftVal, ok := toNum(left)
	if !ok {
		return ok
	}
	rightVal, ok := right.([]float64)
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
