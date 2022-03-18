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

// import "Proyecto1/src/funciones"
import arrayList "github.com/colegno/arraylist"

// import "Proyecto1/src/pruebas"
// import "reflect"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 346,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 58, 10,
	3, 3, 4, 7, 4, 61, 10, 4, 12, 4, 14, 4, 64, 11, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 89, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 106, 10, 6, 3, 7, 6, 7, 109, 10, 7, 13, 7, 14, 7, 110, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 163, 10, 10, 3, 11, 6,
	11, 166, 10, 11, 13, 11, 14, 11, 167, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 230, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 242, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 273, 10, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 325, 10, 18, 12, 18, 14,
	18, 328, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 344, 10, 19, 3, 19, 2,
	3, 34, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 2, 5, 4, 2, 17, 17, 23, 23, 3, 2, 18, 20, 3, 2, 16, 17, 2, 368,
	2, 38, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 88, 3, 2, 2,
	2, 10, 105, 3, 2, 2, 2, 12, 108, 3, 2, 2, 2, 14, 114, 3, 2, 2, 2, 16, 118,
	3, 2, 2, 2, 18, 162, 3, 2, 2, 2, 20, 165, 3, 2, 2, 2, 22, 171, 3, 2, 2,
	2, 24, 178, 3, 2, 2, 2, 26, 185, 3, 2, 2, 2, 28, 191, 3, 2, 2, 2, 30, 229,
	3, 2, 2, 2, 32, 241, 3, 2, 2, 2, 34, 272, 3, 2, 2, 2, 36, 343, 3, 2, 2,
	2, 38, 39, 5, 4, 3, 2, 39, 40, 7, 2, 2, 3, 40, 41, 8, 2, 1, 2, 41, 3, 3,
	2, 2, 2, 42, 43, 7, 46, 2, 2, 43, 44, 7, 38, 2, 2, 44, 45, 7, 3, 2, 2,
	45, 46, 7, 4, 2, 2, 46, 47, 7, 7, 2, 2, 47, 58, 7, 8, 2, 2, 48, 49, 7,
	46, 2, 2, 49, 50, 7, 38, 2, 2, 50, 51, 7, 3, 2, 2, 51, 52, 7, 4, 2, 2,
	52, 53, 7, 7, 2, 2, 53, 54, 5, 6, 4, 2, 54, 55, 7, 8, 2, 2, 55, 56, 8,
	3, 1, 2, 56, 58, 3, 2, 2, 2, 57, 42, 3, 2, 2, 2, 57, 48, 3, 2, 2, 2, 58,
	5, 3, 2, 2, 2, 59, 61, 5, 8, 5, 2, 60, 59, 3, 2, 2, 2, 61, 64, 3, 2, 2,
	2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 65, 66, 8, 4, 1, 2, 66, 7, 3, 2, 2, 2, 67, 68, 5, 30, 16, 2,
	68, 69, 8, 5, 1, 2, 69, 89, 3, 2, 2, 2, 70, 71, 5, 10, 6, 2, 71, 72, 8,
	5, 1, 2, 72, 89, 3, 2, 2, 2, 73, 74, 5, 16, 9, 2, 74, 75, 8, 5, 1, 2, 75,
	89, 3, 2, 2, 2, 76, 77, 5, 18, 10, 2, 77, 78, 8, 5, 1, 2, 78, 89, 3, 2,
	2, 2, 79, 80, 5, 24, 13, 2, 80, 81, 8, 5, 1, 2, 81, 89, 3, 2, 2, 2, 82,
	83, 5, 26, 14, 2, 83, 84, 8, 5, 1, 2, 84, 89, 3, 2, 2, 2, 85, 86, 5, 28,
	15, 2, 86, 87, 8, 5, 1, 2, 87, 89, 3, 2, 2, 2, 88, 67, 3, 2, 2, 2, 88,
	70, 3, 2, 2, 2, 88, 73, 3, 2, 2, 2, 88, 76, 3, 2, 2, 2, 88, 79, 3, 2, 2,
	2, 88, 82, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 89, 9, 3, 2, 2, 2, 90, 91, 7,
	32, 2, 2, 91, 92, 7, 3, 2, 2, 92, 93, 5, 34, 18, 2, 93, 94, 7, 4, 2, 2,
	94, 95, 7, 9, 2, 2, 95, 96, 8, 6, 1, 2, 96, 106, 3, 2, 2, 2, 97, 98, 7,
	32, 2, 2, 98, 99, 7, 3, 2, 2, 99, 100, 5, 34, 18, 2, 100, 101, 5, 12, 7,
	2, 101, 102, 7, 4, 2, 2, 102, 103, 7, 9, 2, 2, 103, 104, 8, 6, 1, 2, 104,
	106, 3, 2, 2, 2, 105, 90, 3, 2, 2, 2, 105, 97, 3, 2, 2, 2, 106, 11, 3,
	2, 2, 2, 107, 109, 5, 14, 8, 2, 108, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2,
	2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	113, 8, 7, 1, 2, 113, 13, 3, 2, 2, 2, 114, 115, 7, 11, 2, 2, 115, 116,
	5, 34, 18, 2, 116, 117, 8, 8, 1, 2, 117, 15, 3, 2, 2, 2, 118, 119, 7, 72,
	2, 2, 119, 120, 7, 15, 2, 2, 120, 121, 5, 34, 18, 2, 121, 122, 7, 9, 2,
	2, 122, 123, 8, 9, 1, 2, 123, 17, 3, 2, 2, 2, 124, 125, 7, 60, 2, 2, 125,
	126, 5, 34, 18, 2, 126, 127, 7, 7, 2, 2, 127, 128, 5, 6, 4, 2, 128, 129,
	7, 8, 2, 2, 129, 130, 8, 10, 1, 2, 130, 163, 3, 2, 2, 2, 131, 132, 7, 60,
	2, 2, 132, 133, 5, 34, 18, 2, 133, 134, 7, 7, 2, 2, 134, 135, 5, 6, 4,
	2, 135, 136, 7, 8, 2, 2, 136, 137, 7, 61, 2, 2, 137, 138, 7, 7, 2, 2, 138,
	139, 5, 6, 4, 2, 139, 140, 7, 8, 2, 2, 140, 141, 8, 10, 1, 2, 141, 163,
	3, 2, 2, 2, 142, 143, 7, 60, 2, 2, 143, 144, 5, 34, 18, 2, 144, 145, 7,
	7, 2, 2, 145, 146, 5, 6, 4, 2, 146, 147, 7, 8, 2, 2, 147, 148, 5, 20, 11,
	2, 148, 149, 8, 10, 1, 2, 149, 163, 3, 2, 2, 2, 150, 151, 7, 60, 2, 2,
	151, 152, 5, 34, 18, 2, 152, 153, 7, 7, 2, 2, 153, 154, 5, 6, 4, 2, 154,
	155, 7, 8, 2, 2, 155, 156, 5, 20, 11, 2, 156, 157, 7, 61, 2, 2, 157, 158,
	7, 7, 2, 2, 158, 159, 5, 6, 4, 2, 159, 160, 7, 8, 2, 2, 160, 161, 8, 10,
	1, 2, 161, 163, 3, 2, 2, 2, 162, 124, 3, 2, 2, 2, 162, 131, 3, 2, 2, 2,
	162, 142, 3, 2, 2, 2, 162, 150, 3, 2, 2, 2, 163, 19, 3, 2, 2, 2, 164, 166,
	5, 22, 12, 2, 165, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 165, 3,
	2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 8, 11, 1,
	2, 170, 21, 3, 2, 2, 2, 171, 172, 7, 62, 2, 2, 172, 173, 5, 34, 18, 2,
	173, 174, 7, 7, 2, 2, 174, 175, 5, 6, 4, 2, 175, 176, 7, 8, 2, 2, 176,
	177, 8, 12, 1, 2, 177, 23, 3, 2, 2, 2, 178, 179, 7, 63, 2, 2, 179, 180,
	5, 34, 18, 2, 180, 181, 7, 7, 2, 2, 181, 182, 5, 6, 4, 2, 182, 183, 7,
	8, 2, 2, 183, 184, 8, 13, 1, 2, 184, 25, 3, 2, 2, 2, 185, 186, 7, 64, 2,
	2, 186, 187, 7, 7, 2, 2, 187, 188, 5, 6, 4, 2, 188, 189, 7, 8, 2, 2, 189,
	190, 8, 14, 1, 2, 190, 27, 3, 2, 2, 2, 191, 192, 7, 65, 2, 2, 192, 193,
	7, 9, 2, 2, 193, 194, 8, 15, 1, 2, 194, 29, 3, 2, 2, 2, 195, 196, 7, 40,
	2, 2, 196, 197, 7, 41, 2, 2, 197, 198, 7, 72, 2, 2, 198, 199, 7, 15, 2,
	2, 199, 200, 5, 34, 18, 2, 200, 201, 7, 9, 2, 2, 201, 202, 8, 16, 1, 2,
	202, 230, 3, 2, 2, 2, 203, 204, 7, 40, 2, 2, 204, 205, 7, 72, 2, 2, 205,
	206, 7, 15, 2, 2, 206, 207, 5, 34, 18, 2, 207, 208, 7, 9, 2, 2, 208, 209,
	8, 16, 1, 2, 209, 230, 3, 2, 2, 2, 210, 211, 7, 40, 2, 2, 211, 212, 7,
	41, 2, 2, 212, 213, 7, 72, 2, 2, 213, 214, 7, 10, 2, 2, 214, 215, 5, 32,
	17, 2, 215, 216, 7, 15, 2, 2, 216, 217, 5, 34, 18, 2, 217, 218, 7, 9, 2,
	2, 218, 219, 8, 16, 1, 2, 219, 230, 3, 2, 2, 2, 220, 221, 7, 40, 2, 2,
	221, 222, 7, 72, 2, 2, 222, 223, 7, 10, 2, 2, 223, 224, 5, 32, 17, 2, 224,
	225, 7, 15, 2, 2, 225, 226, 5, 34, 18, 2, 226, 227, 7, 9, 2, 2, 227, 228,
	8, 16, 1, 2, 228, 230, 3, 2, 2, 2, 229, 195, 3, 2, 2, 2, 229, 203, 3, 2,
	2, 2, 229, 210, 3, 2, 2, 2, 229, 220, 3, 2, 2, 2, 230, 31, 3, 2, 2, 2,
	231, 232, 7, 33, 2, 2, 232, 242, 8, 17, 1, 2, 233, 234, 7, 34, 2, 2, 234,
	242, 8, 17, 1, 2, 235, 236, 7, 35, 2, 2, 236, 242, 8, 17, 1, 2, 237, 238,
	7, 36, 2, 2, 238, 242, 8, 17, 1, 2, 239, 240, 7, 37, 2, 2, 240, 242, 8,
	17, 1, 2, 241, 231, 3, 2, 2, 2, 241, 233, 3, 2, 2, 2, 241, 235, 3, 2, 2,
	2, 241, 237, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 242, 33, 3, 2, 2, 2, 243,
	244, 8, 18, 1, 2, 244, 245, 9, 2, 2, 2, 245, 246, 5, 34, 18, 17, 246, 247,
	8, 18, 1, 2, 247, 273, 3, 2, 2, 2, 248, 249, 7, 66, 2, 2, 249, 250, 7,
	3, 2, 2, 250, 251, 5, 34, 18, 2, 251, 252, 7, 11, 2, 2, 252, 253, 5, 34,
	18, 2, 253, 254, 7, 4, 2, 2, 254, 255, 8, 18, 1, 2, 255, 273, 3, 2, 2,
	2, 256, 257, 7, 67, 2, 2, 257, 258, 7, 3, 2, 2, 258, 259, 5, 34, 18, 2,
	259, 260, 7, 11, 2, 2, 260, 261, 5, 34, 18, 2, 261, 262, 7, 4, 2, 2, 262,
	263, 8, 18, 1, 2, 263, 273, 3, 2, 2, 2, 264, 265, 7, 3, 2, 2, 265, 266,
	5, 34, 18, 2, 266, 267, 7, 4, 2, 2, 267, 268, 8, 18, 1, 2, 268, 273, 3,
	2, 2, 2, 269, 270, 5, 36, 19, 2, 270, 271, 8, 18, 1, 2, 271, 273, 3, 2,
	2, 2, 272, 243, 3, 2, 2, 2, 272, 248, 3, 2, 2, 2, 272, 256, 3, 2, 2, 2,
	272, 264, 3, 2, 2, 2, 272, 269, 3, 2, 2, 2, 273, 326, 3, 2, 2, 2, 274,
	275, 12, 16, 2, 2, 275, 276, 9, 3, 2, 2, 276, 277, 5, 34, 18, 17, 277,
	278, 8, 18, 1, 2, 278, 325, 3, 2, 2, 2, 279, 280, 12, 15, 2, 2, 280, 281,
	9, 4, 2, 2, 281, 282, 5, 34, 18, 16, 282, 283, 8, 18, 1, 2, 283, 325, 3,
	2, 2, 2, 284, 285, 12, 12, 2, 2, 285, 286, 7, 26, 2, 2, 286, 287, 5, 34,
	18, 13, 287, 288, 8, 18, 1, 2, 288, 325, 3, 2, 2, 2, 289, 290, 12, 11,
	2, 2, 290, 291, 7, 27, 2, 2, 291, 292, 5, 34, 18, 12, 292, 293, 8, 18,
	1, 2, 293, 325, 3, 2, 2, 2, 294, 295, 12, 10, 2, 2, 295, 296, 7, 13, 2,
	2, 296, 297, 5, 34, 18, 11, 297, 298, 8, 18, 1, 2, 298, 325, 3, 2, 2, 2,
	299, 300, 12, 9, 2, 2, 300, 301, 7, 12, 2, 2, 301, 302, 5, 34, 18, 10,
	302, 303, 8, 18, 1, 2, 303, 325, 3, 2, 2, 2, 304, 305, 12, 8, 2, 2, 305,
	306, 7, 24, 2, 2, 306, 307, 5, 34, 18, 9, 307, 308, 8, 18, 1, 2, 308, 325,
	3, 2, 2, 2, 309, 310, 12, 7, 2, 2, 310, 311, 7, 25, 2, 2, 311, 312, 5,
	34, 18, 8, 312, 313, 8, 18, 1, 2, 313, 325, 3, 2, 2, 2, 314, 315, 12, 6,
	2, 2, 315, 316, 7, 29, 2, 2, 316, 317, 5, 34, 18, 7, 317, 318, 8, 18, 1,
	2, 318, 325, 3, 2, 2, 2, 319, 320, 12, 5, 2, 2, 320, 321, 7, 28, 2, 2,
	321, 322, 5, 34, 18, 6, 322, 323, 8, 18, 1, 2, 323, 325, 3, 2, 2, 2, 324,
	274, 3, 2, 2, 2, 324, 279, 3, 2, 2, 2, 324, 284, 3, 2, 2, 2, 324, 289,
	3, 2, 2, 2, 324, 294, 3, 2, 2, 2, 324, 299, 3, 2, 2, 2, 324, 304, 3, 2,
	2, 2, 324, 309, 3, 2, 2, 2, 324, 314, 3, 2, 2, 2, 324, 319, 3, 2, 2, 2,
	325, 328, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327,
	35, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 329, 330, 7, 68, 2, 2, 330, 344,
	8, 19, 1, 2, 331, 332, 7, 69, 2, 2, 332, 344, 8, 19, 1, 2, 333, 334, 7,
	44, 2, 2, 334, 344, 8, 19, 1, 2, 335, 336, 7, 45, 2, 2, 336, 344, 8, 19,
	1, 2, 337, 338, 7, 70, 2, 2, 338, 344, 8, 19, 1, 2, 339, 340, 7, 71, 2,
	2, 340, 344, 8, 19, 1, 2, 341, 342, 7, 72, 2, 2, 342, 344, 8, 19, 1, 2,
	343, 329, 3, 2, 2, 2, 343, 331, 3, 2, 2, 2, 343, 333, 3, 2, 2, 2, 343,
	335, 3, 2, 2, 2, 343, 337, 3, 2, 2, 2, 343, 339, 3, 2, 2, 2, 343, 341,
	3, 2, 2, 2, 344, 37, 3, 2, 2, 2, 15, 57, 62, 88, 105, 110, 162, 167, 229,
	241, 272, 324, 326, 343,
}
var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", "','", "'<'",
	"'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", "'&'", "'!'",
	"'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'?'", "'->'", "'println!'",
	"'i64'", "'f64'", "'bool'", "'char'", "'String'", "'main'", "'usize'",
	"'let'", "'mut'", "'struct'", "'as'", "'true'", "'false'", "'fn'", "'return'",
	"'abs'", "'sqrt'", "'to_string'", "'clone'", "'new'", "'len'", "'push'",
	"'remove'", "'contains'", "'insert'", "'capacity'", "'with_capacity'",
	"'if'", "'else'", "'else if'", "'while'", "'loop'", "'break'", "'i64::pow'",
	"'f64::powf'",
}
var symbolicNames = []string{
	"", "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", "TK_CORCHETE_RIGHT",
	"TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA", "TK_DOSPUNTOS", "TK_COMA",
	"TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL", "TK_MAS", "TK_MENOS", "TK_POR",
	"TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA", "TK_AMPERSAND", "TK_ADMIRACION",
	"TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL", "TK_DIFERENTE", "TK_OR",
	"TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION", "TK_INT", "TK_FLOAT",
	"TK_BOOL", "TK_CHAR", "TK_STRING", "TK_MAIN", "TK_USIZE", "TK_LET", "TK_MUT",
	"TK_STRUCT", "TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN", "TK_RETURN", "TK_ABS",
	"TK_SQRT", "TK_TOSTRING", "TK_CLONE", "TK_NEW", "TK_LEN", "TK_PUSH", "TK_REMOVED",
	"TK_CONTAINS", "TK_INSERT", "TK_CAPACITY", "TK_WITHCAPACITY", "TK_IF",
	"TK_ELSE", "TK_ELSE_IF", "TK_WHILE", "TK_LOOP", "TK_BREAK", "TK_POW_I64",
	"TK_POW_F64", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA", "TK_CARACTER", "TK_ID",
	"TK_COMMET", "SPACES",
}

