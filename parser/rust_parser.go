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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 69, 233,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 43, 10, 3, 12, 3, 14, 3, 46, 11, 3,
	3, 3, 3, 3, 5, 3, 50, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 86, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 132, 10, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	144, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5,
	12, 154, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 212, 10, 12, 12, 12, 14,
	12, 215, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 231, 10, 13, 3, 13, 2,
	3, 22, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2, 19,
	21, 3, 2, 17, 18, 2, 253, 2, 26, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 63,
	3, 2, 2, 2, 8, 65, 3, 2, 2, 2, 10, 72, 3, 2, 2, 2, 12, 85, 3, 2, 2, 2,
	14, 87, 3, 2, 2, 2, 16, 93, 3, 2, 2, 2, 18, 131, 3, 2, 2, 2, 20, 143, 3,
	2, 2, 2, 22, 153, 3, 2, 2, 2, 24, 230, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2,
	27, 28, 7, 2, 2, 3, 28, 29, 8, 2, 1, 2, 29, 3, 3, 2, 2, 2, 30, 31, 7, 47,
	2, 2, 31, 32, 7, 39, 2, 2, 32, 33, 7, 4, 2, 2, 33, 34, 7, 5, 2, 2, 34,
	35, 7, 8, 2, 2, 35, 50, 7, 9, 2, 2, 36, 37, 7, 47, 2, 2, 37, 38, 7, 39,
	2, 2, 38, 39, 7, 4, 2, 2, 39, 40, 7, 5, 2, 2, 40, 44, 7, 8, 2, 2, 41, 43,
	5, 6, 4, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2,
	44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 7,
	9, 2, 2, 48, 50, 8, 3, 1, 2, 49, 30, 3, 2, 2, 2, 49, 36, 3, 2, 2, 2, 50,
	5, 3, 2, 2, 2, 51, 52, 5, 18, 10, 2, 52, 53, 8, 4, 1, 2, 53, 64, 3, 2,
	2, 2, 54, 55, 5, 8, 5, 2, 55, 56, 8, 4, 1, 2, 56, 64, 3, 2, 2, 2, 57, 58,
	5, 10, 6, 2, 58, 59, 8, 4, 1, 2, 59, 64, 3, 2, 2, 2, 60, 61, 5, 12, 7,
	2, 61, 62, 8, 4, 1, 2, 62, 64, 3, 2, 2, 2, 63, 51, 3, 2, 2, 2, 63, 54,
	3, 2, 2, 2, 63, 57, 3, 2, 2, 2, 63, 60, 3, 2, 2, 2, 64, 7, 3, 2, 2, 2,
	65, 66, 7, 33, 2, 2, 66, 67, 7, 4, 2, 2, 67, 68, 5, 22, 12, 2, 68, 69,
	7, 5, 2, 2, 69, 70, 7, 10, 2, 2, 70, 71, 8, 5, 1, 2, 71, 9, 3, 2, 2, 2,
	72, 73, 7, 67, 2, 2, 73, 74, 7, 16, 2, 2, 74, 75, 5, 22, 12, 2, 75, 76,
	7, 10, 2, 2, 76, 77, 8, 6, 1, 2, 77, 11, 3, 2, 2, 2, 78, 79, 5, 14, 8,
	2, 79, 80, 8, 7, 1, 2, 80, 86, 3, 2, 2, 2, 81, 82, 5, 14, 8, 2, 82, 83,
	5, 16, 9, 2, 83, 84, 8, 7, 1, 2, 84, 86, 3, 2, 2, 2, 85, 78, 3, 2, 2, 2,
	85, 81, 3, 2, 2, 2, 86, 13, 3, 2, 2, 2, 87, 88, 7, 61, 2, 2, 88, 89, 5,
	22, 12, 2, 89, 90, 7, 8, 2, 2, 90, 91, 7, 9, 2, 2, 91, 92, 8, 8, 1, 2,
	92, 15, 3, 2, 2, 2, 93, 94, 7, 62, 2, 2, 94, 95, 7, 8, 2, 2, 95, 96, 7,
	9, 2, 2, 96, 17, 3, 2, 2, 2, 97, 98, 7, 41, 2, 2, 98, 99, 7, 42, 2, 2,
	99, 100, 7, 67, 2, 2, 100, 101, 7, 16, 2, 2, 101, 102, 5, 22, 12, 2, 102,
	103, 7, 10, 2, 2, 103, 104, 8, 10, 1, 2, 104, 132, 3, 2, 2, 2, 105, 106,
	7, 41, 2, 2, 106, 107, 7, 67, 2, 2, 107, 108, 7, 16, 2, 2, 108, 109, 5,
	22, 12, 2, 109, 110, 7, 10, 2, 2, 110, 111, 8, 10, 1, 2, 111, 132, 3, 2,
	2, 2, 112, 113, 7, 41, 2, 2, 113, 114, 7, 42, 2, 2, 114, 115, 7, 67, 2,
	2, 115, 116, 7, 11, 2, 2, 116, 117, 5, 20, 11, 2, 117, 118, 7, 16, 2, 2,
	118, 119, 5, 22, 12, 2, 119, 120, 7, 10, 2, 2, 120, 121, 8, 10, 1, 2, 121,
	132, 3, 2, 2, 2, 122, 123, 7, 41, 2, 2, 123, 124, 7, 67, 2, 2, 124, 125,
	7, 11, 2, 2, 125, 126, 5, 20, 11, 2, 126, 127, 7, 16, 2, 2, 127, 128, 5,
	22, 12, 2, 128, 129, 7, 10, 2, 2, 129, 130, 8, 10, 1, 2, 130, 132, 3, 2,
	2, 2, 131, 97, 3, 2, 2, 2, 131, 105, 3, 2, 2, 2, 131, 112, 3, 2, 2, 2,
	131, 122, 3, 2, 2, 2, 132, 19, 3, 2, 2, 2, 133, 134, 7, 34, 2, 2, 134,
	144, 8, 11, 1, 2, 135, 136, 7, 35, 2, 2, 136, 144, 8, 11, 1, 2, 137, 138,
	7, 36, 2, 2, 138, 144, 8, 11, 1, 2, 139, 140, 7, 37, 2, 2, 140, 144, 8,
	11, 1, 2, 141, 142, 7, 38, 2, 2, 142, 144, 8, 11, 1, 2, 143, 133, 3, 2,
	2, 2, 143, 135, 3, 2, 2, 2, 143, 137, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2,
	143, 141, 3, 2, 2, 2, 144, 21, 3, 2, 2, 2, 145, 146, 8, 12, 1, 2, 146,
	147, 7, 4, 2, 2, 147, 148, 5, 22, 12, 2, 148, 149, 7, 5, 2, 2, 149, 154,
	3, 2, 2, 2, 150, 151, 5, 24, 13, 2, 151, 152, 8, 12, 1, 2, 152, 154, 3,
	2, 2, 2, 153, 145, 3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 154, 213, 3, 2, 2,
	2, 155, 156, 12, 17, 2, 2, 156, 157, 9, 2, 2, 2, 157, 158, 5, 22, 12, 18,
	158, 159, 8, 12, 1, 2, 159, 212, 3, 2, 2, 2, 160, 161, 12, 16, 2, 2, 161,
	162, 9, 3, 2, 2, 162, 163, 5, 22, 12, 17, 163, 164, 8, 12, 1, 2, 164, 212,
	3, 2, 2, 2, 165, 166, 12, 15, 2, 2, 166, 167, 7, 23, 2, 2, 167, 212, 5,
	22, 12, 16, 168, 169, 12, 14, 2, 2, 169, 170, 7, 3, 2, 2, 170, 212, 5,
	22, 12, 15, 171, 172, 12, 12, 2, 2, 172, 173, 7, 27, 2, 2, 173, 174, 5,
	22, 12, 13, 174, 175, 8, 12, 1, 2, 175, 212, 3, 2, 2, 2, 176, 177, 12,
	11, 2, 2, 177, 178, 7, 28, 2, 2, 178, 179, 5, 22, 12, 12, 179, 180, 8,
	12, 1, 2, 180, 212, 3, 2, 2, 2, 181, 182, 12, 10, 2, 2, 182, 183, 7, 14,
	2, 2, 183, 184, 5, 22, 12, 11, 184, 185, 8, 12, 1, 2, 185, 212, 3, 2, 2,
	2, 186, 187, 12, 9, 2, 2, 187, 188, 7, 13, 2, 2, 188, 189, 5, 22, 12, 10,
	189, 190, 8, 12, 1, 2, 190, 212, 3, 2, 2, 2, 191, 192, 12, 8, 2, 2, 192,
	193, 7, 25, 2, 2, 193, 194, 5, 22, 12, 9, 194, 195, 8, 12, 1, 2, 195, 212,
	3, 2, 2, 2, 196, 197, 12, 7, 2, 2, 197, 198, 7, 26, 2, 2, 198, 199, 5,
	22, 12, 8, 199, 200, 8, 12, 1, 2, 200, 212, 3, 2, 2, 2, 201, 202, 12, 6,
	2, 2, 202, 203, 7, 30, 2, 2, 203, 204, 5, 22, 12, 7, 204, 205, 8, 12, 1,
	2, 205, 212, 3, 2, 2, 2, 206, 207, 12, 5, 2, 2, 207, 208, 7, 29, 2, 2,
	208, 212, 5, 22, 12, 6, 209, 210, 12, 13, 2, 2, 210, 212, 7, 22, 2, 2,
	211, 155, 3, 2, 2, 2, 211, 160, 3, 2, 2, 2, 211, 165, 3, 2, 2, 2, 211,
	168, 3, 2, 2, 2, 211, 171, 3, 2, 2, 2, 211, 176, 3, 2, 2, 2, 211, 181,
	3, 2, 2, 2, 211, 186, 3, 2, 2, 2, 211, 191, 3, 2, 2, 2, 211, 196, 3, 2,
	2, 2, 211, 201, 3, 2, 2, 2, 211, 206, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2,
	212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214,
	23, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 217, 7, 63, 2, 2, 217, 231,
	8, 13, 1, 2, 218, 219, 7, 64, 2, 2, 219, 231, 8, 13, 1, 2, 220, 221, 7,
	45, 2, 2, 221, 231, 8, 13, 1, 2, 222, 223, 7, 46, 2, 2, 223, 231, 8, 13,
	1, 2, 224, 225, 7, 65, 2, 2, 225, 231, 8, 13, 1, 2, 226, 227, 7, 66, 2,
	2, 227, 231, 8, 13, 1, 2, 228, 229, 7, 67, 2, 2, 229, 231, 8, 13, 1, 2,
	230, 216, 3, 2, 2, 2, 230, 218, 3, 2, 2, 2, 230, 220, 3, 2, 2, 2, 230,
	222, 3, 2, 2, 2, 230, 224, 3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 230, 228,
	3, 2, 2, 2, 231, 25, 3, 2, 2, 2, 12, 44, 49, 63, 85, 131, 143, 153, 211,
	213, 230,
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
	"start", "funcionmain", "instrucciones", "impresion", "asignacionVariable",
	"expresionIf", "sintaxisIf", "sintaxisElse", "variable", "tipo", "expresion",
	"valor",
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
	rustParserRULE_start              = 0
	rustParserRULE_funcionmain        = 1
	rustParserRULE_instrucciones      = 2
	rustParserRULE_impresion          = 3
	rustParserRULE_asignacionVariable = 4
	rustParserRULE_expresionIf        = 5
	rustParserRULE_sintaxisIf         = 6
	rustParserRULE_sintaxisElse       = 7
	rustParserRULE_variable           = 8
	rustParserRULE_tipo               = 9
	rustParserRULE_expresion          = 10
	rustParserRULE_valor              = 11
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
		p.SetState(24)

		var _x = p.Funcionmain()

		localctx.(*StartContext)._funcionmain = _x
	}
	{
		p.SetState(25)
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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(29)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(30)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(31)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(32)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(33)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(35)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(36)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(37)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(38)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(rustParserTK_IMPRESION-31))|(1<<(rustParserTK_LET-31))|(1<<(rustParserTK_IF-31)))) != 0) || _la == rustParserTK_ID {
			{
				p.SetState(39)

				var _x = p.Instrucciones()

				localctx.(*FuncionmainContext)._instrucciones = _x
			}
			localctx.(*FuncionmainContext).e = append(localctx.(*FuncionmainContext).e, localctx.(*FuncionmainContext)._instrucciones)

			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(45)
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

	// Get_asignacionVariable returns the _asignacionVariable rule contexts.
	Get_asignacionVariable() IAsignacionVariableContext

	// Set_variable sets the _variable rule contexts.
	Set_variable(IVariableContext)

	// Set_impresion sets the _impresion rule contexts.
	Set_impresion(IImpresionContext)

	// Set_asignacionVariable sets the _asignacionVariable rule contexts.
	Set_asignacionVariable(IAsignacionVariableContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	inst                interfaces.Instruction
	_variable           IVariableContext
	_impresion          IImpresionContext
	_asignacionVariable IAsignacionVariableContext
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

func (s *InstruccionesContext) Get_asignacionVariable() IAsignacionVariableContext {
	return s._asignacionVariable
}

func (s *InstruccionesContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionesContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionesContext) Set_asignacionVariable(v IAsignacionVariableContext) {
	s._asignacionVariable = v
}

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

func (s *InstruccionesContext) AsignacionVariable() IAsignacionVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionVariableContext)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _x = p.Variable()

			localctx.(*InstruccionesContext)._variable = _x
		}
		localctx.(*InstruccionesContext).inst = localctx.(*InstruccionesContext).Get_variable().GetInst()

	case rustParserTK_IMPRESION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)

			var _x = p.Impresion()

			localctx.(*InstruccionesContext)._impresion = _x
		}
		localctx.(*InstruccionesContext).inst = localctx.(*InstruccionesContext).Get_impresion().GetInst()

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)

			var _x = p.AsignacionVariable()

			localctx.(*InstruccionesContext)._asignacionVariable = _x
		}
		localctx.(*InstruccionesContext).inst = localctx.(*InstruccionesContext).Get_asignacionVariable().GetInst()

	case rustParserTK_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
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
		p.SetState(63)
		p.Match(rustParserTK_IMPRESION)
	}
	{
		p.SetState(64)
		p.Match(rustParserTK_PARENTESIS_LEFT)
	}
	{
		p.SetState(65)

		var _x = p.expresion(0)

		localctx.(*ImpresionContext)._expresion = _x
	}
	{
		p.SetState(66)
		p.Match(rustParserTK_PARENTESIS_RIGHT)
	}
	{
		p.SetState(67)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*ImpresionContext).inst = instrucciones.NewImprimir(localctx.(*ImpresionContext).Get_expresion().GetPrimate())

	return localctx
}

