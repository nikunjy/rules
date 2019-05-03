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

func (v *VersionOperation) EQ(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.EQ(rightVer), nil
}

func (v *VersionOperation) NE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.NE(rightVer), nil
}

func (v *VersionOperation) LT(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.LT(rightVer), nil
}

func (v *VersionOperation) GT(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.GT(rightVer), nil
}

func (v *VersionOperation) LE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.LE(rightVer), nil
}
func (v *VersionOperation) GE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, ok := v.get(left, right)
	if !ok {
		return ok, nil
	}
	return leftVer.GE(rightVer), nil
}
