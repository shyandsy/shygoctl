// Code generated from ApiParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package api // ApiParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ApiParserParser struct {
	*antlr.BaseParser
}

var ApiParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func apiparserParserInit() {
	staticData := &ApiParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'('", "')'", "'{'", "'}'", "'*'", "'['", "']'", "'returns'",
		"'-'", "'/'", "'/:'", "'@doc'", "'@handler'", "'interface{}'", "'@server'",
		"'@roles'", "'@features'", "", "", "", "", "", "", "", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "ATDOC", "ATHANDLER",
		"INTERFACE", "ATSERVER", "ATROLES", "ATFEATURES", "WS", "COMMENT", "LINE_COMMENT",
		"STRING", "RAW_STRING", "LINE_VALUE", "ID", "DOT", "LetterOrDigit",
	}
	staticData.RuleNames = []string{
		"api", "spec", "syntaxLit", "importSpec", "importLit", "importBlock",
		"importBlockValue", "importValue", "infoSpec", "typeSpec", "typeLit",
		"typeBlock", "typeLitBody", "typeBlockBody", "typeStruct", "typeAlias",
		"typeBlockStruct", "typeBlockAlias", "field", "normalField", "anonymousFiled",
		"dataType", "qualifiedType", "pointerType", "mapType", "arrayType",
		"serviceSpec", "atServer", "serviceApi", "serviceRoute", "atDoc", "atHandler",
		"atRoles", "atFeatures", "route", "body", "replybody", "kvLit", "serviceName",
		"path", "pathItem",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 400, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 5, 0, 84,
		8, 0, 10, 0, 12, 0, 87, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 94, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 104, 8, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 114, 8, 5, 11, 5, 12,
		5, 115, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 4, 8, 129, 8, 8, 11, 8, 12, 8, 130, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 137,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 147,
		8, 11, 10, 11, 12, 11, 150, 9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 156,
		8, 12, 1, 13, 1, 13, 3, 13, 160, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14, 165,
		8, 14, 1, 14, 1, 14, 5, 14, 169, 8, 14, 10, 14, 12, 14, 172, 9, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 179, 8, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 3, 16, 186, 8, 16, 1, 16, 1, 16, 5, 16, 190, 8, 16, 10, 16,
		12, 16, 193, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 200, 8, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 3, 18, 207, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 213, 8, 19, 1, 20, 3, 20, 216, 8, 20, 1, 20, 1, 20, 3,
		20, 220, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21,
		229, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 26, 3, 26, 251, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 4, 27, 258,
		8, 27, 11, 27, 12, 27, 259, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 5, 28, 269, 8, 28, 10, 28, 12, 28, 272, 9, 28, 1, 28, 1, 28, 1, 29,
		3, 29, 277, 8, 29, 1, 29, 3, 29, 280, 8, 29, 1, 29, 3, 29, 283, 8, 29,
		1, 29, 1, 29, 3, 29, 287, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 293,
		8, 30, 1, 30, 4, 30, 296, 8, 30, 11, 30, 12, 30, 297, 1, 30, 3, 30, 301,
		8, 30, 1, 30, 3, 30, 304, 8, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 3,
		32, 311, 8, 32, 1, 32, 4, 32, 314, 8, 32, 11, 32, 12, 32, 315, 1, 32, 3,
		32, 319, 8, 32, 1, 32, 3, 32, 322, 8, 32, 1, 33, 1, 33, 3, 33, 326, 8,
		33, 1, 33, 4, 33, 329, 8, 33, 11, 33, 12, 33, 330, 1, 33, 3, 33, 334, 8,
		33, 1, 33, 3, 33, 337, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 343, 8,
		34, 1, 34, 3, 34, 346, 8, 34, 1, 35, 1, 35, 3, 35, 350, 8, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 3, 36, 357, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 38, 1, 38, 3, 38, 367, 8, 38, 4, 38, 369, 8, 38, 11, 38,
		12, 38, 370, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 377, 8, 39, 10, 39, 12,
		39, 380, 9, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3, 39, 386, 8, 39, 4, 39, 388,
		8, 39, 11, 39, 12, 39, 389, 1, 39, 3, 39, 393, 8, 39, 1, 40, 4, 40, 396,
		8, 40, 11, 40, 12, 40, 397, 1, 40, 0, 0, 41, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 0, 1, 2, 0,
		25, 25, 27, 27, 417, 0, 85, 1, 0, 0, 0, 2, 93, 1, 0, 0, 0, 4, 95, 1, 0,
		0, 0, 6, 103, 1, 0, 0, 0, 8, 105, 1, 0, 0, 0, 10, 109, 1, 0, 0, 0, 12,
		119, 1, 0, 0, 0, 14, 121, 1, 0, 0, 0, 16, 124, 1, 0, 0, 0, 18, 136, 1,
		0, 0, 0, 20, 138, 1, 0, 0, 0, 22, 142, 1, 0, 0, 0, 24, 155, 1, 0, 0, 0,
		26, 159, 1, 0, 0, 0, 28, 161, 1, 0, 0, 0, 30, 175, 1, 0, 0, 0, 32, 182,
		1, 0, 0, 0, 34, 196, 1, 0, 0, 0, 36, 206, 1, 0, 0, 0, 38, 208, 1, 0, 0,
		0, 40, 219, 1, 0, 0, 0, 42, 228, 1, 0, 0, 0, 44, 230, 1, 0, 0, 0, 46, 234,
		1, 0, 0, 0, 48, 237, 1, 0, 0, 0, 50, 245, 1, 0, 0, 0, 52, 250, 1, 0, 0,
		0, 54, 254, 1, 0, 0, 0, 56, 263, 1, 0, 0, 0, 58, 276, 1, 0, 0, 0, 60, 290,
		1, 0, 0, 0, 62, 305, 1, 0, 0, 0, 64, 308, 1, 0, 0, 0, 66, 323, 1, 0, 0,
		0, 68, 338, 1, 0, 0, 0, 70, 347, 1, 0, 0, 0, 72, 353, 1, 0, 0, 0, 74, 360,
		1, 0, 0, 0, 76, 368, 1, 0, 0, 0, 78, 392, 1, 0, 0, 0, 80, 395, 1, 0, 0,
		0, 82, 84, 3, 2, 1, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83,
		1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 1, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0,
		88, 94, 3, 4, 2, 0, 89, 94, 3, 6, 3, 0, 90, 94, 3, 16, 8, 0, 91, 94, 3,
		18, 9, 0, 92, 94, 3, 52, 26, 0, 93, 88, 1, 0, 0, 0, 93, 89, 1, 0, 0, 0,
		93, 90, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 3, 1, 0,
		0, 0, 95, 96, 6, 2, -1, 0, 96, 97, 5, 25, 0, 0, 97, 98, 5, 1, 0, 0, 98,
		99, 6, 2, -1, 0, 99, 100, 5, 22, 0, 0, 100, 5, 1, 0, 0, 0, 101, 104, 3,
		8, 4, 0, 102, 104, 3, 10, 5, 0, 103, 101, 1, 0, 0, 0, 103, 102, 1, 0, 0,
		0, 104, 7, 1, 0, 0, 0, 105, 106, 6, 4, -1, 0, 106, 107, 5, 25, 0, 0, 107,
		108, 3, 14, 7, 0, 108, 9, 1, 0, 0, 0, 109, 110, 6, 5, -1, 0, 110, 111,
		5, 25, 0, 0, 111, 113, 5, 2, 0, 0, 112, 114, 3, 12, 6, 0, 113, 112, 1,
		0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 3, 0, 0, 118, 11, 1, 0, 0, 0, 119,
		120, 3, 14, 7, 0, 120, 13, 1, 0, 0, 0, 121, 122, 6, 7, -1, 0, 122, 123,
		5, 22, 0, 0, 123, 15, 1, 0, 0, 0, 124, 125, 6, 8, -1, 0, 125, 126, 5, 25,
		0, 0, 126, 128, 5, 2, 0, 0, 127, 129, 3, 74, 37, 0, 128, 127, 1, 0, 0,
		0, 129, 130, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 133, 5, 3, 0, 0, 133, 17, 1, 0, 0, 0, 134, 137, 3,
		20, 10, 0, 135, 137, 3, 22, 11, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0,
		0, 0, 137, 19, 1, 0, 0, 0, 138, 139, 6, 10, -1, 0, 139, 140, 5, 25, 0,
		0, 140, 141, 3, 24, 12, 0, 141, 21, 1, 0, 0, 0, 142, 143, 6, 11, -1, 0,
		143, 144, 5, 25, 0, 0, 144, 148, 5, 2, 0, 0, 145, 147, 3, 26, 13, 0, 146,
		145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 5, 3,
		0, 0, 152, 23, 1, 0, 0, 0, 153, 156, 3, 28, 14, 0, 154, 156, 3, 30, 15,
		0, 155, 153, 1, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 25, 1, 0, 0, 0, 157,
		160, 3, 32, 16, 0, 158, 160, 3, 34, 17, 0, 159, 157, 1, 0, 0, 0, 159, 158,
		1, 0, 0, 0, 160, 27, 1, 0, 0, 0, 161, 162, 6, 14, -1, 0, 162, 164, 5, 25,
		0, 0, 163, 165, 5, 25, 0, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 170, 5, 4, 0, 0, 167, 169, 3, 36, 18, 0, 168,
		167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 5,
		0, 0, 174, 29, 1, 0, 0, 0, 175, 176, 6, 15, -1, 0, 176, 178, 5, 25, 0,
		0, 177, 179, 5, 1, 0, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		180, 1, 0, 0, 0, 180, 181, 3, 42, 21, 0, 181, 31, 1, 0, 0, 0, 182, 183,
		6, 16, -1, 0, 183, 185, 5, 25, 0, 0, 184, 186, 5, 25, 0, 0, 185, 184, 1,
		0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 191, 5, 4, 0,
		0, 188, 190, 3, 36, 18, 0, 189, 188, 1, 0, 0, 0, 190, 193, 1, 0, 0, 0,
		191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0, 193,
		191, 1, 0, 0, 0, 194, 195, 5, 5, 0, 0, 195, 33, 1, 0, 0, 0, 196, 197, 6,
		17, -1, 0, 197, 199, 5, 25, 0, 0, 198, 200, 5, 1, 0, 0, 199, 198, 1, 0,
		0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 3, 42, 21,
		0, 202, 35, 1, 0, 0, 0, 203, 204, 4, 18, 0, 0, 204, 207, 3, 38, 19, 0,
		205, 207, 3, 40, 20, 0, 206, 203, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207,
		37, 1, 0, 0, 0, 208, 209, 6, 19, -1, 0, 209, 210, 5, 25, 0, 0, 210, 212,
		3, 42, 21, 0, 211, 213, 5, 23, 0, 0, 212, 211, 1, 0, 0, 0, 212, 213, 1,
		0, 0, 0, 213, 39, 1, 0, 0, 0, 214, 216, 5, 6, 0, 0, 215, 214, 1, 0, 0,
		0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 220, 3, 44, 22, 0,
		218, 220, 5, 25, 0, 0, 219, 215, 1, 0, 0, 0, 219, 218, 1, 0, 0, 0, 220,
		41, 1, 0, 0, 0, 221, 229, 3, 44, 22, 0, 222, 229, 5, 25, 0, 0, 223, 229,
		3, 48, 24, 0, 224, 229, 3, 50, 25, 0, 225, 229, 5, 15, 0, 0, 226, 229,
		3, 46, 23, 0, 227, 229, 3, 28, 14, 0, 228, 221, 1, 0, 0, 0, 228, 222, 1,
		0, 0, 0, 228, 223, 1, 0, 0, 0, 228, 224, 1, 0, 0, 0, 228, 225, 1, 0, 0,
		0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 43, 1, 0, 0, 0, 230,
		231, 5, 25, 0, 0, 231, 232, 5, 26, 0, 0, 232, 233, 5, 25, 0, 0, 233, 45,
		1, 0, 0, 0, 234, 235, 5, 6, 0, 0, 235, 236, 3, 42, 21, 0, 236, 47, 1, 0,
		0, 0, 237, 238, 6, 24, -1, 0, 238, 239, 5, 25, 0, 0, 239, 240, 5, 7, 0,
		0, 240, 241, 6, 24, -1, 0, 241, 242, 5, 25, 0, 0, 242, 243, 5, 8, 0, 0,
		243, 244, 3, 42, 21, 0, 244, 49, 1, 0, 0, 0, 245, 246, 5, 7, 0, 0, 246,
		247, 5, 8, 0, 0, 247, 248, 3, 42, 21, 0, 248, 51, 1, 0, 0, 0, 249, 251,
		3, 54, 27, 0, 250, 249, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 1,
		0, 0, 0, 252, 253, 3, 56, 28, 0, 253, 53, 1, 0, 0, 0, 254, 255, 5, 16,
		0, 0, 255, 257, 5, 2, 0, 0, 256, 258, 3, 74, 37, 0, 257, 256, 1, 0, 0,
		0, 258, 259, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260,
		261, 1, 0, 0, 0, 261, 262, 5, 3, 0, 0, 262, 55, 1, 0, 0, 0, 263, 264, 6,
		28, -1, 0, 264, 265, 5, 25, 0, 0, 265, 266, 3, 76, 38, 0, 266, 270, 5,
		4, 0, 0, 267, 269, 3, 58, 29, 0, 268, 267, 1, 0, 0, 0, 269, 272, 1, 0,
		0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0,
		272, 270, 1, 0, 0, 0, 273, 274, 5, 5, 0, 0, 274, 57, 1, 0, 0, 0, 275, 277,
		3, 60, 30, 0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 279, 1,
		0, 0, 0, 278, 280, 3, 64, 32, 0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0,
		0, 0, 280, 282, 1, 0, 0, 0, 281, 283, 3, 66, 33, 0, 282, 281, 1, 0, 0,
		0, 282, 283, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 287, 3, 54, 27, 0,
		285, 287, 3, 62, 31, 0, 286, 284, 1, 0, 0, 0, 286, 285, 1, 0, 0, 0, 287,
		288, 1, 0, 0, 0, 288, 289, 3, 68, 34, 0, 289, 59, 1, 0, 0, 0, 290, 292,
		5, 13, 0, 0, 291, 293, 5, 2, 0, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0,
		0, 0, 293, 300, 1, 0, 0, 0, 294, 296, 3, 74, 37, 0, 295, 294, 1, 0, 0,
		0, 296, 297, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298,
		301, 1, 0, 0, 0, 299, 301, 5, 22, 0, 0, 300, 295, 1, 0, 0, 0, 300, 299,
		1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 304, 5, 3, 0, 0, 303, 302, 1, 0,
		0, 0, 303, 304, 1, 0, 0, 0, 304, 61, 1, 0, 0, 0, 305, 306, 5, 14, 0, 0,
		306, 307, 5, 25, 0, 0, 307, 63, 1, 0, 0, 0, 308, 310, 5, 17, 0, 0, 309,
		311, 5, 7, 0, 0, 310, 309, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 318,
		1, 0, 0, 0, 312, 314, 3, 74, 37, 0, 313, 312, 1, 0, 0, 0, 314, 315, 1,
		0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 319, 1, 0, 0,
		0, 317, 319, 5, 22, 0, 0, 318, 313, 1, 0, 0, 0, 318, 317, 1, 0, 0, 0, 319,
		321, 1, 0, 0, 0, 320, 322, 5, 8, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322,
		1, 0, 0, 0, 322, 65, 1, 0, 0, 0, 323, 325, 5, 18, 0, 0, 324, 326, 5, 7,
		0, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 333, 1, 0, 0, 0,
		327, 329, 3, 74, 37, 0, 328, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330,
		328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 334,
		5, 22, 0, 0, 333, 328, 1, 0, 0, 0, 333, 332, 1, 0, 0, 0, 334, 336, 1, 0,
		0, 0, 335, 337, 5, 8, 0, 0, 336, 335, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0,
		337, 67, 1, 0, 0, 0, 338, 339, 6, 34, -1, 0, 339, 340, 5, 25, 0, 0, 340,
		342, 3, 78, 39, 0, 341, 343, 3, 70, 35, 0, 342, 341, 1, 0, 0, 0, 342, 343,
		1, 0, 0, 0, 343, 345, 1, 0, 0, 0, 344, 346, 3, 72, 36, 0, 345, 344, 1,
		0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 69, 1, 0, 0, 0, 347, 349, 5, 2, 0,
		0, 348, 350, 5, 25, 0, 0, 349, 348, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350,
		351, 1, 0, 0, 0, 351, 352, 5, 3, 0, 0, 352, 71, 1, 0, 0, 0, 353, 354, 5,
		9, 0, 0, 354, 356, 5, 2, 0, 0, 355, 357, 3, 42, 21, 0, 356, 355, 1, 0,
		0, 0, 356, 357, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 5, 3, 0, 0,
		359, 73, 1, 0, 0, 0, 360, 361, 5, 25, 0, 0, 361, 362, 6, 37, -1, 0, 362,
		363, 5, 24, 0, 0, 363, 75, 1, 0, 0, 0, 364, 366, 5, 25, 0, 0, 365, 367,
		5, 10, 0, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 369, 1, 0,
		0, 0, 368, 364, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0,
		370, 371, 1, 0, 0, 0, 371, 77, 1, 0, 0, 0, 372, 373, 5, 11, 0, 0, 373,
		378, 3, 80, 40, 0, 374, 375, 5, 10, 0, 0, 375, 377, 3, 80, 40, 0, 376,
		374, 1, 0, 0, 0, 377, 380, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379,
		1, 0, 0, 0, 379, 388, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 381, 382, 5, 12,
		0, 0, 382, 385, 3, 80, 40, 0, 383, 384, 5, 10, 0, 0, 384, 386, 3, 80, 40,
		0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 388, 1, 0, 0, 0, 387,
		372, 1, 0, 0, 0, 387, 381, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 387,
		1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391, 393, 5, 11,
		0, 0, 392, 387, 1, 0, 0, 0, 392, 391, 1, 0, 0, 0, 393, 79, 1, 0, 0, 0,
		394, 396, 7, 0, 0, 0, 395, 394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397,
		395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 81, 1, 0, 0, 0, 51, 85, 93,
		103, 115, 130, 136, 148, 155, 159, 164, 170, 178, 185, 191, 199, 206, 212,
		215, 219, 228, 250, 259, 270, 276, 279, 282, 286, 292, 297, 300, 303, 310,
		315, 318, 321, 325, 330, 333, 336, 342, 345, 349, 356, 366, 370, 378, 385,
		387, 389, 392, 397,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ApiParserParserInit initializes any static state used to implement ApiParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewApiParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ApiParserParserInit() {
	staticData := &ApiParserParserStaticData
	staticData.once.Do(apiparserParserInit)
}

// NewApiParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewApiParserParser(input antlr.TokenStream) *ApiParserParser {
	ApiParserParserInit()
	this := new(ApiParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ApiParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ApiParser.g4"

	return this
}

// ApiParserParser tokens.
const (
	ApiParserParserEOF           = antlr.TokenEOF
	ApiParserParserT__0          = 1
	ApiParserParserT__1          = 2
	ApiParserParserT__2          = 3
	ApiParserParserT__3          = 4
	ApiParserParserT__4          = 5
	ApiParserParserT__5          = 6
	ApiParserParserT__6          = 7
	ApiParserParserT__7          = 8
	ApiParserParserT__8          = 9
	ApiParserParserT__9          = 10
	ApiParserParserT__10         = 11
	ApiParserParserT__11         = 12
	ApiParserParserATDOC         = 13
	ApiParserParserATHANDLER     = 14
	ApiParserParserINTERFACE     = 15
	ApiParserParserATSERVER      = 16
	ApiParserParserATROLES       = 17
	ApiParserParserATFEATURES    = 18
	ApiParserParserWS            = 19
	ApiParserParserCOMMENT       = 20
	ApiParserParserLINE_COMMENT  = 21
	ApiParserParserSTRING        = 22
	ApiParserParserRAW_STRING    = 23
	ApiParserParserLINE_VALUE    = 24
	ApiParserParserID            = 25
	ApiParserParserDOT           = 26
	ApiParserParserLetterOrDigit = 27
)

// ApiParserParser rules.
const (
	ApiParserParserRULE_api              = 0
	ApiParserParserRULE_spec             = 1
	ApiParserParserRULE_syntaxLit        = 2
	ApiParserParserRULE_importSpec       = 3
	ApiParserParserRULE_importLit        = 4
	ApiParserParserRULE_importBlock      = 5
	ApiParserParserRULE_importBlockValue = 6
	ApiParserParserRULE_importValue      = 7
	ApiParserParserRULE_infoSpec         = 8
	ApiParserParserRULE_typeSpec         = 9
	ApiParserParserRULE_typeLit          = 10
	ApiParserParserRULE_typeBlock        = 11
	ApiParserParserRULE_typeLitBody      = 12
	ApiParserParserRULE_typeBlockBody    = 13
	ApiParserParserRULE_typeStruct       = 14
	ApiParserParserRULE_typeAlias        = 15
	ApiParserParserRULE_typeBlockStruct  = 16
	ApiParserParserRULE_typeBlockAlias   = 17
	ApiParserParserRULE_field            = 18
	ApiParserParserRULE_normalField      = 19
	ApiParserParserRULE_anonymousFiled   = 20
	ApiParserParserRULE_dataType         = 21
	ApiParserParserRULE_qualifiedType    = 22
	ApiParserParserRULE_pointerType      = 23
	ApiParserParserRULE_mapType          = 24
	ApiParserParserRULE_arrayType        = 25
	ApiParserParserRULE_serviceSpec      = 26
	ApiParserParserRULE_atServer         = 27
	ApiParserParserRULE_serviceApi       = 28
	ApiParserParserRULE_serviceRoute     = 29
	ApiParserParserRULE_atDoc            = 30
	ApiParserParserRULE_atHandler        = 31
	ApiParserParserRULE_atRoles          = 32
	ApiParserParserRULE_atFeatures       = 33
	ApiParserParserRULE_route            = 34
	ApiParserParserRULE_body             = 35
	ApiParserParserRULE_replybody        = 36
	ApiParserParserRULE_kvLit            = 37
	ApiParserParserRULE_serviceName      = 38
	ApiParserParserRULE_path             = 39
	ApiParserParserRULE_pathItem         = 40
)

// IApiContext is an interface to support dynamic dispatch.
type IApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSpec() []ISpecContext
	Spec(i int) ISpecContext

	// IsApiContext differentiates from other interfaces.
	IsApiContext()
}

type ApiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiContext() *ApiContext {
	var p = new(ApiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_api
	return p
}

func InitEmptyApiContext(p *ApiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_api
}

func (*ApiContext) IsApiContext() {}

func NewApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiContext {
	var p = new(ApiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_api

	return p
}

func (s *ApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiContext) AllSpec() []ISpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISpecContext); ok {
			len++
		}
	}

	tst := make([]ISpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISpecContext); ok {
			tst[i] = t.(ISpecContext)
			i++
		}
	}

	return tst
}

func (s *ApiContext) Spec(i int) ISpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *ApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterApi(s)
	}
}

func (s *ApiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitApi(s)
	}
}

func (s *ApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Api() (localctx IApiContext) {
	localctx = NewApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ApiParserParserRULE_api)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserATSERVER || _la == ApiParserParserID {
		{
			p.SetState(82)
			p.Spec()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SyntaxLit() ISyntaxLitContext
	ImportSpec() IImportSpecContext
	InfoSpec() IInfoSpecContext
	TypeSpec() ITypeSpecContext
	ServiceSpec() IServiceSpecContext

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
	return p
}

func InitEmptySpecContext(p *SpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_spec
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) SyntaxLit() ISyntaxLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxLitContext)
}

func (s *SpecContext) ImportSpec() IImportSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *SpecContext) InfoSpec() IInfoSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfoSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfoSpecContext)
}

