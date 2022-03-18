// Code generated from /Users/mickey/git/MySQLParser-go/parser/antlr4/MySQL.g4 by ANTLR 4.9.2. DO NOT EDIT.

package antlr4 // MySQL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 112, 583, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44, 
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4, 
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55, 
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9, 
	60, 4, 61, 9, 61, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 128, 10, 2, 3, 3, 
	3, 3, 5, 3, 132, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 140, 
	10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 147, 10, 6, 3, 6, 3, 6, 3, 6, 
	3, 6, 3, 6, 3, 6, 5, 6, 155, 10, 6, 3, 7, 3, 7, 7, 7, 159, 10, 7, 12, 7, 
	14, 7, 162, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 169, 10, 8, 12, 
	8, 14, 8, 172, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 181, 
	10, 9, 3, 9, 5, 9, 184, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 
	191, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11, 196, 10, 11, 3, 11, 3, 11, 5, 
	11, 200, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11, 205, 10, 11, 3, 11, 3, 11, 
	5, 11, 209, 10, 11, 3, 12, 3, 12, 3, 12, 5, 12, 214, 10, 12, 3, 12, 3, 
	12, 3, 12, 5, 12, 219, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 225, 
	10, 12, 5, 12, 227, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 
	12, 235, 10, 12, 3, 13, 3, 13, 5, 13, 239, 10, 13, 3, 13, 3, 13, 3, 13, 
	3, 13, 3, 13, 5, 13, 246, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 252, 
	10, 14, 3, 14, 5, 14, 255, 10, 14, 3, 14, 3, 14, 5, 14, 259, 10, 14, 3, 
	15, 3, 15, 3, 15, 7, 15, 264, 10, 15, 12, 15, 14, 15, 267, 11, 15, 3, 16, 
	3, 16, 5, 16, 271, 10, 16, 3, 17, 3, 17, 3, 17, 5, 17, 276, 10, 17, 3, 
	18, 3, 18, 3, 18, 3, 18, 5, 18, 282, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 
	21, 3, 21, 3, 21, 5, 21, 301, 10, 21, 3, 21, 5, 21, 304, 10, 21, 3, 21, 
	5, 21, 307, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 313, 10, 22, 3, 
	22, 3, 22, 5, 22, 317, 10, 22, 5, 22, 319, 10, 22, 3, 23, 3, 23, 3, 23, 
	3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 328, 10, 23, 3, 24, 3, 24, 5, 24, 332, 
	10, 24, 3, 24, 3, 24, 5, 24, 336, 10, 24, 3, 25, 3, 25, 5, 25, 340, 10, 
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 348, 10, 26, 3, 26, 
	3, 26, 5, 26, 352, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 
	27, 360, 10, 27, 3, 28, 3, 28, 3, 28, 7, 28, 365, 10, 28, 12, 28, 14, 28, 
	368, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 374, 10, 29, 3, 30, 3, 
	30, 3, 30, 3, 30, 3, 30, 5, 30, 381, 10, 30, 3, 30, 3, 30, 5, 30, 385, 
	10, 30, 3, 31, 3, 31, 5, 31, 389, 10, 31, 3, 31, 5, 31, 392, 10, 31, 3, 
	32, 3, 32, 3, 32, 7, 32, 397, 10, 32, 12, 32, 14, 32, 400, 11, 32, 3, 33, 
	3, 33, 5, 33, 404, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 411, 
	10, 34, 3, 35, 3, 35, 3, 35, 5, 35, 416, 10, 35, 3, 36, 3, 36, 3, 36, 3, 
	36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 426, 10, 36, 3, 37, 3, 37, 3, 37, 
	3, 37, 3, 38, 3, 38, 5, 38, 434, 10, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 
	38, 3, 39, 3, 39, 3, 39, 5, 39, 444, 10, 39, 3, 39, 3, 39, 3, 40, 3, 40, 
	5, 40, 450, 10, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 5, 
	41, 459, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 5, 42, 467, 
	10, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 
	3, 44, 5, 44, 479, 10, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 5, 
	45, 487, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 
	3, 46, 5, 46, 498, 10, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 
	49, 3, 50, 5, 50, 508, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 
	3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 7, 53, 525, 
	10, 53, 12, 53, 14, 53, 528, 11, 53, 3, 54, 3, 54, 5, 54, 532, 10, 54, 
	3, 55, 5, 55, 535, 10, 55, 3, 55, 3, 55, 3, 56, 3, 56, 5, 56, 541, 10, 
	56, 3, 56, 3, 56, 3, 56, 5, 56, 546, 10, 56, 3, 56, 3, 56, 3, 57, 3, 57, 
	3, 57, 5, 57, 553, 10, 57, 3, 57, 3, 57, 3, 57, 5, 57, 558, 10, 57, 3, 
	58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 
	5, 60, 571, 10, 60, 3, 60, 3, 60, 3, 61, 3, 61, 5, 61, 577, 10, 61, 3, 
	61, 3, 61, 5, 61, 581, 10, 61, 3, 61, 2, 2, 62, 2, 4, 6, 8, 10, 12, 14, 
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 
	88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 
	120, 2, 14, 4, 2, 4, 4, 76, 76, 4, 2, 37, 37, 46, 46, 3, 2, 39, 40, 4, 
	2, 20, 20, 63, 64, 3, 2, 90, 95, 5, 2, 13, 13, 18, 19, 50, 50, 4, 2, 3, 
	3, 65, 65, 7, 2, 4, 4, 13, 13, 50, 50, 74, 74, 79, 80, 4, 2, 18, 19, 76, 
	78, 3, 2, 43, 45, 3, 2, 46, 48, 4, 2, 30, 30, 66, 75, 2, 615, 2, 127, 3, 
	2, 2, 2, 4, 131, 3, 2, 2, 2, 6, 133, 3, 2, 2, 2, 8, 135, 3, 2, 2, 2, 10, 
	137, 3, 2, 2, 2, 12, 160, 3, 2, 2, 2, 14, 170, 3, 2, 2, 2, 16, 180, 3, 
	2, 2, 2, 18, 185, 3, 2, 2, 2, 20, 188, 3, 2, 2, 2, 22, 213, 3, 2, 2, 2, 
	24, 236, 3, 2, 2, 2, 26, 249, 3, 2, 2, 2, 28, 260, 3, 2, 2, 2, 30, 270, 
	3, 2, 2, 2, 32, 275, 3, 2, 2, 2, 34, 277, 3, 2, 2, 2, 36, 285, 3, 2, 2, 
	2, 38, 289, 3, 2, 2, 2, 40, 292, 3, 2, 2, 2, 42, 318, 3, 2, 2, 2, 44, 327, 
	3, 2, 2, 2, 46, 329, 3, 2, 2, 2, 48, 339, 3, 2, 2, 2, 50, 341, 3, 2, 2, 
	2, 52, 353, 3, 2, 2, 2, 54, 361, 3, 2, 2, 2, 56, 369, 3, 2, 2, 2, 58, 375, 
	3, 2, 2, 2, 60, 386, 3, 2, 2, 2, 62, 393, 3, 2, 2, 2, 64, 403, 3, 2, 2, 
	2, 66, 405, 3, 2, 2, 2, 68, 412, 3, 2, 2, 2, 70, 425, 3, 2, 2, 2, 72, 427, 
	3, 2, 2, 2, 74, 431, 3, 2, 2, 2, 76, 440, 3, 2, 2, 2, 78, 447, 3, 2, 2, 
	2, 80, 456, 3, 2, 2, 2, 82, 466, 3, 2, 2, 2, 84, 473, 3, 2, 2, 2, 86, 476, 
	3, 2, 2, 2, 88, 486, 3, 2, 2, 2, 90, 497, 3, 2, 2, 2, 92, 499, 3, 2, 2, 
	2, 94, 501, 3, 2, 2, 2, 96, 503, 3, 2, 2, 2, 98, 507, 3, 2, 2, 2, 100, 
	513, 3, 2, 2, 2, 102, 517, 3, 2, 2, 2, 104, 521, 3, 2, 2, 2, 106, 529, 
	3, 2, 2, 2, 108, 534, 3, 2, 2, 2, 110, 538, 3, 2, 2, 2, 112, 549, 3, 2, 
	2, 2, 114, 559, 3, 2, 2, 2, 116, 562, 3, 2, 2, 2, 118, 567, 3, 2, 2, 2, 
	120, 576, 3, 2, 2, 2, 122, 128, 5, 10, 6, 2, 123, 128, 5, 48, 25, 2, 124, 
	128, 5, 58, 30, 2, 125, 128, 5, 16, 9, 2, 126, 128, 5, 4, 3, 2, 127, 122, 
	3, 2, 2, 2, 127, 123, 3, 2, 2, 2, 127, 124, 3, 2, 2, 2, 127, 125, 3, 2, 
	2, 2, 127, 126, 3, 2, 2, 2, 128, 3, 3, 2, 2, 2, 129, 132, 5, 6, 4, 2, 130, 
	132, 5, 8, 5, 2, 131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 5, 3, 
	2, 2, 2, 133, 134, 7, 54, 2, 2, 134, 7, 3, 2, 2, 2, 135, 136, 7, 55, 2, 
	2, 136, 9, 3, 2, 2, 2, 137, 139, 7, 6, 2, 2, 138, 140, 7, 7, 2, 2, 139, 
	138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 146, 
	7, 79, 2, 2, 142, 143, 7, 85, 2, 2, 143, 144, 5, 12, 7, 2, 144, 145, 7, 
	84, 2, 2, 145, 147, 3, 2, 2, 2, 146, 142, 3, 2, 2, 2, 146, 147, 3, 2, 2, 
	2, 147, 154, 3, 2, 2, 2, 148, 149, 7, 8, 2, 2, 149, 150, 7, 85, 2, 2, 150, 
	151, 5, 14, 8, 2, 151, 152, 7, 84, 2, 2, 152, 155, 3, 2, 2, 2, 153, 155, 
	5, 16, 9, 2, 154, 148, 3, 2, 2, 2, 154, 153, 3, 2, 2, 2, 155, 11, 3, 2, 
	2, 2, 156, 157, 7, 79, 2, 2, 157, 159, 7, 97, 2, 2, 158, 156, 3, 2, 2, 
	2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 
	163, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164, 7, 79, 2, 2, 164, 13, 
	3, 2, 2, 2, 165, 166, 5, 88, 45, 2, 166, 167, 7, 97, 2, 2, 167, 169, 3, 
	2, 2, 2, 168, 165, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 
	2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 
	174, 5, 88, 45, 2, 174, 15, 3, 2, 2, 2, 175, 176, 7, 85, 2, 2, 176, 177, 
	5, 18, 10, 2, 177, 178, 7, 84, 2, 2, 178, 181, 3, 2, 2, 2, 179, 181, 5, 
	18, 10, 2, 180, 175, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 183, 3, 2, 
	2, 2, 182, 184, 5, 24, 13, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3, 2, 2, 
	2, 184, 17, 3, 2, 2, 2, 185, 186, 5, 20, 11, 2, 186, 187, 5, 22, 12, 2, 
	187, 19, 3, 2, 2, 2, 188, 190, 7, 5, 2, 2, 189, 191, 7, 37, 2, 2, 190, 
	189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 195, 
	5, 26, 14, 2, 193, 194, 7, 10, 2, 2, 194, 196, 5, 28, 15, 2, 195, 193, 
	3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 198, 7, 11, 
	2, 2, 198, 200, 5, 64, 33, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 
	2, 200, 204, 3, 2, 2, 2, 201, 202, 7, 25, 2, 2, 202, 203, 7, 26, 2, 2, 
	203, 205, 5, 46, 24, 2, 204, 201, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 
	208, 3, 2, 2, 2, 206, 207, 7, 29, 2, 2, 207, 209, 5, 64, 33, 2, 208, 206, 
	3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 21, 3, 2, 2, 2, 210, 211, 7, 24, 
	2, 2, 211, 212, 7, 26, 2, 2, 212, 214, 5, 46, 24, 2, 213, 210, 3, 2, 2, 
	2, 213, 214, 3, 2, 2, 2, 214, 226, 3, 2, 2, 2, 215, 224, 7, 12, 2, 2, 216, 
	217, 9, 2, 2, 2, 217, 219, 7, 97, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 
	3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 225, 9, 2, 2, 2, 221, 222, 9, 2, 
	2, 2, 222, 223, 7, 38, 2, 2, 223, 225, 9, 2, 2, 2, 224, 218, 3, 2, 2, 2, 
	224, 221, 3, 2, 2, 2, 225, 227, 3, 2, 2, 2, 226, 215, 3, 2, 2, 2, 226, 
	227, 3, 2, 2, 2, 227, 234, 3, 2, 2, 2, 228, 229, 7, 27, 2, 2, 229, 235, 
	7, 22, 2, 2, 230, 231, 7, 51, 2, 2, 231, 232, 7, 15, 2, 2, 232, 233, 7, 
	52, 2, 2, 233, 235, 7, 53, 2, 2, 234, 228, 3, 2, 2, 2, 234, 230, 3, 2, 
	2, 2, 234, 235, 3, 2, 2, 2, 235, 23, 3, 2, 2, 2, 236, 238, 7, 49, 2, 2, 
	237, 239, 9, 3, 2, 2, 238, 237, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 
	245, 3, 2, 2, 2, 240, 241, 7, 85, 2, 2, 241, 242, 5, 16, 9, 2, 242, 243, 
	7, 84, 2, 2, 243, 246, 3, 2, 2, 2, 244, 246, 5, 16, 9, 2, 245, 240, 3, 
	2, 2, 2, 245, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 5, 22, 12, 
	2, 248, 25, 3, 2, 2, 2, 249, 254, 5, 88, 45, 2, 250, 252, 7, 30, 2, 2, 
	251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 
	255, 7, 79, 2, 2, 254, 251, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 258, 
	3, 2, 2, 2, 256, 257, 7, 97, 2, 2, 257, 259, 5, 26, 14, 2, 258, 256, 3, 
	2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 27, 3, 2, 2, 2, 260, 265, 5, 30, 16, 
	2, 261, 262, 7, 97, 2, 2, 262, 264, 5, 30, 16, 2, 263, 261, 3, 2, 2, 2, 
	264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 
	29, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268, 271, 5, 32, 17, 2, 269, 271, 
	5, 38, 20, 2, 270, 268, 3, 2, 2, 2, 270, 269, 3, 2, 2, 2, 271, 31, 3, 2, 
	2, 2, 272, 276, 5, 60, 31, 2, 273, 276, 5, 34, 18, 2, 274, 276, 5, 36, 
	19, 2, 275, 272, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275, 274, 3, 2, 2, 2, 
	276, 33, 3, 2, 2, 2, 277, 278, 7, 85, 2, 2, 278, 279, 5, 16, 9, 2, 279, 
	281, 7, 84, 2, 2, 280, 282, 7, 30, 2, 2, 281, 280, 3, 2, 2, 2, 281, 282, 
	3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284, 7, 79, 2, 2, 284, 35, 3, 2, 
	2, 2, 285, 286, 7, 85, 2, 2, 286, 287, 5, 30, 16, 2, 287, 288, 7, 84, 2, 
	2, 288, 37, 3, 2, 2, 2, 289, 290, 5, 60, 31, 2, 290, 291, 5, 40, 21, 2, 
	291, 39, 3, 2, 2, 2, 292, 293, 5, 42, 22, 2, 293, 300, 7, 33, 2, 2, 294, 
	301, 5, 62, 32, 2, 295, 296, 7, 85, 2, 2, 296, 297, 5, 62, 32, 2, 297, 
	298, 7, 84, 2, 2, 298, 301, 3, 2, 2, 2, 299, 301, 5, 36, 19, 2, 300, 294, 
	3, 2, 2, 2, 300, 295, 3, 2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 303, 3, 2, 
	2, 2, 302, 304, 5, 44, 23, 2, 303, 302, 3, 2, 2, 2, 303, 304, 3, 2, 2, 
	2, 304, 306, 3, 2, 2, 2, 305, 307, 5, 40, 21, 2, 306, 305, 3, 2, 2, 2, 
	306, 307, 3, 2, 2, 2, 307, 41, 3, 2, 2, 2, 308, 319, 7, 31, 2, 2, 309, 
	319, 7, 41, 2, 2, 310, 312, 7, 34, 2, 2, 311, 313, 7, 32, 2, 2, 312, 311, 
	3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 319, 3, 2, 2, 2, 314, 316, 7, 35, 
	2, 2, 315, 317, 7, 32, 2, 2, 316, 315, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 
	317, 319, 3, 2, 2, 2, 318, 308, 3, 2, 2, 2, 318, 309, 3, 2, 2, 2, 318, 
	310, 3, 2, 2, 2, 318, 314, 3, 2, 2, 2, 319, 43, 3, 2, 2, 2, 320, 321, 7, 
	36, 2, 2, 321, 328, 5, 64, 33, 2, 322, 323, 7, 42, 2, 2, 323, 324, 7, 85, 
	2, 2, 324, 325, 5, 12, 7, 2, 325, 326, 7, 84, 2, 2, 326, 328, 3, 2, 2, 
	2, 327, 320, 3, 2, 2, 2, 327, 322, 3, 2, 2, 2, 328, 45, 3, 2, 2, 2, 329, 
	331, 5, 88, 45, 2, 330, 332, 9, 4, 2, 2, 331, 330, 3, 2, 2, 2, 331, 332, 
	3, 2, 2, 2, 332, 335, 3, 2, 2, 2, 333, 334, 7, 97, 2, 2, 334, 336, 5, 46, 
	24, 2, 335, 333, 3, 2, 2, 2, 335, 336, 3, 2, 2, 2, 336, 47, 3, 2, 2, 2, 
	337, 340, 5, 50, 26, 2, 338, 340, 5, 52, 27, 2, 339, 337, 3, 2, 2, 2, 339, 
	338, 3, 2, 2, 2, 340, 49, 3, 2, 2, 2, 341, 342, 7, 22, 2, 2, 342, 343, 
	5, 60, 31, 2, 343, 344, 7, 23, 2, 2, 344, 347, 5, 54, 28, 2, 345, 346, 
	7, 11, 2, 2, 346, 348, 5, 64, 33, 2, 347, 345, 3, 2, 2, 2, 347, 348, 3, 
	2, 2, 2, 348, 351, 3, 2, 2, 2, 349, 350, 7, 12, 2, 2, 350, 352, 9, 2, 2, 
	2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 51, 3, 2, 2, 2, 353, 
	354, 7, 22, 2, 2, 354, 355, 5, 62, 32, 2, 355, 356, 7, 23, 2, 2, 356, 359, 
	5, 54, 28, 2, 357, 358, 7, 11, 2, 2, 358, 360, 5, 64, 33, 2, 359, 357, 
	3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 53, 3, 2, 2, 2, 361, 366, 5, 56, 
	29, 2, 362, 363, 7, 97, 2, 2, 363, 365, 5, 56, 29, 2, 364, 362, 3, 2, 2, 
	2, 365, 368, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 
	55, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 369, 370, 5, 88, 45, 2, 370, 373, 
	7, 90, 2, 2, 371, 374, 5, 88, 45, 2, 372, 374, 7, 21, 2, 2, 373, 371, 3, 
	2, 2, 2, 373, 372, 3, 2, 2, 2, 374, 57, 3, 2, 2, 2, 375, 376, 7, 9, 2, 
	2, 376, 377, 7, 10, 2, 2, 377, 380, 5, 60, 31, 2, 378, 379, 7, 11, 2, 2, 
	379, 381, 5, 64, 33, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 
	384, 3, 2, 2, 2, 382, 383, 7, 12, 2, 2, 383, 385, 9, 2, 2, 2, 384, 382, 
	3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 59, 3, 2, 2, 2, 386, 391, 7, 79, 
	2, 2, 387, 389, 7, 30, 2, 2, 388, 387, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 
	389, 390, 3, 2, 2, 2, 390, 392, 7, 79, 2, 2, 391, 388, 3, 2, 2, 2, 391, 
	392, 3, 2, 2, 2, 392, 61, 3, 2, 2, 2, 393, 398, 5, 60, 31, 2, 394, 395, 
	7, 97, 2, 2, 395, 397, 5, 60, 31, 2, 396, 394, 3, 2, 2, 2, 397, 400, 3, 
	2, 2, 2, 398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399, 63, 3, 2, 2, 
	2, 400, 398, 3, 2, 2, 2, 401, 404, 5, 66, 34, 2, 402, 404, 5, 68, 35, 2, 
	403, 401, 3, 2, 2, 2, 403, 402, 3, 2, 2, 2, 404, 65, 3, 2, 2, 2, 405, 406, 
	7, 85, 2, 2, 406, 407, 5, 64, 33, 2, 407, 410, 7, 84, 2, 2, 408, 409, 9, 
	5, 2, 2, 409, 411, 5, 64, 33, 2, 410, 408, 3, 2, 2, 2, 410, 411, 3, 2, 
	2, 2, 411, 67, 3, 2, 2, 2, 412, 415, 5, 70, 36, 2, 413, 414, 9, 5, 2, 2, 
	414, 416, 5, 64, 33, 2, 415, 413, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 
	69, 3, 2, 2, 2, 417, 426, 5, 72, 37, 2, 418, 426, 5, 74, 38, 2, 419, 426, 
	5, 76, 39, 2, 420, 426, 5, 78, 40, 2, 421, 426, 5, 80, 41, 2, 422, 426, 
	5, 82, 42, 2, 423, 426, 5, 84, 43, 2, 424, 426, 5, 86, 44, 2, 425, 417, 
	3, 2, 2, 2, 425, 418, 3, 2, 2, 2, 425, 419, 3, 2, 2, 2, 425, 420, 3, 2, 
	2, 2, 425, 421, 3, 2, 2, 2, 425, 422, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 
	425, 424, 3, 2, 2, 2, 426, 71, 3, 2, 2, 2, 427, 428, 5, 88, 45, 2, 428, 
	429, 9, 6, 2, 2, 429, 430, 5, 88, 45, 2, 430, 73, 3, 2, 2, 2, 431, 433, 
	5, 88, 45, 2, 432, 434, 7, 65, 2, 2, 433, 432, 3, 2, 2, 2, 433, 434, 3, 
	2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 436, 7, 16, 2, 2, 436, 437, 5, 88, 
	45, 2, 437, 438, 7, 63, 2, 2, 438, 439, 5, 88, 45, 2, 439, 75, 3, 2, 2, 
	2, 440, 441, 5, 88, 45, 2, 441, 443, 7, 14, 2, 2, 442, 444, 7, 65, 2, 2, 
	443, 442, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 
	446, 9, 7, 2, 2, 446, 77, 3, 2, 2, 2, 447, 449, 5, 88, 45, 2, 448, 450, 
	7, 65, 2, 2, 449, 448, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 451, 3, 2, 
	2, 2, 451, 452, 7, 15, 2, 2, 452, 453, 7, 85, 2, 2, 453, 454, 5, 16, 9, 
	2, 454, 455, 7, 84, 2, 2, 455, 79, 3, 2, 2, 2, 456, 458, 5, 88, 45, 2, 
	457, 459, 7, 65, 2, 2, 458, 457, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 
	460, 3, 2, 2, 2, 460, 461, 7, 15, 2, 2, 461, 462, 7, 85, 2, 2, 462, 463, 
	5, 14, 8, 2, 463, 464, 7, 84, 2, 2, 464, 81, 3, 2, 2, 2, 465, 467, 7, 65, 
	2, 2, 466, 465, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 
	468, 469, 7, 17, 2, 2, 469, 470, 7, 85, 2, 2, 470, 471, 5, 16, 9, 2, 471, 
	472, 7, 84, 2, 2, 472, 83, 3, 2, 2, 2, 473, 474, 9, 8, 2, 2, 474, 475, 
	5, 70, 36, 2, 475, 85, 3, 2, 2, 2, 476, 478, 5, 88, 45, 2, 477, 479, 7, 
	65, 2, 2, 478, 477, 3, 2, 2, 2, 478, 479, 3, 2, 2, 2, 479, 480, 3, 2, 2, 
	2, 480, 481, 7, 28, 2, 2, 481, 482, 5, 88, 45, 2, 482, 87, 3, 2, 2, 2, 
	483, 487, 5, 90, 46, 2, 484, 487, 5, 102, 52, 2, 485, 487, 5, 106, 54, 
	2, 486, 483, 3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 486, 485, 3, 2, 2, 2, 487, 
	89, 3, 2, 2, 2, 488, 498, 5, 92, 47, 2, 489, 498, 5, 94, 48, 2, 490, 498, 
	5, 96, 49, 2, 491, 498, 5, 118, 60, 2, 492, 498, 5, 98, 50, 2, 493, 498, 
	5, 110, 56, 2, 494, 498, 5, 100, 51, 2, 495, 498, 5, 114, 58, 2, 496, 498, 
	5, 116, 59, 2, 497, 488, 3, 2, 2, 2, 497, 489, 3, 2, 2, 2, 497, 490, 3, 
	2, 2, 2, 497, 491, 3, 2, 2, 2, 497, 492, 3, 2, 2, 2, 497, 493, 3, 2, 2, 
	2, 497, 494, 3, 2, 2, 2, 497, 495, 3, 2, 2, 2, 497, 496, 3, 2, 2, 2, 498, 
	91, 3, 2, 2, 2, 499, 500, 9, 9, 2, 2, 500, 93, 3, 2, 2, 2, 501, 502, 9, 
	10, 2, 2, 502, 95, 3, 2, 2, 2, 503, 504, 9, 11, 2, 2, 504, 505, 7, 78, 
	2, 2, 505, 97, 3, 2, 2, 2, 506, 508, 9, 12, 2, 2, 507, 506, 3, 2, 2, 2, 
	507, 508, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509, 510, 7, 85, 2, 2, 510, 
	511, 5, 16, 9, 2, 511, 512, 7, 84, 2, 2, 512, 99, 3, 2, 2, 2, 513, 514, 
	7, 85, 2, 2, 514, 515, 5, 88, 45, 2, 515, 516, 7, 84, 2, 2, 516, 101, 3, 
	2, 2, 2, 517, 518, 7, 85, 2, 2, 518, 519, 5, 104, 53, 2, 519, 520, 7, 84, 
	2, 2, 520, 103, 3, 2, 2, 2, 521, 526, 5, 88, 45, 2, 522, 523, 7, 97, 2, 
	2, 523, 525, 5, 88, 45, 2, 524, 522, 3, 2, 2, 2, 525, 528, 3, 2, 2, 2, 
	526, 524, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 105, 3, 2, 2, 2, 528, 
	526, 3, 2, 2, 2, 529, 531, 5, 90, 46, 2, 530, 532, 5, 108, 55, 2, 531, 
	530, 3, 2, 2, 2, 531, 532, 3, 2, 2, 2, 532, 107, 3, 2, 2, 2, 533, 535, 
	9, 13, 2, 2, 534, 533, 3, 2, 2, 2, 534, 535, 3, 2, 2, 2, 535, 536, 3, 2, 
	2, 2, 536, 537, 5, 106, 54, 2, 537, 109, 3, 2, 2, 2, 538, 540, 7, 56, 2, 
	2, 539, 541, 5, 88, 45, 2, 540, 539, 3, 2, 2, 2, 540, 541, 3, 2, 2, 2, 
	541, 542, 3, 2, 2, 2, 542, 545, 5, 112, 57, 2, 543, 544, 7, 59, 2, 2, 544, 
	546, 5, 88, 45, 2, 545, 543, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 547, 
	3, 2, 2, 2, 547, 548, 7, 60, 2, 2, 548, 111, 3, 2, 2, 2, 549, 552, 7, 57, 
	2, 2, 550, 553, 5, 88, 45, 2, 551, 553, 5, 72, 37, 2, 552, 550, 3, 2, 2, 
	2, 552, 551, 3, 2, 2, 2, 553, 554, 3, 2, 2, 2, 554, 555, 7, 58, 2, 2, 555, 
	557, 5, 88, 45, 2, 556, 558, 5, 112, 57, 2, 557, 556, 3, 2, 2, 2, 557, 
	558, 3, 2, 2, 2, 558, 113, 3, 2, 2, 2, 559, 560, 7, 62, 2, 2, 560, 561, 
	5, 90, 46, 2, 561, 115, 3, 2, 2, 2, 562, 563, 7, 61, 2, 2, 563, 564, 7, 
	85, 2, 2, 564, 565, 5, 120, 61, 2, 565, 566, 7, 84, 2, 2, 566, 117, 3, 
	2, 2, 2, 567, 568, 7, 79, 2, 2, 568, 570, 7, 85, 2, 2, 569, 571, 5, 120, 
	61, 2, 570, 569, 3, 2, 2, 2, 570, 571, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 
	572, 573, 7, 84, 2, 2, 573, 119, 3, 2, 2, 2, 574, 577, 5, 88, 45, 2, 575, 
	577, 5, 72, 37, 2, 576, 574, 3, 2, 2, 2, 576, 575, 3, 2, 2, 2, 577, 580, 
	3, 2, 2, 2, 578, 579, 7, 97, 2, 2, 579, 581, 5, 120, 61, 2, 580, 578, 3, 
	2, 2, 2, 580, 581, 3, 2, 2, 2, 581, 121, 3, 2, 2, 2, 73, 127, 131, 139, 
	146, 154, 160, 170, 180, 183, 190, 195, 199, 204, 208, 213, 218, 224, 226, 
	234, 238, 245, 251, 254, 258, 265, 270, 275, 281, 300, 303, 306, 312, 316, 
	318, 327, 331, 335, 339, 347, 351, 359, 366, 373, 380, 384, 388, 391, 398, 
	403, 410, 415, 425, 433, 443, 449, 458, 466, 478, 486, 497, 507, 526, 531, 
	534, 540, 545, 552, 557, 570, 576, 580,
}
var literalNames = []string{
	"", "'!'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'|'", 
	"'&'", "'<<'", "'>>'", "'*'", "'^'", "", "", "", "", "", "'regexp'", "'~'", 
	"'escape'", "')'", "'('", "']'", "'['", "':'", "'.*'", "'='", "'<'", "'>'", 
	"'!='", "'<='", "'>='", "';'", "','", "'.'", "'collate'", "'index'", "'key'", 
	"'use'", "'ignore'", "'partition'", "'straight_join'", "'natural'", "'oj'",
}
var symbolicNames = []string{
	"", "", "PLACEHOLDER", "SELECT", "INSERT", "INTO", "VALUES", "DELETE", 
	"FROM", "WHERE", "LIMIT", "NULL", "IS", "IN", "BETWEEN", "EXISTS", "TRUE", 
	"FALSE", "XOR", "DEFAULT", "UPDATE", "SET", "ORDER", "GROUP", "BY", "FOR", 
	"LIKE", "HAVING", "AS", "INNER", "OUTER", "JOIN", "LEFT", "RIGHT", "ON", 
	"DISTINCT", "OFFSET", "ASC", "DESC", "CROSS", "USING", "DATE", "TIME", 
	"TIMESTAMP", "ALL", "ANY", "SOME", "UNION", "UNKNOWN", "LOCK", "SHARE", 
	"MODE", "COMMIT", "ROLLBACK", "CASE", "WHEN", "THEN", "ELSE", "END", "ROW", 
	"BINARY", "AND", "OR", "NOT", "DIV", "MOD", "PLUS", "MINUS", "VERTBAR", 
	"BITAND", "SHIFT_LEFT", "SHIFT_RIGHT", "ASTERISK", "POWER_OP", "INT", "DECIMAL", 
	"STRING", "ID", "COLUMN_REL", "REGEXP", "NEGATION", "ESCAPE", "RPAREN", 
	"LPAREN", "RBRACK", "LBRACK", "COLON", "ALL_FIELDS", "EQ", "LTH", "GTH", 
	"NOT_EQ", "LET", "GET", "SEMI", "COMMA", "DOT", "COLLATE", "INDEX", "KEY", 
	"USE", "IGNORE", "PARTITION", "STRAIGHT_JOIN", "NATURAL", "OJ", "NEWLINE", 
	"WS", "COMMENT1", "COMMENT2", "USER_VAR",
}

