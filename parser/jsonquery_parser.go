// Code generated from JsonQuery.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 30, 79, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 2, 5, 2, 20, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 36, 10,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 43, 10, 2, 12, 2, 14, 2, 46, 11,
	2, 3, 3, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 60, 10, 5, 3, 5, 3, 5, 5, 5, 64, 10, 5, 3, 5, 5, 5, 67,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 77, 10, 7,
	3, 7, 2, 3, 2, 8, 2, 4, 6, 8, 10, 12, 2, 3, 3, 2, 14, 23, 2, 86, 2, 35,
	3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 51, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10,
	68, 3, 2, 2, 2, 12, 76, 3, 2, 2, 2, 14, 16, 8, 2, 1, 2, 15, 17, 7, 10,
	2, 2, 16, 15, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 20,
	7, 30, 2, 2, 19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2,
	21, 22, 7, 3, 2, 2, 22, 23, 5, 2, 2, 2, 23, 24, 7, 4, 2, 2, 24, 36, 3,
	2, 2, 2, 25, 26, 5, 4, 3, 2, 26, 27, 7, 30, 2, 2, 27, 28, 7, 5, 2, 2, 28,
	36, 3, 2, 2, 2, 29, 30, 5, 4, 3, 2, 30, 31, 7, 30, 2, 2, 31, 32, 9, 2,
	2, 2, 32, 33, 7, 30, 2, 2, 33, 34, 5, 8, 5, 2, 34, 36, 3, 2, 2, 2, 35,
	14, 3, 2, 2, 2, 35, 25, 3, 2, 2, 2, 35, 29, 3, 2, 2, 2, 36, 44, 3, 2, 2,
	2, 37, 38, 12, 5, 2, 2, 38, 39, 7, 30, 2, 2, 39, 40, 7, 11, 2, 2, 40, 41,
	7, 30, 2, 2, 41, 43, 5, 2, 2, 6, 42, 37, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2,
	44, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 3, 3, 2, 2, 2, 46, 44, 3, 2,
	2, 2, 47, 49, 7, 24, 2, 2, 48, 50, 5, 6, 4, 2, 49, 48, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 6, 2, 2, 52, 53, 5, 4, 3,
	2, 53, 7, 3, 2, 2, 2, 54, 67, 7, 12, 2, 2, 55, 67, 7, 13, 2, 2, 56, 67,
	7, 25, 2, 2, 57, 67, 7, 26, 2, 2, 58, 60, 7, 7, 2, 2, 59, 58, 3, 2, 2,
	2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63, 7, 27, 2, 2, 62, 64,
	7, 28, 2, 2, 63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2,
	65, 67, 5, 10, 6, 2, 66, 54, 3, 2, 2, 2, 66, 55, 3, 2, 2, 2, 66, 56, 3,
	2, 2, 2, 66, 57, 3, 2, 2, 2, 66, 59, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67,
	9, 3, 2, 2, 2, 68, 69, 7, 8, 2, 2, 69, 70, 5, 12, 7, 2, 70, 11, 3, 2, 2,
	2, 71, 72, 7, 27, 2, 2, 72, 73, 7, 30, 2, 2, 73, 77, 5, 12, 7, 2, 74, 75,
	7, 27, 2, 2, 75, 77, 7, 9, 2, 2, 76, 71, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2,
	77, 13, 3, 2, 2, 2, 11, 16, 19, 35, 44, 49, 59, 63, 66, 76,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "'null'",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'\n'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN", "NULL",
	"IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW", "ATTRNAME",
	"STRING", "DOUBLE", "INT", "EXP", "NEWLINE", "SP",
}

var ruleNames = []string{
	"query", "attrPath", "subAttr", "value", "listInts", "subListOfInts",
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
	JsonQueryParserATTRNAME         = 22
	JsonQueryParserSTRING           = 23
	JsonQueryParserDOUBLE           = 24
	JsonQueryParserINT              = 25
	JsonQueryParserEXP              = 26
	JsonQueryParserNEWLINE          = 27
	JsonQueryParserSP               = 28
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query         = 0
	JsonQueryParserRULE_attrPath      = 1
	JsonQueryParserRULE_subAttr       = 2
	JsonQueryParserRULE_value         = 3
	JsonQueryParserRULE_listInts      = 4
	JsonQueryParserRULE_subListOfInts = 5
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(13)
				p.Match(JsonQueryParserNOT)
			}

		}
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(16)
				p.Match(JsonQueryParserSP)
			}

		}
		{
			p.SetState(19)
			p.Match(JsonQueryParserT__0)
		}
		{
			p.SetState(20)
			p.query(0)
		}
		{
			p.SetState(21)
			p.Match(JsonQueryParserT__1)
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.AttrPath()
		}
		{
			p.SetState(24)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(25)
			p.Match(JsonQueryParserT__2)
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.AttrPath()
		}
		{
			p.SetState(28)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(29)

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
			p.SetState(30)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(31)
			p.Value()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(42)
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
			p.SetState(35)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(36)
				p.Match(JsonQueryParserSP)
			}
			{
				p.SetState(37)
				p.Match(JsonQueryParserLOGICAL_OPERATOR)
			}
			{
				p.SetState(38)
				p.Match(JsonQueryParserSP)
			}
			{
				p.SetState(39)
				p.query(4)
			}

		}
		p.SetState(44)
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
		p.SetState(45)
		p.Match(JsonQueryParserATTRNAME)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(46)
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
		p.SetState(49)
		p.Match(JsonQueryParserT__3)
	}
	{
		p.SetState(50)
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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonQueryParserBOOLEAN:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(JsonQueryParserBOOLEAN)
		}

	case JsonQueryParserNULL:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Match(JsonQueryParserNULL)
		}

	case JsonQueryParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(JsonQueryParserSTRING)
		}

	case JsonQueryParserDOUBLE:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(55)
			p.Match(JsonQueryParserDOUBLE)
		}

	case JsonQueryParserT__4, JsonQueryParserINT:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__4 {
			{
				p.SetState(56)
				p.Match(JsonQueryParserT__4)
			}

		}
		{
			p.SetState(59)
			p.Match(JsonQueryParserINT)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(60)
				p.Match(JsonQueryParserEXP)
			}

		}

	case JsonQueryParserT__5:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.ListInts()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 8, JsonQueryParserRULE_listInts)

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
		p.SetState(66)
		p.Match(JsonQueryParserT__5)
	}
	{
		p.SetState(67)
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

func (s *SubListOfIntsContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
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
	p.EnterRule(localctx, 10, JsonQueryParserRULE_subListOfInts)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Match(JsonQueryParserINT)
		}
		{
			p.SetState(70)
			p.Match(JsonQueryParserSP)
		}
		{
			p.SetState(71)
			p.SubListOfInts()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.Match(JsonQueryParserINT)
		}
		{
			p.SetState(73)
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
