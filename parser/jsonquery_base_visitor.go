// Code generated from JsonQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJsonQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonQueryVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitLong(ctx *LongContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListInts(ctx *ListIntsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
	return v.VisitChildren(ctx)
}
