package lexer

import (
	"strings"
	"unicode"
)

func (l lexer) not_eof() bool {
	return l.pos.Offset < len(l.contents)
}

func (l lexer) current() byte {
	return l.contents[l.pos.Offset]
}

func (l lexer) peak() byte {
	return l.contents[l.pos.Offset+1]
}

func (l *lexer) advance(amount int) byte {
	if l.pos.Offset < len(l.contents) {
		l.pos.Column += 1
		l.pos.Offset += 1
		return l.contents[l.pos.Offset-1]
	}

	l.errors = append(l.errors, NewLexerError("Unexpected EOF", "EndOfFile Recieved Earlier Than Expected"))
	return 0
}

func (l *lexer) single_tk(kind TokenKind) {
	tk := token(kind, string(l.advance(1)), l.pos)
	l.tokens = append(l.tokens, tk)
}

func (l *lexer) multi_tk(value string, kind TokenKind) {
	tk := Token{kind, value, l.pos}
	l.advance(len(value))
	l.tokens = append(l.tokens, tk)
}

func (l *lexer) push_tk(value string, kind TokenKind) {
	tk := token(kind, value, l.pos)
	l.tokens = append(l.tokens, tk)
}

func (l *lexer) is_eof() bool {
	return !l.not_eof()
}

func (l *lexer) advance_to_newline() {
	for l.not_eof() && l.current() != '\n' {
		l.advance(1)
	}

	if l.is_eof() {
		panic("Unexpected EOF Located")
	}

	l.newline()
}

func (l *lexer) newline() {
	l.pos.Line++
	l.pos.Offset++
	l.pos.Column = 0
}

func (l *lexer) allowed_ident_char() bool {
	ch := l.current()
	return ch == '$' || ch == '_' || unicode.IsDigit(rune(ch)) || unicode.IsLetter(rune(ch))
}

func (l *lexer) build_number() {
	var num strings.Builder
	decimal := false

	for l.not_eof() && (l.current() == '.' || unicode.IsDigit(rune(l.current()))) {
		ch := l.advance(1)

		if ch == '.' {
			if decimal {
				panic("Number cannot contain multiple floating points")
			}

			decimal = true
		}

		num.WriteByte(ch)
	}

	l.push_tk(num.String(), Number)
}

func (l *lexer) build_string() {
	var str strings.Builder
	delim := l.advance(1)

	for l.not_eof() && l.current() != delim {
		str.WriteByte(l.advance(1))
	}

	if l.is_eof() {
		panic("Expected to find ending of string literal instead found EOF")
	}

	l.advance(1)
	l.push_tk(str.String(), String)
}

func (l *lexer) build_identifier() {
	var ident strings.Builder

	for l.not_eof() && l.allowed_ident_char() {
		ident.WriteByte(l.advance(1))
	}

	identifier := ident.String()
	kind, err := RESERVED_KEYWORDS[identifier]

	if err {
		l.push_tk(identifier, Symbol)
	} else {
		l.push_tk(ident.String(), kind)
	}
}
