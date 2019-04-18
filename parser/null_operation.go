package parser

type NullOperation struct {
}

func (o *NullOperation) EQ(left Operand, right Operand) bool {
	return left == nil
}

func (o *NullOperation) NE(left Operand, right Operand) bool {
	return left != nil
}

func (o *NullOperation) GT(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) LT(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) GE(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) LE(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) CO(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) SW(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) EW(left Operand, right Operand) bool {
	panic("not supported")
}

func (o *NullOperation) IN(left Operand, right Operand) bool {
	panic("not supported")
}
