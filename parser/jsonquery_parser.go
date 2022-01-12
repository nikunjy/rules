// Code generated from JsonQuery.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JsonQuery

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 35, 120,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 5, 2, 25,
	10, 2, 3, 2, 5, 2, 28, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 54, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 7, 2, 61, 10, 2, 12, 2, 14, 2, 64, 11, 2, 3, 3, 3, 3, 5, 3, 68, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 79, 10,
	5, 3, 5, 3, 5, 5, 5, 83, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 88, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 98, 10, 7, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 108, 10, 9, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 118, 10, 11, 3, 11, 2, 3,
	2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 5, 3, 2, 14, 23, 3, 2, 24,
	26, 3, 2, 15, 20, 2, 129, 2, 53, 3, 2, 2, 2, 4, 65, 3, 2, 2, 2, 6, 69,
	3, 2, 2, 2, 8, 87, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12, 97, 3, 2, 2, 2,
	14, 99, 3, 2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 109, 3, 2, 2, 2, 20, 117,
	3, 2, 2, 2, 22, 24, 8, 2, 1, 2, 23, 25, 7, 10, 2, 2, 24, 23, 3, 2, 2, 2,
	24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 28, 7, 35, 2, 2, 27, 26, 3,
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 30, 7, 3, 2, 2, 30,
	31, 5, 2, 2, 2, 31, 32, 7, 4, 2, 2, 32, 54, 3, 2, 2, 2, 33, 34, 5, 4, 3,
	2, 34, 35, 7, 35, 2, 2, 35, 36, 7, 5, 2, 2, 36, 54, 3, 2, 2, 2, 37, 38,
	5, 4, 3, 2, 38, 39, 7, 35, 2, 2, 39, 40, 9, 2, 2, 2, 40, 41, 7, 35, 2,
	2, 41, 42, 5, 8, 5, 2, 42, 54, 3, 2, 2, 2, 43, 44, 9, 3, 2, 2, 44, 45,
	7, 35, 2, 2, 45, 46, 7, 3, 2, 2, 46, 47, 5, 4, 3, 2, 47, 48, 7, 4, 2, 2,
	48, 49, 7, 35, 2, 2, 49, 50, 9, 4, 2, 2, 50, 51, 7, 35, 2, 2, 51, 52, 5,
	8, 5, 2, 52, 54, 3, 2, 2, 2, 53, 22, 3, 2, 2, 2, 53, 33, 3, 2, 2, 2, 53,
	37, 3, 2, 2, 2, 53, 43, 3, 2, 2, 2, 54, 62, 3, 2, 2, 2, 55, 56, 12, 6,
	2, 2, 56, 57, 7, 35, 2, 2, 57, 58, 7, 11, 2, 2, 58, 59, 7, 35, 2, 2, 59,
	61, 5, 2, 2, 7, 60, 55, 3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2,
	2, 62, 63, 3, 2, 2, 2, 63, 3, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 67, 7,
	27, 2, 2, 66, 68, 5, 6, 4, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68,
	5, 3, 2, 2, 2, 69, 70, 7, 6, 2, 2, 70, 71, 5, 4, 3, 2, 71, 7, 3, 2, 2,
	2, 72, 88, 7, 12, 2, 2, 73, 88, 7, 13, 2, 2, 74, 88, 7, 28, 2, 2, 75, 88,
	7, 29, 2, 2, 76, 88, 7, 30, 2, 2, 77, 79, 7, 7, 2, 2, 78, 77, 3, 2, 2,
	2, 78, 79, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 7, 31, 2, 2, 81, 83,
	7, 32, 2, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 88, 3, 2, 2, 2,
	84, 88, 5, 18, 10, 2, 85, 88, 5, 14, 8, 2, 86, 88, 5, 10, 6, 2, 87, 72,
	3, 2, 2, 2, 87, 73, 3, 2, 2, 2, 87, 74, 3, 2, 2, 2, 87, 75, 3, 2, 2, 2,
	87, 76, 3, 2, 2, 2, 87, 78, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 87, 85, 3,
	2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 9, 3, 2, 2, 2, 89, 90, 7, 8, 2, 2, 90,
	91, 5, 12, 7, 2, 91, 11, 3, 2, 2, 2, 92, 93, 7, 29, 2, 2, 93, 94, 7, 34,
	2, 2, 94, 98, 5, 12, 7, 2, 95, 96, 7, 29, 2, 2, 96, 98, 7, 9, 2, 2, 97,
	92, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 98, 13, 3, 2, 2, 2, 99, 100, 7, 8,
	2, 2, 100, 101, 5, 16, 9, 2, 101, 15, 3, 2, 2, 2, 102, 103, 7, 30, 2, 2,
	103, 104, 7, 34, 2, 2, 104, 108, 5, 16, 9, 2, 105, 106, 7, 30, 2, 2, 106,
	108, 7, 9, 2, 2, 107, 102, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 17, 3,
	2, 2, 2, 109, 110, 7, 8, 2, 2, 110, 111, 5, 20, 11, 2, 111, 19, 3, 2, 2,
	2, 112, 113, 7, 31, 2, 2, 113, 114, 7, 34, 2, 2, 114, 118, 5, 20, 11, 2,
	115, 116, 7, 31, 2, 2, 116, 118, 7, 9, 2, 2, 117, 112, 3, 2, 2, 2, 117,
	115, 3, 2, 2, 2, 118, 21, 3, 2, 2, 2, 13, 24, 27, 53, 62, 67, 78, 82, 87,
	97, 107, 117,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "'null'",
	"", "", "", "", "", "", "", "", "", "", "'MLP'", "'SUM'", "'SUBTRACT'",
	"", "", "", "", "", "", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN", "NULL",
	"IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW", "ASTERISK",
	"PLUS", "MINUS", "ATTRNAME", "VERSION", "STRING", "DOUBLE", "INT", "EXP",
	"NEWLINE", "COMMA", "SP",
}

