// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Proyecto1/src/interfaces"
import "Proyecto1/src/expressiones"
import "Proyecto1/src/instrucciones"
import arrayList "github.com/colegno/arraylist"
import "Proyecto1/src/pruebas"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 208,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11, 3, 3, 3, 3, 3,
	5, 3, 48, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 75, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 121, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 133, 10, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 143, 10, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 187,
	10, 11, 12, 11, 14, 11, 190, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 206,
	10, 12, 3, 12, 2, 3, 20, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2,
	4, 3, 2, 19, 21, 3, 2, 17, 18, 2, 228, 2, 24, 3, 2, 2, 2, 4, 47, 3, 2,
	2, 2, 6, 58, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 74, 3, 2, 2, 2, 12, 76,
	3, 2, 2, 2, 14, 82, 3, 2, 2, 2, 16, 120, 3, 2, 2, 2, 18, 132, 3, 2, 2,
	2, 20, 142, 3, 2, 2, 2, 22, 205, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26,
	7, 2, 2, 3, 26, 27, 8, 2, 1, 2, 27, 3, 3, 2, 2, 2, 28, 29, 7, 47, 2, 2,
	29, 30, 7, 39, 2, 2, 30, 31, 7, 4, 2, 2, 31, 32, 7, 5, 2, 2, 32, 33, 7,
	8, 2, 2, 33, 48, 7, 9, 2, 2, 34, 35, 7, 47, 2, 2, 35, 36, 7, 39, 2, 2,
	36, 37, 7, 4, 2, 2, 37, 38, 7, 5, 2, 2, 38, 42, 7, 8, 2, 2, 39, 41, 5,
	6, 4, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42,
	43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7, 9, 2,
	2, 46, 48, 8, 3, 1, 2, 47, 28, 3, 2, 2, 2, 47, 34, 3, 2, 2, 2, 48, 5, 3,
	2, 2, 2, 49, 50, 5, 16, 9, 2, 50, 51, 8, 4, 1, 2, 51, 59, 3, 2, 2, 2, 52,
	53, 5, 8, 5, 2, 53, 54, 8, 4, 1, 2, 54, 59, 3, 2, 2, 2, 55, 56, 5, 10,
	6, 2, 56, 57, 8, 4, 1, 2, 57, 59, 3, 2, 2, 2, 58, 49, 3, 2, 2, 2, 58, 52,
	3, 2, 2, 2, 58, 55, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2, 60, 61, 7, 33, 2, 2,
	61, 62, 7, 4, 2, 2, 62, 63, 5, 20, 11, 2, 63, 64, 7, 5, 2, 2, 64, 65, 7,
	10, 2, 2, 65, 66, 8, 5, 1, 2, 66, 9, 3, 2, 2, 2, 67, 68, 5, 12, 7, 2, 68,
	69, 8, 6, 1, 2, 69, 75, 3, 2, 2, 2, 70, 71, 5, 12, 7, 2, 71, 72, 5, 14,
	8, 2, 72, 73, 8, 6, 1, 2, 73, 75, 3, 2, 2, 2, 74, 67, 3, 2, 2, 2, 74, 70,
	3, 2, 2, 2, 75, 11, 3, 2, 2, 2, 76, 77, 7, 61, 2, 2, 77, 78, 5, 20, 11,
	2, 78, 79, 7, 8, 2, 2, 79, 80, 7, 9, 2, 2, 80, 81, 8, 7, 1, 2, 81, 13,
	3, 2, 2, 2, 82, 83, 7, 62, 2, 2, 83, 84, 7, 8, 2, 2, 84, 85, 7, 9, 2, 2,
	85, 15, 3, 2, 2, 2, 86, 87, 7, 41, 2, 2, 87, 88, 7, 42, 2, 2, 88, 89, 7,
	67, 2, 2, 89, 90, 7, 16, 2, 2, 90, 91, 5, 20, 11, 2, 91, 92, 7, 10, 2,
	2, 92, 93, 8, 9, 1, 2, 93, 121, 3, 2, 2, 2, 94, 95, 7, 41, 2, 2, 95, 96,
	7, 67, 2, 2, 96, 97, 7, 16, 2, 2, 97, 98, 5, 20, 11, 2, 98, 99, 7, 10,
	2, 2, 99, 100, 8, 9, 1, 2, 100, 121, 3, 2, 2, 2, 101, 102, 7, 41, 2, 2,
	102, 103, 7, 42, 2, 2, 103, 104, 7, 67, 2, 2, 104, 105, 7, 11, 2, 2, 105,
	106, 5, 18, 10, 2, 106, 107, 7, 16, 2, 2, 107, 108, 5, 20, 11, 2, 108,
	109, 7, 10, 2, 2, 109, 110, 8, 9, 1, 2, 110, 121, 3, 2, 2, 2, 111, 112,
	7, 41, 2, 2, 112, 113, 7, 67, 2, 2, 113, 114, 7, 11, 2, 2, 114, 115, 5,
	18, 10, 2, 115, 116, 7, 16, 2, 2, 116, 117, 5, 20, 11, 2, 117, 118, 7,
	10, 2, 2, 118, 119, 8, 9, 1, 2, 119, 121, 3, 2, 2, 2, 120, 86, 3, 2, 2,
	2, 120, 94, 3, 2, 2, 2, 120, 101, 3, 2, 2, 2, 120, 111, 3, 2, 2, 2, 121,
	17, 3, 2, 2, 2, 122, 123, 7, 34, 2, 2, 123, 133, 8, 10, 1, 2, 124, 125,
	7, 35, 2, 2, 125, 133, 8, 10, 1, 2, 126, 127, 7, 36, 2, 2, 127, 133, 8,
	10, 1, 2, 128, 129, 7, 37, 2, 2, 129, 133, 8, 10, 1, 2, 130, 131, 7, 38,
	2, 2, 131, 133, 8, 10, 1, 2, 132, 122, 3, 2, 2, 2, 132, 124, 3, 2, 2, 2,
	132, 126, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133,
	19, 3, 2, 2, 2, 134, 135, 8, 11, 1, 2, 135, 136, 7, 4, 2, 2, 136, 137,
	5, 20, 11, 2, 137, 138, 7, 5, 2, 2, 138, 143, 3, 2, 2, 2, 139, 140, 5,
	22, 12, 2, 140, 141, 8, 11, 1, 2, 141, 143, 3, 2, 2, 2, 142, 134, 3, 2,
	2, 2, 142, 139, 3, 2, 2, 2, 143, 188, 3, 2, 2, 2, 144, 145, 12, 17, 2,
	2, 145, 146, 9, 2, 2, 2, 146, 147, 5, 20, 11, 18, 147, 148, 8, 11, 1, 2,
	148, 187, 3, 2, 2, 2, 149, 150, 12, 16, 2, 2, 150, 151, 9, 3, 2, 2, 151,
	152, 5, 20, 11, 17, 152, 153, 8, 11, 1, 2, 153, 187, 3, 2, 2, 2, 154, 155,
	12, 15, 2, 2, 155, 156, 7, 23, 2, 2, 156, 187, 5, 20, 11, 16, 157, 158,
	12, 14, 2, 2, 158, 159, 7, 3, 2, 2, 159, 187, 5, 20, 11, 15, 160, 161,
	12, 12, 2, 2, 161, 162, 7, 27, 2, 2, 162, 187, 5, 20, 11, 13, 163, 164,
	12, 11, 2, 2, 164, 165, 7, 28, 2, 2, 165, 187, 5, 20, 11, 12, 166, 167,
	12, 10, 2, 2, 167, 168, 7, 14, 2, 2, 168, 187, 5, 20, 11, 11, 169, 170,
	12, 9, 2, 2, 170, 171, 7, 13, 2, 2, 171, 187, 5, 20, 11, 10, 172, 173,
	12, 8, 2, 2, 173, 174, 7, 25, 2, 2, 174, 187, 5, 20, 11, 9, 175, 176, 12,
	7, 2, 2, 176, 177, 7, 26, 2, 2, 177, 187, 5, 20, 11, 8, 178, 179, 12, 6,
	2, 2, 179, 180, 7, 30, 2, 2, 180, 187, 5, 20, 11, 7, 181, 182, 12, 5, 2,
	2, 182, 183, 7, 29, 2, 2, 183, 187, 5, 20, 11, 6, 184, 185, 12, 13, 2,
	2, 185, 187, 7, 22, 2, 2, 186, 144, 3, 2, 2, 2, 186, 149, 3, 2, 2, 2, 186,
	154, 3, 2, 2, 2, 186, 157, 3, 2, 2, 2, 186, 160, 3, 2, 2, 2, 186, 163,
	3, 2, 2, 2, 186, 166, 3, 2, 2, 2, 186, 169, 3, 2, 2, 2, 186, 172, 3, 2,
	2, 2, 186, 175, 3, 2, 2, 2, 186, 178, 3, 2, 2, 2, 186, 181, 3, 2, 2, 2,
	186, 184, 3, 2, 2, 2, 187, 190, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188,
	189, 3, 2, 2, 2, 189, 21, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 191, 192, 7,
	63, 2, 2, 192, 206, 8, 12, 1, 2, 193, 194, 7, 64, 2, 2, 194, 206, 8, 12,
	1, 2, 195, 196, 7, 45, 2, 2, 196, 206, 8, 12, 1, 2, 197, 198, 7, 46, 2,
	2, 198, 206, 8, 12, 1, 2, 199, 200, 7, 65, 2, 2, 200, 206, 8, 12, 1, 2,
	201, 202, 7, 66, 2, 2, 202, 206, 8, 12, 1, 2, 203, 204, 7, 67, 2, 2, 204,
	206, 8, 12, 1, 2, 205, 191, 3, 2, 2, 2, 205, 193, 3, 2, 2, 2, 205, 195,
	3, 2, 2, 2, 205, 197, 3, 2, 2, 2, 205, 199, 3, 2, 2, 2, 205, 201, 3, 2,
	2, 2, 205, 203, 3, 2, 2, 2, 206, 23, 3, 2, 2, 2, 12, 42, 47, 58, 74, 120,
	132, 142, 186, 188, 205,
}
var literalNames = []string{
	"", "'^'", "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", "','",
	"'<'", "'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'",
	"'!'", "'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'?'", "'->'", "'println!'",
	"'i64'", "'f64'", "'bool'", "'char'", "'String'", "'main'", "'usize'",
	"'let'", "'mut'", "'struct'", "'as'", "'true'", "'false'", "'fn'", "'return'",
	"'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", "'len'", "'push'",
	"'remove'", "'contains'", "'insert'", "'capacity'", "'with_capacity'",
	"'if'", "'else'",
}
var symbolicNames = []string{
	"", "", "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF",
	"TK_CORCHETE_RIGHT", "TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA",
	"TK_DOSPUNTOS", "TK_COMA", "TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL",
	"TK_MAS", "TK_MENOS", "TK_POR", "TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA",
	"TK_AMPERSAND", "TK_ADMIRACION", "TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL",
	"TK_DIFERENTE", "TK_OR", "TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION",
	"TK_INT", "TK_FLOAT", "TK_BOOL", "TK_CHAR", "TK_STRING", "TK_MAIN", "TK_USIZE",
	"TK_LET", "TK_MUT", "TK_STRUCT", "TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN",
	"TK_RETURN", "TK_ABS", "TK_SQRT", "TK_TOSTRING", "TK_CLONE", "TK_NEW",
	"TK_LEN", "TK_PUSH", "TK_REMOVED", "TK_CONTAINS", "TK_INSERT", "TK_CAPACITY",
	"TK_WITHCAPACITY", "TK_IF", "TK_ELSE", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA",
	"TK_CARACTER", "TK_ID", "TK_COMMET", "SPACES",
}

