package ast

type Expr interface {
	expr()
}

type Int struct {
	X int
}

func (x *Int) expr() {}

type Add struct {
	X, Y Expr
}

func (x *Add) expr() {}

type Sub struct {
	X, Y Expr
}

func (x *Sub) expr() {}

type Mul struct {
	X, Y Expr
}

func (x *Mul) expr() {}

type Div struct {
	X, Y Expr
}

func (x *Div) expr() {}

type Ident struct {
	Name string
}

func (x *Ident) expr() {}

type WithDecls struct {
	Decls []*Assign
	Expr  Expr
}

type Assign struct {
	LHS string
	RHS Expr
}
