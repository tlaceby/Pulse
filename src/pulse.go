package main

import (
	"github.com/tlaceby/Pulse/src/ast"
)

func main() {
	parser := ast.MakeParser()
	parser.ProduceAST("fn main () {\nlet x: int = 45;\n let foo: &int = &x;\n}", "test.pls")
}