var ruleNames = []string{
	"stat", "transcationStat", "commit", "rollback", "insertStat", "columnNames", 
	"valueList", "selectStat", "selectInner", "selectPrefix", "selectSuffix", 
	"selectUnionSuffix", "selectExprs", "tables", "tableRel", "tableFactor", 
	"tableSubQuery", "tableRecu", "tableJoin", "tableJoinSuffix", "tableJoinMod", 
	"joinCondition", "gbobExprs", "updateStat", "updateSingleTable", "updateMultipleTable", 
	"setExprs", "setExpr", "deleteStat", "tableNameAndAlias", "tableNameAndAliases", 
	"whereCondition", "whereCondSub", "whereCondOp", "expression", "exprRelational", 
	"exprBetweenAnd", "exprIsOrIsNotNull", "exprInSelect", "exprInValues", 
	"exprExists", "exprNot", "exprLike", "element", "elementOpFactory", "elementText", 
	"elementTextParam", "elementDate", "elementSubQuery", "elementWapperBkt", 
	"elementListFactor", "elementList", "elementOpEle", "elementOpEleSuffix", 
	"elementCase", "caseWhenPart", "elementWithPrefix", "elementRow", "funCall", 
	"paramList",
}
type MySQLParser struct {
	*antlr.BaseParser
}

// NewMySQLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *MySQLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewMySQLParser(input antlr.TokenStream) *MySQLParser {
	this := new(MySQLParser)
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
	this.GrammarFileName = "MySQL.g4"

	return this
}


// MySQLParser tokens.
const (
	MySQLParserEOF = antlr.TokenEOF
	MySQLParserT__0 = 1
	MySQLParserPLACEHOLDER = 2
	MySQLParserSELECT = 3
	MySQLParserINSERT = 4
	MySQLParserINTO = 5
	MySQLParserVALUES = 6
	MySQLParserDELETE = 7
	MySQLParserFROM = 8
	MySQLParserWHERE = 9
	MySQLParserLIMIT = 10
	MySQLParserNULL = 11
	MySQLParserIS = 12
	MySQLParserIN = 13
	MySQLParserBETWEEN = 14
	MySQLParserEXISTS = 15
	MySQLParserTRUE = 16
	MySQLParserFALSE = 17
	MySQLParserXOR = 18
	MySQLParserDEFAULT = 19
	MySQLParserUPDATE = 20
	MySQLParserSET = 21
	MySQLParserORDER = 22
	MySQLParserGROUP = 23
	MySQLParserBY = 24
	MySQLParserFOR = 25
	MySQLParserLIKE = 26
	MySQLParserHAVING = 27
	MySQLParserAS = 28
	MySQLParserINNER = 29
	MySQLParserOUTER = 30
	MySQLParserJOIN = 31
	MySQLParserLEFT = 32
	MySQLParserRIGHT = 33
	MySQLParserON = 34
	MySQLParserDISTINCT = 35
	MySQLParserOFFSET = 36
	MySQLParserASC = 37
	MySQLParserDESC = 38
	MySQLParserCROSS = 39
	MySQLParserUSING = 40
	MySQLParserDATE = 41
	MySQLParserTIME = 42
	MySQLParserTIMESTAMP = 43
	MySQLParserALL = 44
	MySQLParserANY = 45
	MySQLParserSOME = 46
	MySQLParserUNION = 47
	MySQLParserUNKNOWN = 48
	MySQLParserLOCK = 49
	MySQLParserSHARE = 50
	MySQLParserMODE = 51
	MySQLParserCOMMIT = 52
	MySQLParserROLLBACK = 53
	MySQLParserCASE = 54
	MySQLParserWHEN = 55
	MySQLParserTHEN = 56
	MySQLParserELSE = 57
	MySQLParserEND = 58
	MySQLParserROW = 59
	MySQLParserBINARY = 60
	MySQLParserAND = 61
	MySQLParserOR = 62
	MySQLParserNOT = 63
	MySQLParserDIV = 64
	MySQLParserMOD = 65
	MySQLParserPLUS = 66
	MySQLParserMINUS = 67
	MySQLParserVERTBAR = 68
	MySQLParserBITAND = 69
	MySQLParserSHIFT_LEFT = 70
	MySQLParserSHIFT_RIGHT = 71
	MySQLParserASTERISK = 72
	MySQLParserPOWER_OP = 73
	MySQLParserINT = 74
	MySQLParserDECIMAL = 75
	MySQLParserSTRING = 76
	MySQLParserID = 77
	MySQLParserCOLUMN_REL = 78
	MySQLParserREGEXP = 79
	MySQLParserNEGATION = 80
	MySQLParserESCAPE = 81
	MySQLParserRPAREN = 82
	MySQLParserLPAREN = 83
	MySQLParserRBRACK = 84
	MySQLParserLBRACK = 85
	MySQLParserCOLON = 86
	MySQLParserALL_FIELDS = 87
	MySQLParserEQ = 88
	MySQLParserLTH = 89
	MySQLParserGTH = 90
	MySQLParserNOT_EQ = 91
	MySQLParserLET = 92
	MySQLParserGET = 93
	MySQLParserSEMI = 94
	MySQLParserCOMMA = 95
	MySQLParserDOT = 96
	MySQLParserCOLLATE = 97
	MySQLParserINDEX = 98
	MySQLParserKEY = 99
	MySQLParserUSE = 100
	MySQLParserIGNORE = 101
	MySQLParserPARTITION = 102
	MySQLParserSTRAIGHT_JOIN = 103
	MySQLParserNATURAL = 104
	MySQLParserOJ = 105
	MySQLParserNEWLINE = 106
	MySQLParserWS = 107
	MySQLParserCOMMENT1 = 108
	MySQLParserCOMMENT2 = 109
	MySQLParserUSER_VAR = 110
)

