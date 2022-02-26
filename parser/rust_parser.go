// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Proyecto1/src/symbol"
import "Proyecto1/src/interfaces"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 125,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 68, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 80, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 97, 10, 6, 12, 6, 14, 6, 100, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	5, 7, 116, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 2, 3,
	10, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 136, 2, 19, 3, 2, 2, 2, 4, 31,
	3, 2, 2, 2, 6, 67, 3, 2, 2, 2, 8, 79, 3, 2, 2, 2, 10, 81, 3, 2, 2, 2, 12,
	115, 3, 2, 2, 2, 14, 117, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2,
	2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22,
	3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 23, 7, 2, 2, 3, 23, 24, 8, 2, 1, 2,
	24, 3, 3, 2, 2, 2, 25, 26, 5, 6, 4, 2, 26, 27, 8, 3, 1, 2, 27, 32, 3, 2,
	2, 2, 28, 29, 5, 14, 8, 2, 29, 30, 8, 3, 1, 2, 30, 32, 3, 2, 2, 2, 31,
	25, 3, 2, 2, 2, 31, 28, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 7, 39, 2,
	2, 34, 35, 7, 40, 2, 2, 35, 36, 7, 63, 2, 2, 36, 37, 7, 15, 2, 2, 37, 38,
	5, 10, 6, 2, 38, 39, 7, 9, 2, 2, 39, 40, 8, 4, 1, 2, 40, 68, 3, 2, 2, 2,
	41, 42, 7, 39, 2, 2, 42, 43, 7, 63, 2, 2, 43, 44, 7, 15, 2, 2, 44, 45,
	5, 10, 6, 2, 45, 46, 7, 9, 2, 2, 46, 47, 8, 4, 1, 2, 47, 68, 3, 2, 2, 2,
	48, 49, 7, 39, 2, 2, 49, 50, 7, 40, 2, 2, 50, 51, 7, 63, 2, 2, 51, 52,
	7, 10, 2, 2, 52, 53, 5, 8, 5, 2, 53, 54, 7, 15, 2, 2, 54, 55, 5, 10, 6,
	2, 55, 56, 7, 9, 2, 2, 56, 57, 8, 4, 1, 2, 57, 68, 3, 2, 2, 2, 58, 59,
	7, 39, 2, 2, 59, 60, 7, 63, 2, 2, 60, 61, 7, 10, 2, 2, 61, 62, 5, 8, 5,
	2, 62, 63, 7, 15, 2, 2, 63, 64, 5, 10, 6, 2, 64, 65, 7, 9, 2, 2, 65, 66,
	8, 4, 1, 2, 66, 68, 3, 2, 2, 2, 67, 33, 3, 2, 2, 2, 67, 41, 3, 2, 2, 2,
	67, 48, 3, 2, 2, 2, 67, 58, 3, 2, 2, 2, 68, 7, 3, 2, 2, 2, 69, 70, 7, 33,
	2, 2, 70, 80, 8, 5, 1, 2, 71, 72, 7, 34, 2, 2, 72, 80, 8, 5, 1, 2, 73,
	74, 7, 35, 2, 2, 74, 80, 8, 5, 1, 2, 75, 76, 7, 36, 2, 2, 76, 80, 8, 5,
	1, 2, 77, 78, 7, 37, 2, 2, 78, 80, 8, 5, 1, 2, 79, 69, 3, 2, 2, 2, 79,
	71, 3, 2, 2, 2, 79, 73, 3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 79, 77, 3, 2, 2,
	2, 80, 9, 3, 2, 2, 2, 81, 82, 8, 6, 1, 2, 82, 83, 5, 12, 7, 2, 83, 98,
	3, 2, 2, 2, 84, 85, 12, 7, 2, 2, 85, 86, 7, 18, 2, 2, 86, 97, 5, 10, 6,
	8, 87, 88, 12, 6, 2, 2, 88, 89, 7, 19, 2, 2, 89, 97, 5, 10, 6, 7, 90, 91,
	12, 5, 2, 2, 91, 92, 7, 16, 2, 2, 92, 97, 5, 10, 6, 6, 93, 94, 12, 4, 2,
	2, 94, 95, 7, 17, 2, 2, 95, 97, 5, 10, 6, 5, 96, 84, 3, 2, 2, 2, 96, 87,
	3, 2, 2, 2, 96, 90, 3, 2, 2, 2, 96, 93, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2,
	98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 11, 3, 2, 2, 2, 100, 98, 3,
	2, 2, 2, 101, 102, 7, 59, 2, 2, 102, 116, 8, 7, 1, 2, 103, 104, 7, 60,
	2, 2, 104, 116, 8, 7, 1, 2, 105, 106, 7, 43, 2, 2, 106, 116, 8, 7, 1, 2,
	107, 108, 7, 44, 2, 2, 108, 116, 8, 7, 1, 2, 109, 110, 7, 61, 2, 2, 110,
	116, 8, 7, 1, 2, 111, 112, 7, 62, 2, 2, 112, 116, 8, 7, 1, 2, 113, 114,
	7, 63, 2, 2, 114, 116, 8, 7, 1, 2, 115, 101, 3, 2, 2, 2, 115, 103, 3, 2,
	2, 2, 115, 105, 3, 2, 2, 2, 115, 107, 3, 2, 2, 2, 115, 109, 3, 2, 2, 2,
	115, 111, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 13, 3, 2, 2, 2, 117, 118,
	7, 32, 2, 2, 118, 119, 7, 3, 2, 2, 119, 120, 5, 10, 6, 2, 120, 121, 7,
	4, 2, 2, 121, 122, 7, 9, 2, 2, 122, 123, 8, 8, 1, 2, 123, 15, 3, 2, 2,
	2, 9, 19, 31, 67, 79, 96, 98, 115,
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
	"TK_DECIMAL", "TK_CADENA", "TK_CARACTER", "TK_ID", "TK_COMMET", "SPACES",
}