var ruleNames = []string{
	"start", "funcionmain", "instrucciones", "impresion", "expresionIf", "sintaxisIf",
	"sintaxisElse", "variable", "tipo", "expresion", "valor",
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
	rustParserT__0                = 1
	rustParserTK_PARENTESIS_LEFT  = 2
	rustParserTK_PARENTESIS_RIGHT = 3
	rustParserTK_CORCHETE_LETF    = 4
	rustParserTK_CORCHETE_RIGHT   = 5
	rustParserTK_LLAVE_LEFT       = 6
	rustParserTK_LLAVE_RIGHT      = 7
	rustParserTK_PUNTO_COMA       = 8
	rustParserTK_DOSPUNTOS        = 9
	rustParserTK_COMA             = 10
	rustParserTK_MENOR            = 11
	rustParserTK_MAYOR            = 12
	rustParserTK_PUNTO            = 13
	rustParserTK_IGUAL            = 14
	rustParserTK_MAS              = 15
	rustParserTK_MENOS            = 16
	rustParserTK_POR              = 17
	rustParserTK_DIVISION         = 18
	rustParserTK_PORCENTAJE       = 19
	rustParserTK_BARRA            = 20
	rustParserTK_AMPERSAND        = 21
	rustParserTK_ADMIRACION       = 22
	rustParserTK_MAYORIGULA       = 23
	rustParserTK_MENOIGUAL        = 24
	rustParserTK_IGUALIGUAL       = 25
	rustParserTK_DIFERENTE        = 26
	rustParserTK_OR               = 27
	rustParserTK_AND              = 28
	rustParserTK_WHAT             = 29
	rustParserTK_TIPORETURN       = 30
	rustParserTK_IMPRESION        = 31
	rustParserTK_INT              = 32
	rustParserTK_FLOAT            = 33
	rustParserTK_BOOL             = 34
	rustParserTK_CHAR             = 35
	rustParserTK_STRING           = 36
	rustParserTK_MAIN             = 37
	rustParserTK_USIZE            = 38
	rustParserTK_LET              = 39
	rustParserTK_MUT              = 40
	rustParserTK_STRUCT           = 41
	rustParserTK_AS               = 42
	rustParserTK_TRUE             = 43
	rustParserTK_FALSE            = 44
	rustParserTK_FN               = 45
	rustParserTK_RETURN           = 46
	rustParserTK_ABS              = 47
	rustParserTK_SQRT             = 48
	rustParserTK_TOSTRING         = 49
	rustParserTK_CLONE            = 50
	rustParserTK_NEW              = 51
	rustParserTK_LEN              = 52
	rustParserTK_PUSH             = 53
	rustParserTK_REMOVED          = 54
	rustParserTK_CONTAINS         = 55
	rustParserTK_INSERT           = 56
	rustParserTK_CAPACITY         = 57
	rustParserTK_WITHCAPACITY     = 58
	rustParserTK_IF               = 59
	rustParserTK_ELSE             = 60
	rustParserTK_NUMBER           = 61
	rustParserTK_DECIMAL          = 62
	rustParserTK_CADENA           = 63
	rustParserTK_CARACTER         = 64
	rustParserTK_ID               = 65
	rustParserTK_COMMET           = 66
	rustParserSPACES              = 67
)

