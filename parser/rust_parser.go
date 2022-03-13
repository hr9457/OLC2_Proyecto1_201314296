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

// import "reflect"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 71, 293,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 50, 10, 3, 3, 4, 7, 4, 53, 10, 4, 12, 4, 14, 4, 56, 11, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 75, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	128, 10, 8, 3, 9, 6, 9, 131, 10, 9, 13, 9, 14, 9, 132, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 185, 10, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 197, 10, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 212, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 272, 10, 14, 12, 14, 14, 14, 275, 11, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 291, 10, 15, 3, 15, 2, 3, 26, 16, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 2, 5, 4, 2, 18, 18, 24, 24, 3, 2, 19, 21, 3, 2,
	17, 18, 2, 316, 2, 30, 3, 2, 2, 2, 4, 49, 3, 2, 2, 2, 6, 54, 3, 2, 2, 2,
	8, 74, 3, 2, 2, 2, 10, 76, 3, 2, 2, 2, 12, 83, 3, 2, 2, 2, 14, 127, 3,
	2, 2, 2, 16, 130, 3, 2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 143, 3, 2, 2, 2,
	22, 184, 3, 2, 2, 2, 24, 196, 3, 2, 2, 2, 26, 211, 3, 2, 2, 2, 28, 290,
	3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3, 32, 33, 8, 2, 1, 2,
	33, 3, 3, 2, 2, 2, 34, 35, 7, 47, 2, 2, 35, 36, 7, 39, 2, 2, 36, 37, 7,
	4, 2, 2, 37, 38, 7, 5, 2, 2, 38, 39, 7, 8, 2, 2, 39, 50, 7, 9, 2, 2, 40,
	41, 7, 47, 2, 2, 41, 42, 7, 39, 2, 2, 42, 43, 7, 4, 2, 2, 43, 44, 7, 5,
	2, 2, 44, 45, 7, 8, 2, 2, 45, 46, 5, 6, 4, 2, 46, 47, 7, 9, 2, 2, 47, 48,
	8, 3, 1, 2, 48, 50, 3, 2, 2, 2, 49, 34, 3, 2, 2, 2, 49, 40, 3, 2, 2, 2,
	50, 5, 3, 2, 2, 2, 51, 53, 5, 8, 5, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2,
	2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54,
	3, 2, 2, 2, 57, 58, 8, 4, 1, 2, 58, 7, 3, 2, 2, 2, 59, 60, 5, 22, 12, 2,
	60, 61, 8, 5, 1, 2, 61, 75, 3, 2, 2, 2, 62, 63, 5, 10, 6, 2, 63, 64, 8,
	5, 1, 2, 64, 75, 3, 2, 2, 2, 65, 66, 5, 12, 7, 2, 66, 67, 8, 5, 1, 2, 67,
	75, 3, 2, 2, 2, 68, 69, 5, 14, 8, 2, 69, 70, 8, 5, 1, 2, 70, 75, 3, 2,
	2, 2, 71, 72, 5, 20, 11, 2, 72, 73, 8, 5, 1, 2, 73, 75, 3, 2, 2, 2, 74,
	59, 3, 2, 2, 2, 74, 62, 3, 2, 2, 2, 74, 65, 3, 2, 2, 2, 74, 68, 3, 2, 2,
	2, 74, 71, 3, 2, 2, 2, 75, 9, 3, 2, 2, 2, 76, 77, 7, 33, 2, 2, 77, 78,
	7, 4, 2, 2, 78, 79, 5, 26, 14, 2, 79, 80, 7, 5, 2, 2, 80, 81, 7, 10, 2,
	2, 81, 82, 8, 6, 1, 2, 82, 11, 3, 2, 2, 2, 83, 84, 7, 69, 2, 2, 84, 85,
	7, 16, 2, 2, 85, 86, 5, 26, 14, 2, 86, 87, 7, 10, 2, 2, 87, 88, 8, 7, 1,
	2, 88, 13, 3, 2, 2, 2, 89, 90, 7, 61, 2, 2, 90, 91, 5, 26, 14, 2, 91, 92,
	7, 8, 2, 2, 92, 93, 5, 6, 4, 2, 93, 94, 7, 9, 2, 2, 94, 95, 8, 8, 1, 2,
	95, 128, 3, 2, 2, 2, 96, 97, 7, 61, 2, 2, 97, 98, 5, 26, 14, 2, 98, 99,
	7, 8, 2, 2, 99, 100, 5, 6, 4, 2, 100, 101, 7, 9, 2, 2, 101, 102, 7, 62,
	2, 2, 102, 103, 7, 8, 2, 2, 103, 104, 5, 6, 4, 2, 104, 105, 7, 9, 2, 2,
	105, 106, 8, 8, 1, 2, 106, 128, 3, 2, 2, 2, 107, 108, 7, 61, 2, 2, 108,
	109, 5, 26, 14, 2, 109, 110, 7, 8, 2, 2, 110, 111, 5, 6, 4, 2, 111, 112,
	7, 9, 2, 2, 112, 113, 5, 16, 9, 2, 113, 114, 8, 8, 1, 2, 114, 128, 3, 2,
	2, 2, 115, 116, 7, 61, 2, 2, 116, 117, 5, 26, 14, 2, 117, 118, 7, 8, 2,
	2, 118, 119, 5, 6, 4, 2, 119, 120, 7, 9, 2, 2, 120, 121, 5, 16, 9, 2, 121,
	122, 7, 62, 2, 2, 122, 123, 7, 8, 2, 2, 123, 124, 5, 6, 4, 2, 124, 125,
	7, 9, 2, 2, 125, 126, 8, 8, 1, 2, 126, 128, 3, 2, 2, 2, 127, 89, 3, 2,
	2, 2, 127, 96, 3, 2, 2, 2, 127, 107, 3, 2, 2, 2, 127, 115, 3, 2, 2, 2,
	128, 15, 3, 2, 2, 2, 129, 131, 5, 18, 10, 2, 130, 129, 3, 2, 2, 2, 131,
	132, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 135, 8, 9, 1, 2, 135, 17, 3, 2, 2, 2, 136, 137, 7, 63,
	2, 2, 137, 138, 5, 26, 14, 2, 138, 139, 7, 8, 2, 2, 139, 140, 5, 6, 4,
	2, 140, 141, 7, 9, 2, 2, 141, 142, 8, 10, 1, 2, 142, 19, 3, 2, 2, 2, 143,
	144, 7, 64, 2, 2, 144, 145, 5, 26, 14, 2, 145, 146, 7, 8, 2, 2, 146, 147,
	5, 6, 4, 2, 147, 148, 7, 9, 2, 2, 148, 149, 8, 11, 1, 2, 149, 21, 3, 2,
	2, 2, 150, 151, 7, 41, 2, 2, 151, 152, 7, 42, 2, 2, 152, 153, 7, 69, 2,
	2, 153, 154, 7, 16, 2, 2, 154, 155, 5, 26, 14, 2, 155, 156, 7, 10, 2, 2,
	156, 157, 8, 12, 1, 2, 157, 185, 3, 2, 2, 2, 158, 159, 7, 41, 2, 2, 159,
	160, 7, 69, 2, 2, 160, 161, 7, 16, 2, 2, 161, 162, 5, 26, 14, 2, 162, 163,
	7, 10, 2, 2, 163, 164, 8, 12, 1, 2, 164, 185, 3, 2, 2, 2, 165, 166, 7,
	41, 2, 2, 166, 167, 7, 42, 2, 2, 167, 168, 7, 69, 2, 2, 168, 169, 7, 11,
	2, 2, 169, 170, 5, 24, 13, 2, 170, 171, 7, 16, 2, 2, 171, 172, 5, 26, 14,
	2, 172, 173, 7, 10, 2, 2, 173, 174, 8, 12, 1, 2, 174, 185, 3, 2, 2, 2,
	175, 176, 7, 41, 2, 2, 176, 177, 7, 69, 2, 2, 177, 178, 7, 11, 2, 2, 178,
	179, 5, 24, 13, 2, 179, 180, 7, 16, 2, 2, 180, 181, 5, 26, 14, 2, 181,
	182, 7, 10, 2, 2, 182, 183, 8, 12, 1, 2, 183, 185, 3, 2, 2, 2, 184, 150,
	3, 2, 2, 2, 184, 158, 3, 2, 2, 2, 184, 165, 3, 2, 2, 2, 184, 175, 3, 2,
	2, 2, 185, 23, 3, 2, 2, 2, 186, 187, 7, 34, 2, 2, 187, 197, 8, 13, 1, 2,
	188, 189, 7, 35, 2, 2, 189, 197, 8, 13, 1, 2, 190, 191, 7, 36, 2, 2, 191,
	197, 8, 13, 1, 2, 192, 193, 7, 37, 2, 2, 193, 197, 8, 13, 1, 2, 194, 195,
	7, 38, 2, 2, 195, 197, 8, 13, 1, 2, 196, 186, 3, 2, 2, 2, 196, 188, 3,
	2, 2, 2, 196, 190, 3, 2, 2, 2, 196, 192, 3, 2, 2, 2, 196, 194, 3, 2, 2,
	2, 197, 25, 3, 2, 2, 2, 198, 199, 8, 14, 1, 2, 199, 200, 9, 2, 2, 2, 200,
	201, 5, 26, 14, 18, 201, 202, 8, 14, 1, 2, 202, 212, 3, 2, 2, 2, 203, 204,
	7, 4, 2, 2, 204, 205, 5, 26, 14, 2, 205, 206, 7, 5, 2, 2, 206, 207, 8,
	14, 1, 2, 207, 212, 3, 2, 2, 2, 208, 209, 5, 28, 15, 2, 209, 210, 8, 14,
	1, 2, 210, 212, 3, 2, 2, 2, 211, 198, 3, 2, 2, 2, 211, 203, 3, 2, 2, 2,
	211, 208, 3, 2, 2, 2, 212, 273, 3, 2, 2, 2, 213, 214, 12, 17, 2, 2, 214,
	215, 9, 3, 2, 2, 215, 216, 5, 26, 14, 18, 216, 217, 8, 14, 1, 2, 217, 272,
	3, 2, 2, 2, 218, 219, 12, 16, 2, 2, 219, 220, 9, 4, 2, 2, 220, 221, 5,
	26, 14, 17, 221, 222, 8, 14, 1, 2, 222, 272, 3, 2, 2, 2, 223, 224, 12,
	15, 2, 2, 224, 225, 7, 23, 2, 2, 225, 272, 5, 26, 14, 16, 226, 227, 12,
	14, 2, 2, 227, 228, 7, 3, 2, 2, 228, 272, 5, 26, 14, 15, 229, 230, 12,
	12, 2, 2, 230, 231, 7, 27, 2, 2, 231, 232, 5, 26, 14, 13, 232, 233, 8,
	14, 1, 2, 233, 272, 3, 2, 2, 2, 234, 235, 12, 11, 2, 2, 235, 236, 7, 28,
	2, 2, 236, 237, 5, 26, 14, 12, 237, 238, 8, 14, 1, 2, 238, 272, 3, 2, 2,
	2, 239, 240, 12, 10, 2, 2, 240, 241, 7, 14, 2, 2, 241, 242, 5, 26, 14,
	11, 242, 243, 8, 14, 1, 2, 243, 272, 3, 2, 2, 2, 244, 245, 12, 9, 2, 2,
	245, 246, 7, 13, 2, 2, 246, 247, 5, 26, 14, 10, 247, 248, 8, 14, 1, 2,
	248, 272, 3, 2, 2, 2, 249, 250, 12, 8, 2, 2, 250, 251, 7, 25, 2, 2, 251,
	252, 5, 26, 14, 9, 252, 253, 8, 14, 1, 2, 253, 272, 3, 2, 2, 2, 254, 255,
	12, 7, 2, 2, 255, 256, 7, 26, 2, 2, 256, 257, 5, 26, 14, 8, 257, 258, 8,
	14, 1, 2, 258, 272, 3, 2, 2, 2, 259, 260, 12, 6, 2, 2, 260, 261, 7, 30,
	2, 2, 261, 262, 5, 26, 14, 7, 262, 263, 8, 14, 1, 2, 263, 272, 3, 2, 2,
	2, 264, 265, 12, 5, 2, 2, 265, 266, 7, 29, 2, 2, 266, 267, 5, 26, 14, 6,
	267, 268, 8, 14, 1, 2, 268, 272, 3, 2, 2, 2, 269, 270, 12, 13, 2, 2, 270,
	272, 7, 22, 2, 2, 271, 213, 3, 2, 2, 2, 271, 218, 3, 2, 2, 2, 271, 223,
	3, 2, 2, 2, 271, 226, 3, 2, 2, 2, 271, 229, 3, 2, 2, 2, 271, 234, 3, 2,
	2, 2, 271, 239, 3, 2, 2, 2, 271, 244, 3, 2, 2, 2, 271, 249, 3, 2, 2, 2,
	271, 254, 3, 2, 2, 2, 271, 259, 3, 2, 2, 2, 271, 264, 3, 2, 2, 2, 271,
	269, 3, 2, 2, 2, 272, 275, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 273, 274,
	3, 2, 2, 2, 274, 27, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 276, 277, 7, 65,
	2, 2, 277, 291, 8, 15, 1, 2, 278, 279, 7, 66, 2, 2, 279, 291, 8, 15, 1,
	2, 280, 281, 7, 45, 2, 2, 281, 291, 8, 15, 1, 2, 282, 283, 7, 46, 2, 2,
	283, 291, 8, 15, 1, 2, 284, 285, 7, 67, 2, 2, 285, 291, 8, 15, 1, 2, 286,
	287, 7, 68, 2, 2, 287, 291, 8, 15, 1, 2, 288, 289, 7, 69, 2, 2, 289, 291,
	8, 15, 1, 2, 290, 276, 3, 2, 2, 2, 290, 278, 3, 2, 2, 2, 290, 280, 3, 2,
	2, 2, 290, 282, 3, 2, 2, 2, 290, 284, 3, 2, 2, 2, 290, 286, 3, 2, 2, 2,
	290, 288, 3, 2, 2, 2, 291, 29, 3, 2, 2, 2, 13, 49, 54, 74, 127, 132, 184,
	196, 211, 271, 273, 290,
}
var literalNames = []string{
	"", "'^'", "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", "','",
	"'<'", "'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'",
	"'!'", "'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'?'", "'->'", "'println!'",
	"'i64'", "'f64'", "'bool'", "'char'", "'String'", "'main'", "'usize'",
	"'let'", "'mut'", "'struct'", "'as'", "'true'", "'false'", "'fn'", "'return'",
	"'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", "'len'", "'push'",
	"'remove'", "'contains'", "'insert'", "'capacity'", "'with_capacity'",
	"'if'", "'else'", "'else if'", "'while'",
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
	"TK_WITHCAPACITY", "TK_IF", "TK_ELSE", "TK_ELSE_IF", "TK_WHILE", "TK_NUMBER",
	"TK_DECIMAL", "TK_CADENA", "TK_CARACTER", "TK_ID", "TK_COMMET", "SPACES",
}