var ruleNames = []string{
	"start", "instrucciones", "variable", "tipo", "expresion", "valor", "impresion",
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

var tipoValor = interfaces.NULL

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
	rustParserTK_CARACTER         = 60
	rustParserTK_ID               = 61
	rustParserTK_COMMET           = 62
	rustParserSPACES              = 63
)

// rustParser rules.
const (
	rustParserRULE_start         = 0
	rustParserRULE_instrucciones = 1
	rustParserRULE_variable      = 2
	rustParserRULE_tipo          = 3
	rustParserRULE_expresion     = 4
	rustParserRULE_valor         = 5
	rustParserRULE_impresion     = 6
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

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(rustParserEOF, 0)
}

func (s *StartContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *StartContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == rustParserTK_IMPRESION || _la == rustParserTK_LET {
		{
			p.SetState(14)
			p.Instrucciones()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(rustParserEOF)
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_variable returns the _variable rule contexts.
	Get_variable() IVariableContext

	// Set_variable sets the _variable rule contexts.
	Set_variable(IVariableContext)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_variable IVariableContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_variable() IVariableContext { return s._variable }

func (s *InstruccionesContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionesContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *InstruccionesContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *rustParser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, rustParserRULE_instrucciones)

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

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)

			var _x = p.Variable()

			localctx.(*InstruccionesContext)._variable = _x
		}
		fmt.Println(localctx.(*InstruccionesContext).Get_variable().GetNuevaVariable())

	case rustParserTK_IMPRESION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Impresion()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_ID sets the _TK_ID token.
	Set_TK_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// GetNuevaVariable returns the nuevaVariable attribute.
	GetNuevaVariable() interface{}

	// SetNuevaVariable sets the nuevaVariable attribute.
	SetNuevaVariable(interface{})

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	nuevaVariable interface{}
	_TK_ID        antlr.Token
	_expresion    IExpresionContext
	_tipo         ITipoContext
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *VariableContext) Set_TK_ID(v antlr.Token) { s._TK_ID = v }

func (s *VariableContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *VariableContext) Get_tipo() ITipoContext { return s._tipo }

func (s *VariableContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *VariableContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *VariableContext) GetNuevaVariable() interface{} { return s.nuevaVariable }

func (s *VariableContext) SetNuevaVariable(v interface{}) { s.nuevaVariable = v }

func (s *VariableContext) TK_LET() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LET, 0)
}

func (s *VariableContext) TK_MUT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MUT, 0)
}

func (s *VariableContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *VariableContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IGUAL, 0)
}

func (s *VariableContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *VariableContext) TK_PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, 0)
}

func (s *VariableContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DOSPUNTOS, 0)
}

func (s *VariableContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *rustParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, rustParserRULE_variable)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(32)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(33)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(34)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(35)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(36)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).nuevaVariable = symbol.NewSimbolo((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VariableContext).Get_expresion() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*VariableContext).Get_expresion().GetStart(), localctx.(*VariableContext)._expresion.GetStop())
			}
		}()), true, tipoValor)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(40)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(41)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(42)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(43)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).nuevaVariable = symbol.NewSimbolo((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VariableContext).Get_expresion() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*VariableContext).Get_expresion().GetStart(), localctx.(*VariableContext)._expresion.GetStop())
			}
		}()), false, tipoValor)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(47)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(48)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(49)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(50)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(51)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(52)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(53)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).nuevaVariable = symbol.NewSimbolo((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VariableContext).Get_expresion() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*VariableContext).Get_expresion().GetStart(), localctx.(*VariableContext)._expresion.GetStop())
			}
		}()), true, localctx.(*VariableContext).Get_tipo().GetTipoExp())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(57)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(58)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(59)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(60)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(61)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(62)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).nuevaVariable = symbol.NewSimbolo((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VariableContext).Get_expresion() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*VariableContext).Get_expresion().GetStart(), localctx.(*VariableContext)._expresion.GetStop())
			}
		}()), false, localctx.(*VariableContext).Get_tipo().GetTipoExp())

	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipoExp returns the tipoExp attribute.
	GetTipoExp() interfaces.TipoExpression

	// SetTipoExp sets the tipoExp attribute.
	SetTipoExp(interfaces.TipoExpression)

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	tipoExp interfaces.TipoExpression
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) GetTipoExp() interfaces.TipoExpression { return s.tipoExp }

