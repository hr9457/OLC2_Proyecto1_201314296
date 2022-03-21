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
import "Proyecto1/src/arreglos"
import arrayList "github.com/colegno/arraylist"

// import "Proyecto1/src/pruebas"
// import "reflect"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 74, 406,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 64, 10, 3, 3, 4, 7, 4, 67, 10, 4, 12,
	4, 14, 4, 70, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 98, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 115, 10, 6, 3, 7, 6, 7, 118, 10, 7, 13, 7, 14, 7, 119, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 162, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 192, 10, 12, 12, 12, 14, 12, 195, 11, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 235, 10, 13, 3, 14, 6, 14, 238, 10, 14,
	13, 14, 14, 14, 239, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 278, 10, 19,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 5, 20, 315, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 367, 10, 20, 12, 20, 14, 20, 370, 11,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 389, 10, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22,
	401, 10, 22, 12, 22, 14, 22, 404, 11, 22, 3, 22, 2, 5, 22, 38, 42, 23,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 2, 5, 4, 2, 17, 17, 23, 23, 3, 2, 18, 20, 3, 2, 16, 17, 2, 431,
	2, 44, 3, 2, 2, 2, 4, 63, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 97, 3, 2, 2,
	2, 10, 114, 3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14, 123, 3, 2, 2, 2, 16, 161,
	3, 2, 2, 2, 18, 163, 3, 2, 2, 2, 20, 169, 3, 2, 2, 2, 22, 182, 3, 2, 2,
	2, 24, 234, 3, 2, 2, 2, 26, 237, 3, 2, 2, 2, 28, 243, 3, 2, 2, 2, 30, 250,
	3, 2, 2, 2, 32, 257, 3, 2, 2, 2, 34, 263, 3, 2, 2, 2, 36, 277, 3, 2, 2,
	2, 38, 314, 3, 2, 2, 2, 40, 388, 3, 2, 2, 2, 42, 390, 3, 2, 2, 2, 44, 45,
	5, 4, 3, 2, 45, 46, 7, 2, 2, 3, 46, 47, 8, 2, 1, 2, 47, 3, 3, 2, 2, 2,
	48, 49, 7, 46, 2, 2, 49, 50, 7, 38, 2, 2, 50, 51, 7, 3, 2, 2, 51, 52, 7,
	4, 2, 2, 52, 53, 7, 7, 2, 2, 53, 64, 7, 8, 2, 2, 54, 55, 7, 46, 2, 2, 55,
	56, 7, 38, 2, 2, 56, 57, 7, 3, 2, 2, 57, 58, 7, 4, 2, 2, 58, 59, 7, 7,
	2, 2, 59, 60, 5, 6, 4, 2, 60, 61, 7, 8, 2, 2, 61, 62, 8, 3, 1, 2, 62, 64,
	3, 2, 2, 2, 63, 48, 3, 2, 2, 2, 63, 54, 3, 2, 2, 2, 64, 5, 3, 2, 2, 2,
	65, 67, 5, 8, 5, 2, 66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3,
	2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71,
	72, 8, 4, 1, 2, 72, 7, 3, 2, 2, 2, 73, 74, 5, 16, 9, 2, 74, 75, 8, 5, 1,
	2, 75, 98, 3, 2, 2, 2, 76, 77, 5, 10, 6, 2, 77, 78, 8, 5, 1, 2, 78, 98,
	3, 2, 2, 2, 79, 80, 5, 18, 10, 2, 80, 81, 8, 5, 1, 2, 81, 98, 3, 2, 2,
	2, 82, 83, 5, 24, 13, 2, 83, 84, 8, 5, 1, 2, 84, 98, 3, 2, 2, 2, 85, 86,
	5, 30, 16, 2, 86, 87, 8, 5, 1, 2, 87, 98, 3, 2, 2, 2, 88, 89, 5, 32, 17,
	2, 89, 90, 8, 5, 1, 2, 90, 98, 3, 2, 2, 2, 91, 92, 5, 34, 18, 2, 92, 93,
	8, 5, 1, 2, 93, 98, 3, 2, 2, 2, 94, 95, 5, 20, 11, 2, 95, 96, 8, 5, 1,
	2, 96, 98, 3, 2, 2, 2, 97, 73, 3, 2, 2, 2, 97, 76, 3, 2, 2, 2, 97, 79,
	3, 2, 2, 2, 97, 82, 3, 2, 2, 2, 97, 85, 3, 2, 2, 2, 97, 88, 3, 2, 2, 2,
	97, 91, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2, 98, 9, 3, 2, 2, 2, 99, 100, 7,
	32, 2, 2, 100, 101, 7, 3, 2, 2, 101, 102, 5, 38, 20, 2, 102, 103, 7, 4,
	2, 2, 103, 104, 7, 9, 2, 2, 104, 105, 8, 6, 1, 2, 105, 115, 3, 2, 2, 2,
	106, 107, 7, 32, 2, 2, 107, 108, 7, 3, 2, 2, 108, 109, 5, 38, 20, 2, 109,
	110, 5, 12, 7, 2, 110, 111, 7, 4, 2, 2, 111, 112, 7, 9, 2, 2, 112, 113,
	8, 6, 1, 2, 113, 115, 3, 2, 2, 2, 114, 99, 3, 2, 2, 2, 114, 106, 3, 2,
	2, 2, 115, 11, 3, 2, 2, 2, 116, 118, 5, 14, 8, 2, 117, 116, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120,
	121, 3, 2, 2, 2, 121, 122, 8, 7, 1, 2, 122, 13, 3, 2, 2, 2, 123, 124, 7,
	11, 2, 2, 124, 125, 5, 38, 20, 2, 125, 126, 8, 8, 1, 2, 126, 15, 3, 2,
	2, 2, 127, 128, 7, 40, 2, 2, 128, 129, 7, 41, 2, 2, 129, 130, 7, 72, 2,
	2, 130, 131, 7, 15, 2, 2, 131, 132, 5, 38, 20, 2, 132, 133, 7, 9, 2, 2,
	133, 134, 8, 9, 1, 2, 134, 162, 3, 2, 2, 2, 135, 136, 7, 40, 2, 2, 136,
	137, 7, 72, 2, 2, 137, 138, 7, 15, 2, 2, 138, 139, 5, 38, 20, 2, 139, 140,
	7, 9, 2, 2, 140, 141, 8, 9, 1, 2, 141, 162, 3, 2, 2, 2, 142, 143, 7, 40,
	2, 2, 143, 144, 7, 41, 2, 2, 144, 145, 7, 72, 2, 2, 145, 146, 7, 10, 2,
	2, 146, 147, 5, 36, 19, 2, 147, 148, 7, 15, 2, 2, 148, 149, 5, 38, 20,
	2, 149, 150, 7, 9, 2, 2, 150, 151, 8, 9, 1, 2, 151, 162, 3, 2, 2, 2, 152,
	153, 7, 40, 2, 2, 153, 154, 7, 72, 2, 2, 154, 155, 7, 10, 2, 2, 155, 156,
	5, 36, 19, 2, 156, 157, 7, 15, 2, 2, 157, 158, 5, 38, 20, 2, 158, 159,
	7, 9, 2, 2, 159, 160, 8, 9, 1, 2, 160, 162, 3, 2, 2, 2, 161, 127, 3, 2,
	2, 2, 161, 135, 3, 2, 2, 2, 161, 142, 3, 2, 2, 2, 161, 152, 3, 2, 2, 2,
	162, 17, 3, 2, 2, 2, 163, 164, 7, 72, 2, 2, 164, 165, 7, 15, 2, 2, 165,
	166, 5, 38, 20, 2, 166, 167, 7, 9, 2, 2, 167, 168, 8, 10, 1, 2, 168, 19,
	3, 2, 2, 2, 169, 170, 7, 40, 2, 2, 170, 171, 7, 72, 2, 2, 171, 172, 7,
	10, 2, 2, 172, 173, 7, 5, 2, 2, 173, 174, 5, 36, 19, 2, 174, 175, 7, 9,
	2, 2, 175, 176, 5, 40, 21, 2, 176, 177, 7, 6, 2, 2, 177, 178, 7, 15, 2,
	2, 178, 179, 5, 38, 20, 2, 179, 180, 7, 9, 2, 2, 180, 181, 8, 11, 1, 2,
	181, 21, 3, 2, 2, 2, 182, 183, 8, 12, 1, 2, 183, 184, 5, 38, 20, 2, 184,
	185, 8, 12, 1, 2, 185, 193, 3, 2, 2, 2, 186, 187, 12, 4, 2, 2, 187, 188,
	7, 11, 2, 2, 188, 189, 5, 38, 20, 2, 189, 190, 8, 12, 1, 2, 190, 192, 3,
	2, 2, 2, 191, 186, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2,
	2, 193, 194, 3, 2, 2, 2, 194, 23, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 196,
	197, 7, 60, 2, 2, 197, 198, 5, 38, 20, 2, 198, 199, 7, 7, 2, 2, 199, 200,
	5, 6, 4, 2, 200, 201, 7, 8, 2, 2, 201, 202, 8, 13, 1, 2, 202, 235, 3, 2,
	2, 2, 203, 204, 7, 60, 2, 2, 204, 205, 5, 38, 20, 2, 205, 206, 7, 7, 2,
	2, 206, 207, 5, 6, 4, 2, 207, 208, 7, 8, 2, 2, 208, 209, 7, 61, 2, 2, 209,
	210, 7, 7, 2, 2, 210, 211, 5, 6, 4, 2, 211, 212, 7, 8, 2, 2, 212, 213,
	8, 13, 1, 2, 213, 235, 3, 2, 2, 2, 214, 215, 7, 60, 2, 2, 215, 216, 5,
	38, 20, 2, 216, 217, 7, 7, 2, 2, 217, 218, 5, 6, 4, 2, 218, 219, 7, 8,
	2, 2, 219, 220, 5, 26, 14, 2, 220, 221, 8, 13, 1, 2, 221, 235, 3, 2, 2,
	2, 222, 223, 7, 60, 2, 2, 223, 224, 5, 38, 20, 2, 224, 225, 7, 7, 2, 2,
	225, 226, 5, 6, 4, 2, 226, 227, 7, 8, 2, 2, 227, 228, 5, 26, 14, 2, 228,
	229, 7, 61, 2, 2, 229, 230, 7, 7, 2, 2, 230, 231, 5, 6, 4, 2, 231, 232,
	7, 8, 2, 2, 232, 233, 8, 13, 1, 2, 233, 235, 3, 2, 2, 2, 234, 196, 3, 2,
	2, 2, 234, 203, 3, 2, 2, 2, 234, 214, 3, 2, 2, 2, 234, 222, 3, 2, 2, 2,
	235, 25, 3, 2, 2, 2, 236, 238, 5, 28, 15, 2, 237, 236, 3, 2, 2, 2, 238,
	239, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241,
	3, 2, 2, 2, 241, 242, 8, 14, 1, 2, 242, 27, 3, 2, 2, 2, 243, 244, 7, 62,
	2, 2, 244, 245, 5, 38, 20, 2, 245, 246, 7, 7, 2, 2, 246, 247, 5, 6, 4,
	2, 247, 248, 7, 8, 2, 2, 248, 249, 8, 15, 1, 2, 249, 29, 3, 2, 2, 2, 250,
	251, 7, 63, 2, 2, 251, 252, 5, 38, 20, 2, 252, 253, 7, 7, 2, 2, 253, 254,
	5, 6, 4, 2, 254, 255, 7, 8, 2, 2, 255, 256, 8, 16, 1, 2, 256, 31, 3, 2,
	2, 2, 257, 258, 7, 64, 2, 2, 258, 259, 7, 7, 2, 2, 259, 260, 5, 6, 4, 2,
	260, 261, 7, 8, 2, 2, 261, 262, 8, 17, 1, 2, 262, 33, 3, 2, 2, 2, 263,
	264, 7, 65, 2, 2, 264, 265, 7, 9, 2, 2, 265, 266, 8, 18, 1, 2, 266, 35,
	3, 2, 2, 2, 267, 268, 7, 33, 2, 2, 268, 278, 8, 19, 1, 2, 269, 270, 7,
	34, 2, 2, 270, 278, 8, 19, 1, 2, 271, 272, 7, 35, 2, 2, 272, 278, 8, 19,
	1, 2, 273, 274, 7, 36, 2, 2, 274, 278, 8, 19, 1, 2, 275, 276, 7, 37, 2,
	2, 276, 278, 8, 19, 1, 2, 277, 267, 3, 2, 2, 2, 277, 269, 3, 2, 2, 2, 277,
	271, 3, 2, 2, 2, 277, 273, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 278, 37, 3,
	2, 2, 2, 279, 280, 8, 20, 1, 2, 280, 281, 9, 2, 2, 2, 281, 282, 5, 38,
	20, 19, 282, 283, 8, 20, 1, 2, 283, 315, 3, 2, 2, 2, 284, 285, 7, 66, 2,
	2, 285, 286, 7, 3, 2, 2, 286, 287, 5, 38, 20, 2, 287, 288, 7, 11, 2, 2,
	288, 289, 5, 38, 20, 2, 289, 290, 7, 4, 2, 2, 290, 291, 8, 20, 1, 2, 291,
	315, 3, 2, 2, 2, 292, 293, 7, 67, 2, 2, 293, 294, 7, 3, 2, 2, 294, 295,
	5, 38, 20, 2, 295, 296, 7, 11, 2, 2, 296, 297, 5, 38, 20, 2, 297, 298,
	7, 4, 2, 2, 298, 299, 8, 20, 1, 2, 299, 315, 3, 2, 2, 2, 300, 301, 7, 3,
	2, 2, 301, 302, 5, 38, 20, 2, 302, 303, 7, 4, 2, 2, 303, 304, 8, 20, 1,
	2, 304, 315, 3, 2, 2, 2, 305, 306, 7, 5, 2, 2, 306, 307, 5, 22, 12, 2,
	307, 308, 7, 6, 2, 2, 308, 309, 8, 20, 1, 2, 309, 315, 3, 2, 2, 2, 310,
	311, 5, 40, 21, 2, 311, 312, 8, 20, 1, 2, 312, 315, 3, 2, 2, 2, 313, 315,
	3, 2, 2, 2, 314, 279, 3, 2, 2, 2, 314, 284, 3, 2, 2, 2, 314, 292, 3, 2,
	2, 2, 314, 300, 3, 2, 2, 2, 314, 305, 3, 2, 2, 2, 314, 310, 3, 2, 2, 2,
	314, 313, 3, 2, 2, 2, 315, 368, 3, 2, 2, 2, 316, 317, 12, 18, 2, 2, 317,
	318, 9, 3, 2, 2, 318, 319, 5, 38, 20, 19, 319, 320, 8, 20, 1, 2, 320, 367,
	3, 2, 2, 2, 321, 322, 12, 17, 2, 2, 322, 323, 9, 4, 2, 2, 323, 324, 5,
	38, 20, 18, 324, 325, 8, 20, 1, 2, 325, 367, 3, 2, 2, 2, 326, 327, 12,
	14, 2, 2, 327, 328, 7, 26, 2, 2, 328, 329, 5, 38, 20, 15, 329, 330, 8,
	20, 1, 2, 330, 367, 3, 2, 2, 2, 331, 332, 12, 13, 2, 2, 332, 333, 7, 27,
	2, 2, 333, 334, 5, 38, 20, 14, 334, 335, 8, 20, 1, 2, 335, 367, 3, 2, 2,
	2, 336, 337, 12, 12, 2, 2, 337, 338, 7, 13, 2, 2, 338, 339, 5, 38, 20,
	13, 339, 340, 8, 20, 1, 2, 340, 367, 3, 2, 2, 2, 341, 342, 12, 11, 2, 2,
	342, 343, 7, 12, 2, 2, 343, 344, 5, 38, 20, 12, 344, 345, 8, 20, 1, 2,
	345, 367, 3, 2, 2, 2, 346, 347, 12, 10, 2, 2, 347, 348, 7, 24, 2, 2, 348,
	349, 5, 38, 20, 11, 349, 350, 8, 20, 1, 2, 350, 367, 3, 2, 2, 2, 351, 352,
	12, 9, 2, 2, 352, 353, 7, 25, 2, 2, 353, 354, 5, 38, 20, 10, 354, 355,
	8, 20, 1, 2, 355, 367, 3, 2, 2, 2, 356, 357, 12, 8, 2, 2, 357, 358, 7,
	29, 2, 2, 358, 359, 5, 38, 20, 9, 359, 360, 8, 20, 1, 2, 360, 367, 3, 2,
	2, 2, 361, 362, 12, 7, 2, 2, 362, 363, 7, 28, 2, 2, 363, 364, 5, 38, 20,
	8, 364, 365, 8, 20, 1, 2, 365, 367, 3, 2, 2, 2, 366, 316, 3, 2, 2, 2, 366,
	321, 3, 2, 2, 2, 366, 326, 3, 2, 2, 2, 366, 331, 3, 2, 2, 2, 366, 336,
	3, 2, 2, 2, 366, 341, 3, 2, 2, 2, 366, 346, 3, 2, 2, 2, 366, 351, 3, 2,
	2, 2, 366, 356, 3, 2, 2, 2, 366, 361, 3, 2, 2, 2, 367, 370, 3, 2, 2, 2,
	368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 39, 3, 2, 2, 2, 370, 368,
	3, 2, 2, 2, 371, 372, 7, 68, 2, 2, 372, 389, 8, 21, 1, 2, 373, 374, 7,
	69, 2, 2, 374, 389, 8, 21, 1, 2, 375, 376, 7, 44, 2, 2, 376, 389, 8, 21,
	1, 2, 377, 378, 7, 45, 2, 2, 378, 389, 8, 21, 1, 2, 379, 380, 7, 70, 2,
	2, 380, 389, 8, 21, 1, 2, 381, 382, 7, 71, 2, 2, 382, 389, 8, 21, 1, 2,
	383, 384, 7, 72, 2, 2, 384, 389, 8, 21, 1, 2, 385, 386, 5, 42, 22, 2, 386,
	387, 8, 21, 1, 2, 387, 389, 3, 2, 2, 2, 388, 371, 3, 2, 2, 2, 388, 373,
	3, 2, 2, 2, 388, 375, 3, 2, 2, 2, 388, 377, 3, 2, 2, 2, 388, 379, 3, 2,
	2, 2, 388, 381, 3, 2, 2, 2, 388, 383, 3, 2, 2, 2, 388, 385, 3, 2, 2, 2,
	389, 41, 3, 2, 2, 2, 390, 391, 8, 22, 1, 2, 391, 392, 7, 72, 2, 2, 392,
	393, 8, 22, 1, 2, 393, 402, 3, 2, 2, 2, 394, 395, 12, 4, 2, 2, 395, 396,
	7, 5, 2, 2, 396, 397, 5, 38, 20, 2, 397, 398, 7, 6, 2, 2, 398, 399, 8,
	22, 1, 2, 399, 401, 3, 2, 2, 2, 400, 394, 3, 2, 2, 2, 401, 404, 3, 2, 2,
	2, 402, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 43, 3, 2, 2, 2, 404,
	402, 3, 2, 2, 2, 17, 63, 68, 97, 114, 119, 161, 193, 234, 239, 277, 314,
	366, 368, 388, 402,
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
	"expimprimir", "variable", "asignacionVariable", "declarcionarreglo", "listvalores",
	"expresionIf", "listaelif", "elif", "expresionWhile", "expresionLoop",
	"breakInst", "tipo", "expresion", "valor", "listaArreglo",
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
	rustParserRULE_variable           = 7
	rustParserRULE_asignacionVariable = 8
	rustParserRULE_declarcionarreglo  = 9
	rustParserRULE_listvalores        = 10
	rustParserRULE_expresionIf        = 11
	rustParserRULE_listaelif          = 12
	rustParserRULE_elif               = 13
	rustParserRULE_expresionWhile     = 14
	rustParserRULE_expresionLoop      = 15
	rustParserRULE_breakInst          = 16
	rustParserRULE_tipo               = 17
	rustParserRULE_expresion          = 18
	rustParserRULE_valor              = 19
	rustParserRULE_listaArreglo       = 20
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
		p.SetState(42)

		var _x = p.Funcionmain()

		localctx.(*StartContext)._funcionmain = _x
	}
	{
		p.SetState(43)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
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
			p.Match(rustParserTK_LLAVE_RIGHT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(rustParserTK_FN)
		}
		{
			p.SetState(53)
			p.Match(rustParserTK_MAIN)
		}
		{
			p.SetState(54)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(55)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(56)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(57)

			var _x = p.Instrucciones()

			localctx.(*FuncionmainContext)._instrucciones = _x
		}
		{
			p.SetState(58)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(rustParserTK_IMPRESION-30))|(1<<(rustParserTK_LET-30))|(1<<(rustParserTK_IF-30))|(1<<(rustParserTK_WHILE-30)))) != 0) || (((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(rustParserTK_LOOP-62))|(1<<(rustParserTK_BREAK-62))|(1<<(rustParserTK_ID-62)))) != 0) {
		{
			p.SetState(63)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(68)
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

	// Get_declarcionarreglo returns the _declarcionarreglo rule contexts.
	Get_declarcionarreglo() IDeclarcionarregloContext

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

	// Set_declarcionarreglo sets the _declarcionarreglo rule contexts.
	Set_declarcionarreglo(IDeclarcionarregloContext)

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
	_declarcionarreglo  IDeclarcionarregloContext
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

func (s *InstruccionContext) Get_declarcionarreglo() IDeclarcionarregloContext {
	return s._declarcionarreglo
}

func (s *InstruccionContext) Set_variable(v IVariableContext) { s._variable = v }

func (s *InstruccionContext) Set_impresion(v IImpresionContext) { s._impresion = v }

func (s *InstruccionContext) Set_asignacionVariable(v IAsignacionVariableContext) {
	s._asignacionVariable = v
}

func (s *InstruccionContext) Set_expresionIf(v IExpresionIfContext) { s._expresionIf = v }

func (s *InstruccionContext) Set_expresionWhile(v IExpresionWhileContext) { s._expresionWhile = v }

func (s *InstruccionContext) Set_expresionLoop(v IExpresionLoopContext) { s._expresionLoop = v }

func (s *InstruccionContext) Set_breakInst(v IBreakInstContext) { s._breakInst = v }

func (s *InstruccionContext) Set_declarcionarreglo(v IDeclarcionarregloContext) {
	s._declarcionarreglo = v
}

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

func (s *InstruccionContext) Declarcionarreglo() IDeclarcionarregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarcionarregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarcionarregloContext)
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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)

			var _x = p.Variable()

			localctx.(*InstruccionContext)._variable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_variable().GetInst()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)

			var _x = p.Impresion()

			localctx.(*InstruccionContext)._impresion = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_impresion().GetInst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)

			var _x = p.AsignacionVariable()

			localctx.(*InstruccionContext)._asignacionVariable = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_asignacionVariable().GetInst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)

			var _x = p.ExpresionIf()

			localctx.(*InstruccionContext)._expresionIf = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionIf().GetInst()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(83)

			var _x = p.ExpresionWhile()

			localctx.(*InstruccionContext)._expresionWhile = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionWhile().GetInst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(86)

			var _x = p.ExpresionLoop()

			localctx.(*InstruccionContext)._expresionLoop = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_expresionLoop().GetInst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(89)

			var _x = p.BreakInst()

			localctx.(*InstruccionContext)._breakInst = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_breakInst().GetInst()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)

			var _x = p.Declarcionarreglo()

			localctx.(*InstruccionContext)._declarcionarreglo = _x
		}
		localctx.(*InstruccionContext).inst = localctx.(*InstruccionContext).Get_declarcionarreglo().GetInst()

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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(98)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(99)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext)._expresion = _x
		}
		{
			p.SetState(100)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(101)
			p.Match(rustParserTK_PUNTO_COMA)
		}
		localctx.(*ImpresionContext).inst = instrucciones.NewImprimir(localctx.(*ImpresionContext).Get_expresion().GetPrimate(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(rustParserTK_IMPRESION)
		}
		{
			p.SetState(105)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(106)

			var _x = p.expresion(0)

			localctx.(*ImpresionContext)._expresion = _x
		}
		{
			p.SetState(107)

			var _x = p.Listprint()

			localctx.(*ImpresionContext)._listprint = _x
		}
		{
			p.SetState(108)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		{
			p.SetState(109)
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
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rustParserTK_COMA {
		{
			p.SetState(114)

			var _x = p.Expimprimir()

			localctx.(*ListprintContext)._expimprimir = _x
		}
		localctx.(*ListprintContext).list = append(localctx.(*ListprintContext).list, localctx.(*ListprintContext)._expimprimir)

		p.SetState(117)
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
		p.SetState(121)
		p.Match(rustParserTK_COMA)
	}
	{
		p.SetState(122)

		var _x = p.expresion(0)

		localctx.(*ExpimprimirContext)._expresion = _x
	}
	localctx.(*ExpimprimirContext).primate = localctx.(*ExpimprimirContext).Get_expresion().GetPrimate()

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

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(126)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(127)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(128)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(129)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(130)
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
			p.SetState(133)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(134)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(135)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(136)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(137)
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
			p.SetState(140)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(141)
			p.Match(rustParserTK_MUT)
		}
		{
			p.SetState(142)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(143)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(144)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(145)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(146)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(147)
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
			p.SetState(150)
			p.Match(rustParserTK_LET)
		}
		{
			p.SetState(151)

			var _m = p.Match(rustParserTK_ID)

			localctx.(*VariableContext)._TK_ID = _m
		}
		{
			p.SetState(152)
			p.Match(rustParserTK_DOSPUNTOS)
		}
		{
			p.SetState(153)

			var _x = p.Tipo()

			localctx.(*VariableContext)._tipo = _x
		}
		{
			p.SetState(154)
			p.Match(rustParserTK_IGUAL)
		}
		{
			p.SetState(155)

			var _x = p.expresion(0)

			localctx.(*VariableContext)._expresion = _x
		}
		{
			p.SetState(156)
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
	p.EnterRule(localctx, 16, rustParserRULE_asignacionVariable)

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
		p.SetState(161)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*AsignacionVariableContext)._TK_ID = _m
	}
	{
		p.SetState(162)
		p.Match(rustParserTK_IGUAL)
	}
	{
		p.SetState(163)

		var _x = p.expresion(0)

		localctx.(*AsignacionVariableContext)._expresion = _x
	}
	{
		p.SetState(164)
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

// IDeclarcionarregloContext is an interface to support dynamic dispatch.
type IDeclarcionarregloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_ID sets the _TK_ID token.
	Set_TK_ID(antlr.Token)

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_valor returns the _valor rule contexts.
	Get_valor() IValorContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_valor sets the _valor rule contexts.
	Set_valor(IValorContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// IsDeclarcionarregloContext differentiates from other interfaces.
	IsDeclarcionarregloContext()
}

type DeclarcionarregloContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_TK_ID     antlr.Token
	_tipo      ITipoContext
	_valor     IValorContext
	_expresion IExpresionContext
}

func NewEmptyDeclarcionarregloContext() *DeclarcionarregloContext {
	var p = new(DeclarcionarregloContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_declarcionarreglo
	return p
}

func (*DeclarcionarregloContext) IsDeclarcionarregloContext() {}

func NewDeclarcionarregloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarcionarregloContext {
	var p = new(DeclarcionarregloContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_declarcionarreglo

	return p
}

func (s *DeclarcionarregloContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarcionarregloContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *DeclarcionarregloContext) Set_TK_ID(v antlr.Token) { s._TK_ID = v }

func (s *DeclarcionarregloContext) Get_tipo() ITipoContext { return s._tipo }

func (s *DeclarcionarregloContext) Get_valor() IValorContext { return s._valor }

func (s *DeclarcionarregloContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclarcionarregloContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *DeclarcionarregloContext) Set_valor(v IValorContext) { s._valor = v }

func (s *DeclarcionarregloContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclarcionarregloContext) GetInst() interfaces.Instruction { return s.inst }

func (s *DeclarcionarregloContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *DeclarcionarregloContext) TK_LET() antlr.TerminalNode {
	return s.GetToken(rustParserTK_LET, 0)
}

func (s *DeclarcionarregloContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *DeclarcionarregloContext) TK_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(rustParserTK_DOSPUNTOS, 0)
}

func (s *DeclarcionarregloContext) TK_CORCHETE_LETF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_LETF, 0)
}