var ruleNames = []string{
	"start", "funcionmain", "instrucciones", "instruccion", "impresion", "asignacionVariable",
	"expresionIf", "listaelif", "elif", "expresionWhile", "variable", "tipo",
	"expresion", "valor",
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
	rustParserTK_ELSE_IF          = 61
	rustParserTK_WHILE            = 62
	rustParserTK_NUMBER           = 63
	rustParserTK_DECIMAL          = 64
	rustParserTK_CADENA           = 65
	rustParserTK_CARACTER         = 66
	rustParserTK_ID               = 67
	rustParserTK_COMMET           = 68
	rustParserSPACES              = 69
)

// rustParser rules.
const (
	rustParserRULE_start              = 0
	rustParserRULE_funcionmain        = 1
	rustParserRULE_instrucciones      = 2
	rustParserRULE_instruccion        = 3
	rustParserRULE_impresion          = 4
	rustParserRULE_asignacionVariable = 5
	rustParserRULE_expresionIf        = 6
	rustParserRULE_listaelif          = 7
	rustParserRULE_elif               = 8
	rustParserRULE_expresionWhile     = 9
	rustParserRULE_variable           = 10
	rustParserRULE_tipo               = 11
	rustParserRULE_expresion          = 12
	rustParserRULE_valor              = 13
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
		p.SetState(28)

		var _x = p.Funcionmain()

		localctx.(*StartContext)._funcionmain = _x
	}
	{
		p.SetState(29)
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

func (s *FuncionmainContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
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
		{
			p.SetState(37)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(39)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(40)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(41)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(42)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(43)

			var _x = p.Instrucciones()

			localctx.(*FuncionmainContext)._instrucciones = _x
		}
		{
			p.SetState(44)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*FuncionmainContext).lista = localctx.(*FuncionmainContext).Get_instrucciones().GetLista()

	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
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

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetLista() *arrayList.List { return s.lista }

func (s *InstruccionesContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
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
	localctx.(*InstruccionesContext).lista = arrayList.New()
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(rustParserTK_IMPRESION-31))|(1<<(rustParserTK_LET-31))|(1<<(rustParserTK_IF-31))|(1<<(rustParserTK_WHILE-31)))) != 0) || _la == rustParserTK_ID {
		{
			p.SetState(49)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).lista.Add(e.GetInst())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_variable returns the _variable rule contexts.
	Get_variable() IVariableContext

	// Get_impresion returns the _impresion rule contexts.
	Get_impresion() IImpresionContext

	// Get_asignacionVariable returns the _asignacionVariable rule contexts.
	Get_asignacionVariable() IAsignacionVariableContext

	// Get_expresionIf returns the _expresionIf rule contexts.
	Get_expresionIf() IExpresionIfContext

	// Get_expresionWhile returns the _expresionWhile rule contexts.
	Get_expresionWhile() IExpresionWhileContext

	// Set_variable sets the _variable rule contexts.
	Set_variable(IVariableContext)

	// Set_impresion sets the _impresion rule contexts.
	Set_impresion(IImpresionContext)

	// Set_asignacionVariable sets the _asignacionVariable rule contexts.
	Set_asignacionVariable(IAsignacionVariableContext)

	// Set_expresionIf sets the _expresionIf rule contexts.
	Set_expresionIf(IExpresionIfContext)

	// Set_expresionWhile sets the _expresionWhile rule contexts.
	Set_expresionWhile(IExpresionWhileContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	inst                interfaces.Instruction
	_variable           IVariableContext
	_impresion          IImpresionContext
	_asignacionVariable IAsignacionVariableContext
	_expresionIf        IExpresionIfContext
	_expresionWhile     IExpresionWhileContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_variable() IVariableContext { return s._variable }

func (s *InstruccionContext) Get_impresion() IImpresionContext { return s._impresion }

func (s *InstruccionContext) Get_asignacionVariable() IAsignacionVariableContext {
	return s._asignacionVariable
}

func (s *InstruccionContext) Get_expresionIf() IExpresionIfContext { return s._expresionIf }

func (s *InstruccionContext) Get_expresionWhile() IExpresionWhileContext { return s._expresionWhile }

func (s *InstruccionContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_asignacionVariable(v IAsignacionVariableContext) {
	s._asignacionVariable = v
}

func (s *InstruccionContext) Set_expresionIf(v IExpresionIfContext) { s._expresionIf = v }

func (s *InstruccionContext) Set_expresionWhile(v IExpresionWhileContext) { s._expresionWhile = v }

func (s *InstruccionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstruccionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstruccionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *InstruccionContext) Impresion() IImpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpresionContext)
}

func (s *InstruccionContext) AsignacionVariable() IAsignacionVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionVariableContext)
}