var ruleNames = []string{
	"start", "funcionmain", "instrucciones", "instruccion", "impresion", "listprint",
	"expimprimir", "asignacionVariable", "expresionIf", "listaelif", "elif",
	"expresionWhile", "expresionLoop", "breakInst", "variable", "tipo", "expresion",
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
	rustParserTK_MAIN             = 36
	rustParserTK_USIZE            = 37
	rustParserTK_LET              = 38
	rustParserTK_MUT              = 39
	rustParserTK_STRUCT           = 40
	rustParserTK_AS               = 41
	rustParserTK_TRUE             = 42
	rustParserTK_FALSE            = 43
	rustParserTK_FN               = 44
	rustParserTK_RETURN           = 45
	rustParserTK_ABS              = 46
	rustParserTK_SQRT             = 47
	rustParserTK_TOSTRING         = 48
	rustParserTK_CLONE            = 49
	rustParserTK_NEW              = 50
	rustParserTK_LEN              = 51
	rustParserTK_PUSH             = 52
	rustParserTK_REMOVED          = 53
	rustParserTK_CONTAINS         = 54
	rustParserTK_INSERT           = 55
	rustParserTK_CAPACITY         = 56
	rustParserTK_WITHCAPACITY     = 57
	rustParserTK_IF               = 58
	rustParserTK_ELSE             = 59
	rustParserTK_ELSE_IF          = 60
	rustParserTK_WHILE            = 61
	rustParserTK_LOOP             = 62
	rustParserTK_BREAK            = 63
	rustParserTK_POW_I64          = 64
	rustParserTK_POW_F64          = 65
	rustParserTK_NUMBER           = 66
	rustParserTK_DECIMAL          = 67
	rustParserTK_CADENA           = 68
	rustParserTK_CARACTER         = 69
	rustParserTK_ID               = 70
	rustParserTK_COMMET           = 71
	rustParserSPACES              = 72
)