func (s *SpecContext) TypeSpec() ITypeSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpecContext) ServiceSpec() IServiceSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceSpecContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ApiParserParserRULE_spec)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.SyntaxLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.ImportSpec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.InfoSpec()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)
			p.TypeSpec()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(92)
			p.ServiceSpec()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISyntaxLitContext is an interface to support dynamic dispatch.
type ISyntaxLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSyntaxToken returns the syntaxToken token.
	GetSyntaxToken() antlr.Token

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// GetVersion returns the version token.
	GetVersion() antlr.Token

	// SetSyntaxToken sets the syntaxToken token.
	SetSyntaxToken(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// SetVersion sets the version token.
	SetVersion(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsSyntaxLitContext differentiates from other interfaces.
	IsSyntaxLitContext()
}

type SyntaxLitContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	syntaxToken antlr.Token
	assign      antlr.Token
	version     antlr.Token
}

func NewEmptySyntaxLitContext() *SyntaxLitContext {
	var p = new(SyntaxLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_syntaxLit
	return p
}

func InitEmptySyntaxLitContext(p *SyntaxLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_syntaxLit
}

func (*SyntaxLitContext) IsSyntaxLitContext() {}

func NewSyntaxLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxLitContext {
	var p = new(SyntaxLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_syntaxLit

	return p
}

func (s *SyntaxLitContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxLitContext) GetSyntaxToken() antlr.Token { return s.syntaxToken }

func (s *SyntaxLitContext) GetAssign() antlr.Token { return s.assign }

func (s *SyntaxLitContext) GetVersion() antlr.Token { return s.version }

func (s *SyntaxLitContext) SetSyntaxToken(v antlr.Token) { s.syntaxToken = v }

func (s *SyntaxLitContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *SyntaxLitContext) SetVersion(v antlr.Token) { s.version = v }

func (s *SyntaxLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *SyntaxLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *SyntaxLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterSyntaxLit(s)
	}
}

func (s *SyntaxLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitSyntaxLit(s)
	}
}

func (s *SyntaxLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitSyntaxLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) SyntaxLit() (localctx ISyntaxLitContext) {
	localctx = NewSyntaxLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ApiParserParserRULE_syntaxLit)
	p.EnterOuterAlt(localctx, 1)
	match(p, "syntax")
	{
		p.SetState(96)

		var _m = p.Match(ApiParserParserID)

		localctx.(*SyntaxLitContext).syntaxToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)

		var _m = p.Match(ApiParserParserT__0)

		localctx.(*SyntaxLitContext).assign = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	checkVersion(p)
	{
		p.SetState(99)

		var _m = p.Match(ApiParserParserSTRING)

		localctx.(*SyntaxLitContext).version = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportLit() IImportLitContext
	ImportBlock() IImportBlockContext

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importSpec
	return p
}

func InitEmptyImportSpecContext(p *ImportSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importSpec
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) ImportLit() IImportLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportLitContext)
}

func (s *ImportSpecContext) ImportBlock() IImportBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportBlockContext)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportSpec() (localctx IImportSpecContext) {
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ApiParserParserRULE_importSpec)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.ImportLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.ImportBlock()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportLitContext is an interface to support dynamic dispatch.
type IImportLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportToken returns the importToken token.
	GetImportToken() antlr.Token

	// SetImportToken sets the importToken token.
	SetImportToken(antlr.Token)

	// Getter signatures
	ImportValue() IImportValueContext
	ID() antlr.TerminalNode

	// IsImportLitContext differentiates from other interfaces.
	IsImportLitContext()
}

type ImportLitContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	importToken antlr.Token
}

func NewEmptyImportLitContext() *ImportLitContext {
	var p = new(ImportLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importLit
	return p
}

func InitEmptyImportLitContext(p *ImportLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importLit
}

func (*ImportLitContext) IsImportLitContext() {}

func NewImportLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportLitContext {
	var p = new(ImportLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importLit

	return p
}

func (s *ImportLitContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportLitContext) GetImportToken() antlr.Token { return s.importToken }

func (s *ImportLitContext) SetImportToken(v antlr.Token) { s.importToken = v }

func (s *ImportLitContext) ImportValue() IImportValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportValueContext)
}

func (s *ImportLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ImportLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterImportLit(s)
	}
}

func (s *ImportLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitImportLit(s)
	}
}

func (s *ImportLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportLit() (localctx IImportLitContext) {
	localctx = NewImportLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ApiParserParserRULE_importLit)
	p.EnterOuterAlt(localctx, 1)
	match(p, "import")
	{
		p.SetState(106)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportLitContext).importToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.ImportValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportBlockContext is an interface to support dynamic dispatch.
type IImportBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetImportToken returns the importToken token.
	GetImportToken() antlr.Token

	// SetImportToken sets the importToken token.
	SetImportToken(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	AllImportBlockValue() []IImportBlockValueContext
	ImportBlockValue(i int) IImportBlockValueContext

	// IsImportBlockContext differentiates from other interfaces.
	IsImportBlockContext()
}

type ImportBlockContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	importToken antlr.Token
}

func NewEmptyImportBlockContext() *ImportBlockContext {
	var p = new(ImportBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlock
	return p
}

func InitEmptyImportBlockContext(p *ImportBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlock
}

func (*ImportBlockContext) IsImportBlockContext() {}

func NewImportBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockContext {
	var p = new(ImportBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importBlock

	return p
}

func (s *ImportBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportBlockContext) GetImportToken() antlr.Token { return s.importToken }

func (s *ImportBlockContext) SetImportToken(v antlr.Token) { s.importToken = v }

func (s *ImportBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ImportBlockContext) AllImportBlockValue() []IImportBlockValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportBlockValueContext); ok {
			len++
		}
	}

	tst := make([]IImportBlockValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportBlockValueContext); ok {
			tst[i] = t.(IImportBlockValueContext)
			i++
		}
	}

	return tst
}

func (s *ImportBlockContext) ImportBlockValue(i int) IImportBlockValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportBlockValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportBlockValueContext)
}

func (s *ImportBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterImportBlock(s)
	}
}

func (s *ImportBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitImportBlock(s)
	}
}

