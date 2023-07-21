package ast

import (
	"strings"

	"github.com/tlaceby/Pulse/src/lexer"
	"github.com/tlaceby/Pulse/src/pulse"
)

type Parser struct {
	tokens   []lexer.Token
	source   string
	previous lexer.Token
	position int
}

func MakeParser() *Parser {
	return &Parser{
		source:   "",
		tokens:   make([]lexer.Token, 0),
		previous: lexer.Token{},
		position: 0,
	}
}

func (p *Parser) ProduceAST(source string, filename string) Stmt {
	p.source = source
	p.tokens = lexer.Tokenize(source)

	module_name := strings.Split(filename, pulse.PULSE_EXTENSION)[0]

	return Module{Name: module_name}
}