func (s *DeclarcionarregloContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclarcionarregloContext) AllTK_PUNTO_COMA() []antlr.TerminalNode {
	return s.GetTokens(rustParserTK_PUNTO_COMA)
}

func (s *DeclarcionarregloContext) TK_PUNTO_COMA(i int) antlr.TerminalNode {
	return s.GetToken(rustParserTK_PUNTO_COMA, i)
}

func (s *DeclarcionarregloContext) Valor() IValorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValorContext)
}

func (s *DeclarcionarregloContext) TK_CORCHETE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_RIGHT, 0)
}

func (s *DeclarcionarregloContext) TK_IGUAL() antlr.TerminalNode {
	return s.GetToken(rustParserTK_IGUAL, 0)
}

func (s *DeclarcionarregloContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclarcionarregloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarcionarregloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarcionarregloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterDeclarcionarreglo(s)
	}
}

func (s *DeclarcionarregloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitDeclarcionarreglo(s)
	}
}

func (p *rustParser) Declarcionarreglo() (localctx IDeclarcionarregloContext) {
	this := p
	_ = this

	localctx = NewDeclarcionarregloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, rustParserRULE_declarcionarreglo)

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
		p.SetState(167)
		p.Match(rustParserTK_LET)
	}
	{
		p.SetState(168)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*DeclarcionarregloContext)._TK_ID = _m
	}
	{
		p.SetState(169)
		p.Match(rustParserTK_DOSPUNTOS)
	}
	{
		p.SetState(170)
		p.Match(rustParserTK_CORCHETE_LETF)
	}
	{
		p.SetState(171)

		var _x = p.Tipo()

		localctx.(*DeclarcionarregloContext)._tipo = _x
	}
	{
		p.SetState(172)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	{
		p.SetState(173)

		var _x = p.Valor()

		localctx.(*DeclarcionarregloContext)._valor = _x
	}
	{
		p.SetState(174)
		p.Match(rustParserTK_CORCHETE_RIGHT)
	}
	{
		p.SetState(175)
		p.Match(rustParserTK_IGUAL)
	}
	{
		p.SetState(176)

		var _x = p.expresion(0)

		localctx.(*DeclarcionarregloContext)._expresion = _x
	}
	{
		p.SetState(177)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*DeclarcionarregloContext).inst = arreglos.NewArreglo((func() string {
		if localctx.(*DeclarcionarregloContext).Get_TK_ID() == nil {
			return ""
		} else {
			return localctx.(*DeclarcionarregloContext).Get_TK_ID().GetText()
		}
	}()), false, localctx.(*DeclarcionarregloContext).Get_tipo().GetTipoExp(), (func() string {
		if localctx.(*DeclarcionarregloContext).Get_valor() == nil {
			return ""
		} else {
			return p.GetTokenStream().GetTextFromTokens(localctx.(*DeclarcionarregloContext).Get_valor().GetStart(), localctx.(*DeclarcionarregloContext)._valor.GetStop())
		}
	}()), true, localctx.(*DeclarcionarregloContext).Get_expresion().GetPrimate())

	return localctx
}