// MySQLParser rules.
const (
	MySQLParserRULE_stat = 0
	MySQLParserRULE_transcationStat = 1
	MySQLParserRULE_commit = 2
	MySQLParserRULE_rollback = 3
	MySQLParserRULE_insertStat = 4
	MySQLParserRULE_columnNames = 5
	MySQLParserRULE_valueList = 6
	MySQLParserRULE_selectStat = 7
	MySQLParserRULE_selectInner = 8
	MySQLParserRULE_selectPrefix = 9
	MySQLParserRULE_selectSuffix = 10
	MySQLParserRULE_selectUnionSuffix = 11
	MySQLParserRULE_selectExprs = 12
	MySQLParserRULE_tables = 13
	MySQLParserRULE_tableRel = 14
	MySQLParserRULE_tableFactor = 15
	MySQLParserRULE_tableSubQuery = 16
	MySQLParserRULE_tableRecu = 17
	MySQLParserRULE_tableJoin = 18
	MySQLParserRULE_tableJoinSuffix = 19
	MySQLParserRULE_tableJoinMod = 20
	MySQLParserRULE_joinCondition = 21
	MySQLParserRULE_gbobExprs = 22
	MySQLParserRULE_updateStat = 23
	MySQLParserRULE_updateSingleTable = 24
	MySQLParserRULE_updateMultipleTable = 25
	MySQLParserRULE_setExprs = 26
	MySQLParserRULE_setExpr = 27
	MySQLParserRULE_deleteStat = 28
	MySQLParserRULE_tableNameAndAlias = 29
	MySQLParserRULE_tableNameAndAliases = 30
	MySQLParserRULE_whereCondition = 31
	MySQLParserRULE_whereCondSub = 32
	MySQLParserRULE_whereCondOp = 33
	MySQLParserRULE_expression = 34
	MySQLParserRULE_exprRelational = 35
	MySQLParserRULE_exprBetweenAnd = 36
	MySQLParserRULE_exprIsOrIsNotNull = 37
	MySQLParserRULE_exprInSelect = 38
	MySQLParserRULE_exprInValues = 39
	MySQLParserRULE_exprExists = 40
	MySQLParserRULE_exprNot = 41
	MySQLParserRULE_exprLike = 42
	MySQLParserRULE_element = 43
	MySQLParserRULE_elementOpFactory = 44
	MySQLParserRULE_elementText = 45
	MySQLParserRULE_elementTextParam = 46
	MySQLParserRULE_elementDate = 47
	MySQLParserRULE_elementSubQuery = 48
	MySQLParserRULE_elementWapperBkt = 49
	MySQLParserRULE_elementListFactor = 50
	MySQLParserRULE_elementList = 51
	MySQLParserRULE_elementOpEle = 52
	MySQLParserRULE_elementOpEleSuffix = 53
	MySQLParserRULE_elementCase = 54
	MySQLParserRULE_caseWhenPart = 55
	MySQLParserRULE_elementWithPrefix = 56
	MySQLParserRULE_elementRow = 57
	MySQLParserRULE_funCall = 58
	MySQLParserRULE_paramList = 59
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) InsertStat() IInsertStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStatContext)
}

func (s *StatContext) UpdateStat() IUpdateStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStatContext)
}

func (s *StatContext) DeleteStat() IDeleteStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteStatContext)
}

func (s *StatContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *StatContext) TranscationStat() ITranscationStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITranscationStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITranscationStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitStat(s)
	}
}




func (p *MySQLParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MySQLParserRULE_stat)

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

	switch p.GetTokenStream().LA(1) {
	case MySQLParserINSERT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.InsertStat()
		}


	case MySQLParserUPDATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.UpdateStat()
		}


	case MySQLParserDELETE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.DeleteStat()
		}


	case MySQLParserSELECT, MySQLParserLPAREN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			p.SelectStat()
		}


	case MySQLParserCOMMIT, MySQLParserROLLBACK:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(124)
			p.TranscationStat()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITranscationStatContext is an interface to support dynamic dispatch.
type ITranscationStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranscationStatContext differentiates from other interfaces.
	IsTranscationStatContext()
}

type TranscationStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranscationStatContext() *TranscationStatContext {
	var p = new(TranscationStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_transcationStat
	return p
}

func (*TranscationStatContext) IsTranscationStatContext() {}

func NewTranscationStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranscationStatContext {
	var p = new(TranscationStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_transcationStat

	return p
}

func (s *TranscationStatContext) GetParser() antlr.Parser { return s.parser }

func (s *TranscationStatContext) Commit() ICommitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommitContext)
}

func (s *TranscationStatContext) Rollback() IRollbackContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRollbackContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRollbackContext)
}

func (s *TranscationStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranscationStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TranscationStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTranscationStat(s)
	}
}

func (s *TranscationStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTranscationStat(s)
	}
}




func (p *MySQLParser) TranscationStat() (localctx ITranscationStatContext) {
	localctx = NewTranscationStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MySQLParserRULE_transcationStat)

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

	switch p.GetTokenStream().LA(1) {
	case MySQLParserCOMMIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Commit()
		}


	case MySQLParserROLLBACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Rollback()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ICommitContext is an interface to support dynamic dispatch.
type ICommitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommitContext differentiates from other interfaces.
	IsCommitContext()
}

type CommitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommitContext() *CommitContext {
	var p = new(CommitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_commit
	return p
}

func (*CommitContext) IsCommitContext() {}

func NewCommitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommitContext {
	var p = new(CommitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_commit

	return p
}

func (s *CommitContext) GetParser() antlr.Parser { return s.parser }

func (s *CommitContext) COMMIT() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMIT, 0)
}

func (s *CommitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CommitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterCommit(s)
	}
}

func (s *CommitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitCommit(s)
	}
}




func (p *MySQLParser) Commit() (localctx ICommitContext) {
	localctx = NewCommitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MySQLParserRULE_commit)

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
		p.SetState(131)
		p.Match(MySQLParserCOMMIT)
	}



	return localctx
}


// IRollbackContext is an interface to support dynamic dispatch.
type IRollbackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRollbackContext differentiates from other interfaces.
	IsRollbackContext()
}

type RollbackContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRollbackContext() *RollbackContext {
	var p = new(RollbackContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_rollback
	return p
}

func (*RollbackContext) IsRollbackContext() {}

func NewRollbackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RollbackContext {
	var p = new(RollbackContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_rollback

	return p
}

func (s *RollbackContext) GetParser() antlr.Parser { return s.parser }

func (s *RollbackContext) ROLLBACK() antlr.TerminalNode {
	return s.GetToken(MySQLParserROLLBACK, 0)
}

func (s *RollbackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RollbackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RollbackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterRollback(s)
	}
}

func (s *RollbackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitRollback(s)
	}
}




func (p *MySQLParser) Rollback() (localctx IRollbackContext) {
	localctx = NewRollbackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MySQLParserRULE_rollback)

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
		p.SetState(133)
		p.Match(MySQLParserROLLBACK)
	}



	return localctx
}


// IInsertStatContext is an interface to support dynamic dispatch.
type IInsertStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName token.
	GetTableName() antlr.Token 


	// SetTableName sets the tableName token.
	SetTableName(antlr.Token) 


	// IsInsertStatContext differentiates from other interfaces.
	IsInsertStatContext()
}

type InsertStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tableName antlr.Token
}

func NewEmptyInsertStatContext() *InsertStatContext {
	var p = new(InsertStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_insertStat
	return p
}

func (*InsertStatContext) IsInsertStatContext() {}

func NewInsertStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatContext {
	var p = new(InsertStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_insertStat

	return p
}

func (s *InsertStatContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatContext) GetTableName() antlr.Token { return s.tableName }


func (s *InsertStatContext) SetTableName(v antlr.Token) { s.tableName = v }


func (s *InsertStatContext) INSERT() antlr.TerminalNode {
	return s.GetToken(MySQLParserINSERT, 0)
}

func (s *InsertStatContext) ID() antlr.TerminalNode {
	return s.GetToken(MySQLParserID, 0)
}

func (s *InsertStatContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *InsertStatContext) INTO() antlr.TerminalNode {
	return s.GetToken(MySQLParserINTO, 0)
}

func (s *InsertStatContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserLPAREN)
}

func (s *InsertStatContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, i)
}

func (s *InsertStatContext) ColumnNames() IColumnNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNamesContext)
}

func (s *InsertStatContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserRPAREN)
}

func (s *InsertStatContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, i)
}

func (s *InsertStatContext) VALUES() antlr.TerminalNode {
	return s.GetToken(MySQLParserVALUES, 0)
}

func (s *InsertStatContext) ValueList() IValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *InsertStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InsertStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterInsertStat(s)
	}
}

func (s *InsertStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitInsertStat(s)
	}
}




func (p *MySQLParser) InsertStat() (localctx IInsertStatContext) {
	localctx = NewInsertStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MySQLParserRULE_insertStat)
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
		p.SetState(135)
		p.Match(MySQLParserINSERT)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserINTO {
		{
			p.SetState(136)
			p.Match(MySQLParserINTO)
		}

	}
	{
		p.SetState(139)

		var _m = p.Match(MySQLParserID)

		localctx.(*InsertStatContext).tableName = _m
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(140)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(141)
			p.ColumnNames()
		}
		{
			p.SetState(142)
			p.Match(MySQLParserRPAREN)
		}


	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySQLParserVALUES:
		{
			p.SetState(146)
			p.Match(MySQLParserVALUES)
		}
		{
			p.SetState(147)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(148)
			p.ValueList()
		}
		{
			p.SetState(149)
			p.Match(MySQLParserRPAREN)
		}



	case MySQLParserSELECT, MySQLParserLPAREN:
		{
			p.SetState(151)
			p.SelectStat()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IColumnNamesContext is an interface to support dynamic dispatch.
type IColumnNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNamesContext differentiates from other interfaces.
	IsColumnNamesContext()
}

type ColumnNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNamesContext() *ColumnNamesContext {
	var p = new(ColumnNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_columnNames
	return p
}

func (*ColumnNamesContext) IsColumnNamesContext() {}

func NewColumnNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNamesContext {
	var p = new(ColumnNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_columnNames

	return p
}

func (s *ColumnNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNamesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserID)
}

func (s *ColumnNamesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserID, i)
}

func (s *ColumnNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *ColumnNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *ColumnNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ColumnNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterColumnNames(s)
	}
}

func (s *ColumnNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitColumnNames(s)
	}
}




func (p *MySQLParser) ColumnNames() (localctx IColumnNamesContext) {
	localctx = NewColumnNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MySQLParserRULE_columnNames)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(154)
				p.Match(MySQLParserID)
			}
			{
				p.SetState(155)
				p.Match(MySQLParserCOMMA)
			}


		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}
	{
		p.SetState(161)
		p.Match(MySQLParserID)
	}



	return localctx
}


// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_valueList
	return p
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ValueListContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ValueListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *ValueListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitValueList(s)
	}
}




func (p *MySQLParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MySQLParserRULE_valueList)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(163)
				p.Element()
			}
			{
				p.SetState(164)
				p.Match(MySQLParserCOMMA)
			}


		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	{
		p.SetState(171)
		p.Element()
	}



	return localctx
}


// ISelectStatContext is an interface to support dynamic dispatch.
type ISelectStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatContext differentiates from other interfaces.
	IsSelectStatContext()
}

type SelectStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatContext() *SelectStatContext {
	var p = new(SelectStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectStat
	return p
}

func (*SelectStatContext) IsSelectStatContext() {}

func NewSelectStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatContext {
	var p = new(SelectStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectStat

	return p
}

func (s *SelectStatContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *SelectStatContext) SelectInner() ISelectInnerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectInnerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectInnerContext)
}

func (s *SelectStatContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *SelectStatContext) SelectUnionSuffix() ISelectUnionSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectUnionSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectUnionSuffixContext)
}

func (s *SelectStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectStat(s)
	}
}

func (s *SelectStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectStat(s)
	}
}




func (p *MySQLParser) SelectStat() (localctx ISelectStatContext) {
	localctx = NewSelectStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MySQLParserRULE_selectStat)
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
	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySQLParserLPAREN:
		{
			p.SetState(173)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(174)
			p.SelectInner()
		}
		{
			p.SetState(175)
			p.Match(MySQLParserRPAREN)
		}


	case MySQLParserSELECT:
		{
			p.SetState(177)
			p.SelectInner()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserUNION {
		{
			p.SetState(180)
			p.SelectUnionSuffix()
		}

	}



	return localctx
}


// ISelectInnerContext is an interface to support dynamic dispatch.
type ISelectInnerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectInnerContext differentiates from other interfaces.
	IsSelectInnerContext()
}

type SelectInnerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectInnerContext() *SelectInnerContext {
	var p = new(SelectInnerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectInner
	return p
}

func (*SelectInnerContext) IsSelectInnerContext() {}

func NewSelectInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectInnerContext {
	var p = new(SelectInnerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectInner

	return p
}

func (s *SelectInnerContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectInnerContext) SelectPrefix() ISelectPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectPrefixContext)
}

func (s *SelectInnerContext) SelectSuffix() ISelectSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectSuffixContext)
}

func (s *SelectInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectInner(s)
	}
}

func (s *SelectInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectInner(s)
	}
}




func (p *MySQLParser) SelectInner() (localctx ISelectInnerContext) {
	localctx = NewSelectInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MySQLParserRULE_selectInner)

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
		p.SelectPrefix()
	}
	{
		p.SetState(184)
		p.SelectSuffix()
	}



	return localctx
}


// ISelectPrefixContext is an interface to support dynamic dispatch.
type ISelectPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDistinct returns the distinct token.
	GetDistinct() antlr.Token 


	// SetDistinct sets the distinct token.
	SetDistinct(antlr.Token) 


	// GetWhere returns the where rule contexts.
	GetWhere() IWhereConditionContext

	// GetGroupByExprs returns the groupByExprs rule contexts.
	GetGroupByExprs() IGbobExprsContext

	// GetHaving returns the having rule contexts.
	GetHaving() IWhereConditionContext


	// SetWhere sets the where rule contexts.
	SetWhere(IWhereConditionContext)

	// SetGroupByExprs sets the groupByExprs rule contexts.
	SetGroupByExprs(IGbobExprsContext)

	// SetHaving sets the having rule contexts.
	SetHaving(IWhereConditionContext)


	// IsSelectPrefixContext differentiates from other interfaces.
	IsSelectPrefixContext()
}

type SelectPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	distinct antlr.Token
	where IWhereConditionContext 
	groupByExprs IGbobExprsContext 
	having IWhereConditionContext 
}

func NewEmptySelectPrefixContext() *SelectPrefixContext {
	var p = new(SelectPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectPrefix
	return p
}

func (*SelectPrefixContext) IsSelectPrefixContext() {}

func NewSelectPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectPrefixContext {
	var p = new(SelectPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectPrefix

	return p
}

func (s *SelectPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectPrefixContext) GetDistinct() antlr.Token { return s.distinct }


func (s *SelectPrefixContext) SetDistinct(v antlr.Token) { s.distinct = v }


func (s *SelectPrefixContext) GetWhere() IWhereConditionContext { return s.where }

func (s *SelectPrefixContext) GetGroupByExprs() IGbobExprsContext { return s.groupByExprs }

func (s *SelectPrefixContext) GetHaving() IWhereConditionContext { return s.having }


func (s *SelectPrefixContext) SetWhere(v IWhereConditionContext) { s.where = v }

func (s *SelectPrefixContext) SetGroupByExprs(v IGbobExprsContext) { s.groupByExprs = v }

func (s *SelectPrefixContext) SetHaving(v IWhereConditionContext) { s.having = v }


func (s *SelectPrefixContext) SELECT() antlr.TerminalNode {
	return s.GetToken(MySQLParserSELECT, 0)
}

func (s *SelectPrefixContext) SelectExprs() ISelectExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectExprsContext)
}

func (s *SelectPrefixContext) FROM() antlr.TerminalNode {
	return s.GetToken(MySQLParserFROM, 0)
}

func (s *SelectPrefixContext) Tables() ITablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITablesContext)
}

func (s *SelectPrefixContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySQLParserWHERE, 0)
}

func (s *SelectPrefixContext) GROUP() antlr.TerminalNode {
	return s.GetToken(MySQLParserGROUP, 0)
}

func (s *SelectPrefixContext) BY() antlr.TerminalNode {
	return s.GetToken(MySQLParserBY, 0)
}

func (s *SelectPrefixContext) HAVING() antlr.TerminalNode {
	return s.GetToken(MySQLParserHAVING, 0)
}

func (s *SelectPrefixContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(MySQLParserDISTINCT, 0)
}

func (s *SelectPrefixContext) AllWhereCondition() []IWhereConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem())
	var tst = make([]IWhereConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereConditionContext)
		}
	}

	return tst
}

func (s *SelectPrefixContext) WhereCondition(i int) IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *SelectPrefixContext) GbobExprs() IGbobExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGbobExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGbobExprsContext)
}

func (s *SelectPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectPrefix(s)
	}
}

func (s *SelectPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectPrefix(s)
	}
}




func (p *MySQLParser) SelectPrefix() (localctx ISelectPrefixContext) {
	localctx = NewSelectPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MySQLParserRULE_selectPrefix)
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
		p.SetState(186)
		p.Match(MySQLParserSELECT)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserDISTINCT {
		{
			p.SetState(187)

			var _m = p.Match(MySQLParserDISTINCT)

			localctx.(*SelectPrefixContext).distinct = _m
		}

	}
	{
		p.SetState(190)
		p.SelectExprs()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserFROM {
		{
			p.SetState(191)
			p.Match(MySQLParserFROM)
		}
		{
			p.SetState(192)
			p.Tables()
		}

	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserWHERE {
		{
			p.SetState(195)
			p.Match(MySQLParserWHERE)
		}
		{
			p.SetState(196)

			var _x = p.WhereCondition()


			localctx.(*SelectPrefixContext).where = _x
		}

	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserGROUP {
		{
			p.SetState(199)
			p.Match(MySQLParserGROUP)
		}
		{
			p.SetState(200)
			p.Match(MySQLParserBY)
		}
		{
			p.SetState(201)

			var _x = p.GbobExprs()


			localctx.(*SelectPrefixContext).groupByExprs = _x
		}

	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserHAVING {
		{
			p.SetState(204)
			p.Match(MySQLParserHAVING)
		}
		{
			p.SetState(205)

			var _x = p.WhereCondition()


			localctx.(*SelectPrefixContext).having = _x
		}

	}



	return localctx
}


// ISelectSuffixContext is an interface to support dynamic dispatch.
type ISelectSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset token.
	GetOffset() antlr.Token 

	// GetRowCount returns the rowCount token.
	GetRowCount() antlr.Token 

	// GetLock returns the lock token.
	GetLock() antlr.Token 


	// SetOffset sets the offset token.
	SetOffset(antlr.Token) 

	// SetRowCount sets the rowCount token.
	SetRowCount(antlr.Token) 

	// SetLock sets the lock token.
	SetLock(antlr.Token) 


	// GetOrderByExprs returns the orderByExprs rule contexts.
	GetOrderByExprs() IGbobExprsContext


	// SetOrderByExprs sets the orderByExprs rule contexts.
	SetOrderByExprs(IGbobExprsContext)


	// IsSelectSuffixContext differentiates from other interfaces.
	IsSelectSuffixContext()
}

type SelectSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	orderByExprs IGbobExprsContext 
	offset antlr.Token
	rowCount antlr.Token
	lock antlr.Token
}

func NewEmptySelectSuffixContext() *SelectSuffixContext {
	var p = new(SelectSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectSuffix
	return p
}

func (*SelectSuffixContext) IsSelectSuffixContext() {}

func NewSelectSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectSuffixContext {
	var p = new(SelectSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectSuffix

	return p
}

func (s *SelectSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectSuffixContext) GetOffset() antlr.Token { return s.offset }

func (s *SelectSuffixContext) GetRowCount() antlr.Token { return s.rowCount }

func (s *SelectSuffixContext) GetLock() antlr.Token { return s.lock }


func (s *SelectSuffixContext) SetOffset(v antlr.Token) { s.offset = v }

func (s *SelectSuffixContext) SetRowCount(v antlr.Token) { s.rowCount = v }

func (s *SelectSuffixContext) SetLock(v antlr.Token) { s.lock = v }


func (s *SelectSuffixContext) GetOrderByExprs() IGbobExprsContext { return s.orderByExprs }


func (s *SelectSuffixContext) SetOrderByExprs(v IGbobExprsContext) { s.orderByExprs = v }


func (s *SelectSuffixContext) ORDER() antlr.TerminalNode {
	return s.GetToken(MySQLParserORDER, 0)
}

func (s *SelectSuffixContext) BY() antlr.TerminalNode {
	return s.GetToken(MySQLParserBY, 0)
}

func (s *SelectSuffixContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MySQLParserLIMIT, 0)
}

func (s *SelectSuffixContext) GbobExprs() IGbobExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGbobExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGbobExprsContext)
}