func (s *ImportBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportBlock() (localctx IImportBlockContext) {
	localctx = NewImportBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ApiParserParserRULE_importBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	match(p, "import")
	{
		p.SetState(110)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ImportBlockContext).importToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(ApiParserParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserSTRING {
		{
			p.SetState(112)
			p.ImportBlockValue()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(ApiParserParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportBlockValueContext is an interface to support dynamic dispatch.
type IImportBlockValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ImportValue() IImportValueContext

	// IsImportBlockValueContext differentiates from other interfaces.
	IsImportBlockValueContext()
}

type ImportBlockValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportBlockValueContext() *ImportBlockValueContext {
	var p = new(ImportBlockValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlockValue
	return p
}

func InitEmptyImportBlockValueContext(p *ImportBlockValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importBlockValue
}

func (*ImportBlockValueContext) IsImportBlockValueContext() {}

func NewImportBlockValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportBlockValueContext {
	var p = new(ImportBlockValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importBlockValue

	return p
}

func (s *ImportBlockValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportBlockValueContext) ImportValue() IImportValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportValueContext)
}

func (s *ImportBlockValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportBlockValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportBlockValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterImportBlockValue(s)
	}
}

func (s *ImportBlockValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitImportBlockValue(s)
	}
}

func (s *ImportBlockValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportBlockValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportBlockValue() (localctx IImportBlockValueContext) {
	localctx = NewImportBlockValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ApiParserParserRULE_importBlockValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.ImportValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportValueContext is an interface to support dynamic dispatch.
type IImportValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsImportValueContext differentiates from other interfaces.
	IsImportValueContext()
}

type ImportValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportValueContext() *ImportValueContext {
	var p = new(ImportValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importValue
	return p
}

func InitEmptyImportValueContext(p *ImportValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_importValue
}

func (*ImportValueContext) IsImportValueContext() {}

func NewImportValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportValueContext {
	var p = new(ImportValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_importValue

	return p
}

func (s *ImportValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *ImportValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterImportValue(s)
	}
}

func (s *ImportValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitImportValue(s)
	}
}

func (s *ImportValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitImportValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ImportValue() (localctx IImportValueContext) {
	localctx = NewImportValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ApiParserParserRULE_importValue)
	p.EnterOuterAlt(localctx, 1)
	checkImportValue(p)
	{
		p.SetState(122)
		p.Match(ApiParserParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInfoSpecContext is an interface to support dynamic dispatch.
type IInfoSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInfoToken returns the infoToken token.
	GetInfoToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetInfoToken sets the infoToken token.
	SetInfoToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsInfoSpecContext differentiates from other interfaces.
	IsInfoSpecContext()
}

type InfoSpecContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	infoToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyInfoSpecContext() *InfoSpecContext {
	var p = new(InfoSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_infoSpec
	return p
}

func InitEmptyInfoSpecContext(p *InfoSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_infoSpec
}

func (*InfoSpecContext) IsInfoSpecContext() {}

func NewInfoSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfoSpecContext {
	var p = new(InfoSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_infoSpec

	return p
}

func (s *InfoSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *InfoSpecContext) GetInfoToken() antlr.Token { return s.infoToken }

func (s *InfoSpecContext) GetLp() antlr.Token { return s.lp }

func (s *InfoSpecContext) GetRp() antlr.Token { return s.rp }

func (s *InfoSpecContext) SetInfoToken(v antlr.Token) { s.infoToken = v }

func (s *InfoSpecContext) SetLp(v antlr.Token) { s.lp = v }

func (s *InfoSpecContext) SetRp(v antlr.Token) { s.rp = v }

func (s *InfoSpecContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *InfoSpecContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *InfoSpecContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *InfoSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfoSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfoSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterInfoSpec(s)
	}
}

func (s *InfoSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitInfoSpec(s)
	}
}

func (s *InfoSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitInfoSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) InfoSpec() (localctx IInfoSpecContext) {
	localctx = NewInfoSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ApiParserParserRULE_infoSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	match(p, "info")
	{
		p.SetState(125)

		var _m = p.Match(ApiParserParserID)

		localctx.(*InfoSpecContext).infoToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*InfoSpecContext).lp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(127)
			p.KvLit()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*InfoSpecContext).rp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeLit() ITypeLitContext
	TypeBlock() ITypeBlockContext

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeSpec
	return p
}

func InitEmptyTypeSpecContext(p *TypeSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeSpec
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) TypeLit() ITypeLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeLitContext)
}

func (s *TypeSpecContext) TypeBlock() ITypeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ApiParserParserRULE_typeSpec)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.TypeLit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.TypeBlock()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeLitContext is an interface to support dynamic dispatch.
type ITypeLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// Getter signatures
	TypeLitBody() ITypeLitBodyContext
	ID() antlr.TerminalNode

	// IsTypeLitContext differentiates from other interfaces.
	IsTypeLitContext()
}

type TypeLitContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	typeToken antlr.Token
}

func NewEmptyTypeLitContext() *TypeLitContext {
	var p = new(TypeLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLit
	return p
}

func InitEmptyTypeLitContext(p *TypeLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLit
}

func (*TypeLitContext) IsTypeLitContext() {}

func NewTypeLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitContext {
	var p = new(TypeLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeLit

	return p
}

func (s *TypeLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitContext) GetTypeToken() antlr.Token { return s.typeToken }

func (s *TypeLitContext) SetTypeToken(v antlr.Token) { s.typeToken = v }

func (s *TypeLitContext) TypeLitBody() ITypeLitBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeLitBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeLitBodyContext)
}

func (s *TypeLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeLit(s)
	}
}

func (s *TypeLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeLit(s)
	}
}

func (s *TypeLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeLit() (localctx ITypeLitContext) {
	localctx = NewTypeLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ApiParserParserRULE_typeLit)
	p.EnterOuterAlt(localctx, 1)
	match(p, "type")
	{
		p.SetState(139)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeLitContext).typeToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.TypeLitBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeBlockContext is an interface to support dynamic dispatch.
type ITypeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeToken returns the typeToken token.
	GetTypeToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetTypeToken sets the typeToken token.
	SetTypeToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	AllTypeBlockBody() []ITypeBlockBodyContext
	TypeBlockBody(i int) ITypeBlockBodyContext

	// IsTypeBlockContext differentiates from other interfaces.
	IsTypeBlockContext()
}

type TypeBlockContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	typeToken antlr.Token
	lp        antlr.Token
	rp        antlr.Token
}

func NewEmptyTypeBlockContext() *TypeBlockContext {
	var p = new(TypeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlock
	return p
}

func InitEmptyTypeBlockContext(p *TypeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlock
}

func (*TypeBlockContext) IsTypeBlockContext() {}

func NewTypeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockContext {
	var p = new(TypeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlock

	return p
}

func (s *TypeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockContext) GetTypeToken() antlr.Token { return s.typeToken }

func (s *TypeBlockContext) GetLp() antlr.Token { return s.lp }

func (s *TypeBlockContext) GetRp() antlr.Token { return s.rp }

func (s *TypeBlockContext) SetTypeToken(v antlr.Token) { s.typeToken = v }

func (s *TypeBlockContext) SetLp(v antlr.Token) { s.lp = v }

func (s *TypeBlockContext) SetRp(v antlr.Token) { s.rp = v }

func (s *TypeBlockContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeBlockContext) AllTypeBlockBody() []ITypeBlockBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeBlockBodyContext); ok {
			len++
		}
	}

	tst := make([]ITypeBlockBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeBlockBodyContext); ok {
			tst[i] = t.(ITypeBlockBodyContext)
			i++
		}
	}

	return tst
}

func (s *TypeBlockContext) TypeBlockBody(i int) ITypeBlockBodyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockBodyContext)
}

func (s *TypeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeBlock(s)
	}
}

func (s *TypeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeBlock(s)
	}
}

func (s *TypeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlock() (localctx ITypeBlockContext) {
	localctx = NewTypeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ApiParserParserRULE_typeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	match(p, "type")
	{
		p.SetState(143)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockContext).typeToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*TypeBlockContext).lp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ApiParserParserID {
		{
			p.SetState(145)
			p.TypeBlockBody()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*TypeBlockContext).rp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeLitBodyContext is an interface to support dynamic dispatch.
type ITypeLitBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeStruct() ITypeStructContext
	TypeAlias() ITypeAliasContext

	// IsTypeLitBodyContext differentiates from other interfaces.
	IsTypeLitBodyContext()
}

type TypeLitBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeLitBodyContext() *TypeLitBodyContext {
	var p = new(TypeLitBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLitBody
	return p
}

func InitEmptyTypeLitBodyContext(p *TypeLitBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeLitBody
}

func (*TypeLitBodyContext) IsTypeLitBodyContext() {}

func NewTypeLitBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeLitBodyContext {
	var p = new(TypeLitBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeLitBody

	return p
}

func (s *TypeLitBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeLitBodyContext) TypeStruct() ITypeStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
}

func (s *TypeLitBodyContext) TypeAlias() ITypeAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeAliasContext)
}

func (s *TypeLitBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeLitBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeLitBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeLitBody(s)
	}
}

func (s *TypeLitBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeLitBody(s)
	}
}

func (s *TypeLitBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeLitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeLitBody() (localctx ITypeLitBodyContext) {
	localctx = NewTypeLitBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ApiParserParserRULE_typeLitBody)
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.TypeStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.TypeAlias()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeBlockBodyContext is an interface to support dynamic dispatch.
type ITypeBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypeBlockStruct() ITypeBlockStructContext
	TypeBlockAlias() ITypeBlockAliasContext

	// IsTypeBlockBodyContext differentiates from other interfaces.
	IsTypeBlockBodyContext()
}

type TypeBlockBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBlockBodyContext() *TypeBlockBodyContext {
	var p = new(TypeBlockBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockBody
	return p
}

func InitEmptyTypeBlockBodyContext(p *TypeBlockBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockBody
}

func (*TypeBlockBodyContext) IsTypeBlockBodyContext() {}

func NewTypeBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockBodyContext {
	var p = new(TypeBlockBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockBody

	return p
}

func (s *TypeBlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockBodyContext) TypeBlockStruct() ITypeBlockStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockStructContext)
}

func (s *TypeBlockBodyContext) TypeBlockAlias() ITypeBlockAliasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeBlockAliasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeBlockAliasContext)
}

func (s *TypeBlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeBlockBody(s)
	}
}

func (s *TypeBlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeBlockBody(s)
	}
}

func (s *TypeBlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockBody() (localctx ITypeBlockBodyContext) {
	localctx = NewTypeBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ApiParserParserRULE_typeBlockBody)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.TypeBlockStruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.TypeBlockAlias()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeStructContext is an interface to support dynamic dispatch.
type ITypeStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// GetStructToken returns the structToken token.
	GetStructToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// SetStructToken sets the structToken token.
	SetStructToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsTypeStructContext differentiates from other interfaces.
	IsTypeStructContext()
}

type TypeStructContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeStructContext() *TypeStructContext {
	var p = new(TypeStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeStruct
	return p
}

func InitEmptyTypeStructContext(p *TypeStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeStruct
}

func (*TypeStructContext) IsTypeStructContext() {}

func NewTypeStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeStructContext {
	var p = new(TypeStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeStruct

	return p
}

func (s *TypeStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeStructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeStructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeStruct(s)
	}
}

func (s *TypeStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeStruct(s)
	}
}

func (s *TypeStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeStruct() (localctx ITypeStructContext) {
	localctx = NewTypeStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ApiParserParserRULE_typeStruct)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(162)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeStructContext).structName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(163)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeStructContext).structToken = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(166)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeStructContext).lbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(167)
				p.Field()
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(173)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeStructContext).rbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeAliasContext is an interface to support dynamic dispatch.
type ITypeAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext
	ID() antlr.TerminalNode

	// IsTypeAliasContext differentiates from other interfaces.
	IsTypeAliasContext()
}

type TypeAliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeAliasContext() *TypeAliasContext {
	var p = new(TypeAliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeAlias
	return p
}

func InitEmptyTypeAliasContext(p *TypeAliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeAlias
}

func (*TypeAliasContext) IsTypeAliasContext() {}

func NewTypeAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeAliasContext {
	var p = new(TypeAliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeAlias

	return p
}

func (s *TypeAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeAliasContext) GetAssign() antlr.Token { return s.assign }

func (s *TypeAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeAliasContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *TypeAliasContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeAlias(s)
	}
}

func (s *TypeAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeAlias(s)
	}
}

func (s *TypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeAlias() (localctx ITypeAliasContext) {
	localctx = NewTypeAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ApiParserParserRULE_typeAlias)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(176)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeAliasContext).alias = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(177)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeAliasContext).assign = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(180)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeBlockStructContext is an interface to support dynamic dispatch.
type ITypeBlockStructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStructName returns the structName token.
	GetStructName() antlr.Token

	// GetStructToken returns the structToken token.
	GetStructToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetStructName sets the structName token.
	SetStructName(antlr.Token)

	// SetStructToken sets the structToken token.
	SetStructToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllField() []IFieldContext
	Field(i int) IFieldContext

	// IsTypeBlockStructContext differentiates from other interfaces.
	IsTypeBlockStructContext()
}

type TypeBlockStructContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	structName  antlr.Token
	structToken antlr.Token
	lbrace      antlr.Token
	rbrace      antlr.Token
}

func NewEmptyTypeBlockStructContext() *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct
	return p
}

func InitEmptyTypeBlockStructContext(p *TypeBlockStructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct
}

func (*TypeBlockStructContext) IsTypeBlockStructContext() {}

func NewTypeBlockStructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockStructContext {
	var p = new(TypeBlockStructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockStruct

	return p
}

func (s *TypeBlockStructContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockStructContext) GetStructName() antlr.Token { return s.structName }

func (s *TypeBlockStructContext) GetStructToken() antlr.Token { return s.structToken }

func (s *TypeBlockStructContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *TypeBlockStructContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *TypeBlockStructContext) SetStructName(v antlr.Token) { s.structName = v }

func (s *TypeBlockStructContext) SetStructToken(v antlr.Token) { s.structToken = v }

func (s *TypeBlockStructContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *TypeBlockStructContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *TypeBlockStructContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *TypeBlockStructContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *TypeBlockStructContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *TypeBlockStructContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeBlockStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockStructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeBlockStruct(s)
	}
}

func (s *TypeBlockStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeBlockStruct(s)
	}
}

func (s *TypeBlockStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockStruct() (localctx ITypeBlockStructContext) {
	localctx = NewTypeBlockStructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ApiParserParserRULE_typeBlockStruct)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(183)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockStructContext).structName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(184)

			var _m = p.Match(ApiParserParserID)

			localctx.(*TypeBlockStructContext).structToken = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(187)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*TypeBlockStructContext).lbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(188)
				p.Field()
			}

		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(194)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*TypeBlockStructContext).rbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeBlockAliasContext is an interface to support dynamic dispatch.
type ITypeBlockAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// GetAssign returns the assign token.
	GetAssign() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// SetAssign sets the assign token.
	SetAssign(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext
	ID() antlr.TerminalNode

	// IsTypeBlockAliasContext differentiates from other interfaces.
	IsTypeBlockAliasContext()
}

type TypeBlockAliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
	assign antlr.Token
}

func NewEmptyTypeBlockAliasContext() *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias
	return p
}

func InitEmptyTypeBlockAliasContext(p *TypeBlockAliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias
}

func (*TypeBlockAliasContext) IsTypeBlockAliasContext() {}

func NewTypeBlockAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBlockAliasContext {
	var p = new(TypeBlockAliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_typeBlockAlias

	return p
}

func (s *TypeBlockAliasContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBlockAliasContext) GetAlias() antlr.Token { return s.alias }

func (s *TypeBlockAliasContext) GetAssign() antlr.Token { return s.assign }

func (s *TypeBlockAliasContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *TypeBlockAliasContext) SetAssign(v antlr.Token) { s.assign = v }

func (s *TypeBlockAliasContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *TypeBlockAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *TypeBlockAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBlockAliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBlockAliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterTypeBlockAlias(s)
	}
}

func (s *TypeBlockAliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitTypeBlockAlias(s)
	}
}

func (s *TypeBlockAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitTypeBlockAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) TypeBlockAlias() (localctx ITypeBlockAliasContext) {
	localctx = NewTypeBlockAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ApiParserParserRULE_typeBlockAlias)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(197)

		var _m = p.Match(ApiParserParserID)

		localctx.(*TypeBlockAliasContext).alias = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__0 {
		{
			p.SetState(198)

			var _m = p.Match(ApiParserParserT__0)

			localctx.(*TypeBlockAliasContext).assign = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(201)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NormalField() INormalFieldContext
	AnonymousFiled() IAnonymousFiledContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) NormalField() INormalFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormalFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INormalFieldContext)
}

func (s *FieldContext) AnonymousFiled() IAnonymousFiledContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonymousFiledContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonymousFiledContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ApiParserParserRULE_field)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(203)

		if !(isNormal(p)) {
			p.SetError(antlr.NewFailedPredicateException(p, "isNormal(p)", ""))
			goto errorExit
		}
		{
			p.SetState(204)
			p.NormalField()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.AnonymousFiled()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INormalFieldContext is an interface to support dynamic dispatch.
type INormalFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFieldName returns the fieldName token.
	GetFieldName() antlr.Token

	// GetTag returns the tag token.
	GetTag() antlr.Token

	// SetFieldName sets the fieldName token.
	SetFieldName(antlr.Token)

	// SetTag sets the tag token.
	SetTag(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext
	ID() antlr.TerminalNode
	RAW_STRING() antlr.TerminalNode

	// IsNormalFieldContext differentiates from other interfaces.
	IsNormalFieldContext()
}

type NormalFieldContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	fieldName antlr.Token
	tag       antlr.Token
}

func NewEmptyNormalFieldContext() *NormalFieldContext {
	var p = new(NormalFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_normalField
	return p
}

func InitEmptyNormalFieldContext(p *NormalFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_normalField
}

func (*NormalFieldContext) IsNormalFieldContext() {}

func NewNormalFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalFieldContext {
	var p = new(NormalFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_normalField

	return p
}

func (s *NormalFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalFieldContext) GetFieldName() antlr.Token { return s.fieldName }

func (s *NormalFieldContext) GetTag() antlr.Token { return s.tag }

func (s *NormalFieldContext) SetFieldName(v antlr.Token) { s.fieldName = v }

func (s *NormalFieldContext) SetTag(v antlr.Token) { s.tag = v }

func (s *NormalFieldContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *NormalFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *NormalFieldContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserRAW_STRING, 0)
}

func (s *NormalFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterNormalField(s)
	}
}

func (s *NormalFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitNormalField(s)
	}
}

func (s *NormalFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitNormalField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) NormalField() (localctx INormalFieldContext) {
	localctx = NewNormalFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ApiParserParserRULE_normalField)
	p.EnterOuterAlt(localctx, 1)
	checkKeyword(p)
	{
		p.SetState(209)

		var _m = p.Match(ApiParserParserID)

		localctx.(*NormalFieldContext).fieldName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.DataType()
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(211)

			var _m = p.Match(ApiParserParserRAW_STRING)

			localctx.(*NormalFieldContext).tag = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAnonymousFiledContext is an interface to support dynamic dispatch.
type IAnonymousFiledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// Getter signatures
	QualifiedType() IQualifiedTypeContext
	ID() antlr.TerminalNode

	// IsAnonymousFiledContext differentiates from other interfaces.
	IsAnonymousFiledContext()
}

type AnonymousFiledContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptyAnonymousFiledContext() *AnonymousFiledContext {
	var p = new(AnonymousFiledContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_anonymousFiled
	return p
}

func InitEmptyAnonymousFiledContext(p *AnonymousFiledContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_anonymousFiled
}

func (*AnonymousFiledContext) IsAnonymousFiledContext() {}

func NewAnonymousFiledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymousFiledContext {
	var p = new(AnonymousFiledContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_anonymousFiled

	return p
}

func (s *AnonymousFiledContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonymousFiledContext) GetStar() antlr.Token { return s.star }

func (s *AnonymousFiledContext) SetStar(v antlr.Token) { s.star = v }

func (s *AnonymousFiledContext) QualifiedType() IQualifiedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedTypeContext)
}

func (s *AnonymousFiledContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *AnonymousFiledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonymousFiledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonymousFiledContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAnonymousFiled(s)
	}
}

func (s *AnonymousFiledContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAnonymousFiled(s)
	}
}

func (s *AnonymousFiledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAnonymousFiled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AnonymousFiled() (localctx IAnonymousFiledContext) {
	localctx = NewAnonymousFiledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ApiParserParserRULE_anonymousFiled)
	var _la int

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__5 {
			{
				p.SetState(214)

				var _m = p.Match(ApiParserParserT__5)

				localctx.(*AnonymousFiledContext).star = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(217)
			p.QualifiedType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(ApiParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInter returns the inter token.
	GetInter() antlr.Token

	// SetInter sets the inter token.
	SetInter(antlr.Token)

	// Getter signatures
	QualifiedType() IQualifiedTypeContext
	ID() antlr.TerminalNode
	MapType() IMapTypeContext
	ArrayType() IArrayTypeContext
	INTERFACE() antlr.TerminalNode
	PointerType() IPointerTypeContext
	TypeStruct() ITypeStructContext

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	inter  antlr.Token
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_dataType
	return p
}

func InitEmptyDataTypeContext(p *DataTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_dataType
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) GetInter() antlr.Token { return s.inter }

func (s *DataTypeContext) SetInter(v antlr.Token) { s.inter = v }

func (s *DataTypeContext) QualifiedType() IQualifiedTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedTypeContext)
}

func (s *DataTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *DataTypeContext) MapType() IMapTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *DataTypeContext) ArrayType() IArrayTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayTypeContext)
}

func (s *DataTypeContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(ApiParserParserINTERFACE, 0)
}

func (s *DataTypeContext) PointerType() IPointerTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPointerTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPointerTypeContext)
}

func (s *DataTypeContext) TypeStruct() ITypeStructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeStructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeStructContext)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (s *DataTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitDataType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ApiParserParserRULE_dataType)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.QualifiedType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(ApiParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(223)
			p.MapType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(224)
			p.ArrayType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(225)

			var _m = p.Match(ApiParserParserINTERFACE)

			localctx.(*DataTypeContext).inter = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(226)
			p.PointerType()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(227)
			p.TypeStruct()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQualifiedTypeContext is an interface to support dynamic dispatch.
type IQualifiedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPkg returns the pkg token.
	GetPkg() antlr.Token

	// GetDot returns the dot token.
	GetDot() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetPkg sets the pkg token.
	SetPkg(antlr.Token)

	// SetDot sets the dot token.
	SetDot(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsQualifiedTypeContext differentiates from other interfaces.
	IsQualifiedTypeContext()
}

type QualifiedTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	pkg    antlr.Token
	dot    antlr.Token
	name   antlr.Token
}