// rustParser rules.
const (
	rustParserRULE_start              = 0
	rustParserRULE_funcionmain        = 1
	rustParserRULE_instrucciones      = 2
	rustParserRULE_instruccion        = 3
	rustParserRULE_impresion          = 4
	rustParserRULE_listprint          = 5
	rustParserRULE_expimprimir        = 6
	rustParserRULE_asignacionVariable = 7
	rustParserRULE_expresionIf        = 8
	rustParserRULE_listaelif          = 9
	rustParserRULE_elif               = 10
	rustParserRULE_expresionWhile     = 11
	rustParserRULE_expresionLoop      = 12
	rustParserRULE_breakInst          = 13
	rustParserRULE_variable           = 14
	rustParserRULE_tipo               = 15
	rustParserRULE_expresion          = 16
	rustParserRULE_valor              = 17
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

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
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

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

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
		p.SetState(36)

		var _x = p.Funcionmain()

		localctx.(*StartContext)._funcionmain = _x
	}
	{
		p.SetState(37)
		p.Match(rustParserEOF)
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_funcionmain().GetLista()

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(41)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(42)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(43)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(44)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(45)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(47)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(48)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(49)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(50)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(51)

			var _x = p.Instrucciones()

			localctx.(*FuncionmainContext)._instrucciones = _x
		}
		{
			p.SetState(52)
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
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(rustParserTK_IMPRESION-30))|(1<<(rustParserTK_LET-30))|(1<<(rustParserTK_IF-30))|(1<<(rustParserTK_WHILE-30)))) != 0) || (((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(rustParserTK_LOOP-62))|(1<<(rustParserTK_BREAK-62))|(1<<(rustParserTK_ID-62)))) != 0) {
		{
			p.SetState(57)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(62)
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

	// Get_expresionLoop returns the _expresionLoop rule contexts.
	Get_expresionLoop() IExpresionLoopContext

	// Get_breakInst returns the _breakInst rule contexts.
	Get_breakInst() IBreakInstContext

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

	// Set_expresionLoop sets the _expresionLoop rule contexts.
	Set_expresionLoop(IExpresionLoopContext)

	// Set_breakInst sets the _breakInst rule contexts.
	Set_breakInst(IBreakInstContext)

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
	_expresionLoop      IExpresionLoopContext
	_breakInst          IBreakInstContext
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

func (s *InstruccionContext) Get_expresionLoop() IExpresionLoopContext { return s._expresionLoop }

func (s *InstruccionContext) Get_breakInst() IBreakInstContext { return s._breakInst }

func (s *InstruccionContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_asignacionVariable(v IAsignacionVariableContext) {
	s._asignacionVariable = v
}

func (s *InstruccionContext) Set_expresionIf(v IExpresionIfContext) { s._expresionIf = v }

func (s *InstruccionContext) Set_expresionWhile(v IExpresionWhileContext) { s._expresionWhile = v }

func (s *InstruccionContext) Set_expresionLoop(v IExpresionLoopContext) { s._expresionLoop = v }

func (s *InstruccionContext) Set_breakInst(v IBreakInstContext) { s._breakInst = v }

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

func (s *InstruccionContext) ExpresionLoop() IExpresionLoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionLoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionLoopContext)
}

func (s *InstruccionContext) BreakInst() IBreakInstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakInstContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakInstContext)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_LET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)

			var _x = p.Variable()

			localctx.(*InstruccionContext)._variable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_variable().GetInst()

	case rustParserTK_IMPRESION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case rustParserTK_ID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(71)

			var _x = p.AsignacionVariable()

			localctx.(*InstruccionContext)._asignacionVariable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_asignacionVariable().GetInst()

	case rustParserTK_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(74)

			var _x = p.ExpresionIf()

			localctx.(*InstruccionContext)._expresionIf = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionIf().GetInst()

	case rustParserTK_WHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(77)

			var _x = p.ExpresionWhile()

			localctx.(*InstruccionContext)._expresionWhile = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionWhile().GetInst()

	case rustParserTK_LOOP:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(80)

			var _x = p.ExpresionLoop()

			localctx.(*InstruccionContext)._expresionLoop = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionLoop().GetInst()

	case rustParserTK_BREAK:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(83)

			var _x = p.BreakInst()

			localctx.(*InstruccionContext)._breakInst = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_breakInst().GetInst()

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

	// Get_listprint returns the _listprint rule contexts.
	Get_listprint() IListprintContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_listprint sets the _listprint rule contexts.
	Set_listprint(IListprintContext)

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
	_listprint IListprintContext
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

