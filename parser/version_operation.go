package parser

import (
	"github.com/blang/semver"
)

type VersionOperation struct {
	NullOperation
}

func (v *VersionOperation) get(left Operand, right Operand) (semver.Version, semver.Version, bool) {
	var leftVer, rightVer semver.Version
	leftVal, ok := left.(string)
	if !ok {
		return leftVer, rightVer, ok
	}
	rightVal, ok := right.(string)
	if !ok {
		return leftVer, rightVer, ok
	}
	var err error
	leftVer, err = semver.Make(leftVal)
	if err != nil {
		// TODO fix error
		return leftVer, rightVer, false
	}
	rightVer, err = semver.Make(rightVal)
	if err != nil {
		// TODO fix error
		return leftVer, rightVer, false
	}
	return leftVer, rightVer, true
}

func (v *VersionOperation) EQ(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.EQ(rightVer)
}

func (v *VersionOperation) NE(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.NE(rightVer)
}

func (v *VersionOperation) LT(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.LT(rightVer)
}

func (v *VersionOperation) GT(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.GT(rightVer)
}

func (v *VersionOperation) LE(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.LE(rightVer)
}
func (v *VersionOperation) GE(left Operand, right Operand) bool {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok
	}
	return leftVer.GE(rightVer)
}