func NewEmptyQualifiedTypeContext() *QualifiedTypeContext {
	var p = new(QualifiedTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_qualifiedType
	return p
}

func InitEmptyQualifiedTypeContext(p *QualifiedTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_qualifiedType
}

func (*QualifiedTypeContext) IsQualifiedTypeContext() {}

func NewQualifiedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedTypeContext {
	var p = new(QualifiedTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_qualifiedType

	return p
}

func (s *QualifiedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedTypeContext) GetPkg() antlr.Token { return s.pkg }

func (s *QualifiedTypeContext) GetDot() antlr.Token { return s.dot }

func (s *QualifiedTypeContext) GetName() antlr.Token { return s.name }

func (s *QualifiedTypeContext) SetPkg(v antlr.Token) { s.pkg = v }

func (s *QualifiedTypeContext) SetDot(v antlr.Token) { s.dot = v }

func (s *QualifiedTypeContext) SetName(v antlr.Token) { s.name = v }

func (s *QualifiedTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *QualifiedTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *QualifiedTypeContext) DOT() antlr.TerminalNode {
	return s.GetToken(ApiParserParserDOT, 0)
}

func (s *QualifiedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterQualifiedType(s)
	}
}

func (s *QualifiedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitQualifiedType(s)
	}
}

func (s *QualifiedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitQualifiedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) QualifiedType() (localctx IQualifiedTypeContext) {
	localctx = NewQualifiedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ApiParserParserRULE_qualifiedType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)

		var _m = p.Match(ApiParserParserID)

		localctx.(*QualifiedTypeContext).pkg = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)

		var _m = p.Match(ApiParserParserDOT)

		localctx.(*QualifiedTypeContext).dot = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)

		var _m = p.Match(ApiParserParserID)

		localctx.(*QualifiedTypeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPointerTypeContext is an interface to support dynamic dispatch.
type IPointerTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token

	// SetStar sets the star token.
	SetStar(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext

	// IsPointerTypeContext differentiates from other interfaces.
	IsPointerTypeContext()
}

type PointerTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	star   antlr.Token
}

func NewEmptyPointerTypeContext() *PointerTypeContext {
	var p = new(PointerTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_pointerType
	return p
}

func InitEmptyPointerTypeContext(p *PointerTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_pointerType
}

func (*PointerTypeContext) IsPointerTypeContext() {}

func NewPointerTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointerTypeContext {
	var p = new(PointerTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_pointerType

	return p
}

func (s *PointerTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PointerTypeContext) GetStar() antlr.Token { return s.star }

func (s *PointerTypeContext) SetStar(v antlr.Token) { s.star = v }

func (s *PointerTypeContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *PointerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointerTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterPointerType(s)
	}
}

func (s *PointerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitPointerType(s)
	}
}

func (s *PointerTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPointerType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) PointerType() (localctx IPointerTypeContext) {
	localctx = NewPointerTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ApiParserParserRULE_pointerType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)

		var _m = p.Match(ApiParserParserT__5)

		localctx.(*PointerTypeContext).star = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMapToken returns the mapToken token.
	GetMapToken() antlr.Token

	// GetLbrack returns the lbrack token.
	GetLbrack() antlr.Token

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetRbrack returns the rbrack token.
	GetRbrack() antlr.Token

	// SetMapToken sets the mapToken token.
	SetMapToken(antlr.Token)

	// SetLbrack sets the lbrack token.
	SetLbrack(antlr.Token)

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetRbrack sets the rbrack token.
	SetRbrack(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IDataTypeContext

	// SetValue sets the value rule contexts.
	SetValue(IDataTypeContext)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DataType() IDataTypeContext

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	mapToken antlr.Token
	lbrack   antlr.Token
	key      antlr.Token
	rbrack   antlr.Token
	value    IDataTypeContext
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_mapType
	return p
}

func InitEmptyMapTypeContext(p *MapTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_mapType
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) GetMapToken() antlr.Token { return s.mapToken }

func (s *MapTypeContext) GetLbrack() antlr.Token { return s.lbrack }

func (s *MapTypeContext) GetKey() antlr.Token { return s.key }

func (s *MapTypeContext) GetRbrack() antlr.Token { return s.rbrack }

func (s *MapTypeContext) SetMapToken(v antlr.Token) { s.mapToken = v }

func (s *MapTypeContext) SetLbrack(v antlr.Token) { s.lbrack = v }

func (s *MapTypeContext) SetKey(v antlr.Token) { s.key = v }

func (s *MapTypeContext) SetRbrack(v antlr.Token) { s.rbrack = v }

func (s *MapTypeContext) GetValue() IDataTypeContext { return s.value }

func (s *MapTypeContext) SetValue(v IDataTypeContext) { s.value = v }

func (s *MapTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *MapTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *MapTypeContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterMapType(s)
	}
}

func (s *MapTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitMapType(s)
	}
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ApiParserParserRULE_mapType)
	p.EnterOuterAlt(localctx, 1)
	match(p, "map")
	{
		p.SetState(238)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).mapToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)

		var _m = p.Match(ApiParserParserT__6)

		localctx.(*MapTypeContext).lbrack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	checkKey(p)
	{
		p.SetState(241)

		var _m = p.Match(ApiParserParserID)

		localctx.(*MapTypeContext).key = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*MapTypeContext).rbrack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)

		var _x = p.DataType()

		localctx.(*MapTypeContext).value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayTypeContext is an interface to support dynamic dispatch.
type IArrayTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLbrack returns the lbrack token.
	GetLbrack() antlr.Token

	// GetRbrack returns the rbrack token.
	GetRbrack() antlr.Token

	// SetLbrack sets the lbrack token.
	SetLbrack(antlr.Token)

	// SetRbrack sets the rbrack token.
	SetRbrack(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext

	// IsArrayTypeContext differentiates from other interfaces.
	IsArrayTypeContext()
}

type ArrayTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lbrack antlr.Token
	rbrack antlr.Token
}

func NewEmptyArrayTypeContext() *ArrayTypeContext {
	var p = new(ArrayTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_arrayType
	return p
}

func InitEmptyArrayTypeContext(p *ArrayTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_arrayType
}

func (*ArrayTypeContext) IsArrayTypeContext() {}

func NewArrayTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_arrayType

	return p
}

func (s *ArrayTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayTypeContext) GetLbrack() antlr.Token { return s.lbrack }

func (s *ArrayTypeContext) GetRbrack() antlr.Token { return s.rbrack }

func (s *ArrayTypeContext) SetLbrack(v antlr.Token) { s.lbrack = v }

func (s *ArrayTypeContext) SetRbrack(v antlr.Token) { s.rbrack = v }

func (s *ArrayTypeContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ArrayType() (localctx IArrayTypeContext) {
	localctx = NewArrayTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ApiParserParserRULE_arrayType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)

		var _m = p.Match(ApiParserParserT__6)

		localctx.(*ArrayTypeContext).lbrack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)

		var _m = p.Match(ApiParserParserT__7)

		localctx.(*ArrayTypeContext).rbrack = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.DataType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceSpecContext is an interface to support dynamic dispatch.
type IServiceSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ServiceApi() IServiceApiContext
	AtServer() IAtServerContext

	// IsServiceSpecContext differentiates from other interfaces.
	IsServiceSpecContext()
}

type ServiceSpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceSpecContext() *ServiceSpecContext {
	var p = new(ServiceSpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceSpec
	return p
}

func InitEmptyServiceSpecContext(p *ServiceSpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceSpec
}

func (*ServiceSpecContext) IsServiceSpecContext() {}

func NewServiceSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceSpecContext {
	var p = new(ServiceSpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceSpec

	return p
}

func (s *ServiceSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceSpecContext) ServiceApi() IServiceApiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceApiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceApiContext)
}

func (s *ServiceSpecContext) AtServer() IAtServerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtServerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}

func (s *ServiceSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceSpec(s)
	}
}

func (s *ServiceSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceSpec(s)
	}
}

func (s *ServiceSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceSpec() (localctx IServiceSpecContext) {
	localctx = NewServiceSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ApiParserParserRULE_serviceSpec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATSERVER {
		{
			p.SetState(249)
			p.AtServer()
		}

	}
	{
		p.SetState(252)
		p.ServiceApi()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtServerContext is an interface to support dynamic dispatch.
type IAtServerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ATSERVER() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsAtServerContext differentiates from other interfaces.
	IsAtServerContext()
}

type AtServerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtServerContext() *AtServerContext {
	var p = new(AtServerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atServer
	return p
}

func InitEmptyAtServerContext(p *AtServerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atServer
}

func (*AtServerContext) IsAtServerContext() {}

func NewAtServerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtServerContext {
	var p = new(AtServerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atServer

	return p
}

func (s *AtServerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtServerContext) GetLp() antlr.Token { return s.lp }

func (s *AtServerContext) GetRp() antlr.Token { return s.rp }

func (s *AtServerContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtServerContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtServerContext) ATSERVER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATSERVER, 0)
}

func (s *AtServerContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *AtServerContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtServerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtServerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtServerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtServer(s)
	}
}

func (s *AtServerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtServer(s)
	}
}

func (s *AtServerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtServer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtServer() (localctx IAtServerContext) {
	localctx = NewAtServerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ApiParserParserRULE_atServer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(ApiParserParserATSERVER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*AtServerContext).lp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(256)
			p.KvLit()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*AtServerContext).rp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceApiContext is an interface to support dynamic dispatch.
type IServiceApiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetServiceToken returns the serviceToken token.
	GetServiceToken() antlr.Token

	// GetLbrace returns the lbrace token.
	GetLbrace() antlr.Token

	// GetRbrace returns the rbrace token.
	GetRbrace() antlr.Token

	// SetServiceToken sets the serviceToken token.
	SetServiceToken(antlr.Token)

	// SetLbrace sets the lbrace token.
	SetLbrace(antlr.Token)

	// SetRbrace sets the rbrace token.
	SetRbrace(antlr.Token)

	// Getter signatures
	ServiceName() IServiceNameContext
	ID() antlr.TerminalNode
	AllServiceRoute() []IServiceRouteContext
	ServiceRoute(i int) IServiceRouteContext

	// IsServiceApiContext differentiates from other interfaces.
	IsServiceApiContext()
}

type ServiceApiContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	serviceToken antlr.Token
	lbrace       antlr.Token
	rbrace       antlr.Token
}

func NewEmptyServiceApiContext() *ServiceApiContext {
	var p = new(ServiceApiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceApi
	return p
}

func InitEmptyServiceApiContext(p *ServiceApiContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceApi
}

func (*ServiceApiContext) IsServiceApiContext() {}

func NewServiceApiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceApiContext {
	var p = new(ServiceApiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceApi

	return p
}

func (s *ServiceApiContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceApiContext) GetServiceToken() antlr.Token { return s.serviceToken }

func (s *ServiceApiContext) GetLbrace() antlr.Token { return s.lbrace }

func (s *ServiceApiContext) GetRbrace() antlr.Token { return s.rbrace }

func (s *ServiceApiContext) SetServiceToken(v antlr.Token) { s.serviceToken = v }

func (s *ServiceApiContext) SetLbrace(v antlr.Token) { s.lbrace = v }

func (s *ServiceApiContext) SetRbrace(v antlr.Token) { s.rbrace = v }

func (s *ServiceApiContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceApiContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *ServiceApiContext) AllServiceRoute() []IServiceRouteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceRouteContext); ok {
			len++
		}
	}

	tst := make([]IServiceRouteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceRouteContext); ok {
			tst[i] = t.(IServiceRouteContext)
			i++
		}
	}

	return tst
}