func (s *ImpresionContext) Get_listprint() IListprintContext { return s._listprint }

func (s *ImpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ImpresionContext) Set_listprint(v IListprintContext) { s._listprint = v }

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

func (s *ImpresionContext) Listprint() IListprintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListprintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListprintContext)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(89)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(90)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext)._expresion = _x
		}
		{
			p.SetState(91)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(92)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*ImpresionContext).inst = instrucciones.NewImprimir(localctx.(*ImpresionContext).Get_expresion().GetPrimate(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(96)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(97)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext)._expresion = _x
		}
		{
			p.SetState(98)

			var _x = p.Listprint()

			localctx.(*ImpresionContext)._listprint = _x
		}
		{
			p.SetState(99)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(100)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*ImpresionContext).inst = instrucciones.NewImprimir(localctx.(*ImpresionContext).Get_expresion().GetPrimate(), localctx.(*ImpresionContext).Get_listprint().GetLista())

	}

	return localctx
}

// IListprintContext is an interface to support dynamic dispatch.
type IListprintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expimprimir returns the _expimprimir rule contexts.
	Get_expimprimir() IExpimprimirContext

	// Set_expimprimir sets the _expimprimir rule contexts.
	Set_expimprimir(IExpimprimirContext)

	// GetList returns the list rule context list.
	GetList() []IExpimprimirContext

	// SetList sets the list rule context list.
	SetList([]IExpimprimirContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListprintContext differentiates from other interfaces.
	IsListprintContext()
}

type ListprintContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	lista        *arrayList.List
	_expimprimir IExpimprimirContext
	list         []IExpimprimirContext
}

func NewEmptyListprintContext() *ListprintContext {
	var p = new(ListprintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_listprint
	return p
}

func (*ListprintContext) IsListprintContext() {}

func NewListprintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListprintContext {
	var p = new(ListprintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_listprint

	return p
}

func (s *ListprintContext) GetParser() antlr.Parser { return s.parser }

func (s *ListprintContext) Get_expimprimir() IExpimprimirContext { return s._expimprimir }

func (s *ListprintContext) Set_expimprimir(v IExpimprimirContext) { s._expimprimir = v }

func (s *ListprintContext) GetList() []IExpimprimirContext { return s.list }

func (s *ListprintContext) SetList(v []IExpimprimirContext) { s.list = v }

func (s *ListprintContext) GetLista() *arrayList.List { return s.lista }

func (s *ListprintContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListprintContext) AllExpimprimir() []IExpimprimirContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpimprimirContext)(nil)).Elem())
	var tst = make([]IExpimprimirContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpimprimirContext)
		}
	}

	return tst
}

func (s *ListprintContext) Expimprimir(i int) IExpimprimirContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpimprimirContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpimprimirContext)
}

