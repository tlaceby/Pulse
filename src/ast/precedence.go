package ast

import "github.com/tlaceby/Pulse/src/lexer"

type Precedence int

const (
	Default Precedence = iota
	Assignment
	Logical
	Comparison
	Sum
	Product
	Unary
	CallMember
	Primary
)

type LED_FN = func(*Parser, Expr, Precedence)
type NUD_FN = func(*Parser)
type STMT_FN = func(*Parser)

type TokenMetaData struct {
	Precedence Precedence
	LedFN      LED_FN
	NudFN      NUD_FN
	StmtFN     STMT_FN
}

var tk_lookup map[lexer.TokenKind]TokenMetaData

func init_tables() {

}

func reg_literal(kind string) {

}

func reg_symbol(kind string, bp Precedence) {

}

func reg_infix(kind string, bp Precedence, led_fn LED_FN) {

}

func reg_postfix(kind string, bp Precedence, led_fn LED_FN) {

}

func reg_prefix(kind string, bp Precedence, led_fn NUD_FN) {

}

func reg_stmt(kind string, stmt_fn STMT_FN) {

}

func get_bp(kind lexer.TokenKind) Precedence {
	return Default
}