func (s *ServiceApiContext) ServiceRoute(i int) IServiceRouteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceRouteContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceRouteContext)
}

func (s *ServiceApiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceApiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceApiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceApi(s)
	}
}

func (s *ServiceApiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceApi(s)
	}
}

func (s *ServiceApiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceApi(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceApi() (localctx IServiceApiContext) {
	localctx = NewServiceApiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ApiParserParserRULE_serviceApi)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	match(p, "service")
	{
		p.SetState(264)

		var _m = p.Match(ApiParserParserID)

		localctx.(*ServiceApiContext).serviceToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.ServiceName()
	}
	{
		p.SetState(266)

		var _m = p.Match(ApiParserParserT__3)

		localctx.(*ServiceApiContext).lbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&483328) != 0 {
		{
			p.SetState(267)
			p.ServiceRoute()
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(273)

		var _m = p.Match(ApiParserParserT__4)

		localctx.(*ServiceApiContext).rbrace = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceRouteContext is an interface to support dynamic dispatch.
type IServiceRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Route() IRouteContext
	AtServer() IAtServerContext
	AtHandler() IAtHandlerContext
	AtDoc() IAtDocContext
	AtRoles() IAtRolesContext
	AtFeatures() IAtFeaturesContext

	// IsServiceRouteContext differentiates from other interfaces.
	IsServiceRouteContext()
}

type ServiceRouteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceRouteContext() *ServiceRouteContext {
	var p = new(ServiceRouteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceRoute
	return p
}

func InitEmptyServiceRouteContext(p *ServiceRouteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceRoute
}

func (*ServiceRouteContext) IsServiceRouteContext() {}

func NewServiceRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceRouteContext {
	var p = new(ServiceRouteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceRoute

	return p
}

func (s *ServiceRouteContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceRouteContext) Route() IRouteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRouteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRouteContext)
}

func (s *ServiceRouteContext) AtServer() IAtServerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtServerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtServerContext)
}

func (s *ServiceRouteContext) AtHandler() IAtHandlerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtHandlerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtHandlerContext)
}

func (s *ServiceRouteContext) AtDoc() IAtDocContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtDocContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtDocContext)
}

func (s *ServiceRouteContext) AtRoles() IAtRolesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtRolesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtRolesContext)
}

func (s *ServiceRouteContext) AtFeatures() IAtFeaturesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtFeaturesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtFeaturesContext)
}

func (s *ServiceRouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceRouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceRouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceRoute(s)
	}
}

func (s *ServiceRouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceRoute(s)
	}
}