// rustParser rules.
const (
	rustParserRULE_start         = 0
	rustParserRULE_funcionmain   = 1
	rustParserRULE_instrucciones = 2
	rustParserRULE_impresion     = 3
	rustParserRULE_expresionIf   = 4
	rustParserRULE_sintaxisIf    = 5
	rustParserRULE_sintaxisElse  = 6
	rustParserRULE_variable      = 7
	rustParserRULE_tipo          = 8
	rustParserRULE_expresion     = 9
	rustParserRULE_valor         = 10
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcionmain returns the _funcionmain rule contexts.
	Get_funcionmain() IFuncionmainContext

	// Set_funcionmain sets the _funcionmain rule contexts.
	Set_funcionmain(IFuncionmainContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_funcionmain IFuncionmainContext
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

func (s *StartContext) Get_funcionmain() IFuncionmainContext { return s._funcionmain }

func (s *StartContext) Set_funcionmain(v IFuncionmainContext) { s._funcionmain = v }

func (s *StartContext) Funcionmain() IFuncionmainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionmainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionmainContext)
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
		p.SetState(22)

		var _x = p.Funcionmain()

		localctx.(*StartContext)._funcionmain = _x
	}
	{
		p.SetState(23)
		p.Match(rustParserEOF)
	}
	pruebas.Probar(localctx.(*StartContext).Get_funcionmain().GetLista())

	return localctx
}