var ruleNames = []string{
	"query", "attrPath", "subAttr", "value", "listStrings", "subListOfStrings",
	"listDoubles", "subListOfDoubles", "listInts", "subListOfInts",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JsonQueryParser struct {
	*antlr.BaseParser
}

func NewJsonQueryParser(input antlr.TokenStream) *JsonQueryParser {
	this := new(JsonQueryParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JsonQuery.g4"

	return this
}

// JsonQueryParser tokens.
const (
	JsonQueryParserEOF              = antlr.TokenEOF
	JsonQueryParserT__0             = 1
	JsonQueryParserT__1             = 2
	JsonQueryParserT__2             = 3
	JsonQueryParserT__3             = 4
	JsonQueryParserT__4             = 5
	JsonQueryParserT__5             = 6
	JsonQueryParserT__6             = 7
	JsonQueryParserNOT              = 8
	JsonQueryParserLOGICAL_OPERATOR = 9
	JsonQueryParserBOOLEAN          = 10
	JsonQueryParserNULL             = 11
	JsonQueryParserIN               = 12
	JsonQueryParserEQ               = 13
	JsonQueryParserNE               = 14
	JsonQueryParserGT               = 15
	JsonQueryParserLT               = 16
	JsonQueryParserGE               = 17
	JsonQueryParserLE               = 18
	JsonQueryParserCO               = 19
	JsonQueryParserSW               = 20
	JsonQueryParserEW               = 21
	JsonQueryParserASTERISK         = 22
	JsonQueryParserPLUS             = 23
	JsonQueryParserMINUS            = 24
	JsonQueryParserATTRNAME         = 25
	JsonQueryParserVERSION          = 26
	JsonQueryParserSTRING           = 27
	JsonQueryParserDOUBLE           = 28
	JsonQueryParserINT              = 29
	JsonQueryParserEXP              = 30
	JsonQueryParserNEWLINE          = 31
	JsonQueryParserCOMMA            = 32
	JsonQueryParserSP               = 33
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query            = 0
	JsonQueryParserRULE_attrPath         = 1
	JsonQueryParserRULE_subAttr          = 2
	JsonQueryParserRULE_value            = 3
	JsonQueryParserRULE_listStrings      = 4
	JsonQueryParserRULE_subListOfStrings = 5
	JsonQueryParserRULE_listDoubles      = 6
	JsonQueryParserRULE_subListOfDoubles = 7
	JsonQueryParserRULE_listInts         = 8
	JsonQueryParserRULE_subListOfInts    = 9
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulSumExpContext struct {
	*QueryContext
	op antlr.Token
}

func NewMulSumExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulSumExpContext {
	var p = new(MulSumExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *MulSumExpContext) GetOp() antlr.Token { return s.op }

func (s *MulSumExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulSumExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulSumExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *MulSumExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *MulSumExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *MulSumExpContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *MulSumExpContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserASTERISK, 0)
}

func (s *MulSumExpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserPLUS, 0)
}

func (s *MulSumExpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserMINUS, 0)
}

func (s *MulSumExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *MulSumExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *MulSumExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *MulSumExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *MulSumExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *MulSumExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *MulSumExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitMulSumExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	*QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNOT, 0)
}