func (s *SelectSuffixContext) FOR() antlr.TerminalNode {
	return s.GetToken(MySQLParserFOR, 0)
}

func (s *SelectSuffixContext) IN() antlr.TerminalNode {
	return s.GetToken(MySQLParserIN, 0)
}

func (s *SelectSuffixContext) SHARE() antlr.TerminalNode {
	return s.GetToken(MySQLParserSHARE, 0)
}

func (s *SelectSuffixContext) MODE() antlr.TerminalNode {
	return s.GetToken(MySQLParserMODE, 0)
}

func (s *SelectSuffixContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(MySQLParserUPDATE, 0)
}

func (s *SelectSuffixContext) LOCK() antlr.TerminalNode {
	return s.GetToken(MySQLParserLOCK, 0)
}

func (s *SelectSuffixContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(MySQLParserOFFSET, 0)
}

func (s *SelectSuffixContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserINT)
}

func (s *SelectSuffixContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserINT, i)
}

func (s *SelectSuffixContext) AllPLACEHOLDER() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserPLACEHOLDER)
}

func (s *SelectSuffixContext) PLACEHOLDER(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserPLACEHOLDER, i)
}

func (s *SelectSuffixContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, 0)
}

func (s *SelectSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectSuffix(s)
	}
}

func (s *SelectSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectSuffix(s)
	}
}




func (p *MySQLParser) SelectSuffix() (localctx ISelectSuffixContext) {
	localctx = NewSelectSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MySQLParserRULE_selectSuffix)
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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(208)
			p.Match(MySQLParserORDER)
		}
		{
			p.SetState(209)
			p.Match(MySQLParserBY)
		}
		{
			p.SetState(210)

			var _x = p.GbobExprs()


			localctx.(*SelectSuffixContext).orderByExprs = _x
		}


	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(213)
			p.Match(MySQLParserLIMIT)
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			p.SetState(216)
			p.GetErrorHandler().Sync(p)


			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(214)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SelectSuffixContext).offset = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SelectSuffixContext).offset = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(215)
					p.Match(MySQLParserCOMMA)
				}


			}
			{
				p.SetState(218)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SelectSuffixContext).rowCount = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SelectSuffixContext).rowCount = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}



		case 2:
			{
				p.SetState(219)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SelectSuffixContext).rowCount = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SelectSuffixContext).rowCount = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(220)
				p.Match(MySQLParserOFFSET)
			}
			{
				p.SetState(221)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SelectSuffixContext).offset = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SelectSuffixContext).offset = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}


		}


	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(226)
			p.Match(MySQLParserFOR)
		}
		{
			p.SetState(227)

			var _m = p.Match(MySQLParserUPDATE)

			localctx.(*SelectSuffixContext).lock = _m
		}


	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(228)

			var _m = p.Match(MySQLParserLOCK)

			localctx.(*SelectSuffixContext).lock = _m
		}
		{
			p.SetState(229)
			p.Match(MySQLParserIN)
		}
		{
			p.SetState(230)
			p.Match(MySQLParserSHARE)
		}
		{
			p.SetState(231)
			p.Match(MySQLParserMODE)
		}



	}



	return localctx
}


// ISelectUnionSuffixContext is an interface to support dynamic dispatch.
type ISelectUnionSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMethod returns the method token.
	GetMethod() antlr.Token 


	// SetMethod sets the method token.
	SetMethod(antlr.Token) 


	// IsSelectUnionSuffixContext differentiates from other interfaces.
	IsSelectUnionSuffixContext()
}

type SelectUnionSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	method antlr.Token
}

func NewEmptySelectUnionSuffixContext() *SelectUnionSuffixContext {
	var p = new(SelectUnionSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectUnionSuffix
	return p
}

func (*SelectUnionSuffixContext) IsSelectUnionSuffixContext() {}

func NewSelectUnionSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectUnionSuffixContext {
	var p = new(SelectUnionSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectUnionSuffix

	return p
}

func (s *SelectUnionSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectUnionSuffixContext) GetMethod() antlr.Token { return s.method }


func (s *SelectUnionSuffixContext) SetMethod(v antlr.Token) { s.method = v }


func (s *SelectUnionSuffixContext) UNION() antlr.TerminalNode {
	return s.GetToken(MySQLParserUNION, 0)
}

func (s *SelectUnionSuffixContext) SelectSuffix() ISelectSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectSuffixContext)
}

func (s *SelectUnionSuffixContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *SelectUnionSuffixContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *SelectUnionSuffixContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *SelectUnionSuffixContext) ALL() antlr.TerminalNode {
	return s.GetToken(MySQLParserALL, 0)
}

func (s *SelectUnionSuffixContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(MySQLParserDISTINCT, 0)
}

func (s *SelectUnionSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectUnionSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectUnionSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectUnionSuffix(s)
	}
}

func (s *SelectUnionSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectUnionSuffix(s)
	}
}




func (p *MySQLParser) SelectUnionSuffix() (localctx ISelectUnionSuffixContext) {
	localctx = NewSelectUnionSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MySQLParserRULE_selectUnionSuffix)
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
		p.SetState(234)
		p.Match(MySQLParserUNION)
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserDISTINCT || _la == MySQLParserALL {
		{
			p.SetState(235)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SelectUnionSuffixContext).method = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserDISTINCT || _la == MySQLParserALL) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SelectUnionSuffixContext).method = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(238)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(239)
			p.SelectStat()
		}
		{
			p.SetState(240)
			p.Match(MySQLParserRPAREN)
		}


	case 2:
		{
			p.SetState(242)
			p.SelectStat()
		}

	}
	{
		p.SetState(245)
		p.SelectSuffix()
	}



	return localctx
}


// ISelectExprsContext is an interface to support dynamic dispatch.
type ISelectExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token 


	// SetAlias sets the alias token.
	SetAlias(antlr.Token) 


	// IsSelectExprsContext differentiates from other interfaces.
	IsSelectExprsContext()
}

type SelectExprsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias antlr.Token
}

func NewEmptySelectExprsContext() *SelectExprsContext {
	var p = new(SelectExprsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_selectExprs
	return p
}

func (*SelectExprsContext) IsSelectExprsContext() {}

func NewSelectExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectExprsContext {
	var p = new(SelectExprsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_selectExprs

	return p
}

func (s *SelectExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectExprsContext) GetAlias() antlr.Token { return s.alias }


func (s *SelectExprsContext) SetAlias(v antlr.Token) { s.alias = v }


func (s *SelectExprsContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SelectExprsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, 0)
}

func (s *SelectExprsContext) SelectExprs() ISelectExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectExprsContext)
}

func (s *SelectExprsContext) ID() antlr.TerminalNode {
	return s.GetToken(MySQLParserID, 0)
}

func (s *SelectExprsContext) AS() antlr.TerminalNode {
	return s.GetToken(MySQLParserAS, 0)
}

func (s *SelectExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectExprsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSelectExprs(s)
	}
}

func (s *SelectExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSelectExprs(s)
	}
}




func (p *MySQLParser) SelectExprs() (localctx ISelectExprsContext) {
	localctx = NewSelectExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MySQLParserRULE_selectExprs)
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
		p.SetState(247)
		p.Element()
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserAS || _la == MySQLParserID {
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySQLParserAS {
			{
				p.SetState(248)
				p.Match(MySQLParserAS)
			}

		}
		{
			p.SetState(251)

			var _m = p.Match(MySQLParserID)

			localctx.(*SelectExprsContext).alias = _m
		}

	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserCOMMA {
		{
			p.SetState(254)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(255)
			p.SelectExprs()
		}

	}



	return localctx
}


// ITablesContext is an interface to support dynamic dispatch.
type ITablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTablesContext differentiates from other interfaces.
	IsTablesContext()
}

type TablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTablesContext() *TablesContext {
	var p = new(TablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tables
	return p
}

func (*TablesContext) IsTablesContext() {}

func NewTablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TablesContext {
	var p = new(TablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tables

	return p
}

func (s *TablesContext) GetParser() antlr.Parser { return s.parser }

func (s *TablesContext) AllTableRel() []ITableRelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableRelContext)(nil)).Elem())
	var tst = make([]ITableRelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableRelContext)
		}
	}

	return tst
}

func (s *TablesContext) TableRel(i int) ITableRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableRelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableRelContext)
}

func (s *TablesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *TablesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *TablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTables(s)
	}
}

func (s *TablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTables(s)
	}
}




func (p *MySQLParser) Tables() (localctx ITablesContext) {
	localctx = NewTablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MySQLParserRULE_tables)
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
		p.SetState(258)
		p.TableRel()
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySQLParserCOMMA {
		{
			p.SetState(259)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(260)
			p.TableRel()
		}


		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ITableRelContext is an interface to support dynamic dispatch.
type ITableRelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableRelContext differentiates from other interfaces.
	IsTableRelContext()
}

type TableRelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableRelContext() *TableRelContext {
	var p = new(TableRelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableRel
	return p
}

func (*TableRelContext) IsTableRelContext() {}

func NewTableRelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableRelContext {
	var p = new(TableRelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableRel

	return p
}

func (s *TableRelContext) GetParser() antlr.Parser { return s.parser }

func (s *TableRelContext) TableFactor() ITableFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableFactorContext)
}

func (s *TableRelContext) TableJoin() ITableJoinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableJoinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableJoinContext)
}

func (s *TableRelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableRelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableRelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableRel(s)
	}
}

func (s *TableRelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableRel(s)
	}
}




func (p *MySQLParser) TableRel() (localctx ITableRelContext) {
	localctx = NewTableRelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MySQLParserRULE_tableRel)

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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.TableFactor()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.TableJoin()
		}

	}


	return localctx
}


// ITableFactorContext is an interface to support dynamic dispatch.
type ITableFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableFactorContext differentiates from other interfaces.
	IsTableFactorContext()
}

type TableFactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableFactorContext() *TableFactorContext {
	var p = new(TableFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableFactor
	return p
}

func (*TableFactorContext) IsTableFactorContext() {}

func NewTableFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableFactorContext {
	var p = new(TableFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableFactor

	return p
}

func (s *TableFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *TableFactorContext) TableNameAndAlias() ITableNameAndAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasContext)
}

func (s *TableFactorContext) TableSubQuery() ITableSubQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSubQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSubQueryContext)
}

func (s *TableFactorContext) TableRecu() ITableRecuContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableRecuContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableRecuContext)
}

func (s *TableFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableFactor(s)
	}
}

func (s *TableFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableFactor(s)
	}
}




func (p *MySQLParser) TableFactor() (localctx ITableFactorContext) {
	localctx = NewTableFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MySQLParserRULE_tableFactor)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.TableNameAndAlias()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
			p.TableSubQuery()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.TableRecu()
		}

	}


	return localctx
}


// ITableSubQueryContext is an interface to support dynamic dispatch.
type ITableSubQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token 


	// SetAlias sets the alias token.
	SetAlias(antlr.Token) 


	// IsTableSubQueryContext differentiates from other interfaces.
	IsTableSubQueryContext()
}

type TableSubQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias antlr.Token
}

func NewEmptyTableSubQueryContext() *TableSubQueryContext {
	var p = new(TableSubQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableSubQuery
	return p
}

func (*TableSubQueryContext) IsTableSubQueryContext() {}

func NewTableSubQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSubQueryContext {
	var p = new(TableSubQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableSubQuery

	return p
}

func (s *TableSubQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSubQueryContext) GetAlias() antlr.Token { return s.alias }


func (s *TableSubQueryContext) SetAlias(v antlr.Token) { s.alias = v }


func (s *TableSubQueryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *TableSubQueryContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *TableSubQueryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *TableSubQueryContext) ID() antlr.TerminalNode {
	return s.GetToken(MySQLParserID, 0)
}

func (s *TableSubQueryContext) AS() antlr.TerminalNode {
	return s.GetToken(MySQLParserAS, 0)
}

func (s *TableSubQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSubQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableSubQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableSubQuery(s)
	}
}

func (s *TableSubQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableSubQuery(s)
	}
}




func (p *MySQLParser) TableSubQuery() (localctx ITableSubQueryContext) {
	localctx = NewTableSubQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MySQLParserRULE_tableSubQuery)
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
		p.SetState(275)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(276)
		p.SelectStat()
	}
	{
		p.SetState(277)
		p.Match(MySQLParserRPAREN)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserAS {
		{
			p.SetState(278)
			p.Match(MySQLParserAS)
		}

	}
	{
		p.SetState(281)

		var _m = p.Match(MySQLParserID)

		localctx.(*TableSubQueryContext).alias = _m
	}



	return localctx
}


// ITableRecuContext is an interface to support dynamic dispatch.
type ITableRecuContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableRecuContext differentiates from other interfaces.
	IsTableRecuContext()
}

type TableRecuContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableRecuContext() *TableRecuContext {
	var p = new(TableRecuContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableRecu
	return p
}

func (*TableRecuContext) IsTableRecuContext() {}

func NewTableRecuContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableRecuContext {
	var p = new(TableRecuContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableRecu

	return p
}

func (s *TableRecuContext) GetParser() antlr.Parser { return s.parser }

func (s *TableRecuContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *TableRecuContext) TableRel() ITableRelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableRelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableRelContext)
}

func (s *TableRecuContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *TableRecuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableRecuContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableRecuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableRecu(s)
	}
}

func (s *TableRecuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableRecu(s)
	}
}




func (p *MySQLParser) TableRecu() (localctx ITableRecuContext) {
	localctx = NewTableRecuContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MySQLParserRULE_tableRecu)

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
		p.SetState(283)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(284)
		p.TableRel()
	}
	{
		p.SetState(285)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// ITableJoinContext is an interface to support dynamic dispatch.
type ITableJoinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableJoinContext differentiates from other interfaces.
	IsTableJoinContext()
}

type TableJoinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableJoinContext() *TableJoinContext {
	var p = new(TableJoinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableJoin
	return p
}

func (*TableJoinContext) IsTableJoinContext() {}

func NewTableJoinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableJoinContext {
	var p = new(TableJoinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableJoin

	return p
}

func (s *TableJoinContext) GetParser() antlr.Parser { return s.parser }

func (s *TableJoinContext) TableNameAndAlias() ITableNameAndAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasContext)
}

func (s *TableJoinContext) TableJoinSuffix() ITableJoinSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableJoinSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableJoinSuffixContext)
}

func (s *TableJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableJoinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableJoin(s)
	}
}

func (s *TableJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableJoin(s)
	}
}




func (p *MySQLParser) TableJoin() (localctx ITableJoinContext) {
	localctx = NewTableJoinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MySQLParserRULE_tableJoin)

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
		p.SetState(287)
		p.TableNameAndAlias()
	}
	{
		p.SetState(288)
		p.TableJoinSuffix()
	}



	return localctx
}


// ITableJoinSuffixContext is an interface to support dynamic dispatch.
type ITableJoinSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableJoinSuffixContext differentiates from other interfaces.
	IsTableJoinSuffixContext()
}

type TableJoinSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableJoinSuffixContext() *TableJoinSuffixContext {
	var p = new(TableJoinSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableJoinSuffix
	return p
}

func (*TableJoinSuffixContext) IsTableJoinSuffixContext() {}

func NewTableJoinSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableJoinSuffixContext {
	var p = new(TableJoinSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableJoinSuffix

	return p
}

func (s *TableJoinSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *TableJoinSuffixContext) TableJoinMod() ITableJoinModContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableJoinModContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableJoinModContext)
}

func (s *TableJoinSuffixContext) JOIN() antlr.TerminalNode {
	return s.GetToken(MySQLParserJOIN, 0)
}

func (s *TableJoinSuffixContext) TableNameAndAliases() ITableNameAndAliasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasesContext)
}

func (s *TableJoinSuffixContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *TableJoinSuffixContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *TableJoinSuffixContext) TableRecu() ITableRecuContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableRecuContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableRecuContext)
}

func (s *TableJoinSuffixContext) JoinCondition() IJoinConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJoinConditionContext)
}

func (s *TableJoinSuffixContext) TableJoinSuffix() ITableJoinSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableJoinSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableJoinSuffixContext)
}

func (s *TableJoinSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableJoinSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableJoinSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableJoinSuffix(s)
	}
}

func (s *TableJoinSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableJoinSuffix(s)
	}
}




func (p *MySQLParser) TableJoinSuffix() (localctx ITableJoinSuffixContext) {
	localctx = NewTableJoinSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MySQLParserRULE_tableJoinSuffix)
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
		p.SetState(290)
		p.TableJoinMod()
	}
	{
		p.SetState(291)
		p.Match(MySQLParserJOIN)
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(292)
			p.TableNameAndAliases()
		}


	case 2:
		{
			p.SetState(293)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(294)
			p.TableNameAndAliases()
		}
		{
			p.SetState(295)
			p.Match(MySQLParserRPAREN)
		}


	case 3:
		{
			p.SetState(297)
			p.TableRecu()
		}

	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserON || _la == MySQLParserUSING {
		{
			p.SetState(300)
			p.JoinCondition()
		}

	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 29)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 29))) & ((1 << (MySQLParserINNER - 29)) | (1 << (MySQLParserLEFT - 29)) | (1 << (MySQLParserRIGHT - 29)) | (1 << (MySQLParserCROSS - 29)))) != 0) {
		{
			p.SetState(303)
			p.TableJoinSuffix()
		}

	}



	return localctx
}


// ITableJoinModContext is an interface to support dynamic dispatch.
type ITableJoinModContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableJoinModContext differentiates from other interfaces.
	IsTableJoinModContext()
}

type TableJoinModContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableJoinModContext() *TableJoinModContext {
	var p = new(TableJoinModContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableJoinMod
	return p
}

func (*TableJoinModContext) IsTableJoinModContext() {}

func NewTableJoinModContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableJoinModContext {
	var p = new(TableJoinModContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableJoinMod

	return p
}

func (s *TableJoinModContext) GetParser() antlr.Parser { return s.parser }

func (s *TableJoinModContext) INNER() antlr.TerminalNode {
	return s.GetToken(MySQLParserINNER, 0)
}

func (s *TableJoinModContext) CROSS() antlr.TerminalNode {
	return s.GetToken(MySQLParserCROSS, 0)
}

func (s *TableJoinModContext) LEFT() antlr.TerminalNode {
	return s.GetToken(MySQLParserLEFT, 0)
}

func (s *TableJoinModContext) OUTER() antlr.TerminalNode {
	return s.GetToken(MySQLParserOUTER, 0)
}

func (s *TableJoinModContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(MySQLParserRIGHT, 0)
}

func (s *TableJoinModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableJoinModContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableJoinModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableJoinMod(s)
	}
}

func (s *TableJoinModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableJoinMod(s)
	}
}




