package lexer

type TokenPos struct {
	Line   int
	Column int
	Offset int
}

type Token struct {
	Kind     TokenKind
	Value    string
	Position TokenPos
}

func token(kind TokenKind, value string, pos TokenPos) Token {
	return Token{
		Kind:     kind,
		Value:    value,
		Position: pos,
	}
}

type TokenKind int

const (
	EOF_TK TokenKind = iota
	String
	Number
	Symbol
	// Keywords
	Var
	Fn
	Module
	Using
	Struct
	For
	While
	If
	Else
	Elif
	In
	Typeof
	Return
	Break
	Impl
	// Operators
	Plus
	Minus
	Star
	Slash
	Ampersand
	Not
	NotEq
	Eq
	EqEq
	And
	Or
	Less
	LessEq
	Greater
	GreaterEq
	PlusPlus
	MinusMinus
	// Grouping
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket
	Comma
	Dot
	Colon
	Resolution
	Question
	Percent
	Semicolon
)

var RESERVED_KEYWORDS = map[string]TokenKind{
	"var":    Var,
	"module": Var,
	"fn":     Fn,
	"using":  Using,
	"struct": Struct,
	"for":    For,
	"while":  While,
	"if":     If,
	"else":   Else,
	"elif":   Elif,
	"in":     In,
	"typeof": Typeof,
	"return": Return,
	"break":  Break,
	"impl":   Impl,
	"and":    And,
	"or":     Or,
}