func (s *ListprintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListprintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListprintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterListprint(s)
	}
}

func (s *ListprintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitListprint(s)
	}
}

func (p *rustParser) Listprint() (localctx IListprintContext) {
	this := p
	_ = this

	localctx = NewListprintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, rustParserRULE_listprint)
	localctx.(*ListprintContext).lista = arrayList.New()
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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rustParserTK_COMA {
		{
			p.SetState(105)

			var _x = p.Expimprimir()

			localctx.(*ListprintContext)._expimprimir = _x
		}
		localctx.(*ListprintContext).list = append(localctx.(*ListprintContext).list, localctx.(*ListprintContext)._expimprimir)

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	localList := localctx.(*ListprintContext).GetList()
	for _, list := range localList {
		localctx.(*ListprintContext).lista.Add(list.GetPrimate())
	}

	return localctx
}

// IExpimprimirContext is an interface to support dynamic dispatch.
type IExpimprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetPrimate returns the primate attribute.
	GetPrimate() interfaces.Expresion

	// SetPrimate sets the primate attribute.
	SetPrimate(interfaces.Expresion)

	// IsExpimprimirContext differentiates from other interfaces.
	IsExpimprimirContext()
}

type ExpimprimirContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	primate    interfaces.Expresion
	_expresion IExpresionContext
}

func NewEmptyExpimprimirContext() *ExpimprimirContext {
	var p = new(ExpimprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_expimprimir
	return p
}

func (*ExpimprimirContext) IsExpimprimirContext() {}

func NewExpimprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpimprimirContext {
	var p = new(ExpimprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_expimprimir

	return p
}

func (s *ExpimprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpimprimirContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpimprimirContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpimprimirContext) GetPrimate() interfaces.Expresion { return s.primate }

func (s *ExpimprimirContext) SetPrimate(v interfaces.Expresion) { s.primate = v }

func (s *ExpimprimirContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_COMA, 0)
}

func (s *ExpimprimirContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpimprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpimprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpimprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterExpimprimir(s)
	}
}

func (s *ExpimprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitExpimprimir(s)
	}
}