func (s *ParenExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	*QueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicalExpContext struct {
	*QueryContext
}

func NewLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpContext {
	var p = new(LogicalExpContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *LogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *LogicalExpContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *LogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *LogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *LogicalExpContext) LOGICAL_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLOGICAL_OPERATOR, 0)
}

func (s *LogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *JsonQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, JsonQueryParserRULE_query, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(21)
				p.Match(JsonQueryParserNOT)
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(24)
				p.Match(JsonQueryParserSP)
			}

		}
		{
			p.SetState(27)
			p.Match(JsonQueryParserT__0)
		}
		{
			p.SetState(28)
			p.query(0)
		}
		{
			p.SetState(29)
			p.Match(JsonQueryParserT__1)
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.AttrPath()
		}
		{
			p.SetState(32)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(33)
			p.Match(JsonQueryParserT__2)
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.AttrPath()
		}
		{
			p.SetState(36)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(37)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonQueryParserIN)|(1<<JsonQueryParserEQ)|(1<<JsonQueryParserNE)|(1<<JsonQueryParserGT)|(1<<JsonQueryParserLT)|(1<<JsonQueryParserGE)|(1<<JsonQueryParserLE)|(1<<JsonQueryParserCO)|(1<<JsonQueryParserSW)|(1<<JsonQueryParserEW))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(38)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(39)
			p.Value()
		}

	case 4:
		localctx = NewMulSumExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonQueryParserASTERISK)|(1<<JsonQueryParserPLUS)|(1<<JsonQueryParserMINUS))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(42)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(43)
			p.Match(JsonQueryParserT__0)
		}
		{
			p.SetState(44)
			p.AttrPath()
		}
		{
			p.SetState(45)
			p.Match(JsonQueryParserT__1)
		}
		{
			p.SetState(46)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(47)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MulSumExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonQueryParserEQ)|(1<<JsonQueryParserNE)|(1<<JsonQueryParserGT)|(1<<JsonQueryParserLT)|(1<<JsonQueryParserGE)|(1<<JsonQueryParserLE))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MulSumExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(48)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(49)
			p.Value()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
			p.SetState(53)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(54)
				p.Match(JsonQueryParserSP)
			}
			{
				p.SetState(55)
				p.Match(JsonQueryParserLOGICAL_OPERATOR)
			}
			{
				p.SetState(56)
				p.Match(JsonQueryParserSP)
			}
			{
				p.SetState(57)
				p.query(5)
			}

		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
	return p
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonQueryParserRULE_attrPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(JsonQueryParserATTRNAME)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(64)
			p.SubAttr()
		}

	}

	return localctx
}

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
	return p
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonQueryParserRULE_subAttr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(JsonQueryParserT__3)
	}
	{
		p.SetState(68)
		p.AttrPath()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListOfDoublesContext struct {
	*ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfStringsContext struct {
	*ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	*ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	*ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNULL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	*ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionContext struct {
	*ValueContext
}

func NewVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionContext {
	var p = new(VersionContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserVERSION, 0)
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

type LongContext struct {
	*ValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEXP, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfIntsContext struct {
	*ValueContext
}

func NewListOfIntsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntsContext {
	var p = new(ListOfIntsContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntsContext) ListInts() IListIntsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListIntsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *ListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonQueryParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(JsonQueryParserBOOLEAN)
		}

	case 2:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Match(JsonQueryParserNULL)
		}

	case 3:
		localctx = NewVersionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(JsonQueryParserVERSION)
		}

	case 4:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Match(JsonQueryParserSTRING)
		}

	case 5:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.Match(JsonQueryParserDOUBLE)
		}

	case 6:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__4 {
			{
				p.SetState(75)
				p.Match(JsonQueryParserT__4)
			}

		}
		{
			p.SetState(78)
			p.Match(JsonQueryParserINT)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(79)
				p.Match(JsonQueryParserEXP)
			}

		}

	case 7:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(82)
			p.ListInts()
		}

	case 8:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(83)
			p.ListDoubles()
		}

	case 9:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(84)
			p.ListStrings()
		}

	}

	return localctx
}

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
	return p
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_listStrings)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(JsonQueryParserT__5)
	}
	{
		p.SetState(88)
		p.SubListOfStrings()
	}

	return localctx
}

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
	return p
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfStringsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonQueryParserRULE_subListOfStrings)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(JsonQueryParserSTRING)
		}
		{
			p.SetState(91)
			p.Match(JsonQueryParserCOMMA)
		}
		{
			p.SetState(92)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(JsonQueryParserSTRING)
		}
		{
			p.SetState(94)
			p.Match(JsonQueryParserT__6)
		}

	}

	return localctx
}

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
	return p
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonQueryParserRULE_listDoubles)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(JsonQueryParserT__5)
	}
	{
		p.SetState(98)
		p.SubListOfDoubles()
	}

	return localctx
}