func (s *InstruccionContext) ExpresionIf() IExpresionIfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionIfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionIfContext)
}

func (s *InstruccionContext) ExpresionWhile() IExpresionWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionWhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionWhileContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *rustParser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, rustParserRULE_instruccion)

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

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)

			var _x = p.Variable()

			localctx.(*InstruccionContext)._variable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_variable().GetInst()

	case rustParserTK_IMPRESION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)

			var _x = p.AsignacionVariable()

			localctx.(*InstruccionContext)._asignacionVariable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_asignacionVariable().GetInst()

	case rustParserTK_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(66)

			var _x = p.ExpresionIf()

			localctx.(*InstruccionContext)._expresionIf = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionIf().GetInst()

	case rustParserTK_WHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)

			var _x = p.ExpresionWhile()

			localctx.(*InstruccionContext)._expresionWhile = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionWhile().GetInst()

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
	p.EnterRule(localctx, 8, rustParserRULE_impresion)

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
		p.Match(rustParserTK_IMPRESION)
	}
	{
		p.SetState(75)
		p.Match(rustParserTK_PARENTESIS_LEFT)
	}
	{
		p.SetState(76)

		var _x = p.expresion(0)

		localctx.(*ImpresionContext)._expresion = _x
	}
	{
		p.SetState(77)
		p.Match(rustParserTK_PARENTESIS_RIGHT)
	}
	{
		p.SetState(78)
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
	p.EnterRule(localctx, 10, rustParserRULE_asignacionVariable)

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
		p.SetState(81)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*AsignacionVariableContext)._TK_ID = _m
	}
	{
		p.SetState(82)
		p.Match(rustParserTK_IGUAL)
	}
	{
		p.SetState(83)

		var _x = p.expresion(0)

		localctx.(*AsignacionVariableContext)._expresion = _x
	}
	{
		p.SetState(84)
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

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetBloqueif returns the bloqueif rule contexts.
	GetBloqueif() IInstruccionesContext

	// GetBloqueelse returns the bloqueelse rule contexts.
	GetBloqueelse() IInstruccionesContext

	// GetIfblock returns the ifblock rule contexts.
	GetIfblock() IInstruccionesContext

	// Get_listaelif returns the _listaelif rule contexts.
	Get_listaelif() IListaelifContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetBloqueif sets the bloqueif rule contexts.
	SetBloqueif(IInstruccionesContext)

	// SetBloqueelse sets the bloqueelse rule contexts.
	SetBloqueelse(IInstruccionesContext)

	// SetIfblock sets the ifblock rule contexts.
	SetIfblock(IInstruccionesContext)

	// Set_listaelif sets the _listaelif rule contexts.
	Set_listaelif(IListaelifContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsExpresionIfContext differentiates from other interfaces.
	IsExpresionIfContext()
}

type ExpresionIfContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_expresion IExpresionContext
	bloqueif   IInstruccionesContext
	bloqueelse IInstruccionesContext
	ifblock    IInstruccionesContext
	_listaelif IListaelifContext
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

func (s *ExpresionIfContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionIfContext) GetBloqueif() IInstruccionesContext { return s.bloqueif }

func (s *ExpresionIfContext) GetBloqueelse() IInstruccionesContext { return s.bloqueelse }

func (s *ExpresionIfContext) GetIfblock() IInstruccionesContext { return s.ifblock }

func (s *ExpresionIfContext) Get_listaelif() IListaelifContext { return s._listaelif }

func (s *ExpresionIfContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionIfContext) SetBloqueif(v IInstruccionesContext) { s.bloqueif = v }

func (s *ExpresionIfContext) SetBloqueelse(v IInstruccionesContext) { s.bloqueelse = v }

func (s *ExpresionIfContext) SetIfblock(v IInstruccionesContext) { s.ifblock = v }

func (s *ExpresionIfContext) Set_listaelif(v IListaelifContext) { s._listaelif = v }

func (s *ExpresionIfContext) GetInst() interfaces.Instruction { return s.inst }

func (s *ExpresionIfContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *ExpresionIfContext) TK_IF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IF, 0)
}