// IListvaloresContext is an interface to support dynamic dispatch.
type IListvaloresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IListvaloresContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetL sets the l rule contexts.
	SetL(IListvaloresContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsListvaloresContext differentiates from other interfaces.
	IsListvaloresContext()
}

type ListvaloresContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	lista      *arrayList.List
	l          IListvaloresContext
	_expresion IExpresionContext
}

func NewEmptyListvaloresContext() *ListvaloresContext {
	var p = new(ListvaloresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_listvalores
	return p
}

func (*ListvaloresContext) IsListvaloresContext() {}

func NewListvaloresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListvaloresContext {
	var p = new(ListvaloresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_listvalores

	return p
}

func (s *ListvaloresContext) GetParser() antlr.Parser { return s.parser }

func (s *ListvaloresContext) GetL() IListvaloresContext { return s.l }

func (s *ListvaloresContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ListvaloresContext) SetL(v IListvaloresContext) { s.l = v }

func (s *ListvaloresContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ListvaloresContext) GetLista() *arrayList.List { return s.lista }

func (s *ListvaloresContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *ListvaloresContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ListvaloresContext) TK_COMA() antlr.TerminalNode {
	return s.GetToken(rustParserTK_COMA, 0)
}

func (s *ListvaloresContext) Listvalores() IListvaloresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListvaloresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListvaloresContext)
}

