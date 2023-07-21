package ast

import (
	"fmt"

	"github.com/tlaceby/Pulse/src/helpers"
)

func (p Module) debug(depth int) {
	println("ModuleStmt: %s", p.Name)

	for _, stmt := range p.Body {
		stmt.debug(1)
	}
}

func (p BlockStmt) debug(depth int) {
	println("BlockStmt")

	for _, stmt := range p.Body {
		stmt.debug(1)
	}
}

// Primary Expressions

func (p NumericExpr) debug(depth int) {
	helpers.RepeatIndent(depth)
	fmt.Printf("Numeric: %s\n", p.Value)
}

func (p StringExpr) debug(depth int) {
	helpers.RepeatIndent(depth)
	fmt.Printf("String: %s\n", p.Value)
}

func (p SymbolExpr) debug(depth int) {
	helpers.RepeatIndent(depth)
	fmt.Printf("Symbol: %s\n", p.Name)
}
