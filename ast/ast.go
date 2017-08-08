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