func (s *ListvaloresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListvaloresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListvaloresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterListvalores(s)
	}
}

func (s *ListvaloresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitListvalores(s)
	}
}

func (p *rustParser) Listvalores() (localctx IListvaloresContext) {
	return p.listvalores(0)
}

func (p *rustParser) listvalores(_p int) (localctx IListvaloresContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListvaloresContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListvaloresContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, rustParserRULE_listvalores, _p)

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
		p.SetState(181)

		var _x = p.expresion(0)

		localctx.(*ListvaloresContext)._expresion = _x
	}

	localctx.(*ListvaloresContext).lista = arrayList.New()
	localctx.(*ListvaloresContext).lista.Add(localctx.(*ListvaloresContext).Get_expresion().GetPrimate())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListvaloresContext(p, _parentctx, _parentState)
			localctx.(*ListvaloresContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, rustParserRULE_listvalores)
			p.SetState(184)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(185)
				p.Match(rustParserTK_COMA)
			}
			{
				p.SetState(186)

				var _x = p.expresion(0)

				localctx.(*ListvaloresContext)._expresion = _x
			}

			localctx.(*ListvaloresContext).GetL().GetLista().Add(localctx.(*ListvaloresContext).Get_expresion().GetPrimate())
			localctx.(*ListvaloresContext).lista = localctx.(*ListvaloresContext).GetL().GetLista()

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

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
	p.EnterRule(localctx, 22, rustParserRULE_expresionIf)

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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(195)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(196)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(197)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(198)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(202)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(203)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(204)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueif = _x
		}
		{
			p.SetState(205)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(206)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(207)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(208)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(209)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetBloqueif().GetLista(), localctx.(*ExpresionIfContext).GetBloqueelse().GetLista())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(213)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(214)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(215)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(216)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(217)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		localctx.(*ExpresionIfContext).inst = instrucciones.NewIf2(localctx.(*ExpresionIfContext).Get_expresion().GetPrimate(), localctx.(*ExpresionIfContext).GetIfblock().GetLista(), nil, localctx.(*ExpresionIfContext).Get_listaelif().GetLista())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(rustParserTK_IF)
		}
		{
			p.SetState(221)

			var _x = p.expresion(0)

			localctx.(*ExpresionIfContext)._expresion = _x
		}
		{
			p.SetState(222)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(223)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).ifblock = _x
		}
		{
			p.SetState(224)
			p.Match(rustParserTK_LLAVE_RIGHT)
		}
		{
			p.SetState(225)

			var _x = p.Listaelif()

			localctx.(*ExpresionIfContext)._listaelif = _x
		}
		{
			p.SetState(226)
			p.Match(rustParserTK_ELSE)
		}
		{
			p.SetState(227)
			p.Match(rustParserTK_LLAVE_LEFT)
		}
		{
			p.SetState(228)

			var _x = p.Instrucciones()

			localctx.(*ExpresionIfContext).bloqueelse = _x
		}
		{
			p.SetState(229)
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
	p.EnterRule(localctx, 24, rustParserRULE_listaelif)
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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == rustParserTK_ELSE_IF {
		{
			p.SetState(234)

			var _x = p.Elif()

			localctx.(*ListaelifContext)._elif = _x
		}
		localctx.(*ListaelifContext).l = append(localctx.(*ListaelifContext).l, localctx.(*ListaelifContext)._elif)

		p.SetState(237)
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
	p.EnterRule(localctx, 26, rustParserRULE_elif)

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
		p.SetState(241)
		p.Match(rustParserTK_ELSE_IF)
	}
	{
		p.SetState(242)

		var _x = p.expresion(0)

		localctx.(*ElifContext)._expresion = _x
	}
	{
		p.SetState(243)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(244)

		var _x = p.Instrucciones()

		localctx.(*ElifContext).elifblock = _x
	}
	{
		p.SetState(245)
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
	p.EnterRule(localctx, 28, rustParserRULE_expresionWhile)

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
		p.SetState(248)
		p.Match(rustParserTK_WHILE)
	}
	{
		p.SetState(249)

		var _x = p.expresion(0)

		localctx.(*ExpresionWhileContext).exp = _x
	}
	{
		p.SetState(250)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(251)

		var _x = p.Instrucciones()

		localctx.(*ExpresionWhileContext)._instrucciones = _x
	}
	{
		p.SetState(252)
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
	p.EnterRule(localctx, 30, rustParserRULE_expresionLoop)

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
		p.SetState(255)
		p.Match(rustParserTK_LOOP)
	}
	{
		p.SetState(256)
		p.Match(rustParserTK_LLAVE_LEFT)
	}
	{
		p.SetState(257)

		var _x = p.Instrucciones()

		localctx.(*ExpresionLoopContext)._instrucciones = _x
	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 32, rustParserRULE_breakInst)

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
		p.SetState(261)
		p.Match(rustParserTK_BREAK)
	}
	{
		p.SetState(262)
		p.Match(rustParserTK_PUNTO_COMA)
	}
	localctx.(*BreakInstContext).inst = instrucciones.NewBreak(interfaces.BREAK)

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
	p.EnterRule(localctx, 34, rustParserRULE_tipo)

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

	p.SetState(275)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case rustParserTK_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Match(rustParserTK_INT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.INTEGER

	case rustParserTK_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(rustParserTK_FLOAT)
		}
		localctx.(*TipoContext).tipoExp = interfaces.FLOAT

	case rustParserTK_BOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(269)
			p.Match(rustParserTK_BOOL)
		}
		localctx.(*TipoContext).tipoExp = interfaces.BOOLEAN

	case rustParserTK_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(271)
			p.Match(rustParserTK_CHAR)
		}
		localctx.(*TipoContext).tipoExp = interfaces.CHAR

	case rustParserTK_STRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(273)
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

	// Get_listvalores returns the _listvalores rule contexts.
	Get_listvalores() IListvaloresContext

	// Get_valor returns the _valor rule contexts.
	Get_valor() IValorContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExpresionContext)

	// SetRight sets the right rule contexts.
	SetRight(IExpresionContext)

	// Set_listvalores sets the _listvalores rule contexts.
	Set_listvalores(IListvaloresContext)

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
	parser       antlr.Parser
	primate      interfaces.Expresion
	left         IExpresionContext
	op           antlr.Token
	right        IExpresionContext
	_TK_POW_I64  antlr.Token
	_TK_POW_F64  antlr.Token
	_listvalores IListvaloresContext
	_valor       IValorContext
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