func (p *MySQLParser) TableJoinMod() (localctx ITableJoinModContext) {
	localctx = NewTableJoinModContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MySQLParserRULE_tableJoinMod)
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

	p.SetState(316)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySQLParserINNER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.Match(MySQLParserINNER)
		}


	case MySQLParserCROSS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Match(MySQLParserCROSS)
		}


	case MySQLParserLEFT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(308)
			p.Match(MySQLParserLEFT)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySQLParserOUTER {
			{
				p.SetState(309)
				p.Match(MySQLParserOUTER)
			}

		}


	case MySQLParserRIGHT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Match(MySQLParserRIGHT)
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySQLParserOUTER {
			{
				p.SetState(313)
				p.Match(MySQLParserOUTER)
			}

		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IJoinConditionContext is an interface to support dynamic dispatch.
type IJoinConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinConditionContext differentiates from other interfaces.
	IsJoinConditionContext()
}

type JoinConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinConditionContext() *JoinConditionContext {
	var p = new(JoinConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_joinCondition
	return p
}

func (*JoinConditionContext) IsJoinConditionContext() {}

func NewJoinConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinConditionContext {
	var p = new(JoinConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_joinCondition

	return p
}

func (s *JoinConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinConditionContext) ON() antlr.TerminalNode {
	return s.GetToken(MySQLParserON, 0)
}

func (s *JoinConditionContext) WhereCondition() IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *JoinConditionContext) USING() antlr.TerminalNode {
	return s.GetToken(MySQLParserUSING, 0)
}

func (s *JoinConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *JoinConditionContext) ColumnNames() IColumnNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNamesContext)
}

func (s *JoinConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *JoinConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *JoinConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterJoinCondition(s)
	}
}

func (s *JoinConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitJoinCondition(s)
	}
}




func (p *MySQLParser) JoinCondition() (localctx IJoinConditionContext) {
	localctx = NewJoinConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MySQLParserRULE_joinCondition)

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

	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySQLParserON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(MySQLParserON)
		}
		{
			p.SetState(319)
			p.WhereCondition()
		}


	case MySQLParserUSING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(MySQLParserUSING)
		}
		{
			p.SetState(321)
			p.Match(MySQLParserLPAREN)
		}
		{
			p.SetState(322)
			p.ColumnNames()
		}
		{
			p.SetState(323)
			p.Match(MySQLParserRPAREN)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IGbobExprsContext is an interface to support dynamic dispatch.
type IGbobExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSc returns the sc token.
	GetSc() antlr.Token 


	// SetSc sets the sc token.
	SetSc(antlr.Token) 


	// IsGbobExprsContext differentiates from other interfaces.
	IsGbobExprsContext()
}

type GbobExprsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sc antlr.Token
}

func NewEmptyGbobExprsContext() *GbobExprsContext {
	var p = new(GbobExprsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_gbobExprs
	return p
}

func (*GbobExprsContext) IsGbobExprsContext() {}

func NewGbobExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GbobExprsContext {
	var p = new(GbobExprsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_gbobExprs

	return p
}

func (s *GbobExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *GbobExprsContext) GetSc() antlr.Token { return s.sc }


func (s *GbobExprsContext) SetSc(v antlr.Token) { s.sc = v }


func (s *GbobExprsContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *GbobExprsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, 0)
}

func (s *GbobExprsContext) GbobExprs() IGbobExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGbobExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGbobExprsContext)
}

func (s *GbobExprsContext) ASC() antlr.TerminalNode {
	return s.GetToken(MySQLParserASC, 0)
}

func (s *GbobExprsContext) DESC() antlr.TerminalNode {
	return s.GetToken(MySQLParserDESC, 0)
}

func (s *GbobExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GbobExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GbobExprsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterGbobExprs(s)
	}
}

func (s *GbobExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitGbobExprs(s)
	}
}




func (p *MySQLParser) GbobExprs() (localctx IGbobExprsContext) {
	localctx = NewGbobExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MySQLParserRULE_gbobExprs)
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
		p.SetState(327)
		p.Element()
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserASC || _la == MySQLParserDESC {
		{
			p.SetState(328)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*GbobExprsContext).sc = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserASC || _la == MySQLParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*GbobExprsContext).sc = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserCOMMA {
		{
			p.SetState(331)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(332)
			p.GbobExprs()
		}

	}



	return localctx
}


// IUpdateStatContext is an interface to support dynamic dispatch.
type IUpdateStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateStatContext differentiates from other interfaces.
	IsUpdateStatContext()
}

type UpdateStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatContext() *UpdateStatContext {
	var p = new(UpdateStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_updateStat
	return p
}

func (*UpdateStatContext) IsUpdateStatContext() {}

func NewUpdateStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatContext {
	var p = new(UpdateStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_updateStat

	return p
}

func (s *UpdateStatContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatContext) UpdateSingleTable() IUpdateSingleTableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateSingleTableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateSingleTableContext)
}

func (s *UpdateStatContext) UpdateMultipleTable() IUpdateMultipleTableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateMultipleTableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateMultipleTableContext)
}

func (s *UpdateStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterUpdateStat(s)
	}
}

func (s *UpdateStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitUpdateStat(s)
	}
}




func (p *MySQLParser) UpdateStat() (localctx IUpdateStatContext) {
	localctx = NewUpdateStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MySQLParserRULE_updateStat)

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

	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.UpdateSingleTable()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.UpdateMultipleTable()
		}

	}


	return localctx
}


// IUpdateSingleTableContext is an interface to support dynamic dispatch.
type IUpdateSingleTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRowCount returns the rowCount token.
	GetRowCount() antlr.Token 


	// SetRowCount sets the rowCount token.
	SetRowCount(antlr.Token) 


	// IsUpdateSingleTableContext differentiates from other interfaces.
	IsUpdateSingleTableContext()
}

type UpdateSingleTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rowCount antlr.Token
}

func NewEmptyUpdateSingleTableContext() *UpdateSingleTableContext {
	var p = new(UpdateSingleTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_updateSingleTable
	return p
}

func (*UpdateSingleTableContext) IsUpdateSingleTableContext() {}

func NewUpdateSingleTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateSingleTableContext {
	var p = new(UpdateSingleTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_updateSingleTable

	return p
}

func (s *UpdateSingleTableContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateSingleTableContext) GetRowCount() antlr.Token { return s.rowCount }


func (s *UpdateSingleTableContext) SetRowCount(v antlr.Token) { s.rowCount = v }


func (s *UpdateSingleTableContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(MySQLParserUPDATE, 0)
}

func (s *UpdateSingleTableContext) TableNameAndAlias() ITableNameAndAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasContext)
}

func (s *UpdateSingleTableContext) SET() antlr.TerminalNode {
	return s.GetToken(MySQLParserSET, 0)
}

func (s *UpdateSingleTableContext) SetExprs() ISetExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetExprsContext)
}

func (s *UpdateSingleTableContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySQLParserWHERE, 0)
}

func (s *UpdateSingleTableContext) WhereCondition() IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *UpdateSingleTableContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MySQLParserLIMIT, 0)
}

func (s *UpdateSingleTableContext) INT() antlr.TerminalNode {
	return s.GetToken(MySQLParserINT, 0)
}

func (s *UpdateSingleTableContext) PLACEHOLDER() antlr.TerminalNode {
	return s.GetToken(MySQLParserPLACEHOLDER, 0)
}

func (s *UpdateSingleTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateSingleTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateSingleTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterUpdateSingleTable(s)
	}
}

func (s *UpdateSingleTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitUpdateSingleTable(s)
	}
}




func (p *MySQLParser) UpdateSingleTable() (localctx IUpdateSingleTableContext) {
	localctx = NewUpdateSingleTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MySQLParserRULE_updateSingleTable)
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
		p.SetState(339)
		p.Match(MySQLParserUPDATE)
	}
	{
		p.SetState(340)
		p.TableNameAndAlias()
	}
	{
		p.SetState(341)
		p.Match(MySQLParserSET)
	}
	{
		p.SetState(342)
		p.SetExprs()
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserWHERE {
		{
			p.SetState(343)
			p.Match(MySQLParserWHERE)
		}
		{
			p.SetState(344)
			p.WhereCondition()
		}

	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserLIMIT {
		{
			p.SetState(347)
			p.Match(MySQLParserLIMIT)
		}
		{
			p.SetState(348)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UpdateSingleTableContext).rowCount = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UpdateSingleTableContext).rowCount = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}



	return localctx
}


// IUpdateMultipleTableContext is an interface to support dynamic dispatch.
type IUpdateMultipleTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateMultipleTableContext differentiates from other interfaces.
	IsUpdateMultipleTableContext()
}

type UpdateMultipleTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateMultipleTableContext() *UpdateMultipleTableContext {
	var p = new(UpdateMultipleTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_updateMultipleTable
	return p
}

func (*UpdateMultipleTableContext) IsUpdateMultipleTableContext() {}

func NewUpdateMultipleTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateMultipleTableContext {
	var p = new(UpdateMultipleTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_updateMultipleTable

	return p
}

func (s *UpdateMultipleTableContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateMultipleTableContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(MySQLParserUPDATE, 0)
}

func (s *UpdateMultipleTableContext) TableNameAndAliases() ITableNameAndAliasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasesContext)
}

func (s *UpdateMultipleTableContext) SET() antlr.TerminalNode {
	return s.GetToken(MySQLParserSET, 0)
}

func (s *UpdateMultipleTableContext) SetExprs() ISetExprsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExprsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetExprsContext)
}

func (s *UpdateMultipleTableContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySQLParserWHERE, 0)
}

func (s *UpdateMultipleTableContext) WhereCondition() IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *UpdateMultipleTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateMultipleTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateMultipleTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterUpdateMultipleTable(s)
	}
}

func (s *UpdateMultipleTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitUpdateMultipleTable(s)
	}
}




func (p *MySQLParser) UpdateMultipleTable() (localctx IUpdateMultipleTableContext) {
	localctx = NewUpdateMultipleTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MySQLParserRULE_updateMultipleTable)
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
		p.SetState(351)
		p.Match(MySQLParserUPDATE)
	}
	{
		p.SetState(352)
		p.TableNameAndAliases()
	}
	{
		p.SetState(353)
		p.Match(MySQLParserSET)
	}
	{
		p.SetState(354)
		p.SetExprs()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserWHERE {
		{
			p.SetState(355)
			p.Match(MySQLParserWHERE)
		}
		{
			p.SetState(356)
			p.WhereCondition()
		}

	}



	return localctx
}


// ISetExprsContext is an interface to support dynamic dispatch.
type ISetExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetExprsContext differentiates from other interfaces.
	IsSetExprsContext()
}

type SetExprsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetExprsContext() *SetExprsContext {
	var p = new(SetExprsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_setExprs
	return p
}

func (*SetExprsContext) IsSetExprsContext() {}

func NewSetExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetExprsContext {
	var p = new(SetExprsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_setExprs

	return p
}

func (s *SetExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *SetExprsContext) AllSetExpr() []ISetExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISetExprContext)(nil)).Elem())
	var tst = make([]ISetExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISetExprContext)
		}
	}

	return tst
}

func (s *SetExprsContext) SetExpr(i int) ISetExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISetExprContext)
}

func (s *SetExprsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *SetExprsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *SetExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SetExprsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSetExprs(s)
	}
}

func (s *SetExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSetExprs(s)
	}
}




func (p *MySQLParser) SetExprs() (localctx ISetExprsContext) {
	localctx = NewSetExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MySQLParserRULE_setExprs)
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
		p.SetState(359)
		p.SetExpr()
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySQLParserCOMMA {
		{
			p.SetState(360)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(361)
			p.SetExpr()
		}


		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISetExprContext is an interface to support dynamic dispatch.
type ISetExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRightDefault returns the rightDefault token.
	GetRightDefault() antlr.Token 


	// SetRightDefault sets the rightDefault token.
	SetRightDefault(antlr.Token) 


	// GetLeft returns the left rule contexts.
	GetLeft() IElementContext

	// GetRight returns the right rule contexts.
	GetRight() IElementContext


	// SetLeft sets the left rule contexts.
	SetLeft(IElementContext)

	// SetRight sets the right rule contexts.
	SetRight(IElementContext)


	// IsSetExprContext differentiates from other interfaces.
	IsSetExprContext()
}

type SetExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left IElementContext 
	right IElementContext 
	rightDefault antlr.Token
}

func NewEmptySetExprContext() *SetExprContext {
	var p = new(SetExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_setExpr
	return p
}

func (*SetExprContext) IsSetExprContext() {}

func NewSetExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetExprContext {
	var p = new(SetExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_setExpr

	return p
}

func (s *SetExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SetExprContext) GetRightDefault() antlr.Token { return s.rightDefault }


func (s *SetExprContext) SetRightDefault(v antlr.Token) { s.rightDefault = v }


func (s *SetExprContext) GetLeft() IElementContext { return s.left }

func (s *SetExprContext) GetRight() IElementContext { return s.right }


func (s *SetExprContext) SetLeft(v IElementContext) { s.left = v }

func (s *SetExprContext) SetRight(v IElementContext) { s.right = v }


func (s *SetExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(MySQLParserEQ, 0)
}

func (s *SetExprContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *SetExprContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SetExprContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MySQLParserDEFAULT, 0)
}

func (s *SetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SetExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterSetExpr(s)
	}
}

func (s *SetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitSetExpr(s)
	}
}




func (p *MySQLParser) SetExpr() (localctx ISetExprContext) {
	localctx = NewSetExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, MySQLParserRULE_setExpr)

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
		p.SetState(367)

		var _x = p.Element()


		localctx.(*SetExprContext).left = _x
	}
	{
		p.SetState(368)
		p.Match(MySQLParserEQ)
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySQLParserPLACEHOLDER, MySQLParserNULL, MySQLParserTRUE, MySQLParserFALSE, MySQLParserDATE, MySQLParserTIME, MySQLParserTIMESTAMP, MySQLParserALL, MySQLParserANY, MySQLParserSOME, MySQLParserUNKNOWN, MySQLParserCASE, MySQLParserROW, MySQLParserBINARY, MySQLParserASTERISK, MySQLParserINT, MySQLParserDECIMAL, MySQLParserSTRING, MySQLParserID, MySQLParserCOLUMN_REL, MySQLParserLPAREN:
		{
			p.SetState(369)

			var _x = p.Element()


			localctx.(*SetExprContext).right = _x
		}


	case MySQLParserDEFAULT:
		{
			p.SetState(370)

			var _m = p.Match(MySQLParserDEFAULT)

			localctx.(*SetExprContext).rightDefault = _m
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}



	return localctx
}


// IDeleteStatContext is an interface to support dynamic dispatch.
type IDeleteStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRowCount returns the rowCount token.
	GetRowCount() antlr.Token 


	// SetRowCount sets the rowCount token.
	SetRowCount(antlr.Token) 


	// IsDeleteStatContext differentiates from other interfaces.
	IsDeleteStatContext()
}

type DeleteStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rowCount antlr.Token
}

func NewEmptyDeleteStatContext() *DeleteStatContext {
	var p = new(DeleteStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_deleteStat
	return p
}

func (*DeleteStatContext) IsDeleteStatContext() {}

func NewDeleteStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatContext {
	var p = new(DeleteStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_deleteStat

	return p
}

func (s *DeleteStatContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatContext) GetRowCount() antlr.Token { return s.rowCount }


func (s *DeleteStatContext) SetRowCount(v antlr.Token) { s.rowCount = v }


func (s *DeleteStatContext) DELETE() antlr.TerminalNode {
	return s.GetToken(MySQLParserDELETE, 0)
}

func (s *DeleteStatContext) FROM() antlr.TerminalNode {
	return s.GetToken(MySQLParserFROM, 0)
}

func (s *DeleteStatContext) TableNameAndAlias() ITableNameAndAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasContext)
}

func (s *DeleteStatContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySQLParserWHERE, 0)
}

func (s *DeleteStatContext) WhereCondition() IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *DeleteStatContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MySQLParserLIMIT, 0)
}

func (s *DeleteStatContext) INT() antlr.TerminalNode {
	return s.GetToken(MySQLParserINT, 0)
}

func (s *DeleteStatContext) PLACEHOLDER() antlr.TerminalNode {
	return s.GetToken(MySQLParserPLACEHOLDER, 0)
}

func (s *DeleteStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeleteStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterDeleteStat(s)
	}
}

func (s *DeleteStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitDeleteStat(s)
	}
}




func (p *MySQLParser) DeleteStat() (localctx IDeleteStatContext) {
	localctx = NewDeleteStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, MySQLParserRULE_deleteStat)
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
		p.SetState(373)
		p.Match(MySQLParserDELETE)
	}
	{
		p.SetState(374)
		p.Match(MySQLParserFROM)
	}
	{
		p.SetState(375)
		p.TableNameAndAlias()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserWHERE {
		{
			p.SetState(376)
			p.Match(MySQLParserWHERE)
		}
		{
			p.SetState(377)
			p.WhereCondition()
		}

	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserLIMIT {
		{
			p.SetState(380)
			p.Match(MySQLParserLIMIT)
		}
		{
			p.SetState(381)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeleteStatContext).rowCount = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserINT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeleteStatContext).rowCount = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}



	return localctx
}


// ITableNameAndAliasContext is an interface to support dynamic dispatch.
type ITableNameAndAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token 

	// GetAlias returns the alias token.
	GetAlias() antlr.Token 


	// SetName sets the name token.
	SetName(antlr.Token) 

	// SetAlias sets the alias token.
	SetAlias(antlr.Token) 


	// IsTableNameAndAliasContext differentiates from other interfaces.
	IsTableNameAndAliasContext()
}

type TableNameAndAliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name antlr.Token
	alias antlr.Token
}

func NewEmptyTableNameAndAliasContext() *TableNameAndAliasContext {
	var p = new(TableNameAndAliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableNameAndAlias
	return p
}

func (*TableNameAndAliasContext) IsTableNameAndAliasContext() {}

func NewTableNameAndAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameAndAliasContext {
	var p = new(TableNameAndAliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableNameAndAlias

	return p
}

func (s *TableNameAndAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameAndAliasContext) GetName() antlr.Token { return s.name }

func (s *TableNameAndAliasContext) GetAlias() antlr.Token { return s.alias }


func (s *TableNameAndAliasContext) SetName(v antlr.Token) { s.name = v }

func (s *TableNameAndAliasContext) SetAlias(v antlr.Token) { s.alias = v }


func (s *TableNameAndAliasContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserID)
}

func (s *TableNameAndAliasContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserID, i)
}

func (s *TableNameAndAliasContext) AS() antlr.TerminalNode {
	return s.GetToken(MySQLParserAS, 0)
}

func (s *TableNameAndAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameAndAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableNameAndAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableNameAndAlias(s)
	}
}

func (s *TableNameAndAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableNameAndAlias(s)
	}
}




func (p *MySQLParser) TableNameAndAlias() (localctx ITableNameAndAliasContext) {
	localctx = NewTableNameAndAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MySQLParserRULE_tableNameAndAlias)
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
		p.SetState(384)

		var _m = p.Match(MySQLParserID)

		localctx.(*TableNameAndAliasContext).name = _m
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserAS || _la == MySQLParserID {
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySQLParserAS {
			{
				p.SetState(385)
				p.Match(MySQLParserAS)
			}

		}
		{
			p.SetState(388)

			var _m = p.Match(MySQLParserID)

			localctx.(*TableNameAndAliasContext).alias = _m
		}

	}



	return localctx
}