// IFuncionmainContext is an interface to support dynamic dispatch.
type IFuncionmainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionesContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsFuncionmainContext differentiates from other interfaces.
	IsFuncionmainContext()
}

type FuncionmainContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
	e              []IInstruccionesContext
}

func NewEmptyFuncionmainContext() *FuncionmainContext {
	var p = new(FuncionmainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_funcionmain
	return p
}

func (*FuncionmainContext) IsFuncionmainContext() {}

func NewFuncionmainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionmainContext {
	var p = new(FuncionmainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_funcionmain

	return p
}

func (s *FuncionmainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionmainContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *FuncionmainContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *FuncionmainContext) GetE() []IInstruccionesContext { return s.e }

func (s *FuncionmainContext) SetE(v []IInstruccionesContext) { s.e = v }

func (s *FuncionmainContext) GetLista() *arrayList.List { return s.lista }

func (s *FuncionmainContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *FuncionmainContext) TK_FN() antlr.TerminalNode {
	return s.GetToken(rustParserTK_FN, 0)
}

func (s *FuncionmainContext) TK_MAIN() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MAIN, 0)
}

func (s *FuncionmainContext) TK_PARENTESIS_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_LEFT, 0)
}

func (s *FuncionmainContext) TK_PARENTESIS_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_RIGHT, 0)
}