func (s *ExpresionIfContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionIfContext) AllTK_LLAVE_LEFT() []antlr.TerminalNode {
	return s.GetTokens(rustParserTK_LLAVE_LEFT)
}

func (s *ExpresionIfContext) TK_LLAVE_LEFT(i int) antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, i)
}

func (s *ExpresionIfContext) AllTK_LLAVE_RIGHT() []antlr.TerminalNode {
	return s.GetTokens(rustParserTK_LLAVE_RIGHT)
}

func (s *ExpresionIfContext) TK_LLAVE_RIGHT(i int) antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, i)
}

func (s *ExpresionIfContext) AllInstrucciones() []IInstruccionesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem())
	var tst = make([]IInstruccionesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionesContext)
		}
	}

	return tst
}

func (s *ExpresionIfContext) Instrucciones(i int) IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ExpresionIfContext) TK_ELSE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ELSE, 0)
}

func (s *ExpresionIfContext) Listaelif() IListaelifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaelifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaelifContext)
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
	p.EnterRule(localctx, 12, rustParserRULE_expresionIf)

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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(88)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(89)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(90)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(91)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(95)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(96)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(97)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(98)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(99)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(100)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(101)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(102)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), localctx.(*ExpresionIfContext).GetBloqueelse().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(106)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(107)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(108)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(109)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(110)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf2(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetIfblock().GetLista(), nil, localctx.(*ExpresionIfContext).Get_listaelif().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(113)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(114)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(115)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(116)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(117)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(118)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		{
			p.SetState(119)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(120)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(121)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(122)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf2(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetIfblock().GetLista(), localctx.(*ExpresionIfContext).GetBloqueelse().GetLista(), localctx.(*ExpresionIfContext).Get_listaelif().GetLista())

	}

	return localctx
}