func (s *TipoContext) SetTipoExp(v interfaces.TipoExpression) { s.tipoExp = v }

func (s *TipoContext) TK_INT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_INT, 0)
}

func (s *TipoContext) TK_FLOAT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_FLOAT, 0)
}

func (s *TipoContext) TK_BOOL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_BOOL, 0)
}

func (s *TipoContext) TK_CHAR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CHAR, 0)
}

func (s *TipoContext) TK_STRING() antlr.TerminalNode {
	return s.GetToken(rustParserTK_STRING, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *rustParser) Tipo() (localctx ITipoContext) {
	this := p
	_ = this

	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, rustParserRULE_tipo)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(rustParserTK_STRING)
		}
		localctx.(*TipoContext).tipoExp = interfaces.STRING

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Valor() IValorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) TK_POR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_POR, 0)
}

func (s *ExpresionContext) TK_DIVISION() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DIVISION, 0)
}

func (s *ExpresionContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MAS, 0)
}

func (s *ExpresionContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MENOS, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *rustParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *rustParser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, rustParserRULE_expresion, _p)

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
	{
		p.SetState(80)
		p.Valor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(83)
					p.Match(rustParserTK_POR)
				}
				{
					p.SetState(84)
					p.expresion(6)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(86)
					p.Match(rustParserTK_DIVISION)
				}
				{
					p.SetState(87)
					p.expresion(5)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(89)
					p.Match(rustParserTK_MAS)
				}
				{
					p.SetState(90)
					p.expresion(4)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(92)
					p.Match(rustParserTK_MENOS)
				}
				{
					p.SetState(93)
					p.expresion(3)
				}

			}

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValorContext() *ValorContext {
	var p = new(ValorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_valor
	return p
}

func (*ValorContext) IsValorContext() {}

func NewValorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValorContext {
	var p = new(ValorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_valor

	return p
}

func (s *ValorContext) GetParser() antlr.Parser { return s.parser }

func (s *ValorContext) TK_NUMBER() antlr.TerminalNode {
	return s.GetToken(rustParserTK_NUMBER, 0)
}

func (s *ValorContext) TK_DECIMAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DECIMAL, 0)
}

func (s *ValorContext) TK_TRUE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_TRUE, 0)
}

func (s *ValorContext) TK_FALSE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_FALSE, 0)
}

func (s *ValorContext) TK_CADENA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CADENA, 0)
}

func (s *ValorContext) TK_CARACTER() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CARACTER, 0)
}

func (s *ValorContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *ValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterValor(s)
	}
}

func (s *ValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitValor(s)
	}
}

func (p *rustParser) Valor() (localctx IValorContext) {
	this := p
	_ = this

	localctx = NewValorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, rustParserRULE_valor)

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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(rustParserTK_NUMBER)
		}
		tipoValor = interfaces.INTEGER

	case rustParserTK_DECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(rustParserTK_DECIMAL)
		}
		tipoValor = interfaces.FLOAT

	case rustParserTK_TRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(rustParserTK_TRUE)
		}
		tipoValor = interfaces.BOOLEAN

	case rustParserTK_FALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Match(rustParserTK_FALSE)
		}
		tipoValor = interfaces.BOOLEAN

	case rustParserTK_CADENA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(rustParserTK_CADENA)
		}
		tipoValor = interfaces.STRING

	case rustParserTK_CARACTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(109)
			p.Match(rustParserTK_CARACTER)
		}
		tipoValor = interfaces.CHAR

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.Match(rustParserTK_ID)
		}
		tipoValor = interfaces.IDENTIFICADOR

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ImpresionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ImpresionContext) TK_PARENTESIS_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_RIGHT, 0)
}

func (s *ImpresionContext) TK_PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, 0)
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
	p.EnterRule(localctx, 12, rustParserRULE_impresion)

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
		p.SetState(115)
		p.Match(rustParserTK_IMPRESION)
	}
	{
		p.SetState(116)
		p.Match(rustParserTK_PARENTESIS_LEFT)
	}
	{
		p.SetState(117)
		p.expresion(0)
	}
	{
		p.SetState(118)
		p.Match(rustParserTK_PARENTESIS_RIGHT)
	}
	{
		p.SetState(119)
		p.Match(rustParserTK_PUNTO_COMA)
	}

	return localctx
}

func (p *rustParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *rustParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