func (s *FuncionmainContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *FuncionmainContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *FuncionmainContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *FuncionmainContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *FuncionmainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionmainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionmainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterFuncionmain(s)
	}
}

func (s *FuncionmainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitFuncionmain(s)
	}
}

func (p *rustParser) Funcionmain() (localctx IFuncionmainContext) {
	this := p
	_ = this

	localctx = NewFuncionmainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, rustParserRULE_funcionmain)
	localctx.(*FuncionmainContext).lista = arrayList.New()
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

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(27)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(28)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(29)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(30)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(31)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(33)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(34)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(35)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(36)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(rustParserTK_IMPRESION-31))|(1<<(rustParserTK_LET-31))|(1<<(rustParserTK_IF-31)))) != 0 {
			{
				p.SetState(37)

				var _x = p.Instrucciones()

				localctx.(*FuncionmainContext)._instrucciones = _x
			}
			localctx.(*FuncionmainContext).e = append(localctx.(*FuncionmainContext).e, localctx.(*FuncionmainContext)._instrucciones)

			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(43)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

		listInt := localctx.(*FuncionmainContext).GetE()
		for _, e := range listInt {
			localctx.(*FuncionmainContext).lista.Add(e.GetInst())
		}

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

	// Get_impresion returns the _impresion rule contexts.
	Get_impresion() IImpresionContext

	// Set_variable sets the _variable rule contexts.
	Set_variable(IVariableContext)

	// Set_impresion sets the _impresion rule contexts.
	Set_impresion(IImpresionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_variable  IVariableContext
	_impresion IImpresionContext
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

func (s *InstruccionesContext) Get_impresion() IImpresionContext { return s._impresion }

func (s *InstruccionesContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionesContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionesContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstruccionesContext) SetInst(v interfaces.Instruction) { s.inst = v }

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

func (s *InstruccionesContext) ExpresionIf() IExpresionIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionIfContext)
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
	p.EnterRule(localctx, 4, rustParserRULE_instrucciones)

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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _x = p.Variable()

			localctx.(*InstruccionesContext)._variable = _x
		}
		localctx.(*InstruccionesContext).inst = localctx.(*InstruccionesContext).Get_variable().GetInst()

	case rustParserTK_IMPRESION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _x = p.Impresion()

			localctx.(*InstruccionesContext)._impresion = _x
		}
		localctx.(*InstruccionesContext).inst = localctx.(*InstruccionesContext).Get_impresion().GetInst()

	case rustParserTK_IF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.ExpresionIf()
		}

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

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsImpresionContext differentiates from other interfaces.
	IsImpresionContext()
}

type ImpresionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_expresion IExpresionContext
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

func (s *ImpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ImpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ImpresionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *ImpresionContext) SetInst(v interfaces.Instruction) { s.inst = v }

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
	p.EnterRule(localctx, 6, rustParserRULE_impresion)

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
		p.SetState(58)
		p.Match(rustParserTK_IMPRESION)
	}
	{
		p.SetState(59)
		p.Match(rustParserTK_PARENTESIS_LEFT)
	}
	{
		p.SetState(60)

		var _x = p.expresion(0)

		localctx.(*ImpresionContext)._expresion = _x
	}
	{
		p.SetState(61)
		p.Match(rustParserTK_PARENTESIS_RIGHT)
	}
	{
		p.SetState(62)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*ImpresionContext).inst = instrucciones.NewImprimir(localctx.(*ImpresionContext).Get_expresion().GetPrimate())

	return localctx
}

// IExpresionIfContext is an interface to support dynamic dispatch.
type IExpresionIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresionIfContext differentiates from other interfaces.
	IsExpresionIfContext()
}

type ExpresionIfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionIfContext() *ExpresionIfContext {
	var p = new(ExpresionIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_expresionIf
	return p
}

func (*ExpresionIfContext) IsExpresionIfContext() {}

func NewExpresionIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionIfContext {
	var p = new(ExpresionIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_expresionIf

	return p
}

func (s *ExpresionIfContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionIfContext) SintaxisIf() ISintaxisIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISintaxisIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISintaxisIfContext)
}

func (s *ExpresionIfContext) SintaxisElse() ISintaxisElseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISintaxisElseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISintaxisElseContext)
}

func (s *ExpresionIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterExpresionIf(s)
	}
}

func (s *ExpresionIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitExpresionIf(s)
	}
}

func (p *rustParser) ExpresionIf() (localctx IExpresionIfContext) {
	this := p
	_ = this

	localctx = NewExpresionIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, rustParserRULE_expresionIf)

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

	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.SintaxisIf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.SintaxisIf()
		}
		{
			p.SetState(69)
			p.SintaxisElse()
		}

	}

	return localctx
}

// ISintaxisIfContext is an interface to support dynamic dispatch.
type ISintaxisIfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSintaxisIfContext differentiates from other interfaces.
	IsSintaxisIfContext()
}

type SintaxisIfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySintaxisIfContext() *SintaxisIfContext {
	var p = new(SintaxisIfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_sintaxisIf
	return p
}

func (*SintaxisIfContext) IsSintaxisIfContext() {}

func NewSintaxisIfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SintaxisIfContext {
	var p = new(SintaxisIfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_sintaxisIf

	return p
}

func (s *SintaxisIfContext) GetParser() antlr.Parser { return s.parser }

func (s *SintaxisIfContext) TK_IF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IF, 0)
}

func (s *SintaxisIfContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SintaxisIfContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *SintaxisIfContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *SintaxisIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SintaxisIfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SintaxisIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterSintaxisIf(s)
	}
}

func (s *SintaxisIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitSintaxisIf(s)
	}
}

func (p *rustParser) SintaxisIf() (localctx ISintaxisIfContext) {
	this := p
	_ = this

	localctx = NewSintaxisIfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, rustParserRULE_sintaxisIf)

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
		p.SetState(74)
		p.Match(rustParserTK_IF)
	}
	{
		p.SetState(75)
		p.expresion(0)
	}
	{
		p.SetState(76)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(77)
		p.Match(rustParserTK_LLAVE_RIGHT)
	}

	return localctx
}

// ISintaxisElseContext is an interface to support dynamic dispatch.
type ISintaxisElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSintaxisElseContext differentiates from other interfaces.
	IsSintaxisElseContext()
}

type SintaxisElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySintaxisElseContext() *SintaxisElseContext {
	var p = new(SintaxisElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_sintaxisElse
	return p
}

func (*SintaxisElseContext) IsSintaxisElseContext() {}

func NewSintaxisElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SintaxisElseContext {
	var p = new(SintaxisElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_sintaxisElse

	return p
}

func (s *SintaxisElseContext) GetParser() antlr.Parser { return s.parser }

func (s *SintaxisElseContext) TK_ELSE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ELSE, 0)
}

func (s *SintaxisElseContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *SintaxisElseContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *SintaxisElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SintaxisElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SintaxisElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterSintaxisElse(s)
	}
}

func (s *SintaxisElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitSintaxisElse(s)
	}
}

