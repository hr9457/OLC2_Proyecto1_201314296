// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 22, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 2, 2, 4, 2, 4, 2, 2,
	2, 20, 2, 6, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7,
	2, 2, 3, 8, 3, 3, 2, 2, 2, 9, 10, 7, 32, 2, 2, 10, 11, 7, 3, 2, 2, 11,
	12, 7, 59, 2, 2, 12, 13, 7, 4, 2, 2, 13, 20, 7, 9, 2, 2, 14, 15, 7, 32,
	2, 2, 15, 16, 7, 3, 2, 2, 16, 17, 7, 62, 2, 2, 17, 18, 7, 4, 2, 2, 18,
	20, 7, 9, 2, 2, 19, 9, 3, 2, 2, 2, 19, 14, 3, 2, 2, 2, 20, 5, 3, 2, 2,
	2, 3, 19,
}
var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", "','", "'<'",
	"'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'",
	"'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'?'", "'->'", "'println!'",
	"'i64'", "'f64'", "'bool'", "'char'", "'String'", "'usize'", "'let'", "'mut'",
	"'struct'", "'as'", "'true'", "'false'", "'fn'", "'return'", "'abs'", "'sqrt'",
	"'to_string'", "'clone'", "'new'", "'len'", "'push'", "'remove'", "'contains'",
	"'insert'", "'capacity'", "'with_capacity'",
}
var symbolicNames = []string{
	"", "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", "TK_CORCHETE_RIGHT",
	"TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA", "TK_DOSPUNTOS", "TK_COMA",
	"TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL", "TK_MAS", "TK_MENOS", "TK_POR",
	"TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA", "TK_AMPERSAND", "TK_ADMIRACION",
	"TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL", "TK_DIFERENTE", "TK_OR",
	"TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION", "TK_INT", "TK_FLOAT",
	"TK_BOOL", "TK_CHAR", "TK_STRING", "TK_USIZE", "TK_LET", "TK_MUT", "TK_STRUCT",
	"TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN", "TK_RETURN", "TK_ABS", "TK_SQRT",
	"TK_TOSTRING", "TK_CLONE", "TK_NEW", "TK_LEN", "TK_PUSH", "TK_REMOVED",
	"TK_CONTAINS", "TK_INSERT", "TK_CAPACITY", "TK_WITHCAPACITY", "TK_NUMBER",
	"TK_DECIMAL", "TK_CADENA", "TK_ID", "TK_COMMET", "SPACES",
}

var ruleNames = []string{
	"start", "impresion",
}

type rustParser struct {
	*antlr.BaseParser
}

// NewrustParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *rustParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewrustParser(input antlr.TokenStream) *rustParser {
	this := new(rustParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "rust.g4"

	return this
}

// rustParser tokens.
const (
	rustParserEOF                 = antlr.TokenEOF
	rustParserTK_PARENTESIS_LEFT  = 1
	rustParserTK_PARENTESIS_RIGHT = 2
	rustParserTK_CORCHETE_LETF    = 3
	rustParserTK_CORCHETE_RIGHT   = 4
	rustParserTK_LLAVE_LEFT       = 5
	rustParserTK_LLAVE_RIGHT      = 6
	rustParserTK_PUNTO_COMA       = 7
	rustParserTK_DOSPUNTOS        = 8
	rustParserTK_COMA             = 9
	rustParserTK_MENOR            = 10
	rustParserTK_MAYOR            = 11
	rustParserTK_PUNTO            = 12
	rustParserTK_IGUAL            = 13
	rustParserTK_MAS              = 14
	rustParserTK_MENOS            = 15
	rustParserTK_POR              = 16
	rustParserTK_DIVISION         = 17
	rustParserTK_PORCENTAJE       = 18
	rustParserTK_BARRA            = 19
	rustParserTK_AMPERSAND        = 20
	rustParserTK_ADMIRACION       = 21
	rustParserTK_MAYORIGULA       = 22
	rustParserTK_MENOIGUAL        = 23
	rustParserTK_IGUALIGUAL       = 24
	rustParserTK_DIFERENTE        = 25
	rustParserTK_OR               = 26
	rustParserTK_AND              = 27
	rustParserTK_WHAT             = 28
	rustParserTK_TIPORETURN       = 29
	rustParserTK_IMPRESION        = 30
	rustParserTK_INT              = 31
	rustParserTK_FLOAT            = 32
	rustParserTK_BOOL             = 33
	rustParserTK_CHAR             = 34
	rustParserTK_STRING           = 35
	rustParserTK_USIZE            = 36
	rustParserTK_LET              = 37
	rustParserTK_MUT              = 38
	rustParserTK_STRUCT           = 39
	rustParserTK_AS               = 40
	rustParserTK_TRUE             = 41
	rustParserTK_FALSE            = 42
	rustParserTK_FN               = 43
	rustParserTK_RETURN           = 44
	rustParserTK_ABS              = 45
	rustParserTK_SQRT             = 46
	rustParserTK_TOSTRING         = 47
	rustParserTK_CLONE            = 48
	rustParserTK_NEW              = 49
	rustParserTK_LEN              = 50
	rustParserTK_PUSH             = 51
	rustParserTK_REMOVED          = 52
	rustParserTK_CONTAINS         = 53
	rustParserTK_INSERT           = 54
	rustParserTK_CAPACITY         = 55
	rustParserTK_WITHCAPACITY     = 56
	rustParserTK_NUMBER           = 57
	rustParserTK_DECIMAL          = 58
	rustParserTK_CADENA           = 59
	rustParserTK_ID               = 60
	rustParserTK_COMMET           = 61
	rustParserSPACES              = 62
)

// rustParser rules.
const (
	rustParserRULE_start     = 0
	rustParserRULE_impresion = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(rustParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *rustParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, rustParserRULE_start)

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
		p.SetState(4)
		p.Impresion()
	}
	{
		p.SetState(5)
		p.Match(rustParserEOF)
	}

	return localctx
}

// IImpresionContext is an interface to support dynamic dispatch.
type IImpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpresionContext differentiates from other interfaces.
	IsImpresionContext()
}

type ImpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpresionContext() *ImpresionContext {
	var p = new(ImpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_impresion
	return p
}

func (*ImpresionContext) IsImpresionContext() {}

func NewImpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpresionContext {
	var p = new(ImpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_impresion

	return p
}

func (s *ImpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpresionContext) TK_IMPRESION() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IMPRESION, 0)
}

func (s *ImpresionContext) TK_PARENTESIS_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_LEFT, 0)
}

func (s *ImpresionContext) TK_NUMBER() antlr.TerminalNode {
	return s.GetToken(rustParserTK_NUMBER, 0)
}

func (s *ImpresionContext) TK_PARENTESIS_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_RIGHT, 0)
}

func (s *ImpresionContext) TK_PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, 0)
}

func (s *ImpresionContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *ImpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterImpresion(s)
	}
}

func (s *ImpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitImpresion(s)
	}
}

func (p *rustParser) Impresion() (localctx IImpresionContext) {
	this := p
	_ = this

	localctx = NewImpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, rustParserRULE_impresion)

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

	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(7)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(8)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(9)
			p.Match(rustParserTK_NUMBER)
		}
		{
			p.SetState(10)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(11)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(12)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(13)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(14)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(15)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(16)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	}

	return localctx
}