func (p *rustParser) Expimprimir() (localctx IExpimprimirContext) {
	this := p
	_ = this

	localctx = NewExpimprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, rustParserRULE_expimprimir)

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
		p.SetState(112)
		p.Match(rustParserTK_COMA)
	}
	{
		p.SetState(113)

		var _x = p.expresion(0)

		localctx.(*ExpimprimirContext)._expresion = _x
	}
	localctx.(*ExpimprimirContext).primate = localctx.(*ExpimprimirContext).Get_expresion().GetPrimate()

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
	p.EnterRule(localctx, 14, rustParserRULE_asignacionVariable)

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
		p.SetState(116)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*AsignacionVariableContext)._TK_ID = _m
	}
	{
		p.SetState(117)
		p.Match(rustParserTK_IGUAL)
	}
	{
		p.SetState(118)

		var _x = p.expresion(0)

		localctx.(*AsignacionVariableContext)._expresion = _x
	}
	{
		p.SetState(119)
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
	p.EnterRule(localctx, 16, rustParserRULE_expresionIf)

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

	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(123)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(124)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(125)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(126)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(130)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(131)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(132)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(133)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(134)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(135)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(136)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(137)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), localctx.(*ExpresionIfContext).GetBloqueelse().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(141)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(142)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(143)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(144)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(145)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf2(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetIfblock().GetLista(), nil, localctx.(*ExpresionIfContext).Get_listaelif().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(149)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(150)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(151)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(152)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(153)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		{
			p.SetState(154)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(155)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(156)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(157)
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
	p.EnterRule(localctx, 18, rustParserRULE_listaelif)
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
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rustParserTK_ELSE_IF {
		{
			p.SetState(162)

			var _x = p.Elif()

			localctx.(*ListaelifContext)._elif = _x
		}
		localctx.(*ListaelifContext).l = append(localctx.(*ListaelifContext).l, localctx.(*ListaelifContext)._elif)

		p.SetState(165)
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
	p.EnterRule(localctx, 20, rustParserRULE_elif)

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
		p.SetState(169)
		p.Match(rustParserTK_ELSE_IF)
	}
	{
		p.SetState(170)

		var _x = p.expresion(0)

		localctx.(*ElifContext)._expresion = _x
	}
	{
		p.SetState(171)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(172)

		var _x = p.Instrucciones()

		localctx.(*ElifContext).elifblock = _x
	}
	{
		p.SetState(173)
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
	p.EnterRule(localctx, 22, rustParserRULE_expresionWhile)

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
		p.SetState(176)
		p.Match(rustParserTK_WHILE)
	}
	{
		p.SetState(177)

		var _x = p.expresion(0)

		localctx.(*ExpresionWhileContext).exp = _x
	}
	{
		p.SetState(178)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(179)

		var _x = p.Instrucciones()

		localctx.(*ExpresionWhileContext)._instrucciones = _x
	}
	{
		p.SetState(180)
		p.Match(rustParserTK_LLAVE_RIGHT)
	}
	localctx.(*ExpresionWhileContext).inst = instrucciones.NewWhile(localctx.(*ExpresionWhileContext).GetExp().GetPrimate(), localctx.(*ExpresionWhileContext).Get_instrucciones().GetLista())

	return localctx
}

// IExpresionLoopContext is an interface to support dynamic dispatch.
type IExpresionLoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsExpresionLoopContext differentiates from other interfaces.
	IsExpresionLoopContext()
}

type ExpresionLoopContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	inst           interfaces.Instruction
	_instrucciones IInstruccionesContext
}

func NewEmptyExpresionLoopContext() *ExpresionLoopContext {
	var p = new(ExpresionLoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_expresionLoop
	return p
}

func (*ExpresionLoopContext) IsExpresionLoopContext() {}

func NewExpresionLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionLoopContext {
	var p = new(ExpresionLoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_expresionLoop

	return p
}

func (s *ExpresionLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionLoopContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *ExpresionLoopContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *ExpresionLoopContext) GetInst() interfaces.Instruction { return s.inst }

func (s *ExpresionLoopContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *ExpresionLoopContext) TK_LOOP() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LOOP, 0)
}

func (s *ExpresionLoopContext) TK_LLAVE_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_LEFT, 0)
}

func (s *ExpresionLoopContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *ExpresionLoopContext) TK_LLAVE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LLAVE_RIGHT, 0)
}

func (s *ExpresionLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterExpresionLoop(s)
	}
}

func (s *ExpresionLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitExpresionLoop(s)
	}
}

func (p *rustParser) ExpresionLoop() (localctx IExpresionLoopContext) {
	this := p
	_ = this

	localctx = NewExpresionLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, rustParserRULE_expresionLoop)

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
		p.SetState(183)
		p.Match(rustParserTK_LOOP)
	}
	{
		p.SetState(184)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(185)

		var _x = p.Instrucciones()

		localctx.(*ExpresionLoopContext)._instrucciones = _x
	}
	{
		p.SetState(186)
		p.Match(rustParserTK_LLAVE_RIGHT)
	}
	localctx.(*ExpresionLoopContext).inst = instrucciones.NewLoop(localctx.(*ExpresionLoopContext).Get_instrucciones().GetLista())

	return localctx
}

// IBreakInstContext is an interface to support dynamic dispatch.
type IBreakInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsBreakInstContext differentiates from other interfaces.
	IsBreakInstContext()
}

type BreakInstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	inst   interfaces.Instruction
}

func NewEmptyBreakInstContext() *BreakInstContext {
	var p = new(BreakInstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_breakInst
	return p
}

func (*BreakInstContext) IsBreakInstContext() {}

func NewBreakInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakInstContext {
	var p = new(BreakInstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_breakInst

	return p
}

func (s *BreakInstContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakInstContext) GetInst() interfaces.Instruction { return s.inst }

func (s *BreakInstContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *BreakInstContext) TK_BREAK() antlr.TerminalNode {
	return s.GetToken(rustParserTK_BREAK, 0)
}

func (s *BreakInstContext) TK_PUNTO_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, 0)
}

func (s *BreakInstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakInstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakInstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterBreakInst(s)
	}
}

func (s *BreakInstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitBreakInst(s)
	}
}