// IListaelifContext is an interface to support dynamic dispatch.
type IListaelifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_elif returns the _elif rule contexts.
	Get_elif() IElifContext

	// Set_elif sets the _elif rule contexts.
	Set_elif(IElifContext)

	// GetL returns the l rule context list.
	GetL() []IElifContext

	// SetL sets the l rule context list.
	SetL([]IElifContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListaelifContext differentiates from other interfaces.
	IsListaelifContext()
}

type ListaelifContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayList.List
	_elif  IElifContext
	l      []IElifContext
}

func NewEmptyListaelifContext() *ListaelifContext {
	var p = new(ListaelifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_listaelif
	return p
}

func (*ListaelifContext) IsListaelifContext() {}

func NewListaelifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaelifContext {
	var p = new(ListaelifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_listaelif

	return p
}

func (s *ListaelifContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaelifContext) Get_elif() IElifContext { return s._elif }

func (s *ListaelifContext) Set_elif(v IElifContext) { s._elif = v }

func (s *ListaelifContext) GetL() []IElifContext { return s.l }

func (s *ListaelifContext) SetL(v []IElifContext) { s.l = v }

func (s *ListaelifContext) GetLista() *arrayList.List { return s.lista }

func (s *ListaelifContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListaelifContext) AllElif() []IElifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElifContext)(nil)).Elem())
	var tst = make([]IElifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElifContext)
		}
	}

	return tst
}