// ITableNameAndAliasesContext is an interface to support dynamic dispatch.
type ITableNameAndAliasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameAndAliasesContext differentiates from other interfaces.
	IsTableNameAndAliasesContext()
}

type TableNameAndAliasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameAndAliasesContext() *TableNameAndAliasesContext {
	var p = new(TableNameAndAliasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_tableNameAndAliases
	return p
}

func (*TableNameAndAliasesContext) IsTableNameAndAliasesContext() {}

func NewTableNameAndAliasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameAndAliasesContext {
	var p = new(TableNameAndAliasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_tableNameAndAliases

	return p
}

func (s *TableNameAndAliasesContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameAndAliasesContext) AllTableNameAndAlias() []ITableNameAndAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem())
	var tst = make([]ITableNameAndAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableNameAndAliasContext)
		}
	}

	return tst
}

func (s *TableNameAndAliasesContext) TableNameAndAlias(i int) ITableNameAndAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameAndAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableNameAndAliasContext)
}

func (s *TableNameAndAliasesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *TableNameAndAliasesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *TableNameAndAliasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameAndAliasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableNameAndAliasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterTableNameAndAliases(s)
	}
}

func (s *TableNameAndAliasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitTableNameAndAliases(s)
	}
}




func (p *MySQLParser) TableNameAndAliases() (localctx ITableNameAndAliasesContext) {
	localctx = NewTableNameAndAliasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MySQLParserRULE_tableNameAndAliases)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.TableNameAndAlias()
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(392)
				p.Match(MySQLParserCOMMA)
			}
			{
				p.SetState(393)
				p.TableNameAndAlias()
			}


		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}



	return localctx
}


// IWhereConditionContext is an interface to support dynamic dispatch.
type IWhereConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereConditionContext differentiates from other interfaces.
	IsWhereConditionContext()
}

type WhereConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereConditionContext() *WhereConditionContext {
	var p = new(WhereConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_whereCondition
	return p
}

func (*WhereConditionContext) IsWhereConditionContext() {}

func NewWhereConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereConditionContext {
	var p = new(WhereConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_whereCondition

	return p
}

func (s *WhereConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereConditionContext) WhereCondSub() IWhereCondSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereCondSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereCondSubContext)
}

func (s *WhereConditionContext) WhereCondOp() IWhereCondOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereCondOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereCondOpContext)
}

func (s *WhereConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterWhereCondition(s)
	}
}

func (s *WhereConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitWhereCondition(s)
	}
}




func (p *MySQLParser) WhereCondition() (localctx IWhereConditionContext) {
	localctx = NewWhereConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MySQLParserRULE_whereCondition)

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

	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			p.WhereCondSub()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.WhereCondOp()
		}

	}


	return localctx
}


// IWhereCondSubContext is an interface to support dynamic dispatch.
type IWhereCondSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpressionOperator returns the expressionOperator token.
	GetExpressionOperator() antlr.Token 


	// SetExpressionOperator sets the expressionOperator token.
	SetExpressionOperator(antlr.Token) 


	// GetWc1 returns the wc1 rule contexts.
	GetWc1() IWhereConditionContext

	// GetWc2 returns the wc2 rule contexts.
	GetWc2() IWhereConditionContext


	// SetWc1 sets the wc1 rule contexts.
	SetWc1(IWhereConditionContext)

	// SetWc2 sets the wc2 rule contexts.
	SetWc2(IWhereConditionContext)


	// IsWhereCondSubContext differentiates from other interfaces.
	IsWhereCondSubContext()
}

type WhereCondSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	wc1 IWhereConditionContext 
	expressionOperator antlr.Token
	wc2 IWhereConditionContext 
}

func NewEmptyWhereCondSubContext() *WhereCondSubContext {
	var p = new(WhereCondSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_whereCondSub
	return p
}

func (*WhereCondSubContext) IsWhereCondSubContext() {}

func NewWhereCondSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereCondSubContext {
	var p = new(WhereCondSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_whereCondSub

	return p
}

func (s *WhereCondSubContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereCondSubContext) GetExpressionOperator() antlr.Token { return s.expressionOperator }


func (s *WhereCondSubContext) SetExpressionOperator(v antlr.Token) { s.expressionOperator = v }


func (s *WhereCondSubContext) GetWc1() IWhereConditionContext { return s.wc1 }

func (s *WhereCondSubContext) GetWc2() IWhereConditionContext { return s.wc2 }


func (s *WhereCondSubContext) SetWc1(v IWhereConditionContext) { s.wc1 = v }

func (s *WhereCondSubContext) SetWc2(v IWhereConditionContext) { s.wc2 = v }


func (s *WhereCondSubContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *WhereCondSubContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *WhereCondSubContext) AllWhereCondition() []IWhereConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem())
	var tst = make([]IWhereConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereConditionContext)
		}
	}

	return tst
}

func (s *WhereCondSubContext) WhereCondition(i int) IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *WhereCondSubContext) AND() antlr.TerminalNode {
	return s.GetToken(MySQLParserAND, 0)
}

func (s *WhereCondSubContext) XOR() antlr.TerminalNode {
	return s.GetToken(MySQLParserXOR, 0)
}

func (s *WhereCondSubContext) OR() antlr.TerminalNode {
	return s.GetToken(MySQLParserOR, 0)
}

func (s *WhereCondSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereCondSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereCondSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterWhereCondSub(s)
	}
}

func (s *WhereCondSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitWhereCondSub(s)
	}
}




func (p *MySQLParser) WhereCondSub() (localctx IWhereCondSubContext) {
	localctx = NewWhereCondSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MySQLParserRULE_whereCondSub)
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
		p.SetState(403)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(404)

		var _x = p.WhereCondition()


		localctx.(*WhereCondSubContext).wc1 = _x
	}
	{
		p.SetState(405)
		p.Match(MySQLParserRPAREN)
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserXOR || _la == MySQLParserAND || _la == MySQLParserOR {
		{
			p.SetState(406)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*WhereCondSubContext).expressionOperator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserXOR || _la == MySQLParserAND || _la == MySQLParserOR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*WhereCondSubContext).expressionOperator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(407)

			var _x = p.WhereCondition()


			localctx.(*WhereCondSubContext).wc2 = _x
		}

	}



	return localctx
}


// IWhereCondOpContext is an interface to support dynamic dispatch.
type IWhereCondOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpressionOperator returns the expressionOperator token.
	GetExpressionOperator() antlr.Token 


	// SetExpressionOperator sets the expressionOperator token.
	SetExpressionOperator(antlr.Token) 


	// IsWhereCondOpContext differentiates from other interfaces.
	IsWhereCondOpContext()
}

type WhereCondOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expressionOperator antlr.Token
}

func NewEmptyWhereCondOpContext() *WhereCondOpContext {
	var p = new(WhereCondOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_whereCondOp
	return p
}

func (*WhereCondOpContext) IsWhereCondOpContext() {}

func NewWhereCondOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereCondOpContext {
	var p = new(WhereCondOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_whereCondOp

	return p
}

func (s *WhereCondOpContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereCondOpContext) GetExpressionOperator() antlr.Token { return s.expressionOperator }


func (s *WhereCondOpContext) SetExpressionOperator(v antlr.Token) { s.expressionOperator = v }


func (s *WhereCondOpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhereCondOpContext) WhereCondition() IWhereConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereConditionContext)
}

func (s *WhereCondOpContext) AND() antlr.TerminalNode {
	return s.GetToken(MySQLParserAND, 0)
}

func (s *WhereCondOpContext) XOR() antlr.TerminalNode {
	return s.GetToken(MySQLParserXOR, 0)
}

func (s *WhereCondOpContext) OR() antlr.TerminalNode {
	return s.GetToken(MySQLParserOR, 0)
}

func (s *WhereCondOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereCondOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *WhereCondOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterWhereCondOp(s)
	}
}

func (s *WhereCondOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitWhereCondOp(s)
	}
}




func (p *MySQLParser) WhereCondOp() (localctx IWhereCondOpContext) {
	localctx = NewWhereCondOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MySQLParserRULE_whereCondOp)
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
		p.SetState(410)
		p.Expression()
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserXOR || _la == MySQLParserAND || _la == MySQLParserOR {
		{
			p.SetState(411)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*WhereCondOpContext).expressionOperator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserXOR || _la == MySQLParserAND || _la == MySQLParserOR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*WhereCondOpContext).expressionOperator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(412)
			p.WhereCondition()
		}

	}



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ExprRelational() IExprRelationalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprRelationalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprRelationalContext)
}

func (s *ExpressionContext) ExprBetweenAnd() IExprBetweenAndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprBetweenAndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprBetweenAndContext)
}

func (s *ExpressionContext) ExprIsOrIsNotNull() IExprIsOrIsNotNullContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprIsOrIsNotNullContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprIsOrIsNotNullContext)
}

func (s *ExpressionContext) ExprInSelect() IExprInSelectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprInSelectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprInSelectContext)
}

func (s *ExpressionContext) ExprInValues() IExprInValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprInValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprInValuesContext)
}

func (s *ExpressionContext) ExprExists() IExprExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprExistsContext)
}

func (s *ExpressionContext) ExprNot() IExprNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprNotContext)
}

func (s *ExpressionContext) ExprLike() IExprLikeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprLikeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprLikeContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExpression(s)
	}
}




func (p *MySQLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MySQLParserRULE_expression)

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

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			p.ExprRelational()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.ExprBetweenAnd()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(417)
			p.ExprIsOrIsNotNull()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(418)
			p.ExprInSelect()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(419)
			p.ExprInValues()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(420)
			p.ExprExists()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(421)
			p.ExprNot()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(422)
			p.ExprLike()
		}

	}


	return localctx
}


// IExprRelationalContext is an interface to support dynamic dispatch.
type IExprRelationalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRelationalOp returns the relationalOp token.
	GetRelationalOp() antlr.Token 


	// SetRelationalOp sets the relationalOp token.
	SetRelationalOp(antlr.Token) 


	// GetLeft returns the left rule contexts.
	GetLeft() IElementContext

	// GetRight returns the right rule contexts.
	GetRight() IElementContext


	// SetLeft sets the left rule contexts.
	SetLeft(IElementContext)

	// SetRight sets the right rule contexts.
	SetRight(IElementContext)


	// IsExprRelationalContext differentiates from other interfaces.
	IsExprRelationalContext()
}

type ExprRelationalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left IElementContext 
	relationalOp antlr.Token
	right IElementContext 
}

func NewEmptyExprRelationalContext() *ExprRelationalContext {
	var p = new(ExprRelationalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprRelational
	return p
}

func (*ExprRelationalContext) IsExprRelationalContext() {}

func NewExprRelationalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprRelationalContext {
	var p = new(ExprRelationalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprRelational

	return p
}

func (s *ExprRelationalContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprRelationalContext) GetRelationalOp() antlr.Token { return s.relationalOp }


func (s *ExprRelationalContext) SetRelationalOp(v antlr.Token) { s.relationalOp = v }


func (s *ExprRelationalContext) GetLeft() IElementContext { return s.left }

func (s *ExprRelationalContext) GetRight() IElementContext { return s.right }


func (s *ExprRelationalContext) SetLeft(v IElementContext) { s.left = v }

func (s *ExprRelationalContext) SetRight(v IElementContext) { s.right = v }


func (s *ExprRelationalContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ExprRelationalContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprRelationalContext) EQ() antlr.TerminalNode {
	return s.GetToken(MySQLParserEQ, 0)
}

func (s *ExprRelationalContext) LTH() antlr.TerminalNode {
	return s.GetToken(MySQLParserLTH, 0)
}

func (s *ExprRelationalContext) GTH() antlr.TerminalNode {
	return s.GetToken(MySQLParserGTH, 0)
}

func (s *ExprRelationalContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT_EQ, 0)
}

func (s *ExprRelationalContext) LET() antlr.TerminalNode {
	return s.GetToken(MySQLParserLET, 0)
}

func (s *ExprRelationalContext) GET() antlr.TerminalNode {
	return s.GetToken(MySQLParserGET, 0)
}

func (s *ExprRelationalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprRelationalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprRelationalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprRelational(s)
	}
}

func (s *ExprRelationalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprRelational(s)
	}
}




func (p *MySQLParser) ExprRelational() (localctx IExprRelationalContext) {
	localctx = NewExprRelationalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MySQLParserRULE_exprRelational)
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
		p.SetState(425)

		var _x = p.Element()


		localctx.(*ExprRelationalContext).left = _x
	}
	{
		p.SetState(426)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ExprRelationalContext).relationalOp = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 88)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 88))) & ((1 << (MySQLParserEQ - 88)) | (1 << (MySQLParserLTH - 88)) | (1 << (MySQLParserGTH - 88)) | (1 << (MySQLParserNOT_EQ - 88)) | (1 << (MySQLParserLET - 88)) | (1 << (MySQLParserGET - 88)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ExprRelationalContext).relationalOp = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(427)

		var _x = p.Element()


		localctx.(*ExprRelationalContext).right = _x
	}



	return localctx
}


// IExprBetweenAndContext is an interface to support dynamic dispatch.
type IExprBetweenAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 


	// GetEl returns the el rule contexts.
	GetEl() IElementContext

	// GetLeft returns the left rule contexts.
	GetLeft() IElementContext

	// GetRight returns the right rule contexts.
	GetRight() IElementContext


	// SetEl sets the el rule contexts.
	SetEl(IElementContext)

	// SetLeft sets the left rule contexts.
	SetLeft(IElementContext)

	// SetRight sets the right rule contexts.
	SetRight(IElementContext)


	// IsExprBetweenAndContext differentiates from other interfaces.
	IsExprBetweenAndContext()
}

type ExprBetweenAndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	el IElementContext 
	not antlr.Token
	left IElementContext 
	right IElementContext 
}

func NewEmptyExprBetweenAndContext() *ExprBetweenAndContext {
	var p = new(ExprBetweenAndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprBetweenAnd
	return p
}

func (*ExprBetweenAndContext) IsExprBetweenAndContext() {}

func NewExprBetweenAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprBetweenAndContext {
	var p = new(ExprBetweenAndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprBetweenAnd

	return p
}

func (s *ExprBetweenAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprBetweenAndContext) GetNot() antlr.Token { return s.not }


func (s *ExprBetweenAndContext) SetNot(v antlr.Token) { s.not = v }


func (s *ExprBetweenAndContext) GetEl() IElementContext { return s.el }

func (s *ExprBetweenAndContext) GetLeft() IElementContext { return s.left }

func (s *ExprBetweenAndContext) GetRight() IElementContext { return s.right }


func (s *ExprBetweenAndContext) SetEl(v IElementContext) { s.el = v }

func (s *ExprBetweenAndContext) SetLeft(v IElementContext) { s.left = v }

func (s *ExprBetweenAndContext) SetRight(v IElementContext) { s.right = v }


func (s *ExprBetweenAndContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(MySQLParserBETWEEN, 0)
}

func (s *ExprBetweenAndContext) AND() antlr.TerminalNode {
	return s.GetToken(MySQLParserAND, 0)
}

func (s *ExprBetweenAndContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ExprBetweenAndContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprBetweenAndContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprBetweenAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprBetweenAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprBetweenAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprBetweenAnd(s)
	}
}

func (s *ExprBetweenAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprBetweenAnd(s)
	}
}




func (p *MySQLParser) ExprBetweenAnd() (localctx IExprBetweenAndContext) {
	localctx = NewExprBetweenAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MySQLParserRULE_exprBetweenAnd)
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
		p.SetState(429)

		var _x = p.Element()


		localctx.(*ExprBetweenAndContext).el = _x
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(430)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprBetweenAndContext).not = _m
		}

	}
	{
		p.SetState(433)
		p.Match(MySQLParserBETWEEN)
	}
	{
		p.SetState(434)

		var _x = p.Element()


		localctx.(*ExprBetweenAndContext).left = _x
	}
	{
		p.SetState(435)
		p.Match(MySQLParserAND)
	}
	{
		p.SetState(436)

		var _x = p.Element()


		localctx.(*ExprBetweenAndContext).right = _x
	}



	return localctx
}


// IExprIsOrIsNotNullContext is an interface to support dynamic dispatch.
type IExprIsOrIsNotNullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 

	// GetWhat returns the what token.
	GetWhat() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 

	// SetWhat sets the what token.
	SetWhat(antlr.Token) 


	// IsExprIsOrIsNotNullContext differentiates from other interfaces.
	IsExprIsOrIsNotNullContext()
}

type ExprIsOrIsNotNullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	not antlr.Token
	what antlr.Token
}

func NewEmptyExprIsOrIsNotNullContext() *ExprIsOrIsNotNullContext {
	var p = new(ExprIsOrIsNotNullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprIsOrIsNotNull
	return p
}

func (*ExprIsOrIsNotNullContext) IsExprIsOrIsNotNullContext() {}

func NewExprIsOrIsNotNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprIsOrIsNotNullContext {
	var p = new(ExprIsOrIsNotNullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprIsOrIsNotNull

	return p
}

func (s *ExprIsOrIsNotNullContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprIsOrIsNotNullContext) GetNot() antlr.Token { return s.not }

func (s *ExprIsOrIsNotNullContext) GetWhat() antlr.Token { return s.what }


func (s *ExprIsOrIsNotNullContext) SetNot(v antlr.Token) { s.not = v }

func (s *ExprIsOrIsNotNullContext) SetWhat(v antlr.Token) { s.what = v }


func (s *ExprIsOrIsNotNullContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprIsOrIsNotNullContext) IS() antlr.TerminalNode {
	return s.GetToken(MySQLParserIS, 0)
}

func (s *ExprIsOrIsNotNullContext) NULL() antlr.TerminalNode {
	return s.GetToken(MySQLParserNULL, 0)
}

func (s *ExprIsOrIsNotNullContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MySQLParserTRUE, 0)
}

func (s *ExprIsOrIsNotNullContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MySQLParserFALSE, 0)
}

func (s *ExprIsOrIsNotNullContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(MySQLParserUNKNOWN, 0)
}

func (s *ExprIsOrIsNotNullContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprIsOrIsNotNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIsOrIsNotNullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprIsOrIsNotNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprIsOrIsNotNull(s)
	}
}

func (s *ExprIsOrIsNotNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprIsOrIsNotNull(s)
	}
}




func (p *MySQLParser) ExprIsOrIsNotNull() (localctx IExprIsOrIsNotNullContext) {
	localctx = NewExprIsOrIsNotNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MySQLParserRULE_exprIsOrIsNotNull)
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
		p.SetState(438)
		p.Element()
	}
	{
		p.SetState(439)
		p.Match(MySQLParserIS)
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(440)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprIsOrIsNotNullContext).not = _m
		}

	}
	{
		p.SetState(443)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ExprIsOrIsNotNullContext).what = _lt

		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MySQLParserNULL) | (1 << MySQLParserTRUE) | (1 << MySQLParserFALSE))) != 0) || _la == MySQLParserUNKNOWN) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ExprIsOrIsNotNullContext).what = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IExprInSelectContext is an interface to support dynamic dispatch.
type IExprInSelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 


	// IsExprInSelectContext differentiates from other interfaces.
	IsExprInSelectContext()
}

type ExprInSelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	not antlr.Token
}

func NewEmptyExprInSelectContext() *ExprInSelectContext {
	var p = new(ExprInSelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprInSelect
	return p
}

func (*ExprInSelectContext) IsExprInSelectContext() {}

func NewExprInSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprInSelectContext {
	var p = new(ExprInSelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprInSelect

	return p
}

func (s *ExprInSelectContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprInSelectContext) GetNot() antlr.Token { return s.not }


func (s *ExprInSelectContext) SetNot(v antlr.Token) { s.not = v }


func (s *ExprInSelectContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprInSelectContext) IN() antlr.TerminalNode {
	return s.GetToken(MySQLParserIN, 0)
}

func (s *ExprInSelectContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ExprInSelectContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *ExprInSelectContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ExprInSelectContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprInSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprInSelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprInSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprInSelect(s)
	}
}

func (s *ExprInSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprInSelect(s)
	}
}




func (p *MySQLParser) ExprInSelect() (localctx IExprInSelectContext) {
	localctx = NewExprInSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MySQLParserRULE_exprInSelect)
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
		p.SetState(445)
		p.Element()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(446)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprInSelectContext).not = _m
		}

	}
	{
		p.SetState(449)
		p.Match(MySQLParserIN)
	}
	{
		p.SetState(450)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(451)
		p.SelectStat()
	}
	{
		p.SetState(452)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IExprInValuesContext is an interface to support dynamic dispatch.
type IExprInValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 


	// IsExprInValuesContext differentiates from other interfaces.
	IsExprInValuesContext()
}

type ExprInValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	not antlr.Token
}

func NewEmptyExprInValuesContext() *ExprInValuesContext {
	var p = new(ExprInValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprInValues
	return p
}

func (*ExprInValuesContext) IsExprInValuesContext() {}

func NewExprInValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprInValuesContext {
	var p = new(ExprInValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprInValues

	return p
}

func (s *ExprInValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprInValuesContext) GetNot() antlr.Token { return s.not }


func (s *ExprInValuesContext) SetNot(v antlr.Token) { s.not = v }


func (s *ExprInValuesContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprInValuesContext) IN() antlr.TerminalNode {
	return s.GetToken(MySQLParserIN, 0)
}

func (s *ExprInValuesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ExprInValuesContext) ValueList() IValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *ExprInValuesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ExprInValuesContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprInValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprInValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprInValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprInValues(s)
	}
}

func (s *ExprInValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprInValues(s)
	}
}




func (p *MySQLParser) ExprInValues() (localctx IExprInValuesContext) {
	localctx = NewExprInValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MySQLParserRULE_exprInValues)
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
		p.SetState(454)
		p.Element()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(455)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprInValuesContext).not = _m
		}

	}
	{
		p.SetState(458)
		p.Match(MySQLParserIN)
	}
	{
		p.SetState(459)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(460)
		p.ValueList()
	}
	{
		p.SetState(461)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IExprExistsContext is an interface to support dynamic dispatch.
type IExprExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 


	// IsExprExistsContext differentiates from other interfaces.
	IsExprExistsContext()
}

type ExprExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	not antlr.Token
}

func NewEmptyExprExistsContext() *ExprExistsContext {
	var p = new(ExprExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprExists
	return p
}

func (*ExprExistsContext) IsExprExistsContext() {}

func NewExprExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprExistsContext {
	var p = new(ExprExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprExists

	return p
}

func (s *ExprExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprExistsContext) GetNot() antlr.Token { return s.not }


func (s *ExprExistsContext) SetNot(v antlr.Token) { s.not = v }


func (s *ExprExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(MySQLParserEXISTS, 0)
}

func (s *ExprExistsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ExprExistsContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *ExprExistsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ExprExistsContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprExists(s)
	}
}

func (s *ExprExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprExists(s)
	}
}




func (p *MySQLParser) ExprExists() (localctx IExprExistsContext) {
	localctx = NewExprExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MySQLParserRULE_exprExists)
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
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(463)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprExistsContext).not = _m
		}

	}
	{
		p.SetState(466)
		p.Match(MySQLParserEXISTS)
	}
	{
		p.SetState(467)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(468)
		p.SelectStat()
	}
	{
		p.SetState(469)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IExprNotContext is an interface to support dynamic dispatch.
type IExprNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprNotContext differentiates from other interfaces.
	IsExprNotContext()
}

type ExprNotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprNotContext() *ExprNotContext {
	var p = new(ExprNotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprNot
	return p
}

func (*ExprNotContext) IsExprNotContext() {}

func NewExprNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprNotContext {
	var p = new(ExprNotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprNot

	return p
}

func (s *ExprNotContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprNotContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprNotContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprNot(s)
	}
}

func (s *ExprNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprNot(s)
	}
}




func (p *MySQLParser) ExprNot() (localctx IExprNotContext) {
	localctx = NewExprNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, MySQLParserRULE_exprNot)
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
		p.SetState(471)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySQLParserT__0 || _la == MySQLParserNOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(472)
		p.Expression()
	}



	return localctx
}


// IExprLikeContext is an interface to support dynamic dispatch.
type IExprLikeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token 


	// SetNot sets the not token.
	SetNot(antlr.Token) 


	// GetLeft returns the left rule contexts.
	GetLeft() IElementContext

	// GetRight returns the right rule contexts.
	GetRight() IElementContext


	// SetLeft sets the left rule contexts.
	SetLeft(IElementContext)

	// SetRight sets the right rule contexts.
	SetRight(IElementContext)


	// IsExprLikeContext differentiates from other interfaces.
	IsExprLikeContext()
}

type ExprLikeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left IElementContext 
	not antlr.Token
	right IElementContext 
}

func NewEmptyExprLikeContext() *ExprLikeContext {
	var p = new(ExprLikeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_exprLike
	return p
}

func (*ExprLikeContext) IsExprLikeContext() {}

func NewExprLikeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprLikeContext {
	var p = new(ExprLikeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_exprLike

	return p
}

func (s *ExprLikeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprLikeContext) GetNot() antlr.Token { return s.not }


func (s *ExprLikeContext) SetNot(v antlr.Token) { s.not = v }


func (s *ExprLikeContext) GetLeft() IElementContext { return s.left }

func (s *ExprLikeContext) GetRight() IElementContext { return s.right }


func (s *ExprLikeContext) SetLeft(v IElementContext) { s.left = v }

func (s *ExprLikeContext) SetRight(v IElementContext) { s.right = v }


func (s *ExprLikeContext) LIKE() antlr.TerminalNode {
	return s.GetToken(MySQLParserLIKE, 0)
}

func (s *ExprLikeContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ExprLikeContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ExprLikeContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLParserNOT, 0)
}

func (s *ExprLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLikeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprLikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterExprLike(s)
	}
}

func (s *ExprLikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitExprLike(s)
	}
}




func (p *MySQLParser) ExprLike() (localctx IExprLikeContext) {
	localctx = NewExprLikeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, MySQLParserRULE_exprLike)
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
		p.SetState(474)

		var _x = p.Element()


		localctx.(*ExprLikeContext).left = _x
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserNOT {
		{
			p.SetState(475)

			var _m = p.Match(MySQLParserNOT)

			localctx.(*ExprLikeContext).not = _m
		}

	}
	{
		p.SetState(478)
		p.Match(MySQLParserLIKE)
	}
	{
		p.SetState(479)

		var _x = p.Element()


		localctx.(*ExprLikeContext).right = _x
	}



	return localctx
}


// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) ElementOpFactory() IElementOpFactoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpFactoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpFactoryContext)
}

func (s *ElementContext) ElementListFactor() IElementListFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementListFactorContext)
}

func (s *ElementContext) ElementOpEle() IElementOpEleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpEleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpEleContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElement(s)
	}
}




func (p *MySQLParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, MySQLParserRULE_element)

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

	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(481)
			p.ElementOpFactory()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)
			p.ElementListFactor()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(483)
			p.ElementOpEle()
		}

	}


	return localctx
}


// IElementOpFactoryContext is an interface to support dynamic dispatch.
type IElementOpFactoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementOpFactoryContext differentiates from other interfaces.
	IsElementOpFactoryContext()
}

type ElementOpFactoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementOpFactoryContext() *ElementOpFactoryContext {
	var p = new(ElementOpFactoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementOpFactory
	return p
}

func (*ElementOpFactoryContext) IsElementOpFactoryContext() {}

func NewElementOpFactoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementOpFactoryContext {
	var p = new(ElementOpFactoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementOpFactory

	return p
}

func (s *ElementOpFactoryContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementOpFactoryContext) ElementText() IElementTextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementTextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementTextContext)
}

func (s *ElementOpFactoryContext) ElementTextParam() IElementTextParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementTextParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementTextParamContext)
}

func (s *ElementOpFactoryContext) ElementDate() IElementDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementDateContext)
}

func (s *ElementOpFactoryContext) FunCall() IFunCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunCallContext)
}

func (s *ElementOpFactoryContext) ElementSubQuery() IElementSubQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementSubQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementSubQueryContext)
}

func (s *ElementOpFactoryContext) ElementCase() IElementCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementCaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementCaseContext)
}

func (s *ElementOpFactoryContext) ElementWapperBkt() IElementWapperBktContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementWapperBktContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementWapperBktContext)
}

func (s *ElementOpFactoryContext) ElementWithPrefix() IElementWithPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementWithPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementWithPrefixContext)
}

func (s *ElementOpFactoryContext) ElementRow() IElementRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementRowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementRowContext)
}

func (s *ElementOpFactoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementOpFactoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementOpFactoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementOpFactory(s)
	}
}

func (s *ElementOpFactoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementOpFactory(s)
	}
}




func (p *MySQLParser) ElementOpFactory() (localctx IElementOpFactoryContext) {
	localctx = NewElementOpFactoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, MySQLParserRULE_elementOpFactory)

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

	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.ElementText()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.ElementTextParam()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(488)
			p.ElementDate()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(489)
			p.FunCall()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(490)
			p.ElementSubQuery()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(491)
			p.ElementCase()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(492)
			p.ElementWapperBkt()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(493)
			p.ElementWithPrefix()
		}


	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(494)
			p.ElementRow()
		}

	}


	return localctx
}


// IElementTextContext is an interface to support dynamic dispatch.
type IElementTextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementTextContext differentiates from other interfaces.
	IsElementTextContext()
}

type ElementTextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementTextContext() *ElementTextContext {
	var p = new(ElementTextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementText
	return p
}

func (*ElementTextContext) IsElementTextContext() {}

func NewElementTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementTextContext {
	var p = new(ElementTextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementText

	return p
}

func (s *ElementTextContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementTextContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MySQLParserASTERISK, 0)
}

func (s *ElementTextContext) PLACEHOLDER() antlr.TerminalNode {
	return s.GetToken(MySQLParserPLACEHOLDER, 0)
}

func (s *ElementTextContext) COLUMN_REL() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOLUMN_REL, 0)
}

func (s *ElementTextContext) ID() antlr.TerminalNode {
	return s.GetToken(MySQLParserID, 0)
}

func (s *ElementTextContext) NULL() antlr.TerminalNode {
	return s.GetToken(MySQLParserNULL, 0)
}

func (s *ElementTextContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(MySQLParserUNKNOWN, 0)
}

func (s *ElementTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementTextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementText(s)
	}
}

func (s *ElementTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementText(s)
	}
}




func (p *MySQLParser) ElementText() (localctx IElementTextContext) {
	localctx = NewElementTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, MySQLParserRULE_elementText)
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
		p.SetState(497)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySQLParserPLACEHOLDER || _la == MySQLParserNULL || ((((_la - 48)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 48))) & ((1 << (MySQLParserUNKNOWN - 48)) | (1 << (MySQLParserASTERISK - 48)) | (1 << (MySQLParserID - 48)) | (1 << (MySQLParserCOLUMN_REL - 48)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IElementTextParamContext is an interface to support dynamic dispatch.
type IElementTextParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementTextParamContext differentiates from other interfaces.
	IsElementTextParamContext()
}

type ElementTextParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementTextParamContext() *ElementTextParamContext {
	var p = new(ElementTextParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementTextParam
	return p
}

func (*ElementTextParamContext) IsElementTextParamContext() {}

func NewElementTextParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementTextParamContext {
	var p = new(ElementTextParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementTextParam

	return p
}

func (s *ElementTextParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementTextParamContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(MySQLParserDECIMAL, 0)
}

func (s *ElementTextParamContext) STRING() antlr.TerminalNode {
	return s.GetToken(MySQLParserSTRING, 0)
}

func (s *ElementTextParamContext) INT() antlr.TerminalNode {
	return s.GetToken(MySQLParserINT, 0)
}

func (s *ElementTextParamContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MySQLParserTRUE, 0)
}

func (s *ElementTextParamContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MySQLParserFALSE, 0)
}

func (s *ElementTextParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementTextParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementTextParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementTextParam(s)
	}
}

func (s *ElementTextParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementTextParam(s)
	}
}




func (p *MySQLParser) ElementTextParam() (localctx IElementTextParamContext) {
	localctx = NewElementTextParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, MySQLParserRULE_elementTextParam)
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
		p.SetState(499)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySQLParserTRUE || _la == MySQLParserFALSE || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (MySQLParserINT - 74)) | (1 << (MySQLParserDECIMAL - 74)) | (1 << (MySQLParserSTRING - 74)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IElementDateContext is an interface to support dynamic dispatch.
type IElementDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDt returns the dt token.
	GetDt() antlr.Token 


	// SetDt sets the dt token.
	SetDt(antlr.Token) 


	// IsElementDateContext differentiates from other interfaces.
	IsElementDateContext()
}

type ElementDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dt antlr.Token
}

func NewEmptyElementDateContext() *ElementDateContext {
	var p = new(ElementDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementDate
	return p
}

func (*ElementDateContext) IsElementDateContext() {}

func NewElementDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementDateContext {
	var p = new(ElementDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementDate

	return p
}

func (s *ElementDateContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementDateContext) GetDt() antlr.Token { return s.dt }


func (s *ElementDateContext) SetDt(v antlr.Token) { s.dt = v }


func (s *ElementDateContext) STRING() antlr.TerminalNode {
	return s.GetToken(MySQLParserSTRING, 0)
}

func (s *ElementDateContext) DATE() antlr.TerminalNode {
	return s.GetToken(MySQLParserDATE, 0)
}

func (s *ElementDateContext) TIME() antlr.TerminalNode {
	return s.GetToken(MySQLParserTIME, 0)
}

func (s *ElementDateContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MySQLParserTIMESTAMP, 0)
}

func (s *ElementDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementDate(s)
	}
}

func (s *ElementDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementDate(s)
	}
}




func (p *MySQLParser) ElementDate() (localctx IElementDateContext) {
	localctx = NewElementDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, MySQLParserRULE_elementDate)
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
		p.SetState(501)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ElementDateContext).dt = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 41)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 41))) & ((1 << (MySQLParserDATE - 41)) | (1 << (MySQLParserTIME - 41)) | (1 << (MySQLParserTIMESTAMP - 41)))) != 0)) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ElementDateContext).dt = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(502)
		p.Match(MySQLParserSTRING)
	}



	return localctx
}


// IElementSubQueryContext is an interface to support dynamic dispatch.
type IElementSubQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSqWith returns the sqWith token.
	GetSqWith() antlr.Token 


	// SetSqWith sets the sqWith token.
	SetSqWith(antlr.Token) 


	// IsElementSubQueryContext differentiates from other interfaces.
	IsElementSubQueryContext()
}

type ElementSubQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	sqWith antlr.Token
}

func NewEmptyElementSubQueryContext() *ElementSubQueryContext {
	var p = new(ElementSubQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementSubQuery
	return p
}

func (*ElementSubQueryContext) IsElementSubQueryContext() {}

func NewElementSubQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementSubQueryContext {
	var p = new(ElementSubQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementSubQuery

	return p
}

func (s *ElementSubQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementSubQueryContext) GetSqWith() antlr.Token { return s.sqWith }


func (s *ElementSubQueryContext) SetSqWith(v antlr.Token) { s.sqWith = v }


func (s *ElementSubQueryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ElementSubQueryContext) SelectStat() ISelectStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatContext)
}

func (s *ElementSubQueryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ElementSubQueryContext) ANY() antlr.TerminalNode {
	return s.GetToken(MySQLParserANY, 0)
}

func (s *ElementSubQueryContext) SOME() antlr.TerminalNode {
	return s.GetToken(MySQLParserSOME, 0)
}

func (s *ElementSubQueryContext) ALL() antlr.TerminalNode {
	return s.GetToken(MySQLParserALL, 0)
}

func (s *ElementSubQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementSubQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementSubQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementSubQuery(s)
	}
}

func (s *ElementSubQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementSubQuery(s)
	}
}




func (p *MySQLParser) ElementSubQuery() (localctx IElementSubQueryContext) {
	localctx = NewElementSubQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, MySQLParserRULE_elementSubQuery)
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
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 44)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 44))) & ((1 << (MySQLParserALL - 44)) | (1 << (MySQLParserANY - 44)) | (1 << (MySQLParserSOME - 44)))) != 0) {
		{
			p.SetState(504)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ElementSubQueryContext).sqWith = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((((_la - 44)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 44))) & ((1 << (MySQLParserALL - 44)) | (1 << (MySQLParserANY - 44)) | (1 << (MySQLParserSOME - 44)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ElementSubQueryContext).sqWith = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(507)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(508)
		p.SelectStat()
	}
	{
		p.SetState(509)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IElementWapperBktContext is an interface to support dynamic dispatch.
type IElementWapperBktContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementWapperBktContext differentiates from other interfaces.
	IsElementWapperBktContext()
}

type ElementWapperBktContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementWapperBktContext() *ElementWapperBktContext {
	var p = new(ElementWapperBktContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementWapperBkt
	return p
}

func (*ElementWapperBktContext) IsElementWapperBktContext() {}

func NewElementWapperBktContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementWapperBktContext {
	var p = new(ElementWapperBktContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementWapperBkt

	return p
}

func (s *ElementWapperBktContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementWapperBktContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ElementWapperBktContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ElementWapperBktContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ElementWapperBktContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementWapperBktContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementWapperBktContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementWapperBkt(s)
	}
}

func (s *ElementWapperBktContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementWapperBkt(s)
	}
}




func (p *MySQLParser) ElementWapperBkt() (localctx IElementWapperBktContext) {
	localctx = NewElementWapperBktContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, MySQLParserRULE_elementWapperBkt)

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
		p.SetState(511)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(512)
		p.Element()
	}
	{
		p.SetState(513)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IElementListFactorContext is an interface to support dynamic dispatch.
type IElementListFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListFactorContext differentiates from other interfaces.
	IsElementListFactorContext()
}

type ElementListFactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListFactorContext() *ElementListFactorContext {
	var p = new(ElementListFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementListFactor
	return p
}

func (*ElementListFactorContext) IsElementListFactorContext() {}

func NewElementListFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListFactorContext {
	var p = new(ElementListFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementListFactor

	return p
}

func (s *ElementListFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListFactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ElementListFactorContext) ElementList() IElementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ElementListFactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ElementListFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementListFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementListFactor(s)
	}
}

