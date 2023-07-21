package lexer

import (
	"fmt"
	"os"
	"unicode"
)

type LexerError struct {
	Name     string
	Message  string
	Position TokenPos
}

func (err LexerError) display() {
	occured := fmt.Sprintf("Offset: %d, Line: %d, Col: %d", err.Position.Offset, err.Position.Line, err.Position.Column)
	fmt.Printf("LexicalError::%s\nMessage: %s\nLocation: %s\n", err.Name, err.Message, occured)
}

func NewLexerError(name string, message string) LexerError {
	return LexerError{Name: name, Message: message}
}

type lexer struct {
	contents string
	tokens   []Token
	errors   []LexerError
	pos      TokenPos
}

func new_lexer(contents string) *lexer {
	return &lexer{
		contents: contents,
		tokens:   make([]Token, 0),
		errors:   make([]LexerError, 0),
		pos: TokenPos{
			Line:   1,
			Offset: 0,
			Column: 0,
		},
	}
}

func (l *lexer) tk(kind TokenKind, value string) {
	l.tokens = append(l.tokens, token(kind, value, l.pos))
}

func (l *lexer) display_errors() {
	println("An error occured during lexical analysis!\n")
	for _, err := range l.errors {
		err.display()
	}
	println("")
}

func Tokenize(source string) []Token {
	lexer := new_lexer(source)

	for lexer.not_eof() {
		lexer.produce_token()
	}

	lexer.tk(EOF_TK, "eof")

	if len(lexer.errors) > 0 {
		lexer.display_errors()
		os.Exit(1)
	}

	return lexer.tokens
}

func (l *lexer) produce_token() {
	ch := l.current()
	switch ch {
	case '\n':
		l.newline()
	case ' ':
		l.advance(1)
	case '(':
		l.single_tk(LParen)
	case ')':
		l.single_tk(RParen)
	case '[':
		l.single_tk(LBracket)
	case ']':
		l.single_tk(RBracket)
	case '{':
		l.single_tk(LBrace)
	case '}':
		l.single_tk(RBrace)
	case ',':
		l.single_tk(Comma)
	case '.':
		l.single_tk(Dot)
	case '?':
		l.single_tk(Question)
	case '&':
		l.single_tk(Ampersand)
	case ';':
		l.single_tk(Semicolon)
	case ':':
		if l.peak() == ':' {
			l.multi_tk("::", Resolution)
		} else {
			l.single_tk(Colon)
		}
	case '=':
		if l.peak() == '=' {
			l.multi_tk("==", EqEq)
		} else {
			l.single_tk(Eq)
		}
	case '!':
		if l.peak() == '=' {
			l.multi_tk("!=", NotEq)
		} else {
			l.single_tk(Not)
		}
	case '<':
		if l.peak() == '=' {
			l.multi_tk("<=", LessEq)
		} else {
			l.single_tk(Less)
		}
	case '>':
		if l.peak() == '=' {
			l.multi_tk(">=", GreaterEq)
		} else {
			l.single_tk(Greater)
		}
	case '+':
		if l.peak() == '+' {
			l.multi_tk("++", PlusPlus)
		} else {
			l.single_tk(Plus)
		}
	case '-':
		if l.peak() == '-' {
			l.multi_tk("--", MinusMinus)
		} else {
			l.single_tk(Minus)
		}
	case '/':
		l.single_tk(Slash)
	case '*':
		l.single_tk(Star)
	case '%':
		l.single_tk(Percent)
	case '\'', '"':
		l.build_string()
	case '#':
		l.advance_to_newline()
	default:
		if unicode.IsDigit(rune(ch)) {
			l.build_number()
		} else {
			l.build_identifier()
		}
	}
}