func (s *ListaelifContext) Elif(i int) IElifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElifContext)
}

func (s *ListaelifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaelifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaelifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterListaelif(s)
	}
}

func (s *ListaelifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitListaelif(s)
	}
}

func (p *rustParser) Listaelif() (localctx IListaelifContext) {
	this := p
	_ = this

	localctx = NewListaelifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, rustParserRULE_listaelif)
	localctx.(*ListaelifContext).lista = arrayList.New()
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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rustParserTK_ELSE_IF {
		{
			p.SetState(127)

			var _x = p.Elif()

			localctx.(*ListaelifContext)._elif = _x
		}
		localctx.(*ListaelifContext).l = append(localctx.(*ListaelifContext).l, localctx.(*ListaelifContext)._elif)

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listLocal := localctx.(*ListaelifContext).GetL()
	for _, l := range listLocal {
		localctx.(*ListaelifContext).lista.Add(l.GetInst())
	}

	return localctx
}

// IElifContext is an interface to support dynamic dispatch.
type IElifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetElifblock returns the elifblock rule contexts.
	GetElifblock() IInstruccionesContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetElifblock sets the elifblock rule contexts.
	SetElifblock(IInstruccionesContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsElifContext differentiates from other interfaces.
	IsElifContext()
}

type ElifContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_expresion IExpresionContext
	elifblock  IInstruccionesContext
}

