package parser

import (
	"errors"
	"fmt"
	"reflect"
)

type Operand interface{}

var (
	ErrInvalidOperation   = errors.New("Invalid operation on the type")
	ErrEvalOperandMissing = errors.New("Operand not present")
)

type ErrInvalidOperand struct {
	Val     interface{}
	typeObj interface{}
}

func newErrInvalidOperand(val Operand, typeObj interface{}) *ErrInvalidOperand {
	return &ErrInvalidOperand{
		Val:     val,
		typeObj: typeObj,
	}
}

func (e *ErrInvalidOperand) Error() string {
	return fmt.Sprintf("Operand %v is not of type %s", e.Val, reflect.TypeOf(e.typeObj).String())
}

type Operation interface {
	EQ(left Operand, right Operand) (bool, error)
	NE(left Operand, right Operand) (bool, error)
	GT(left Operand, right Operand) (bool, error)
	LT(left Operand, right Operand) (bool, error)
	GE(left Operand, right Operand) (bool, error)
	LE(left Operand, right Operand) (bool, error)
	CO(left Operand, right Operand) (bool, error)
	SW(left Operand, right Operand) (bool, error)
	EW(left Operand, right Operand) (bool, error)
	IN(left Operand, right Operand) (bool, error)
}
