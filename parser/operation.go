package parser

type Operand interface{}

type Operation interface {
	EQ(left Operand, right Operand) bool
	NE(left Operand, right Operand) bool
	GT(left Operand, right Operand) bool
	LT(left Operand, right Operand) bool
	GE(left Operand, right Operand) bool
	LE(left Operand, right Operand) bool
	CO(left Operand, right Operand) bool
	SW(left Operand, right Operand) bool
	EW(left Operand, right Operand) bool
}