func NewEmptyElifContext() *ElifContext {
	var p = new(ElifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_elif
	return p
}

func (*ElifContext) IsElifContext() {}

func NewElifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifContext {
	var p = new(ElifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_elif

	return p
}

func (s *ElifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ElifContext) GetElifblock() IInstruccionesContext { return s.elifblock }

func (s *ElifContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ElifContext) SetElifblock(v IInstruccionesContext) { s.elifblock = v }

func (s *ElifContext) GetInst() interfaces.Instruction { return s.inst }

func (s *ElifContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *ElifContext) TK_ELSE_IF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ELSE_IF, 0)
}

func (s *ElifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ElifContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *ElifContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *ElifContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ElifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterElif(s)
	}
}

func (s *ElifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitElif(s)
	}
}

func (p *rustParser) Elif() (localctx IElifContext) {
	this := p
	_ = this

	localctx = NewElifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, rustParserRULE_elif)

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
		p.SetState(134)
		p.Match(rustParserTK_ELSE_IF)
	}
	{
		p.SetState(135)

		var _x = p.expresion(0)

		localctx.(*ElifContext)._expresion = _x
	}
	{
		p.SetState(136)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(137)

		var _x = p.Instrucciones()

		localctx.(*ElifContext).elifblock = _x
	}
	{
		p.SetState(138)
		p.Match(rustParserTK_LLAVE_RIGHT)
	}
	localctx.(*ElifContext).inst = instrucciones.NewIf(localctx.(*ElifContext).Get_expresion().GetPrimate(), localctx.(*ElifContext).GetElifblock().GetLista(), nil)

	return localctx
}

// IExpresionWhileContext is an interface to support dynamic dispatch.
type IExpresionWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExpresionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpresionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsExpresionWhileContext differentiates from other interfaces.
	IsExpresionWhileContext()
}

type ExpresionWhileContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	inst           interfaces.Instruction
	exp            IExpresionContext
	_instrucciones IInstruccionesContext
}

func NewEmptyExpresionWhileContext() *ExpresionWhileContext {
	var p = new(ExpresionWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_expresionWhile
	return p
}

func (*ExpresionWhileContext) IsExpresionWhileContext() {}

func NewExpresionWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionWhileContext {
	var p = new(ExpresionWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_expresionWhile

	return p
}

func (s *ExpresionWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionWhileContext) GetExp() IExpresionContext { return s.exp }

func (s *ExpresionWhileContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *ExpresionWhileContext) SetExp(v IExpresionContext) { s.exp = v }

func (s *ExpresionWhileContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *ExpresionWhileContext) GetInst() interfaces.Instruction { return s.inst }

func (s *ExpresionWhileContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *ExpresionWhileContext) TK_WHILE() antlr.TerminalNode {
	return s.GetToken(rustParserTK_WHILE, 0)
}

func (s *ExpresionWhileContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *ExpresionWhileContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ExpresionWhileContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *ExpresionWhileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionWhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterExpresionWhile(s)
	}
}

func (s *ExpresionWhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitExpresionWhile(s)
	}
}

func (p *rustParser) ExpresionWhile() (localctx IExpresionWhileContext) {
	this := p
	_ = this

	localctx = NewExpresionWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, rustParserRULE_expresionWhile)

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
		p.SetState(141)
		p.Match(rustParserTK_WHILE)
	}
	{
		p.SetState(142)

		var _x = p.expresion(0)

		localctx.(*ExpresionWhileContext).exp = _x
	}
	{
		p.SetState(143)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(144)

		var _x = p.Instrucciones()

		localctx.(*ExpresionWhileContext)._instrucciones = _x
	}
	{
		p.SetState(145)
		p.Match(rustParserTK_LLAVE_RIGHT)
	}
	localctx.(*ExpresionWhileContext).inst = instrucciones.NewWhile(localctx.(*ExpresionWhileContext).GetExp().GetPrimate(), localctx.(*ExpresionWhileContext).Get_instrucciones().GetLista())

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
	_expresion IExpresionContext
	_tipo      ITipoContext
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
	p.EnterRule(localctx, 20, rustParserRULE_variable)

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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(149)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(150)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(151)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(152)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(153)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), true, interfaces.NULL, localctx.(*VariableContext).Get_expresion().GetPrimate())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(157)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(158)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(159)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(160)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), false, interfaces.NULL, localctx.(*VariableContext).Get_expresion().GetPrimate())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(164)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(165)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(166)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(167)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(168)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(169)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(170)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), true, localctx.(*VariableContext).Get_tipo().GetTipoExp(), localctx.(*VariableContext).Get_expresion().GetPrimate())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(174)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(175)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(176)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(177)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(178)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(179)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*VariableContext).inst = instrucciones.NewDeclaracion((func() string {
			if localctx.(*VariableContext).Get_TK_ID() == nil {
				return ""
			} else {
				return localctx.(*VariableContext).Get_TK_ID().GetText()
			}
		}()), false, localctx.(*VariableContext).Get_tipo().GetTipoExp(), localctx.(*VariableContext).Get_expresion().GetPrimate())

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
	p.EnterRule(localctx, 22, rustParserRULE_tipo)

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

	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(190)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(192)
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

	// GetRight returns the right rule contexts.
	GetRight() IExpresionContext

	// Get_valor returns the _valor rule contexts.
	Get_valor() IValorContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpresionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpresionContext)

	// Set_valor sets the _valor rule contexts.
	Set_valor(IValorContext)

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
	op      antlr.Token
	right   IExpresionContext
	_valor  IValorContext
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

