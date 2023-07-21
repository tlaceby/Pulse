package ast

type Stmt interface {
	debug(depth int)
}

type Expr interface{ Stmt }

// Statements
type Module struct {
	Stmt
	Name string
	Body []Stmt
}

type BlockStmt struct {
	Stmt
	Body []Stmt
}

// Expressions

type NumericExpr struct {
	Expr
	FloatingPoint bool
	Value         string
}

type StringExpr struct {
	Expr
	Value string
}

type SymbolExpr struct {
	Expr
	Name string
}