func (s *ElementListFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementListFactor(s)
	}
}




func (p *MySQLParser) ElementListFactor() (localctx IElementListFactorContext) {
	localctx = NewElementListFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, MySQLParserRULE_elementListFactor)

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
		p.SetState(515)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(516)
		p.ElementList()
	}
	{
		p.SetState(517)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementList
	return p
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ElementListContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ElementListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLParserCOMMA)
}

func (s *ElementListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, i)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementList(s)
	}
}




func (p *MySQLParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, MySQLParserRULE_elementList)
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
		p.SetState(519)
		p.Element()
	}
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySQLParserCOMMA {
		{
			p.SetState(520)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(521)
			p.Element()
		}


		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IElementOpEleContext is an interface to support dynamic dispatch.
type IElementOpEleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementOpEleContext differentiates from other interfaces.
	IsElementOpEleContext()
}

type ElementOpEleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementOpEleContext() *ElementOpEleContext {
	var p = new(ElementOpEleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementOpEle
	return p
}

func (*ElementOpEleContext) IsElementOpEleContext() {}

func NewElementOpEleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementOpEleContext {
	var p = new(ElementOpEleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementOpEle

	return p
}

func (s *ElementOpEleContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementOpEleContext) ElementOpFactory() IElementOpFactoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpFactoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpFactoryContext)
}

func (s *ElementOpEleContext) ElementOpEleSuffix() IElementOpEleSuffixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpEleSuffixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpEleSuffixContext)
}

func (s *ElementOpEleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementOpEleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementOpEleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementOpEle(s)
	}
}

func (s *ElementOpEleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementOpEle(s)
	}
}




func (p *MySQLParser) ElementOpEle() (localctx IElementOpEleContext) {
	localctx = NewElementOpEleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, MySQLParserRULE_elementOpEle)

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
		p.SetState(527)
		p.ElementOpFactory()
	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(528)
			p.ElementOpEleSuffix()
		}


	}



	return localctx
}


// IElementOpEleSuffixContext is an interface to support dynamic dispatch.
type IElementOpEleSuffixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token 


	// SetOp sets the op token.
	SetOp(antlr.Token) 


	// IsElementOpEleSuffixContext differentiates from other interfaces.
	IsElementOpEleSuffixContext()
}

type ElementOpEleSuffixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op antlr.Token
}

func NewEmptyElementOpEleSuffixContext() *ElementOpEleSuffixContext {
	var p = new(ElementOpEleSuffixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementOpEleSuffix
	return p
}

func (*ElementOpEleSuffixContext) IsElementOpEleSuffixContext() {}

func NewElementOpEleSuffixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementOpEleSuffixContext {
	var p = new(ElementOpEleSuffixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementOpEleSuffix

	return p
}

func (s *ElementOpEleSuffixContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementOpEleSuffixContext) GetOp() antlr.Token { return s.op }


func (s *ElementOpEleSuffixContext) SetOp(v antlr.Token) { s.op = v }


func (s *ElementOpEleSuffixContext) ElementOpEle() IElementOpEleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpEleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpEleContext)
}

func (s *ElementOpEleSuffixContext) VERTBAR() antlr.TerminalNode {
	return s.GetToken(MySQLParserVERTBAR, 0)
}

func (s *ElementOpEleSuffixContext) BITAND() antlr.TerminalNode {
	return s.GetToken(MySQLParserBITAND, 0)
}

func (s *ElementOpEleSuffixContext) SHIFT_LEFT() antlr.TerminalNode {
	return s.GetToken(MySQLParserSHIFT_LEFT, 0)
}

func (s *ElementOpEleSuffixContext) SHIFT_RIGHT() antlr.TerminalNode {
	return s.GetToken(MySQLParserSHIFT_RIGHT, 0)
}

func (s *ElementOpEleSuffixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MySQLParserPLUS, 0)
}

func (s *ElementOpEleSuffixContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MySQLParserMINUS, 0)
}

func (s *ElementOpEleSuffixContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(MySQLParserASTERISK, 0)
}

func (s *ElementOpEleSuffixContext) DIV() antlr.TerminalNode {
	return s.GetToken(MySQLParserDIV, 0)
}

func (s *ElementOpEleSuffixContext) MOD() antlr.TerminalNode {
	return s.GetToken(MySQLParserMOD, 0)
}

func (s *ElementOpEleSuffixContext) POWER_OP() antlr.TerminalNode {
	return s.GetToken(MySQLParserPOWER_OP, 0)
}

func (s *ElementOpEleSuffixContext) AS() antlr.TerminalNode {
	return s.GetToken(MySQLParserAS, 0)
}

func (s *ElementOpEleSuffixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementOpEleSuffixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementOpEleSuffixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementOpEleSuffix(s)
	}
}

func (s *ElementOpEleSuffixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementOpEleSuffix(s)
	}
}




func (p *MySQLParser) ElementOpEleSuffix() (localctx IElementOpEleSuffixContext) {
	localctx = NewElementOpEleSuffixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, MySQLParserRULE_elementOpEleSuffix)
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
	p.SetState(532)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(531)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ElementOpEleSuffixContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLParserAS || ((((_la - 64)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 64))) & ((1 << (MySQLParserDIV - 64)) | (1 << (MySQLParserMOD - 64)) | (1 << (MySQLParserPLUS - 64)) | (1 << (MySQLParserMINUS - 64)) | (1 << (MySQLParserVERTBAR - 64)) | (1 << (MySQLParserBITAND - 64)) | (1 << (MySQLParserSHIFT_LEFT - 64)) | (1 << (MySQLParserSHIFT_RIGHT - 64)) | (1 << (MySQLParserASTERISK - 64)) | (1 << (MySQLParserPOWER_OP - 64)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ElementOpEleSuffixContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	}
	{
		p.SetState(534)
		p.ElementOpEle()
	}



	return localctx
}


// IElementCaseContext is an interface to support dynamic dispatch.
type IElementCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IElementContext

	// GetElseEl returns the elseEl rule contexts.
	GetElseEl() IElementContext


	// SetValue sets the value rule contexts.
	SetValue(IElementContext)

	// SetElseEl sets the elseEl rule contexts.
	SetElseEl(IElementContext)


	// IsElementCaseContext differentiates from other interfaces.
	IsElementCaseContext()
}

type ElementCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value IElementContext 
	elseEl IElementContext 
}

func NewEmptyElementCaseContext() *ElementCaseContext {
	var p = new(ElementCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementCase
	return p
}

func (*ElementCaseContext) IsElementCaseContext() {}

func NewElementCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementCaseContext {
	var p = new(ElementCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementCase

	return p
}

func (s *ElementCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementCaseContext) GetValue() IElementContext { return s.value }

func (s *ElementCaseContext) GetElseEl() IElementContext { return s.elseEl }


func (s *ElementCaseContext) SetValue(v IElementContext) { s.value = v }

func (s *ElementCaseContext) SetElseEl(v IElementContext) { s.elseEl = v }


func (s *ElementCaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(MySQLParserCASE, 0)
}

func (s *ElementCaseContext) CaseWhenPart() ICaseWhenPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseWhenPartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseWhenPartContext)
}

func (s *ElementCaseContext) END() antlr.TerminalNode {
	return s.GetToken(MySQLParserEND, 0)
}

func (s *ElementCaseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MySQLParserELSE, 0)
}

func (s *ElementCaseContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *ElementCaseContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ElementCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementCase(s)
	}
}

func (s *ElementCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementCase(s)
	}
}




func (p *MySQLParser) ElementCase() (localctx IElementCaseContext) {
	localctx = NewElementCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, MySQLParserRULE_elementCase)
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
		p.SetState(536)
		p.Match(MySQLParserCASE)
	}
	p.SetState(538)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MySQLParserPLACEHOLDER) | (1 << MySQLParserNULL) | (1 << MySQLParserTRUE) | (1 << MySQLParserFALSE))) != 0) || ((((_la - 41)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 41))) & ((1 << (MySQLParserDATE - 41)) | (1 << (MySQLParserTIME - 41)) | (1 << (MySQLParserTIMESTAMP - 41)) | (1 << (MySQLParserALL - 41)) | (1 << (MySQLParserANY - 41)) | (1 << (MySQLParserSOME - 41)) | (1 << (MySQLParserUNKNOWN - 41)) | (1 << (MySQLParserCASE - 41)) | (1 << (MySQLParserROW - 41)) | (1 << (MySQLParserBINARY - 41)) | (1 << (MySQLParserASTERISK - 41)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (MySQLParserINT - 74)) | (1 << (MySQLParserDECIMAL - 74)) | (1 << (MySQLParserSTRING - 74)) | (1 << (MySQLParserID - 74)) | (1 << (MySQLParserCOLUMN_REL - 74)) | (1 << (MySQLParserLPAREN - 74)))) != 0) {
		{
			p.SetState(537)

			var _x = p.Element()


			localctx.(*ElementCaseContext).value = _x
		}

	}
	{
		p.SetState(540)
		p.CaseWhenPart()
	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserELSE {
		{
			p.SetState(541)
			p.Match(MySQLParserELSE)
		}
		{
			p.SetState(542)

			var _x = p.Element()


			localctx.(*ElementCaseContext).elseEl = _x
		}

	}
	{
		p.SetState(545)
		p.Match(MySQLParserEND)
	}



	return localctx
}


// ICaseWhenPartContext is an interface to support dynamic dispatch.
type ICaseWhenPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhenEl returns the whenEl rule contexts.
	GetWhenEl() IElementContext

	// GetWhenRe returns the whenRe rule contexts.
	GetWhenRe() IExprRelationalContext

	// GetThen returns the then rule contexts.
	GetThen() IElementContext

	// GetSuffix returns the suffix rule contexts.
	GetSuffix() ICaseWhenPartContext


	// SetWhenEl sets the whenEl rule contexts.
	SetWhenEl(IElementContext)

	// SetWhenRe sets the whenRe rule contexts.
	SetWhenRe(IExprRelationalContext)

	// SetThen sets the then rule contexts.
	SetThen(IElementContext)

	// SetSuffix sets the suffix rule contexts.
	SetSuffix(ICaseWhenPartContext)


	// IsCaseWhenPartContext differentiates from other interfaces.
	IsCaseWhenPartContext()
}

type CaseWhenPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	whenEl IElementContext 
	whenRe IExprRelationalContext 
	then IElementContext 
	suffix ICaseWhenPartContext 
}

func NewEmptyCaseWhenPartContext() *CaseWhenPartContext {
	var p = new(CaseWhenPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_caseWhenPart
	return p
}

func (*CaseWhenPartContext) IsCaseWhenPartContext() {}

func NewCaseWhenPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseWhenPartContext {
	var p = new(CaseWhenPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_caseWhenPart

	return p
}

func (s *CaseWhenPartContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseWhenPartContext) GetWhenEl() IElementContext { return s.whenEl }

func (s *CaseWhenPartContext) GetWhenRe() IExprRelationalContext { return s.whenRe }

func (s *CaseWhenPartContext) GetThen() IElementContext { return s.then }

func (s *CaseWhenPartContext) GetSuffix() ICaseWhenPartContext { return s.suffix }


func (s *CaseWhenPartContext) SetWhenEl(v IElementContext) { s.whenEl = v }

func (s *CaseWhenPartContext) SetWhenRe(v IExprRelationalContext) { s.whenRe = v }

func (s *CaseWhenPartContext) SetThen(v IElementContext) { s.then = v }

func (s *CaseWhenPartContext) SetSuffix(v ICaseWhenPartContext) { s.suffix = v }


func (s *CaseWhenPartContext) WHEN() antlr.TerminalNode {
	return s.GetToken(MySQLParserWHEN, 0)
}

func (s *CaseWhenPartContext) THEN() antlr.TerminalNode {
	return s.GetToken(MySQLParserTHEN, 0)
}

func (s *CaseWhenPartContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *CaseWhenPartContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *CaseWhenPartContext) ExprRelational() IExprRelationalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprRelationalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprRelationalContext)
}

func (s *CaseWhenPartContext) CaseWhenPart() ICaseWhenPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseWhenPartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseWhenPartContext)
}

func (s *CaseWhenPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseWhenPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CaseWhenPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterCaseWhenPart(s)
	}
}

func (s *CaseWhenPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitCaseWhenPart(s)
	}
}




func (p *MySQLParser) CaseWhenPart() (localctx ICaseWhenPartContext) {
	localctx = NewCaseWhenPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, MySQLParserRULE_caseWhenPart)
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
		p.SetState(547)
		p.Match(MySQLParserWHEN)
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(548)

			var _x = p.Element()


			localctx.(*CaseWhenPartContext).whenEl = _x
		}


	case 2:
		{
			p.SetState(549)

			var _x = p.ExprRelational()


			localctx.(*CaseWhenPartContext).whenRe = _x
		}

	}
	{
		p.SetState(552)
		p.Match(MySQLParserTHEN)
	}
	{
		p.SetState(553)

		var _x = p.Element()


		localctx.(*CaseWhenPartContext).then = _x
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserWHEN {
		{
			p.SetState(554)

			var _x = p.CaseWhenPart()


			localctx.(*CaseWhenPartContext).suffix = _x
		}

	}



	return localctx
}


// IElementWithPrefixContext is an interface to support dynamic dispatch.
type IElementWithPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token 


	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token) 


	// IsElementWithPrefixContext differentiates from other interfaces.
	IsElementWithPrefixContext()
}

type ElementWithPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	prefix antlr.Token
}

func NewEmptyElementWithPrefixContext() *ElementWithPrefixContext {
	var p = new(ElementWithPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementWithPrefix
	return p
}

func (*ElementWithPrefixContext) IsElementWithPrefixContext() {}

func NewElementWithPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementWithPrefixContext {
	var p = new(ElementWithPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementWithPrefix

	return p
}

func (s *ElementWithPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementWithPrefixContext) GetPrefix() antlr.Token { return s.prefix }


func (s *ElementWithPrefixContext) SetPrefix(v antlr.Token) { s.prefix = v }


func (s *ElementWithPrefixContext) ElementOpFactory() IElementOpFactoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementOpFactoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementOpFactoryContext)
}

func (s *ElementWithPrefixContext) BINARY() antlr.TerminalNode {
	return s.GetToken(MySQLParserBINARY, 0)
}

func (s *ElementWithPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementWithPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementWithPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementWithPrefix(s)
	}
}

func (s *ElementWithPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementWithPrefix(s)
	}
}




func (p *MySQLParser) ElementWithPrefix() (localctx IElementWithPrefixContext) {
	localctx = NewElementWithPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, MySQLParserRULE_elementWithPrefix)

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
		p.SetState(557)

		var _m = p.Match(MySQLParserBINARY)

		localctx.(*ElementWithPrefixContext).prefix = _m
	}
	{
		p.SetState(558)
		p.ElementOpFactory()
	}



	return localctx
}


// IElementRowContext is an interface to support dynamic dispatch.
type IElementRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementRowContext differentiates from other interfaces.
	IsElementRowContext()
}

type ElementRowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementRowContext() *ElementRowContext {
	var p = new(ElementRowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_elementRow
	return p
}

func (*ElementRowContext) IsElementRowContext() {}

func NewElementRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementRowContext {
	var p = new(ElementRowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_elementRow

	return p
}

func (s *ElementRowContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementRowContext) ROW() antlr.TerminalNode {
	return s.GetToken(MySQLParserROW, 0)
}

func (s *ElementRowContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *ElementRowContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ElementRowContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *ElementRowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementRowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElementRowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterElementRow(s)
	}
}

func (s *ElementRowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitElementRow(s)
	}
}




func (p *MySQLParser) ElementRow() (localctx IElementRowContext) {
	localctx = NewElementRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, MySQLParserRULE_elementRow)

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
		p.SetState(560)
		p.Match(MySQLParserROW)
	}
	{
		p.SetState(561)
		p.Match(MySQLParserLPAREN)
	}
	{
		p.SetState(562)
		p.ParamList()
	}
	{
		p.SetState(563)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IFunCallContext is an interface to support dynamic dispatch.
type IFunCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFunName returns the funName token.
	GetFunName() antlr.Token 


	// SetFunName sets the funName token.
	SetFunName(antlr.Token) 


	// IsFunCallContext differentiates from other interfaces.
	IsFunCallContext()
}

type FunCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	funName antlr.Token
}

func NewEmptyFunCallContext() *FunCallContext {
	var p = new(FunCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_funCall
	return p
}

func (*FunCallContext) IsFunCallContext() {}

func NewFunCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunCallContext {
	var p = new(FunCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_funCall

	return p
}

func (s *FunCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunCallContext) GetFunName() antlr.Token { return s.funName }


func (s *FunCallContext) SetFunName(v antlr.Token) { s.funName = v }


func (s *FunCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserLPAREN, 0)
}

func (s *FunCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLParserRPAREN, 0)
}

func (s *FunCallContext) ID() antlr.TerminalNode {
	return s.GetToken(MySQLParserID, 0)
}

func (s *FunCallContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterFunCall(s)
	}
}

func (s *FunCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitFunCall(s)
	}
}




func (p *MySQLParser) FunCall() (localctx IFunCallContext) {
	localctx = NewFunCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, MySQLParserRULE_funCall)
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
		p.SetState(565)

		var _m = p.Match(MySQLParserID)

		localctx.(*FunCallContext).funName = _m
	}
	{
		p.SetState(566)
		p.Match(MySQLParserLPAREN)
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MySQLParserPLACEHOLDER) | (1 << MySQLParserNULL) | (1 << MySQLParserTRUE) | (1 << MySQLParserFALSE))) != 0) || ((((_la - 41)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 41))) & ((1 << (MySQLParserDATE - 41)) | (1 << (MySQLParserTIME - 41)) | (1 << (MySQLParserTIMESTAMP - 41)) | (1 << (MySQLParserALL - 41)) | (1 << (MySQLParserANY - 41)) | (1 << (MySQLParserSOME - 41)) | (1 << (MySQLParserUNKNOWN - 41)) | (1 << (MySQLParserCASE - 41)) | (1 << (MySQLParserROW - 41)) | (1 << (MySQLParserBINARY - 41)) | (1 << (MySQLParserASTERISK - 41)))) != 0) || ((((_la - 74)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 74))) & ((1 << (MySQLParserINT - 74)) | (1 << (MySQLParserDECIMAL - 74)) | (1 << (MySQLParserSTRING - 74)) | (1 << (MySQLParserID - 74)) | (1 << (MySQLParserCOLUMN_REL - 74)) | (1 << (MySQLParserLPAREN - 74)))) != 0) {
		{
			p.SetState(567)
			p.ParamList()
		}

	}
	{
		p.SetState(570)
		p.Match(MySQLParserRPAREN)
	}



	return localctx
}


// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySQLParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) Element() IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *ParamListContext) ExprRelational() IExprRelationalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprRelationalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprRelationalContext)
}

func (s *ParamListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySQLParserCOMMA, 0)
}

func (s *ParamListContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLListener); ok {
		listenerT.ExitParamList(s)
	}
}




func (p *MySQLParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, MySQLParserRULE_paramList)
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
	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(572)
			p.Element()
		}


	case 2:
		{
			p.SetState(573)
			p.ExprRelational()
		}

	}
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySQLParserCOMMA {
		{
			p.SetState(576)
			p.Match(MySQLParserCOMMA)
		}
		{
			p.SetState(577)
			p.ParamList()
		}

	}



	return localctx
}