func (p *rustParser) SintaxisElse() (localctx ISintaxisElseContext) {
	this := p
	_ = this

	localctx = NewSintaxisElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, rustParserRULE_sintaxisElse)

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
		p.SetState(80)
		p.Match(rustParserTK_ELSE)
	}
	{
		p.SetState(81)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(82)
		p.Match(rustParserTK_LLAVE_RIGHT)
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

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_TK_ID     antlr.Token
	_tipo      ITipoContext
	_expresion IExpresionContext
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

func (s *VariableContext) Get_tipo() ITipoContext { return s._tipo }

func (s *VariableContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *VariableContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *VariableContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *VariableContext) GetInst() interfaces.Instruction { return s.inst }

func (s *VariableContext) SetInst(v interfaces.Instruction) { s.inst = v }

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
	p.EnterRule(localctx, 14, rustParserRULE_variable)

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

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(85)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(86)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(87)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(88)
			p.expresion(0)
		}
		{
			p.SetState(89)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(93)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(94)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(95)
			p.expresion(0)
		}
		{
			p.SetState(96)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(99)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(100)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(101)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(102)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(103)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(104)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(105)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(106)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), localctx.(*VariableContext).Get_tipo().GetTipoExp(), localctx.(*VariableContext).Get_expresion().GetPrimate())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(110)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(111)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(112)
			p.Tipo()
		}
		{
			p.SetState(113)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(114)
			p.expresion(0)
		}
		{
			p.SetState(115)
			p.Match(rustParserTK_PUNTO_COMA)
		}

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
	p.EnterRule(localctx, 16, rustParserRULE_tipo)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
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

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExpresionContext

	// Get_valor returns the _valor rule contexts.
	Get_valor() IValorContext

	// GetRight returns the right rule contexts.
	GetRight() IExpresionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpresionContext)

	// Set_valor sets the _valor rule contexts.
	Set_valor(IValorContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpresionContext)

	// GetPrimate returns the primate attribute.
	GetPrimate() interfaces.Expresion

	// SetPrimate sets the primate attribute.
	SetPrimate(interfaces.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	primate interfaces.Expresion
	left    IExpresionContext
	_valor  IValorContext
	op      antlr.Token
	right   IExpresionContext
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

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) GetLeft() IExpresionContext { return s.left }

func (s *ExpresionContext) Get_valor() IValorContext { return s._valor }

func (s *ExpresionContext) GetRight() IExpresionContext { return s.right }

func (s *ExpresionContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExpresionContext) Set_valor(v IValorContext) { s._valor = v }

func (s *ExpresionContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExpresionContext) GetPrimate() interfaces.Expresion { return s.primate }

func (s *ExpresionContext) SetPrimate(v interfaces.Expresion) { s.primate = v }

func (s *ExpresionContext) TK_PARENTESIS_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_LEFT, 0)
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

func (s *ExpresionContext) TK_PARENTESIS_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_RIGHT, 0)
}

func (s *ExpresionContext) Valor() IValorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *ExpresionContext) TK_POR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_POR, 0)
}

func (s *ExpresionContext) TK_DIVISION() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DIVISION, 0)
}

func (s *ExpresionContext) TK_PORCENTAJE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PORCENTAJE, 0)
}

func (s *ExpresionContext) TK_MAS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MAS, 0)
}

func (s *ExpresionContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MENOS, 0)
}

func (s *ExpresionContext) TK_AMPERSAND() antlr.TerminalNode {
	return s.GetToken(rustParserTK_AMPERSAND, 0)
}

func (s *ExpresionContext) TK_IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IGUALIGUAL, 0)
}

func (s *ExpresionContext) TK_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DIFERENTE, 0)
}

func (s *ExpresionContext) TK_MAYOR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MAYOR, 0)
}

func (s *ExpresionContext) TK_MENOR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MENOR, 0)
}

func (s *ExpresionContext) TK_MAYORIGULA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MAYORIGULA, 0)
}

func (s *ExpresionContext) TK_MENOIGUAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MENOIGUAL, 0)
}

func (s *ExpresionContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(rustParserTK_AND, 0)
}

func (s *ExpresionContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(rustParserTK_OR, 0)
}