func (s *ExpresionContext) GetRight() IExpresionContext { return s.right }

func (s *ExpresionContext) Get_valor() IValorContext { return s._valor }

func (s *ExpresionContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExpresionContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExpresionContext) Set_valor(v IValorContext) { s._valor = v }

func (s *ExpresionContext) GetPrimate() interfaces.Expresion { return s.primate }

func (s *ExpresionContext) SetPrimate(v interfaces.Expresion) { s.primate = v }

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

func (s *ExpresionContext) TK_ADMIRACION() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ADMIRACION, 0)
}

func (s *ExpresionContext) TK_MENOS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_MENOS, 0)
}

func (s *ExpresionContext) TK_PARENTESIS_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_LEFT, 0)
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
	_startState := 24
	p.EnterRecursionRule(localctx, 24, rustParserRULE_expresion, _p)
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
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_MENOS, rustParserTK_ADMIRACION:
		{
			p.SetState(197)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpresionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == rustParserTK_MENOS || _la == rustParserTK_ADMIRACION) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpresionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(198)

			var _x = p.expresion(16)

			localctx.(*ExpresionContext).right = _x
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(nil, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case rustParserTK_PARENTESIS_LEFT:
		{
			p.SetState(201)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(202)
			p.expresion(0)
		}
		{
			p.SetState(203)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	case rustParserTK_TRUE, rustParserTK_FALSE, rustParserTK_NUMBER, rustParserTK_DECIMAL, rustParserTK_CADENA, rustParserTK_CARACTER, rustParserTK_ID:
		{
			p.SetState(206)

			var _x = p.Valor()

			localctx.(*ExpresionContext)._valor = _x
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(212)

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
					p.SetState(213)

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
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(217)

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
					p.SetState(218)

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
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(222)
					p.Match(rustParserTK_AMPERSAND)
				}
				{
					p.SetState(223)
					p.expresion(14)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(225)
					p.Match(rustParserT__0)
				}
				{
					p.SetState(226)
					p.expresion(13)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(228)

					var _m = p.Match(rustParserTK_IGUALIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(229)

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
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(233)

					var _m = p.Match(rustParserTK_DIFERENTE)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(234)

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
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(238)

					var _m = p.Match(rustParserTK_MAYOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(239)

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
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(243)

					var _m = p.Match(rustParserTK_MENOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(244)

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
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(248)

					var _m = p.Match(rustParserTK_MAYORIGULA)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(249)

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
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(253)

					var _m = p.Match(rustParserTK_MENOIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(254)

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
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(258)

					var _m = p.Match(rustParserTK_AND)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(259)

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
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(263)

					var _m = p.Match(rustParserTK_OR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(264)

					var _x = p.expresion(4)

					localctx.(*ExpresionContext).right = _x
				}
				localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

			case 13:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(267)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(268)
					p.Match(rustParserTK_BARRA)
				}

			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, rustParserRULE_valor)

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

	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)

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
			p.SetState(276)

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
			p.SetState(278)

			var _m = p.Match(rustParserTK_TRUE)

			localctx.(*ValorContext)._TK_TRUE = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextBool((func() string {
			if localctx.(*ValorContext).Get_TK_TRUE() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_TRUE().GetText()
			}
		}())), interfaces.BOOLEAN)

	case rustParserTK_FALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)

			var _m = p.Match(rustParserTK_FALSE)

			localctx.(*ValorContext)._TK_FALSE = _m
		}
		localctx.(*ValorContext).primate = expressiones.NewPrimito(interfaces.ConvertTextBool((func() string {
			if localctx.(*ValorContext).Get_TK_FALSE() == nil {
				return ""
			} else {
				return localctx.(*ValorContext).Get_TK_FALSE().GetText()
			}
		}())), interfaces.BOOLEAN)

	case rustParserTK_CADENA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)

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
			p.SetState(284)

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
			p.SetState(286)

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
	case 12:
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