func (s *ExpresionContext) Get_listvalores() IListvaloresContext { return s._listvalores }

func (s *ExpresionContext) Get_valor() IValorContext { return s._valor }

func (s *ExpresionContext) SetLeft(v IExpresionContext) { s.left = v }

func (s *ExpresionContext) SetRight(v IExpresionContext) { s.right = v }

func (s *ExpresionContext) Set_listvalores(v IListvaloresContext) { s._listvalores = v }

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

func (s *ExpresionContext) TK_CORCHETE_LETF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_LETF, 0)
}

func (s *ExpresionContext) Listvalores() IListvaloresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListvaloresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListvaloresContext)
}

func (s *ExpresionContext) TK_CORCHETE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_RIGHT, 0)
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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, rustParserRULE_expresion, _p)
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
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(278)

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
			p.SetState(279)

			var _x = p.expresion(17)

			localctx.(*ExpresionContext).right = _x
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(nil, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case 2:
		{
			p.SetState(282)

			var _m = p.Match(rustParserTK_POW_I64)

			localctx.(*ExpresionContext)._TK_POW_I64 = _m
		}
		{
			p.SetState(283)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(284)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).left = _x
		}
		{
			p.SetState(285)
			p.Match(rustParserTK_COMA)
		}
		{
			p.SetState(286)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).right = _x
		}
		{
			p.SetState(287)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
			if localctx.(*ExpresionContext).Get_TK_POW_I64() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TK_POW_I64().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case 3:
		{
			p.SetState(290)

			var _m = p.Match(rustParserTK_POW_F64)

			localctx.(*ExpresionContext)._TK_POW_F64 = _m
		}
		{
			p.SetState(291)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(292)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).left = _x
		}
		{
			p.SetState(293)
			p.Match(rustParserTK_COMA)
		}
		{
			p.SetState(294)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).right = _x
		}
		{
			p.SetState(295)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = expressiones.NewOperacion(localctx.(*ExpresionContext).GetLeft().GetPrimate(), (func() string {
			if localctx.(*ExpresionContext).Get_TK_POW_F64() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TK_POW_F64().GetText()
			}
		}()), localctx.(*ExpresionContext).GetRight().GetPrimate())

	case 4:
		{
			p.SetState(298)
			p.Match(rustParserTK_PARENTESIS_LEFT)
		}
		{
			p.SetState(299)
			p.expresion(0)
		}
		{
			p.SetState(300)
			p.Match(rustParserTK_PARENTESIS_RIGHT)
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	case 5:
		{
			p.SetState(303)
			p.Match(rustParserTK_CORCHETE_LETF)
		}
		{
			p.SetState(304)

			var _x = p.listvalores(0)

			localctx.(*ExpresionContext)._listvalores = _x
		}
		{
			p.SetState(305)
			p.Match(rustParserTK_CORCHETE_RIGHT)
		}
		localctx.(*ExpresionContext).primate = arreglos.NewArray(localctx.(*ExpresionContext).Get_listvalores().GetLista())

	case 6:
		{
			p.SetState(308)

			var _x = p.Valor()

			localctx.(*ExpresionContext)._valor = _x
		}
		localctx.(*ExpresionContext).primate = localctx.(*ExpresionContext).Get_valor().GetPrimate()

	case 7:

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(364)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(315)

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
					p.SetState(316)

					var _x = p.expresion(17)

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
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(320)

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
					p.SetState(321)

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

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, rustParserRULE_expresion)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(325)

					var _m = p.Match(rustParserTK_IGUALIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(326)

					var _x = p.expresion(13)

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
				p.SetState(329)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(330)

					var _m = p.Match(rustParserTK_DIFERENTE)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(331)

					var _x = p.expresion(12)

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
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(335)

					var _m = p.Match(rustParserTK_MAYOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(336)

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
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(340)

					var _m = p.Match(rustParserTK_MENOR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(341)

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
				p.SetState(344)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(345)

					var _m = p.Match(rustParserTK_MAYORIGULA)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(346)

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
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(350)

					var _m = p.Match(rustParserTK_MENOIGUAL)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(351)

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
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(355)

					var _m = p.Match(rustParserTK_AND)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(356)

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
				p.SetState(359)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(360)

					var _m = p.Match(rustParserTK_OR)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(361)

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

			}

		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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

	// GetLista returns the lista rule contexts.
	GetLista() IListaArregloContext

	// SetLista sets the lista rule contexts.
	SetLista(IListaArregloContext)

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
	lista        IListaArregloContext
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

func (s *ValorContext) GetLista() IListaArregloContext { return s.lista }

func (s *ValorContext) SetLista(v IListaArregloContext) { s.lista = v }

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

func (s *ValorContext) ListaArreglo() IListaArregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaArregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaArregloContext)
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
	p.EnterRule(localctx, 38, rustParserRULE_valor)

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

	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)

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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)

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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)

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

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(375)

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

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(377)

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

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(379)

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

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(381)

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

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(383)

			var _x = p.listaArreglo(0)

			localctx.(*ValorContext).lista = _x
		}
		localctx.(*ValorContext).primate = localctx.(*ValorContext).GetLista().GetPrimate()

	}

	return localctx
}