func (p *rustParser) BreakInst() (localctx IBreakInstContext) {
	this := p
	_ = this

	localctx = NewBreakInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, rustParserRULE_breakInst)

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
		p.SetState(189)
		p.Match(rustParserTK_BREAK)
	}
	{
		p.SetState(190)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*BreakInstContext).inst = instrucciones.NewBreak(interfaces.BREAK)

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
	p.EnterRule(localctx, 28, rustParserRULE_variable)

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

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(194)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(195)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(196)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(197)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(198)
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
			p.SetState(201)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(202)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(203)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(204)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(205)
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
			p.SetState(208)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(209)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(210)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(211)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(212)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(213)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(214)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(215)
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
			p.SetState(218)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(219)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(220)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(221)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(222)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(223)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(224)
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
	p.EnterRule(localctx, 30, rustParserRULE_tipo)

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

	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(235)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(237)
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

	// Get_TK_POW_I64 returns the _TK_POW_I64 token.
	Get_TK_POW_I64() antlr.Token

	// Get_TK_POW_F64 returns the _TK_POW_F64 token.
	Get_TK_POW_F64() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_TK_POW_I64 sets the _TK_POW_I64 token.
	Set_TK_POW_I64(antlr.Token)

	// Set_TK_POW_F64 sets the _TK_POW_F64 token.
	Set_TK_POW_F64(antlr.Token)

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
	parser      antlr.Parser
	primate     interfaces.Expresion
	left        IExpresionContext
	op          antlr.Token
	right       IExpresionContext
	_TK_POW_I64 antlr.Token
	_TK_POW_F64 antlr.Token
	_valor      IValorContext
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

func (s *ExpresionContext) Get_TK_POW_I64() antlr.Token { return s._TK_POW_I64 }

func (s *ExpresionContext) Get_TK_POW_F64() antlr.Token { return s._TK_POW_F64 }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) Set_TK_POW_I64(v antlr.Token) { s._TK_POW_I64 = v }

func (s *ExpresionContext) Set_TK_POW_F64(v antlr.Token) { s._TK_POW_F64 = v }

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

func (s *ExpresionContext) TK_POW_I64() antlr.TerminalNode {
	return s.GetToken(rustParserTK_POW_I64, 0)
}

func (s *ExpresionContext) TK_PARENTESIS_LEFT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_LEFT, 0)
}

func (s *ExpresionContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_COMA, 0)
}

func (s *ExpresionContext) TK_PARENTESIS_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_PARENTESIS_RIGHT, 0)
}

func (s *ExpresionContext) TK_POW_F64() antlr.TerminalNode {
	return s.GetToken(rustParserTK_POW_F64, 0)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, rustParserRULE_expresion, _p)
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
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_MENOS, rustParserTK_ADMIRACION:
		{
			p.SetState(242)

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
			p.SetState(243)

			var _x = p.expresion(15)

			localctx.(*ExpresionContext).right = _x
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(nil, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case rustParserTK_POW_I64:
		{
			p.SetState(246)

			var _m = p.Match(rustParserTK_POW_I64)

			localctx.(*ExpresionContext)._TK_POW_I64 = _m
		}
		{
			p.SetState(247)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(248)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).left = _x
		}
		{
			p.SetState(249)
			p.Match(rustParserTK_COMA)
		}
		{
			p.SetState(250)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).right = _x
		}
		{
			p.SetState(251)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
			if localctx.(*ExpresionContext).Get_TK_POW_I64() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TK_POW_I64().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case rustParserTK_POW_F64:
		{
			p.SetState(254)

			var _m = p.Match(rustParserTK_POW_F64)

			localctx.(*ExpresionContext)._TK_POW_F64 = _m
		}
		{
			p.SetState(255)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(256)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).left = _x
		}
		{
			p.SetState(257)
			p.Match(rustParserTK_COMA)
		}
		{
			p.SetState(258)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).right = _x
		}
		{
			p.SetState(259)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
			if localctx.(*ExpresionContext).Get_TK_POW_F64() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TK_POW_F64().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case rustParserTK_PARENTESIS_LEFT:
		{
			p.SetState(262)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(263)
			p.expresion(0)
		}
		{
			p.SetState(264)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	case rustParserTK_TRUE, rustParserTK_FALSE, rustParserTK_NUMBER, rustParserTK_DECIMAL, rustParserTK_CADENA, rustParserTK_CARACTER, rustParserTK_ID:
		{
			p.SetState(267)

			var _x = p.Valor()

			localctx.(*ExpresionContext)._valor = _x
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(273)

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
					p.SetState(274)

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

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(278)

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
					p.SetState(279)

					var _x = p.expresion(14)

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
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(283)

					var _m = p.Match(rustParserTK_IGUALIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(284)

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

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(288)

					var _m = p.Match(rustParserTK_DIFERENTE)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(289)

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

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(293)

					var _m = p.Match(rustParserTK_MAYOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(294)

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

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(298)

					var _m = p.Match(rustParserTK_MENOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(299)

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

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(303)

					var _m = p.Match(rustParserTK_MAYORIGULA)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(304)

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

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(308)

					var _m = p.Match(rustParserTK_MENOIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(309)

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

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(312)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(313)

					var _m = p.Match(rustParserTK_AND)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(314)

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

			case 10:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(317)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(318)

					var _m = p.Match(rustParserTK_OR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(319)

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

			}

		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, rustParserRULE_valor)

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

	p.SetState(341)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_NUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)

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
			p.SetState(329)

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
			p.SetState(331)

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
			p.SetState(333)

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
			p.SetState(335)

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
			p.SetState(337)

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
			p.SetState(339)

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
	case 16:
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