func (s *ServiceRouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceRoute() (localctx IServiceRouteContext) {
	localctx = NewServiceRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ApiParserParserRULE_serviceRoute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATDOC {
		{
			p.SetState(275)
			p.AtDoc()
		}

	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATROLES {
		{
			p.SetState(278)
			p.AtRoles()
		}

	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserATFEATURES {
		{
			p.SetState(281)
			p.AtFeatures()
		}

	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserATSERVER:
		{
			p.SetState(284)
			p.AtServer()
		}

	case ApiParserParserATHANDLER:
		{
			p.SetState(285)
			p.AtHandler()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(288)
		p.Route()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtDocContext is an interface to support dynamic dispatch.
type IAtDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ATDOC() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsAtDocContext differentiates from other interfaces.
	IsAtDocContext()
}

type AtDocContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtDocContext() *AtDocContext {
	var p = new(AtDocContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atDoc
	return p
}

func InitEmptyAtDocContext(p *AtDocContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atDoc
}

func (*AtDocContext) IsAtDocContext() {}

func NewAtDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDocContext {
	var p = new(AtDocContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atDoc

	return p
}

func (s *AtDocContext) GetParser() antlr.Parser { return s.parser }

func (s *AtDocContext) GetLp() antlr.Token { return s.lp }

func (s *AtDocContext) GetRp() antlr.Token { return s.rp }

func (s *AtDocContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtDocContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtDocContext) ATDOC() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATDOC, 0)
}

func (s *AtDocContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *AtDocContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *AtDocContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtDocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtDocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtDoc(s)
	}
}

func (s *AtDocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtDoc(s)
	}
}

func (s *AtDocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtDoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtDoc() (localctx IAtDocContext) {
	localctx = NewAtDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ApiParserParserRULE_atDoc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(ApiParserParserATDOC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__1 {
		{
			p.SetState(291)

			var _m = p.Match(ApiParserParserT__1)

			localctx.(*AtDocContext).lp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserID:
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserID {
			{
				p.SetState(294)
				p.KvLit()
			}

			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ApiParserParserSTRING:
		{
			p.SetState(299)
			p.Match(ApiParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__2 {
		{
			p.SetState(302)

			var _m = p.Match(ApiParserParserT__2)

			localctx.(*AtDocContext).rp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtHandlerContext is an interface to support dynamic dispatch.
type IAtHandlerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATHANDLER() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsAtHandlerContext differentiates from other interfaces.
	IsAtHandlerContext()
}

type AtHandlerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtHandlerContext() *AtHandlerContext {
	var p = new(AtHandlerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atHandler
	return p
}

func InitEmptyAtHandlerContext(p *AtHandlerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atHandler
}

func (*AtHandlerContext) IsAtHandlerContext() {}

func NewAtHandlerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtHandlerContext {
	var p = new(AtHandlerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atHandler

	return p
}

func (s *AtHandlerContext) GetParser() antlr.Parser { return s.parser }

func (s *AtHandlerContext) ATHANDLER() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATHANDLER, 0)
}

func (s *AtHandlerContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *AtHandlerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtHandlerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtHandlerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtHandler(s)
	}
}

func (s *AtHandlerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtHandler(s)
	}
}

func (s *AtHandlerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtHandler(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtHandler() (localctx IAtHandlerContext) {
	localctx = NewAtHandlerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ApiParserParserRULE_atHandler)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(ApiParserParserATHANDLER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(ApiParserParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtRolesContext is an interface to support dynamic dispatch.
type IAtRolesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ATROLES() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsAtRolesContext differentiates from other interfaces.
	IsAtRolesContext()
}

type AtRolesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtRolesContext() *AtRolesContext {
	var p = new(AtRolesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atRoles
	return p
}

func InitEmptyAtRolesContext(p *AtRolesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atRoles
}

func (*AtRolesContext) IsAtRolesContext() {}

func NewAtRolesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtRolesContext {
	var p = new(AtRolesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atRoles

	return p
}

func (s *AtRolesContext) GetParser() antlr.Parser { return s.parser }

func (s *AtRolesContext) GetLp() antlr.Token { return s.lp }

func (s *AtRolesContext) GetRp() antlr.Token { return s.rp }

func (s *AtRolesContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtRolesContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtRolesContext) ATROLES() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATROLES, 0)
}

func (s *AtRolesContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *AtRolesContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *AtRolesContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtRolesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtRolesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtRolesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtRoles(s)
	}
}

func (s *AtRolesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtRoles(s)
	}
}

func (s *AtRolesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtRoles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtRoles() (localctx IAtRolesContext) {
	localctx = NewAtRolesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ApiParserParserRULE_atRoles)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(ApiParserParserATROLES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__6 {
		{
			p.SetState(309)

			var _m = p.Match(ApiParserParserT__6)

			localctx.(*AtRolesContext).lp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserID:
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserID {
			{
				p.SetState(312)
				p.KvLit()
			}

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ApiParserParserSTRING:
		{
			p.SetState(317)
			p.Match(ApiParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__7 {
		{
			p.SetState(320)

			var _m = p.Match(ApiParserParserT__7)

			localctx.(*AtRolesContext).rp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtFeaturesContext is an interface to support dynamic dispatch.
type IAtFeaturesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ATFEATURES() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllKvLit() []IKvLitContext
	KvLit(i int) IKvLitContext

	// IsAtFeaturesContext differentiates from other interfaces.
	IsAtFeaturesContext()
}

type AtFeaturesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyAtFeaturesContext() *AtFeaturesContext {
	var p = new(AtFeaturesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atFeatures
	return p
}

func InitEmptyAtFeaturesContext(p *AtFeaturesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_atFeatures
}

func (*AtFeaturesContext) IsAtFeaturesContext() {}

func NewAtFeaturesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtFeaturesContext {
	var p = new(AtFeaturesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_atFeatures

	return p
}

func (s *AtFeaturesContext) GetParser() antlr.Parser { return s.parser }

func (s *AtFeaturesContext) GetLp() antlr.Token { return s.lp }

func (s *AtFeaturesContext) GetRp() antlr.Token { return s.rp }

func (s *AtFeaturesContext) SetLp(v antlr.Token) { s.lp = v }

func (s *AtFeaturesContext) SetRp(v antlr.Token) { s.rp = v }

func (s *AtFeaturesContext) ATFEATURES() antlr.TerminalNode {
	return s.GetToken(ApiParserParserATFEATURES, 0)
}

func (s *AtFeaturesContext) STRING() antlr.TerminalNode {
	return s.GetToken(ApiParserParserSTRING, 0)
}

func (s *AtFeaturesContext) AllKvLit() []IKvLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKvLitContext); ok {
			len++
		}
	}

	tst := make([]IKvLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKvLitContext); ok {
			tst[i] = t.(IKvLitContext)
			i++
		}
	}

	return tst
}

func (s *AtFeaturesContext) KvLit(i int) IKvLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKvLitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKvLitContext)
}

func (s *AtFeaturesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtFeaturesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtFeaturesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterAtFeatures(s)
	}
}

func (s *AtFeaturesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitAtFeatures(s)
	}
}

func (s *AtFeaturesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitAtFeatures(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) AtFeatures() (localctx IAtFeaturesContext) {
	localctx = NewAtFeaturesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ApiParserParserRULE_atFeatures)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(ApiParserParserATFEATURES)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__6 {
		{
			p.SetState(324)

			var _m = p.Match(ApiParserParserT__6)

			localctx.(*AtFeaturesContext).lp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ApiParserParserID:
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserID {
			{
				p.SetState(327)
				p.KvLit()
			}

			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ApiParserParserSTRING:
		{
			p.SetState(332)
			p.Match(ApiParserParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__7 {
		{
			p.SetState(335)

			var _m = p.Match(ApiParserParserT__7)

			localctx.(*AtFeaturesContext).rp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRouteContext is an interface to support dynamic dispatch.
type IRouteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHttpMethod returns the httpMethod token.
	GetHttpMethod() antlr.Token

	// SetHttpMethod sets the httpMethod token.
	SetHttpMethod(antlr.Token)

	// GetRequest returns the request rule contexts.
	GetRequest() IBodyContext

	// GetResponse returns the response rule contexts.
	GetResponse() IReplybodyContext

	// SetRequest sets the request rule contexts.
	SetRequest(IBodyContext)

	// SetResponse sets the response rule contexts.
	SetResponse(IReplybodyContext)

	// Getter signatures
	Path() IPathContext
	ID() antlr.TerminalNode
	Body() IBodyContext
	Replybody() IReplybodyContext

	// IsRouteContext differentiates from other interfaces.
	IsRouteContext()
}

type RouteContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	httpMethod antlr.Token
	request    IBodyContext
	response   IReplybodyContext
}

func NewEmptyRouteContext() *RouteContext {
	var p = new(RouteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_route
	return p
}

func InitEmptyRouteContext(p *RouteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_route
}

func (*RouteContext) IsRouteContext() {}

func NewRouteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RouteContext {
	var p = new(RouteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_route

	return p
}

func (s *RouteContext) GetParser() antlr.Parser { return s.parser }

func (s *RouteContext) GetHttpMethod() antlr.Token { return s.httpMethod }

func (s *RouteContext) SetHttpMethod(v antlr.Token) { s.httpMethod = v }

func (s *RouteContext) GetRequest() IBodyContext { return s.request }

func (s *RouteContext) GetResponse() IReplybodyContext { return s.response }

func (s *RouteContext) SetRequest(v IBodyContext) { s.request = v }

func (s *RouteContext) SetResponse(v IReplybodyContext) { s.response = v }

func (s *RouteContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *RouteContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *RouteContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *RouteContext) Replybody() IReplybodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReplybodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReplybodyContext)
}

func (s *RouteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RouteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RouteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterRoute(s)
	}
}

func (s *RouteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitRoute(s)
	}
}

func (s *RouteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitRoute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Route() (localctx IRouteContext) {
	localctx = NewRouteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ApiParserParserRULE_route)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	checkHTTPMethod(p)
	{
		p.SetState(339)

		var _m = p.Match(ApiParserParserID)

		localctx.(*RouteContext).httpMethod = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(340)
		p.Path()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__1 {
		{
			p.SetState(341)

			var _x = p.Body()

			localctx.(*RouteContext).request = _x
		}

	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserT__8 {
		{
			p.SetState(344)

			var _x = p.Replybody()

			localctx.(*RouteContext).response = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
	rp     antlr.Token
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) GetLp() antlr.Token { return s.lp }

func (s *BodyContext) GetRp() antlr.Token { return s.rp }

func (s *BodyContext) SetLp(v antlr.Token) { s.lp = v }

func (s *BodyContext) SetRp(v antlr.Token) { s.rp = v }

func (s *BodyContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitBody(s)
	}
}

func (s *BodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ApiParserParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*BodyContext).lp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ApiParserParserID {
		{
			p.SetState(348)
			p.Match(ApiParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(351)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*BodyContext).rp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReplybodyContext is an interface to support dynamic dispatch.
type IReplybodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetReturnToken returns the returnToken token.
	GetReturnToken() antlr.Token

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// GetRp returns the rp token.
	GetRp() antlr.Token

	// SetReturnToken sets the returnToken token.
	SetReturnToken(antlr.Token)

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// SetRp sets the rp token.
	SetRp(antlr.Token)

	// Getter signatures
	DataType() IDataTypeContext

	// IsReplybodyContext differentiates from other interfaces.
	IsReplybodyContext()
}

type ReplybodyContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	returnToken antlr.Token
	lp          antlr.Token
	rp          antlr.Token
}

func NewEmptyReplybodyContext() *ReplybodyContext {
	var p = new(ReplybodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_replybody
	return p
}

func InitEmptyReplybodyContext(p *ReplybodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_replybody
}

func (*ReplybodyContext) IsReplybodyContext() {}

func NewReplybodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplybodyContext {
	var p = new(ReplybodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_replybody

	return p
}

func (s *ReplybodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplybodyContext) GetReturnToken() antlr.Token { return s.returnToken }

func (s *ReplybodyContext) GetLp() antlr.Token { return s.lp }

func (s *ReplybodyContext) GetRp() antlr.Token { return s.rp }

func (s *ReplybodyContext) SetReturnToken(v antlr.Token) { s.returnToken = v }

func (s *ReplybodyContext) SetLp(v antlr.Token) { s.lp = v }

func (s *ReplybodyContext) SetRp(v antlr.Token) { s.rp = v }

func (s *ReplybodyContext) DataType() IDataTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ReplybodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplybodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplybodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterReplybody(s)
	}
}

func (s *ReplybodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitReplybody(s)
	}
}

func (s *ReplybodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitReplybody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Replybody() (localctx IReplybodyContext) {
	localctx = NewReplybodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ApiParserParserRULE_replybody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)

		var _m = p.Match(ApiParserParserT__8)

		localctx.(*ReplybodyContext).returnToken = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(354)

		var _m = p.Match(ApiParserParserT__1)

		localctx.(*ReplybodyContext).lp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33587392) != 0 {
		{
			p.SetState(355)
			p.DataType()
		}

	}
	{
		p.SetState(358)

		var _m = p.Match(ApiParserParserT__2)

		localctx.(*ReplybodyContext).rp = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKvLitContext is an interface to support dynamic dispatch.
type IKvLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	LINE_VALUE() antlr.TerminalNode

	// IsKvLitContext differentiates from other interfaces.
	IsKvLitContext()
}

type KvLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  antlr.Token
}

func NewEmptyKvLitContext() *KvLitContext {
	var p = new(KvLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_kvLit
	return p
}

func InitEmptyKvLitContext(p *KvLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_kvLit
}

func (*KvLitContext) IsKvLitContext() {}

func NewKvLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KvLitContext {
	var p = new(KvLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_kvLit

	return p
}

func (s *KvLitContext) GetParser() antlr.Parser { return s.parser }

func (s *KvLitContext) GetKey() antlr.Token { return s.key }

func (s *KvLitContext) GetValue() antlr.Token { return s.value }

func (s *KvLitContext) SetKey(v antlr.Token) { s.key = v }

func (s *KvLitContext) SetValue(v antlr.Token) { s.value = v }

func (s *KvLitContext) ID() antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, 0)
}

func (s *KvLitContext) LINE_VALUE() antlr.TerminalNode {
	return s.GetToken(ApiParserParserLINE_VALUE, 0)
}

func (s *KvLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KvLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KvLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterKvLit(s)
	}
}

func (s *KvLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitKvLit(s)
	}
}

func (s *KvLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitKvLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) KvLit() (localctx IKvLitContext) {
	localctx = NewKvLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ApiParserParserRULE_kvLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)

		var _m = p.Match(ApiParserParserID)

		localctx.(*KvLitContext).key = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	checkKeyValue(p)
	{
		p.SetState(362)

		var _m = p.Match(ApiParserParserLINE_VALUE)

		localctx.(*KvLitContext).value = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceName
	return p
}

func InitEmptyServiceNameContext(p *ServiceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_serviceName
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *ServiceNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterServiceName(s)
	}
}

func (s *ServiceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitServiceName(s)
	}
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ApiParserParserRULE_serviceName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID {
		{
			p.SetState(364)
			p.Match(ApiParserParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ApiParserParserT__9 {
			{
				p.SetState(365)
				p.Match(ApiParserParserT__9)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPathItem() []IPathItemContext
	PathItem(i int) IPathItemContext

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPathItem() []IPathItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathItemContext); ok {
			len++
		}
	}

	tst := make([]IPathItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathItemContext); ok {
			tst[i] = t.(IPathItemContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) PathItem(i int) IPathItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathItemContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ApiParserParserRULE_path)
	var _la int

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ApiParserParserT__10 || _la == ApiParserParserT__11 {
			p.SetState(387)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case ApiParserParserT__10:
				{
					p.SetState(372)
					p.Match(ApiParserParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(373)
					p.PathItem()
				}
				p.SetState(378)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for _la == ApiParserParserT__9 {
					{
						p.SetState(374)
						p.Match(ApiParserParserT__9)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(375)
						p.PathItem()
					}

					p.SetState(380)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}

			case ApiParserParserT__11:
				{
					p.SetState(381)
					p.Match(ApiParserParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(382)
					p.PathItem()
				}
				p.SetState(385)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ApiParserParserT__9 {
					{
						p.SetState(383)
						p.Match(ApiParserParserT__9)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(384)
						p.PathItem()
					}

				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(389)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)
			p.Match(ApiParserParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPathItemContext is an interface to support dynamic dispatch.
type IPathItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllLetterOrDigit() []antlr.TerminalNode
	LetterOrDigit(i int) antlr.TerminalNode

	// IsPathItemContext differentiates from other interfaces.
	IsPathItemContext()
}

type PathItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathItemContext() *PathItemContext {
	var p = new(PathItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_pathItem
	return p
}

func InitEmptyPathItemContext(p *PathItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ApiParserParserRULE_pathItem
}

func (*PathItemContext) IsPathItemContext() {}

func NewPathItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathItemContext {
	var p = new(PathItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ApiParserParserRULE_pathItem

	return p
}

func (s *PathItemContext) GetParser() antlr.Parser { return s.parser }

func (s *PathItemContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserID)
}

func (s *PathItemContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserID, i)
}

func (s *PathItemContext) AllLetterOrDigit() []antlr.TerminalNode {
	return s.GetTokens(ApiParserParserLetterOrDigit)
}

func (s *PathItemContext) LetterOrDigit(i int) antlr.TerminalNode {
	return s.GetToken(ApiParserParserLetterOrDigit, i)
}

func (s *PathItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.EnterPathItem(s)
	}
}

func (s *PathItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ApiParserListener); ok {
		listenerT.ExitPathItem(s)
	}
}

func (s *PathItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ApiParserVisitor:
		return t.VisitPathItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ApiParserParser) PathItem() (localctx IPathItemContext) {
	localctx = NewPathItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ApiParserParserRULE_pathItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ApiParserParserID || _la == ApiParserParserLetterOrDigit {
		{
			p.SetState(394)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ApiParserParserID || _la == ApiParserParserLetterOrDigit) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ApiParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *FieldContext = nil
		if localctx != nil {
			t = localctx.(*FieldContext)
		}
		return p.Field_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ApiParserParser) Field_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return isNormal(p)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