// IListaArregloContext is an interface to support dynamic dispatch.
type IListaArregloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_TK_ID returns the _TK_ID token.
	Get_TK_ID() antlr.Token

	// Set_TK_ID sets the _TK_ID token.
	Set_TK_ID(antlr.Token)

	// GetLista returns the lista rule contexts.
	GetLista() IListaArregloContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetLista sets the lista rule contexts.
	SetLista(IListaArregloContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetPrimate returns the primate attribute.
	GetPrimate() interfaces.Expresion

	// SetPrimate sets the primate attribute.
	SetPrimate(interfaces.Expresion)

	// IsListaArregloContext differentiates from other interfaces.
	IsListaArregloContext()
}

type ListaArregloContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	primate    interfaces.Expresion
	lista      IListaArregloContext
	_TK_ID     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyListaArregloContext() *ListaArregloContext {
	var p = new(ListaArregloContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = rustParserRULE_listaArreglo
	return p
}

func (*ListaArregloContext) IsListaArregloContext() {}

func NewListaArregloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListaArregloContext {
	var p = new(ListaArregloContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = rustParserRULE_listaArreglo

	return p
}

func (s *ListaArregloContext) GetParser() antlr.Parser { return s.parser }

func (s *ListaArregloContext) Get_TK_ID() antlr.Token { return s._TK_ID }

func (s *ListaArregloContext) Set_TK_ID(v antlr.Token) { s._TK_ID = v }

func (s *ListaArregloContext) GetLista() IListaArregloContext { return s.lista }

func (s *ListaArregloContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ListaArregloContext) SetLista(v IListaArregloContext) { s.lista = v }

func (s *ListaArregloContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ListaArregloContext) GetPrimate() interfaces.Expresion { return s.primate }

func (s *ListaArregloContext) SetPrimate(v interfaces.Expresion) { s.primate = v }

func (s *ListaArregloContext) TK_ID() antlr.TerminalNode {
	return s.GetToken(rustParserTK_ID, 0)
}

func (s *ListaArregloContext) TK_CORCHETE_LETF() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_LETF, 0)
}

