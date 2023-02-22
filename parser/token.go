package parser

type Token struct {
	Type  int
	Value string
	Pos   Position
	Ind   Indices
}

type Position struct {
	Line   int
	Column int
}

type Indices struct {
	Start int
	End   int
}

const (
	TOKEN_VAR = iota
	TOKEN_DATA
	TOKEN_RESOURCE
	TOKEN_CUSTOM_TYPE
	TOKEN_HARDCODED_STRING

	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET

	TOKEN_POINT
	TOKEN_ASSIGN
	TOKEN_QUOTE

	TOKEN_STRING

	TOKEN_NEWLINE
)

var Keywords = map[string]int{
	"resource": TOKEN_RESOURCE,
	"data":     TOKEN_DATA,
	"var":      TOKEN_VAR,
}

var Chars = map[string]int{
	"{": TOKEN_LEFT_BRACKET,
	"}": TOKEN_RIGHT_BRACKET,
	".": TOKEN_POINT,
	"=": TOKEN_ASSIGN,
	`"`: TOKEN_QUOTE,
}

var DebugTokens = map[int]string{
	TOKEN_RESOURCE:      "resource",
	TOKEN_DATA:          "data",
	TOKEN_VAR:           "var",
	TOKEN_LEFT_BRACKET:  "{",
	TOKEN_RIGHT_BRACKET: "}",
	TOKEN_POINT:         ".",
	TOKEN_ASSIGN:        "=",
	TOKEN_QUOTE:         `"`,
}