func (s *ExpresionContext) TK_BARRA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_BARRA, 0)
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
	_startState := 18
	p.EnterRecursionRule(localctx, 18, rustParserRULE_expresion, _p)
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
	p.SetState(140)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_PARENTESIS_LEFT:
		{
			p.SetState(133)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(134)
			p.expresion(0)
		}
		{
			p.SetState(135)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}

	case rustParserTK_TRUE, rustParserTK_FALSE, rustParserTK_NUMBER, rustParserTK_DECIMAL, rustParserTK_CADENA, rustParserTK_CARACTER, rustParserTK_ID:
		{
			p.SetState(137)

			var _x = p.Valor()

			localctx.(*ExpresionContext)._valor = _x
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(143)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<rustParserTK_POR)|(1<<rustParserTK_DIVISION)|(1<<rustParserTK_PORCENTAJE))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(144)

					var _x = p.expresion(16)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(148)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == rustParserTK_MAS || _la == rustParserTK_MENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(149)

					var _x = p.expresion(15)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(153)
					p.Match(rustParserTK_AMPERSAND)
				}
				{
					p.SetState(154)
					p.expresion(14)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(156)
					p.Match(rustParserT__0)
				}
				{
					p.SetState(157)
					p.expresion(13)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(159)
					p.Match(rustParserTK_IGUALIGUAL)
				}
				{
					p.SetState(160)
					p.expresion(11)
				}

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(162)
					p.Match(rustParserTK_DIFERENTE)
				}
				{
					p.SetState(163)
					p.expresion(10)
				}

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(165)
					p.Match(rustParserTK_MAYOR)
				}
				{
					p.SetState(166)
					p.expresion(9)
				}

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(168)
					p.Match(rustParserTK_MENOR)
				}
				{
					p.SetState(169)
					p.expresion(8)
				}

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(171)
					p.Match(rustParserTK_MAYORIGULA)
				}
				{
					p.SetState(172)
					p.expresion(7)
				}

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(174)
					p.Match(rustParserTK_MENOIGUAL)
				}
				{
					p.SetState(175)
					p.expresion(6)
				}

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(177)
					p.Match(rustParserTK_AND)
				}
				{
					p.SetState(178)
					p.expresion(5)
				}

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(180)
					p.Match(rustParserTK_OR)
				}
				{
					p.SetState(181)
					p.expresion(4)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(183)
					p.Match(rustParserTK_BARRA)
				}

			}

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IValorContext is an interface to support dynamic dispatch.
type IValorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_NUMBER returns the _TK_NUMBER token.
	Get_TK_NUMBER() antlr.Token

	// Get_TK_DECIMAL returns the _TK_DECIMAL token.
	Get_TK_DECIMAL() antlr.Token

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_NUMBER sets the _TK_NUMBER token.
	Set_TK_NUMBER(antlr.Token)

	// Set_TK_DECIMAL sets the _TK_DECIMAL token.
	Set_TK_DECIMAL(antlr.Token)

	// Set_TK_ID sets the _TK_ID token.
	Set_TK_ID(antlr.Token)

	// GetPrimate returns the primate attribute.
	GetPrimate() interfaces.Expresion

	// SetPrimate sets the primate attribute.
	SetPrimate(interfaces.Expresion)

	// IsValorContext differentiates from other interfaces.
	IsValorContext()
}

type ValorContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	primate     interfaces.Expresion
	_TK_NUMBER  antlr.Token
	_TK_DECIMAL antlr.Token
	_TK_ID      antlr.Token
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

func (s *ValorContext) Get_TK_NUMBER() antlr.Token { return s._TK_NUMBER }

func (s *ValorContext) Get_TK_DECIMAL() antlr.Token { return s._TK_DECIMAL }

func (s *ValorContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *ValorContext) Set_TK_NUMBER(v antlr.Token) { s._TK_NUMBER = v }

func (s *ValorContext) Set_TK_DECIMAL(v antlr.Token) { s._TK_DECIMAL = v }

func (s *ValorContext) Set_TK_ID(v antlr.Token) { s._TK_ID = v }

func (s *ValorContext) GetPrimate() interfaces.Expresion { return s.primate }

func (s *ValorContext) SetPrimate(v interfaces.Expresion) { s.primate = v }

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
	p.EnterRule(localctx, 20, rustParserRULE_valor)

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

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)

			var _m = p.Match(rustParserTK_NUMBER)

			localctx.(*ValorContext)._TK_NUMBER = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextInt((func() string {
			if localctx.(*ValorContext).Get_TK_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_NUMBER().GetText()
			}
		}())), interfaces.INTEGER)

	case rustParserTK_DECIMAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)

			var _m = p.Match(rustParserTK_DECIMAL)

			localctx.(*ValorContext)._TK_DECIMAL = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextFloat((func() string {
			if localctx.(*ValorContext).Get_TK_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_DECIMAL().GetText()
			}
		}())), interfaces.FLOAT)

	case rustParserTK_TRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Match(rustParserTK_TRUE)
		}

	case rustParserTK_FALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(195)
			p.Match(rustParserTK_FALSE)
		}

	case rustParserTK_CADENA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(197)
			p.Match(rustParserTK_CADENA)
		}

	case rustParserTK_CARACTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(199)
			p.Match(rustParserTK_CARACTER)
		}

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(201)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*ValorContext)._TK_ID = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewLLamadoVariable((func() string {
			if localctx.(*ValorContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_ID().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *rustParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
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
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