func (s *ListaArregloContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ListaArregloContext) TK_CORCHETE_RIGHT() antlr.TerminalNode {
	return s.GetToken(rustParserTK_CORCHETE_RIGHT, 0)
}

func (s *ListaArregloContext) ListaArreglo() IListaArregloContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListaArregloContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListaArregloContext)
}

func (s *ListaArregloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaArregloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListaArregloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.EnterListaArreglo(s)
	}
}

func (s *ListaArregloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(rustListener); ok {
		listenerT.ExitListaArreglo(s)
	}
}

func (p *rustParser) ListaArreglo() (localctx IListaArregloContext) {
	return p.listaArreglo(0)
}

func (p *rustParser) listaArreglo(_p int) (localctx IListaArregloContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListaArregloContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListaArregloContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, rustParserRULE_listaArreglo, _p)

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
		p.SetState(389)

		var _m = p.Match(rustParserTK_ID)

		localctx.(*ListaArregloContext)._TK_ID = _m
	}

	localctx.(*ListaArregloContext).primate = expressiones.NewLLamadoVariable((func() string {
		if localctx.(*ListaArregloContext).Get_TK_ID() == nil {
			return ""
		} else {
			return localctx.(*ListaArregloContext).Get_TK_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListaArregloContext(p, _parentctx, _parentState)
			localctx.(*ListaArregloContext).lista = _prevctx
			p.PushNewRecursionContext(localctx, _startState, rustParserRULE_listaArreglo)
			p.SetState(392)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(393)
				p.Match(rustParserTK_CORCHETE_LETF)
			}
			{
				p.SetState(394)

				var _x = p.expresion(0)

				localctx.(*ListaArregloContext)._expresion = _x
			}
			{
				p.SetState(395)
				p.Match(rustParserTK_CORCHETE_RIGHT)
			}

			localctx.(*ListaArregloContext).primate = arreglos.NewArrayAccess(localctx.(*ListaArregloContext).GetLista().GetPrimate(), localctx.(*ListaArregloContext).Get_expresion().GetPrimate())

		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

func (p *rustParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ListvaloresContext = nil
		if localctx != nil {
			t = localctx.(*ListvaloresContext)
		}
		return p.Listvalores_Sempred(t, predIndex)

	case 18:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	case 20:
		var t *ListaArregloContext = nil
		if localctx != nil {
			t = localctx.(*ListaArregloContext)
		}
		return p.ListaArreglo_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *rustParser) Listvalores_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *rustParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *rustParser) ListaArreglo_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