// ISubListOfDoublesContext is an interface to support dynamic dispatch.
type ISubListOfDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfDoublesContext differentiates from other interfaces.
	IsSubListOfDoublesContext()
}

type SubListOfDoublesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfDoublesContext() *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
	return p
}

func (*SubListOfDoublesContext) IsSubListOfDoublesContext() {}

func NewSubListOfDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles

	return p
}

func (s *SubListOfDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfDoublesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *SubListOfDoublesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfDoublesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *SubListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfDoubles() (localctx ISubListOfDoublesContext) {
	localctx = NewSubListOfDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_subListOfDoubles)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(JsonQueryParserDOUBLE)
		}
		{
			p.SetState(101)
			p.Match(JsonQueryParserCOMMA)
		}
		{
			p.SetState(102)
			p.SubListOfDoubles()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(JsonQueryParserDOUBLE)
		}
		{
			p.SetState(104)
			p.Match(JsonQueryParserT__6)
		}

	}

	return localctx
}

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
	return p
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfIntsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_listInts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(JsonQueryParserT__5)
	}
	{
		p.SetState(108)
		p.SubListOfInts()
	}

	return localctx
}

// ISubListOfIntsContext is an interface to support dynamic dispatch.
type ISubListOfIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubListOfIntsContext differentiates from other interfaces.
	IsSubListOfIntsContext()
}

type SubListOfIntsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIntsContext() *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
	return p
}

func (*SubListOfIntsContext) IsSubListOfIntsContext() {}

func NewSubListOfIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfInts

	return p
}

func (s *SubListOfIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIntsContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *SubListOfIntsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubListOfIntsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *SubListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfInts() (localctx ISubListOfIntsContext) {
	localctx = NewSubListOfIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonQueryParserRULE_subListOfInts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.Match(JsonQueryParserINT)
		}
		{
			p.SetState(111)
			p.Match(JsonQueryParserCOMMA)
		}
		{
			p.SetState(112)
			p.SubListOfInts()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(JsonQueryParserINT)
		}
		{
			p.SetState(114)
			p.Match(JsonQueryParserT__6)
		}

	}

	return localctx
}

func (p *JsonQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