// IAsignacionVariableContext is an interface to support dynamic dispatch.
type IAsignacionVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_ID sets the _TK_ID token.
	Set_TK_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsAsignacionVariableContext differentiates from other interfaces.
	IsAsignacionVariableContext()
}

type AsignacionVariableContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_TK_ID     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyAsignacionVariableContext() *AsignacionVariableContext {
	var p = new(AsignacionVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_asignacionVariable
	return p
}

func (*AsignacionVariableContext) IsAsignacionVariableContext() {}

func NewAsignacionVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionVariableContext {
	var p = new(AsignacionVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_asignacionVariable

	return p
}

func (s *AsignacionVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionVariableContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *AsignacionVariableContext) Set_TK_ID(v antlr.Token) { s._TK_ID = v }

func (s *AsignacionVariableContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AsignacionVariableContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionVariableContext) GetInst() interfaces.Instruction { return s.inst }

func (s *AsignacionVariableContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *AsignacionVariableContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *AsignacionVariableContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IGUAL, 0)
}

func (s *AsignacionVariableContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionVariableContext) TK_PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, 0)
}

func (s *AsignacionVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterAsignacionVariable(s)
	}
}

func (s *AsignacionVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitAsignacionVariable(s)
	}
}

func (p *rustParser) AsignacionVariable() (localctx IAsignacionVariableContext) {
	this := p
	_ = this

	localctx = NewAsignacionVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, rustParserRULE_asignacionVariable)

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
		p.SetState(70)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*AsignacionVariableContext)._TK_ID = _m
	}
	{
		p.SetState(71)
		p.Match(rustParserTK_IGUAL)
	}
	{
		p.SetState(72)

		var _x = p.expresion(0)

		localctx.(*AsignacionVariableContext)._expresion = _x
	}
	{
		p.SetState(73)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*AsignacionVariableContext).inst = instrucciones.NewAsignacion((func() string {
		if localctx.(*AsignacionVariableContext).Get_TK_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionVariableContext).Get_TK_ID().GetText()
		}
	}()), localctx.(*AsignacionVariableContext).Get_expresion().GetPrimate())

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
	p.EnterRule(localctx, 10, rustParserRULE_expresionIf)

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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.SintaxisIf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.SintaxisIf()
		}
		{
			p.SetState(80)
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
	p.EnterRule(localctx, 12, rustParserRULE_sintaxisIf)

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
		p.SetState(85)
		p.Match(rustParserTK_IF)
	}
	{
		p.SetState(86)
		p.expresion(0)
	}
	{
		p.SetState(87)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(88)
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
	p.EnterRule(localctx, 14, rustParserRULE_sintaxisElse)

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
		p.SetState(91)
		p.Match(rustParserTK_ELSE)
	}
	{
		p.SetState(92)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(93)
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
	p.EnterRule(localctx, 16, rustParserRULE_variable)

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(96)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(97)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(98)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(99)
			p.expresion(0)
		}
		{
			p.SetState(100)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(104)
			p.Match(rustParserTK_ID)
		}
		{
			p.SetState(105)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(106)
			p.expresion(0)
		}
		{
			p.SetState(107)
			p.Match(rustParserTK_PUNTO_COMA)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(111)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(112)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(113)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(114)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(115)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(116)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(117)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), false, localctx.(*VariableContext).Get_tipo().GetTipoExp(), localctx.(*VariableContext).Get_expresion().GetPrimate())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(121)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(122)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(123)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(124)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(125)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(126)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), true, localctx.(*VariableContext).Get_tipo().GetTipoExp(), localctx.(*VariableContext).Get_expresion().GetPrimate())

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
	p.EnterRule(localctx, 18, rustParserRULE_tipo)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(137)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, rustParserRULE_expresion, _p)
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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_PARENTESIS_LEFT:
		{
			p.SetState(144)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(145)
			p.expresion(0)
		}
		{
			p.SetState(146)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}

	case rustParserTK_TRUE, rustParserTK_FALSE, rustParserTK_NUMBER, rustParserTK_DECIMAL, rustParserTK_CADENA, rustParserTK_CARACTER, rustParserTK_ID:
		{
			p.SetState(148)

			var _x = p.Valor()

			localctx.(*ExpresionContext)._valor = _x
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(154)

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
					p.SetState(155)

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
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(159)

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
					p.SetState(160)

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
				p.SetState(163)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(164)
					p.Match(rustParserTK_AMPERSAND)
				}
				{
					p.SetState(165)
					p.expresion(14)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(166)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(167)
					p.Match(rustParserT__0)
				}
				{
					p.SetState(168)
					p.expresion(13)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(170)

					var _m = p.Match(rustParserTK_IGUALIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(171)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(175)

					var _m = p.Match(rustParserTK_DIFERENTE)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(176)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(180)

					var _m = p.Match(rustParserTK_MAYOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(181)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(185)

					var _m = p.Match(rustParserTK_MENOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(186)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(190)

					var _m = p.Match(rustParserTK_MAYORIGULA)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(191)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(195)

					var _m = p.Match(rustParserTK_MENOIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(196)

					var _x = p.expresion(6)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 11:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(200)

					var _m = p.Match(rustParserTK_AND)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(201)

					var _x = p.expresion(5)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 12:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(205)
					p.Match(rustParserTK_OR)
				}
				{
					p.SetState(206)
					p.expresion(4)
				}

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(208)
					p.Match(rustParserTK_BARRA)
				}

			}

		}
		p.SetState(213)
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

	// Get_TK_TRUE returns the _TK_TRUE token.
	Get_TK_TRUE() antlr.Token

	// Get_TK_FALSE returns the _TK_FALSE token.
	Get_TK_FALSE() antlr.Token

	// Get_TK_CADENA returns the _TK_CADENA token.
	Get_TK_CADENA() antlr.Token

	// Get_TK_CARACTER returns the _TK_CARACTER token.
	Get_TK_CARACTER() antlr.Token

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_NUMBER sets the _TK_NUMBER token.
	Set_TK_NUMBER(antlr.Token)

	// Set_TK_DECIMAL sets the _TK_DECIMAL token.
	Set_TK_DECIMAL(antlr.Token)

	// Set_TK_TRUE sets the _TK_TRUE token.
	Set_TK_TRUE(antlr.Token)

	// Set_TK_FALSE sets the _TK_FALSE token.
	Set_TK_FALSE(antlr.Token)

	// Set_TK_CADENA sets the _TK_CADENA token.
	Set_TK_CADENA(antlr.Token)

	// Set_TK_CARACTER sets the _TK_CARACTER token.
	Set_TK_CARACTER(antlr.Token)

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
	parser       antlr.Parser
	primate      interfaces.Expresion
	_TK_NUMBER   antlr.Token
	_TK_DECIMAL  antlr.Token
	_TK_TRUE     antlr.Token
	_TK_FALSE    antlr.Token
	_TK_CADENA   antlr.Token
	_TK_CARACTER antlr.Token
	_TK_ID       antlr.Token
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

func (s *ValorContext) Get_TK_TRUE() antlr.Token { return s._TK_TRUE }

func (s *ValorContext) Get_TK_FALSE() antlr.Token { return s._TK_FALSE }

func (s *ValorContext) Get_TK_CADENA() antlr.Token { return s._TK_CADENA }

func (s *ValorContext) Get_TK_CARACTER() antlr.Token { return s._TK_CARACTER }

func (s *ValorContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *ValorContext) Set_TK_NUMBER(v antlr.Token) { s._TK_NUMBER = v }

func (s *ValorContext) Set_TK_DECIMAL(v antlr.Token) { s._TK_DECIMAL = v }

func (s *ValorContext) Set_TK_TRUE(v antlr.Token) { s._TK_TRUE = v }

func (s *ValorContext) Set_TK_FALSE(v antlr.Token) { s._TK_FALSE = v }

func (s *ValorContext) Set_TK_CADENA(v antlr.Token) { s._TK_CADENA = v }

func (s *ValorContext) Set_TK_CARACTER(v antlr.Token) { s._TK_CARACTER = v }

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
	p.EnterRule(localctx, 22, rustParserRULE_valor)

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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)

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
			p.SetState(216)

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
			p.SetState(218)

			var _m = p.Match(rustParserTK_TRUE)

			localctx.(*ValorContext)._TK_TRUE = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito((func() string {
			if localctx.(*ValorContext).Get_TK_TRUE() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_TRUE().GetText()
			}
		}()), interfaces.BOOLEAN)

	case rustParserTK_FALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)

			var _m = p.Match(rustParserTK_FALSE)

			localctx.(*ValorContext)._TK_FALSE = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito((func() string {
			if localctx.(*ValorContext).Get_TK_FALSE() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_FALSE().GetText()
			}
		}()), interfaces.BOOLEAN)

	case rustParserTK_CADENA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)

			var _m = p.Match(rustParserTK_CADENA)

			localctx.(*ValorContext)._TK_CADENA = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextString((func() string {
			if localctx.(*ValorContext).Get_TK_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_CADENA().GetText()
			}
		}())), interfaces.STRING)

	case rustParserTK_CARACTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)

			var _m = p.Match(rustParserTK_CARACTER)

			localctx.(*ValorContext)._TK_CARACTER = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextString((func() string {
			if localctx.(*ValorContext).Get_TK_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_CARACTER().GetText()
			}
		}())), interfaces.CHAR)

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(226)

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
	case 10:
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
